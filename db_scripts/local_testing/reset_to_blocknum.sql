delete from account_operations where block_num>18246321;
delete from allowed_protocols where block_num>18246321;
delete from allowed_tokens where block_num>18246321;
delete from credit_manager_stats where block_num>18246321;
delete from credit_session_snapshots where block_num>18246321;
delete from dao_operations where block_num> 18246321;
delete from blocks where id > 18246321;


delete from debts where block_num> 18246321;
update debt_sync set last_calculated_at=18246321;
delete from token_oracle where  block_num> 18246321;
delete from transfer_account_allowed  where  block_num> 18246321;
delete from no_session_transfers where  block_num> 18246321;
delete from fast_check_params where  block_num> 18246321;
delete from parameters where  block_num> 18246321;
delete from pool_stats where  block_num> 18246321;
delete from pool_ledger where  block_num> 18246321;
delete from current_debts where block_num > 18246321;  
delete from liquidable_accounts where block_num> 18246321; -- doesn't matter if it is closed or open
delete from token_ltramp where block_num> 18246321; 
delete from quota_details where block_num> 18246321; 

-- a
update sync_adapters set last_sync = 18246321 where type in ('CreditConfigurator',  'AccountManager','CreditFilter', 'CreditManager', 'Pool', 'PoolKeeper') and last_sync > 18246321;
-- 
delete from credit_sessions where since> 18246321;
update  credit_sessions set closed_at=0,liquidator='', remaining_funds='0', close_transfers='{}',status=0 where closed_at> 18246321;

-- the borrower is not accounted for ; i.e. transfer of acocunt is not handled and should be used carefully
update credit_sessions cs set borrowed_amount=css.borrowed_amount_bi, balances=css.balances from  (select distinct on(session_id) * from credit_session_snapshots order by session_id, block_num desc) css where css.session_id=cs.id  and cs.status=0;

-- for creditmanager

update credit_managers cm set 
    opened_accounts_count=cms.opened_accounts_count,
    total_opened_accounts=cms.total_opened_accounts,
    total_repaid_accounts=cms.total_repaid_accounts,
    total_closed_accounts=cms.total_closed_accounts,
    total_liquidated_accounts=cms.total_liquidated_accounts,
    total_borrowed_bi=cms.total_borrowed_bi,
    cumulative_borrowed_bi=cms.cumulative_borrowed_bi,
    total_repaid_bi=cms.total_repaid_bi,
    total_losses_bi=cms.total_losses_bi,
    total_profit_bi=cms.total_profit_bi,
    total_borrowed=cms.total_borrowed,
    cumulative_borrowed=cms.cumulative_borrowed,
    total_repaid=cms.total_repaid,
    total_losses=cms.total_losses,
    total_profit=cms.total_profit,
    unique_users=cms.unique_users from (select distinct on (credit_manager) * from credit_manager_stats order by credit_manager, block_num desc) cms 
    where cm.address= cms.credit_manager;




delete from price_feeds where block_num >  18246321;
delete from token_oracle where block_num > 18246321;
update sync_adapters set last_sync=18246321 where type in ('PriceOracle', 'ChainlinkPriceFeed', 'CompositeChainlinkPF', 'QueryPriceFeed');
update sync_adapters set last_sync=18246321 where type in ('AddressProvider',  'ACL', 'AccountFactory');
delete from rebase_details where block_num > 18246321;



update sync_adapters set last_sync=18246321 where type in ('ContractRegister', 'ACL');
delete from debts where block_num > 18246321;
delete from current_debts where block_num > 18246321;
update debt_sync set last_calculated_at=18246321;


delete from closed_trading_sessions;
----

-- sync adapter for treasury is not updated.
-- tokens and token_current_price, schema_migrations not are needed to be updated.
-- faucet and operations are reduntant tables.
-- gear_balances , no need to udpate as GearToken syncadpater is not updated.
-- PooLMReward syncadapter is for table diesel_transfers, lm_rewards, diesel_balances.


-- ALTER DATABASE gearbox RENAME TO old_gearbox;
-- ALTER DATABASE sample RENAME TO gearbox;
-- createdb -O sample -T sample tmp_sample