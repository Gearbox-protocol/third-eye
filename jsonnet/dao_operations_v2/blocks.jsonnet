{
  '3': {
    allowedTokens: [
      {
        BlockNumber: 3,
        Configurator: '#CreditConfigurator_1',
        CreditManager: '#CreditManager_1',
        DisableBlock: 4,
        LiquidityThreshold: '7500',
        Token: '#Token_1',
      },
    ],
    daoOperations: [
      {
        Args: {
          creditManager: '#CreditManager_1',
          prevLiquidationThreshold: '0',
          token: '#Token_1',
        },
        BlockNumber: 3,
        Contract: '#CreditConfigurator_1',
        LogID: 0,
        TxHash: '#Hash_1',
        Type: 22,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          liquidityThreshold: '7500',
          prevLiquidationThreshold: '0',
          token: '#Token_1',
        },
        BlockNumber: 3,
        Contract: '#CreditConfigurator_1',
        LogID: 1,
        TxHash: '#Hash_2',
        Type: 27,
      },
    ],
  },
  '4': {
    allowedTokens: [
      {
        BlockNumber: 4,
        Configurator: '#CreditConfigurator_1',
        CreditManager: '#CreditManager_1',
        DisableBlock: 5,
        LiquidityThreshold: '8000',
        Token: '#Token_1',
      },
    ],
    daoOperations: [
      {
        Args: {
          creditManager: '#CreditManager_1',
          liquidityThreshold: '8000',
          prevLiquidationThreshold: '7500',
          token: '#Token_1',
        },
        BlockNumber: 4,
        Contract: '#CreditConfigurator_1',
        LogID: 0,
        TxHash: '#Hash_3',
        Type: 27,
      },
    ],
  },
  '5': {
    allowedTokens: [
      {
        BlockNumber: 5,
        Configurator: '#CreditConfigurator_1',
        CreditManager: '#CreditManager_1',
        DisableBlock: 0,
        LiquidityThreshold: '8000',
        Token: '#Token_1',
      },
    ],
    daoOperations: [
      {
        Args: {
          configurator: '#CreditConfigurator_2',
          oldConfigurator: '#CreditConfigurator_1',
        },
        BlockNumber: 5,
        Contract: '#CreditManager_1',
        LogID: 7,
        TxHash: '#Hash_13',
        Type: 26,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          token: '#Token_1',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 0,
        TxHash: '#Hash_4',
        Type: 1,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          liquidityThreshold: '8000',
          prevLiquidationThreshold: '8000',
          token: '#Token_1',
          type: 'reEnabled',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 1,
        TxHash: '#Hash_5',
        Type: 22,
      },
      {
        Args: {
          adapter: '#Adapter_1',
          creditManager: '#CreditManager_1',
          protocol: '#Protocol_1',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 2,
        TxHash: '#Hash_6',
        Type: 2,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          protocol: '#Protocol_1',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 3,
        TxHash: '#Hash_7',
        Type: 3,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          newPriceOracle: '#PriceOracle_2',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 4,
        TxHash: '#Hash_9',
        Type: 6,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          maxAmount: [
            '0',
            '5000000000',
          ],
          minAmount: [
            '0',
            '1000000000',
          ],
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 5,
        TxHash: '#Hash_11',
        Type: 23,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          feeInterest: [
            0,
            1000,
          ],
          feeLiquidation: [
            0,
            200,
          ],
          feeLiquidationExpired: [
            0,
            300,
          ],
          liquidationDiscount: [
            0,
            9500,
          ],
          liquidationDiscountExpired: [
            0,
            9400,
          ],
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 6,
        TxHash: '#Hash_12',
        Type: 24,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          facade: '#CreditFacade_1',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 8,
        TxHash: '#Hash_14',
        Type: 25,
      },
    ],
  },
  '6': {
    daoOperations: [
      {
        Args: {
          creditManager: '#CreditManager_1',
          newPriceOracle: '#PriceOracle_1',
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 0,
        TxHash: '#Hash_15',
        Type: 6,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          increaseDebtForbiddenMode: 1,
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 1,
        TxHash: '#Hash_16',
        Type: 28,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          date: 123456789,
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 2,
        TxHash: '#Hash_16',
        Type: 29,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          maxEnabledTokens: 10,
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 3,
        TxHash: '#Hash_16',
        Type: 30,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          limitPerBlock: 10,
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 4,
        TxHash: '#Hash_16',
        Type: 31,
      },
      {
        Args: {
          contract: '#Admin_1',
          creditManager: '#CreditManager_1',
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 5,
        TxHash: '#Hash_17',
        Type: 32,
      },
      {
        Args: {
          contract: '#Admin_1',
          creditManager: '#CreditManager_1',
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 6,
        TxHash: '#Hash_17',
        Type: 33,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          emergencyLiquidator: '#Emergencyiquidator_1',
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 7,
        TxHash: '#Hash_18',
        Type: 34,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          emergencyLiquidator: '#Emergencyiquidator_1',
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 8,
        TxHash: '#Hash_18',
        Type: 35,
      },
      {
        Args: {
          adapter: '#Adapter_1',
          creditManager: '#CreditManager_1',
        },
        BlockNumber: 6,
        Contract: '#CreditConfigurator_2',
        LogID: 9,
        TxHash: '#Hash_18',
        Type: 36,
      },
    ],
  },
}
