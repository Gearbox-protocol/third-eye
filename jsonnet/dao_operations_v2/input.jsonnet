local utils = import '../utils.libsonnet';
local initialAmount = 1000;
local borrowedAmount = 4000;
{
  mocks: {
    syncAdapters: 'mocks/syncAdapterV2.json',
    tokens: '../inputs/mocks/tokens.json',
  },
  states: {
    otherCalls: {
      '54fd4d50': { '#CreditConfigurator_2': '2' },
    },
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
  },
  blocks: {
    '2': {
      calls: {  // any of the below call from blocknumber 2 will return listed response
        others: {
          '2f7a1881': { '#CreditManager_1': '#CreditFacade_1' },
          f9aa028a: { '#CreditManager_1': '#CreditConfigurator_1' },
        },
      },
    },
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
            'TokenLiquidationThresholdUpdated(address,uint16)',
            '#Token_1',
          ],
          data: [
            'uint16:7500',
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
            'TokenLiquidationThresholdUpdated(address,uint16)',
            '#Token_1',
          ],
          data: [
            'uint16:8000',
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
        // {
        //   // chidelay and hfcheck
        //   address: '#CreditConfigurator_1',
        //   topics: [
        //     'FastCheckParametersUpdated(uint256,uint256)',
        //   ],
        //   data: [
        //     'bigint:7500',
        //     'bigint:7500',
        //   ],
        //   txHash: '!#Hash_8',
        // },
        {
          // credit manager on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'PriceOracleUpgraded(address)',
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
          txHash: '!#Hash_12',
        },
        {
          // credit manager on usdc
          address: '#CreditManager_1',
          topics: [
            'NewConfigurator(address)',
            '#CreditConfigurator_2',
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
      calls: {
        cms: [{
          address: '#CreditManager_1',
          isWETH: false,
          minAmount: utils.bigInt(1000, 6),
          maxAmount: utils.bigInt(5000, 6),
          availableLiquidity: utils.bigInt(1000, 6),
          borrowRate: '0',
        }],
      },
    },
    '6': {  // new v2 events
      events: [
        {
          address: '#CreditConfigurator_2',
          topics: [
            'PriceOracleUpgraded(address)',
            '#PriceOracle_1',
          ],
          txHash: '!#Hash_15',
        },
        {
          address: '#CreditConfigurator_2',
          topics: [
            'IncreaseDebtForbiddenModeChanged(bool)',
          ],
          data: [
            // bool
            'bool:1',
          ],
          txHash: '!#Hash_16',
        },
        {
          address: '#CreditConfigurator_2',
          topics: [
            'ExpirationDateUpdated(uint40)',
          ],
          data: [
            'bigint:123456789',
          ],
          txHash: '!#Hash_16',
        },
        {
          address: '#CreditConfigurator_2',
          topics: [
            'MaxEnabledTokensUpdated(uint8)',
          ],
          data: [
            'uint8:10',
          ],
          txHash: '!#Hash_16',
        },
        {
          address: '#CreditConfigurator_2',
          topics: [
            'LimitPerBlockUpdated(uint128)',
          ],
          data: [
            'bigint:10',
          ],
          txHash: '!#Hash_16',
        },
        {
          address: '#CreditConfigurator_2',
          topics: [
            'AddedToUpgradeable(address)',
            '#Admin_1',
          ],
          txHash: '!#Hash_17',
        },
        {
          address: '#CreditConfigurator_2',
          topics: [
            'RemovedFromUpgradeable(address)',
            '#Admin_1',
          ],
          txHash: '!#Hash_17',
        },
        {
          address: '#CreditConfigurator_2',
          topics: [
            'EmergencyLiquidatorAdded(address)',
          ],
          data: [
            '#Emergencyiquidator_1',
          ],
          txHash: '!#Hash_18',
        },
        {
          address: '#CreditConfigurator_2',
          topics: [
            'EmergencyLiquidatorRemoved(address)',
          ],
          data: [
            '#Emergencyiquidator_1',
          ],
          txHash: '!#Hash_18',
        },
        {
          // adapter forbidden
          address: '#CreditConfigurator_2',
          topics: [
            'AdapterForbidden(address)',
            '#Adapter_1',
          ],
          txHash: '!#Hash_18',
        },
      ],
      calls: {
        cms: [{
          address: '#CreditManager_1',
          isWETH: false,
          minAmount: utils.bigInt(1000, 6),
          maxAmount: utils.bigInt(5000, 6),
          availableLiquidity: utils.bigInt(1000, 6),
          borrowRate: '0',
        }],
      },
    },
  },
}
