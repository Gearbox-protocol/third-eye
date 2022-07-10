create table liquidable_accounts (
    session_id varchar(100) PRIMARY KEY,
    notified_if_liquidable boolean,
    block_num integer);

create table profiles (profile text);

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

create TABLE parameters (
    block_num integer,
    credit_manager varchar(42),
    min_amount varchar(42),
    max_amount varchar(42),
    max_leverage varchar(42),
    fee_interest varchar(42),
    fee_liquidation varchar(42),
    liq_discount varchar(42),
    PRIMARY KEY (block_num, credit_manager));

ALTER TABLE ONLY parameters
    ADD CONSTRAINT parameters_block_num_fkey FOREIGN KEY (block_num) REFERENCES blocks(id) ON DELETE CASCADE;