package models

import (
	"time"
)

type NewsDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventDay   time.Time `json:"event_day" xorm:"not null comment('事件日期') DATETIME"`
	Channel    string    `json:"channel" xorm:"VARCHAR(128)"`
	EventType  string    `json:"event_type" xorm:"VARCHAR(128)"`
	Pv         int       `json:"pv" xorm:"comment('PV用户数') INT(11)"`
	Uv         int       `json:"uv" xorm:"comment('UV用户数') INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`
	Detail     string    `json:"detail" xorm:"TEXT"`
}

type NewsDetailWeb struct {
	Id         int       `json:"id" `
	EventDay   string `json:"event_day" `
	Channel    string    `json:"channel" `
	EventType  string    `json:"event_type" `
	Pv         string       `json:"pv" `
	Uv         string       `json:"uv" `
	LastUpdate string `json:"last_update" `
	Detail     string    `json:"detail"`
}

func (t NewsDetail) TableName() string {
	return "news_detail"
}


//Cols ...用户web显示使用
func (t NewsDetail) Cols() []map[string]string {
	var cols []map[string]string
	col_event_day := map[string]string{
		"name": "事件日期",
		"key":  "event_day",
	}
	cols = append(cols, col_event_day)
	col_active_mode := map[string]string{
		"name": "渠道名称",
		"key":  "channel",
	}
	cols = append(cols, col_active_mode)

	col_event_type := map[string]string{
		"name": "消息类型",
		"key":  "event_type",
	}
	cols = append(cols, col_event_type)

	col_pv := map[string]string{
		"name": "pv",
		"key":  "pv",
	}
	cols = append(cols, col_pv)
	col_uv := map[string]string{
		"name": "uv",
		"key":  "uv",
	}
	cols = append(cols, col_uv)
	col_last_update := map[string]string{
		"name": "更新时间",
		"key":  "last_update",
	}
	cols = append(cols, col_last_update)
	return cols
}