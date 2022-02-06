update sync_adapters set last_sync=firstlog_at-1 where type='PriceOracle'; 
DELETE FROM price_feeds;
DELETE FROM sync_adapters where type in ('YearnPriceFeed', 'ChainlinkPriceFeed');
delete from uniswap_chainlink_relations;
delete from uniswap_pools;
delete from token_oracle;