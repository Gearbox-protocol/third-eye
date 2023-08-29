update sync_adapters set last_sync=18020893-1 where type='PriceOracle'; -- priceoracle
delete from token_oracle where block_num>=18020893;
delete from price_feeds where block_num>=18020893;
delete from sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'CompositeChainlinkPF') and discovered_at>=18020893;
update sync_adapters set last_sync= 18020893-1 , disabled='f' where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'CompositeChainlinkPF') and last_sync>=18020893;


delete from debts where block_num>=18020893;
delete from current_debts where block_num>=18020893;
update debt_sync set last_calculated_at=18020893-1 where last_calculated_at>=18020893;

-- check if for QueryPriceFeed there is an entry of closure for tokens