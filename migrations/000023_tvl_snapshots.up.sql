create table tvl_snapshots (
    block_num integer PRIMARY KEY,
    available_liquidity DOUBLE PRECISION,
    ca_total_value DOUBLE PRECISION);