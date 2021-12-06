package core

import ()

type Debt struct {
	Id                           int64  `gorm:"primaryKey;column:id;autoincrement:true"`
	BlockNumber                  int64  `gorm:"column:block_num"`
	SessionId                    string `gorm:"column:session_id"`
	HealthFactor                 int64  `gorm:"column:health_factor"`
	TotalValueBI                 string `gorm:"column:total_value"`
	CalculatedTotalValueBI       string `gorm:"column:cal_total_value"`
	BorrowedAmountPlusInterestBI string `gorm:"column:borrowed_amt_with_interest"`
}
