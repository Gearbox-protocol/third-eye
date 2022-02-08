{
  adapters: [
    {
      type: 'AddressProvider',
      address: '#AddressProvider_1',
      lastSync: 1,
    },
    {
      type: 'Pool',
      address: '#Pool_1',
      lastSync: 2,
    },
    {
      type: 'CreditManager',
      address: '#CreditManager_1',
      lastSync: 2,
    },
    {
      type: 'CreditFilter',
      address: '#CreditFilter_1',
      details: {
        creditManager: '@CreditManager_1',
      },
      lastSync: 2,
    },
    {
      type: 'ChainlinkPriceFeed',
      address: '#ChainlinkPriceFeed_1',
      details: {
        oracle: '#Oracle_1',
        token: '#Token_1',
        underlyingToken: '#Token_3',
      },
      lastSync: 2,
    },
    {
      type: 'ChainlinkPriceFeed',
      address: '#ChainlinkPriceFeed_2',
      details: {
        oracle: '#Oracle_2',
        token: '#Token_2',
      },
      lastSync: 2,
    },
  ],
  poolState: [
    {
      address: '@Pool_1',
      dieselToken: '@Token_1',
      underlyingToken: '@Token_3',
    },
  ],
  cmState: [
    {
      underlyingToken: '@Token_3',
      pool: '@Pool_1',
      address: '@CreditManager_1',
    },
  ],
  tokens: [
    { address: '@Token_1', symbol: 'USDC', decimals: 6 },
    { address: '@Token_2', symbol: 'DAI', decimals: 8 },
    { address: '#Token_3', symbol: 'WETH', decimals: 18 },
  ],
}
