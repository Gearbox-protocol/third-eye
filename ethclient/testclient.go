package ethclient

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sort"
)

type TestClient struct {
	// Blocks map[int64]BlockInput
	blockNums []int64
	events    map[int64]map[string][]types.Log
	prices    map[int64]map[string]*big.Int
	masks     map[int64]map[string]*big.Int
	USDCAddr  string
	WETHAddr  string
}

func (t *TestClient) SetUSDC(addr string) {
	t.USDCAddr = addr
}
func (t *TestClient) SetWETH(addr string) {
	t.WETHAddr = addr
}
func NewTestClient() *TestClient {
	return &TestClient{
		events: make(map[int64]map[string][]types.Log),
	}
}
func (t *TestClient) SetEvents(obj map[int64]map[string][]types.Log) {
	t.events = obj
	for blockNum := range obj {
		t.blockNums = append(t.blockNums, blockNum)
	}
	sort.Slice(t.blockNums, func(i, j int) bool { return t.blockNums[i] < t.blockNums[j] })
}

func (t *TestClient) SetPrices(obj map[int64]map[string]*big.Int) {
	t.prices = obj
}

func (t *TestClient) SetMasks(masks map[int64]map[string]*big.Int) {
	t.masks = masks
}

func (t *TestClient) ChainID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1337), nil
}
func (t *TestClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return types.NewBlock(&types.Header{Time: uint64(number.Int64()) * 86400},
		[]*types.Transaction{},
		[]*types.Header{},
		[]*types.Receipt{}, nil), nil
}
func (t *TestClient) BlockNumber(ctx context.Context) (uint64, error) {
	if len(t.blockNums) == 0 {
		return 1, nil
	}
	return uint64(t.blockNums[len(t.blockNums)-1]), nil
}

func (t *TestClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	toBlock := query.ToBlock.Int64()
	txLogs := []types.Log{}
	for i := query.FromBlock.Int64(); i < toBlock; i++ {
		for _, address := range query.Addresses {
			if t.events[i] != nil {
				txLogs = append(txLogs, t.events[i][address.Hex()]...)
			}
		}
	}
	return txLogs, nil
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
	// convert on priceOracle
	sig := hex.EncodeToString(call.Data[:4])
	blockNum := blockNumber.Int64()
	if sig == "b66102df" {
		s := 4
		amount, ok := new(big.Int).SetString(hex.EncodeToString(call.Data[s:s+32]), 16)
		if !ok {
			log.Fatal("failed in parsing int")
		}
		s += 32
		token0 := common.BytesToAddress(call.Data[s : s+32]).Hex()
		s += 32
		token1 := common.BytesToAddress(call.Data[s : s+32]).Hex()
		price0 := t.getPrice(blockNum, token0)
		price1 := t.getPrice(blockNum, token1)
		newAmount := new(big.Int).Mul(amount, price0)
		newAmount = new(big.Int).Quo(newAmount, price1)
		return common.HexToHash(fmt.Sprintf("%x", newAmount)).Bytes(), nil
		// enabledmask on creditfilter for account
	} else if sig == "b451cecc" {
		s := 4
		account := common.BytesToAddress(call.Data[s : s+32]).Hex()
		mask := t.masks[blockNum][account]
		return common.HexToHash(fmt.Sprintf("%x", mask)).Bytes(), nil
	}
	return nil, nil
}
func (t *TestClient) getPrice(blockNum int64, addr string) *big.Int {
	if addr == "WETH" {
		value, _ := new(big.Int).SetString("1000000000000000000", 10)
		return value
	} else {
		return t.prices[blockNum][addr]
	}
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

func (t *TestClient) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}
