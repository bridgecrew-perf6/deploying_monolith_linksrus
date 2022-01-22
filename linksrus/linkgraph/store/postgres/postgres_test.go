package postgres

import (
	"database/sql"
	"testing"

	"deploying_monolith_linksrus/linksrus/linkgraph/graph/graphtest"
	gc "gopkg.in/check.v1"
)

var _ = gc.Suite(new(PostgresGraphTestSuite))

func Test(t *testing.T) { gc.TestingT(t) }

type PostgresGraphTestSuite struct {
	graphtest.SuiteBase
	db *sql.DB
}

func (s *PostgresGraphTestSuite) SetUpSuite(c *gc.C) {
	psqlConArgs := PostgresConArgs{"localhost", 5432, "admin", "secret", "postgres", "false" }

	g, err := NewPostgresGraph(psqlConArgs.CreateString())
	c.Assert(err, gc.IsNil)
	s.SetGraph(g)
	s.db = g.db
}

func (s *PostgresGraphTestSuite) SetUpTest(c *gc.C) {
	s.flushDB(c)
}

func (s *PostgresGraphTestSuite) TearDownSuite(c *gc.C) {
	if s.db != nil {
		s.flushDB(c)
		c.Assert(s.db.Close(), gc.IsNil)
	}
}

func (s *PostgresGraphTestSuite) flushDB(c *gc.C) {
	_, err := s.db.Exec("DELETE FROM links")
	c.Assert(err, gc.IsNil)
	_, err = s.db.Exec("DELETE FROM edges")
	c.Assert(err, gc.IsNil)
}
