package models

import (
	"time"
)

type SsxfDetail struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Channel       string    `json:"channel" xorm:"not null VARCHAR(255)"`
	EventDay      time.Time `json:"event_day" xorm:"DATETIME"`
	SetuserdataPv int       `json:"SetUserData_pv" xorm:"INT(11)"`
	SetuserdataUv int       `json:"SetUserData_uv" xorm:"INT(11)"`
	AdvanceuserPv int       `json:"AdvanceUser_pv" xorm:"INT(11)"`
	AdvanceuserUv int       `json:"AdvanceUser_uv" xorm:"INT(11)"`
	RefectivePv   int       `json:"refective_pv" xorm:"INT(11)"`
	RefectiveUv   int       `json:"refective_uv" xorm:"INT(11)"`
	LastUpdate    time.Time `json:"last_update" xorm:"DATETIME"`
}

func (t SsxfDetail) TableName() string {
	return "ssxf_detail"
}
