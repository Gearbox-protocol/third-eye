package core

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/eRC20"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Address  string            `gorm:"primaryKey;column:address" json:"address"`
	Symbol   string            `gorm:"column:symbol" json:"symbol"`
	Decimals int8              `gorm:"column:decimals" json:"decimals"`
	client   ethclient.ClientI `gorm:"-" json:"-"`
}

func (Token) TableName() string {
	return "tokens"
}

func NewToken(addr string, client ethclient.ClientI) (*Token, error) {
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
	return utils.AlmostSameBigInt(a.Convert(), b.Convert(), token.Decimals-precision)
}

type TokenTransfer struct {
	BlockNum      int64   `gorm:"column:block_num;primaryKey"`
	LogID         uint    `gorm:"column:log_id;primaryKey"`
	TxHash        string  `gorm:"column:tx_hash"`
	Token         string  `gorm:"column:token"`
	From          string  `gorm:"column:source"`
	To            string  `gorm:"column:destination"`
	Amount        *BigInt `gorm:"column:amount"`
	IsFromAccount bool    `gorm:"column:isfrom_account"`
	IsToAccount   bool    `gorm:"column:isto_account"`
}

type TokenTransferList []*TokenTransfer

func (ts TokenTransferList) Len() int {
	return len(ts)
}
func (ts TokenTransferList) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// sort in increasing order by blockNumber,index
func (ts TokenTransferList) Less(i, j int) bool {
	return ts[i].LogID < ts[j].LogID
}
func (a *TokenTransfer) equal(b *TokenTransfer) bool {
	return a.BlockNum == b.BlockNum &&
		a.LogID == b.LogID &&
		a.TxHash == b.TxHash &&
		a.Token == b.Token &&
		a.From == b.From &&
		a.To == b.To &&
		a.Amount == b.Amount &&
		a.IsFromAccount == b.IsFromAccount &&
		a.IsToAccount == b.IsToAccount
}

func (TokenTransfer) TableName() string {
	return "no_session_transfers"
}
