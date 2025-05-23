{
  '8': {
    blockNum: 8,
    css: [
      {
        balance: {
          '#Token_1': {
            BI: '1500000000',
            F: 1500,
            isEnabled: true,
            isForbidden: false,  // changed
            ind: 0,
          },
          '#Token_3': {
            BI: '3000000000000000000',
            F: 3,
            isEnabled: true,
            isForbidden: false,  // changed
            ind: 1,
          },
          // '#Token_2': {
          //   BI: '100000000000000000',
          //   F: 0.1,
          //   isEnabled: false,
          //   isForbidden: false, // changed
          //   ind: 2,
          // },
        },
        blockNum: 8,
        borrowedAmount: 4500,
        borrowedAmountBI: '4500000000',
        borrower: '#User_1',
        collateralInUSD: 6500,
        collateralInUnderlying: 6500,
        instCollateralUSD: 6500,
        instCollateralUnderlying: 6500,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '16333',
        sessionId: '#Account_1_4_0',
        totalValue: 9000,
        totalValueBI: '9000000000',
        quotaFees: '0',
      },
    ],
    timestamp: 691200,
  },
  '9': {
    accountOperations: [
      {
        action: 'CloseCreditAccount(address,address)',
        adapterCall: false,
        args: {
          _order: [
            'borrower',
            'to',
          ],
          borrower: '#User_1',
          to: '#User_2',
          remainingFunds: '5500000000',
        },
        blockNum: 9,
        borrower: '#User_1',
        dapp: '#CreditManager_1',
        logId: 3,
        multicalls: [
          {
            action: 'swapExactTokensForTokens(uint256,uint256,address[],address,uint256)',
            adapterCall: true,
            args: {
              '': '#Account_1',
              _order: [
                'amountIn',
                'amountOutMin',
                'path',
                '',
                'deadline',
              ],
              amountIn: 'bigint:2000000000000000000',
              amountOutMin: 'bigint:1000000000000000000',
              deadline: 0,
              path: [
                '#Token_3',
                '#Token_1',
              ],
            },
            blockNum: 9,
            borrower: '#User_1',
            dapp: '#Uniswapv2_1',
            logId: 1,
            sessionId: '#Account_1_4_0',
            transfers: {
              '#Token_1': 4000000000,
              '#Token_3': -2000000000000000000,
            },
            txHash: '#Hash_12',
          },
        ],
        sessionId: '#Account_1_4_0',
        transfers: {
          '#Token_1': 1000000000,
          '#Token_2': 100000000000000000,
          '#Token_3': 1000000000000000000,
        },
        txHash: '#Hash_12',
      },
    ],
    blockNum: 9,
    cmStats: [
      {
        Address: '#CreditManager_1',
        // AvailableLiquidity: 10000,
        // AvailableLiquidityBI: '10000000000',
        BlockNum: 9,
        CumulativeBorrowed: 0,
        ID: 0,
        OpenedAccountsCount: 0,
        TotalBorrowed: 0,
        TotalBorrowedBI: '0',
        TotalClosedAccounts: 1,
        TotalLiquidatedAccounts: 0,
        TotalLosses: 0,
        TotalLossesBI: '0',
        TotalOpenedAccounts: 1,
        TotalProfit: 0,
        TotalProfitBI: '0',
        TotalRepaid: 5000,
        TotalRepaidAccounts: 0,
        TotalRepaidBI: '5000000000',
        UniqueUsers: 0,
      },
    ],
    poolLedgers: [
      {
        amount: 4500,
        blockNum: 9,
        event: 'Repay',
        logId: 3,
        pool: '#Pool_1',
        sessionId: '#Account_1_4_0',
        txHash: '#Hash_12',
        user: '#User_1',
      },
    ],
    poolStats: [
      {
        availableLiquidity: 10000,
        availableLiquidityBI: '10000000000',
        blockNum: 9,
        baseBorrowAPY: 0,
        baseBorrowAPYBI: '0',
        cumulativeIndexRAY: '1000000000000000000000000000',
        depositAPY: 0,
        depositAPYBI: '0',
        dieselRate: 0,
        dieselRateBI: '0',
        expectedLiquidity: 10000,
        expectedLiquidityBI: '10000000000',
        pool: '#Pool_1',
        totalBorrowed: 0,
        totalBorrowedBI: '0',
        totalLosses: 0,
        totalProfit: 0,
        uniqueUsers: 0,
        withdrawFee: 0,
      },
    ],
    timestamp: 777600,
  },
}
