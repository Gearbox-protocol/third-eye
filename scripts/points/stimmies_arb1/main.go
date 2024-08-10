package main

import (
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/repository"
)

type RewardAndTs struct {
	User   string  `gorm:"column:user_address" json:"user"`
	Points float64 `gorm:"column:point" json:"amount"`
}

func main() {
	cfg := config.NewConfig()
	db := repository.NewDBClient(cfg)
	cam := "Arb-Campaign2"
	data := []RewardAndTs{}
	err := db.Raw(`(select user_address, point from user_points where campaign=?) union all (select user_address, points from hrlyrate_points where campaign=?)`, cam, cam).Find(&data).Error
	log.CheckFatal(err)
	ans := map[string]float64{}
	for _, d := range data {
		ans[d.User] += d.Points
	}
	fmt.Println("user, points")
	for k, v := range ans {
		fmt.Printf("%s, %f\n", k, v)
	}
}
