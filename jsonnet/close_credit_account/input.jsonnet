local utils = import '../utils.libsonnet';
local borrowedAmount = 4000;
local extraBorrowedAmount = 1000;
{
  blocks: {
    '7': {
      events: [
        {
          // close credit account
          address: '#CreditManager_1',
          topics: [
            'CloseCreditAccount(address,address,uint256)',
            '#User_1',
            '#User_1',
          ],
          data: [
            utils.bigIntTopic(4500, 0),
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
            totalBorrowed: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
            expectedLiquidity: utils.bigInt(borrowedAmount + extraBorrowedAmount, 6),
            availableLiquidity: utils.bigInt(1000, 6),
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
            availableLiquidity: utils.bigInt(1000, 6),
            borrowRate: '0',
          }],
        },
    },
  },
}
