{
  blockNum: 6,
  daoOperations: [
    {
      Args: {
        priceFeed: '#Oracle_5',
        token: '#Token_1',
        reserve: false,
      },
      BlockNumber: 6,
      Contract: '#PriceOracle_2',
      LogID: 1,
      TxHash: '#Hash_9',
      Type: 12,
    },
  ],
  priceFeeds: [
    {
      blockNum: 6,
      feed: '#ChainlinkPriceFeed_1',
      price: 0.0004,
      priceBI: '400000000000000',
      roundId: 1,
    },
    {
      blockNum: 6,
      feed: '#ChainlinkPriceFeed_5',
      price: 1,
      priceBI: '100000000',
      roundId: 1,
    },
  ],
  timestamp: 518400,
  tokenOracles: [
    {
      blockNum: 6,
      priceOracle: '#PriceOracle_2',
      feed: '#ChainlinkPriceFeed_5',
      oracle: '#Oracle_5',
      DisabledAt: 0,
      token: '#Token_1',
      reserve: false,
      version: 2,
    },
  ],
}
