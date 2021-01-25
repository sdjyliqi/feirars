package models

import (
	"time"
)

type Md5Check struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SoftName       string    `json:"soft_name" xorm:"VARCHAR(255)"`
	Channel        string    `json:"channel" xorm:"not null VARCHAR(255)"`
	EventDay       time.Time `json:"event_day" xorm:"DATETIME"`
	ApplistokPv    int       `json:"applistok_pv" xorm:"INT(11)"`
	ApplistokUv    int       `json:"applistok_uv" xorm:"INT(11)"`
	ApplistshowPv  int       `json:"applistshow_pv" xorm:"INT(11)"`
	ApplistshowUv  int       `json:"applistshow_uv" xorm:"INT(11)"`
	ApplistcloseUv int       `json:"applistclose_uv" xorm:"INT(11)"`
	ApplistclosePv int       `json:"applistclose_pv" xorm:"INT(11)"`
	Appdown1Uv     int       `json:"appdown1_uv" xorm:"INT(11)"`
	Appdown1Pv     int       `json:"appdown1_pv" xorm:"INT(11)"`
	Appdown0Pv     int       `json:"appdown0_pv" xorm:"INT(11)"`
	Appdown0Uv     int       `json:"appdown0_uv" xorm:"INT(11)"`
	Apprun0Pv      int       `json:"apprun0_pv" xorm:"INT(11)"`
	Apprun0Uv      int       `json:"apprun0_uv" xorm:"INT(11)"`
	Apprun1Pv      int       `json:"apprun1_pv" xorm:"INT(11)"`
	Apprun1Uv      int       `json:"apprun1_uv" xorm:"INT(11)"`
	LastUpdate     time.Time `json:"last_update" xorm:"DATETIME"`
}

func (t Md5Check) TableName() string {
	return "md5_check"
}
