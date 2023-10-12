///////
// TEST
///////
// test open with multicall and additional multicall in same tx in block 4
// test open without multicall and close with multicall in same tx in block 6

local utils = import '../utils.libsonnet';
local initialAmount = 1000;
local borrowedAmount = 4000;
local moreCollateral = 1000;
// price are in usd
{
  executeParser: {
    '4': {
      mainEventLogs: {
        '!#Hash_1': [{
          name: 'openCreditAccountMulticall',
          len: 1,
        }, {
          name: 'multicall',
          len: 1,
        }],
      },
    },
    '5': {
      executeTransfers: {
        '!#Hash_2': {
          '#Token_1': utils.bigInt(3000, 6),
        },
      },
    },
    '6': {
      mainEventLogs: {
        '!#Hash_2': [{
          name: 'closeCreditAccount',
          len: 1,
        }],
      },
    },
  },
  states: {
    otherCalls: {
      '54fd4d50': { '#AddressProvider_1': '1' },
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
          txHash: '!#Hash_1',
        },
        {
          // multicall start
          address: '#CreditFacade_1',
          topics: [
            'MultiCallStarted(address)',
            '#User_1',
          ],
          txHash: '!#Hash_1',
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
          txHash: '!#Hash_1',
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
          txHash: '!#Hash_1',
        },
        {
          // multicall end
          address: '#CreditFacade_1',
          topics: [
            'MultiCallFinished()',
          ],
          txHash: '!#Hash_1',
        },
        {
          // multicall start
          address: '#CreditFacade_1',
          topics: [
            'MultiCallStarted(address)',
            '#User_1',
          ],
          txHash: '!#Hash_1',
        },
        {
          // pool on usdc
          address: '#CreditFacade_1',
          topics: [
            'AddCollateral(address,address,uint256)',
            '#User_1',
            '#Token_1',
          ],
          data: [
            utils.bigIntTopic(moreCollateral, 6),
          ],
          txHash: '!#Hash_1',
        },
        {
          // multicall end
          address: '#CreditFacade_1',
          topics: [
            'MultiCallFinished()',
          ],
          txHash: '!#Hash_1',
        },
      ],
      calls:
        {
          pools: [{
            address: '#Pool_1',
            totalBorrowed: utils.bigInt(borrowedAmount, 6),
            expectedLiquidity: utils.bigInt(borrowedAmount + 5000, 6),
            availableLiquidity: utils.bigInt(5000, 6),
            depositAPY: utils.bigInt(0),
            baseBorrowRate: utils.bigInt(0),
            dieselRate: utils.bigInt(0),
            withdrawFee: '0',
            cumulativeIndex: utils.bigInt(1, 27),
          }],
          accounts: [{
            address: '#Account_1',
            creditManager: '#CreditManager_1',
            borrower: '#User_1',
            healthFactor: '13500',  //  .9*6000/4000 //borrowed twice after opening
            totalValue: utils.bigInt(borrowedAmount + moreCollateral + initialAmount, 6),
            repayAmount: utils.bigInt(borrowedAmount, 6),
            cumulativeIndexAtOpen: utils.bigInt(1, 27),
            borrowedAmount: utils.bigInt(borrowedAmount, 6),
            borrowedAmountPlusInterest: utils.bigInt(borrowedAmount, 6),
            balances: [{
              token: '#Token_1',
              BI: utils.bigInt(6000, 6),
              isForbidden: false,  // changed
              isEnabled: true,
            }],
            version: 2,
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
    '5': {
      calls: {
        accounts: [{
          address: '#Account_1',
          creditManager: '#CreditManager_1',
          borrower: '#User_1',
          healthFactor: '15750',  // .9 * 7000/ 4000
          totalValue: utils.bigInt(borrowedAmount + moreCollateral * 2 + initialAmount, 6),
          repayAmount: utils.bigInt(borrowedAmount, 6),
          cumulativeIndexAtOpen: utils.bigInt(1, 27),
          borrowedAmount: utils.bigInt(borrowedAmount, 6),
          borrowedAmountPlusInterest: utils.bigInt(borrowedAmount, 6),
          balances: [
            {
              token: '#Token_1',
              BI: utils.bigInt(7000, 6),
              isForbidden: false,  // changed
              isEnabled: true,
            },
          ],
          version: 2,
        }],
      },
    },
    '6': {
      events: [
        {
          // multicall start
          address: '#CreditFacade_1',
          topics: [
            'MultiCallStarted(address)',
            '#User_1',
          ],
          txHash: '!#Hash_2',
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
            utils.bigIntTopic(moreCollateral, 6),
          ],
          txHash: '!#Hash_2',
        },
        {
          // multicall end
          address: '#CreditFacade_1',
          topics: [
            'MultiCallFinished()',
          ],
          txHash: '!#Hash_2',
        },
        {
          // close credit account
          address: '#CreditManager_1',
          topics: [
            'CloseCreditAccount(address,address)',
            '#User_1',  // borrower
            // to
            '#User_1',
          ],
          txHash: '!#Hash_2',
        },
        {
          // open account without multicall
          address: '#CreditFacade_1',
          topics: [
            'OpenCreditAccount(address,address,uint256,uint16)',
            '#User_2',
            '#Account_2',
          ],
          data: [
            utils.bigIntTopic(borrowedAmount, 6),
            'uint16:0',
          ],
          txHash: '!#Hash_2',
        },
        {
          // pool on usdc
          address: '#Pool_1',
          topics: [
            'Borrow(address,address,uint256)',
            '#CreditManager_1',
            '#Account_2',
          ],
          data: [
            utils.bigIntTopic(borrowedAmount, 6),
          ],
          txHash: '!#Hash_2',
        },
        {
          // add collateral
          address: '#CreditFacade_1',
          topics: [
            'AddCollateral(address,address,uint256)',
            '#User_2',
            '#Token_1',
          ],
          data: [
            utils.bigIntTopic(moreCollateral, 6),
          ],
          txHash: '!#Hash_2',
        },
        {
          // add collateral
          address: '#CreditFacade_1',
          topics: [
            'AddCollateral(address,address,uint256)',
            '#User_2',
            '#Token_1',
          ],
          data: [
            utils.bigIntTopic(moreCollateral, 6),
          ],
          txHash: '!#Hash_2',
        },
      ],
      calls: {
        pools: [{
          address: '#Pool_1',
          totalBorrowed: utils.bigInt(borrowedAmount, 6),
          expectedLiquidity: utils.bigInt(borrowedAmount + 5000, 6),
          availableLiquidity: utils.bigInt(5000, 6),
          depositAPY: utils.bigInt(0),
          baseBorrowRate: utils.bigInt(0),
          dieselRate: utils.bigInt(0),
          withdrawFee: '0',
          cumulativeIndex: utils.bigInt(1, 27),
        }],
        accounts: [{
          address: '#Account_2',
          creditManager: '#CreditManager_1',
          borrower: '#User_2',
          healthFactor: '13500',  // .9 * 6000/ 4000
          totalValue: utils.bigInt(borrowedAmount + moreCollateral + initialAmount, 6),  //borrowed once more after opening
          repayAmount: utils.bigInt(borrowedAmount, 6),
          cumulativeIndexAtOpen: utils.bigInt(1, 27),
          borrowedAmount: utils.bigInt(borrowedAmount, 6),
          borrowedAmountPlusInterest: utils.bigInt(borrowedAmount, 6),
          balances: [
            {
              token: '#Token_1',
              BI: utils.bigInt(6000, 6),
              isForbidden: false,  // changed
              isEnabled: true,
            },
          ],
          version: 2,
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
  },
}
