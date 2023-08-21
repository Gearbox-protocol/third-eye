alter table pool_ledger add shares integer, shares_bi varchar(80), executor varchar(42), receiver varchar(42);

create table quota_details (
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
);

create table account_quota_info (
    block_num integer,
    session_id varchar(100),
    token varchar(42),
    pool_quota_keeper varchar(42),
    --
    quota_index varchar(80),
    quota varchar(80),
    fees varchar(80),
    interest varchar(80),
    PRIMARY KEY (block_num, session_id, token)
);