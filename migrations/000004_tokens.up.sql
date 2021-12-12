CREATE TABLE tokens (
    address varchar(42) PRIMARY KEY,
    symbol varchar(40),
    decimals integer
);

CREATE TABLE allowed_tokens (
    id SERIAL PRIMARY KEY,
    credit_manager varchar(42),
    token varchar(42),
    liquiditythreshold character varying(80),
    block_num integer
);
ALTER TABLE ONLY allowed_tokens
    ADD CONSTRAINT allowed_tokens_credit_manager_fkey FOREIGN KEY (credit_manager) REFERENCES credit_managers(address);
ALTER TABLE ONLY allowed_tokens
    ADD CONSTRAINT allowed_tokens_token_fkey FOREIGN KEY (token) REFERENCES tokens(address);
