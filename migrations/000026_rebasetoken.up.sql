create table rebase_details(
    total_eth varchar(80),
    total_shares varchar(80),
    block_num integer PRIMARY KEY
);

INSERT INTO sync_adapters(address,discovered_at,last_sync,firstlog_at,type ) 
    SELECT * FROM (SELECT '0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84', 17266004,17266004,17266004, 'RebaseToken') as tmp WHERE 
    EXISTS (SELECT 1 FROM sync_Adapters where address='0xcF64698AFF7E5f27A11dff868AF228653ba53be0'); -- if mainnet addressProvier exist then add steth rebase token
