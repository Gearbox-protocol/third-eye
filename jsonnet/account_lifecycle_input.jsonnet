local utils = import 'utils.libsonnet';
local initialAmount = 1000;
local borrowedAmount = 4000;
local extraBorrowedAmount = 1000;
local newCollateral = 1;
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
    // block with open and borrow more underlying asset
    '3': {
      events: [
        {
          // credit filter on usdc
          address: '#CreditFilter_1',
          topics: [
            'TokenAllowed(address,uint256)',
            '#Token_1',
          ],
          data: [
            utils.bigIntTopic(9000, 0),
          ],
          txHash: '!#Hash_1',
        },
        {
          // price chainlink on usdc
          address: '#CreditManager_1',
          txHash: '!#Hash_2',
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
        },
        {
          // price chainlink on usdc
          address: '#ChainlinkPriceFeed_1',
          txHash: '!#Hash_3',
          topics: [
            'AnswerUpdated(int256,uint256,uint256)',
            // 0.0004
            utils.bigIntTopic(0.0004, 18),
            // roundid
            utils.bigIntTopic(1, 0),
          ],
          data: [],
        },
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
          txHash: '!#Hash_5',
        },
        {
          // credit manager on usdc increaseborrow
          address: '#CreditManager_1',
          topics: [
            'IncreaseBorrowedAmount(address,uint256)',
            '#User_1',  // borrower
          ],
          data: [
            utils.bigIntTopic(extraBorrowedAmount, 6),  // amount
          ],
          txHash: '!#Hash_6',
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
            utils.bigIntTopic(extraBorrowedAmount, 6),
          ],
          txHash: '!#Hash_7',
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
            expectedLiquidity: utils.bigInt(borrowedAmount + extraBorrowedAmount + 1000, 6),
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
            maxAmount: utils.bigInt(5000, 6),
            availableLiquidity: utils.bigInt(1000, 6),
            borrowRate: '0',
          }],
        },
    },
    // new block with add collateral
    '4': {
      events: [
        {
          // credit filter on usdc
          address: '#CreditFilter_1',
          topics: [
            'TokenAllowed(address,uint256)',
            '#Token_3',
          ],
          data: [
            utils.bigIntTopic(8000, 0),
          ],
          txHash: '!#Hash_8',
        },
        {
          // credit manager on usdc
          address: '#CreditManager_1',
          topics: [
            'AddCollateral(address,address,uint256)',
            '#User_1',
            '#Token_3',
          ],
          data: [
            utils.bigIntTopic(newCollateral, 18),
          ],
          txHash: '!#Hash_9',
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
          healthFactor: '14800',
          totalValue: utils.bigInt(borrowedAmount + initialAmount + extraBorrowedAmount + newCollateral * 2500, 6),
          repayAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
          cumulativeIndexAtOpen: utils.bigInt(1, 27),
          borrowedAmount: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
          balances: [{
            token: '#Token_1',
            balance: utils.bigInt(6000, 6),
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
          maxAmount: utils.bigInt(5000, 6),
          availableLiquidity: utils.bigInt(1000, 6),
          borrowRate: '0',
        }],
      },
    },
  },
}
