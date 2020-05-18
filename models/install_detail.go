package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
	"time"
)

type InstallDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventDay   time.Time `json:"event_day" xorm:"not null comment('事件日期') DATETIME"`
	Channel    string    `json:"channel" xorm:"VARCHAR(64)"`
	Pv         int       `json:"pv" xorm:"comment('PV用户数') INT(11)"`
	Uv         int       `json:"uv" xorm:"comment('UV用户数') INT(11)"`
	Day1Active int       `json:"day1_active" xorm:"INT(11)"`
	Day2Active int       `json:"day2_active" xorm:"INT(11)"`
	Day3Active int       `json:"day3_active" xorm:"INT(11)"`
	Day4Active int       `json:"day4_active" xorm:"INT(11)"`
	Day5Active int       `json:"day5_active" xorm:"INT(11)"`
	Day6Active int       `json:"day6_active" xorm:"INT(11)"`
	WeekActive int       `json:"week_active" xorm:"INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`
	Detail     string    `json:"detail" xorm:"TEXT"`
}

type InstallDetailWeb struct {
	Id         string `json:"id" `
	EventDay   string `json:"event_day"`
	Channel    string `json:"channel"`
	Pv         string `json:"pv"`
	Uv         string `json:"uv"`
	Day1Active string `json:"day1_active"`
	Day2Active string `json:"day2_active" `
	Day3Active string `json:"day3_active" `
	Day4Active string `json:"day4_active"`
	Day5Active string `json:"day5_active" `
	Day6Active string `json:"day6_active"`
	WeekActive string `json:"week_active" `
	LastUpdate string `json:"last_update"`
	Detail     string `json:"detail"`
}

func (t InstallDetail) TableName() string {
	return "install_detail"
}

func (t InstallDetail) CovertWebItem(item *InstallDetail) InstallDetailWeb {
	webItem := InstallDetailWeb{
		EventDay:   item.EventDay.Format(utils.DayTime),
		Channel:    item.Channel,
		Pv:         fmt.Sprintf("%d", item.Pv),
		Uv:         fmt.Sprintf("%d", item.Uv),
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
func (t InstallDetail) Cols() []map[string]string {
	var cols []map[string]string
	col_event_day := map[string]string{
		"name":  "日期",
		"key":   "event_day",
		"click": "0",
	}
	cols = append(cols, col_event_day)

	col_channel := map[string]string{
		"name":  "渠道",
		"key":   "channel",
		"click": "1",
	}
	cols = append(cols, col_channel)

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
		"name":  "更新时间",
		"key":   "last_update",
		"click": "0",
	}
	cols = append(cols, col_last_update)
	return cols
}

//GetAllChannels ...获取所有渠道
func (t InstallDetail) GetAllChannels(client *xorm.Engine) ([]string, error) {
	var items []*InstallDetail
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

func (t InstallDetail) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*InstallDetail, int64, error) {
	var items []*InstallDetail
	item := &InstallDetail{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
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

func (t InstallDetail) GetChartItems(client *xorm.Engine, chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	chartXvalue := []string{}
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*InstallDetail
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE)
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

	chartPVValue := map[string]utils.ChartLineSeries{}
	chartUVValue := map[string]utils.ChartLineSeries{}
	for _, v := range items {
		//时间正序计算x轴的日期
		xValue := v.EventDay.Format(utils.DayTime)
		_, ok := chartXDic[xValue]
		if !ok {
			chartXDic[xValue] = true
			chartXvalue = append(chartXvalue, xValue)
		}
		//计算pv chart数据
		idx := fmt.Sprintf("%s%s%s", v.Channel, utils.SepChar, "-")
		val, ok := chartPVValue[idx]
		//pv chart
		if ok {
			val.Data = append(val.Data, float64(v.Pv))
			val.EventTime = append(val.EventTime, xValue)
			chartPVValue[idx] = val
		} else {
			chartPVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.Pv)},
				EventTime: []string{xValue},
			}
		}
		//计算UV chart
		val, ok = chartUVValue[idx]
		//pv chart
		if ok {
			val.Data = append(val.Data, float64(v.Uv))
			val.EventTime = append(val.EventTime, xValue)
			chartUVValue[idx] = val
		} else {
			chartUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.Uv)},
				EventTime: []string{xValue},
			}
		}
	}

	var chartYlines []utils.ChartSeriesYValue
	for k, v := range chartPVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("渠道%sPV趋势图", infos[0])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}
	for k, v := range chartUVValue {
		infos := strings.Split(k, utils.SepChar)
		//chan_
		lineTitle := fmt.Sprintf("渠道%sUV趋势图", infos[0])
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
