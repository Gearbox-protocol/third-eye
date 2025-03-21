alter table debt_sync add debt_block integer, add tvl_block integer;
update debt_sync set debt_block=last_calculated_at;
update debt_sync set tvl_block=0;
alter table debt_sync drop column last_calculated_at;

alter table tvl_snapshots add expected_liq DOUBLE PRECISION;