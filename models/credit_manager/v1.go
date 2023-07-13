package credit_manager

import (
	"math/big"

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
)

// sets underlying state on init
// pool, and underlying token address
func (mdl *CreditManager) CommonInit(version core.VersionType) {
	// do state changes
	// create underlying token
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(mdl.DiscoveredAt),
	}
	var underlyingToken common.Address
	var err error
	cmContract, err := creditManager.NewCreditManager(common.HexToAddress(mdl.Address), mdl.Client)
	log.CheckFatal(err)
	if version.IsGBv1() {
		underlyingToken, err = cmContract.UnderlyingToken(opts)
		log.CheckFatal(err)
	} else {
		contract, err := creditManagerv2.NewCreditManagerv2(common.HexToAddress(mdl.Address), mdl.Client)
		log.CheckFatal(err)
		underlyingToken, err = contract.Underlying(opts)
		log.CheckFatal(err)
	}
	mdl.Repo.GetToken(underlyingToken.Hex())
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
		Version:         version,
	})
}

func (cm *CreditManager) addCreditFilterAdapter(blockNum int64) {
	creditFilter, err := cm.contractETHV1.CreditFilter(&bind.CallOpts{BlockNumber: big.NewInt(blockNum)})
	if err != nil {
		log.Fatal(err)
	}
	cm.Repo.GetDCWrapper().AddCreditManagerToFilter(cm.Address, creditFilter.Hex())
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
		paramsEvent, err := mdl.contractETHV1.ParseNewParameters(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack NewParameters event", err)
		}
		mdl.State.MinAmount = (*core.BigInt)(paramsEvent.MinAmount)
		mdl.State.MaxAmount = (*core.BigInt)(paramsEvent.MaxAmount)
		mdl.State.MaxLeverageFactor = paramsEvent.MaxLeverage.Int64()
		params := &schemas.Parameters{
			BlockNum:            int64(txLog.BlockNumber),
			CreditManager:       mdl.GetAddress(),
			MinAmount:           (*core.BigInt)(paramsEvent.MinAmount),
			MaxAmount:           (*core.BigInt)(paramsEvent.MaxAmount),
			MaxLeverage:         (*core.BigInt)(paramsEvent.MaxLeverage),
			FeeInterest:         uint16(paramsEvent.FeeInterest.Int64()),
			FeeLiquidation:      uint16(paramsEvent.FeeLiquidation.Int64()),
			LiquidationDiscount: uint16(paramsEvent.LiquidationDiscount.Int64()),
		}
		mdl.Repo.AddParameters(txLog.Index, txLog.TxHash.Hex(), params, mdl.State.UnderlyingToken)
		mdl.setParams(params)
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
