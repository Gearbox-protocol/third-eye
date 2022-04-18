{
    bigInt(x, decimals=0)::
        if decimals <18 
        then
        std.toString(x * std.pow(10, decimals))
        else 
        std.toString(x * std.pow(10, 18))+std.toString(std.pow(10,decimals-18))[1:],
    bigIntTopic(x, decimals)::
        'bigint:' + self.bigInt(x, decimals),
    strToHex(msg)::
        std.join("", ['%x'% std.codepoint(i) for i in std.stringChars(msg)]),
}