package models

import (
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"time"
)

type FeirarDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventDay   time.Time `json:"event_day" xorm:"not null comment('事件日期') DATETIME"`
	Channel    string    `json:"channel" xorm:"VARCHAR(64)"`
	EventKey   string    `json:"event_key" xorm:"comment('时间名称') VARCHAR(128)"`
	Pv         int       `json:"pv" xorm:"comment('PV用户数') INT(11)"`
	Uv         int       `json:"uv" xorm:"comment('UV用户数') INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`
	Detail     string    `json:"detail" xorm:"TEXT"`
}

func (t FeirarDetail) TableName() string {
	return "feirar_detail"
}
type FeirarDetailWeb struct {
	Id         int       `json:"id" `
	EventDay   time.Time `json:"event_day" `
	Channel    string    `json:"channel" `
	EventKey   string    `json:"event_key"`
	Pv         int       `json:"pv" `
	Uv         int       `json:"uv" `
	LastUpdate time.Time `json:"last_update"`
	Detail     string    `json:"detail"`
}


func (t FeirarDetail) Cols() []map[string]string {
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

	col_event_key := map[string]string{
		"name": "api名称",
		"key":  "event_key",
	}
	cols = append(cols, col_event_key)


	cols = append(cols, col_active_mode)
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

func (t FeirarDetail) GetItemsByPage(client *xorm.Engine, pageCount,pageID int) ([]*FeirarDetail, error) {
	var items []*FeirarDetail
	err := client.Desc("event_day").Limit(pageCount,pageCount*(pageID-1)).Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}


