{
  '3': {
    allowedTokens: [
      {
        BlockNumber: 3,
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
        CreditManager: '#CreditManager_1',
        DisableBlock: 0,
        LiquidityThreshold: '8000',
        Token: '#Token_1',
      },
    ],
    daoOperations: [
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
        Type: 27,
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
          chiThreshold: [
            '0',
            '7500',
          ],
          creditManager: '#CreditManager_1',
          fastDelay: [
            '0',
            '7500',
          ],
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 4,
        TxHash: '#Hash_8',
        Type: 4,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          newPriceOracle: '#PriceOracle_2',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 5,
        TxHash: '#Hash_9',
        Type: 6,
      },
      {
        Args: {
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
        LogID: 6,
        TxHash: '#Hash_11',
        Type: 23,
      },
      {
        Args: {
          LiquidationDiscount: [
            '0',
            '9500',
          ],
          feeInterest: [
            '0',
            '1000',
          ],
          feeLiquidation: [
            '0',
            '200',
          ],
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 7,
        TxHash: '#Hash_12',
        Type: 24,
      },
      {
        Args: {
          configurator: '#CreditConfigurator_1',
          oldConfigurator: '#CreditConfigurator_1',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 8,
        TxHash: '#Hash_13',
        Type: 4,
      },
      {
        Args: {
          facade: '#CreditFacade_1',
        },
        BlockNumber: 5,
        Contract: '#CreditConfigurator_1',
        LogID: 9,
        TxHash: '#Hash_14',
        Type: 4,
      },
    ],
  },
}
