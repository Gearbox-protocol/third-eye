local utils = import '../utils.libsonnet';
local borrowedAmount = 4000;
local extraBorrowedAmount = 1000;
{
  blocks: {
    '7': {
      events: [
        {
          // repay credit account
          address: '#CreditManager_1',
          topics: [
            'RepayCreditAccount(address,address)',
            '#User_1',
            '#User_1',
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
            utils.bigIntTopic(5000, 6),
            // profit
            utils.bigIntTopic(0, 0),
            // loss
            utils.bigIntTopic(0, 0),
          ],
          txHash: '!#Hash_13',
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
            totalBorrowed: utils.bigInt(0, 6),
            expectedLiquidity: utils.bigInt(borrowedAmount + extraBorrowedAmount + 1000, 6),
            availableLiquidity: utils.bigInt(6000, 6),
            depositAPY: utils.bigInt(0),
            baseBorrowRate: utils.bigInt(0),
            dieselRate: utils.bigInt(0),
            withdrawFee: '0',
            linearCumulativeIndex: utils.bigInt(1, 27),
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
