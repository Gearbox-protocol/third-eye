create table farm_v3 (
    pool varchar(42),
    farm varchar(42),
    --
    checkpoint integer,
    farmed_per_token varchar(80),
    --
    reward varchar(80),
    period integer,
    end_ts integer,
    --  
    total_supply varchar(80),
    diesel_token varchar(42),
    PRIMARY KEY (farm)
);

create table user_lmdetails_v3 (
    balances_bi varchar(80),
    balances double precision,
    correction varchar(80),
    account varchar(42),
    farm varchar(42),
    diesel_sym varchar(42),
    PRIMARY KEY(farm, account)
);

-- 1inch lm rewards are stored with gear.

update sync_adapters set type='LMRewardsv2', address='0x00000000000000000000000000000000000beef2' where type = 'PoolLMRewards';
--
-- WITH data as (select * from sync_adapters where type = 'AddressProvider')
-- insert into sync_adapters(address, type, last_sync, firstlog_at, version, discovered_at, disabled) values (
--     '0x00000000000000000000000000000000000beef3', 
--     'LMRewardsv3', 
--     (SELECT firstlog_at-1 from data),
--     (SELECT firstlog_at from data),
--     '300',
--     (SELECT firstlog_at from data),
--     'f'
--     );

