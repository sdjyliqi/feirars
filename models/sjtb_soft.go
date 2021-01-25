package models

import (
	"time"
)

type SjtbSoft struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SoftName      string    `json:"soft_name" xorm:"VARCHAR(255)"`
	Channel       string    `json:"channel" xorm:"not null VARCHAR(255)"`
	EventDay      time.Time `json:"event_day" xorm:"DATETIME"`
	ApplistshowPv int       `json:"applistshow_pv" xorm:"INT(11)"`
	ApplistshowUv int       `json:"applistshow_uv" xorm:"INT(11)"`
	ApplistokPv   int       `json:"applistok_pv" xorm:"INT(11)"`
	ApplistokUv   int       `json:"applistok_uv" xorm:"INT(11)"`
	Apprun1Pv     int       `json:"apprun1_pv" xorm:"INT(11)"`
	Apprun1Uv     int       `json:"apprun1_uv" xorm:"INT(11)"`
	LastUpdate    time.Time `json:"last_update" xorm:"DATETIME"`
}

func (t SjtbSoft) TableName() string {
	return "sjtb_soft"
}
