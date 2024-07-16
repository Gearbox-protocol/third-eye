alter table diesel_balances add pool varchar(42);
alter table diesel_balances drop constraint diesel_balances_pkey;
update diesel_balances db set pool=pa from (select symbol diesel_sym, p.address pa from pools p join tokens t on t.address=p.diesel_token ) t1 where db.diesel_sym =t1.diesel_sym;
alter table diesel_balances add PRIMARY KEY (user_address, pool);
ALTER TABLE diesel_balances ALTER COLUMN diesel_sym DROP NOT NULL;

INSERT INTO diesel_balances (balance_bi, balance, pool, user_address, diesel_sym)
SELECT diesel_balance_bi balance_bi, diesel_balance_bi::float4/power(10, decimals), p.address, account, diesel_sym
FROM user_lmdetails_v3 ul join tokens t on t.symbol=ul.diesel_sym join pools p on p.diesel_token=t.address where p.address!='0x1DC0F3359a254f876B37906cFC1000A35Ce2d717';


-- alter table user_lmdetails_v3 drop column diesel_sym, drop column diesel_balance_bi;
-- alter table  diesel_balances drop column diesel_sym;