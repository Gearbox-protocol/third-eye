{
  accountOperations: [
    {
      action: 'OpenCreditAccount(address,address,address,uint256,uint256,uint256)',
      adapterCall: false,
      args: {
        _order: [
          'sender',
          'onBehalfOf',
          'creditAccount',
          'amount',
          'borrowAmount',
          'referralCode',
        ],
        amount: '1000000000',
        userFunds: '1000000000',
        borrowAmount: '4000000000',
        creditAccount: '#Account_1',
        onBehalfOf: '#User_1',
        referralCode: '0',
        sender: '#User_1',
      },
      blockNum: 3,
      borrower: '#User_1',
      dapp: '#CreditManager_1',
      logId: 0,
      sessionId: '#Account_1_3_0',
      transfers: {
        '#Token_1': 5000000000,
      },
      txHash: '#Hash_1',
    },
  ],
  allowedTokens: [
    {
      BlockNumber: 3,
      Configurator: '#CreditFilter_1',
      CreditManager: '#CreditManager_1',
      DisableBlock: 0,
      LiquidityThreshold: '9000',
      LogID: 0,
      Token: '#Token_1',
    },
  ],
  blockNum: 3,
  cmStats: [
    {
      Address: '#CreditManager_1',
      BlockNum: 3,
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
          BI: '5000000000',
          F: 5000,
          ind: 0,
          isEnabled: true,
          isForbidden: false,
        },
      },
      blockNum: 3,
      borrowedAmount: 4000,
      borrowedAmountBI: '4000000000',
      borrower: '#User_1',
      collateralInUSD: 1000,
      collateralInUnderlying: 1000,
      instCollateralUSD: 1000,
      instCollateralUnderlying: 1000,
      cumulativeIndexAtOpen: '1000000000000000000000000000',
      healthFactor: '11250',
      quotaFees: '0',
      sessionId: '#Account_1_3_0',
      totalValue: 5000,
      totalValueBI: '5000000000',
      extraQuotaAPY: 0,
    },
  ],
  daoOperations: [
    {
      Args: {
        feeInterest: [
          0,
          1000,
        ],
        feeLiquidation: [
          0,
          200,
        ],
        liquidationDiscount: [
          0,
          9500,
        ],
        maxAmount: [
          '0',
          '5000000000',
        ],
        maxLeverage: [
          '0',
          '400000000',
        ],
        minAmount: [
          '0',
          '1000000000',
        ],
        token: '#Token_1',
      },
      BlockNumber: 3,
      Contract: '#CreditManager_1',
      LogID: 4,
      TxHash: '#Hash_3',
      Type: 21,
    },
    {
      Args: {
        creditManager: '#CreditManager_1',
        liquidityThreshold: '9000',
        prevLiquidationThreshold: '0',
        token: '#Token_1',
      },
      BlockNumber: 3,
      Contract: '#CreditFilter_1',
      LogID: 1,
      TxHash: '#Hash_3',
      Type: 0,
    },
  ],
  params: [
    {
      BlockNum: 3,
      CreditManager: '#CreditManager_1',
      EmergencyLiqDiscount: 0,
      FeeInterest: 1000,
      FeeLiquidation: 200,
      FeeLiquidationExpired: 0,
      LiquidationDiscount: 9500,
      LiquidationDiscountExpired: 0,
      MaxAmount: '5000000000',
      MaxLeverage: '400000000',
      MinAmount: '1000000000',
    },
  ],
  poolLedgers: [
    {
      amount: 4000,
      blockNum: 3,
      event: 'Borrow',
      logId: 0,
      pool: '#Pool_1',
      sessionId: '#Account_1_3_0',
      txHash: '#Hash_1',
      user: '#User_1',
    },
  ],
  poolStats: [
    {
      availableLiquidity: 1000,
      availableLiquidityBI: '1000000000',
      blockNum: 3,
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
      totalBorrowed: 4000,
      totalBorrowedBI: '4000000000',
      totalLosses: 0,
      totalProfit: 0,
      uniqueUsers: 0,
      withdrawFee: 0,
    },
  ],
  priceFeeds: [
    {
      blockNum: 3,
      feed: '#ChainlinkPriceFeed_1',
      price: 1e-18,
      priceBI: '1',
      roundId: 300000000000000,
    },
  ],
  timestamp: 259200,
}
