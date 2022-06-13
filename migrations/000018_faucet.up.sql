CREATE TABLE faucet (
    id varchar(69) PRIMARY KEY,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    address varchar(42) NOT NULL,
    token varchar(42) NOT NULL,
    next_update integer DEFAULT 0 NOT NULL,
    total double precision NOT NULL
);