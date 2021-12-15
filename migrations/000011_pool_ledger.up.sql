alter table pool_ledger add tx_hash varchar(66), add session_id varchar(100), add user_address varchar(42), add amount double precision, add amount_bi varchar(80);
alter table pool_ledger drop column liquidity, drop column address;
