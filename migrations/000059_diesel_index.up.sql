ALTER TABLE diesel_balances  DROP CONSTRAINT diesel_balances_pkey;
ALTER TABLE diesel_balances  ADD PRIMARY KEY(pool, user_address);