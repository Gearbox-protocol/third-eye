delete from farm_V3;
delete from user_lmdetails_v3 ;
delete from sync_adapters where type='LMRewardsv3';

WITH data as (select * from sync_adapters where type = 'AddressProvider')
insert into sync_adapters(address, type, last_sync, firstlog_at, version, discovered_at, disabled) values (
    '0x00000000000000000000000000000000000beef3', 
    'LMRewardsv3', 
    (SELECT firstlog_at-1 from data),
    (SELECT firstlog_at from data),
    '300',
    (SELECT firstlog_at from data),
    'f'
    );