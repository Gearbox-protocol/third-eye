local utils = import '../utils.libsonnet';
local borrowedAmount = 4000;
local extraBorrowedAmount = 1000;
local initialAmount = 1000;
local newCollateral = 1;
{
  mocks: {
    syncAdapters: 'mocks/yearn_feed.json',
  },
  blocks: {
    '4': {
      calls: {
        others: {
          feaf968c: {
            '#YearnFeed_1': utils.bigIntTopic(1004, 0),
            '#YearnFeed_3': utils.bigIntTopic(3004, 0),
          },
        },
      },
    },
    '11': {
      calls: {
        others: {
          feaf968c: {
            '#YearnFeed_1': utils.bigIntTopic(1011, 0),
          },
        },
      },
    },
    '26': {
      calls: {
        others: {
          feaf968c: {
            '#YearnFeed_1': utils.bigIntTopic(1026, 0),
            '#YearnFeed_3': utils.bigIntTopic(3026, 0),
          },
        },
      },
    },
    '31': {
      calls: {
        others: {
          feaf968c: {
            '#YearnFeed_3': utils.bigIntTopic(3031, 0),
          },
        },
      },
    },
    '51': {
      calls: {
        others: {
          feaf968c: {
            '#YearnFeed_1': utils.bigIntTopic(1051, 0),
            '#YearnFeed_3': utils.bigIntTopic(3051, 0),
          },
        },
      },
    },
    '53': {
      calls: {
        others: {
          feaf968c: {
            '#YearnFeed_1': utils.bigIntTopic(1053, 0),
            '#YearnFeed_3': utils.bigIntTopic(3053, 0),
          },
        },
      },
    },
    '56': {
      calls: {
        others: {
          feaf968c: {
            '#YearnFeed_4': utils.latestRoundData('4056'),
          },
        },
      },
    },
    '58': {
      calls: {
        others: {
          feaf968c: {
            '#YearnFeed_3': utils.bigIntTopic(3058, 0),
            '#YearnFeed_4': utils.bigIntTopic(4058, 0),
          },
        },
      },
    },
  },
}
