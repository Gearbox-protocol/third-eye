update sync_adapters set last_sync = firstlog_at - 1 where type not in ('YearnPriceFeed', 'ChainlinkPriceFeed', 'AddressProvider', 'PriceOracle');
delete from pool_ledger;
delete from account_operations;
delete from allowed_protocols;
delete from allowed_tokens;
delete from credit_manager_stats;
delete from credit_session_snapshots;
delete from parameters;
delete from credit_sessions;
insert into credit_managers(address, pool_address,underlying_token, is_weth) 
 select right(address, length(address) -1), pool_address, underlying_token , is_weth from credit_managers;
 delete from credit_managers where address like '0x%';
update credit_managers set address = concat('0',address);
update pools set expected_liq_limit='0', withdraw_fee='0';
delete from pool_stats;
delete from dao_operations;delete from fast_check_params;delete from gear_balances;delete from treasury_snapshots;delete from treasury_transfers;
delete from debts; delete from liquidable_accounts; delete from profiles;delete from current_debts; delete from debt_sync;