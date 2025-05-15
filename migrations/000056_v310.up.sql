alter table token_oracle add price_oracle varchar(42), add disabled_at integer;
alter table token_oracle drop constraint token_oracle_pkey;

update token_oracle set price_oracle= (select address from sync_adapters where type='PriceOracle' and version=1) where version=1;
update token_oracle set price_oracle= (select address from sync_adapters where type='PriceOracle' and version=2) where version=2;
-- UPDATE
update token_oracle set price_oracle= (select address from sync_adapters where type='PriceOracle' and version=300) where version=300;
alter table token_oracle add PRIMARY KEY (price_oracle, block_num, token, reserve);

--
alter table token_current_price add price_oracle varchar(42);
alter table token_current_price drop constraint token_current_price_pkey;
update token_current_price set price_oracle= (select address from sync_adapters where type='PriceOracle' and version=300) where price_source in ('chainlink', 'gearbox');
update token_current_price set price_oracle= '0x0000000000000000000000000000000000000000' where price_source='spot';
alter table token_current_price add PRIMARY KEY (price_oracle, price_source, token);

--
alter table price_feeds drop column merged_pf_version, drop column token; 

alter table pools add market varchar(42), add price_oracle varchar(42);

-- alter table price_feeds add PRIMARY KEY (block_num, feed);

CREATE TABLE relations (
    block_num integer,
    owner varchar(42),
    dependent varchar(42),
    category character varying(10),
    id SERIAL PRIMARY KEY
);

-- UPDATE
update pools p set price_oracle=sa.address from  sync_adapters sa where type='PriceOracle' and (case when p._version=1 then 2 else p._version end)=sa.version;

update token_oracle set disabled_at=19752044 where version=2; -- don't disable for v1
update token_oracle set disabled_at=13856183 where feed='0xc170DC3C2e8809AC6197D56b86bF421c8a7f8c67'; -- all for v1
update token_oracle set disabled_at=18577104 where feed='0x172971182351e00C2D700bA1e8c5586Ad2CFa38c';
update token_oracle set disabled_at=18577104 where feed='0x614f9486Ab9C7a217526c097656D2F6bD2DB631C';
update token_oracle set disabled_at=14769098 where feed='0x1a8AC67A1B64F7fd71bB91c21581f036AbE6AEc2';
update token_oracle set disabled_at=14956928 where feed='0x91401cedCBFd9680cE193A5F54E716504233e998'; -- all for v1



-- insert into relations(block_num, owner, dependent, category) (select discovered_at, pool,oracle , 'PoolOracle'  from (select address pool from pools where _version=1) l join (select address oracle, discovered_at from sync_adapters where type='PriceOracle' and version=2) on true);
-- insert into relations(block_num, owner, dependent, category)  (select discovered_at, pool,oracle , 'PoolOracle'  from (select address pool from pools where _version=1) l join (select address oracle, discovered_at from sync_adapters where type='PriceOracle' and version=1) on true);
insert into relations(block_num, owner, dependent, category)  (select discovered_at, pool,oracle , 'PoolOracle'  from (select address pool from pools where _version=300) l join (select address oracle, discovered_at from sync_adapters where type='PriceOracle' and version=300) on true);

alter table tvl_snapshots add market varchar(42);
update tvl_snapshots set market='0x0000000000000000000000000000000000000000';
alter table tvl_snapshots drop constraint tvl_snapshots_pkey;
alter table tvl_snapshots add PRIMARY KEY (market, block_num);


alter table price_feeds add PRIMARY KEY (block_num, feed);

-- do after the legacy market configurator address is found.
-- update legacy address in marketconfigurator.
-- update tvl_snapshots market=? where market='0x0000000000000000000000000000000000000000';

insert into debt_sync values ('t', 0 ,0 );