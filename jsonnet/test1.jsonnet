local bigInt(x, decimals=0) = std.toString(x * std.pow(10, decimals));
local bigIntTopic(x, decimals) = 'bigint:' + bigInt(x, decimals);
{
  _order:["addresses", "blocks"],
  addresses:[
    "#User_1",
    "#Account_1",
  ],
  mocks: {
    syncAdapters: 'mocks/syncAdapter1.json',
  },
  blocks: {
    '3': {
      events: [
        {
          // credit manager on usdc
          address: '@CreditManager_1',
          topics: [
            'OpenCreditAccount(address,address,address,uint256,uint256,uint256)',
            '@User_1',
            '@User_1',
            '@Account_1',
          ],
          data: [
            bigIntTopic(1000, 6),
            bigIntTopic(400, 6),
            'bigint:0',
          ],
          txHash: '!#Hash_1',
        },
        {
          // price chainlink on usdc
          address: '@ChainlinkPriceFeed_1',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // roundid
            bigIntTopic(1, 0),
            // 0.0003
            bigIntTopic(0.0003, 18),
          ],
          data: [],
        },
      ],
      calls:
        {
          masks: [{
            account: "@Account_1",
            mask: "1",
          }],
          pools: [{
            address: "@Pool_1",
            totalBorrowed: bigInt(4000, 6),
            expectedLiquidity: bigInt(4000, 6),
            availableLiquidity: bigInt(1000, 6),
            depositAPY: bigInt(0),
            borrowAPY: bigInt(0),
            dieselRate: bigInt(0),
            withdrawFee: '0',
            linearCumulativeIndex: bigInt(1, 27),
          }],
          accounts: [{
            address: "@Account_1",
            creditManager: "@CreditManager_1",
            borrower: "@User_1",
            healthFactor: '12500',
            totalValue: bigInt(5000, 6),
            repayAmount: bigInt(4000, 6),
            cumulativeIndexAtOpen: bigInt(1, 27),
            borrowedAmount: bigInt(4000,6),
            balances: [{
              token: '@Token_1',
              balance: bigInt(5000, 6),
              isAllowed: true,
            }],
          }],
          cms: [{
            address: "@CreditManager_1",
            isWETH: false,
            minAmount: bigInt(1000, 6),
            maxAmount: bigInt(5000, 6),
            availableLiquidity: bigInt(1000, 6),
            borrowRate: '0',
          }],
        },
    },
  },
}
