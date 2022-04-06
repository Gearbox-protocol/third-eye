-- price in usd and eth
alter table price_feeds add price_bi varchar(80),add  price DOUBLE PRECISION, add price_in_usd boolean;
update  price_feeds set price_bi=price_eth_bi, price=price_eth, price_in_usd='f';
alter table price_feeds drop price_eth_bi, drop price_eth;

select max(last_calculated_at) from debt_sync;
alter table debt_sync add field_set boolean;
delete from debt_sync;
insert into debt_sync(last_calculated_at, field_set) values(30859574,'t');
-- version 
alter table credit_sessions add version smallint;
update credit_sessions set version=1;
alter table  sync_adapters add version smallint;
update sync_adapters set version=1;
alter table token_oracle add version smallint;
update token_oracle set version = 1;

-- transfer allowed by the receiver
create table transfer_account_allowed(
    log_id integer,
    block_num integer,
    sender varchar(42),
    receiver varchar(42),
    allowed boolean,
    PRIMARY KEY(block_num, log_id)
);

