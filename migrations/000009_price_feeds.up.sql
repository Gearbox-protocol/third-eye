CREATE TABLE price_feeds (
    id SERIAL PRIMARY KEY,
    block_num integer,
    token character varying(42),
    feed character varying(42),
    price_bi varchar(80),
    price DOUBLE PRECISION, 
    price_in_usd boolean,
    round_id integer,
    uniswapv2_price double precision,
    uniswapv3_twap double precision,
    uniswapv3_price double precision,
    uni_price_fetch_block integer);

CREATE TABLE uniswap_pool_prices (
    id SERIAL PRIMARY KEY,
    uniswapv2_price double precision,
    uniswapv3_twap double precision,
    uniswapv3_price double precision,
    block_num integer,
    token character varying(42)
);

CREATE TABLE uniswap_chainlink_relations (
    block_num integer,
    chainlink_block_num integer,
    token character varying(42),
    feed character varying(42)
);


CREATE TABLE token_oracle (
    token character varying(42),
    oracle character varying(42),
    feed character varying(42),
    block_num integer NOT NULL,
    version smallint,
    PRIMARY KEY (block_num, token));

CREATE TABLE uniswap_pools (
    token character varying(42) PRIMARY KEY,
    pool_v2 character varying(42),
    pool_v3 character varying(42));