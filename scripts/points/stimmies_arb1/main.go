package main

import (
	"fmt"
	"sort"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/repository"
)

type RewardAndTs struct {
	User   string  `gorm:"column:user_address" json:"user"`
	Points float64 `gorm:"column:point" json:"amount"`
}

type UserPo struct {
	User   string  `gorm:"column:user_address" json:"user"`
	Points float64 `gorm:"column:points" json:"amount"`
}
func main() {
	cfg := config.NewConfig()
	db := repository.NewDBClient(cfg)
	cam := "Arb-Campaign4"
	data := []RewardAndTs{}
	err := db.Raw(`(select user_address, point from user_points where campaign=?) union all (select user_address, points from hrlyrate_points where campaign=?)`, cam, cam).Find(&data).Error
	log.CheckFatal(err)
	ans := map[string]float64{}
	for _, d := range data {
		ans[d.User] += d.Points
	}
	l := []UserPo{}
	for k, v := range ans {
		l = append(l, UserPo{User: k, Points: v})
	}
	sort.Slice(l, func(i, j int) bool { return l[i].Points > l[j].Points })
	fmt.Println("user, points")
	for _, v := range l {
		fmt.Printf("%s, %f\n", v.User, v.Points)
	}
}
