alter table price_feeds add price_bi varchar(80),add  price DOUBLE PRECISION, add price_in_usd boolean;
update  price_feeds set price_bi=price_eth_bi, price=price_eth, price_in_usd='f';
alter table price_feeds drop price_eth_bi, drop price_eth;

alter table credit_sessions add version smallint;
alter table debt_sync add field_set boolean;
update credit_sessions set version=1;