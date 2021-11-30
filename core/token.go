package core

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/eRC20"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Address  string            `gorm:"primaryKey;column:address"`
	Symbol   string            `gorm:"column:symbol"`
	Decimals uint8             `gorm:"column:decimals"`
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
		log.Fatal(err)
	}
	if symbol, err := contract.Symbol(&bind.CallOpts{}); err != nil {
		log.Fatal(err)
	} else {
		t.Symbol = symbol
	}
	if decimals, err := contract.Decimals(&bind.CallOpts{}); err != nil {
		log.Fatal(err)
	} else {
		t.Decimals = decimals
	}
}

type AllowedToken struct {
	Id                 int64  `gorm:"primaryKey;column:id;autoIncrement:true"`
	CreditManager      string `gorm:"column:credit_manager"`
	Token              string `gorm:"column:token"`
	LiquidityThreshold string `gorm:"column:liquiditythreshold"`
}

func (AllowedToken) TableName() string {
	return "allowed_tokens"
}
