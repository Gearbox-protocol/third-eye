{ data: [
  {
    Args: {
      priceFeed: '#Oracle_3',
      token: '#Token_4',
    },
    BlockNumber: 3,
    Contract: '#PriceOracle_1',
    LogID: 12,
    TxHash: '#Hash_13',
    Type: 12,  //  NewPriceFeed
  },
  {
    Args: {
      admin: '#Admin_1',
    },
    BlockNumber: 3,
    Contract: '#ACL_1',
    LogID: 14,
    TxHash: '#Hash_15',
    Type: 14,  // PausableAdminAdded
  },
  {
    Args: {
      admin: '#Admin_1',
    },
    BlockNumber: 3,
    Contract: '#ACL_1',
    LogID: 15,
    TxHash: '#Hash_16',
    Type: 15,  // PausableAdminRemoved
  },
  {
    Args: {
      admin: '#Admin_1',
    },
    BlockNumber: 3,
    Contract: '#ACL_1',
    LogID: 16,
    TxHash: '#Hash_17',
    Type: 16,  // UnpausableAdminAdded
  },
  {
    Args: {
      admin: '#Admin_1',
    },
    BlockNumber: 3,
    Contract: '#ACL_1',
    LogID: 17,
    TxHash: '#Hash_18',
    Type: 17,  // UnpausableAdminRemoved
  },
  {
    Args: {
      newOwner: '#Admin_2',
      oldOwner: '#Owner_1',
    },
    BlockNumber: 3,
    Contract: '#ACL_1',
    LogID: 18,
    TxHash: '#Hash_19',
    Type: 18,  // OwnershipTransferred
  },
  {
    BlockNumber: 3,
    Contract: '#ACL_1',
    LogID: 19,
    TxHash: '#Hash_20',
    Type: 19,  // Paused
  },
  {
    BlockNumber: 3,
    Contract: '#ACL_1',
    LogID: 20,
    TxHash: '#Hash_21',
    Type: 20,  // Unpaused
  },
  {
    Args: {
      creditAccount: '#Account_10',
      to: '#To_1',
    },
    BlockNumber: 3,
    Contract: '#AccountFactory_1',
    LogID: 13,
    TxHash: '#Hash_14',
    Type: 13,  // TakeForever
  },
  {
    Args: {
      newInterestRateModel: '#IntereestRateModel_1',
    },
    BlockNumber: 3,
    Contract: '#Pool_1',
    LogID: 7,
    TxHash: '#Hash_8',
    Type: 7,  // NewInterestRateModel
  },
  {
    Args: {
      creditManager: '#CreditManager_1',
    },
    BlockNumber: 3,
    Contract: '#Pool_1',
    LogID: 8,
    TxHash: '#Hash_9',
    Type: 8,  // NewCreditManagerConnected
  },
  {
    Args: {
      newLimit: '10000000000',
      token: '#Token_1',
    },
    BlockNumber: 3,
    Contract: '#Pool_1',
    LogID: 9,
    TxHash: '#Hash_10',
    Type: 9,  // NewExpectedLiquidityLimit
  },
  {
    Args: {
      creditManager: '#CreditManager_2',
    },
    BlockNumber: 3,
    Contract: '#Pool_1',
    LogID: 10,
    TxHash: '#Hash_11',
    Type: 10,  // BorrowForbidden
  },
  {
    Args: {
      newFee: '100',
      oldFee: '0',
      token: '#Token_1',
    },
    BlockNumber: 3,
    Contract: '#Pool_1',
    LogID: 11,
    TxHash: '#Hash_12',
    Type: 11,  // NewWithdrawFee
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
    LogID: 21,
    TxHash: '#Hash_22',
    Type: 21,
  },
  {
    Args: {
      creditManager: '#CreditManager_1',
      liquidityThreshold: '7500',
      prevLiquidationThreshold: '0',
      token: '#Token_1',
    },
    BlockNumber: 3,
    Contract: '#CreditFilter_1',
    LogID: 0,
    TxHash: '#Hash_1',
    Type: 0,  // TokenAllowed
  },
  {
    Args: {
      creditManager: '#CreditManager_1',
      token: '#Token_1',
    },
    BlockNumber: 3,
    Contract: '#CreditFilter_1',
    LogID: 1,
    TxHash: '#Hash_2',
    Type: 1,  // TokenForbidden
  },
  {
    Args: {
      adapter: '#Adapter_1',
      creditManager: '#CreditManager_1',
      protocol: '#Protocol_1',
    },
    BlockNumber: 3,
    Contract: '#CreditFilter_1',
    LogID: 2,
    TxHash: '#Hash_3',
    Type: 2,  // ContractAllowed
  },
  {
    Args: {
      creditManager: '#CreditManager_1',
      protocol: '#Protocol_1',
    },
    BlockNumber: 3,
    Contract: '#CreditFilter_1',
    LogID: 3,
    TxHash: '#Hash_4',
    Type: 3,  // ContractForbidden
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
    BlockNumber: 3,
    Contract: '#CreditFilter_1',
    LogID: 4,
    TxHash: '#Hash_5',
    Type: 4,  // NewFastCheckParameters
  },
  {
    Args: {
      creditManager: '#CreditManager_1',
      plugin: '#Plugin_1',
      state: true,
    },
    BlockNumber: 3,
    Contract: '#CreditFilter_1',
    LogID: 5,
    TxHash: '#Hash_6',
    Type: 5,  // TransferPluginAllowed
  },
  {
    Args: {
      creditManager: '#CreditManager_1',
      newPriceOracle: '#PriceOracle_2',
    },
    BlockNumber: 3,
    Contract: '#CreditFilter_1',
    LogID: 6,
    TxHash: '#Hash_7',
    Type: 6,  // PriceOracleUpdated
  },
] }
