-- drop all tables in db;
drop table account_operations;
drop table allowed_protocols;
drop table allowed_tokens;

drop table credit_manager_stats;

drop table credit_session_snapshots;
drop table credit_sessions;
drop table current_debts;
drop table dao_descriptions;
drop table dao_operations;
drop table debt_sync;
drop table debts;
drop table fast_check_params;
drop table faucet;
drop table gear_balances;
drop table liquidable_accounts;
drop table no_session_transfers;
drop table operations;
drop table parameters;
drop table pool_ledger;
drop table pool_stats;
drop table pools;
drop table price_feeds;
drop table profiles;
drop table schema_migrations;
drop table sync_adapters;
drop table token_oracle;
drop table tokens;
drop table transfer_account_allowed;
drop table treasury_snapshots;
drop table treasury_transfers;
drop table uniswap_chainlink_relations;
drop table uniswap_pool_prices;
drop table uniswap_pools;
drop table token_current_price;

drop materialized view ranking_30d;
drop materialized view ranking_7d;
drop table credit_managers;
drop table blocks;

--
drop table diesel_balances; 
drop table diesel_transfers;
drop table lm_rewards; 