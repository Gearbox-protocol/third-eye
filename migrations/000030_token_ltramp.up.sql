create table token_ltramp (
    block_num integer,
    credit_manager varchar(42),
    token varchar(42),
    lt_initial integer,
    lt_final integer,
    ramp_start integer,
    ramp_end integer,
    PRIMARY KEY (block_num, credit_manager, token)
);

