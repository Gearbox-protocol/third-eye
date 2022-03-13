CREATE TABLE debt_sync (
    last_calculated_at integer,
    field_set boolean,
);

CREATE TABLE debts (
    id SERIAL NOT NULL,
    block_num integer,
    session_id character varying(100),
    cal_health_factor varchar(80),
    cal_threshold_value character varying(80),
    cal_borrowed_amt_with_interest character varying(80),
    cal_total_value character varying(80),
    profit_underlying DOUBLE PRECISION,
    profit_usd DOUBLE PRECISION,
    total_value_usd DOUBLE PRECISION,
    collateral_usd DOUBLE PRECISION,
    collateral_underlying DOUBLE PRECISION
);
create index debts_session_id_index on debts using BTREE (session_id,block_num);