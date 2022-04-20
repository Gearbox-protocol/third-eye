{
  '3': {
    blockNum: 3,
    priceFeeds: [
      {
        blockNum: 3,
        feed: '#ChainlinkPriceFeed_1',
        price: 0.0004,
        priceBI: '400000000000000',
        isPriceInUSD: false,
        roundId: 1,
        token: '#Token_1',
        uniPriceFetchBlock: 0,
        uniswapv2Price: 0,
        uniswapv3Price: 0,
        uniswapv3Twap: 0,
      },
      {
        blockNum: 3,
        feed: '#ChainlinkPriceFeed_2',
        price: 8,
        priceBI: '8000000000000000000',
        isPriceInUSD: false,
        roundId: 1,
        token: '#Token_2',
        uniPriceFetchBlock: 0,
        uniswapv2Price: 0,
        uniswapv3Price: 0,
        uniswapv3Twap: 0,
      },
    ],
    timestamp: 259200,
  },
  '4': {
    blockNum: 4,
    timestamp: 345600,
    tokenOracles: [
      {
        blockNum: 4,
        feed: '#ChainlinkPriceFeed_4',
        oracle: '#Oracle_2',
        token: '#Token_2',
        version: 1,
      },
    ],
  },
  '5': {
    blockNum: 5,
    daoOperations: [
      {
        Args: {
          priceFeed: '#Oracle_3',
          token: '#Token_1',
        },
        BlockNumber: 5,
        Contract: '#PriceOracle_1',
        LogID: 0,
        TxHash: '#Hash_3',
        Type: 12,
      },
    ],
    priceFeeds: [
      {
        blockNum: 5,
        feed: '#ChainlinkPriceFeed_4',
        price: 10,
        priceBI: '10000000000000000000',
        isPriceInUSD: false,
        roundId: 1,
        token: '#Token_2',
        uniPriceFetchBlock: 0,
        uniswapv2Price: 0,
        uniswapv3Price: 0,
        uniswapv3Twap: 0,
      },
    ],
    timestamp: 432000,
    tokenOracles: [
      {
        blockNum: 5,
        feed: '#ChainlinkPriceFeed_3',
        oracle: '#Oracle_3',
        token: '#Token_1',
        version: 1,
      },
    ],
  },
}
