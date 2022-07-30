///////
// TEST
///////
// open credit account without multicall. Block 4
// multicall for decreaseBorrowedAmount and add collateral. Block 5
// multicall for token swapping. Block 6
// direct token transfers with yif not linked and usdc linked 7

local utils = import '../utils.libsonnet';
local initialAmount = 1000;
local borrowedAmount = 4000;
local extraBorrowedAmount = 1000;
local newCollateral = 1;
// price are in usd
{
  executeParser: {
    '5': {
      mainEventLogs: {
        '!#Hash_7': [{
          name: 'multicall',
          len: 2,
        }],
      },
    },
    '6': {
      executeOnCM: {
        '!#Hash_8': [{
          name: 'swapExactTokensForTokens(uint256,uint256,address[],address,uint256)',
          args: {
            _order: ['amountIn', 'amountOutMin', 'path', '', 'deadline'],
            amountIn: utils.bigIntTopic(2500, 6),
            amountOutMin: utils.bigIntTopic(1, 18),
            path: ['#Token_1', '#Token_3'],
            '': '#Account_1',
            deadline: 0,
          },
          depth: 0,
          transfers: {
            '#Token_1': utils.bigInt(-2500, 6),
            '#Token_3': utils.bigInt(1, 18),
          },
        }, {
          name: 'swapExactTokensForTokens(uint256,uint256,address[],address,uint256)',
          args: {
            _order: ['amountIn', 'amountOutMin', 'path', '', 'deadline'],
            amountIn: utils.bigIntTopic(2500, 6),
            amountOutMin: utils.bigIntTopic(1, 18),
            path: ['#Token_1', '#Token_3'],
            '': '#Account_1',
            deadline: 0,
          },
          depth: 0,
          transfers: {
            '#Token_1': utils.bigInt(-2500, 6),
            '#Token_3': utils.bigInt(1, 18),
          },
        }],
      },
      mainEventLogs: {
        '!#Hash_8': [{
          name: 'multicall',
          len: 2,
        }],
      },
    },
  },
  blocks: {
    // block with open and borrow more underlying asset
    '4': {
      events: [
        {
          // open account without multicall
          address: '#CreditFacade_1',
          topics: [
            'OpenCreditAccount(address,address,uint256,uint16)',
            '#User_1',
            '#Account_1',
          ],
          data: [
            utils.bigIntTopic(borrowedAmount, 6),
            'uint16:0',
          ],
          txHash: '!#Hash_4',
        },
        {
          // add collateral
          address: '#CreditFacade_1',
          topics: [
            'AddCollateral(address,address,uint256)',
            '#User_1',
            '#Token_1',
          ],
          data: [
            utils.bigIntTopic(initialAmount, 6),
          ],
          txHash: '!#Hash_4',
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
          txHash: '!#Hash_4',
        },
        {
          // credit manager on usdc increaseborrow
          address: '#CreditFacade_1',
          topics: [
            'IncreaseBorrowedAmount(address,uint256)',
            // borrower
            '#User_1',
          ],
          data: [
            // amount
            utils.bigIntTopic(extraBorrowedAmount, 6),
          ],
          txHash: '!#Hash_5',
        },
        {
          // pool on usdc
          address: '#Pool_1',
          topics: [
            'Borrow(address,address,uint256)',
            '#CreditFacade_1',
            '#Account_1',
          ],
          data: [
            utils.bigIntTopic(extraBorrowedAmount, 6),
          ],
          txHash: '!#Hash_5',
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
            expectedLiquidity: utils.bigInt(borrowedAmount + extraBorrowedAmount + 5000, 6),
            availableLiquidity: utils.bigInt(5000, 6),
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
            healthFactor: '10800',
            totalValue: utils.bigInt(borrowedAmount + initialAmount + extraBorrowedAmount, 6),
            repayAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
            cumulativeIndexAtOpen: utils.bigInt(1, 27),
            borrowedAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
            balances: [{
              token: '#Token_1',
              balance: utils.bigInt(6000, 6),
              isAllowed: true,
            }],
          }],
          cms: [{
            address: '#CreditManager_1',
            isWETH: false,
            minAmount: utils.bigInt(1000, 6),
            maxAmount: utils.bigInt(6000, 6),
            availableLiquidity: utils.bigInt(5000, 6),
            borrowRate: '0',
          }],
        },
    },
    // new block with multicall add collateral and decreaseBorrowedAmount
    '5': {
      events: [
        {
          // WETH price chainlink
          address: '#ChainlinkPriceFeed_3',
          txHash: '!#Hash_6',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // 8
            utils.bigIntTopic(2500, 8),
            // roundid
            utils.bigIntTopic(1, 0),
          ],
        },
        {
          // credit filter on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'TokenAllowed(address)',
            '#Token_3',
          ],
          txHash: '!#Hash_6',
        },
        {
          // credit filter on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'TokenLiquidationThresholdUpdated(address,uint16)',
            '#Token_3',
          ],
          data: [
            'uint16:8000',
          ],
          txHash: '!#Hash_6',
        },
        {
          // multicall start
          address: '#CreditFacade_1',
          topics: [
            'MultiCallStarted(address)',
            '#User_1',
          ],
          txHash: '!#Hash_7',
        },
        {
          // credit manager on usdc
          address: '#CreditFacade_1',
          topics: [
            'AddCollateral(address,address,uint256)',
            '#User_1',
            '#Token_3',
          ],
          data: [
            utils.bigIntTopic(newCollateral, 18),
          ],
          txHash: '!#Hash_7',
        },
        {
          // credit manager on usdc increaseborrow
          address: '#CreditFacade_1',
          topics: [
            'DecreaseBorrowedAmount(address,uint256)',
            // borrower
            '#User_1',
          ],
          data: [
            // amount
            utils.bigIntTopic(extraBorrowedAmount / 2, 6),
          ],
          txHash: '!#Hash_7',
        },
        {
          // multicall end
          address: '#CreditFacade_1',
          topics: [
            'MultiCallFinished()',
          ],
          txHash: '!#Hash_7',
        },
      ],
      calls: {
        masks: [{
          account: '#Account_1',
          mask: '3',
        }],
        accounts: [{
          address: '#Account_1',
          creditManager: '#CreditManager_1',
          borrower: '#User_1',
          healthFactor: '15444',  // 6950/4500
          totalValue: utils.bigInt(borrowedAmount + initialAmount + (extraBorrowedAmount / 2) + newCollateral * 2500, 6),
          repayAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount / 2, 6),
          cumulativeIndexAtOpen: utils.bigInt(1, 27),
          borrowedAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount / 2, 6),
          balances: [{
            token: '#Token_1',
            balance: utils.bigInt(5500, 6),
            isAllowed: true,
          }, {
            token: '#Token_3',
            balance: utils.bigInt(1, 18),
            isAllowed: true,
          }],
        }],
        cms: [{
          address: '#CreditManager_1',
          isWETH: false,
          minAmount: utils.bigInt(1000, 6),
          maxAmount: utils.bigInt(6000, 6),
          availableLiquidity: utils.bigInt(5500, 6),
          borrowRate: '0',
        }],
      },
    },
    // swap on uniswap v2
    '6': {
      events: [
        {
          // multicall start
          address: '#CreditFacade_1',
          topics: [
            'MultiCallStarted(address)',
            '#User_1',
          ],
          txHash: '!#Hash_8',
        },
        {
          // credit filter on usdc
          address: '#CreditManager_1',
          topics: [
            'ExecuteOrder(address,address)',
            '#User_1',
            '#Uniswapv2_1',
          ],
          txHash: '!#Hash_8',
        },
        {
          // credit filter on usdc
          address: '#CreditManager_1',
          topics: [
            'ExecuteOrder(address,address)',
            '#User_1',
            '#Uniswapv2_1',
          ],
          txHash: '!#Hash_8',
        },
        {
          // multicall end
          address: '#CreditFacade_1',
          topics: [
            'MultiCallFinished()',
          ],
          txHash: '!#Hash_8',
        },
      ],
      calls: {
        masks: [{
          account: '#Account_1',
          mask: '3',
        }],
        accounts: [{
          address: '#Account_1',
          creditManager: '#CreditManager_1',
          borrower: '#User_1',
          healthFactor: '14333',
          totalValue: utils.bigInt(500 + 3 * 2500, 6),
          repayAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount / 2, 6),
          cumulativeIndexAtOpen: utils.bigInt(1, 27),
          borrowedAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount / 2, 6),
          balances: [{
            token: '#Token_1',
            balance: utils.bigInt(500, 6),
            isAllowed: true,
          }, {
            token: '#Token_3',
            balance: utils.bigInt(3, 18),
            isAllowed: true,
          }],
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
    // direct token transfer for yfi and usdc
    // yfi is enabled on creditfacade but not linked
    // hence will not be considered while calculating total value
    '7': {
      events: [
        {
          // credit filter on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'TokenAllowed(address)',
            '#Token_2',
          ],
          txHash: '!#Hash_9',
        },
        {
          // credit filter on usdc
          address: '#CreditConfigurator_1',
          topics: [
            'TokenLiquidationThresholdUpdated(address,uint16)',
            '#Token_2',
          ],
          data: [
            'uint16:9000',
          ],
          txHash: '!#Hash_9',
        },
        {
          // YFI price chainlink
          address: '#ChainlinkPriceFeed_2',
          txHash: '!#Hash_9',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // 8
            utils.bigIntTopic(20000, 8),
            // roundid
            utils.bigIntTopic(1, 0),
          ],
        },
        {
          // direc token transfer for USDC
          address: '#Token_1',
          topics: [
            'Transfer(address,address,uint256)',
            '#User_3',
            '#Account_1',
          ],
          data: [
            utils.bigIntTopic(1000, 6),
          ],
          txHash: '!#Hash_10',
        },
        {
          // direc token transfer for YFI
          address: '#Token_2',
          topics: [
            'Transfer(address,address,uint256)',
            '#User_3',
            '#Account_1',
          ],
          data: [
            utils.bigIntTopic(0.1, 18),
          ],
          txHash: '!#Hash_11',
        },
      ],
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
