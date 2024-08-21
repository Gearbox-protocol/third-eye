update  farm_v3 set pool_synced_till=0 ;
delete from diesel_balances ;
update sync_adapters set last_sync = (select min(discovered_at) from sync_adapters where type='Pool' and version=300 ) where type = 'LMRewardsv3';