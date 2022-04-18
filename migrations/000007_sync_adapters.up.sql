CREATE TABLE sync_adapters (
    address character varying(42) PRIMARY KEY,
    discovered_at integer,
    last_sync integer,
    firstlog_at integer,
    details jsonb,
    type character varying(40),
    error character varying(200),
    version smallint,
    disabled boolean
);