alter table pool_stats add base_borrow_apy double precision, add base_borrow_apy_bi varchar(80);
update pool_stats set base_borrow_apy=borrow_apy, base_borrow_apy_bi=borrow_apy_bi;
alter table pool_stats drop column borrow_apy , drop column borrow_apy_bi;

--
alter table pools add base_borrow_apy_bi double precision;
update pools set base_borrow_apy_bi=borrow_apy_bi;
alter table pools drop column borrow_apy_bi;

--
alter table credit_session_snapshots add extra_quota_apy integer;