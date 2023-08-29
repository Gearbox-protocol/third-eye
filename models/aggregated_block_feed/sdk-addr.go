package aggregated_block_feed

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

type TokenSymMap struct {
	symToAddr map[string]common.Address
	addrToSym map[common.Address]string
	chainId   int64
}

func (x TokenSymMap) UpdateForTest(sym string, token common.Address) {
	x.addrToSym[token] = sym
	x.symToAddr[sym] = token
}

func newTokenSymMap(chainId int64) TokenSymMap {
	if chainId == 1337 {
		return TokenSymMap{chainId: chainId}
	} else {
		return TokenSymMapFromchainId(chainId)
	}
}

func (m *TokenSymMap) updateIfTest(repo repoI) {
	if m.chainId != 1337 {
		return
	}
	symToAddr := map[string]common.Address{}
	addrToSym := map[common.Address]string{}
	for _, tokenStr := range repo.GetTokens() {
		sym := repo.GetToken(tokenStr).Symbol
		token := common.HexToAddress(tokenStr)
		symToAddr[sym] = token
		addrToSym[token] = sym
	}
	m.symToAddr = symToAddr
	m.addrToSym = addrToSym
}

func TokenSymMapFromchainId(chainId int64) TokenSymMap {
	symToAddr := core.GetSymToAddrByChainId(chainId)
	addrToSym := map[common.Address]string{}
	for sym, addr := range symToAddr.Tokens {
		addrToSym[addr] = sym
	}
	return TokenSymMap{
		symToAddr: symToAddr.Tokens,
		addrToSym: addrToSym,
		chainId:   chainId,
	}
}

func (x TokenSymMap) getTokenAddr(sym string) string {
	addr := x.symToAddr[sym]
	if addr == core.NULL_ADDR {
		log.Fatal("Token sym not found  in embedded jsonnet", sym)
	}
	return addr.Hex()
}
func (x TokenSymMap) getTokenSym(addr string) string {
	sym := x.addrToSym[common.HexToAddress(addr)]
	if sym == "" {
		log.Fatal("Token addr not found in embedded jsonnet", addr)
	}
	return sym
}
