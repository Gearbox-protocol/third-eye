package getter

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
)

func TestTxInternalCalls(t *testing.T) {
	rpc := config.GetEnv("GOERLI_ETH_PROVIDER", "")
	if rpc == "" {
		return
	}
	// fetcher := NewInternalCallFetcher(rpc)
	data, err := os.ReadFile("internal_goerli.json")
	if err != nil {
		t.Fatal(err)
	}
	rpcTraces := []RPCTrace{}
	err = json.Unmarshal(data, &rpcTraces)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(utils.ToJson(rpcTraces))

	tenderlyTrace := getTenderlyCall(rpcTraces, "0x63c2fee94a94379de941aab4950c51a505eea652c89b9ad3757d1480092fb330")

	t.Log(utils.ToJson(tenderlyTrace))
}
