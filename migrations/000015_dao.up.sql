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
    user_address varchar(42) PRIMARY KEY);

create table treasury_snapshots (
    date_str varchar(20),
    block_num integer PRIMARY KEY,
    prices_in_usd jsonb,
    balances jsonb,
    value_in_usd DOUBLE PRECISION);

create table treasury_transfers (
    amount varchar(80),
    token varchar(42),
    log_id integer,
    block_num integer,
    PRIMARY KEY (block_num, log_id));

create table no_session_transfers (
    amount varchar(80),
    token varchar(42),
    from varchar(42),
    to varchar(42),
    block_num integer,
    log_id integer,
    tx_hash varchar(66));