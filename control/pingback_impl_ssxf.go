package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//GetSjtbFullItems ...获取升级托包全部数据
func (pc *pingbackCenter) GetSSXFItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.SsxfDetailWeb, int64, error) {
	items, count, err := pc.ssxfDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.SsxfDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.ssxfDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetSjtbFullCols ...前端显示的表头
func (pc *pingbackCenter) GetSSXFCols() []map[string]string {
	return pc.ssxfDetail.Cols()
}

func (pc *pingbackCenter) GetSSXFChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.ssxfDetail.GetAllChannels(pc.db)
	}
	chnList := strings.Split(item.Chn, ",")
	return chnList, nil
}

func (pc *pingbackCenter) GetSSXFChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.ssxfDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
