{
  '4': {
    accountOperations: [
      {
        action: 'OpenCreditAccount(address,address,uint256,uint16)',
        adapterCall: false,
        args: {
          _order: [
            'onBehalfOf',
            'creditAccount',
            'borrowAmount',
            'referralCode',
          ],
          amount: '1000000000',
          borrowAmount: 4000000000,
          creditAccount: '#Account_1',
          onBehalfOf: '#User_1',
          referralCode: 0,
        },
        blockNum: 4,
        borrower: '#User_1',
        dapp: '#CreditFacade_1',
        logId: 0,
        multicalls: [
          {
            action: 'AddCollateral(address,address,uint256)',
            adapterCall: false,
            args: {
              _order: [
                'onBehalfOf',
                'token',
                'value',
              ],
              onBehalfOf: '#User_1',
              token: '#Token_1',
              value: 1000000000,
            },
            blockNum: 4,
            borrower: '#User_1',
            dapp: '#CreditFacade_1',
            logId: 2,
            sessionId: '#Account_1_4_0',
            transfers: {
              '#Token_1': 1000000000,
            },
            txHash: '#Hash_1',
          },
        ],
        sessionId: '#Account_1_4_0',
        transfers: {
          '#Token_1': 4000000000,
        },
        txHash: '#Hash_1',
      },
      {
        action: 'MultiCallStarted(address)',
        adapterCall: false,
        blockNum: 4,
        borrower: '#User_1',
        dapp: '#CreditFacade_1',
        logId: 0,
        multicalls: [
          {
            action: 'AddCollateral(address,address,uint256)',
            adapterCall: false,
            args: {
              _order: [
                'onBehalfOf',
                'token',
                'value',
              ],
              onBehalfOf: '#User_1',
              token: '#Token_1',
              value: 1000000000,
            },
            blockNum: 4,
            borrower: '#User_1',
            dapp: '#CreditFacade_1',
            logId: 6,
            sessionId: '#Account_1_4_0',
            transfers: {
              '#Token_1': 1000000000,
            },
            txHash: '#Hash_1',
          },
        ],
        sessionId: '#Account_1_4_0',
        txHash: '#Hash_1',
      },
    ],
    blockNum: 4,
    cmStats: [
      {
        Address: '#CreditManager_1',
        // AvailableLiquidity: 5000,
        // AvailableLiquidityBI: '5000000000',
        BlockNum: 4,
        CumulativeBorrowed: 0,
        ID: 0,
        OpenedAccountsCount: 1,
        TotalBorrowed: 4000,
        TotalBorrowedBI: '4000000000',
        TotalClosedAccounts: 0,
        TotalLiquidatedAccounts: 0,
        TotalLosses: 0,
        TotalLossesBI: '0',
        TotalOpenedAccounts: 1,
        TotalProfit: 0,
        TotalProfitBI: '0',
        TotalRepaid: 0,
        TotalRepaidAccounts: 0,
        TotalRepaidBI: '0',
        UniqueUsers: 0,
      },
    ],
    css: [
      {
        balance: {
          '#Token_1': {
            BI: '6000000000',
            F: 6000,
            ind: 0,
            isForbidden: false,  // changed
            isEnabled: true,
          },
        },
        blockNum: 4,
        borrowedAmount: 4000,
        borrowedAmountBI: '4000000000',
        borrower: '#User_1',
        collateralInUSD: 2000,
        collateralInUnderlying: 2000,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '13500',
        sessionId: '#Account_1_4_0',
        totalValue: 6000,
        totalValueBI: '6000000000',
      },
    ],
    poolLedgers: [
      {
        amount: 4000,
        blockNum: 4,
        event: 'Borrow',
        logId: 0,
        pool: '#Pool_1',
        sessionId: '#Account_1_4_0',
        txHash: '#Hash_1',
        user: '#User_1',
      },
    ],
    poolStats: [
      {
        availableLiquidity: 5000,
        availableLiquidityBI: '5000000000',
        blockNum: 4,
        borrowAPY: 0,
        borrowAPYBI: '0',
        cumulativeIndexRAY: '1000000000000000000000000000',
        depositAPY: 0,
        depositAPYBI: '0',
        dieselRate: 0,
        dieselRateBI: '0',
        expectedLiquidity: 9000,
        expectedLiquidityBI: '9000000000',
        pool: '#Pool_1',
        totalBorrowed: 4000,
        totalBorrowedBI: '4000000000',
        totalLosses: 0,
        totalProfit: 0,
        uniqueUsers: 0,
        withdrawFee: 0,
      },
    ],
    timestamp: 345600,
  },
  '5': {
    blockNum: 5,
    css: [
      {
        balance: {
          '#Token_1': {
            BI: '7000000000',
            F: 7000,
            ind: 0,
            isForbidden: false,  // changed
            isEnabled: true,
          },
        },
        blockNum: 5,
        borrowedAmount: 4000,
        borrowedAmountBI: '4000000000',
        borrower: '#User_1',
        collateralInUSD: 3000,
        collateralInUnderlying: 3000,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '15750',
        sessionId: '#Account_1_4_0',
        totalValue: 7000,
        totalValueBI: '7000000000',
      },
    ],
    timestamp: 432000,
  },
  '6': {
    accountOperations: [
      {
        action: 'AddCollateral(address,address,uint256)',
        adapterCall: false,
        args: {
          _order: [
            'onBehalfOf',
            'token',
            'value',
          ],
          onBehalfOf: '#User_2',
          token: '#Token_1',
          value: 1000000000,
        },
        blockNum: 6,
        borrower: '#User_2',
        dapp: '#CreditFacade_1',
        logId: 7,
        sessionId: '#Account_2_6_4',
        transfers: {
          '#Token_1': 1000000000,
        },
        txHash: '#Hash_2',
      },
      {
        action: 'OpenCreditAccount(address,address,uint256,uint16)',
        adapterCall: false,
        args: {
          _order: [
            'onBehalfOf',
            'creditAccount',
            'borrowAmount',
            'referralCode',
          ],
          amount: '1000000000',
          borrowAmount: 4000000000,
          creditAccount: '#Account_2',
          onBehalfOf: '#User_2',
          referralCode: 0,
        },
        blockNum: 6,
        borrower: '#User_2',
        dapp: '#CreditFacade_1',
        logId: 4,
        multicalls: [
          {
            action: 'AddCollateral(address,address,uint256)',
            adapterCall: false,
            args: {
              _order: [
                'onBehalfOf',
                'token',
                'value',
              ],
              onBehalfOf: '#User_2',
              token: '#Token_1',
              value: 1000000000,
            },
            blockNum: 6,
            borrower: '#User_2',
            dapp: '#CreditFacade_1',
            logId: 6,
            sessionId: '#Account_2_6_4',
            transfers: {
              '#Token_1': 1000000000,
            },
            txHash: '#Hash_2',
          },
        ],
        sessionId: '#Account_2_6_4',
        transfers: {
          '#Token_1': 4000000000,
        },
        txHash: '#Hash_2',
      },
      {
        action: 'CloseCreditAccount(address,address)',
        adapterCall: false,
        args: {
          _order: [
            'borrower',
            'to',
          ],
          borrower: '#User_1',
          remainingFunds: '3000000000',
          to: '#User_1',
        },
        blockNum: 6,
        borrower: '#User_1',
        dapp: '#CreditManager_1',
        logId: 3,
        multicalls: [
          {
            action: 'AddCollateral(address,address,uint256)',
            adapterCall: false,
            args: {
              _order: [
                'onBehalfOf',
                'token',
                'value',
              ],
              onBehalfOf: '#User_1',
              token: '#Token_1',
              value: 1000000000,
            },
            blockNum: 6,
            borrower: '#User_1',
            dapp: '#CreditFacade_1',
            logId: 1,
            sessionId: '#Account_1_4_0',
            transfers: {
              '#Token_1': 1000000000,
            },
            txHash: '#Hash_2',
          },
        ],
        sessionId: '#Account_1_4_0',
        transfers: {
          '#Token_1': 3000000000,
        },
        txHash: '#Hash_2',
      },
    ],
    blockNum: 6,
    cmStats: [
      {
        Address: '#CreditManager_1',
        // AvailableLiquidity: 5000,
        // AvailableLiquidityBI: '5000000000',
        BlockNum: 6,
        CumulativeBorrowed: 0,
        ID: 0,
        OpenedAccountsCount: 1,
        TotalBorrowed: 8000,
        TotalBorrowedBI: '8000000000',
        TotalClosedAccounts: 1,
        TotalLiquidatedAccounts: 0,
        TotalLosses: 0,
        TotalLossesBI: '0',
        TotalOpenedAccounts: 2,
        TotalProfit: 0,
        TotalProfitBI: '0',
        TotalRepaid: 4000,
        TotalRepaidAccounts: 0,
        TotalRepaidBI: '4000000000',
        UniqueUsers: 0,
      },
    ],
    css: [
      {
        balance: {
          '#Token_1': {
            BI: '6000000000',
            F: 6000,
            ind: 0,
            isForbidden: false,  // changed
            isEnabled: true,
          },
        },
        blockNum: 6,
        borrowedAmount: 4000,
        borrowedAmountBI: '4000000000',
        borrower: '#User_2',
        collateralInUSD: 2000,
        collateralInUnderlying: 2000,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '13500',
        sessionId: '#Account_2_6_4',
        totalValue: 6000,
        totalValueBI: '6000000000',
      },
    ],
    poolLedgers: [
      {
        amount: 4000,
        blockNum: 6,
        event: 'Borrow',
        logId: 4,
        pool: '#Pool_1',
        sessionId: '#Account_2_6_4',
        txHash: '#Hash_2',
        user: '#User_2',
      },
      {
        amount: 4000,
        blockNum: 6,
        event: 'Repay',
        logId: 3,
        pool: '#Pool_1',
        sessionId: '#Account_1_4_0',
        txHash: '#Hash_2',
        user: '#User_1',
      },
    ],
    poolStats: [
      {
        availableLiquidity: 5000,
        availableLiquidityBI: '5000000000',
        blockNum: 6,
        borrowAPY: 0,
        borrowAPYBI: '0',
        cumulativeIndexRAY: '1000000000000000000000000000',
        depositAPY: 0,
        depositAPYBI: '0',
        dieselRate: 0,
        dieselRateBI: '0',
        expectedLiquidity: 9000,
        expectedLiquidityBI: '9000000000',
        pool: '#Pool_1',
        totalBorrowed: 4000,
        totalBorrowedBI: '4000000000',
        totalLosses: 0,
        totalProfit: 0,
        uniqueUsers: 0,
        withdrawFee: 0,
      },
    ],
    timestamp: 518400,
  },
}
