package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"time"
)

type PreserveDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventTime  time.Time `json:"event_time" xorm:"not null DATETIME"`
	Channel    string    `json:"channel" xorm:"VARCHAR(64)"`
	Uv         int       `json:"uv" xorm:"INT(11)"`
	NewUv      int       `json:"new_uv" xorm:"INT(11)"`
	Day1Active int       `json:"day1_active" xorm:"INT(11)"`
	Day2Active int       `json:"day2_active" xorm:"INT(11)"`
	Day3Active int       `json:"day3_active" xorm:"INT(11)"`
	Day4Active int       `json:"day4_active" xorm:"INT(11)"`
	Day5Active int       `json:"day5_active" xorm:"INT(11)"`
	Day6Active int       `json:"day6_active" xorm:"INT(11)"`
	WeekActive int       `json:"week_active" xorm:"INT(11)"`
	Detail     string    `json:"detail" xorm:"TEXT"`
	LastUpdate time.Time `json:"last_update" xorm:"DATETIME"`
}
type PreserveDetailWeb struct {
	Id         string `json:"id" `
	EventTime  string `json:"event_time" `
	Channel    string `json:"channel" `
	Uv         string `json:"uv" `
	NewUv      string `json:"new_uv" `
	Day1Active string `json:"day1_active"`
	Day2Active string `json:"day2_active" `
	Day3Active string `json:"day3_active" `
	Day4Active string `json:"day4_active"`
	Day5Active string `json:"day5_active" `
	Day6Active string `json:"day6_active"`
	WeekActive string `json:"week_active" `
	Detail     string `json:"detail" `
	LastUpdate string `json:"last_update" `
}

func (t PreserveDetail) TableName() string {
	return "preserve_detail"
}

func (t PreserveDetail) CovertWebItem(item *PreserveDetail) PreserveDetailWeb {
	webItem := PreserveDetailWeb{
		EventTime:  item.EventTime.Format(utils.DayTime),
		Channel:    item.Channel,
		Uv:         fmt.Sprintf("%d", item.Uv),
		NewUv:      fmt.Sprintf("%d", item.NewUv),
		Day1Active: fmt.Sprintf("%d", item.Day1Active),
		Day2Active: fmt.Sprintf("%d", item.Day2Active),
		Day3Active: fmt.Sprintf("%d", item.Day3Active),
		Day4Active: fmt.Sprintf("%d", item.Day4Active),
		Day5Active: fmt.Sprintf("%d", item.Day5Active),
		Day6Active: fmt.Sprintf("%d", item.Day6Active),
		WeekActive: fmt.Sprintf("%d", item.WeekActive),
		LastUpdate: item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols ...用户web显示使用
func (t PreserveDetail) Cols() []map[string]string {
	var cols []map[string]string
	col_event_day := map[string]string{
		"name": "事件日期",
		"key":  "event_time",
	}
	cols = append(cols, col_event_day)

	col_client_channel := map[string]string{
		"name":  "渠道名称",
		"key":   "channel",
		"click": "1",
	}
	cols = append(cols, col_client_channel)

	col_uv := map[string]string{
		"name": "uv",
		"key":  "uv",
	}
	cols = append(cols, col_uv)

	col_new_uv := map[string]string{
		"name": "新增用户",
		"key":  "new_uv",
	}
	cols = append(cols, col_new_uv)
	//次日留存
	col_day1_active := map[string]string{
		"name": "1日留存",
		"key":  "day1_active",
	}
	cols = append(cols, col_day1_active)

	//二日留存
	col_day2_active := map[string]string{
		"name": "2日留存",
		"key":  "day2_active",
	}
	cols = append(cols, col_day2_active)

	//三日留存
	col_day3_active := map[string]string{
		"name": "3日留存",
		"key":  "day3_active",
	}
	cols = append(cols, col_day3_active)

	//四日留存
	col_day4_active := map[string]string{
		"name": "4日留存",
		"key":  "day4_active",
	}
	cols = append(cols, col_day4_active)

	//五日留存
	col_day5_active := map[string]string{
		"name": "5日留存",
		"key":  "day5_active",
	}
	cols = append(cols, col_day5_active)

	//六日留存
	col_day6_active := map[string]string{
		"name": "6日留存",
		"key":  "day6_active",
	}
	cols = append(cols, col_day6_active)

	//周留存
	col_week_active := map[string]string{
		"name": "周留存",
		"key":  "week_active",
	}
	cols = append(cols, col_week_active)

	col_last_update := map[string]string{
		"name": "更新时间",
		"key":  "last_update",
	}
	cols = append(cols, col_last_update)
	return cols
}

func (t PreserveDetail) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*PreserveDetail, int64, error) {
	var items []*PreserveDetail
	item := &PreserveDetail{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	session := client.Where("event_time>=?", timeTS).And("event_time<=?", timeTE)
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}

	err := session.Desc("event_time").
		Limit(pageCount, pageCount*(pageID-1)).
		Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, 0, err
	}
	session = client.Where("event_time>=?", timeTS).And("event_time<=?", timeTE)
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
