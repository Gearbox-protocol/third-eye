# Third-eye

Third-eye is used for getting gearbox contracts-related data from evm compatible chains, calculating interesting parameters, and storing them. It uses [uber/fx](https://github.com/uber-go/fx), a dependency injection framework that prevents the need for maintaining global variables. It has two different engines which can be used standalone/together: `Sync` and `Debt` engines.

## Sync Engine
It syncs data from the chain using sync_adapters. Sync adapters are individual or bunch of contracts that can be queried individually over a range of block numbers without affecting other adapters. 

Sometimes, there are dependencies btw different contracts for a given block. This means contract `A` needs to be synced before `B` syncs for blocknumber `x`. For allowing multiple sync adapters to sync in parallel while maintaining these dependencies, we use another component: adapter kit. The adapter kit organises sync adapters in different levels based on `adapter types` i.e. CreditManager, CreditFacade. So that independent contracts in a level can sync in parallel. 

For eg:

> AccountFactory `syncs before`  AccountManager `syncs before`  CreditManager 

>`AccountFactory` gets the account address which are used by `AccountManager` for getting all token transfers on `Credit Accounts`, and `CreditManager` filter transfer that are not related to any gearbox `CreditManager`/`CreditFacade` event.

### Adapter Type

Adapters are of two types: `Event` and `Query` based.

#### Event Based
- `AccountFactory`: Adapter for getting credit account addresses and dao event `TakeForever`.
- `AccountManager`: Fetching token transfers (except for GearToken) on the all credit account.
- `ACL`: Getting dao events required to account control list: `OwnershipTransferred` and adding/ removing admins.
- `AddressProvider`: Address of other contracts.
- `ChainlinkPriceFeed(CPF)`: For syncing prices of tokens that have chainlink Feed available. 
    * PriceOracle has tokens and corresponding ChainlinkOracle addresses. Chainlink oracle contract internally manages a list of versioned AggregatedInterface contracts that stores the token's price. CPF adapter gets the latest AggregatedInterface address and listens for `AnswerUpdated` event. 
    * CPF checks for the new Aggregated interface(Feed) on the `ChainlinkOracle` contract. If a new interface is detected, another CPF adapter is created with the feed set to the new interface address and the same oracle address while the previous adapter is disabled.
- `CompositeChainlinkPF`: Composite chainlink price feed uses 2 chainlink price feeds for getting the token price in usd. One pair is token/ETH(A) and another is ETH/USD(B) price oracle. This adapter monitors `AnswerUpdated` event on the phaseAggregator behind these two priceFeeds. And uses below formula for calculating the price.
```
  Token's USD price = A*B/10^18
```

- `ContractRegister`: Getting addresses of CreditManager and pools.
- `GearToken`: Getting token transfers for Gear Token.
- `Pool`: Getting Pool operations like add/remove liquidity and borrow/repay etc.
- `PriceOracle`: Getting token and their registered PriceFeed. These priceFeeds can be ChainlinkPF, YearnPF, CurvePF, ZeroPF, AlmostZeroPF. 
- `Treasury`: Getting all token transfers on Treasury.
- `CreditManager`: This adapter is reponsible for bulk of operations. 
    * With help of `AccountManager` checks if there is any token transfers to credit account which doesn't emit Gearbox CreditManager/CreditFacade event. 
    * Using events on CreditManager, we can find all operations that are executed by individual CreditAccount. 
    * CreditManager is internally divided in v1 and v2 modules.  

#### Query Based
- `AggregatedBlockFeed`: Maintains list of yearn and curve price feeds, and uniswap v2/v3 pools for token/eth pairs. It uses multicall for getting involving prices after a set interval:
    * prices from yearn, curve feeds.
    * yearn and curve feeds internally uses chainlink price feed. If any major change occurs in the chainlink feed, yearn/curve token's price will also be affected. To account for this, aggregatedBlockFeed maintains a dependency graph of chainlink-based tokens and dependent yearn/curve tokens. Since chainlink price feed syncs before aggregatedBlockFeed, aggregatedblockFeed can get a list of block numbers where the chainlink-based token's price changed. If aggregatedBlockFeed missed getting the price for any of the chainlink feeds' updated blocks, it would fetch the price for dependent tokens for these remaining block numbers.

## Repo
Repo is used by both sync and debt engines to get data from DB and store fetched data to DB. It acts as an intermediate between different adapters. Repo syncs all this data to DB using a single transaction in `save.go`, this provides the atomic guarantee that either all data is saved or nothing. On crash, the sync engine will restart with the previous state fetched from DB.

## Debt Engine

This engine calculates debt parameters. These are:
 - total value
 - borrowed amount with interest
 - repayAmount 
 - pnl in underlying asset and USD denomination

This requires tracking the state of Gearbox protocols so that the calculated values are in sync with on-chain data. All this data is fetched using the sync engine and stored in DB. Repo provides an interface for this data, which the debt engine uses. Following values are tracked for calculating required parameters:

- CreditAccount `balance` and current `cumulative Index RAY`.
- `Price` of tokens
- `Liquidity Threshold`.
- `Pool Cumulative Index` and `blockNumber` of last operation on pool(add/remove liquidity and repay/borrow)

### Debt throttling conditions
There were around 20k active accounts in Kovan on 15 Jan 2022, storing entries on every block or even on the block where any value changes would result in hundreds of millions of rows in the debt table. To prevent the debt table from unexpected growth, we added throttling to slow the calculation of the debt parameters. The conditions for throttling are:

- Calculate for the block where we have an event for the credit Account. In such a case, `Credit Session Snapshot` will be present for that block. This is a `forced condition` and entry will always be added to DB.
- in the config, we have `ThrottleByHrsStr` variable which set the limit for how many times values can be calculated in a day.
- if the current debt entry's totalValue or borrowedAmountWithInterest deviates more than 5% from the previous debt entry present in DB.  
- if the previous debt entry's healthFactor is on different side of 10k than the current debt's healthFactor. 

### Checks for consistency in Debt Engine

The debt parameters calculated by the debt engine can deviate from on-chain data. Debt engine has to check in place to prevent this from happening and be alerted if that's the case. 
- If Credit Session Snapshot is present for that block, then the calculated totalValue and borrowedAmountWithInterest are checked to be within the precision of the underlying token. For eg, for WBTC we can tolerate a deviation of 0.001 BTC.
- If `DebtDCMatching` is enabled then for each credit account whenever debt parameters are calculated data from data compressor is fetched for that account. And values are checked to be within tolerable precision error.


### Current debt table

We also want to have snapshots of the latest state for each credit session. This goal can't be achieved with the debt table as due to throttling the debt parameters are calculated at different block numbers, and it can happen last debt blocks are different for each credit session. Apart from this, debt table is still very large with millions of entries. So queries on this debt takes some time. 

To solve this, another table for debts called `current_debts` is used. After the sync engine completes syncing till block number `x`, the debt engine firstly calculates debt snapshots for credit accounts for blocks that the sync engine fetched. Then for block number `x`, the debt parameters are calculated and stored in `current_debts` table. If there are accounts that are closed before `x`, then the current_debt logic skips those accounts. These accounts have their current debt calculated in the debt logic. Refer: [debt module](https://github.com/Gearbox-protocol/third-eye/blob/master/debts/engine.go#L221-L224) and [current_debt module](https://github.com/Gearbox-protocol/third-eye/blob/master/debts/current_debt.go#L36-L37).


### Notification for liquidations

Whenever the healthfactor of any credit account goes below or above 10k, the account becomes liquidable and safe respectively.
We store at which block the account becomes liquidable so that we can calculate the average liquidation time. This block number and whether the liquidable notification has been sent are stored in `liquidable_accounts` table.


__Note__: At some places, credit account and credit sessions are used interchangeably. As credit session is an abstract concept there is no contract representing it. Once a credit account is assigned by the credit manager to a borrower, it gets initialized with balance and some state, resulting in a credit session.