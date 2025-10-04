create table diesel_sync (
    pool varchar(42) PRIMARY KEY,
    pool_synced_till integer
);

insert into diesel_sync (select distinct on (pool) pool,   pool_synced_till from farm_V3 order by pool);