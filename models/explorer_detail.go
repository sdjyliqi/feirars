package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
	"time"
)

type ExplorerDetail struct {
	Id                   int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Channel              string    `json:"channel" xorm:"not null VARCHAR(255)"`
	EventDay             time.Time `json:"event_day" xorm:"DATETIME"`
	ShellconfigThreadPv  int       `json:"shellconfig_thread_pv" xorm:"INT(11)"`
	ShellconfigThreadUv  int       `json:"shellconfig_thread_uv" xorm:"INT(11)"`
	HelpconfigOpenfilePv int       `json:"helpconfig_openfile_pv" xorm:"INT(11)"`
	HelpconfigOpenfileUv int       `json:"helpconfig_openfile_uv" xorm:"INT(11)"`
	HelpconfigThreadPv   int       `json:"helpconfig_thread_pv" xorm:"INT(11)"`
	HelpconfigThreadUv   int       `json:"helpconfig_thread_uv" xorm:"INT(11)"`
	KeepalivePv          int       `json:"keepalive_pv" xorm:"INT(11)"`
	KeepaliveUv          int       `json:"keepalive_uv" xorm:"INT(11)"`
	LastUpdate           time.Time `json:"last_update" xorm:"DATETIME"`
}

func (t ExplorerDetail) TableName() string {
	return "explorer_detail"
}

type ExplorerDetailWeb struct {
	Id                   string `json:"id"`
	EventDay             string `json:"event_day"`
	Channel              string `json:"channel" `
	ShellconfigThreadPv  int    `json:"shellconfig_thread_pv" xorm:"INT(11)"`
	ShellconfigThreadUv  int    `json:"shellconfig_thread_uv" xorm:"INT(11)"`
	HelpconfigOpenfilePv int    `json:"helpconfig_openfile_pv" xorm:"INT(11)"`
	HelpconfigOpenfileUv int    `json:"helpconfig_openfile_uv" xorm:"INT(11)"`
	HelpconfigThreadPv   int    `json:"helpconfig_thread_pv" xorm:"INT(11)"`
	HelpconfigThreadUv   int    `json:"helpconfig_thread_uv" xorm:"INT(11)"`
	KeepalivePv          int    `json:"keepalive_pv" xorm:"INT(11)"`
	KeepaliveUv          int    `json:"keepalive_uv" xorm:"INT(11)"`
	LastUpdate           string `json:"last_update" `
}

func (t ExplorerDetail) CovertWebItem(item *ExplorerDetail) ExplorerDetailWeb {
	webItem := ExplorerDetailWeb{
		Id:                   "",
		EventDay:             item.EventDay.Format(utils.DayTime),
		Channel:              item.Channel,
		ShellconfigThreadPv:  item.ShellconfigThreadPv,
		ShellconfigThreadUv:  item.ShellconfigThreadUv,
		HelpconfigOpenfilePv: item.HelpconfigOpenfilePv,
		HelpconfigOpenfileUv: item.HelpconfigOpenfileUv,
		HelpconfigThreadPv:   item.HelpconfigThreadPv,
		HelpconfigThreadUv:   item.HelpconfigThreadUv,
		KeepalivePv:          item.KeepalivePv,
		KeepaliveUv:          item.KeepaliveUv,
		LastUpdate:           item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols 获取所有的列名
func (t ExplorerDetail) Cols() []map[string]string {
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
		"name":  "ShellconfigThreadPv",
		"key":   "shellconfig_thread_pv",
		"click": "0",
		"raw":   "ShellconfigThreadPv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ShellconfigThreadUv",
		"key":   "shellconfig_thread_uv",
		"click": "0",
		"raw":   "ShellconfigThreadUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "HelpconfigOpenfilePv",
		"key":   "helpconfig_openfile_pv",
		"click": "0",
		"raw":   "HelpconfigOpenfilePv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "HelpconfigOpenfileUv",
		"key":   "helpconfig_openfile_uv",
		"click": "0",
		"raw":   "HelpconfigOpenfileUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "HelpconfigThreadPv",
		"key":   "helpconfig_thread_pv",
		"click": "0",
		"raw":   "HelpconfigThreadPv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "HelpconfigThreadUv",
		"key":   "helpconfig_thread_uv",
		"click": "0",
		"raw":   "HelpconfigThreadUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "KeepalivePv",
		"key":   "keepalive_pv",
		"click": "0",
		"raw":   "KeepalivePv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "KeepaliveUv",
		"key":   "keepalive_uv",
		"click": "0",
		"raw":   "KeepaliveUv",
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
func (t ExplorerDetail) GetAllChannels(client *xorm.Engine) ([]string, error) {
	var items []*ExplorerDetail
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
func (t ExplorerDetail) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*ExplorerDetail, int64, error) {
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*ExplorerDetail
	item := &ExplorerDetail{}
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

func (t ExplorerDetail) GetChartItems(client *xorm.Engine, chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	chartXvalue := make([]string, 0)
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*ExplorerDetail
	session := client.Where("event_day>=?", timeTS).And("event_day<=?", timeTE) //.And(fmt.Sprintf("event_type ='%s'", eventKey))
	if chn != "" {
		//chnList := utils.ChannelList(chn)
		session = session.In("channel", []string{"all"})
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
		idx := fmt.Sprintf("%s%s%s", v.Channel, utils.SepChar, "-")
		//计算ShellconfigThreadPv数据
		val, ok := chartPVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.ShellconfigThreadPv))
			val.EventTime = append(val.EventTime, xValue)
			chartPVValue[idx] = val
		} else {
			chartPVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.ShellconfigThreadPv)},
				EventTime: []string{xValue},
			}
		}
		//计算ShellconfigThreadUV chart
		val, ok = chartUVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.ShellconfigThreadUv))
			val.EventTime = append(val.EventTime, xValue)
			chartUVValue[idx] = val
		} else {
			chartUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.ShellconfigThreadUv)},
				EventTime: []string{xValue},
			}
		}
	}
	var chartYlines []utils.ChartSeriesYValue
	//添加第一条线
	for k, v := range chartPVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("%s渠道shellconfig_thread_pv趋势图", infos[0])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}
	//添加第二条线
	for k, v := range chartUVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("%s渠道shellconfig_thread_Uv趋势图", infos[0])
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
