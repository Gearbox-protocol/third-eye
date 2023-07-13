local utils = import '../utils.libsonnet';
// tests
// check priceOracle v1 and v2's  oracle are working together. Block 6
// check if upgrade working properly from v1 to v2 for creditmanager/priceoracle/chainlinkPriceFeed/dc: 5 and 6
{
  mocks: {
    syncAdapters: 'mocks/syncAdapterInit.json',
    tokens: '../inputs/mocks/tokens.json',
  },
  states: {
    oracles: {
      '#Oracle_1': [{
        oracle: '#Oracle_1',
        block: 1,
        feed: '#ChainlinkPriceFeed_1',
      }],
      '#Oracle_5': [{
        oracle: '#Oracle_5',
        block: 1,
        feed: '#ChainlinkPriceFeed_5',
      }],
    },
  },
  blocks: {
    // v1 blocks: 3 and 4
    // v2 blocks: 5 and 6
    '3': {
      calls: {
        others: {
          '54fd4d50': {
            '#ACL_1': '1',
            '#ContractRegister_1': '1',
            '#DC_1': '1',
            '#PriceOracle_1': '1',
          },
        },
      },
      //
      events: [
        {
          // contract register
          address: '#AddressProvider_1',
          txHash: '!#Hash_1',
          topics: [
            'AddressSet(bytes32,address)',
            utils.strToHex('ACL'),
            '#ACL_1',
          ],
          data: [],
        },
        {
          // contract register
          address: '#AddressProvider_1',
          txHash: '!#Hash_2',
          topics: [
            'AddressSet(bytes32,address)',
            utils.strToHex('CONTRACTS_REGISTER'),
            '#ContractRegister_1',
          ],
          data: [],
        },
        {
          // price oracle
          address: '#AddressProvider_1',
          txHash: '!#Hash_3',
          topics: [
            'AddressSet(bytes32,address)',
            utils.strToHex('PRICE_ORACLE'),
            '#PriceOracle_1',
          ],
          data: [],
        },
        {
          // price oracle
          address: '#PriceOracle_1',
          txHash: '!#Hash_3',
          topics: [
            'NewPriceFeed(address,address)',
            '#Token_1',
            '#Oracle_1',
          ],
          data: [],
        },
        {
          // data compressor
          address: '#AddressProvider_1',
          txHash: '!#Hash_4',
          topics: [
            'AddressSet(bytes32,address)',
            utils.strToHex('DATA_COMPRESSOR'),
            '#DC_1',
          ],
          data: [],
        },
      ],
    },
    '4': {
      //
      events: [
        {
          // pool
          address: '#ContractRegister_1',
          txHash: '!#Hash_5',
          topics: [
            'NewPoolAdded(address)',
            '#Pool_1',
          ],
          data: [],
        },
        {
          // credit manager
          address: '#ContractRegister_1',
          txHash: '!#Hash_6',
          topics: [
            'NewCreditManagerAdded(address)',
            '#CreditManager_1',
          ],
          data: [],
        },
      ],
      calls: {
        others: {
          '2495a599': { '#CreditManager_1': '#Token_1', '#Pool_1': '#Token_1' },
          '36dda7d5': { '#Pool_1': '#DieselToken_1' },
          f93f515b: { '#CreditManager_1': '#CreditFilter_1' },
          '570a7af2': { '#CreditManager_1': '#Pool_1' },
          '54fd4d50': {
            '#CreditManager_1': '1',
            '#Pool_1': '1',
            '#CreditFilter_1': '1',
          },
        },
        pools: [{
          address: '#Pool_1',
          totalBorrowed: utils.bigInt(0, 6),
          expectedLiquidity: utils.bigInt(6000, 6),
          availableLiquidity: utils.bigInt(6000, 6),
          depositAPY: utils.bigInt(0),
          borrowAPY: utils.bigInt(0),
          dieselRate: utils.bigInt(0),
          withdrawFee: '0',
          linearCumulativeIndex: utils.bigInt(1, 27),
        }],
      },
    },
    '5': {
      //
      events: [
        {
          // price oracle
          address: '#AddressProvider_1',
          txHash: '!#Hash_7',
          topics: [
            'AddressSet(bytes32,address)',
            utils.strToHex('PRICE_ORACLE'),
            '#PriceOracle_2',
          ],
          data: [],
        },
        {
          // data compressor
          address: '#AddressProvider_1',
          txHash: '!#Hash_8',
          topics: [
            'AddressSet(bytes32,address)',
            utils.strToHex('DATA_COMPRESSOR'),
            '#DC_2',
          ],
          data: [],
        },
      ],
      calls: {
        others: {
          '54fd4d50': {
            '#PriceOracle_2': '2',  // version 2 for price oracle
            '#DC_2': '2',
          },
        },
      },
    },
    '6': {
      events: [
        {
          // chainlinkPricefeed
          address: '#ChainlinkPriceFeed_1',
          txHash: '!#Hash_8',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // 8
            utils.bigIntTopic(0.0004, 18),
            // roundid
            utils.bigIntTopic(1, 0),
          ],
          data: [],
        },
        {
          // price oracle
          address: '#PriceOracle_2',
          txHash: '!#Hash_9',
          topics: [
            'NewPriceFeed(address,address)',
            '#Token_1',
            '#Oracle_5',
          ],
          data: [],
        },
        {
          // chainlinkPricefeed
          address: '#ChainlinkPriceFeed_5',
          txHash: '!#Hash_10',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // 8
            utils.bigIntTopic(1, 8),
            // roundid
            utils.bigIntTopic(1, 0),
          ],
          data: [],
        },
      ],
    },
    '7': {
      //
      events: [
        // version doesn't matter for pool v2
        {
          // pool
          address: '#ContractRegister_1',
          txHash: '!#Hash_9',
          topics: [
            'NewPoolAdded(address)',
            '#Pool_2',
          ],
          data: [],
        },
        {
          // credit manager
          address: '#ContractRegister_1',
          txHash: '!#Hash_10',
          topics: [
            'NewCreditManagerAdded(address)',
            '#CreditManager_2',
          ],
          data: [],
        },
      ],
      calls: {
        others: {
          '2495a599': {
            '#Pool_2': '#Token_1',
            '#CreditManager_2': '#Token_1',
          },  // underlyingTOken used by pool
          '6f307dc3': { '#CreditManager_2': '#Token_1' },  // underlying used by credit manager version 2
          '36dda7d5': { '#Pool_2': '#DieselToken_1' },  // get dieeltoken on pool
          '570a7af2': { '#CreditManager_2': '#Pool_2' },  // poolservice
          '2f7a1881': { '#CreditManager_2': '#CreditFacade_2' },  // creditfacade
          f9aa028a: { '#CreditManager_2': '#CreditConfigurator_2' },  // getcreditconfigurator
          '54fd4d50': {  // version
            '#CreditManager_2': '2',
            '#CreditConfigurator_2': '2',
            '#Pool_2': '2',
          },
        },
        pools: [{
          address: '#Pool_2',
          totalBorrowed: utils.bigInt(0, 6),
          expectedLiquidity: utils.bigInt(6000, 6),
          availableLiquidity: utils.bigInt(6000, 6),
          depositAPY: utils.bigInt(0),
          borrowAPY: utils.bigInt(0),
          dieselRate: utils.bigInt(0),
          withdrawFee: '0',
          linearCumulativeIndex: utils.bigInt(1, 27),
        }],
      },
    },
  },
}
