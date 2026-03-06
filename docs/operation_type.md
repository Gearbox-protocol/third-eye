# All account operations with their corresponding operation names

## V1 Operations (GetV1OperationOnLog)
 - `OpenCreditAccount(address,address,address,uint256,uint256,uint256)` → **OpenCreditAccount**
 - `CloseCreditAccount(address,address,uint256)` → **CloseCreditAccount**
 - `DirectTokenTransfer` → **DirectTokenTransfer**
 - `LiquidateCreditAccount(address,address,uint256)` → **LiquidateCreditAccount**
 - `IncreaseBorrowedAmount(address,uint256)` → **IncreaseBorrowedAmount**
 - `AddCollateral(address,address,uint256)` → **AddCollateral**
 - `TransferAccount(address,address)` → **TransferAccount**
 - `RepayCreditAccount(address,address)` → **RepayCreditAccount**

## V2 Operations (GetV2OperationOnLog)
 - `OpenCreditAccount(address,address,uint256,uint16)` → **OpenCreditAccount**
 - `LiquidateCreditAccount(address,address,address,uint256)` → **LiquidateCreditAccount**
 - `DecreaseBorrowedAmount(address,uint256)` → **DecreaseBorrowedAmount**
 - `TokenEnabled(address,address)` → **TokenEnabled**
 - `TokenDisabled(address,address)` → **TokenDisabled**
 - `CloseCreditAccount(address,address)` → **CloseCreditAccount**
 - `MultiCallStarted(address)` → **MultiCall**

## Adapter Operations (GetAdapterOperationOnLog)

### Uniswap/DEX Swaps
 - `swapExactTokensForTokens(uint256,uint256,address[],address,uint256)` → **Swap**
 - `swapTokensForExactTokens(uint256,uint256,address[],address,uint256)` → **Swap**
 - `exactOutput((bytes,address,uint256,uint256,uint256))` → **Swap**
 - `exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160))` → **Swap**
 - `exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160))` → **Swap**
 - `exactInput((bytes,address,uint256,uint256,uint256))` → **Swap**
 - `swapExactPtForToken(address,address,uint256,(address,uint256,address,address,(uint8,address,bytes,bool)),(address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes))` → **Swap**
 - `swapExactTokenForPt(address,address,uint256,(uint256,uint256,uint256,uint256,uint256),(address,uint256,address,address,(uint8,address,bytes,bool)),(address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes))` → **Swap**

### Yearn
 - `deposit(uint256)` → **YearnDeposit**
 - `withdraw(uint256)` → **YearnWithDrawn**

### Mellow
 - `deposit(address,uint256[],uint256,uint256)` → **Swap** (MellowDeposit)

### Curve
 - `exchange(int128,int128,uint256,uint256)` → **CurveExchange**
 - `exchange(uint256,uint256,uint256,uint256)` → **CurveExchange**
 - `remove_liquidity_one_coin(uint256,int128,uint256)` → **CurveRemoveLiquidityOneCoin**
 - `remove_liquidity(uint256,uint256[4])` → **CurveRemoveLiquidity**
 - `add_liquidity(uint256[2],uint256)` → **CurveAddLiquidity**
 - `add_liquidity(uint256[3],uint256)` → **CurveAddLiquidity**
 - `add_liquidity(uint256[4],uint256)` → **CurveAddLiquidity**

### Convex
 - `depositAll(uint256,bool)` → **ConvexDepositAndStake** (if _stake=true) or **ConvexDeposit** (if _stake=false)
 - `deposit(uint256,uint256,bool)` → **ConvexDepositAndStake** (if _stake=true) or **ConvexDeposit** (if _stake=false)
 - `RewardClaimed` → **RewardClaimed**
 - `getReward()` → **GetReward**
 - `stake(uint256)` → **ConvexStake**
 - `withdrawAndUnwrap(uint256,bool)` → **ConvexWithdraw** (if claim=false) or **ConvexWithdrawAndClaim** (if claim=true)
 - `withdrawAllAndUnwrap(bool)` → **ConvexWithdraw** (if claim=false) or **ConvexWithdrawAndClaim** (if claim=true)

### wstETH/Lido
 - `unwrap(uint256)` → **WstETHUnwrap**
 - `wrap(uint256)` → **WstETHWrap**
 - `submit(uint256,address)` → **LidoSubmit**

## V3 Operations (GetV3OperationOnLog)
 - `StartMultiCall(address,address)` → **MultiCall**
 - `IncreaseDebt(address,uint256)` → **IncreaseBorrowedAmount**
 - `DecreaseDebt(address,uint256)` → **DecreaseBorrowedAmount**
 - `WithdrawCollateral(address,address,uint256,address)` → **WithdrawCollateral**
 - `UpdateQuota` → **UpdateQuota**
 - `OpenCreditAccount(address,address,address,uint256)` → **OpenCreditAccount**

## V3 Adapter Operations (GetAdapterOperationOnLogv3)

### Maker/MakerDAO
 - `redeem(uint256,address,address)` → **MakerRedeem**
 - `usdsToDai(address,uint256)` → **MakerRedeem**
 - `daiToUsds(address,uint256)` → **VaultDeposit**

### Additional Curve Operations
 - `remove_liquidity_one_coin(uint256,uint256,uint256)` → **CurveRemoveLiquidityOneCoin**
 - `exchange_underlying(int128,int128,uint256,uint256)` → **CurveExchange**
 - `withdrawAll(bool)` → **CurveClaims**
 - `withdrawAll(uint256)` → **CurveWithdrawal**

### Additional Uniswap
 - `exactInputSingle((address,address,address,uint256,uint256,uint256,uint160))` → **Swap**

### Balancer
 - `swap((bytes32,uint8,address,address,uint256,bytes),(address,bool,address,bool),uint256,uint256)` → **BalancerSwap**

### Additional Swaps/Generic
 - `add_liquidity(uint256[],uint256)` → **Swap** (llamaethena swap / manualSwap)
 - `redeemPyToToken(address,address,uint256,(address,uint256,address,address,(uint8,address,bytes,bool)))` → **Swap**

## Unmapped Operations (Still not found in operation files)
 - `swapIn(bool,uint256,uint256,address)`
 - `withdraw(uint256,bool)`
 - `multiAcceptAndClaim(address,uint256[],uint256[][],address,uint256)`
 - `migrate((address,address,address,(address,uint256,uint96,bool,bool,(bool,address,uint256))[],uint256,(address,bytes)[],(address,bytes)[],address[],uint256,uint256,uint256,uint256))`
 - `PartialLiquidation`
 - `deposit(uint256,address)`
 - `deposit(address,uint256,address,address,address)`
 - `swapSingleTokenExactIn(address,address,address,uint256,uint256,uint256,bool,bytes)`
 - `removeLiquiditySingleTokenExactIn(address,uint256,address,uint256,bool,bytes)`

## Notes
- These unique account operation signatures can be fetched using `select distinct (action) from account_operations;
`
- Code for mapping in [account_operation.go](https://github.com/Gearbox-protocol/charts_server/blob/master/core/account_operation.go), [operation_type(v1/v2)](https://github.com/Gearbox-protocol/charts_server/blob/master/core/operation_type.go),  [operation_type(v3)](https://github.com/Gearbox-protocol/charts_server/blob/master/core/operation_type_v3.go)
