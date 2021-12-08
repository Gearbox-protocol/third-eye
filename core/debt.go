package core

import ()

type Debt struct {
	Id                              int64  `gorm:"primaryKey;column:id;autoincrement:true"`
	BlockNumber                     int64  `gorm:"column:block_num"`
	SessionId                       string `gorm:"column:session_id"`
	HealthFactor                    int64  `gorm:"column:health_factor"`
	ThresholdValueBI                string `gorm:"column:threshold_value"`
	CalThresholdValueBI             string `gorm:"column:cal_threshold_value"`
	BorrowedAmountPlusInterestBI    string `gorm:"column:borrowed_amt_with_interest"`
	CalBorrowedAmountPlusInterestBI string `gorm:"column:cal_borrowed_amt_with_interest"`
	CalHealthFactor                 int64  `gorm:"column:cal_health_factor"`
}
