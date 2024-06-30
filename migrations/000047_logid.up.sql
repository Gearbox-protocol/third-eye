alter table allowed_tokens add log_id integer;
alter table token_ltramp add log_id integer;
-- delete from debts where block_num>=221760836;
-- update debt_sync set last_calculated_at=221760835;