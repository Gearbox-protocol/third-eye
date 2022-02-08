{adapters:[
    {
        type: "AddressProvider",
        address: "#AddressProvider1",
        lastSync: 1,
    },
    {
        type: "Pool",
        address: "#Pool1",
        lastSync: 2,
    },
    {
        type: "CreditManager",
        address: "#CreditManager1",
        lastSync: 2,
    },
    {
        type: "CreditFilter",
        address: "#CreditFilter1",
        details: {
            creditManager: "@CreditManager1",
        },
        lastSync: 2,
    },
    {
        type: "ChainlinkPriceFeed",
        address: "#ChainlinkPriceFeed1",
        details: {
            oracle: "#Oracle1",
            token: "#Token1",
        },
        lastSync: 2,
    },
    {
        type: "ChainlinkPriceFeed",
        address: "#ChainlinkPriceFeed2",
        details: {
            oracle: "#Oracle2",
            token: "#Token2",
        },
        lastSync: 2,
    },
]
}