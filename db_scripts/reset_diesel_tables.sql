delete from diesel_balances; delete from diesel_transfers; delete from lm_rewards; 
update sync_adapters set details='{}'::jsonb, last_sync=(SELECT firstlog_at-1 FROM sync_adapters WHERE type='AddressProvider') WHERE type='LMRewardsv2';