{
  adapters: [
    {
      type: 'QueryPriceFeed',
      address: '#YearnFeed_1',
      details: {
        token: { '#Token_1': [1] },  // removed at 50
        pfType: 'YearnPF',
        mergedPFVersion: 2,
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#CurvePriceFeed_1',
      details: {
        token: { '#Token_2': [1] },
        pfType: 'CurvePF',
        mergedPFVersion: 4,
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#SingleAsset_1',
      details: {
        token: { '#Token_3': [1] },
        pfType: 'SingleAssetPF',
        mergedPFVersion: 4,
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#Redstone_1',
      details: {
        token: { '0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599': [1] },
        pfType: 'RedStonePF',
        mergedPFVersion: 4,
        info: {
          '#Redstone_1': {
            type: 15,
            dataServiceId: 'redstone-primary-prod',
            dataId: 'BTC',
            signersThreshold: 5,
            token: '0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599',
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
        token: { '0xCd5fE23C85820F7B72D0926FC9b05b43E359b7ee': [1] },
        pfType: 'CompositeRedStonePF',
        mergedPFVersion: 4,
        info: {
          '#CompositeRedstone_1': {
            type: 15,
            dataServiceId: 'redstone-primary-prod',
            dataId: 'weETH_FUNDAMENTAL',
            signersThreshold: 5,
            token: '0x8C23b9E4CB9884e807294c4b4C33820333cC613c',
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
}
