alter table farm_v3 add reward_token varchar(42);
alter table lm_rewards add reward_token varchar(42);
alter table lm_rewards drop constraint lm_rewards_pkey;
alter table lm_rewards add PRIMARY KEY (reward_token, user_address, pool);