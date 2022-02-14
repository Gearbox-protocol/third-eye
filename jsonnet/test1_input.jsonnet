local bigInt(x, decimals=0) = std.toString(x * std.pow(10, decimals));
local bigIntTopic(x, decimals) = 'bigint:' + bigInt(x, decimals);
local initialAmount = 1000;
local borrowedAmount = 4000;
{
  mocks: {
    syncAdapters: 'mocks/syncAdapter1.json',
  },
  states: {
    oracles: [{
      oracle: '#Oracle_1',
      block: 1,
      feed: '#ChainlinkPriceFeed_1',
    }, {
      oracle: '#Oracle_2',
      block: 1,
      feed: '#ChainlinkPriceFeed_2',
    }],
  },
  blocks: {
    '3': {
      events: [
        {
          // credit manager on usdc
          address: '#CreditManager_1',
          topics: [
            'OpenCreditAccount(address,address,address,uint256,uint256,uint256)',
            '#User_1',
            '#User_1',
            '#Account_1',
          ],
          data: [
            bigIntTopic(initialAmount, 6),
            bigIntTopic(borrowedAmount, 6),
            'bigint:0',
          ],
          txHash: '!#Hash_1',
        },
        {
          // pool on usdc
          address: '#CreditFilter_1',
          topics: [
            'TokenAllowed(address,uint256)',
            '#Token_1',
          ],
          data: [
            bigIntTopic(9000, 0),
          ],
          txHash: '!#Hash_3',
        },
        {
          // pool on usdc
          address: '#Pool_1',
          topics: [
            'Borrow(address,address,uint256)',
            '#CreditManager_1',
            '#Account_1',
          ],
          data: [
            bigIntTopic(borrowedAmount, 6),
          ],
          txHash: '!#Hash_2',
        },
        {
          // price chainlink on usdc
          address: '#ChainlinkPriceFeed_1',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // roundid
            bigIntTopic(1, 0),
            // 0.0003
            bigIntTopic(0.0003, 18),
          ],
          data: [],
        },
        {
          // price chainlink on usdc
          address: '#CreditManager_1',
          topics: [
            'NewParameters(uint256,uint256,uint256,uint256,uint256,uint256)',
          ],
          data: [
            // minAnount
            bigIntTopic(1000, 6),
            // maxAmount
            bigIntTopic(5000, 6),
            // maxLeverage
            bigIntTopic(400, 6),
            // feeInterest
            bigIntTopic(1000, 0),
            // feeLiquidation
            bigIntTopic(200, 0),
            // liquidationDiscount
            bigIntTopic(9500, 0),
          ],
        },
      ],
      calls:
        {
          masks: [{
            account: '#Account_1',
            mask: '1',
          }],
          pools: [{
            address: '#Pool_1',
            totalBorrowed: bigInt(borrowedAmount, 6),
            expectedLiquidity: bigInt(borrowedAmount + 1000, 6),
            availableLiquidity: bigInt(1000, 6),
            depositAPY: bigInt(0),
            borrowAPY: bigInt(0),
            dieselRate: bigInt(0),
            withdrawFee: '0',
            linearCumulativeIndex: bigInt(1, 27),
          }],
          accounts: [{
            address: '#Account_1',
            creditManager: '#CreditManager_1',
            borrower: '#User_1',
            healthFactor: '11250',
            totalValue: bigInt(borrowedAmount + initialAmount, 6),
            repayAmount: bigInt(borrowedAmount, 6),
            cumulativeIndexAtOpen: bigInt(1, 27),
            borrowedAmount: bigInt(borrowedAmount, 6),
            balances: [{
              token: '#Token_1',
              balance: bigInt(5000, 6),
              isAllowed: true,
            }],
          }],
          cms: [{
            address: '#CreditManager_1',
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
