-- for mainnet

-- update sync_adapters set  details=( details || jsonb_build_object('pfType','CompositeRedStonePF')) where address='0x1BC358C61D3d62439Abe4883e3BeF30c9b0a9c63';
-- update token_oracle set feed_type='CompositeRedStonePF' where oracle='0x1BC358C61D3d62439Abe4883e3BeF30c9b0a9c63';


-- for arbitrum
-- latest ezETH/ETH/USD
update sync_adapters set  details=( details || jsonb_build_object('pfType','CompositeRedStonePF')) where address='0xCbeCfA4017965939805Da5a2150E3DB1BeDD0364';
update token_oracle set feed_type='CompositeRedStonePF' where oracle='0xCbeCfA4017965939805Da5a2150E3DB1BeDD0364';
-- previous ezETH/ETH/USD
update sync_adapters set  details=( details || jsonb_build_object('pfType','CompositeRedStonePF')) where address='0x814E6564e8cda436c1ab25041C10bfdb21dEC519';
update token_oracle set feed_type='CompositeRedStonePF' where oracle='0x814E6564e8cda436c1ab25041C10bfdb21dEC519';