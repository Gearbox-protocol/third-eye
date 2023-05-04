local utils = import '../utils.libsonnet';
local initialAmount = 1000;
local borrowedAmount = 4000;
{
  mocks: {
    syncAdapters: 'mocks/syncAdapter1.json',
    tokens: '../inputs/mocks/tokens.json',
  },
  states: {
    oracles: {
      '#Oracle_1': [{
        oracle: '#Oracle_1',
        block: 1,
        feed: '#ChainlinkPriceFeed_1',
      }],
      '#Oracle_2': [{
        oracle: '#Oracle_2',
        block: 1,
        feed: '#ChainlinkPriceFeed_2',
      }],
    },
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
            utils.bigIntTopic(initialAmount, 6),
            utils.bigIntTopic(borrowedAmount, 6),
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
            utils.bigIntTopic(9000, 0),
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
            utils.bigIntTopic(borrowedAmount, 6),
          ],
          txHash: '!#Hash_2',
        },
        {
          // price chainlink on usdc
          address: '#ChainlinkPriceFeed_1',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // roundid
            utils.bigIntTopic(1, 0),
            // 0.0003
            utils.bigIntTopic(0.0003, 18),
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
            utils.bigIntTopic(1000, 6),
            // maxAmount
            utils.bigIntTopic(5000, 6),
            // maxLeverage
            utils.bigIntTopic(400, 6),
            // feeInterest
            utils.bigIntTopic(1000, 0),
            // feeLiquidation
            utils.bigIntTopic(200, 0),
            // liquidationDiscount
            utils.bigIntTopic(9500, 0),
          ],
          txHash: '!#Hash_3',
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
            totalBorrowed: utils.bigInt(borrowedAmount, 6),
            expectedLiquidity: utils.bigInt(borrowedAmount + 1000, 6),
            availableLiquidity: utils.bigInt(1000, 6),
            depositAPY: utils.bigInt(0),
            borrowAPY: utils.bigInt(0),
            dieselRate: utils.bigInt(0),
            withdrawFee: '0',
            linearCumulativeIndex: utils.bigInt(1, 27),
          }],
          accounts: [{
            address: '#Account_1',
            creditManager: '#CreditManager_1',
            borrower: '#User_1',
            healthFactor: '11250',
            totalValue: utils.bigInt(borrowedAmount + initialAmount, 6),
            repayAmount: utils.bigInt(borrowedAmount, 6),
            cumulativeIndexAtOpen: utils.bigInt(1, 27),
            borrowedAmount: utils.bigInt(borrowedAmount, 6),
            borrowedAmountPlusInterest: utils.bigInt(borrowedAmount, 6),
            balances: [{
              token: '#Token_1',
              balance: utils.bigInt(5000, 6),
              isAllowed: true,
            }],
            version: 1,
          }],
          cms: [{
            address: '#CreditManager_1',
            isWETH: false,
            minAmount: utils.bigInt(1000, 6),
            maxAmount: utils.bigInt(5000, 6),
            availableLiquidity: utils.bigInt(1000, 6),
            borrowRate: '0',
          }],
        },
    },
  },
}
