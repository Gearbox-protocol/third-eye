alter table user_lmdetails_v3 add farmed_balance_bi varchar(80), add farmed_balance double precision, add diesel_balance_bi varchar(80);
update user_lmdetails_v3 set farmed_balance_bi=balances_bi , farmed_balance=balances, diesel_balance_bi=diesel_balance;
alter table user_lmdetails_v3 drop column balances_bi, drop column balances, drop column diesel_balance;
