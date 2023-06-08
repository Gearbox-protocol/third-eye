CREATE TABLE price_feeds (
    id SERIAL PRIMARY KEY,
    block_num integer,
    token character varying(42),
    feed character varying(42),
    price_bi varchar(80),
    price DOUBLE PRECISION, 
    price_in_usd boolean,
    round_id integer);


CREATE TABLE token_oracle (
    token character varying(42),
    oracle character varying(42),
    feed character varying(42),
    block_num integer NOT NULL,
    version smallint,
    feed_type varchar(25),
    PRIMARY KEY (block_num, token));


CREATE TABLE token_current_price (
    token varchar(42) PRIMARY KEY,
    price DOUBLE PRECISION,
    price_bi varchar(80),
    block_num integer);

alter table token_current_price add column price_source varchar(10);
update token_current_price set price_source='chainlink';
alter table token_current_price drop  constraint  token_current_price_pkey;
alter table token_current_price add PRIMARY KEY (price_source, token);

-- insert into token_current_price(token, price, block_num) select distinct on (token)  token, price, block_num from price_feeds order by token, block_num DESC;