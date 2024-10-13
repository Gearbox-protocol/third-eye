{
  '3': {
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
        logId: 3,
        sessionId: '#Account_1_3_3',
        transfers: {
          '#Token_1': 5000000000,
        },
        txHash: '#Hash_4',
      },
      {
        action: 'IncreaseBorrowedAmount(address,uint256)',
        adapterCall: false,
        args: {
          _order: [
            'borrower',
            'amount',
          ],
          amount: '1000000000',
          borrower: '#User_1',
        },
        blockNum: 3,
        borrower: '#User_1',
        dapp: '#CreditManager_1',
        logId: 5,
        sessionId: '#Account_1_3_3',
        transfers: {
          '#Token_1': 1000000000,
        },
        txHash: '#Hash_6',
      },
    ],
    allowedTokens: [
      {
        BlockNumber: 3,
        Configurator: '#CreditFilter_1',
        CreditManager: '#CreditManager_1',
        DisableBlock: 0,
        LogID: 0,
        LiquidityThreshold: '9000',
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
        TotalBorrowed: 5000,
        TotalBorrowedBI: '5000000000',
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
            isEnabled: true,
            isForbidden: false,
          },
        },
        blockNum: 3,
        borrowedAmount: 5000,
        borrowedAmountBI: '5000000000',
        borrower: '#User_1',
        collateralInUSD: 1000,
        collateralInUnderlying: 1000,
        instCollateralUSD: 1000,
        instCollateralUnderlying: 1000,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '10800',
        quotaFees: '0',
        sessionId: '#Account_1_3_3',
        totalValue: 6000,
        totalValueBI: '6000000000',
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
        LogID: 1,
        TxHash: '#Hash_2',
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
        LogID: 0,
        TxHash: '#Hash_1',
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
        logId: 3,
        pool: '#Pool_1',
        sessionId: '#Account_1_3_3',
        txHash: '#Hash_4',
        user: '#User_1',
      },
      {
        amount: 1000,
        blockNum: 3,
        event: 'Borrow',
        logId: 5,
        pool: '#Pool_1',
        sessionId: '#Account_1_3_3',
        txHash: '#Hash_6',
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
        expectedLiquidity: 6000,
        expectedLiquidityBI: '6000000000',
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
        price: 0.0004,
        priceBI: '400000000000000',
        roundId: 1,
      },
    ],
    timestamp: 259200,
  },
  '4': {
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
          onBehalfOf: '#User_1',
          token: '#Token_3',
          value: '1000000000000000000',
        },
        blockNum: 4,
        borrower: '#User_1',
        dapp: '#CreditManager_1',
        logId: 1,
        sessionId: '#Account_1_3_3',
        transfers: {
          '#Token_3': 1000000000000000000,
        },
        txHash: '#Hash_9',
      },
    ],
    allowedTokens: [
      {
        BlockNumber: 4,
        Configurator: '#CreditFilter_1',
        CreditManager: '#CreditManager_1',
        DisableBlock: 0,
        LogID: 0,
        LiquidityThreshold: '8000',
        Token: '#Token_3',
      },
    ],
    blockNum: 4,
    cmStats: [
      {
        Address: '#CreditManager_1',
        BlockNum: 4,
        CumulativeBorrowed: 0,
        ID: 0,
        OpenedAccountsCount: 1,
        TotalBorrowed: 5000,
        TotalBorrowedBI: '5000000000',
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
            isEnabled: true,
            isForbidden: false,
          },
          '#Token_3': {
            BI: '1000000000000000000',
            F: 1,
            ind: 1,
            isEnabled: true,
            isForbidden: false,
          },
        },
        blockNum: 4,
        borrowedAmount: 5000,
        borrowedAmountBI: '5000000000',
        borrower: '#User_1',
        collateralInUSD: 3500,
        collateralInUnderlying: 3500,
        instCollateralUSD: 3500,
        instCollateralUnderlying: 3500,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '14800',
        quotaFees: '0',
        sessionId: '#Account_1_3_3',
        totalValue: 8500,
        totalValueBI: '8500000000',
        extraQuotaAPY: 0,
      },
    ],
    daoOperations: [
      {
        Args: {
          creditManager: '#CreditManager_1',
          liquidityThreshold: '8000',
          prevLiquidationThreshold: '0',
          token: '#Token_3',
        },
        BlockNumber: 4,
        Contract: '#CreditFilter_1',
        LogID: 0,
        TxHash: '#Hash_8',
        Type: 0,
      },
    ],
    timestamp: 345600,
  },
  '5': {
    accountOperations: [
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
          amountIn: 'bigint:5000000000',
          amountOutMin: 'bigint:2000000000000000000',
          deadline: 0,
          path: [
            '#Token_1',
            '#Token_3',
          ],
        },
        blockNum: 5,
        borrower: '#User_1',
        dapp: '#Uniswapv2_1',
        logId: 0,
        sessionId: '#Account_1_3_3',
        transfers: {
          '#Token_1': -5000000000,
          '#Token_3': 2000000000000000000,
        },
        txHash: '#Hash_8',
      },
    ],
    blockNum: 5,
    cmStats: [
      {
        Address: '#CreditManager_1',
        BlockNum: 5,
        CumulativeBorrowed: 0,
        ID: 0,
        OpenedAccountsCount: 1,
        TotalBorrowed: 5000,
        TotalBorrowedBI: '5000000000',
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
            BI: '1000000000',
            F: 1000,
            ind: 0,
            isEnabled: true,
            isForbidden: false,
          },
          '#Token_3': {
            BI: '3000000000000000000',
            F: 3,
            ind: 1,
            isEnabled: true,
            isForbidden: false,
          },
        },
        blockNum: 5,
        borrowedAmount: 5000,
        borrowedAmountBI: '5000000000',
        borrower: '#User_1',
        collateralInUSD: 3500,
        collateralInUnderlying: 3500,
        instCollateralUSD: 3500,
        instCollateralUnderlying: 3500,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '13800',
        quotaFees: '0',
        sessionId: '#Account_1_3_3',
        totalValue: 8500,
        totalValueBI: '8500000000',
        extraQuotaAPY: 0,
      },
    ],
    timestamp: 432000,
  },
  '6': {
    accountOperations: [
      {
        action: 'DirectTokenTransfer',
        adapterCall: false,
        args: {
          amount: '1000000000',
          from: '#User_3',
          to: '#Account_1',
        },
        blockNum: 6,
        borrower: '#User_1',
        dapp: '#Token_1',
        logId: 2,
        sessionId: '#Account_1_3_3',
        transfers: {
          '#Token_1': 1000000000,
        },
        txHash: '#Hash_10',
      },
      {
        action: 'DirectTokenTransfer',
        adapterCall: false,
        args: {
          amount: '100000000000000000',
          from: '#User_3',
          to: '#Account_1',
        },
        blockNum: 6,
        borrower: '#User_1',
        dapp: '#Token_2',
        logId: 3,
        sessionId: '#Account_1_3_3',
        transfers: {
          '#Token_2': 100000000000000000,
        },
        txHash: '#Hash_11',
      },
    ],
    allowedTokens: [
      {
        BlockNumber: 6,
        Configurator: '#CreditFilter_1',
        CreditManager: '#CreditManager_1',
        DisableBlock: 0,
        LogID: 0,
        LiquidityThreshold: '9000',
        Token: '#Token_2',
      },
    ],
    blockNum: 6,
    css: [
      {
        balance: {
          '#Token_1': {
            BI: '2000000000',
            F: 2000,
            ind: 0,
            isEnabled: true,
            isForbidden: false,
          },
          '#Token_2': {
            BI: '100000000000000000',
            F: 0.1,
            ind: 2,
            isEnabled: false,
            isForbidden: false,
          },
          '#Token_3': {
            BI: '3000000000000000000',
            F: 3,
            ind: 1,
            isEnabled: true,
            isForbidden: true,
          },
        },
        blockNum: 6,
        borrowedAmount: 5000,
        borrowedAmountBI: '5000000000',
        borrower: '#User_1',
        collateralInUSD: 6500,
        collateralInUnderlying: 6500,
        instCollateralUSD: 6500,
        instCollateralUnderlying: 6500,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '15600',
        quotaFees: '0',
        sessionId: '#Account_1_3_3',
        totalValue: 9500,
        totalValueBI: '9500000000',
        extraQuotaAPY: 0,
      },
    ],
    daoOperations: [
      {
        Args: {
          creditManager: '#CreditManager_1',
          liquidityThreshold: '9000',
          prevLiquidationThreshold: '0',
          token: '#Token_2',
        },
        BlockNumber: 6,
        Contract: '#CreditFilter_1',
        LogID: 0,
        TxHash: '#Hash_1',
        Type: 0,
      },
    ],
    priceFeeds: [
      {
        blockNum: 6,
        feed: '#ChainlinkPriceFeed_2',
        price: 8,
        priceBI: '8000000000000000000',
        roundId: 1,
      },
    ],
    timestamp: 518400,
  },
}
