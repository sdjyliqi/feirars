package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
	"time"
)

type UpdateSsxf struct {
	Id                int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Channel           string    `json:"channel" xorm:"not null VARCHAR(255)"`
	EventDay          time.Time `json:"event_day" xorm:"DATETIME"`
	RefectivePv       int       `json:"refective_pv" xorm:"INT(11)"`
	RefectiveUv       int       `json:"refective_uv" xorm:"INT(11)"`
	GetupdaterecommPv int       `json:"getUpdateRecomm_pv" xorm:"INT(11)"`
	GetupdaterecommUv int       `json:"getUpdateRecomm_uv" xorm:"INT(11)"`
	UpdatePv          int       `json:"update_pv" xorm:"INT(11)"`
	UpdateUv          int       `json:"update_uv" xorm:"INT(11)"`
	ReportconfigPv    int       `json:"reportconfig_pv" xorm:"INT(11)"`
	ReportconfigUv    int       `json:"reportconfig_uv" xorm:"INT(11)"`

	LastUpdate time.Time `json:"last_update" xorm:"DATETIME"`
}

func (t UpdateSsxf) TableName() string {
	return "update_ssxf"
}

type UpdateSsxfWeb struct {
	Id                string `json:"id"`
	Channel           string `json:"channel" `
	EventDay          string `json:"event_day"`
	RefectivePv       int    `json:"refective_pv" `
	RefectiveUv       int    `json:"refective_uv" `
	GetupdaterecommPv int    `json:"getUpdateRecomm_pv" `
	GetupdaterecommUv int    `json:"getUpdateRecomm_uv"`
	UpdatePv          int    `json:"update_pv" `
	UpdateUv          int    `json:"update_uv" `
	ReportconfigPv    int    `json:"reportconfig_pv" `
	ReportconfigUv    int    `json:"reportconfig_uv"`
	LastUpdate        string `json:"last_update" `
}

func (t UpdateSsxf) CovertWebItem(item *UpdateSsxf) UpdateSsxfWeb {
	webItem := UpdateSsxfWeb{
		Id:                "",
		EventDay:          item.EventDay.Format(utils.DayTime),
		Channel:           item.Channel,
		RefectivePv:       item.RefectivePv,
		RefectiveUv:       item.RefectiveUv,
		GetupdaterecommPv: item.GetupdaterecommPv,
		GetupdaterecommUv: item.GetupdaterecommUv,
		UpdatePv:          item.UpdatePv,
		UpdateUv:          item.UpdateUv,
		ReportconfigPv:    item.ReportconfigPv,
		ReportconfigUv:    item.ReportconfigUv,
		LastUpdate:        item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols 获取所有的列名
func (t UpdateSsxf) Cols() []map[string]string {
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
		"name":  "RefectivePv",
		"key":   "refective_pv",
		"click": "0",
		"raw":   "RefectivePv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "RefectiveUv",
		"key":   "refective_uv",
		"click": "0",
		"raw":   "RefectiveUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "GetupdaterecommPv",
		"key":   "getUpdateRecomm_pv",
		"click": "0",
		"raw":   "GetupdaterecommPv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "GetupdaterecommUv",
		"key":   "getUpdateRecomm_uv",
		"click": "0",
		"raw":   "GetupdaterecommUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "UpdatePv",
		"key":   "update_pv",
		"click": "0",
		"raw":   "UpdatePv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "UpdateUv",
		"key":   "update_uv",
		"click": "0",
		"raw":   "UpdateUv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ReportconfigPv",
		"key":   "reportconfig_pv",
		"click": "0",
		"raw":   "ReportconfigPv",
	}
	cols = append(cols, col)

	col = map[string]string{
		"name":  "ReportconfigUv",
		"key":   "reportconfig_uv",
		"click": "0",
		"raw":   "ReportconfigUv",
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
func (t UpdateSsxf) GetAllChannels(client *xorm.Engine) ([]string, error) {
	var items []*UpdateSsxf
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
func (t UpdateSsxf) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*UpdateSsxf, int64, error) {
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*UpdateSsxf
	item := &UpdateSsxf{}
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

func (t UpdateSsxf) GetChartItems(client *xorm.Engine, chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	chartXvalue := make([]string, 0)
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*UpdateSsxf
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
			val.Data = append(val.Data, float64(v.UpdatePv))
			val.EventTime = append(val.EventTime, xValue)
			chartPVValue[idx] = val
		} else {
			chartPVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.UpdatePv)},
				EventTime: []string{xValue},
			}
		}
		//计算ShellconfigThreadUV chart
		val, ok = chartUVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.UpdateUv))
			val.EventTime = append(val.EventTime, xValue)
			chartUVValue[idx] = val
		} else {
			chartUVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.UpdateUv)},
				EventTime: []string{xValue},
			}
		}
	}
	var chartYlines []utils.ChartSeriesYValue
	//添加第一条线
	for k, v := range chartPVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("%s渠道UpdatePv趋势图", infos[0])
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
		lineTitle := fmt.Sprintf("%s渠道shellconfig_thread_pvUv趋势图", infos[0])
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
