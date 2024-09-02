local utils = import '../utils.libsonnet';
local borrowedAmount = 4000;
local extraBorrowedAmount = 1000;
local initialAmount = 1000;
local newCollateral = 1;
{
  mocks: {
    syncAdapters: 'mocks/aqf_multiple_adapters.json',
    tokens: '../inputs/mocks/aqf_redstone_tokens.json',
  },
  states: {
    otherCalls: {
      '385aee1b': { '#CompositeRedstone_1': '#PriceFeed_0' },
      ab0ca0e1: { '#CompositeRedstone_1': '#PriceFeed_1' },
      '313ce567': { '#PriceFeed_0': '8' },
      feaf968c: { '#CompositeRedstone_1': '6' },
    },
  },
  blocks: {
    '26': {
      calls: {
        others: {
          feaf968c: {
            '#YearnFeed_1': utils.bigIntTopic(1, 8),
            '#CurvePriceFeed_1': utils.bigIntTopic(10000, 8),
            '#SingleAsset_1': utils.bigIntTopic(1, 8),
            '#Redstone_1': utils.bigIntTopic(20, 8),
            '#PriceFeed_1': utils.bigIntTopic(3000, 8),
          },
        },
      },
    },
  },
}
