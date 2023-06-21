{
  '3': {
    daoOperations: [
      {
        Args: {},
        BlockNumber: 3,
        Contract: '#CreditConfigurator_1',
        LogID: 1,
        TxHash: '#Hash_4',
        Type: 38,
      },
      {
        Args: {
          creditManager: '#CreditManager_1',
          emergencyLiqDiscount: [
            0,
            9600,
          ],
        },
        BlockNumber: 3,
        Contract: '#CreditConfigurator_1',
        LogID: 2,
        TxHash: '#Hash_1',
        Type: 39,
      },
    ],
  },
  '4': {
    daoOperations: [
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
        BlockNumber: 4,
        Contract: '#CreditConfigurator_1',
        LogID: 0,
        TxHash: '#Hash_2',
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
        BlockNumber: 4,
        Contract: '#CreditConfigurator_1',
        LogID: 1,
        TxHash: '#Hash_3',
        Type: 24,
      },
    ],
  },
}
