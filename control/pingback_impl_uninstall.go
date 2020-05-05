package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
)

func (pc *pingbackCenter) GetUninstallDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.UninstallDetailWeb, int64, error) {
	items, count, err := pc.uninstallDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.UninstallDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.uninstallDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetUninstallDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetUninstallDetailCols() []map[string]string {
	return pc.uninstallDetail.Cols()
}

//GetUninstallChannel...获取卸载事件的所有客户端渠道
func (pc *pingbackCenter) GetUninstallChannel() ([]string, error) {
	return pc.uninstallDetail.GetAllChannels(pc.db)
}

//GetUninstallChart ...获取卸载统计趋势图数据
func (pc *pingbackCenter) GetUninstallChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.uninstallDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
