package core

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/eRC20"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Token struct {
	Address  string            `gorm:"primaryKey;column:address"`
	Symbol   string            `gorm:"column:symbol"`
	Decimals int8              `gorm:"column:decimals"`
	client   *ethclient.Client `gorm:"-"`
}

func (Token) TableName() string {
	return "tokens"
}

func NewToken(addr string, client *ethclient.Client) *Token {
	token := &Token{
		Address: addr,
		client:  client,
	}
	token.init()
	return token
}

func (t *Token) init() {
	contract, err := eRC20.NewERC20(common.HexToAddress(t.Address), t.client)
	if err != nil {
		log.Fatal(err, t.Address)
	}
	if symbol, err := contract.Symbol(&bind.CallOpts{}); err != nil {
		log.Fatal(err, t.Address)
	} else {
		t.Symbol = symbol
	}
	if decimals, err := contract.Decimals(&bind.CallOpts{}); err != nil {
		log.Fatal(err, t.Address)
	} else {
		t.Decimals = int8(decimals)
	}
}

type AllowedToken struct {
	Id                 int64   `gorm:"primaryKey;column:id;autoIncrement:true"`
	BlockNumber        int64   `gorm:"column:block_num"`
	CreditManager      string  `gorm:"column:credit_manager"`
	Token              string  `gorm:"column:token"`
	LiquidityThreshold *BigInt `gorm:"column:liquiditythreshold"`
}

func (AllowedToken) TableName() string {
	return "allowed_tokens"
}

func CompareBalance(a, b *big.Int, token *Token) bool {
	precision := utils.GetPrecision(token.Symbol)
	return utils.AlmostSameBigInt(a, b, token.Decimals, precision)
}
