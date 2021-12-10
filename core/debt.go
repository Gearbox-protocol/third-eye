package core

import ()

type Debt struct {
	Id                              int64  `gorm:"primaryKey;column:id;autoincrement:true"`
	BlockNumber                     int64  `gorm:"column:block_num"`
	SessionId                       string `gorm:"column:session_id"`
	HealthFactor                    int64  `gorm:"column:health_factor"`
	TotalValueBI                    string `gorm:"column:total_value"`
	BorrowedAmountPlusInterestBI    string `gorm:"column:borrowed_amt_with_interest"`
	CalHealthFactor                 int64  `gorm:"column:cal_health_factor"`
	CalTotalValue                   string `gorm:"column:cal_total_value"`
	CalBorrowedAmountPlusInterestBI string `gorm:"column:cal_borrowed_amt_with_interest"`
	CalThresholdValueBI             string `gorm:"column:cal_threshold_value"`
}
