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
          borrowAmount: 4000000000,
          creditAccount: '#Account_1',
          onBehalfOf: '#User_1',
          referralCode: 0,
          amount: '1000000000',
        },
        blockNum: 4,
        borrower: '#User_1',
        dapp: '#CreditFacade_1',
        depth: 0,
        logId: 0,
        sessionId: '#Account_1_4_0',
        transfers: {
          '#Token_1': 4000000000,
        },
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
            depth: 0,
            logId: 1,
            sessionId: '#Account_1_4_0',
            transfers: {
              '#Token_1': 1000000000,
            },
            txHash: '#Hash_4',
          },
        ],
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
          amount: 1000000000,
          borrower: '#User_1',
        },
        blockNum: 4,
        borrower: '#User_1',
        dapp: '#CreditFacade_1',
        depth: 0,
        logId: 3,
        sessionId: '#Account_1_4_0',
        transfers: {
          '#Token_1': 1000000000,
        },
        txHash: '#Hash_5',
      },
    ],
    blockNum: 4,
    cmStats: [
      {
        Address: '#CreditManager_1',
        AvailableLiquidity: 5000,
        AvailableLiquidityBI: '5000000000',
        BlockNum: 4,
        BorrowRate: 0,
        BorrowRateBI: '0',
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
            linked: true,
          },
        },
        blockNum: 4,
        borrowedAmount: 5000,
        borrowedAmountBI: '5000000000',
        borrower: '#User_1',
        collateralInUSD: 1000,
        collateralInUnderlying: 1000,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '10800',
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
        txHash: '#Hash_4',
        user: '#User_1',
      },
      {
        amount: 1000,
        blockNum: 4,
        event: 'Borrow',
        logId: 3,
        pool: '#Pool_1',
        sessionId: '#Account_1_4_0',
        txHash: '#Hash_5',
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
        expectedLiquidity: 10000,
        expectedLiquidityBI: '10000000000',
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
    accountOperations: [
      {
        action: 'MultiCallStarted(address)',
        adapterCall: false,
        blockNum: 5,
        borrower: '#User_1',
        dapp: '#CreditFacade_1',
        depth: 0,
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
              token: '#Token_3',
              value: 1000000000000000000,
            },
            blockNum: 5,
            borrower: '#User_1',
            dapp: '#CreditFacade_1',
            depth: 0,
            logId: 4,
            sessionId: '#Account_1_4_0',
            transfers: {
              '#Token_3': 1000000000000000000,
            },
            txHash: '#Hash_7',
          },
          {
            action: 'DecreaseBorrowedAmount(address,uint256)',
            adapterCall: false,
            args: {
              _order: [
                'borrower',
                'amount',
              ],
              amount: 500000000,
              borrower: '#User_1',
            },
            blockNum: 5,
            borrower: '#User_1',
            dapp: '#CreditFacade_1',
            depth: 0,
            logId: 5,
            sessionId: '#Account_1_4_0',
            transfers: {
              '#Token_1': -500000000,
            },
            txHash: '#Hash_7',
          },
        ],
        sessionId: '#Account_1_4_0',
        txHash: '#Hash_7',
      },
    ],
    allowedTokens: [
      {
        BlockNumber: 5,
        CreditManager: '#CreditManager_1',
        DisableBlock: 0,
        LiquidityThreshold: '8000',
        Token: '#Token_3',
      },
    ],
    blockNum: 5,
    cmStats: [
      {
        Address: '#CreditManager_1',
        AvailableLiquidity: 5500,
        AvailableLiquidityBI: '5500000000',
        BlockNum: 5,
        BorrowRate: 0,
        BorrowRateBI: '0',
        CumulativeBorrowed: 0,
        ID: 0,
        OpenedAccountsCount: 1,
        TotalBorrowed: 4500,
        TotalBorrowedBI: '4500000000',
        TotalClosedAccounts: 0,
        TotalLiquidatedAccounts: 0,
        TotalLosses: 0,
        TotalLossesBI: '0',
        TotalOpenedAccounts: 1,
        TotalProfit: 0,
        TotalProfitBI: '0',
        TotalRepaid: 500,
        TotalRepaidAccounts: 0,
        TotalRepaidBI: '500000000',
        UniqueUsers: 0,
      },
    ],
    css: [
      {
        balance: {
          '#Token_1': {
            BI: '5500000000',
            F: 5500,
            linked: true,
          },
          '#Token_3': {
            BI: '1000000000000000000',
            F: 1,
            linked: true,
          },
        },
        blockNum: 5,
        borrowedAmount: 4500,
        borrowedAmountBI: '4500000000',
        borrower: '#User_1',
        collateralInUSD: 3500,
        collateralInUnderlying: 3500,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '15444',
        sessionId: '#Account_1_4_0',
        totalValue: 8000,
        totalValueBI: '8000000000',
      },
    ],
    daoOperations: [
      {
        Args: {
          creditManager: '#CreditManager_1',
          token: '#Token_3',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 1,
        TxHash: '#Hash_6',
        Type: 22,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          liquidityThreshold: '8000',
          prevLiquidationThreshold: '0',
          token: '#Token_3',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 2,
        TxHash: '#Hash_6',
        Type: 27,
      },
    ],
    poolLedgers: [
      {
        amount: 500,
        blockNum: 5,
        event: 'Repay',
        logId: 5,
        pool: '#Pool_1',
        sessionId: '#Account_1_4_0',
        txHash: '#Hash_7',
        user: '#User_1',
      },
    ],
    priceFeeds: [
      {
        blockNum: 5,
        feed: '#ChainlinkPriceFeed_3',
        isPriceInUSD: true,
        price: 2500,
        priceBI: '250000000000',
        roundId: 1,
        token: '#Token_3',
        uniPriceFetchBlock: 0,
        uniswapv2Price: 0,
        uniswapv3Price: 0,
        uniswapv3Twap: 0,
      },
    ],
    timestamp: 432000,
  },
  '6': {
    accountOperations: [
      {
        action: 'MultiCallStarted(address)',
        adapterCall: false,
        blockNum: 6,
        borrower: '#User_1',
        dapp: '#CreditFacade_1',
        depth: 0,
        logId: 0,
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
              amountIn: 'bigint:2500000000',
              amountOutMin: 'bigint:1000000000000000000',
              deadline: 0,
              path: [
                '#Token_1',
                '#Token_3',
              ],
            },
            blockNum: 6,
            borrower: '#User_1',
            dapp: '#Uniswapv2_1',
            depth: 0,
            logId: 1,
            sessionId: '#Account_1_4_0',
            transfers: {
              '#Token_1': -2500000000,
              '#Token_3': 1000000000000000000,
            },
            txHash: '#Hash_8',
          },
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
              amountIn: 'bigint:2500000000',
              amountOutMin: 'bigint:1000000000000000000',
              deadline: 0,
              path: [
                '#Token_1',
                '#Token_3',
              ],
            },
            blockNum: 6,
            borrower: '#User_1',
            dapp: '#Uniswapv2_1',
            depth: 0,
            logId: 2,
            sessionId: '#Account_1_4_0',
            transfers: {
              '#Token_1': -2500000000,
              '#Token_3': 1000000000000000000,
            },
            txHash: '#Hash_8',
          },
        ],
        sessionId: '#Account_1_4_0',
        txHash: '#Hash_8',
      },
    ],
    blockNum: 6,
    cmStats: [
      {
        Address: '#CreditManager_1',
        AvailableLiquidity: 1000,
        AvailableLiquidityBI: '1000000000',
        BlockNum: 6,
        BorrowRate: 0,
        BorrowRateBI: '0',
        CumulativeBorrowed: 0,
        ID: 0,
        OpenedAccountsCount: 1,
        TotalBorrowed: 4500,
        TotalBorrowedBI: '4500000000',
        TotalClosedAccounts: 0,
        TotalLiquidatedAccounts: 0,
        TotalLosses: 0,
        TotalLossesBI: '0',
        TotalOpenedAccounts: 1,
        TotalProfit: 0,
        TotalProfitBI: '0',
        TotalRepaid: 500,
        TotalRepaidAccounts: 0,
        TotalRepaidBI: '500000000',
        UniqueUsers: 0,
      },
    ],
    css: [
      {
        balance: {
          '#Token_1': {
            BI: '500000000',
            F: 500,
            linked: true,
          },
          '#Token_3': {
            BI: '3000000000000000000',
            F: 3,
            linked: true,
          },
        },
        blockNum: 6,
        borrowedAmount: 4500,
        borrowedAmountBI: '4500000000',
        borrower: '#User_1',
        collateralInUSD: 3500,
        collateralInUnderlying: 3500,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '14333',
        sessionId: '#Account_1_4_0',
        totalValue: 8000,
        totalValueBI: '8000000000',
      },
    ],
    timestamp: 518400,
  },
  '7': {
    accountOperations: [
      {
        action: 'DirectTokenTransfer',
        adapterCall: false,
        args: {
          amount: 1000000000,
          from: '#User_3',
          to: '#Account_1',
        },
        blockNum: 7,
        borrower: '#User_1',
        dapp: '#Token_1',
        depth: 0,
        logId: 3,
        sessionId: '#Account_1_4_0',
        transfers: {
          '#Token_1': 1000000000,
        },
        txHash: '#Hash_10',
      },
      {
        action: 'DirectTokenTransfer',
        adapterCall: false,
        args: {
          amount: 100000000000000000,
          from: '#User_3',
          to: '#Account_1',
        },
        blockNum: 7,
        borrower: '#User_1',
        dapp: '#Token_2',
        depth: 0,
        logId: 4,
        sessionId: '#Account_1_4_0',
        transfers: {
          '#Token_2': 100000000000000000,
        },
        txHash: '#Hash_11',
      },
    ],
    allowedTokens: [
      {
        BlockNumber: 7,
        CreditManager: '#CreditManager_1',
        DisableBlock: 0,
        LiquidityThreshold: '9000',
        Token: '#Token_2',
      },
    ],
    blockNum: 7,
    css: [
      {
        balance: {
          '#Token_1': {
            BI: '1500000000',
            F: 1500,
            linked: true,
          },
          '#Token_2': {
            BI: '100000000000000000',
            F: 0.1,
            linked: false,
          },
          '#Token_3': {
            BI: '3000000000000000000',
            F: 3,
            linked: true,
          },
        },
        blockNum: 7,
        borrowedAmount: 4500,
        borrowedAmountBI: '4500000000',
        borrower: '#User_1',
        collateralInUSD: 6500,
        collateralInUnderlying: 6500,
        cumulativeIndexAtOpen: '1000000000000000000000000000',
        healthFactor: '16333',
        sessionId: '#Account_1_4_0',
        totalValue: 9000,
        totalValueBI: '9000000000',
      },
    ],
    daoOperations: [
      {
        Args: {
          creditManager: '#CreditManager_1',
          token: '#Token_2',
        },
        BlockNumber: 7,
        Contract: '#CreditConfigurator_1',
        LogID: 0,
        TxHash: '#Hash_9',
        Type: 22,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          liquidityThreshold: '9000',
          prevLiquidationThreshold: '0',
          token: '#Token_2',
        },
        BlockNumber: 7,
        Contract: '#CreditConfigurator_1',
        LogID: 1,
        TxHash: '#Hash_9',
        Type: 27,
      },
    ],
    priceFeeds: [
      {
        blockNum: 7,
        feed: '#ChainlinkPriceFeed_2',
        isPriceInUSD: true,
        price: 20000,
        priceBI: '2000000000000',
        roundId: 1,
        token: '#Token_2',
        uniPriceFetchBlock: 0,
        uniswapv2Price: 0,
        uniswapv3Price: 0,
        uniswapv3Twap: 0,
      },
    ],
    timestamp: 604800,
  },
}
