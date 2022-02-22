package framework

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"sort"
)

type TestClient struct {
	// Blocks map[int64]BlockInput
	blockNums []int64
	events    map[int64]map[string][]types.Log
	prices    map[string]map[int64]*big.Int
	masks     map[int64]map[string]*big.Int
	state     *StateStore
	USDCAddr  string
	WETHAddr  string
	token     map[string]int8
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
		token:  map[string]int8{},
	}
}
func (t *TestClient) AddToken(tokenAddr string, decimals int8) {
	t.token[tokenAddr] = decimals
}
func (t *TestClient) setEvents(obj map[int64]map[string][]types.Log) {
	t.events = obj
	for blockNum := range obj {
		t.blockNums = append(t.blockNums, blockNum)
	}
	sort.Slice(t.blockNums, func(i, j int) bool { return t.blockNums[i] < t.blockNums[j] })
}

func (t *TestClient) setPrices(obj map[string]map[int64]*big.Int) {
	t.prices = obj
}

func (t *TestClient) setMasks(masks map[int64]map[string]*big.Int) {
	t.masks = masks
}
func (t *TestClient) setState(state *StateStore) {
	t.state = state
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
func topic(v string) common.Hash {
	return crypto.Keccak256Hash([]byte(v))
}
func ContainsHash(list []common.Hash, v common.Hash) bool {
	for _, hash := range list {
		if hash == v {
			return true
		}
	}
	return false
}
func (t *TestClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	toBlock := query.ToBlock.Int64()
	txLogs := []types.Log{}
	for i := query.FromBlock.Int64(); i < toBlock; i++ {
		for _, address := range query.Addresses {
			if t.events[i] != nil {
				if len(query.Topics) > 0 && query.Topics[0][0] == topic("Transfer(address,address,uint256)") {
					for _, txLog := range t.events[i][address.Hex()] {
						if ContainsHash(query.Topics[2], txLog.Topics[2]) {
							txLogs = append(txLogs, txLog)
						}
					}
				} else {
					txLogs = append(txLogs, t.events[i][address.Hex()]...)
				}
			}
		}
	}
	txLogList := TxLogList(txLogs)
	sort.Sort(txLogList)
	return []types.Log(txLogList), nil
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
	sig := hex.EncodeToString(call.Data[:4])
	blockNum := blockNumber.Int64()
	// convert on priceOracle
	if sig == "b66102df" {
		s := 4
		amount, ok := new(big.Int).SetString(hex.EncodeToString(call.Data[s:s+32]), 16)
		if !ok {
			log.Fatal("failed in parsing int")
		}
		s += 32
		token0 := common.BytesToAddress(call.Data[s : s+32]).Hex()
		decimalT0 := t.token[token0]
		s += 32
		token1 := common.BytesToAddress(call.Data[s : s+32]).Hex()
		decimalT1 := t.token[token1]
		price0 := t.getPrice(blockNum, token0)
		price1 := t.getPrice(blockNum, token1)
		newAmount := new(big.Int).Mul(amount, price0)
		newAmount = utils.GetInt64(newAmount, decimalT0-decimalT1)
		newAmount = new(big.Int).Quo(newAmount, price1)
		return common.HexToHash(fmt.Sprintf("%x", newAmount)).Bytes(), nil
		// enabledmask on creditfilter for account
	} else if sig == "b451cecc" {
		s := 4
		account := common.BytesToAddress(call.Data[s : s+32]).Hex()
		mask := t.masks[blockNum][account]
		return common.HexToHash(fmt.Sprintf("%x", mask)).Bytes(), nil
		// phaseId
	} else if sig == "58303b10" {
		oracle := call.To.Hex()
		index := t.state.Oracle.GetIndex(oracle, blockNum)
		return common.HexToHash(fmt.Sprintf("%x", index)).Bytes(), nil
		// current phase aggregator
	} else if sig == "c1597304" {
		s := 4
		index, ok := new(big.Int).SetString(hex.EncodeToString(call.Data[s:s+32]), 16)
		if !ok {
			log.Fatal("oracle:%s data: %s", call.To, call.Data)
		}
		oracle := call.To.Hex()
		feed := t.state.Oracle.GetState(oracle, int(index.Int64()))
		return common.HexToHash(feed.Feed).Bytes(), nil
	}
	return nil, nil
}
func (t *TestClient) getPrice(blockNum int64, tokenAddr string) *big.Int {
	if tokenAddr == t.WETHAddr {
		value, _ := new(big.Int).SetString("1000000000000000000", 10)
		return value
	} else {
		if t.prices[tokenAddr] != nil {
			var lastprice *big.Int
			for currentNum, price := range t.prices[tokenAddr] {
				if currentNum <= blockNum {
					lastprice = price
				}
			}
			return lastprice
		} else {
			log.Fatalf("token(%s) price not present", tokenAddr)
		}
	}
	return nil
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
