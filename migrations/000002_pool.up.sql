CREATE TABLE pools (
    address varchar(42) PRIMARY KEY,
    underlying_token varchar(42),
    diesel_token varchar(42),
    is_weth boolean,
    expected_liq_limit varchar(80),
    withdraw_fee varchar(80)
);

CREATE TABLE pool_stats (
    id SERIAL PRIMARY KEY,
    pool varchar(42),
    block_num integer,
    unique_users integer,
    total_borrowed double precision,
    total_borrowed_bi varchar(80),
    total_profit double precision,
    total_profit_bi varchar(80),
    total_losses double precision,
    total_losses_bi varchar(80),
    deposit_apy double precision,
    deposit_apy_bi varchar(80),
    borrow_apy double precision,
    borrow_apy_bi varchar(80),
    expected_liquidity double precision,
    expected_liquidity_bi varchar(80),
    expected_liquidity_limit_bi varchar(80),
    available_liquidity double precision,
    available_liquidity_bi varchar(80),
    withdraw_fee integer,
    diesel_rate_bi varchar(80),
    diesel_rate double precision,
    cumulative_index_ray character varying(80)
);

ALTER TABLE ONLY pool_stats
    ADD CONSTRAINT pool_stats_block_num_fkey FOREIGN KEY (block_num) REFERENCES blocks(id) ON DELETE CASCADE;
ALTER TABLE ONLY pool_stats
    ADD CONSTRAINT pool_stats_pool_fkey FOREIGN KEY (pool) REFERENCES pools(address);


CREATE TABLE pool_ledger (
    id SERIAL PRIMARY KEY,
    block_num integer,
    log_id integer,
    pool character varying(42),
    address character varying(42),
    event character varying(50),
    liquidity character varying(80)
);
