update sync_adapters set address=details->>'oracle' where type='CompositeChainlinkPF';


--  update pool_ledger pl set amount=amount_bi::float4/power(10,pp.decimals) from  (select p.*, t,decimals from (select address , underlying_token from pools where _version=300) p join (select address,decimals from tokens) t on t.address=p.underlying_token) pp where pp.address = pl.pool and pl.event='AddLiquidity';

-- update pool_ledger pl set shares=shares_bi::float4/power(10,pp.decimals), amount=amount_bi::float4/power(10,pp.decimals) from  (select p.*, t,decimals from (select address , underlying_token from pools where _version=300) p join (select address,decimals from tokens) t on t.address=p.underlying_token) pp where pp.address = pl.pool and pl.event='RemoveLiquidity';