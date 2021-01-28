package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
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

type SjtbSoftWeb struct {
	Id            string `json:"id"`
	EventDay      string `json:"event_day"`
	Channel       string `json:"channel" `
	ApplistshowPv int    `json:"applistshow_pv" xorm:"INT(11)"`
	ApplistshowUv int    `json:"applistshow_uv" xorm:"INT(11)"`
	ApplistokPv   int    `json:"applistok_pv" xorm:"INT(11)"`
	ApplistokUv   int    `json:"applistok_uv" xorm:"INT(11)"`
	Apprun1Pv     int    `json:"apprun1_pv" xorm:"INT(11)"`
	Apprun1Uv     int    `json:"apprun1_uv" xorm:"INT(11)"`
	LastUpdate    string `json:"last_update" `
}

func (t SjtbSoft) CovertWebItem(item *SjtbSoft) SjtbSoftWeb {
	webItem := SjtbSoftWeb{
		Id:            "",
		EventDay:      item.EventDay.Format(utils.DayTime),
		Channel:       item.Channel,
		ApplistshowPv: item.ApplistshowPv,
		ApplistshowUv: item.ApplistshowUv,
		ApplistokPv:   item.ApplistokPv,
		ApplistokUv:   item.ApplistshowUv,
		Apprun1Pv:     item.Apprun1Pv,
		Apprun1Uv:     item.Apprun1Uv,
		LastUpdate:    item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols 获取所有的列名
func (t SjtbSoft) Cols() []map[string]string {
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
		"name":  "ApplistshowPv",
		"key":   "applistshow_pv",
		"click": "0",
		"raw":   "ApplistshowPv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ApplistshowUv",
		"key":   "applistshow_uv",
		"click": "0",
		"raw":   "ApplistshowUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ApplistokPv",
		"key":   "applistok_pv",
		"click": "0",
		"raw":   "ApplistokPv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ApplistokUv",
		"key":   "applistok_uv",
		"click": "0",
		"raw":   "ApplistokUv",
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
func (t SjtbSoft) GetAllChannels(client *xorm.Engine) ([]string, error) {
	var items []*SjtbSoft
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

func (t SjtbSoft) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*SjtbSoft, int64, error) {
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*SjtbSoft
	item := &SjtbSoft{}
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

func (t SjtbSoft) GetChartItems(client *xorm.Engine, chn string, tsStart, tsEnd int64, eventKey string) (*utils.ChartDetail, error) {
	chartXvalue := make([]string, 0)
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*SjtbSoft
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE).And(fmt.Sprintf("event_type ='%s'", eventKey))
	if chn != "" {
		chnList := utils.ChannelList(chn)
		session = session.In("channel", chnList)
	}
	err := session.OrderBy("event_day,channel").
		Find(&items)
	if err != nil {
		glog.Errorf("[mysql]Get the items for from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	chartApplistshowPVValue := map[string]utils.ChartLineSeries{}
	chartApplistshowUVValue := map[string]utils.ChartLineSeries{}

	for _, v := range items {
		//时间正序计算x轴的日期
		xValue := v.EventDay.Format(utils.DayTime)
		_, ok := chartXDic[xValue]
		if !ok {
			chartXDic[xValue] = true
			chartXvalue = append(chartXvalue, xValue)
		}

		idx := fmt.Sprintf("%s%s%s", v.Channel, utils.SepChar, "-")
		//计算chartApplistokPVValue数据
		val, ok := chartApplistshowPVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.ApplistshowPv))
			val.EventTime = append(val.EventTime, xValue)
			chartApplistshowPVValue[idx] = val
		} else {
			chartApplistshowPVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.ApplistshowPv)},
				EventTime: []string{xValue},
			}
		}
		//计算chartApplistokUVValue chart
		val, ok = chartApplistshowUVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.ApplistshowUv))
			val.EventTime = append(val.EventTime, xValue)
			chartApplistshowUVValue[idx] = val
		} else {
			chartApplistshowUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.ApplistshowUv)},
				EventTime: []string{xValue},
			}
		}
	}
	var chartYlines []utils.ChartSeriesYValue
	//添加第一条线
	for k, v := range chartApplistshowPVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("%s渠道ApplistshowPV趋势图", infos[0])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}

	//添加第二条线
	for k, v := range chartApplistshowUVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("%s渠道ApplistshowUv趋势图", infos[0])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}

	chartItems := &utils.ChartDetail{
		XAxis:  chartXvalue,
		Series: chartYlines,
	}
	return utils.ChartItemsMend(chartItems), err
}
