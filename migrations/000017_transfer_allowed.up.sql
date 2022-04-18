create table transfer_account_allowed(
    log_id integer,
    block_num integer,
    sender varchar(42),
    receiver varchar(42),
    allowed boolean,
    PRIMARY KEY(block_num, log_id)
);