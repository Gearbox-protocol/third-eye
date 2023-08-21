create table token_ltramp (
    block_num integer,
    credit_manager varchar(42),
    token varchar(42),
    lt_initial integer,
    lt_final integer,
    ramp_start integer,
    ramp_end integer,
    PRIMARY KEY (block_num, credit_manager, token)
);


create table quota_details {
    block_num integer,
    pool_quota_keeper varchar(42),
    token varchar(42),
    --
    pool varchar(42),
    cum_quota_index varchar(42),
    --
    limit varchar(80),
    rate integer,
    increase_fee integer,
    PRIMARY KEY (block_num, pool_quota_keeper, token)
};

create table account_quota_info {
    block_num integer,
    session_id varchar(100),
    token varchar(42),
    --
    quota_index varchar(80),
    quota varchar(80),
    fees varchar(80),
    interest varchar(80),
};