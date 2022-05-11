-- price in usd and eth
alter table price_feeds add price_bi varchar(80),add  price DOUBLE PRECISION, add price_in_usd boolean;
update  price_feeds set price_bi=price_eth_bi, price=price_eth, price_in_usd='f';
alter table price_feeds drop price_eth_bi, drop price_eth;

select max(last_calculated_at) from debt_sync;
alter table debt_sync add field_set boolean;
delete from debt_sync;
insert into debt_sync(last_calculated_at, field_set) values(30859574,'t');
alter table debt_sync add primary key(field_set);
-- version 
alter table credit_sessions add version smallint;
update credit_sessions set version=1;
alter table  sync_adapters add version smallint;
update sync_adapters set version=1;
alter table token_oracle add version smallint;
update token_oracle set version = 1;

-- 
alter table account_operations add main_action integer ;
alter table account_operations add constraint main_action_c FOREIGN KEY(main_action) references account_operations(id);

-- transfer allowed by the receiver
create table transfer_account_allowed(
    log_id integer,
    block_num integer,
    sender varchar(42),
    receiver varchar(42),
    allowed boolean,
    PRIMARY KEY(block_num, log_id)
);

-- update sync_adapters set details='{"priceOracles":[]}' where type='AddressProvider';


--
-- new v2 updates
update sync_adapters set type='QueryPriceFeed' where type='YearnPriceFeed';
select * from sync_adapters  where type='QueryPriceFeed';
update sync_adapters set details=replace(details::TEXT,'"token"','"pfType":"YearnPF", "token"')::jsonb    WHERE type='QueryPriceFeed';


-- alter table current_debts add repay_amount DOUBLE PRECISION, add    repay_amount_bi varchar(80);
-- update table current_debts set repay_amount=amount_to_pool, repay_amount_bi=amount_to_pool_bi ;
-- alter table current_debts drop column amount_to_pool, drop column amount_to_pool_bi;