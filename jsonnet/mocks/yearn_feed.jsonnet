{
  adapters: [
    {
      type: 'QueryPriceFeed',
      address: '#YearnFeed_1',
      details: {
        pfType: 'YearnPF',
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#YearnFeed_3',
      details: {
        pfType: 'YearnPF',
      },
      lastSync: 1,
      version: 2,
    },
    {
      type: 'QueryPriceFeed',
      address: '#YearnFeed_4',
      details: {
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
  tokens: [
    { addr: '#Token_1', symbol: 'USDC', decimals: 6 },
    { addr: '#Token_2', symbol: 'YFI', decimals: 18 },
    { addr: '#Token_3', symbol: 'WETH', decimals: 18 },
    { addr: '#Token_4', symbol: 'DAI', decimals: 8 },
  ],
}
