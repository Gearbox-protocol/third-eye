{
  '7': {
    blockNum: 7,
    css: [
      {
        balance: {
          '#Token_1': {
            BI: '2000000000',
            F: 2000,
            isEnabled: true,
            isForbidden: false,  // changed
            ind: 0,
          },
          '#Token_2': {
            BI: '100000000000000000',
            F: 0.1,
            isEnabled: false,
            isForbidden: false,  // changed
            ind: 2,
          },
          '#Token_3': {
            BI: '3000000000000000000',
            F: 3,
            isEnabled: true,
            isForbidden: false,  // changed
            ind: 1,
          },
        },
        blockNum: 7,
        borrowedAmount: 5000,
        borrowedAmountBI: '5000000000',
        borrower: '#User_1',
        collateralInUSD: 6500,
        collateralInUnderlying: 6500,
        instCollateralUSD: 6500,
        instCollateralUnderlying: 6500,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '15600',
        sessionId: '#Account_1_3_3',
        totalValue: 9500,
        totalValueBI: '9500000000',
        extraQuotaAPY: 0,
        quotaFees: '0',
      },
    ],
    timestamp: 604800,
  },
  '8': {
    accountOperations: [
      {
        action: 'CloseCreditAccount(address,address,uint256)',
        adapterCall: false,
        args: {
          _order: [
            'owner',
            'to',
            'remainingFunds',
          ],
          owner: '#User_1',
          remainingFunds: '4500000000',
          to: '#User_1',
        },
        blockNum: 8,
        borrower: '#User_1',
        dapp: '#CreditManager_1',
        logId: 0,
        sessionId: '#Account_1_3_3',
        transfers: {
          '#Token_1': 4500000000,
        },
        txHash: '#Hash_12',
      },
    ],
    blockNum: 8,
    cmStats: [
      {
        Address: '#CreditManager_1',
        // AvailableLiquidity: 1000,
        // AvailableLiquidityBI: '1000000000',
        BlockNum: 8,
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
        amount: 5000,
        blockNum: 8,
        event: 'Repay',
        logId: 0,
        pool: '#Pool_1',
        sessionId: '#Account_1_3_3',
        txHash: '#Hash_12',
        user: '#User_1',
      },
    ],
    poolStats: [
      {
        availableLiquidity: 1000,
        availableLiquidityBI: '1000000000',
        blockNum: 8,
        baseBorrowAPY: 0,
        baseBorrowAPYBI: '0',
        cumulativeIndexRAY: '1000000000000000000000000000',
        depositAPY: 0,
        depositAPYBI: '0',
        dieselRate: 0,
        dieselRateBI: '0',
        expectedLiquidity: 5000,
        expectedLiquidityBI: '5000000000',
        pool: '#Pool_1',
        totalBorrowed: 5000,
        totalBorrowedBI: '5000000000',
        totalLosses: 0,
        totalProfit: 0,
        uniqueUsers: 0,
        withdrawFee: 0,
      },
    ],
    timestamp: 691200,
  },
}
