DELETE FROM price_feeds WHERE token IN (select address from sync_adapters where type in ('YearnPriceFeed', 'ChainlinkPriceFeed'), disabled='f')
update sync_adapters set last_sync = firstlog_at - 1 where type in ('YearnPriceFeed', 'ChainlinkPriceFeed'), disabled='f';

-- alter table price_feeds add uniswapv2_price DOUBLE PRECISION,uniswapv3_twap DOUBLE PRECISION,
-- uniswapv3_price DOUBLE PRECISION,uni_price_fetch_block DOUBLE PRECISION;