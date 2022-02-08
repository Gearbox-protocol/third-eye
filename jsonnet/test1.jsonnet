{
  mocks: {
    syncAdapters: 'mocks/syncAdapter1.json',
  },
  blocks: {
    '1': {
      events: [{
        address: '@CreditManager_1',
        topics: [
          'OpenCreditAccount(address,address,address,uint256,uint256,uint256)',
          '#User_1',
          '@User_1',
          '#Account_1',
        ],
        data: [
          'bigint:1000000000',
          'bigint:4000000000',
          'bigint:0',
        ],
        txHash: '!#Hash_1',
      }],
      calls:
        {
          pools: [],
          accounts: [],
          cms: [],
        },
    },
  },
}
