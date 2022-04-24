DELETE FROM price_feeds WHERE feed IN (select address from sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed') and disabled='f');
DELETE FROM uniswap_pool_prices WHERE token IN (select details->>'token' from sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed') and disabled='f');
update sync_adapters set last_sync = firstlog_at - 1 where type in ('QueryPriceFeed', 'ChainlinkPriceFeed') and disabled='f';
delete from uniswap_pools;
delete from uniswap_chainlink_relations;


-- DELETE FROM price_feeds WHERE block_num >= 14029900 and feed IN (select address from sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed') and disabled='f');
-- DELETE FROM uniswap_pool_prices WHERE block_num >= 14029900 and token IN (select details->>'token' from sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed') and disabled='f');
-- update sync_adapters set last_sync = 14029900 - 1 where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'Treasury') and disabled='f';
-- delete from treasury_snapshots where block_num >=14029900; delete from treasury_transfers where block_num >=14029900;
-- delete from uniswap_pools where block_num >=14029900;
-- delete from uniswap_chainlink_relations where block_num >=14029900;