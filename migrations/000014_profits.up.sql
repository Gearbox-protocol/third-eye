alter table credit_session_snapshots add collateral_in_usd varchar(80);
alter table credit_managers drop column allowed_contracts;

alter table current_debts drop column liq_amount, drop column loss, drop column repay_amount;
alter table current_debts add amount_to_pool DOUBLE PRECISION,add amount_to_pool_bi varchar(80);
alter table current_debts add  collateral_usd DOUBLE PRECISION, add profit_usd DOUBLE PRECISION;
alter table debts add  collateral_usd DOUBLE PRECISION, add profit_usd DOUBLE PRECISION, add total_value_usd DOUBLE PRECISION;
alter table debts drop column health_factor, drop column total_value, drop column borrowed_amt_with_interest;

-- alter table current_debts add collateral_usd DOUBLE PRECISION, add profit_usd DOUBLE PRECISION;
-- update current_debts set collateral_usd=collateral_usd_bi::float4/100000000, 
-- profit_usd=profit_usd_bi::float4/100000000;
-- alter table current_debts drop column  collateral_usd_bi, drop column profit_usd_bi;
-- alter table debts add collateral_usd DOUBLE PRECISION, add total_value_usd DOUBLE PRECISION, add profit_usd DOUBLE PRECISION;
-- update debts set collateral_usd=collateral_usd_bi::float4/100000000, 
-- total_value_usd=total_value_usd_bi::float4/100000000,
-- profit_usd=profit_usd_bi::float4/100000000;
-- alter table debts drop column  collateral_usd_bi, drop column total_value_usd_bi, drop column profit_usd_bi;
