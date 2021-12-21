create index debts_session_id_index on debts using BTREE (session_id,block_num);

create table liquidable_accounts (
    session_id varchar(100) PRIMARY KEY,
    block_num integer);

create table profiles (profile text);