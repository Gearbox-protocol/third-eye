-- reset all price_oracles, so that new token_oracles are synced, aggregatedBlockFeed refetches all the blocks and chainlink/compositeChainlink resyncs. Works if some token/feed pairs are missed.
update sync_adapters set last_sync=firstlog_at-1 where type='PriceOracle'; 
DELETE FROM price_feeds;
DELETE FROM sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'CompositeChainlinkPF');
delete from uniswap_chainlink_relations;
delete from uniswap_pools;
delete from token_oracle;