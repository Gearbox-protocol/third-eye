local utils = import '../utils.libsonnet';
local borrowedAmount = 4000;
local extraBorrowedAmount = 1000;
{
  blocks: {
    '7': {
      events: [
        {
          // price chainlink on usdc
          address: '#ChainlinkPriceFeed_1',
          txHash: '!#Hash_12',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // 0.0008
            utils.bigIntTopic(0.0008, 18),
            // roundid
            utils.bigIntTopic(2, 0),
          ],
          data: [],
        },
      ],
      calls:
        {
          masks: [{
            account: '#Account_1',
            mask: '3',
          }],
          accounts: [{
            address: '#Account_1',
            creditManager: '#CreditManager_1',
            borrower: '#User_1',
            healthFactor: '9600',
            // 1000 is for direct token transfer of token 1 usdc
            // .1 YFI = 8*ETH *.1= 2000 USDC  // yfi is not linked so it not included in yfi
            totalValue: utils.bigInt(3 * 1250 + 2000, 6),
            repayAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
            liquidationAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
            cumulativeIndexAtOpen: utils.bigInt(1, 27),
            borrowedAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
            borrowedAmountPlusInterest: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
            balances: [{
              token: '#Token_1',
              balance: utils.bigInt(2000, 6),
              isAllowed: true,
            }, {
              token: '#Token_3',
              balance: utils.bigInt(3, 18),
              isAllowed: true,
            }, {
              // token 2 yfi is allowed.but  its not linked to account
              token: '#Token_2',
              balance: utils.bigInt(0.1, 18),
              isAllowed: true,
            }],
          }],
        },
    },
    '8': {
      events: [
        {
          // close credit account
          address: '#CreditManager_1',
          topics: [
            'LiquidateCreditAccount(address,address,uint256)',
            '#User_1',
            '#User_2',
          ],
          data: [
            utils.bigIntTopic(5750 - 5109.25, 6),
          ],
          txHash: '!#Hash_13',
        },
        {
          // credit filter on usdc
          address: '#Pool_1',
          topics: [
            'Repay(address,uint256,uint256,uint256)',
            '#CreditManager_1',
          ],
          data: [
            // borrowedamount
            utils.bigIntTopic(5000, 6),
            // profit
            utils.bigIntTopic(0, 0),
            // loss
            utils.bigIntTopic(0, 0),
          ],
          txHash: '!#Hash_13',
        },
      ],
      calls: {
        pools: [{
          address: '#Pool_1',
          totalBorrowed: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
          expectedLiquidity: utils.bigInt(borrowedAmount + extraBorrowedAmount + 1000, 6),
          availableLiquidity: utils.bigInt(6000, 6),
          depositAPY: utils.bigInt(0),
          borrowAPY: utils.bigInt(0),
          dieselRate: utils.bigInt(0),
          withdrawFee: '0',
          linearCumulativeIndex: utils.bigInt(1, 27),
        }],
        cms: [{
          address: '#CreditManager_1',
          isWETH: false,
          minAmount: utils.bigInt(1000, 6),
          maxAmount: utils.bigInt(5000, 6),
          availableLiquidity: utils.bigInt(6000, 6),
          borrowRate: '0',
        }],
      },
    },
  },
}
