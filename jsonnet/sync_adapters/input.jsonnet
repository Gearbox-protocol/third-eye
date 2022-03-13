local utils = import '../utils.libsonnet';
{
  mocks: {
    syncAdapters: 'mocks/syncAdapterInit.json',
  },
  blocks: {
    '3': {
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
          // data compressor
          address: '#AddressProvider_1',
          txHash: '!#Hash_3',
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
          // price oracle
          address: '#AddressProvider_1',
          txHash: '!#Hash_4',
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
          txHash: '!#Hash_5',
          topics: [
            'AddressSet(bytes32,address)',
            utils.strToHex('DATA_COMPRESSOR'),
            '#DC_2',
          ],
          data: [],
        },
      ],
    },
    '5': {
      //
      events: [
        {
          // pool
          address: '#ContractRegister_1',
          txHash: '!#Hash_6',
          topics: [
            'NewPoolAdded(address)',
            '#Pool_1',
          ],
          data: [],
        },
        {
          // credit manager
          address: '#ContractRegister_1',
          txHash: '!#Hash_7',
          topics: [
            'NewCreditManagerAdded(address)',
            '#CreditManager_1',
          ],
          data: [],
        },
      ],
      calls: {
        others: {
          '2495a599': ['#Token_1'],
          '36dda7d5': ['#DieselToken_1'],
          f93f515b: ['#CreditFilter_1'],
          '570a7af2': ['#Pool_1'],
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
  },
}
