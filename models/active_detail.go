package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"time"
)

type ActiveDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventDay   time.Time `json:"event_day" xorm:"not null comment('事件日期') DATETIME"`
	Channel    string    `json:"channel" xorm:"VARCHAR(64)"`
	ActiveMode string    `json:"active_mode" xorm:"VARCHAR(64)"`
	Pv         int       `json:"pv" xorm:"comment('PV用户数') INT(11)"`
	Uv         int       `json:"uv" xorm:"comment('UV用户数') INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`
	Detail     string    `json:"detail" xorm:"TEXT"`
}

type ActiveDetailWeb struct {
	Id         string `json:"id"`
	EventDay   string `json:"event_day"`
	Channel    string `json:"channel" `
	ActiveMode string `json:"active_mode" xorm:"VARCHAR(64)"`
	Pv         string `json:"pv"`
	Uv         string `json:"uv"`
	LastUpdate string `json:"last_update" `
}

func (t ActiveDetail) TableName() string {
	return "active_detail"
}

func (t ActiveDetail) CovertWebItem(item *ActiveDetail) ActiveDetailWeb {
	webItem := ActiveDetailWeb{
		EventDay:   item.EventDay.Format(utils.DayTime),
		Channel:    item.Channel,
		ActiveMode: item.ActiveMode,
		Pv:         fmt.Sprintf("%d", item.Pv),
		Uv:         fmt.Sprintf("%d", item.Uv),
		LastUpdate: item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols ...用户web显示使用
func (t ActiveDetail) Cols() []map[string]string {
	var cols []map[string]string
	col_event_day := map[string]string{
		"name":  "事件日期",
		"key":   "event_day",
		"click": "0",
	}
	cols = append(cols, col_event_day)

	col_client_channel := map[string]string{
		"name":  "渠道名称",
		"key":   "channel",
		"click": "1",
	}
	cols = append(cols, col_client_channel)

	col_active_mode := map[string]string{
		"name":  "激活方式",
		"key":   "active_mode",
		"click": "0",
	}
	cols = append(cols, col_active_mode)
	col_pv := map[string]string{
		"name":  "pv",
		"key":   "pv",
		"click": "0",
	}
	cols = append(cols, col_pv)
	col_uv := map[string]string{
		"name":  "uv",
		"key":   "uv",
		"click": "0",
	}
	cols = append(cols, col_uv)
	col_last_update := map[string]string{
		"name":  "更新时间",
		"key":   "last_update",
		"click": "0",
	}
	cols = append(cols, col_last_update)
	return cols
}

func (t ActiveDetail) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*ActiveDetail, int64, error) {
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*ActiveDetail
	item := &ActiveDetail{}
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	err := session.Desc("event_day").
		Limit(pageCount, pageCount*(pageID-1)).
		Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, 0, err
	}

	session = client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	cnt, err := session.Count(item)
	if err != nil {
		glog.Errorf("[mysql]Get the count of items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, 0, err
	}
	return items, cnt, nil
}
