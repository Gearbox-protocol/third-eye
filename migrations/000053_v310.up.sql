alter table token_oracle add price_oracle varchar(42), add disabled_at integer;
alter table token_oracle drop constraint token_oracle_pkey;

update token_oracle set price_oracle= (select address from sync_adapters where type='PriceOracle' and version=1) where version=1;
update token_oracle set price_oracle= (select address from sync_adapters where type='PriceOracle' and version=2) where version=2;
update token_oracle set price_oracle= (select address from sync_adapters where type='PriceOracle' and version=300) where version=300;

alter table token_oracle add PRIMARY KEY (price_oracle, block_num, token, reserve);
--
alter table price_feeds drop column merged_pf_version, drop column token; 

alter table pools add market varchar(42), add price_oracle varchar(42);

alter table price_feeds add PRIMARY KEY (block_num, feed);

CREATE TABLE relations (
    block_num integer,
    owner varchar(42),
    dependent varchar(42),
    category character varying(10),
    id SERIAL PRIMARY KEY
);

update pools p set price_oracle=sa.address from  sync_adapters sa where type='PriceOracle' and p._version=sa.version;

update token_oracle set disabled_at=19752044 where version=1 or version=2;



-- insert into relations(block_num, owner, dependent, category) (select discovered_at, pool,oracle , 'PoolOracle'  from (select address pool from pools where _version=1) l join (select address oracle, discovered_at from sync_adapters where type='PriceOracle' and version=2) on true);
-- insert into relations(block_num, owner, dependent, category)  (select discovered_at, pool,oracle , 'PoolOracle'  from (select address pool from pools where _version=1) l join (select address oracle, discovered_at from sync_adapters where type='PriceOracle' and version=1) on true);
insert into relations(block_num, owner, dependent, category)  (select discovered_at, pool,oracle , 'PoolOracle'  from (select address pool from pools where _version=300) l join (select address oracle, discovered_at from sync_adapters where type='PriceOracle' and version=300) on true);
