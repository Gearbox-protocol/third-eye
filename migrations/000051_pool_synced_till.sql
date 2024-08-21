alter table farm_v3 add pool_synced_till integer not null default 0;
update farm_v3 set pool_synced_till = synced_till;