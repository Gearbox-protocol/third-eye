package credit_manager
import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
)

func (cm *CreditManager) addCreditConfigurator() {
	creditConfigurator, err := cm.contractETHV2.CreditConfigurator(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	cm.Repo.AddCreditManagerToFilter(cm.Address, creditConfigurator.Hex())
	cf := credit_filter.NewCreditFilter(creditConfigurator.Hex(), core.CreditConfigurator, cm.Address, cm.DiscoveredAt, cm.Client, cm.Repo)
	cm.Repo.AddSyncAdapter(cf)
}

func (mdl *CreditManager) checkLogV2(txLog types.Log) {
	//-- for credit manager stats
	switch txLog.Topics[0] {
	case core.Topic("OpenCreditAccount(address,address,uint256,uint256)"):
		openCreditAccountEvent, err := mdl.facadeContractV2.ParseOpenCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack OpenCreditAccount event", err)
		}
		mdl.onOpenCreditAccountV2(&txLog,
			openCreditAccountEvent.OnBehalfOf.Hex(),
			openCreditAccountEvent.CreditAccount.Hex(),
			openCreditAccountEvent.BorrowAmount,
			openCreditAccountEvent.ReferralCode)
	case core.Topic("CloseCreditAccount(address,address,uint256)"):
		closeCreditAccountEvent, err := mdl.facadeContractV2.ParseCloseCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack CloseCreditAccount event", err)
		}

		mdl.onCloseCreditAccountV2(&txLog,
			closeCreditAccountEvent.Owner.Hex(),
			closeCreditAccountEvent.To.Hex())
	}
}