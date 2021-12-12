
CREATE TABLE IF NOT EXISTS operations (
    id serial primary key,
    address varchar(42),
    tx_hash text,
    block_num integer,
    log_id integer,
    operation text,
    foreign key (block_num) references blocks(id));