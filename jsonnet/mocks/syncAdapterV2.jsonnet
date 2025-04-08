{
  adapters: [
    {
      type: 'AddressProvider',
      address: '#AddressProvider_1',
      details: {
        dc:
          // data compressor should be deployed before it is used
          { '1': '#DC_1' },
        priceOracles: { '1': '#PriceOracle_0', '2': '#PriceOracle_1' },
      },
      lastSync: 1,
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
    // v1 priceoracle
    {
      type: 'PriceOracle',
      address: '#PriceOracle_0',
      lastSync: 1,
      version: 1,
      disabled_at: 2,
      disabled: true,
    },
    {
      type: 'ChainlinkPriceFeed',
      address: '#ChainlinkPriceFeed_0',
      details: {
        oracle: '#Oracle_0',
      },
      lastSync: 2,
      version: 1,

    },
    // v2
    {
      // on usdc
      type: 'Pool',
      address: '#Pool_1',
      lastSync: 2,
      version: 2,
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
    // v2 priceoracle
    {
      type: 'PriceOracle',
      address: '#PriceOracle_1',
      lastSync: 2,
      version: 2,
    },
    {
      type: 'ChainlinkPriceFeed',
      address: '#ChainlinkPriceFeed_1',
      details: {
        oracle: '#Oracle_1',
      },
      lastSync: 2,
      version: 2,

    },
    {
      type: 'ChainlinkPriceFeed',
      address: '#ChainlinkPriceFeed_2',
      details: {
        oracle: '#Oracle_2',
      },
      lastSync: 2,
      version: 2,
    },
    {
      type: 'ChainlinkPriceFeed',
      address: '#ChainlinkPriceFeed_3',
      details: {
        oracle: '#Oracle_3',
        token: '#Token_3',
        mergedPFVersion: 2,
      },
      lastSync: 2,
      version: 2,
    },
  ],
  poolState: [
    {
      address: '#Pool_1',
      dieselToken: '#DieselToken_1',
      underlyingToken: '#Token_1',
      priceOracle: '#PriceOracle_0',
    },
  ],
  cmState: [
    {
      underlyingToken: '#Token_1',
      pool: '#Pool_1',
      address: '#CreditManager_1',
    },
  ],
  poToTokenOracles: {
    '#PriceOracle_0': {
      '#Token_1': {
        feed: '#ChainlinkPriceFeed_0',
        type: 'ChainlinkPriceFeed',
        blockNum: 2,
        version: 1,
      },
    },
    '#PriceOracle_1': {
      '#Token_1': {
        feed: '#ChainlinkPriceFeed_1',
        type: 'ChainlinkPriceFeed',
        blockNum: 2,
        version: 2,
      },
      '#Token_2': {
        feed: '#ChainlinkPriceFeed_2',
        type: 'ChainlinkPriceFeed',
        blockNum: 2,
        version: 2,
      },
      '#Token_3': {
        feed: '#ChainlinkPriceFeed_3',
        type: 'ChainlinkPriceFeed',
        blockNum: 2,
        version: 2,
      },
    },
  },
}
