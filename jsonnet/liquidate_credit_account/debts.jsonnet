{
  currentDebts: [
    {
      amountToPool: 5000 + 5750 * 0.95 * 0.02,  // borrowedamount+ totalvalue*feeLiquidation*liquidationCount, 5109.25
      blockNum: 7,
      calBorrowedAmountPlusInterest: 5000,
      calHealthFactor: '9600',
      calThresholdValue: 4800,  //1250*.8*3+2000*.9
      calTotalValue: 5750,  // totalvalue
      collateralUSD: 6500,  // 2000 yfi + 2000 usdc + 2500 worth eth
      collateralUnderlying: 6500,
      profitUSD: 5750 - (5109.25 + 6500),  // (total value  amountToPool , - collateral)
      profitUnderlying: 5750 - (5109.25 + 6500),
      sessionId: '#Account_1_3_3',
      repayAmount: 5750,
    },
  ],
  debts: [{
    blockNum: 7,
    calBorrowedAmountPlusInterest: '5000000000',
    calHealthFactor: '9600',
    calThresholdValue: '4800000000',
    calTotalValue: '5750000000',
    collateralUSD: 6500,
    collateralUnderlying: 6500,
    profitUSD: 5750 - (5109.25 + 6500),
    profitUnderlying: 5750 - (5109.25 + 6500),
    sessionId: '#Account_1_3_3',
    totalValueInUSD: 5750,
  }],
}
