package v2

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MockRepo struct {
	ds.DummyRepo
	tokendata map[string]*schemas.Token
	data      map[string]*schemas.UTokenAndPool
	client    core.ClientI
}

func NewMockRepo(client core.ClientI, data map[string]*schemas.UTokenAndPool) *MockRepo {
	return &MockRepo{
		tokendata: map[string]*schemas.Token{},
		data:      data,
		client:    client,
	}
}

func (r MockRepo) GetToken(tokenAddr string) *schemas.Token {
	if r.tokendata[tokenAddr] == nil {
		tokenInfo, err := schemas.NewToken(tokenAddr, r.client)
		log.CheckFatal(err)
		r.tokendata[tokenAddr] = tokenInfo
	}
	return r.tokendata[tokenAddr]
}

func (r MockRepo) GetDieselTokens() map[string]*schemas.UTokenAndPool {
	return r.data
}

func skipIfETHProviderMissing(t *testing.T) string {
	ethProvider := utils.GetEnvOrDefault("ETH_PROVIDER", "")
	if ethProvider == "" {
		t.Skip("Skipping testing if eth provider missing")
	}
	return ethProvider
}

func TestRewardCalc(t *testing.T) {
	data := map[string]*schemas.UTokenAndPool{
		"0xc411dB5f5Eb3f7d552F9B8454B2D74097ccdE6E3": {Pool: "0x86130bDD69143D8a4E5fc50bf4323D48049E98E4"}, //usdc
		// "0x8A1112AFef7F4FC7c066a77AABBc01b3Fff31D47": {Pool: "0x79012c8d491DcF3A30Db20d1f449b14CAF01da6C"},
		// "0x2158034dB06f06dcB9A786D2F1F8c38781bA779d": {Pool: "0xB8cf3Ed326bB0E51454361Fb37E9E8df6DC5C286"},
		// "0xF21fc650C1B34eb0FDE786D52d23dA99Db3D6278": {Pool: "0xB03670c20F87f2169A7c4eBE35746007e9575901"},
		// "0xe753260F1955e8678DCeA8887759e07aa57E8c54": {Pool: "0xB2A015c71c17bCAC6af36645DEad8c572bA08A08"},
		// "0x6CFaF95457d7688022FC53e7AbE052ef8DFBbdBA": {Pool: "0x24946bCbBd028D5ABb62ad9B635EB1b1a67AF668"},
	}
	//
	client, err := ethclient.Dial(skipIfETHProviderMissing(t))
	log.CheckFatal(err)
	repo := NewMockRepo(client, data)
	//
	addrs := []common.Address{}
	for addr := range data {
		addrs = append(addrs, common.HexToAddress(addr))
	}
	//
	obj := NewLMRewardsv2(core.NULL_ADDR.Hex(), 13810899, client, repo)
	obj.GetAllAddrsForLogs()
	//
	var till int64 = 16925064
	var batch int64 = 50_000
	var batchStart int64 = 13810899
	//
	for ; batchStart < till; batchStart += batch {
		batchEnd := batchStart + batch
		if batchStart+batch >= till {
			batchEnd = till
		}
		txLogs, err := pkg.Node{Client: client}.GetLogs(batchStart+1, batchEnd, addrs, [][]common.Hash{{core.Topic("Transfer(address,address,uint256)")}})
		log.CheckFatal(err)
		for _, txLog := range txLogs {
			obj.OnLog(txLog)
		}
		obj.AfterSyncHook(batchEnd)
		log.Info(batchEnd)
	}
	user := common.HexToAddress("0x3b2367e13a5835570e82f81c3f7aff5d32a4470b").Hex()
	userReward := obj.rewards["0x86130bDD69143D8a4E5fc50bf4323D48049E98E4"][user] // pool and user
	if userReward.Cmp(utils.StringToInt("292193677819107573777")) != 0 {
		t.Fatal("wrong rewards calculated for user: ")
	}
}
