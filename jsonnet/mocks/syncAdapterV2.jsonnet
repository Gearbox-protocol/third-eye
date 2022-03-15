{
  adapters: [
    {
      type: 'AddressProvider',
      address: '#AddressProvider_1',
      details: {
        dc:
          // data compressor should be deployed before it is used
          { '1': '#DC_1' },
        priceOracles: ['#PriceOracle_1'],
      },
      lastSync: 1,
    },
    {
      // on usdc
      type: 'Pool',
      address: '#Pool_1',
      lastSync: 2,
    },
    {
      type: 'AccountFactory',
      address: '#AccountFactory_1',
      lastSync: 2,
    },
    {
      type: 'ACL',
      address: '#ACL_1',
      lastSync: 2,
    },
    {
      // on usdc
      type: 'CreditManager',
      address: '#CreditManager_1',
      details: {
        facade: '#CreditFacade_1',
        configurator: '#CreditConfigurator_1',
      },
      version: 2,
      lastSync: 2,
    },
    {
      type: 'CreditConfigurator',
      address: '#CreditConfigurator_1',
      details: {
        creditManager: '#CreditManager_1',
      },
      version: 2,
      lastSync: 2,
    },
    {
      type: 'PriceOracle',
      address: '#PriceOracle_1',
      lastSync: 2,
    },
    {
      type: 'ChainlinkPriceFeed',
      address: '#ChainlinkPriceFeed_1',
      details: {
        oracle: '#Oracle_1',
        token: '#Token_1',
        dieselToken: '#DieselToken_1',
      },
      lastSync: 2,
    },
    {
      type: 'AccountManager',
      address: '#AccountManager_1',
      details: {
        accounts: ['#Account_1'],
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
      address: '#Pool_1',
      dieselToken: '#DieselToken_1',
      underlyingToken: '#Token_1',
    },
  ],
  cmState: [
    {
      underlyingToken: '#Token_1',
      pool: '#Pool_1',
      address: '#CreditManager_1',
    },
  ],
  tokens: [
    { address: '#Token_1', symbol: 'USDC', decimals: 6 },
    { address: '#Token_2', symbol: 'YFI', decimals: 18 },
    { address: '#Token_3', symbol: 'WETH', decimals: 18 },
    { address: '#Token_4', symbol: 'DAI', decimals: 8 },
  ],
}
