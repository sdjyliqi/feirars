package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
)

//GetPreserveDetailItems ...获取留存统计相关数据
func (pc *pingbackCenter) GetPreserveDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.PreserveDetailWeb, int64, error) {
	items, count, err := pc.preserveDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.PreserveDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.preserveDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetPreserveDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetPreserveDetailCols() []map[string]string {
	return pc.preserveDetail.Cols()
}

//GetPreserveChannel ...获取留存的所有渠道
func (pc *pingbackCenter) GetPreserveChannel() ([]string, error) {
	return pc.preserveDetail.GetAllChannels(pc.db)
}

//GetPreserveChart...基于渠道的留存统计图
func (pc *pingbackCenter) GetPreserveChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.preserveDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
