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
    fee_interest integer,
    fee_liquidation integer,
    liq_discount integer,
    liq_discount_expired  integer,
    fee_liquidation_expired  integer,
    PRIMARY KEY (block_num, credit_manager));

ALTER TABLE ONLY parameters
    ADD CONSTRAINT parameters_block_num_fkey FOREIGN KEY (block_num) REFERENCES blocks(id) ON DELETE CASCADE;

-- alter table parameters add liq_discount_expired integer, add fee_liquidation_expired integer;
-- alter table parameters  add fee_interest_x integer,add fee_liquidation_x integer,add liq_discount_x integer;
-- update parameters set fee_interest_x=cast(fee_interest as int),fee_liquidation_x=cast(fee_liquidation as int),liq_discount_x=cast(liq_discount as int);

-- alter table parameters drop column fee_interest, drop column fee_liquidation,drop column liq_discount;
-- alter table parameters add fee_interest integer, add fee_liquidation integer,add liq_discount integer;
-- update parameters set fee_interest=fee_interest_x,fee_liquidation=fee_liquidation_x,liq_discount=liq_discount_x;

-- alter table parameters drop column fee_interest_x, drop column fee_liquidation_x,drop column liq_discount_x;
