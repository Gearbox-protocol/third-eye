{
  bigInt(x, decimals=0)::
    if decimals < 18
    then
      std.toString(x * std.pow(10, decimals))
    else
      std.toString(x * std.pow(10, 18)) + std.toString(std.pow(10, decimals - 18))[1:],
  bigIntTopic(x, decimals)::
    'bigint:' + self.bigInt(x, decimals),
  strToHex(msg)::
    std.join('', ['%x' % std.codepoint(i) for i in std.stringChars(msg)]),
  zeros(x)::
    std.join('', std.makeArray(x, function(x) '0')),
  latestRoundData(x)::
    self.zeros(64) + self.zeros(64 - std.length(x)) + x + self.zeros(64 * 3),
}
