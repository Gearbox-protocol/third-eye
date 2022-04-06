package utils

import (
	"context"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"time"
)

func TimeToDateEndTs(t time.Time) int64 {
	return TimeToDateEndTime(t).Unix()
}

func TimeToDateEndTime(t time.Time) time.Time {
	year, month, day := t.UTC().Date()
	return time.Date(year, month, day, 23, 59, 59, 0, time.UTC)
}

func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func GetTimeoutCtx(sec int) (context.Context, context.CancelFunc) {
	//https://blog.golang.org/context
	timeout, err := time.ParseDuration(fmt.Sprintf("%ds", sec))
	if err != nil {
		log.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.TODO(), timeout*time.Second)
	return ctx, cancel
}

func GetTimeoutOpts(blockNum int64) (*bind.CallOpts, context.CancelFunc) {
	ctx, cancel := GetTimeoutCtx(20)
	return &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
		Context:     ctx,
	}, cancel
}

func Elapsed(what string) func() {
    start := time.Now()
    return func() {
        log.InfoStack3(fmt.Sprintf("%s took %v", what, time.Since(start)))
    }
}