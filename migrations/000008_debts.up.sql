CREATE TABLE debt_sync (
    last_calculated_at integer
);

CREATE TABLE debts (
    id SERIAL PRIMARY KEY,
    block_num integer,
    session_id character varying(100),
    health_factor varchar(80),
    cal_health_factor varchar(80),
    cal_threshold_value character varying(80),
    borrowed_amt_with_interest character varying(80),
    cal_borrowed_amt_with_interest character varying(80),
    cal_total_value character varying(80),
    total_value character varying(80),
    profit_in_underlying DOUBLE PRECISION,
     collateral_in_underlying DOUBLE PRECISION
);
create index debts_session_id_index on debts using BTREE (session_id,block_num);