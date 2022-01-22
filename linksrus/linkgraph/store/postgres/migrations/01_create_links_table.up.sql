CREATE TABLE IF NOT EXISTS links (
	id UUID PRIMARY KEY default gen_random_uuid(),
	url varchar(100) UNIQUE,
	retrieved_at TIMESTAMP with time zone
);
