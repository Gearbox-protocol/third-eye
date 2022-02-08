package ethclient

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type TestClient struct {
	// Blocks map[int64]BlockInput
	blockNums []int64
	events    map[int64]map[string][]types.Log
}

func NewTestClient() *TestClient {
	return &TestClient{
		events: make(map[int64]map[string][]types.Log),
	}
}
func (t *TestClient) SetEvents(obj map[int64]map[string][]types.Log) {
	t.events = obj
}
func (t *TestClient) ChainID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1337), nil
}
func (t *TestClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return nil, nil
}
func (t *TestClient) BlockNumber(ctx context.Context) (uint64, error) {
	return 0, nil
}
func (t *TestClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return nil, false, nil
}
func (t *TestClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}

func (t *TestClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}
func (t *TestClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}
func (t *TestClient) PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error) {
	return nil, nil
}
func (t *TestClient) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	return nil, nil
}
func (t *TestClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return nil, nil
}
func (t *TestClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, nil
}
func (t *TestClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
func (t *TestClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
func (t *TestClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return 0, nil
}
func (t *TestClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}

func (t *TestClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

func (t *TestClient) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}
