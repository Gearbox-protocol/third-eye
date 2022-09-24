package services

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
)

func TestGetMainEventLogs(t *testing.T) {
	ep := NewExecuteParser(&config.Config{ChainId: 42})
	actionWithMulticall := ep.GetMainEventLogs("0xfbbfbca8d6300adc20c1fd9b2bf2074a14cad0873ad5ed8492ef226861f7c0cc", "0x5aacdab79aa2d30f4242898ecdafda2ed2216db2")
	if len(actionWithMulticall) != 1 || actionWithMulticall[0].Name != "openCreditAccountMulticall" || actionWithMulticall[0].Len() != 1 {
		log.Fatal(utils.ToJson(actionWithMulticall))
	}
}
