CREATE EXTENSION  IF NOT EXISTS hstore;
CREATE TABLE credit_managers (
    address varchar(42) PRIMARY KEY,
    pool_address varchar(42),
    underlying_token varchar(42),
    max_leverage integer,
    disabled boolean,
    is_weth boolean,
    min_amount character varying(80),
    max_amount character varying(80),
    borrow_rate_bi character varying(80),
    available_liquidity_bi character varying(80),
    total_borrowed_bi character varying(80),
    cumulative_borrowed_bi character varying(80),
    total_repaid_bi character varying(80),
    total_profit_bi character varying(80),
    total_losses_bi character varying(80),
    borrow_rate double precision,
    available_liquidity double precision,
    total_borrowed double precision,
    cumulative_borrowed double precision,
    total_repaid double precision,
    total_profit double precision,
    total_losses double precision,
    opened_accounts_count integer,
    total_opened_accounts integer,
    total_closed_accounts integer,
    total_repaid_accounts integer,
    total_liquidated_accounts integer,
    unique_users integer,
    _version integer,
    paused boolean
);

CREATE TABLE credit_manager_stats (
    id SERIAL PRIMARY KEY,
    credit_manager varchar(42),
    opened_accounts_count integer,
    total_opened_accounts integer,
    total_closed_accounts integer,
    total_repaid_accounts integer,
    total_liquidated_accounts integer,
    total_borrowed_bi varchar(80),
    cumulative_borrowed_bi varchar(80),
    total_repaid_bi varchar(80),
    total_profit_bi varchar(80),
    total_losses_bi varchar(80),
    block_num integer,
    total_borrowed double precision,
    cumulative_borrowed double precision,
    total_repaid double precision,
    total_profit double precision,
    total_losses double precision,
    unique_users integer,
    available_liquidity_bi varchar(80),
    available_liquidity double precision,
    borrow_rate double precision,
    borrow_rate_bi varchar(80)
);
ALTER TABLE ONLY credit_manager_stats
    ADD CONSTRAINT credit_manager_stats_block_num_fkey FOREIGN KEY (block_num) REFERENCES blocks(id) ON DELETE CASCADE;
ALTER TABLE ONLY credit_manager_stats
    ADD CONSTRAINT credit_manager_stats_credit_manager_fkey FOREIGN KEY (credit_manager) REFERENCES credit_managers(address);

CREATE TABLE allowed_protocols (
    id SERIAL PRIMARY KEY,
    protocol character varying(42),
    adapter character varying(42),
    block_num integer,
    credit_manager character varying(42)
);
