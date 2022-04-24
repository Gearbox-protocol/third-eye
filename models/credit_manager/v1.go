package credit_manager

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (mdl *CreditManager) CommonInit(version int16) {
	// do state changes
	// create underlying token
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(mdl.DiscoveredAt),
	}
	var underlyingToken common.Address
	var err error
	cmContract, err := creditManager.NewCreditManager(common.HexToAddress(mdl.Address), mdl.Client)
	log.CheckFatal(err)

	switch version {
	case 1:
		underlyingToken, err = cmContract.UnderlyingToken(opts)
		log.CheckFatal(err)
	case 2:
		contract, err := creditManagerv2.NewCreditManagerv2(common.HexToAddress(mdl.Address), mdl.Client)
		log.CheckFatal(err)
		underlyingToken, err = contract.Underlying(opts)
		log.CheckFatal(err)
	}
	mdl.Repo.AddToken(underlyingToken.Hex())
	//
	poolAddr, err := cmContract.PoolService(opts)
	if err != nil {
		log.Fatal(err)
	}
	mdl.SetUnderlyingState(&schemas.CreditManagerState{
		Address:         mdl.Address,
		PoolAddress:     poolAddr.Hex(),
		UnderlyingToken: underlyingToken.Hex(),
		Sessions:        map[string]string{},
	})
}

func (cm *CreditManager) addCreditFilter(blockNum int64) {
	creditFilter, err := cm.contractETHV1.CreditFilter(&bind.CallOpts{BlockNumber: big.NewInt(blockNum)})
	if err != nil {
		log.Fatal(err)
	}
	cm.Repo.AddCreditManagerToFilter(cm.Address, creditFilter.Hex())
	cf := credit_filter.NewCreditFilter(creditFilter.Hex(), ds.CreditFilter, cm.Address, cm.DiscoveredAt, cm.Client, cm.Repo)
	cm.Repo.AddSyncAdapter(cf)
}

func (mdl *CreditManager) checkLogV1(txLog types.Log) {
	//-- for credit manager stats
	switch txLog.Topics[0] {
	case core.Topic("OpenCreditAccount(address,address,address,uint256,uint256,uint256)"):
		openCreditAccountEvent, err := mdl.contractETHV1.ParseOpenCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack OpenCreditAccount event", err)
		}

		mdl.onOpenCreditAccount(&txLog,
			openCreditAccountEvent.OnBehalfOf.Hex(),
			openCreditAccountEvent.CreditAccount.Hex(),
			openCreditAccountEvent.Amount,
			openCreditAccountEvent.BorrowAmount,
			openCreditAccountEvent.ReferralCode)
	case core.Topic("CloseCreditAccount(address,address,uint256)"):
		closeCreditAccountEvent, err := mdl.contractETHV1.ParseCloseCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack CloseCreditAccount event", err)
		}

		mdl.onCloseCreditAccount(&txLog,
			closeCreditAccountEvent.Owner.Hex(),
			closeCreditAccountEvent.To.Hex(),
			closeCreditAccountEvent.RemainingFunds)
	case core.Topic("LiquidateCreditAccount(address,address,uint256)"):
		liquidateCreditAccountEvent, err := mdl.contractETHV1.ParseLiquidateCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack LiquidateCreditAccount event", err)
		}

		mdl.onLiquidateCreditAccount(&txLog,
			liquidateCreditAccountEvent.Owner.Hex(),
			liquidateCreditAccountEvent.Liquidator.Hex(),
			liquidateCreditAccountEvent.RemainingFunds)
	case core.Topic("RepayCreditAccount(address,address)"):
		repayCreditAccountEvent, err := mdl.contractETHV1.ParseRepayCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack RepayCreditAccount event", err)
		}

		mdl.onRepayCreditAccount(&txLog,
			repayCreditAccountEvent.Owner.Hex(),
			repayCreditAccountEvent.To.Hex())
	case core.Topic("IncreaseBorrowedAmount(address,uint256)"):
		increaseBorrowEvent, err := mdl.contractETHV1.ParseIncreaseBorrowedAmount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack IncreaseBorrowedAmount event", err)
		}

		mdl.onIncreaseBorrowedAmount(&txLog,
			increaseBorrowEvent.Borrower.Hex(), increaseBorrowEvent.Amount)
	case core.Topic("AddCollateral(address,address,uint256)"):
		addCollateralEvent, err := mdl.contractETHV1.ParseAddCollateral(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack AddCollateral event", err)
		}

		mdl.onAddCollateral(&txLog, addCollateralEvent.OnBehalfOf.Hex(),
			addCollateralEvent.Token.Hex(),
			addCollateralEvent.Value)
	case core.Topic("ExecuteOrder(address,address)"):
		execute, err := mdl.contractETHV1.ParseExecuteOrder(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack ExecuteOrder event", err)
		}
		mdl.AddExecuteParams(&txLog, execute.Borrower, execute.Target)
	case core.Topic("NewParameters(uint256,uint256,uint256,uint256,uint256,uint256)"):
		params, err := mdl.contractETHV1.ParseNewParameters(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack NewParameters event", err)
		}
		mdl.State.MinAmount = (*core.BigInt)(params.MinAmount)
		mdl.State.MaxAmount = (*core.BigInt)(params.MaxAmount)
		mdl.State.MaxLeverageFactor = params.MaxLeverage.Int64()
		mdl.State.FeeInterest = params.FeeInterest.Int64()
		mdl.Repo.AddParameters(txLog.Index, txLog.TxHash.Hex(), &schemas.Parameters{
			BlockNum:            int64(txLog.BlockNumber),
			CreditManager:       mdl.GetAddress(),
			MinAmount:           (*core.BigInt)(params.MinAmount),
			MaxAmount:           (*core.BigInt)(params.MaxAmount),
			MaxLeverage:         (*core.BigInt)(params.MaxLeverage),
			FeeInterest:         (*core.BigInt)(params.FeeInterest),
			FeeLiquidation:      (*core.BigInt)(params.FeeLiquidation),
			LiquidationDiscount: (*core.BigInt)(params.LiquidationDiscount),
		}, mdl.State.UnderlyingToken)
	case core.Topic("TransferAccount(address,address)"):
		if len(txLog.Data) == 0 { // oldowner and newowner are indexed
			transferAccount, err := mdl.contractETHV1.ParseTransferAccount(txLog)
			if err != nil {
				log.Fatal("[CreditManagerModel]: Cant unpack TransferAccount event", err)
			}
			mdl.onTransferAccount(&txLog, transferAccount.OldOwner.Hex(), transferAccount.NewOwner.Hex())
		} else {
			oldOwner := common.BytesToAddress(txLog.Data[:32])
			newOwner := common.BytesToAddress(txLog.Data[32:])
			mdl.onTransferAccount(&txLog, oldOwner.Hex(), newOwner.Hex())
		}
	}
}
