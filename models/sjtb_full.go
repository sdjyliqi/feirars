package models

import (
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"time"
)

type SjtbFull struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
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

func (t SjtbFull) TableName() string {
	return "sjtb_full"
}

type SjtbFullWeb struct {
	Id             string `json:"id"`
	EventDay       string `json:"event_day"`
	Channel        string `json:"channel" `
	ApplistokPv    int    `json:"applistok_pv" `
	ApplistokUv    int    `json:"applistok_uv" `
	ApplistshowPv  int    `json:"applistshow_pv"`
	ApplistshowUv  int    `json:"applistshow_uv"`
	ApplistcloseUv int    `json:"applistclose_uv"`
	ApplistclosePv int    `json:"applistclose_pv"`
	Appdown1Uv     int    `json:"appdown1_uv"`
	Appdown1Pv     int    `json:"appdown1_pv"`
	Appdown0Pv     int    `json:"appdown0_pv" `
	Appdown0Uv     int    `json:"appdown0_uv" `
	Apprun0Pv      int    `json:"apprun0_pv" `
	Apprun0Uv      int    `json:"apprun0_uv"`
	Apprun1Pv      int    `json:"apprun1_pv"`
	Apprun1Uv      int    `json:"apprun1_uv"`
	LastUpdate     string `json:"last_update" `
}

func (t SjtbFull) CovertWebItem(item *SjtbFull) SjtbFullWeb {
	webItem := SjtbFullWeb{
		EventDay:       item.EventDay.Format(utils.DayTime),
		Channel:        item.Channel,
		ApplistokPv:    item.ApplistokPv,
		ApplistokUv:    item.ApplistokUv,
		ApplistshowPv:  item.ApplistshowPv,
		ApplistshowUv:  item.ApplistshowUv,
		ApplistcloseUv: item.ApplistcloseUv,
		ApplistclosePv: item.ApplistclosePv,
		Appdown1Uv:     item.Appdown1Uv,
		Appdown1Pv:     item.Appdown1Pv,
		Appdown0Pv:     item.Appdown0Pv,
		Appdown0Uv:     item.Appdown0Uv,
		Apprun0Pv:      item.Apprun0Pv,
		Apprun0Uv:      item.Apprun0Uv,
		Apprun1Pv:      item.Apprun1Pv,
		Apprun1Uv:      item.Apprun1Uv,
		LastUpdate:     item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols 获取所有的列名
func (t SjtbFull) Cols() []map[string]string {
	var cols []map[string]string
	col := map[string]string{
		"name":  "日期",
		"key":   "event_day",
		"click": "0",
		"raw":   "EventDay",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "渠道",
		"key":   "channel",
		"click": "1",
		"raw":   "Channel",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ApplistokPV",
		"key":   "applistok_pv",
		"click": "0",
		"raw":   "ApplistokPv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ApplistokUV",
		"key":   "applistok_uv",
		"click": "0",
		"raw":   "ApplistokUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ApplistshowPV",
		"key":   "applistshow_pv",
		"click": "0",
		"raw":   "ApplistshowPv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ApplistshowUV",
		"key":   "applistshow_uv",
		"click": "0",
		"raw":   "ApplistshowUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ApplistclosePv",
		"key":   "applistclose_pv",
		"click": "0",
		"raw":   "ApplistclosePv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ApplistcloseUv",
		"key":   "applistclose_uv",
		"click": "0",
		"raw":   "ApplistcloseUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Appdown1Uv",
		"key":   "appdown1_uv",
		"click": "0",
		"raw":   "Appdown1Uv",
	}

	cols = append(cols, col)
	col = map[string]string{
		"name":  "Appdown1Pv",
		"key":   "appdown1_pv",
		"click": "0",
		"raw":   "Appdown1Pv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Appdown0Pv",
		"key":   "appdown0_pv",
		"click": "0",
		"raw":   "Appdown0Pv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Appdown0Uv",
		"key":   "appdown0_uv",
		"click": "0",
		"raw":   "Appdown0Uv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Apprun0Pv",
		"key":   "apprun0_pv",
		"click": "0",
		"raw":   "Apprun0Pv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Apprun0Uv",
		"key":   "apprun0_uv",
		"click": "0",
		"raw":   "Apprun0Uv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Apprun1Pv",
		"key":   "apprun1_pv",
		"click": "0",
		"raw":   "Apprun1Pv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Apprun1Uv",
		"key":   "apprun1_uv",
		"click": "0",
		"raw":   "Apprun1Uv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "更新时间",
		"key":   "last_update",
		"click": "0",
		"raw":   "LastUpdate",
	}
	cols = append(cols, col)
	return cols
}

//GetAllChannels ...获取所有渠道
func (t SjtbFull) GetAllChannels(client *xorm.Engine) ([]string, error) {
	var items []*SjtbFull
	var channels []string
	err := client.Distinct("channel").OrderBy("channel").Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the channel  from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	for _, v := range items {
		channels = append(channels, v.Channel)
	}
	return channels, nil
}

func (t SjtbFull) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*SjtbFull, int64, error) {
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*SjtbFull
	item := &SjtbFull{}
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	session = session.Desc("event_day")
	if pageCount > 0 {
		session = session.Limit(pageCount, pageCount*(pageID-1))
	}

	err := session.Find(&items)
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
