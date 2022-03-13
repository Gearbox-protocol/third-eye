alter table price_feeds add price_bi varchar(80),add  price DOUBLE PRECISION, add price_in_eth boolean;
update  price_feeds set price_bi=price_eth_bi, price=price_eth, price_in_eth='t';
alter table price_feeds drop price_eth_bi, drop price_eth;