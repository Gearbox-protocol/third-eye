CREATE TABLE price_feeds (
    id SERIAL PRIMARY KEY,
    block_num integer,
    token character varying(42),
    feed character varying(42),
    price_eth_bi character varying(80),
    round_id integer,
    price_eth double precision,
    uniswapv2_price double precision,
    uniswapv3_twap double precision,
    uniswapv3_price double precision,
    uni_price_fetch_block double precision,
);


CREATE TABLE token_oracle (
    token character varying(42),
    oracle character varying(42),
    feed character varying(42),
    block_num integer NOT NULL,
    PRIMARY KEY (block_num, token));