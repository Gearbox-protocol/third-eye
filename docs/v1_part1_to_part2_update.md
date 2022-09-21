# ChangeLog

Changes btw gearbox v2 part 1(which was created before April) and part2(which was created till July):

- Signature for feesUpdated, liquidationThresholdUpdated, opencreditAccount is different and uses uint16 instead of uint256 for same fields.
- DataCompressor call failed for 0x4 or the 8th datacompressor added to gearbox v2 part 1 gearbox deployment. So, we have to manually calculate the `credit account data`. This has been fixed in the 9th dc.
- In priceOracle, for supporting NFT the convert takes creditAccount as the first param. Third-eye uses convert call in getValueInCurrency(used for collateral value calculation in usd and underlying token) and for getPriceInUSD which has v1 and v2 logic. v1 logic converts token price to USDC to get the USD denominated value. For v2, price is fetched using latestRoundData.
- TokenLiquidationThresholUpdated is not emitted for underlying token as it is set in the FeesUpdated event on creditConfigurator. [This method](https://github.com/Gearbox-protocol/contracts-v2/blob/581000e1948ef6008e8faa5dce3fc2177d17488d/contracts/credit/CreditConfigurator.sol#L433) will has set the liquidation threshold for tokens if the previous lt is more than the underlying token lt. The new value will be equal to underlying token lt. 
- In the creditconfigurator `_setParams` is called in the constructor with default values and no event is emitted. There is no feesupdated in current kovan setup. And for underlying token and fees we will have to calc by making call.
- Another case, where event is not emitted in some cases. 
https://github.com/Gearbox-protocol/contracts-v2/blob/main/contracts/credit/CreditConfigurator.sol#L114-L116 while credit facade is updated we haven’t emitted creditfacadeupgraded.
- addcollateral in executeorder can have onbehalOf different than the mainAction user/borrower. handled with issue #37.
