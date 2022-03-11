local utils = import '../utils.libsonnet';
local initialAmount = 1000;
local borrowedAmount = 4000;
{
  mocks: {
    syncAdapters: 'mocks/syncAdapterV2.json',
  },
  states: {
    oracles: [{
      oracle: '#Oracle_1',
      block: 1,
      feed: '#ChainlinkPriceFeed_1',
    }, {
      oracle: '#Oracle_2',
      block: 1,
      feed: '#ChainlinkPriceFeed_2',
    }, {
      oracle: '#Oracle_4',
      block: 1,
      feed: '#ChainlinkPriceFeed_4',
    }],
  },
  blocks: {
    '3': {
      events: [
        // credit configurator allowed token test
        // - First time token enabled: the dao operation added. not added allowed_tokens table wait for liquiditythrshold event
        // - After liquidityThreshold: another dao operation. Added to allowed_tokens table, allowedTokenstate also updated
        // - Again liquiditythreshold: disable previous allowed token value, generate dao operation and to allowed_tokens table and
        //   allowed token state
        // - On token forbidden: dao operation added. add disabled to allowed tokens state, add to allowed token.
        // - on reenable token: dao operation added. add to allowed_tokens table and add allowed token state.
        {
          // credit configurator enable token
          address: '#CreditConfigurator_1',
          topics: [
            'TokenAllowed(address)',
            '#Token_1',
          ],
          txHash: '!#Hash_1',
        },
        {
          // credit configurator add liquidity threshold
          address: '#CreditConfigurator_1',
          topics: [
            'TokenLiquidationThresholdUpdated(address,uint256)',
            '#Token_1',
          ],
          data: [
            'bigint:7500',
          ],
          txHash: '!#Hash_2',
        },
      ],
    },
    '4': {
      events: [
        {
          // credit configurator updated liquidity threshold
          address: '#CreditConfigurator_1',
          topics: [
            'TokenLiquidationThresholdUpdated(address,uint256)',
            '#Token_1',
          ],
          data: [
            'bigint:8000',
          ],
          txHash: '!#Hash_3',
        },
      ],
    },
    '5': {
      events: [
        {
          // token forbidden
          address: '#CreditConfigurator_1',
          topics: [
            'TokenForbidden(address)',
            '#Token_1',
          ],
          txHash: '!#Hash_4',
        },
        {
          // token reenable
          address: '#CreditConfigurator_1',
          topics: [
            'TokenAllowed(address)',
            '#Token_1',
          ],
          txHash: '!#Hash_5',
        },
        {
          // contract allowed
          address: '#CreditConfigurator_1',
          topics: [
            'ContractAllowed(address,address)',
            '#Protocol_1',
            '#Adapter_1',
          ],
          txHash: '!#Hash_6',
        },
        {
          // contract forbidden
          address: '#CreditConfigurator_1',
          topics: [
            'ContractForbidden(address)',
            '#Protocol_1',
          ],
          txHash: '!#Hash_7',
        },
        {
          // chidelay and hfcheck
          address: '#CreditConfigurator_1',
          topics: [
            'FastCheckParametersUpdated(uint256,uint256)',
          ],
          data: [
            'bigint:7500',
            'bigint:7500',
          ],
          txHash: '!#Hash_8',
        },
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'PriceOracleUpdated(address)',
            '#PriceOracle_2',
          ],
          txHash: '!#Hash_9',
        },
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
          txHash: '!#Hash_11',
        },
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'FeesUpdated(uint256,uint256,uint256)',
          ],
          data: [
            // feeInterest
            utils.bigIntTopic(1000, 0),
            // feeLiquidation
            utils.bigIntTopic(200, 0),
            // liquidationDiscount
            utils.bigIntTopic(9500, 0),
          ],
          txHash: '!#Hash_12',
        },
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'CreditConfiguratorUpgraded(address)',
            '#CreditConfigurator_1',
          ],
          txHash: '!#Hash_13',
        },
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'CreditFacadeUpgraded(address)',
            '#CreditFacade_1',
          ],
          txHash: '!#Hash_14',
        },
      ],
    },
  },
}
