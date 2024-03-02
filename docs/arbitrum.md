## Changes
- geth client doesn't support different tx types on arbitrum. Use offchainlabs geth fork where they added webhook for support the tx and there know custom types for blocks etc. that are compatible with geth.
- stETH contract is not available on arbitrum. wstETH is used. so, no need to track the rebase token such as stETH and don't use logic for converting shares to balance of stETH on arbitrum.
- Multicall2 address on arbitrum is different from ethereum. Ethereum multicall2 is deployed by aave and has different address. https://arbiscan.io/address/0x842eC2c7D803033Edf55E478F461FC547Bc54EB2#code
- There are USDC, USDC.e on arbitrum due to bridged and native versions of USDC. Change chainlink adapter to support multiple tokens.
- Since we use zappers, the liquidity providing user is replaced with zapper address in the AddLiquidity/RemoveLiquidity  event, we have to track user call to get the actual user address and replace in the event. We have logic for dUSDC-> farmedUSDCv3 zappers too, but these zappers are not present on arbitrum disable them.
- farming pools info is not added to sdk-gov , like it is added for mainnet. In the tokens.ts for mainnet we have sdUSDCv3 , sdWETHv3 and other token details that are farmed tokens. This info is missing handle this in the ts-> go config generator. 
- disable lm rewards v2 logic for arbitrum, disable the sync adapter for lmrewards v2.
- On mainnet the price of gearbox token is fetched using curve pool, 0x0E9B5B092caD6F1c5E6bc7f89Ffe1abb5c95F1C2 , that is missing on arbitrum, need to add logic for that in the inch oracle in sdk-go. [TODO]
- set price in spot oracle (1inch oracle) for USDC_e, USDC token on networks other table mainnet.