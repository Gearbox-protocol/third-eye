CREATE TABLE account_operations (
    id SERIAL PRIMARY KEY,
    depth integer,
    tx_hash character varying(66),
    block_num integer,
    log_id integer,
    borrower character varying(42),
    session_id character varying(100),
    dapp character varying(42),
    action text,
    adapter_call boolean,
    args jsonb,
    transfers hstore
);