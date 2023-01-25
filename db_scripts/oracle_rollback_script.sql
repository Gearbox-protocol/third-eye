-- delete all prices fetched from price_feeds and reset them to sync again from firstlog_at-1.
DELETE FROM price_feeds WHERE feed IN (select address from sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'CompositeChainlinkPF') and disabled='f');
DELETE FROM uniswap_pool_prices WHERE token IN (select details->>'token' from sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'CompositeChainlinkPF') and disabled='f');
update sync_adapters set last_sync = firstlog_at - 1 where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'CompositeChainlinkPF') and disabled='f';
delete from uniswap_pools;
delete from uniswap_chainlink_relations;

