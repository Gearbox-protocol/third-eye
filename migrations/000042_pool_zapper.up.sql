update sync_adapters set  details=( details || jsonb_build_object('USDC-farmedUSDCv3','0x6f42B83eCda76A8313Fd8a45ca18a3FdFD37bBc7') || jsonb_build_object('farmedUSDCv3','0xE2037090f896A858E3168B978668F22026AC52e7')) where details->>'USDC-farmedUSDCv3' ='0x0A8abaCEaaf0786DFe0AB420Ca540C9958d9D4F2';
update sync_adapters set  details=( details || jsonb_build_object('USDC-farmedUSDCv3','0x1aD0780a152fE66FAf7c44A7F875A36b1bf790F0') || jsonb_build_object('farmedUSDCv3','0xC853E4DA38d9Bd1d01675355b8c8f3BBC1451973')) where details->>'USDC-farmedUSDCv3' ='0xa010Fb889700986Dcf43d5fE110C968B7Dc05dAD';