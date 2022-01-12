create table dao_operations (
    log_id   integer,
    tx_hash varchar(66),
    block_num integer,
    contract varchar(42),
    type integer,
    args jsonb,
    PRIMARY KEY (log_id, block_num));

create table fast_check_params (
    block_num integer,
    credit_manager varchar(42),
    chi_threshold varchar(80),
    hf_checkinterval varchar(80),
    PRIMARY KEY (block_num, credit_manager));

create table gear_balances (
    balance varchar(80),
    user varchar(42) PRIMARY KEY);