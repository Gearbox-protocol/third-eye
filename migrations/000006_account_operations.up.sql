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
alter table account_operations add main_action integer ;
alter table account_operations add constraint main_action_c FOREIGN KEY(main_action) references account_operations(id);
