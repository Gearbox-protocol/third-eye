package yearn_price_feed

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// only used in v1,v2

func (mdl *YearnPriceFeed) CalculateYearnPFInternally(blockNum int64) (*schemas.PriceFeed, error) {
	return mdl.yearnPFInternal.calculatePrice(blockNum, mdl.Client, mdl.GetVersion())
}

type yearnPFInternal struct {
	mainPFAddress        common.Address // for yearn manual price calculation
	yVaultAddr           common.Address //for yearn manual price calculation
	underlyingPFContract *priceFeed.PriceFeed
	decimalDivider       *big.Int
	version              core.VersionType
}

// only used in v1 and v2 for calculating price if latestRoudnData execution is Reverted
func (mdl *yearnPFInternal) calculatePrice(blockNum int64, client core.ClientI, version core.VersionType) (*schemas.PriceFeed, error) {
	if mdl.underlyingPFContract == nil {
		if err := mdl.setContracts(blockNum, client); err != nil {
			return nil, err
		}
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	//
	pps, err := core.CallFuncWithExtraBytes(client, "99530b06", mdl.yVaultAddr, blockNum, nil) // pps
	if err != nil {
		return nil, err
	}
	pricePerShare := new(big.Int).SetBytes(pps)
	//
	roundData, err := mdl.underlyingPFContract.LatestRoundData(opts)
	if err != nil {
		return nil, err
	}

	// for yearn it is based on the vault. https://github.com/Gearbox-protocol/integrations-v2/blob/main/contracts/oracles/yearn/YearnPriceFeed.sol#L62
	var newAnswer *big.Int
	if pricePerShare == nil || roundData.Answer == nil || mdl.decimalDivider == nil {
		newAnswer = new(big.Int)
		log.Errorf("failing to get price internally", mdl.mainPFAddress)
	} else {

		newAnswer = new(big.Int).Quo(
			new(big.Int).Mul(pricePerShare, roundData.Answer),
			mdl.decimalDivider,
		)
	}
	pfVersion := schemas.VersionToPFVersion(version, false)
	return &schemas.PriceFeed{
		RoundId:         roundData.RoundId.Int64(),
		PriceBI:         (*core.BigInt)(newAnswer),
		Price:           utils.GetFloat64Decimal(newAnswer, pfVersion.Decimals()),
		MergedPFVersion: schemas.MergedPFVersion(pfVersion), // only used for v1,v2 so can convert from pfVersion to MergedPFVersion
	}, nil
}

func (mdl *yearnPFInternal) setContracts(blockNum int64, client core.ClientI) error {
	// set the price feed contract
	underlyingPFAddrBytes, err := core.CallFuncWithExtraBytes(client, "741bef1a", mdl.mainPFAddress, blockNum, nil) // priceFeed
	if err != nil {
		return err
	}
	// underlying price feed not found
	if common.BytesToAddress(underlyingPFAddrBytes) == core.NULL_ADDR {
		return fmt.Errorf("address for underlying pf for yearn feed(%d) not found at %d",
			mdl.mainPFAddress, blockNum)
	}
	mdl.underlyingPFContract, err = priceFeed.NewPriceFeed(common.BytesToAddress(underlyingPFAddrBytes), client)
	log.CheckFatal(err)

	// set the yvault contract
	if mdl.version == core.NewVersion(300) {
		// https://github.com/Gearbox-protocol/oracles-v3/blob/2ac6d1ba1108df949222084791699d821096bc8c/contracts/oracles/yearn/YearnPriceFeed.sol#L19
		// https://github.com/Gearbox-protocol/oracles-v3/blob/2ac6d1ba1108df949222084791699d821096bc8c/contracts/oracles/SingleAssetLPPriceFeed.sol#L24
		//https://github.com/Gearbox-protocol/oracles-v3/blob/2ac6d1ba1108df949222084791699d821096bc8c/contracts/oracles/LPPriceFeed.sol#L69C9-
		// LPCONTRACT_LOGIC
		lpCOntractBytes, err := core.CallFuncWithExtraBytes(client, "8acee3cf", mdl.mainPFAddress, blockNum, nil) // lpContract
		if err != nil {
			return err
		}
		mdl.yVaultAddr = common.BytesToAddress(lpCOntractBytes)
	} else {
		yVaultAddrBytes, err := core.CallFuncWithExtraBytes(client, "33303f8e", mdl.mainPFAddress, blockNum, nil) // yVault
		if err != nil {
			return err
		}
		mdl.yVaultAddr = common.BytesToAddress(yVaultAddrBytes)
	}
	//

	// set the decimals
	decimalsBytes, err := core.CallFuncWithExtraBytes(client, "313ce567", mdl.yVaultAddr, blockNum, nil) // decimals
	if err != nil {
		return err
	}
	mdl.decimalDivider = utils.GetExpInt(int8(
		new(big.Int).SetBytes(decimalsBytes).Int64(),
	))
	//
	return nil
}
