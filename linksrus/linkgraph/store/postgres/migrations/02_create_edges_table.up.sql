CREATE TABLE IF NOT EXISTS edges (
	id UUID PRIMARY KEY default gen_random_uuid(),
	src UUID NOT NULL REFERENCES links(id) ON DELETE CASCADE,
	dst UUID NOT NULL REFERENCES links(id) ON DELETE CASCADE,
	updated_at TIMESTAMP with time zone,
	CONSTRAINT edge_links UNIQUE(src,dst)
);
