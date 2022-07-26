update sync_adapters set last_sync=discovered_at-1 where type='Treasury';

delete from treasury_transfers;
delete from treasury_snapshots;