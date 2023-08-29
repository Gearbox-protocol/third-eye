{
  adapters: [
    {
      type: 'QueryPriceFeed',
      address: '#YearnFeed_1',
      details: {
        token: { '#Token_2': [1, 50], '#Token_3': [1, 50], '#Token_4': [1, 50] },  // removed at 50
        pfType: 'YearnPF',
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#YearnFeed_3',
      details: {
        token: { '#Token_2': [50], '#Token_3': [50], '#Token_4': [50] },
        pfType: 'YearnPF',
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
