local utils = import '../utils.libsonnet';
{
  mocks: {
    syncAdapters: 'mocks/syncAdapterV2.json',
  },
  states: {
    oracles: [
      {  // for v1
        oracle: '#Oracle_0',
        block: 1,
        feed: '#ChainlinkPriceFeed_0',
      },  // for v2
      {
        oracle: '#Oracle_1',
        block: 1,
        feed: '#ChainlinkPriceFeed_1',
      },
      {
        oracle: '#Oracle_2',
        block: 1,
        feed: '#ChainlinkPriceFeed_2',
      },
      {
        oracle: '#Oracle_3',
        block: 1,
        feed: '#ChainlinkPriceFeed_3',
      },
      {
        oracle: '#Oracle_4',
        block: 1,
        feed: '#ChainlinkPriceFeed_4',
      },
    ],
  },
  blocks: {
    // block with open and borrow more underlying asset
    '3': {
      events: [
        // init setup
        {
          // credit filter on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'TokenAllowed(address)',
            '#Token_1',
          ],
          txHash: '!#Hash_1',
        },
        {
          // credit filter on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'TokenLiquidationThresholdUpdated(address,uint16)',
            '#Token_1',
          ],
          data: [
            'uint16:9000',
          ],
          txHash: '!#Hash_1',
        },
        {
          // price chainlink on usdc
          address: '#CreditConfigurator_1',
          txHash: '!#Hash_2',
          topics: [
            'LimitsUpdated(uint256,uint256)',
          ],
          data: [
            // minAnount
            utils.bigIntTopic(1000, 6),
            // maxAmount
            utils.bigIntTopic(5000, 6),
          ],
        },
        {
          // price chainlink on usdc
          address: '#CreditConfigurator_1',
          txHash: '!#Hash_2',
          topics: [
            'FeesUpdated(uint16,uint16,uint16)',
          ],
          data: [
            // feeInterest
            'uint16:1000',
            // feeLiquidation
            'uint16:200',
            // liquidationDiscount
            'uint16:9500',
          ],
        },
        {
          // price chainlink on usdc
          address: '#ChainlinkPriceFeed_1',
          txHash: '!#Hash_3',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // usdc price
            utils.bigIntTopic(1, 8),
            // roundid
            utils.bigIntTopic(1, 0),
          ],
          data: [],
        },
      ],
    },
  },
}
