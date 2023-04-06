create table tvl_snapshots (
    block_num integer PRIMARY KEY,
    available_liquidity DOUBLE PRECISION,
    ca_total_value DOUBLE PRECISION);

alter table token_current_price add column price_source varchar(10);
update token_current_price set price_source='chainlink';
alter table token_current_price drop  constraint  token_current_price_pkey;
alter table token_current_price add PRIMARY KEY (price_source, token);