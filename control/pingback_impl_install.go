package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
)

//获取安装统计数据
func (pc *pingbackCenter) GetInstallDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.InstallDetailWeb, int64, error) {
	items, count, err := pc.installDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.InstallDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.installDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetInstallDetailCols ...
func (pc *pingbackCenter) GetInstallDetailCols() []map[string]string {
	return pc.installDetail.Cols()
}

//GetInstallChannel ...获取安装统计项的所有渠道列表
func (pc *pingbackCenter) GetInstallChannel() ([]string, error) {
	return pc.installDetail.GetAllChannels(pc.db)
}

//GetInstallNewsChart ...获取渠道统计安装的趋势图数据
func (pc *pingbackCenter) GetInstallChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.installDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
