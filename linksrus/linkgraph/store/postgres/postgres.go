package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"deploying_monolith_linksrus/linksrus/linkgraph/graph"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/xerrors"
)

var (
	upsertLinkQuery = `
INSERT INTO links (url, retrieved_at) VALUES ($1, $2) 
ON CONFLICT (url) DO UPDATE SET retrieved_at=GREATEST(links.retrieved_at, $2)
RETURNING id, retrieved_at
`
	findLinkQuery         = "SELECT url, retrieved_at FROM links WHERE id=$1"
	linksInPartitionQuery = "SELECT id, url, retrieved_at FROM links WHERE id >= $1 AND id < $2 AND retrieved_at < $3"

	upsertEdgeQuery = `
INSERT INTO edges (src, dst, updated_at) VALUES ($1, $2, $3)
ON CONFLICT (src,dst) DO UPDATE SET updated_at=$3
RETURNING id, updated_at
`
	edgesInPartitionQuery = "SELECT id, src, dst, updated_at FROM edges WHERE src >= $1 AND src < $2 AND updated_at < $3"
	removeStaleEdgesQuery = "DELETE FROM edges WHERE src=$1 AND updated_at < $2"

	// Compile-time check for ensuring PostgresGraph implements Graph.
	_ graph.Graph = (*PostgresGraph)(nil)
)

// PostgresGraph implements a graph that persists its links and edges to a
// Postgres instance.
type PostgresGraph struct {
	db *sql.DB
}

type PostgresConArgs struct {
	host string
	port int32
	user string
	password string
	dbname string
	sslmode string
}

func (pConArgs *PostgresConArgs) CreateString() (res string){
	res = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pConArgs.host, pConArgs.port, pConArgs.user, pConArgs.password, pConArgs.dbname)

	return
}

// NewPostgresGraph returns a PostgresGraph instance that connects to the Postgres
// instance specified by dsn.
func NewPostgresGraph(cs string) (*PostgresGraph, error) {
	db, err := sql.Open("postgres", cs)
	if err != nil {
		return nil, err
	}

	return &PostgresGraph{db: db}, nil
}

// Close terminates the connection to the backing Postgres instance.
func (c *PostgresGraph) Close() error {
	return c.db.Close()
}

// UpsertLink creates a new link or updates an existing link.
func (c *PostgresGraph) UpsertLink(link *graph.Link) error {
	row := c.db.QueryRow(upsertLinkQuery, link.URL, link.RetrievedAt.UTC())
	if err := row.Scan(&link.ID, &link.RetrievedAt); err != nil {
		return xerrors.Errorf("upsert link: %w", err)
	}

	link.RetrievedAt = link.RetrievedAt.UTC()
	return nil
}

// FindLink looks up a link by its ID.
func (c *PostgresGraph) FindLink(id uuid.UUID) (*graph.Link, error) {
	row := c.db.QueryRow(findLinkQuery, id)
	link := &graph.Link{ID: id}
	if err := row.Scan(&link.URL, &link.RetrievedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, xerrors.Errorf("find link: %w", graph.ErrNotFound)
		}

		return nil, xerrors.Errorf("find link: %w", err)
	}

	link.RetrievedAt = link.RetrievedAt.UTC()
	return link, nil
}

// Links returns an iterator for the set of links whose IDs belong to the
// [fromID, toID) range and were last accessed before the provided value.
func (c *PostgresGraph) Links(fromID, toID uuid.UUID, accessedBefore time.Time) (graph.LinkIterator, error) {
	rows, err := c.db.Query(linksInPartitionQuery, fromID, toID, accessedBefore.UTC())
	if err != nil {
		return nil, xerrors.Errorf("links: %w", err)
	}

	return &linkIterator{rows: rows}, nil
}

// UpsertEdge creates a new edge or updates an existing edge.
func (c *PostgresGraph) UpsertEdge(edge *graph.Edge) error {
	row := c.db.QueryRow(upsertEdgeQuery, edge.Src, edge.Dst, time.Now().UTC())
	if err := row.Scan(&edge.ID, &edge.UpdatedAt); err != nil {
		if isForeignKeyViolationError(err) {
			err = graph.ErrUnknownEdgeLinks
		}
		return xerrors.Errorf("upsert edge: %w", err)
	}

	return nil
}

// Edges returns an iterator for the set of edges whose source vertex IDs
// belong to the [fromID, toID) range and were last updated before the provided
// value.
func (c *PostgresGraph) Edges(fromID, toID uuid.UUID, updatedBefore time.Time) (graph.EdgeIterator, error) {
	updatedBefore = updatedBefore.UTC()
	rows, err := c.db.Query(edgesInPartitionQuery, fromID, toID, updatedBefore)
	if err != nil {
		return nil, xerrors.Errorf("edges: %w", err)
	}

	return &edgeIterator{rows: rows}, nil
}

// RemoveStaleEdges removes any edge that originates from the specified link ID
// and was updated before the specified timestamp.
func (c *PostgresGraph) RemoveStaleEdges(fromID uuid.UUID, updatedBefore time.Time) error {
	_, err := c.db.Exec(removeStaleEdgesQuery, fromID, updatedBefore.UTC())
	if err != nil {
		return xerrors.Errorf("remove stale edges: %w", err)
	}

	return nil
}

// isForeignKeyViolationError returns true if err indicates a foreign key
// constraint violation.
func isForeignKeyViolationError(err error) bool {
	pqErr, valid := err.(*pq.Error)
	if !valid {
		return false
	}

	return pqErr.Code.Name() == "foreign_key_violation"
}
