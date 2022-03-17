{
  adapters: [
    {
      type: 'AddressProvider',
      address: '#AddressProvider_1',
      details: {
        dc:
          // data compressor should be deployed before it is used
          { '1': '0x0050b1ABD1DD2D9b01ce954E663ff3DbCa9193B1' },
        priceOracles: ['#PriceOracle_1'],
      },
      lastSync: 1,
      version: 1,
    },
    {
      // on usdc
      type: 'Pool',
      address: '#Pool_1',
      lastSync: 2,
      version: 1,
    },
    {
      type: 'AccountFactory',
      address: '#AccountFactory_1',
      lastSync: 2,
      version: 1,
    },
    {
      type: 'ACL',
      address: '#ACL_1',
      lastSync: 2,
      version: 1,
    },
    {
      // on usdc
      type: 'CreditManager',
      address: '#CreditManager_1',
      lastSync: 2,
      version: 1,
    },
    {
      type: 'CreditFilter',
      address: '#CreditFilter_1',
      details: {
        creditManager: '#CreditManager_1',
      },
      lastSync: 2,
      version: 1,
    },
    {
      type: 'PriceOracle',
      address: '#PriceOracle_1',
      lastSync: 2,
      version: 1,

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
      version: 1,

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
      version: 1,

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
