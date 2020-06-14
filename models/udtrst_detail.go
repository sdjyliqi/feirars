package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
	"time"
)

type UdtrstDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventDay   time.Time `json:"event_day" xorm:"not null comment('事件日期') DATETIME"`
	Channel    string    `json:"channel" xorm:"VARCHAR(64)"`
	Rst0Pv     int       `json:"rst0_pv" xorm:"comment('rst 0的数量') INT(11)"`
	Rst0Uv     int       `json:"rst0_uv" xorm:"comment('rst 0的数量') INT(11)"`
	Rst1Pv     int       `json:"rst1_pv" xorm:"INT(11)"`
	Rst1Uv     int       `json:"rst1_uv" xorm:"INT(11)"`
	Rst3Pv     int       `json:"rst3_pv" xorm:"INT(11)"`
	Rst3Uv     int       `json:"rst3_uv" xorm:"INT(11)"`
	Rst4Pv     int       `json:"rst4_pv" xorm:"INT(11)"`
	Rst4Uv     int       `json:"rst4_uv" xorm:"INT(11)"`
	LastUpdate time.Time `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`
	Detail     string    `json:"detail" xorm:"LONGTEXT"`
}

type UdtrstDetailWeb struct {
	Id         int    `json:"id" `
	EventDay   string `json:"event_day" `
	Channel    string `json:"channel" `
	Rst0Pv     string       `json:"rst0_pv" xorm:"comment('rst 0的数量') INT(11)"`
	Rst0Uv     string       `json:"rst0_uv" xorm:"comment('rst 0的数量') INT(11)"`
	Rst1Pv     string       `json:"rst1_pv" xorm:"INT(11)"`
	Rst1Uv     string       `json:"rst1_uv" xorm:"INT(11)"`
	Rst3Pv     string       `json:"rst3_pv" xorm:"INT(11)"`
	Rst3Uv     string       `json:"rst3_uv" xorm:"INT(11)"`
	Rst4Pv     string       `json:"rst4_pv" xorm:"INT(11)"`
	Rst4Uv     string       `json:"rst4_uv" xorm:"INT(11)"`
	LastUpdate string `json:"last_update" xorm:"not null comment('更新数据时间') DATETIME"`

}

func (t UdtrstDetail) TableName() string {
	return "udtrst_detail"
}


func (t UdtrstDetail) CovertWebItem(item *UdtrstDetail) UdtrstDetailWeb {
	webItem := UdtrstDetailWeb{
		Id:         item.Id,
		EventDay:   item.EventDay.Format(utils.DayTime),
		Channel:    item.Channel,
		Rst0Pv:     fmt.Sprintf("%d",item.Rst0Pv),
		Rst0Uv:     fmt.Sprintf("%d",item.Rst0Uv),
		Rst1Pv:     fmt.Sprintf("%d",item.Rst1Pv),
		Rst1Uv:     fmt.Sprintf("%d",item.Rst1Uv),
		Rst3Pv:     fmt.Sprintf("%d",item.Rst3Pv),
		Rst3Uv:     fmt.Sprintf("%d",item.Rst3Uv),
		Rst4Pv:     fmt.Sprintf("%d",item.Rst4Pv),
		Rst4Uv:     fmt.Sprintf("%d",item.Rst4Uv),
		LastUpdate: item.LastUpdate.Format(utils.FullTime),
	}
	return webItem
}

//Cols ...用户web显示使用
func (t UdtrstDetail) Cols() []map[string]string {
	var cols []map[string]string
	colEventDay := map[string]string{
		"name": "日期",
		"key":  "event_day",
	}
	cols = append(cols, colEventDay)

	colClientChannel := map[string]string{
		"name":  "渠道",
		"key":   "channel",
		"click": "1",
	}
	cols = append(cols, colClientChannel)

	colRstPv := map[string]string{
		"name":  "rst0_pv",
		"key":   "rst0_pv",
		"click": "0",
	}
	cols = append(cols, colRstPv)

	colRstUv := map[string]string{
		"name":  "rst0_uv",
		"key":   "rst0_uv",
		"click": "0",
	}
	cols = append(cols, colRstUv)


	colRstPv = map[string]string{
		"name":  "rst1_pv",
		"key":   "rst1_pv",
		"click": "0",
	}
	cols = append(cols, colRstPv)

	colRstUv = map[string]string{
		"name":  "rst1_uv",
		"key":   "rst1_uv",
		"click": "0",
	}
	cols = append(cols, colRstUv)


	colRstPv = map[string]string{
		"name":  "rst3_pv",
		"key":   "rst3_pv",
		"click": "0",
	}
	cols = append(cols, colRstPv)

	colRstUv = map[string]string{
		"name":  "rst3_uv",
		"key":   "rst3_uv",
		"click": "0",
	}
	cols = append(cols, colRstUv)

	colRstPv = map[string]string{
		"name":  "rst4_pv",
		"key":   "rst4_pv",
		"click": "0",
	}
	cols = append(cols, colRstPv)

	colRstUv = map[string]string{
		"name":  "rst4_uv",
		"key":   "rst4_uv",
		"click": "0",
	}
	cols = append(cols, colRstUv)

	colLastUpdate := map[string]string{
		"name":  "更新时间",
		"key":   "last_update",
		"click": "0",
	}
	cols = append(cols, colLastUpdate)
	return cols
}

//GetAllChannels ...获取所有渠道
func (t UdtrstDetail) GetAllChannels(client *xorm.Engine) ([]string, error) {
	var items []*NewsDetail
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

func (t UdtrstDetail) GetItemsByPage(client *xorm.Engine, chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]*UdtrstDetail, int64, error) {
	var items []*UdtrstDetail
	item := &UdtrstDetail{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
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

func (t UdtrstDetail) GetChartItems(client *xorm.Engine, chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	chartXvalue := []string{}
	chartXDic := map[string]bool{}
	timeTS, timeTE := utils.ConvertToTime(tsStart), utils.ConvertToTime(tsEnd)
	var items []*UdtrstDetail
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

	chartRST0PVValue := map[string]utils.ChartLineSeries{}
	chartRST0UVValue := map[string]utils.ChartLineSeries{}

	for _, v := range items {
		//时间正序计算x轴的日期
		xValue := v.EventDay.Format(utils.DayTime)
		_, ok := chartXDic[xValue]
		if !ok {
			chartXDic[xValue] = true
			chartXvalue = append(chartXvalue, xValue)
		}

		idx := fmt.Sprintf("%s%s%s", v.Channel, utils.SepChar, "-")
		//计算rst0 PV chart数据
		val, ok := chartRST0PVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.Rst0Uv))
			val.EventTime = append(val.EventTime, xValue)
			chartRST0PVValue[idx] = val
		} else {
			chartRST0PVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.Rst0Uv)},
				EventTime: []string{xValue},
			}
		}
		//计算RST0 UV chart
		val, ok = chartRST0UVValue[idx]
		if ok {
			val.Data = append(val.Data, float64(v.Rst0Uv))
			val.EventTime = append(val.EventTime, xValue)
			chartRST0UVValue[idx] = val
		} else {
			chartRST0UVValue[idx] = utils.ChartLineSeries{
				Data:      []float64{float64(v.Rst0Uv)},
				EventTime: []string{xValue},
			}
		}
	}
	var chartYlines []utils.ChartSeriesYValue
	//添加第一条线
	for k, v := range chartRST0PVValue {
		infos := strings.Split(k, utils.SepChar)
		lineTitle := fmt.Sprintf("%s渠道rst0_pv趋势图", infos[0])
		chartYLine := utils.ChartSeriesYValue{
			Name:      lineTitle,
			ChartType: "line",
			Data:      v.Data,
			EventTime: v.EventTime,
		}
		chartYlines = append(chartYlines, chartYLine)
	}
	//添加第二条线
	for k, v := range chartRST0UVValue {
		infos := strings.Split(k, utils.SepChar)
		//chan_
		lineTitle := fmt.Sprintf("%s渠道rst0_UV趋势图", infos[0])
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

