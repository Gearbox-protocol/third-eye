package core

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/eRC20"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

func NewToken(addr string, client *ethclient.Client) (*Token, error) {
	token := &Token{
		Address: addr,
		client:  client,
	}
	err := token.init()
	return token, err
}

func (t *Token) init() error {
	contract, err := eRC20.NewERC20(common.HexToAddress(t.Address), t.client)
	if err != nil {
		return err
	}
	if symbol, err := contract.Symbol(&bind.CallOpts{}); err != nil {
		return err
	} else {
		t.Symbol = symbol
	}
	if decimals, err := contract.Decimals(&bind.CallOpts{}); err != nil {
		return err
	} else {
		t.Decimals = int8(decimals)
	}
	return nil
}

type AllowedToken struct {
	BlockNumber        int64   `gorm:"column:block_num;primaryKey"`
	CreditManager      string  `gorm:"column:credit_manager;primaryKey"`
	Token              string  `gorm:"column:token;primaryKey"`
	LiquidityThreshold *BigInt `gorm:"column:liquiditythreshold"`
	DisableBlock       int64   `gorm:"column:disable_block"`
}

func (AllowedToken) TableName() string {
	return "allowed_tokens"
}

func CompareBalance(a, b *BigInt, token *CumIndexAndUToken) bool {
	precision := utils.GetPrecision(token.Symbol)
	return utils.AlmostSameBigInt(a.Convert(), b.Convert(), token.Decimals, precision)
}
