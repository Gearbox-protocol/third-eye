

insert into sync_adapters(address, type, last_sync) values('0x000000000000000000000000000000000000beef','PoolLMRewards',0);


CREATE TABLE diesel_balances (
    user varchar(42) ,
    diesel_sym varchar(42),
    balance varchar(80),
    PRIMARY KEY(diesel_sym, user));


CREATE TABLE lm_rewards (
    user varchar(42),
    pool varchar(42) 
    reward varchar(80),
    PRIMARY KEY(pool, user));
