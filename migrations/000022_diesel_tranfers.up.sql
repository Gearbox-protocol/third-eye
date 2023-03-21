create table diesel_transfers (
    block_num integer,
    from_user varchar(42),
    to_user varchar(42),
    token_sym varchar(10),
    amount DOUBLE PRECISION,
    log_id integer,
    PRIMARY KEY (block_num, log_id)
);

ALTER TABLE pools add deposit_apy_bi varchar(80);