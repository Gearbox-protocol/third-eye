local utils = import '../utils.libsonnet';
local initialAmount = 1000;
local borrowedAmount = 4000;
{
  mocks: {
    syncAdapters: 'mocks/syncAdapterV2.json',
    tokens: '../inputs/mocks/tokens.json',
  },
  states: {
    oracles: {
      '#Oracle_0': [{  // for v1
        oracle: '#Oracle_0',
        block: 1,
        feed: '#ChainlinkPriceFeed_0',
      }],  // for v2
      '#Oracle_1': [{
        oracle: '#Oracle_1',
        block: 1,
        feed: '#ChainlinkPriceFeed_1',
      }],
      '#Oracle_2': [{
        oracle: '#Oracle_2',
        block: 1,
        feed: '#ChainlinkPriceFeed_2',
      }],
      '#Oracle_3': [{
        oracle: '#Oracle_3',
        block: 1,
        feed: '#ChainlinkPriceFeed_3',
      }],
      '#Oracle_4': [{
        oracle: '#Oracle_4',
        block: 1,
        feed: '#ChainlinkPriceFeed_4',
      }],
    },
    otherCalls: {
      '54fd4d50': { '#AddressProvider_1': '1' },
    },
  },
  blocks: {
    '3': {
      events: [
        // if maxcumulativeloss and reset are handled.
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'NewMaxCumulativeLoss(uint128)',
          ],
          data: [
            // 100k loss in usdc
            utils.bigIntTopic(100000, 6),
          ],
          txHash: '!#Hash_5',
        },
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'CumulativeLossReset()',
          ],
          txHash: '!#Hash_4',
        },
        // if emergency liq discount is getting set.
        // if limits updated and fees updated after emergencyliqdiscount are working current.
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'NewEmergencyLiquidationDiscount(uint16)',
          ],
          data: [
            // emergency liq discount
            'uint16:9600',
          ],
          txHash: '!#Hash_1',
        },
      ],
    },
    '4': {
      events: [
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'LimitsUpdated(uint256,uint256)',
          ],
          data: [
            // minAnount
            utils.bigIntTopic(1000, 6),
            // maxAmount
            utils.bigIntTopic(5000, 6),
          ],
          txHash: '!#Hash_2',
        },
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'FeesUpdated(uint16,uint16,uint16,uint16,uint16)',
          ],
          data: [
            // feeInterest
            'uint16:1000',
            // feeLiquidation
            'uint16:200',
            // liquidationPremium
            'uint16:500',
            // feeLiquidationexpired
            'uint16:300',
            // liquidationPremiumexpired
            'uint16:600',
          ],
          txHash: '!#Hash_3',
        },
      ],
    },
  },
}
