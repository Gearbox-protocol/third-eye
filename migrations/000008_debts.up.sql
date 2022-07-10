CREATE TABLE debt_sync (
    last_calculated_at integer,
    field_set boolean PRIMARY KEY
);

CREATE TABLE debts (
    block_num integer,
    cal_health_factor varchar(80),
    cal_threshold_value_bi character varying(80),
    cal_borrowed_amt_with_interest_bi character varying(80),
    cal_total_value_bi character varying(80),
    profit_underlying DOUBLE PRECISION,
    profit_usd DOUBLE PRECISION,
    collateral_usd DOUBLE PRECISION,
    collateral_underlying DOUBLE PRECISION,
    id SERIAL NOT NULL,
    session_id character varying(100),
    total_value_usd DOUBLE PRECISION
);
create index debts_block_num on debts using BTREE (block_num);
create index debts_session_id on debts using BTREE (session_id);