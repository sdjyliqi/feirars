package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//GetMd5ChkItems ...获取升级托包全部数据
func (pc *pingbackCenter) GetUpdateSsxfItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.UpdateSsxfWeb, int64, error) {
	items, count, err := pc.updateSSXF.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.UpdateSsxfWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.updateSSXF.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetMd5ChkCols ...前端显示的表头
func (pc *pingbackCenter) GetUpdateSsxfCols() []map[string]string {
	return pc.updateSSXF.Cols()
}

func (pc *pingbackCenter) GetUpdateSsxfChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.updateSSXF.GetAllChannels(pc.db)
	}
	chnList := strings.Split(item.Chn, ",")
	return chnList, nil
}

func (pc *pingbackCenter) GetUpdateSsxfChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.updateSSXF.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
