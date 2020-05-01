package models

import (
	"github.com/go-xorm/xorm"
	"time"
)

type ActiveDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventDay   time.Time `json:"event_day" xorm:"not null comment('事件日期') DATETIME"`
	ActiveMode string    `json:"active_mode" xorm:"VARCHAR(64)"`
	Pv         int       `json:"pv" xorm:"comment('PV用户数') INT(11)"`
	Uv         int       `json:"uv" xorm:"comment('UV用户数') INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`
	Detail     string    `json:"detail" xorm:"TEXT"`
}

type ActiveDetailWeb struct {
	Id         string `json:"id"`
	EventDay   string `json:"event_day"`
	ActiveMode string `json:"active_mode" xorm:"VARCHAR(64)"`
	Pv         string `json:"pv"`
	Uv         string `json:"uv"`
	LastUpdate string `json:"last_update" `
	Detail     string `json:"detail"`
}

func (t ActiveDetail) TableName() string {
	return "active_detail"
}

//Cols ...用户web显示使用
func (t ActiveDetail) Cols() []map[string]string {
	var cols []map[string]string
	col_event_day := map[string]string{
		"name": "事件日期",
		"key":  "event_day",
	}
	cols = append(cols, col_event_day)
	col_active_mode := map[string]string{
		"name": "激活方式",
		"key":  "active_mode",
	}
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

func (t ActiveDetail) GetItemsByPage(client *xorm.Engine, pageCount,pageID int) ([]*ActiveDetail, error) {
	var items []*ActiveDetail
	err := client.Desc("event_day").Limit(pageCount,pageCount*(pageID-1)).Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
