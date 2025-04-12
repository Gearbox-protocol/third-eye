{
  adapters: [
    {
      type: 'QueryPriceFeed',
      address: '#YearnFeed_1',
      details: {
        token: { '#Token_1': [1] },  // removed at 50
        pfType: 'YearnPF',
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#CurvePriceFeed_1',
      details: {
        pfType: 'CurvePF',
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#SingleAsset_1',
      details: {
        pfType: 'SingleAssetPF',
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#Redstone_1',
      details: {
        pfType: 'RedStonePF',
        info: {
          '#Redstone_1': {
            type: 15,
            dataServiceId: 'redstone-primary-prod',
            dataId: 'BTC',
            signersThreshold: 5,
            token: '0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599',
            feed: '#Redstone_1',
          },
        },
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#CompositeRedstone_1',
      details: {
        pfType: 'CompositeRedStonePF',
        info: {
          '#CompositeRedstone_1': {
            type: 15,
            dataServiceId: 'redstone-primary-prod',
            dataId: 'weETH_FUNDAMENTAL',
            signersThreshold: 5,
            token: '0x8C23b9E4CB9884e807294c4b4C33820333cC613c',
            feed: '#Redstone_1',
          },
        },
      },
      lastSync: 1,
      version: 2,
    },

  ],
  poolState: [
    {
      address: '#Token_1',
    },
  ],
  poToTokenOracles: {
    '#PriceOracle_1': {
      '#Token_1': {
        feed: '#YearnFeed_1',
        type: 'YearnFeed',
        blockNum: 1,
        version: 2,
      },
      '#Token_2': {
        feed: '#CurvePriceFeed_1',
        type: 'CurvePF',
        blockNum: 1,
        version: 2,
      },
      '#Token_3': {
        feed: '#SingleAsset_1',
        type: 'SingleAsset',
        blockNum: 1,
        version: 2,
      },
      '0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599': {
        feed: '#Redstone_1',
        type: 'Redstone',
        blockNum: 1,
        version: 2,
      },
      '0xCd5fE23C85820F7B72D0926FC9b05b43E359b7ee': {
        feed: '#CompositeRedstone_1',
        type: 'CompositeRedstone',
        blockNum: 1,
        version: 2,
      },
    },
  },
}
