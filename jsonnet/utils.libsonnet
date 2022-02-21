{
    bigInt(x, decimals=0)::
        std.toString(x * std.pow(10, decimals)),
    bigIntTopic(x, decimals)::
        'bigint:' + self.bigInt(x, decimals),
}