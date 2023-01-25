update sync_adapters set last_sync=15997386-1 where type='PriceOracle'; -- priceoracle
delete from token_oracle where block_num>=15997386;
delete from price_feeds where block_num>=15997386;
delete from sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'CompositeChainlinkPF') and discovered_at>=15997386;
update sync_adapters set last_sync= 15997386-1 , disabled='f' where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'CompositeChainlinkPF') and last_sync>=15997386;


delete from debts where block_num>=15997386;
delete from current_debts where block_num>=15997386;
update debt_sync set last_calculated_at=15997386-1 where last_calculated_at>=15997386;

-- check if for QueryPriceFeed there is an entry of closure for tokens