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
- `ContractRegister`: Getting addresses of CreditManager and pools.
- `GearToken`: Getting token transfers for Gear Token.
- `Pool`: Getting Pool operations like add/remove liquidity and borrow/repay etc.
- `PriceOracle`: Getting token and their registered PriceFeed. These priceFeeds can be ChainlinkPF, YearnPF, CurvePF, ZeroPF. 
- `Treasury`: Getting all token transfers on Treasury.
- `CreditManager`: This adapter is reponsible for bulk of operations. 
    * With help of `AccountManager` checks if there is any token transfers to credit account which doesn't emit Gearbox CreditManager/CreditFacade event. 
    * Using events on CreditManager, we can find all operations that are executed by individual CreditAccount. 
    * CreditManager is internally divided in v1 and v2 modules.  

#### Query Based
- `AggregatedBlockFeed`: Maintains list of yearn and curve price feeds, and uniswap v2/v3 pools for token/eth pairs. It uses multicall for getting involving prices after a set interval:
    * prices from yearn and curve feeds.
    * uni v3 TWAP, uniswap v2 quoted price and uniswap v3 quoted price. 


## Debt Engine

