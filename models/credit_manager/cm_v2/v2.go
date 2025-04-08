package cm_v2

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CMv2) checkLogV2(txLog types.Log) {
	//-- for credit manager stats
	switch txLog.Topics[0] {
	case core.Topic("AddCollateral(address,address,uint256)"):
		addCollateralEvent, err := mdl.facadeContractv2.ParseAddCollateral(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack AddCollateral event", err)
		}

		mdl.onAddCollateralV2(&txLog, addCollateralEvent.OnBehalfOf.Hex(),
			addCollateralEvent.Token.Hex(),
			addCollateralEvent.Value)
	case core.Topic("OpenCreditAccount(address,address,uint256,uint16)"):
		openCreditAccountEvent, err := mdl.facadeContractv2.ParseOpenCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack OpenCreditAccount event", err)
		}
		mdl.onOpenCreditAccountV2(&txLog,
			openCreditAccountEvent.OnBehalfOf.Hex(),
			openCreditAccountEvent.CreditAccount.Hex(),
			openCreditAccountEvent.BorrowAmount,
			openCreditAccountEvent.ReferralCode)
	case core.Topic("CloseCreditAccount(address,address)"):
		closeCreditAccountEvent, err := mdl.facadeContractv2.ParseCloseCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack CloseCreditAccount event", err)
		}

		mdl.onCloseCreditAccountV2(&txLog,
			closeCreditAccountEvent.Borrower.Hex(),
			closeCreditAccountEvent.To.Hex())
	// for getting correct liquidation status
	case core.Topic("Paused(address)"):
		if txLog.Address.Hex() == mdl.Address { // set pause on cm, if Paused event is emitted only on cm address
			mdl.State.Paused = true
		}
	case core.Topic("Unpaused(address)"):
		if txLog.Address.Hex() == mdl.Address { // unset pause on cm, if Unpaused event is emitted only on cm address
			mdl.State.Paused = false
		}
	case core.Topic("LiquidateCreditAccount(address,address,address,uint256)"):
		liquidateCreditAccountEvent, err := mdl.facadeContractv2.ParseLiquidateCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack LiquidateCreditAccount event", err)
		}

		mdl.onLiquidateCreditAccountV2(&txLog,
			liquidateCreditAccountEvent.Borrower.Hex(),
			liquidateCreditAccountEvent.Liquidator.Hex(),
			liquidateCreditAccountEvent.RemainingFunds)
	case core.Topic("TokenEnabled(address,address)"):
		mdl.enableOrDisableToken(txLog, "TokenEnabled(address,address)")
	case core.Topic("TokenDisabled(address,address)"):
		mdl.enableOrDisableToken(txLog, "TokenDisabled(address,address)")
	case core.Topic("MultiCallStarted(address)"):
		borrower := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		sessionId := mdl.GetCreditOwnerSession(borrower)
		mdl.MulticallMgr.Start(txLog.TxHash.Hex(), &schemas.AccountOperation{
			TxHash:      txLog.TxHash.Hex(),
			BlockNumber: int64(txLog.BlockNumber),
			SessionId:   sessionId,
			Borrower:    borrower,
			Dapp:        txLog.Address.Hex(),
			LogId:       txLog.Index,
			Action:      "MultiCallStarted(address)",
		})
	case core.Topic("MultiCallFinished()"):
		mdl.MulticallMgr.End(txLog.Index, nil, mdl.GetUnderlyingToken())
	case core.Topic("IncreaseBorrowedAmount(address,uint256)"):
		increaseBorrowEvent, err := mdl.facadeContractv2.ParseIncreaseBorrowedAmount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack IncreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV2(&txLog, increaseBorrowEvent.Borrower.Hex(),
			increaseBorrowEvent.Amount, "IncreaseBorrowedAmount")
	case core.Topic("DecreaseBorrowedAmount(address,uint256)"):
		decreaseBorrowEvent, err := mdl.facadeContractv2.ParseDecreaseBorrowedAmount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack DecreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV2(&txLog, decreaseBorrowEvent.Borrower.Hex(),
			new(big.Int).Neg(decreaseBorrowEvent.Amount), "DecreaseBorrowedAmount")
	case core.Topic("TransferAccount(address,address)"):
		if len(txLog.Data) == 0 { // oldowner and newowner are indexed
			transferAccount, err := mdl.facadeContractv2.ParseTransferAccount(txLog)
			if err != nil {
				log.Fatal("[CreditManagerModel]: Cant unpack TransferAccount event", err)
			}
			mdl.onTransferAccountV2(&txLog, transferAccount.OldOwner.Hex(), transferAccount.NewOwner.Hex())
		} else {
			oldOwner := common.BytesToAddress(txLog.Data[:32])
			newOwner := common.BytesToAddress(txLog.Data[32:])
			mdl.onTransferAccountV2(&txLog, oldOwner.Hex(), newOwner.Hex())
		}
	case core.Topic("TransferAccountAllowed(address,address,bool)"):
		transferAccount, err := mdl.facadeContractv2.ParseTransferAccountAllowed(txLog)
		log.CheckFatal(err)
		mdl.Repo.TransferAccountAllowed(&schemas.TransferAccountAllowed{
			From:        transferAccount.From.Hex(),
			To:          transferAccount.To.Hex(),
			Allowed:     transferAccount.State,
			LogId:       int64(txLog.Index),
			BlockNumber: int64(txLog.BlockNumber),
		})
	// on credit manager
	case core.Topic("ExecuteOrder(address,address)"):
		execute, err := mdl.cmContractv2.ParseExecuteOrder(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack ExecuteOrder event", err)
		}
		mdl.AddExecuteParamsV2(&txLog, execute.Borrower, execute.Target)
	case core.Topic("NewConfigurator(address)"):
		newConfigurator := utils.ChecksumAddr(txLog.Topics[1].Hex())
		oldConfigurator := mdl.GetDetailsByKey("configurator")
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Args:        &core.Json{"oldConfigurator": oldConfigurator, "configurator": newConfigurator},
			Type:        schemas.NewConfigurator,
		})
		if oldConfigurator != newConfigurator {
			mdl.Repo.GetAdapter(oldConfigurator).SetBlockToDisableOn(int64(txLog.BlockNumber))
			mdl.addCreditConfiguratorAdapter(newConfigurator)
			mdl.setConfiguratorSyncer(newConfigurator)
		}
	}
}
