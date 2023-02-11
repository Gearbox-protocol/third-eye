

insert into sync_adapters(address, type, last_sync, disabled) values('0x000000000000000000000000000000000000beef','PoolLMRewards',(select discovered_at from sync_Adapters where type='AddressProvider') , 'f');


CREATE TABLE diesel_balances (
    user_address varchar(42) ,
    diesel_sym varchar(42),
    balance_bi varchar(80),
    balance double precision,
    PRIMARY KEY(diesel_sym, user_address)
);


CREATE TABLE lm_rewards (
    user_address varchar(42),
    pool varchar(42),
    reward varchar(80),
    PRIMARY KEY(pool, user_address));

create index pool_ledger_user on pool_ledger using BTREE (user);