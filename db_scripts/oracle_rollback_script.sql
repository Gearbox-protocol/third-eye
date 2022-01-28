DELETE FROM price_feeds WHERE feed IN (select address from sync_adapters where type in ('YearnPriceFeed', 'ChainlinkPriceFeed') and disabled='f');
DELETE FROM uniswap_pool_prices WHERE token IN (select details->>'token' from sync_adapters where type in ('YearnPriceFeed', 'ChainlinkPriceFeed') and disabled='f');
update sync_adapters set last_sync = firstlog_at - 1 where type in ('YearnPriceFeed', 'ChainlinkPriceFeed', 'Treasury') and disabled='f';
delete from treasury_snapshots;delete from treasury_transfers;
-- alter table price_feeds add uniswapv2_price DOUBLE PRECISION, add uniswapv3_twap DOUBLE PRECISION,
-- add uniswapv3_price DOUBLE PRECISION, add uni_price_fetch_block integer;