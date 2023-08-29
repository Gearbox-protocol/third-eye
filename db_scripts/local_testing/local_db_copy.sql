
--------
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


CREATE TABLE current_debts (
    session_id varchar(80) PRIMARY KEY,
    block_num integer,
    cal_health_factor varchar(80),
    cal_threshold_value DOUBLE PRECISION,
    cal_threshold_value_bi varchar(80),
    cal_borrowed_amt_with_interest DOUBLE PRECISION,
    cal_borrowed_amt_with_interest_bi varchar(80),
    cal_total_value DOUBLE PRECISION,
    cal_total_value_bi varchar(80),
    profit_usd DOUBLE PRECISION,
    profit_underlying DOUBLE PRECISION,
    collateral_usd DOUBLE PRECISION,
    collateral_underlying DOUBLE PRECISION,
    repay_amount DOUBLE PRECISION,
    repay_amount_bi varchar(80));

ALTER TABLE ONLY current_debts
    ADD CONSTRAINT current_debts_block_num_fkey FOREIGN KEY (block_num) REFERENCES blocks(id) ON DELETE CASCADE;