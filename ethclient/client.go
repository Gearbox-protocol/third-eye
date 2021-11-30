package ethclient

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	urls   []string
	index  int
	client *ethclient.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	urls := strings.Split(rawurl, ",")
	c := &Client{urls: urls}
	var err error
	c.client, err = ethclient.Dial(c.urls[c.index])
	return c, err
}

func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	urls := strings.Split(rawurl, ",")
	c := &Client{urls: urls}
	var err error
	c.client, err = ethclient.DialContext(ctx, c.urls[c.index])
	return c, err
}
func (rc *Client) UpdateClient() error {
	// close eth client
	rc.client.Close()
	// get new eth client
	rc.index = (rc.index + 1) % len(rc.urls)
	log.Info("New rpc url: ", rc.urls[rc.index])
	var err error
	rc.client, err = ethclient.Dial(rc.urls[rc.index])
	return err
}

func (rc *Client) Close() {
	rc.client.Close()
}
func (rc *Client) errorHandler(err error) bool {
	if err != nil {
		if err.Error() == "execution aborted (timeout = 10s)" {
			log.Error("sleeping due to execution aborted (timeout = 10s)")
			time.Sleep(2 * time.Second)
		}
		if strings.HasPrefix(err.Error(), "403") {
			log.Error("Retry because of error: ", err)
			if err = rc.UpdateClient(); err != nil {
				log.Error("Next rpc connection failed: ", err)
			}
		} else if strings.HasPrefix(err.Error(), "429") {
			log.Error("sleep because of error: ", err)
			time.Sleep(20 * time.Second)
		}
	} else {
		return false
	}
	return true
}
func (rc *Client) ChainID(ctx context.Context) (*big.Int, error) {
	v, err := rc.client.ChainID(ctx)
	if rc.errorHandler(err) {
		v, err = rc.client.ChainID(ctx)
	}
	return v, err
}

func (rc *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	v, err := rc.client.BlockByHash(ctx, hash)
	if rc.errorHandler(err) {
		v, err = rc.client.BlockByHash(ctx, hash)
	}
	return v, err
}

func (rc *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	v, err := rc.client.BlockByNumber(ctx, number)
	if rc.errorHandler(err) {
		v, err = rc.client.BlockByNumber(ctx, number)
	}
	return v, err
}

func (rc *Client) BlockNumber(ctx context.Context) (uint64, error) {
	v, err := rc.client.BlockNumber(ctx)
	if rc.errorHandler(err) {
		v, err = rc.client.BlockNumber(ctx)
	}
	return v, err
}

func (rc *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	v, err := rc.client.HeaderByHash(ctx, hash)
	if rc.errorHandler(err) {
		v, err = rc.client.HeaderByHash(ctx, hash)
	}
	return v, err
}

func (rc *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	v, err := rc.client.HeaderByNumber(ctx, number)
	if rc.errorHandler(err) {
		v, err = rc.client.HeaderByNumber(ctx, number)
	}
	return v, err
}

// TransactionByHash returns the transaction with the given hash.
func (rc *Client) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	a, b, err := rc.client.TransactionByHash(ctx, hash)
	if rc.errorHandler(err) {
		a, b, err = rc.client.TransactionByHash(ctx, hash)
	}
	return a, b, err
}

func (rc *Client) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	v, err := rc.client.TransactionSender(ctx, tx, block, index)
	if rc.errorHandler(err) {
		v, err = rc.client.TransactionSender(ctx, tx, block, index)
	}
	return v, err
}

// TransactionCount returns the total number of transactions in the given block.
func (rc *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	v, err := rc.client.TransactionCount(ctx, blockHash)
	if rc.errorHandler(err) {
		v, err = rc.client.TransactionCount(ctx, blockHash)
	}
	return v, err
}

// TransactionInBlock returns a single transaction at index in the given block.
func (rc *Client) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	v, err := rc.client.TransactionInBlock(ctx, blockHash, index)
	if rc.errorHandler(err) {
		v, err = rc.client.TransactionInBlock(ctx, blockHash, index)
	}
	return v, err
}

func (rc *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	v, err := rc.client.TransactionReceipt(ctx, txHash)
	if rc.errorHandler(err) {
		v, err = rc.client.TransactionReceipt(ctx, txHash)
	}
	return v, err
}

func (rc *Client) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	v, err := rc.client.SyncProgress(ctx)
	if rc.errorHandler(err) {
		v, err = rc.client.SyncProgress(ctx)
	}
	return v, err
}

func (rc *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	v, err := rc.client.SubscribeNewHead(ctx, ch)
	if rc.errorHandler(err) {
		v, err = rc.client.SubscribeNewHead(ctx, ch)
	}
	return v, err
}

func (rc *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	v, err := rc.client.NetworkID(ctx)
	if rc.errorHandler(err) {
		v, err = rc.client.NetworkID(ctx)
	}
	return v, err
}

func (rc *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	v, err := rc.client.BalanceAt(ctx, account, blockNumber)
	if rc.errorHandler(err) {
		v, err = rc.client.BalanceAt(ctx, account, blockNumber)
	}
	return v, err
}

func (rc *Client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	v, err := rc.client.StorageAt(ctx, account, key, blockNumber)
	if rc.errorHandler(err) {
		v, err = rc.client.StorageAt(ctx, account, key, blockNumber)
	}
	return v, err
}

func (rc *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	v, err := rc.client.CodeAt(ctx, account, blockNumber)
	if rc.errorHandler(err) {
		v, err = rc.client.CodeAt(ctx, account, blockNumber)
	}
	return v, err
}

func (rc *Client) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	v, err := rc.client.NonceAt(ctx, account, blockNumber)
	if rc.errorHandler(err) {
		v, err = rc.client.NonceAt(ctx, account, blockNumber)
	}
	return v, err
}

func (rc *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	v, err := rc.client.FilterLogs(ctx, q)
	if rc.errorHandler(err) {
		v, err = rc.client.FilterLogs(ctx, q)
	}
	return v, err
}

func (rc *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	v, err := rc.client.SubscribeFilterLogs(ctx, q, ch)
	if rc.errorHandler(err) {
		v, err = rc.client.SubscribeFilterLogs(ctx, q, ch)
	}
	return v, err
}

func (rc *Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	v, err := rc.client.PendingBalanceAt(ctx, account)
	if rc.errorHandler(err) {
		v, err = rc.client.PendingBalanceAt(ctx, account)
	}
	return v, err
}

func (rc *Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	v, err := rc.client.PendingStorageAt(ctx, account, key)
	if rc.errorHandler(err) {
		v, err = rc.client.PendingStorageAt(ctx, account, key)
	}
	return v, err
}

func (rc *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	v, err := rc.client.PendingCodeAt(ctx, account)
	if rc.errorHandler(err) {
		v, err = rc.client.PendingCodeAt(ctx, account)
	}
	return v, err
}

func (rc *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	v, err := rc.client.PendingNonceAt(ctx, account)
	if rc.errorHandler(err) {
		v, err = rc.client.PendingNonceAt(ctx, account)
	}
	return v, err
}

func (rc *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
	v, err := rc.client.PendingTransactionCount(ctx)
	if rc.errorHandler(err) {
		v, err = rc.client.PendingTransactionCount(ctx)
	}
	return v, err
}

func (rc *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	v, err := rc.client.CallContract(ctx, msg, blockNumber)
	if rc.errorHandler(err) {
		v, err = rc.client.CallContract(ctx, msg, blockNumber)
	}
	return v, err
}

func (rc *Client) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	v, err := rc.client.PendingCallContract(ctx, msg)
	if rc.errorHandler(err) {
		v, err = rc.client.PendingCallContract(ctx, msg)
	}
	return v, err
}

func (rc *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	v, err := rc.client.SuggestGasPrice(ctx)
	if rc.errorHandler(err) {
		v, err = rc.client.SuggestGasPrice(ctx)
	}
	return v, err
}

func (rc *Client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	v, err := rc.client.SuggestGasTipCap(ctx)
	if rc.errorHandler(err) {
		v, err = rc.client.SuggestGasTipCap(ctx)
	}
	return v, err
}

func (rc *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	v, err := rc.client.EstimateGas(ctx, msg)
	if rc.errorHandler(err) {
		v, err = rc.client.EstimateGas(ctx, msg)
	}
	return v, err
}

func (rc *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := rc.client.SendTransaction(ctx, tx)
	if rc.errorHandler(err) {
		err = rc.client.SendTransaction(ctx, tx)
	}
	return err
}
