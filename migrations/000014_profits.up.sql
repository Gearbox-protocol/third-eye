alter table credit_session_snapshots add collateral_in_usd varchar(80);
alter table credit_managers drop column allowed_contracts, drop column score;

alter table current_debts drop column liq_amount, drop column loss, drop column repay_amount;
alter table current_debts add amount_to_pool DOUBLE PRECISION,add amount_to_pool_bi varchar(80);
alter table current_debts add  collateral_usd_bi varchar(80), add profit_usd_bi varchar(80);
alter table debts add  collateral_usd_bi varchar(80), add profit_usd_bi varchar(80), add total_value_usd_bi varchar(80);
alter table debts drop column health_factor, drop column total_value, drop column borrowed_amt_with_interest;