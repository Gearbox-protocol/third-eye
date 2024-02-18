--
alter table token_oracle add reserve boolean;
update token_oracle set reserve= 'f';


-- upgraded to v3 at 18798084
--
alter table price_feeds add merged_pf_version integer;
update price_feeds set merged_pf_version=1 where price_in_usd='f';
update price_feeds set merged_pf_version= 2 where price_in_usd='t' and block_num < 18798084 ;
update price_feeds set merged_pf_version= 2 where price_in_usd='t' and block_num >= 18798084  and feed='0x4A7b3F6c4aaB7Bc5617D6c30C3f22bAeBbc34F43'; -- yvDAI
update price_feeds set merged_pf_version=4 where price_in_usd='t' and block_num >= 18798084 and feed in (select feed from token_oracle where version=300); -- v3PF_main

update price_feeds set merged_pf_version=6 where block_num >= 18798084 and feed in (select distinct feed from (select distinct on (token) * from token_oracle where version=2 order by token, block_num desc) t where  token not in  (select token from token_oracle where version=300)); -- tokens that don't have entyr for v3 not are present in v2.

update price_feeds set merged_pf_version=6 where price_in_usd='t' and block_num >= 18798084 and feed in (select feed from (select * from token_oracle where version=300  union all select * from (select distinct on (token) * from token_oracle where version=2 order by token, block_num desc) k ) t group by feed,token having count(*)>1); -- feed that have both active v2 and v3 price oracle


delete from price_feeds where feed='0x37bC7498f4FF12C19678ee8fE19d713b87F6a9e6' and block_num > 17217055; -- // as there is already another chainlink activated 0xE62B71cf983019BFf55bC83B48601ce8419650CC

--
update sync_adapters set details=( details || jsonb_build_object('mergedPFVersion',6))  where address in (select distinct feed from (select distinct on (token) * from token_oracle where version=2 order by token, block_num desc) t where  token not in  (select token from token_oracle where version=300));

update sync_adapters set details=( details || jsonb_build_object('mergedPFVersion',4)) where version=300 and type in ('ChainlinkPriceFeed', 'CompositeChainlinkPF', 'QueryPriceFeed');

update sync_adapters set  details=( details || jsonb_build_object('mergedPFVersion',2)) where version =2 and address not in (select distinct feed from (select distinct on (token) * from token_oracle where version=2 order by token, block_num desc) t where  token not in  (select token from token_oracle where version=300)) and address!='0x6385892aCB085eaa24b745a712C9e682d80FF681';
-- below query overrides the above query when a feed is active in both 2 and 6 .
update sync_adapters set details=( details || jsonb_build_object('mergedPFVersion',6))  where address in (select feed from (select * from token_oracle where version=300  union all select * from (select distinct on (token) * from token_oracle where version=2 order by token, block_num desc) k ) t group by feed,token having count(*)>1); 


update sync_adapters set details=( details || jsonb_build_object('mergedPFVersion',1)) where version=1 and type in ('ChainlinkPriceFeed', 'CompositeChainlinkPF', 'QueryPriceFeed') and details is not null;

--

create  table t as (select * from (select distinct on (token) * from token_oracle where version=2 order by token, block_num desc) t where  token not in  (select token from token_oracle where version=300));
update t set block_num=18798084 , version=300 ;
insert into token_oracle  select * from t;
drop table t;

alter table price_feeds drop column price_in_usd;

alter table token_oracle drop constraint token_oracle_pkey;
alter table token_oracle add PRIMARY KEY (block_num, token, version, reserve);
alter table price_feeds drop constraint price_feeds_pkey;
delete from price_feeds where feed = '0x91401cedCBFd9680cE193A5F54E716504233e998';
alter table price_feeds add PRIMARY KEY (block_num, token, merged_pf_version);


update  price_feeds  set merged_pf_version =6 where block_num > 18798084 and token in ('0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599','0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84','0x5A98FcBEA516Cf06857215779Fd812CA3beF1B32') and merged_pf_version!=1;
update sync_adapters set details=( details || jsonb_build_object('mergedPFVersion',6)) where type = 'CompositeChainlinkPF' and disabled='f' and version =300 and details->>'token'  in ('0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599','0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84','0x5A98FcBEA516Cf06857215779Fd812CA3beF1B32') ;