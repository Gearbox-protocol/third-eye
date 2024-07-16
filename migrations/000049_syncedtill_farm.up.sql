alter table farm_v3 add synced_till integer;
update farm_v3 set synced_till=(select last_sync from sync_adapters where type='LMRewardsv3');