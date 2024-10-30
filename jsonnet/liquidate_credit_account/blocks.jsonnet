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
        instCollateralUSD: 4250,
        instCollateralUnderlying: 4250,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '9600',
        sessionId: '#Account_1_3_3',
        totalValue: 5750,
        totalValueBI: '5750000000',
        quotaFees: '0',
      },
    ],
    priceFeeds: [
      {
        blockNum: 7,
        feed: '#ChainlinkPriceFeed_1',
        price: 0.0008,
        priceBI: '800000000000000',
        roundId: 2,
      },
    ],
    timestamp: 604800,
  },
  '8': {
    accountOperations: [
      {
        action: 'LiquidateCreditAccount(address,address,uint256)',
        adapterCall: false,
        args: {
          _order: [
            'owner',
            'liquidator',
            'remainingFunds',
          ],
          liquidator: '#User_2',
          owner: '#User_1',
          remainingFunds: '640750000',
        },
        blockNum: 8,
        borrower: '#User_1',
        dapp: '#CreditManager_1',
        logId: 0,
        sessionId: '#Account_1_3_3',
        transfers: {
          '#Token_1': 640750000,
        },
        txHash: '#Hash_13',
      },
    ],
    blockNum: 8,
    cmStats: [
      {
        Address: '#CreditManager_1',
        // AvailableLiquidity: 6000,
        // AvailableLiquidityBI: '6000000000',
        BlockNum: 8,
        CumulativeBorrowed: 0,
        ID: 0,
        OpenedAccountsCount: 0,
        TotalBorrowed: 0,
        TotalBorrowedBI: '0',
        TotalClosedAccounts: 0,
        TotalLiquidatedAccounts: 1,
        TotalLosses: 0,
        TotalLossesBI: '0',
        TotalOpenedAccounts: 1,
        TotalProfit: 109.25,
        TotalProfitBI: '109250000',
        TotalRepaid: 5109.25,  // has feeliquidation
        TotalRepaidAccounts: 0,
        TotalRepaidBI: '5109250000',
        UniqueUsers: 0,
      },
    ],
    poolLedgers: [
      {
        amount: 5109.25,
        blockNum: 8,
        event: 'Repay',
        logId: 0,
        pool: '#Pool_1',
        sessionId: '#Account_1_3_3',
        txHash: '#Hash_13',
        user: '#User_1',
      },
    ],
    poolStats: [
      {
        availableLiquidity: 6000,
        availableLiquidityBI: '6000000000',
        blockNum: 8,
        baseBorrowAPY: 0,
        baseBorrowAPYBI: '0',
        cumulativeIndexRAY: '1000000000000000000000000000',
        depositAPY: 0,
        depositAPYBI: '0',
        dieselRate: 0,
        dieselRateBI: '0',
        expectedLiquidity: 6000,
        expectedLiquidityBI: '6000000000',
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
