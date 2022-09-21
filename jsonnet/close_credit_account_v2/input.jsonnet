local utils = import '../utils.libsonnet';
local borrowedAmount = 4000;
local extraBorrowedAmount = 1000;
{
  executeParser: {
    '9': {
      executeOnCM: {
        '!#Hash_12': [{
          name: 'swapExactTokensForTokens(uint256,uint256,address[],address,uint256)',
          args: {
            _order: ['amountIn', 'amountOutMin', 'path', '', 'deadline'],
            amountIn: utils.bigIntTopic(2, 18),
            amountOutMin: utils.bigIntTopic(1, 18),
            path: ['#Token_3', '#Token_1'],
            '': '#Account_1',
            deadline: 0,
          },
          depth: 0,
          transfers: {
            '#Token_3': utils.bigInt(-2, 18),
            '#Token_1': utils.bigInt(4000, 6),
          },
        }],
      },
      executeTransfers: {
        '!#Hash_12': {
          '#Token_1': utils.bigInt(1000, 6),
          '#Token_2': utils.bigInt(0.1, 18),
          '#Token_3': utils.bigInt(1, 18),
        },
      },
      mainEventLogs: {
        '!#Hash_12': [{
          name: 'closeCreditAccount',
          len: 1,
        }],
      },
    },
  },
  blocks: {
    '8': {
      calls: {
        masks: [{
          account: '#Account_1',
          mask: '3',
        }],
        accounts: [{
          address: '#Account_1',
          creditManager: '#CreditManager_1',
          borrower: '#User_1',
          healthFactor: '16333',  // 7350/4500 = (.9*1500+ 3*2500*.8)/(4000+1000-500)
          // 1000 is for direct token transfer of token 1 usdc
          // .1 YFI = 8*ETH *.1= 2000 USDC  // yfi is not linked so it not included in token value
          totalValue: utils.bigInt(500 + 3 * 2500 + 1000, 6),
          repayAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount / 2, 6),
          cumulativeIndexAtOpen: utils.bigInt(1, 27),
          borrowedAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount / 2, 6),
          borrowedAmountPlusInterest: utils.bigInt(borrowedAmount + extraBorrowedAmount / 2, 6),
          balances: [{
            token: '#Token_1',
            balance: utils.bigInt(1500, 6),
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
    '9': {
      events: [
        {
          // multicall start
          address: '#CreditFacade_1',
          topics: [
            'MultiCallStarted(address)',
            '#User_1',
          ],
          txHash: '!#Hash_12',
        },
        {
          // credit filter on usdc
          address: '#CreditManager_1',
          topics: [
            'ExecuteOrder(address,address)',
            '#User_1',
            '#Uniswapv2_1',
          ],
          txHash: '!#Hash_12',
        },
        {
          // multicall end
          address: '#CreditFacade_1',
          topics: [
            'MultiCallFinished()',
          ],
          txHash: '!#Hash_12',
        },
        {
          // close credit account
          address: '#CreditManager_1',
          topics: [
            'CloseCreditAccount(address,address)',
            '#User_1',  // borrower
            // to
            '#User_2',
          ],
          txHash: '!#Hash_12',
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
            utils.bigIntTopic(4500, 6),
            // profit
            utils.bigIntTopic(0, 0),
            // loss
            utils.bigIntTopic(0, 0),
          ],
          txHash: '!#Hash_12',
        },
      ],
      calls:
        {
          masks: [{
            account: '#Account_1',
            mask: '3',
          }],
          pools: [{
            address: '#Pool_1',
            totalBorrowed: utils.bigInt(0, 6),
            expectedLiquidity: utils.bigInt(10000, 6),
            availableLiquidity: utils.bigInt(10000, 6),
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
            maxAmount: utils.bigInt(6000, 6),
            availableLiquidity: utils.bigInt(10000, 6),
            borrowRate: '0',
          }],
        },
    },
  },
}
