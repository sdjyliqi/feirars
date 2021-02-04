package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
	"time"
)

type Md5Check struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SoftName   string    `json:"soft_name" xorm:"VARCHAR(255)"`
	Channel    string    `json:"channel" xorm:"not null VARCHAR(255)"`
	EventDay   time.Time `json:"event_day" xorm:"DATETIME"`
	Mdcheck1Pv int       `json:"mdcheck1_pv" xorm:"INT(11)"`
	Mdcheck1Uv int       `json:"mdcheck1_uv" xorm:"INT(11)"`
	Mdcheck0Pv int       `json:"mdcheck0_pv" xorm:"INT(11)"`
	Mdcheck0Uv int       `json:"mdcheck0_uv" xorm:"INT(11)"`
	Appdown1Pv int       `json:"appdown1_pv" xorm:"INT(11)"`
	Appdown1Uv int       `json:"appdown1_uv" xorm:"INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"DATETIME"`
}

func (t Md5Check) TableName() string {
	return "md5_check"
}

type Md5CheckWeb struct {
	EventDay   string `json:"event_day"`
	SoftName   string `json:"soft_name"`
	Channel    string `json:"channel" `
	Mdcheck1Pv int    `json:"mdcheck1_pv" `
	Mdcheck1Uv int    `json:"mdcheck1_uv"`
	Mdcheck0Pv int    `json:"mdcheck0_pv"`
	Mdcheck0Uv int    `json:"mdcheck0_uv" `
	Appdown1Pv int    `json:"appdown1_pv"`
	Appdown1Uv int    `json:"appdown1_uv" `
	LastUpdate string `json:"last_update" `
}

func (t Md5Check) CovertWebItem(item *Md5Check) Md5CheckWeb {
	webItem := Md5CheckWeb{
		EventDay:   item.EventDay.Format(utils.DayTime),
		SoftName:   item.SoftName,
		Channel:    item.Channel,
		Mdcheck1Pv: item.Mdcheck1Pv,
		Mdcheck1Uv: item.Mdcheck1Uv,
		Mdcheck0Pv: item.Mdcheck0Pv,
		Mdcheck0Uv: item.Mdcheck0Uv,
		Appdown1Pv: item.Appdown1Pv,
		Appdown1Uv: item.Appdown1Uv,
		LastUpdate: item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols 获取所有的列名
func (t Md5Check) Cols() []map[string]string {
	var cols []map[string]string
	col := map[string]string{
		"name":  "日期",
		"key":   "event_day",
		"click": "0",
		"raw":   "EventDay",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "SoftName",
		"key":   "soft_name",
		"click": "1",
		"raw":   "SoftName",
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
		"name":  "Mdcheck1Pv",
		"key":   "mdcheck1_pv",
		"click": "0",
		"raw":   "Mdcheck1Pv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Mdcheck1Uv",
		"key":   "mdcheck1_uv",
		"click": "0",
		"raw":   "Mdcheck1Uv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Mdcheck0Pv",
		"key":   "mdcheck0_pv",
		"click": "0",
		"raw":   "Mdcheck0Pv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "Mdcheck0Uv",
		"key":   "mdcheck0_uv",
		"click": "0",
		"raw":   "Mdcheck0Uv",
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
		"name":  "Appdown1Uv",
		"key":   "appdown1_uv",
		"click": "0",
		"raw":   "Appdown1Uv",
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
func (t Md5Check) GetAllChannels(client *xorm.Engine) ([]string, error) {
	var items []*Md5Check
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

//GetItemsByPage...按照日期区间&渠道列表获取数据
func (t Md5Check) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*Md5Check, int64, error) {
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*Md5Check
	item := &Md5Check{}
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

func (t Md5Check) GetChartItems(client *xorm.Engine, chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	chartXvalue := make([]string, 0)
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*Md5Check
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
	if chn != "" {
		session = session.In("channel", []string{"all"})
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
		idx := fmt.Sprintf("%s%s%s%s", v.SoftName, utils.SepChar, v.Channel, utils.SepChar)
		//计算chartApplistokPVValue数据
		val, ok := chartApplistshowPVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.Appdown1Pv))
			val.EventTime = append(val.EventTime, xValue)
			chartApplistshowPVValue[idx] = val
		} else {
			chartApplistshowPVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.Appdown1Pv)},
				EventTime: []string{xValue},
			}
		}
		//计算chartApplistokUVValue chart
		val, ok = chartApplistshowUVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.Appdown1Uv))
			val.EventTime = append(val.EventTime, xValue)
			chartApplistshowUVValue[idx] = val
		} else {
			chartApplistshowUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.Appdown1Uv)},
				EventTime: []string{xValue},
			}
		}
	}
	var chartYlines []utils.ChartSeriesYValue
	//添加第一条线
	for k, v := range chartApplistshowPVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("%s%s渠道Appdown1PV趋势图", infos[0], infos[1])
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
		lineTitle := fmt.Sprintf("%s%s渠道Appdown1Uv趋势图", infos[0], infos[1])
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
