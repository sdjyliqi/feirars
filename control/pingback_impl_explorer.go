package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//GetMd5ChkItems ...获取升级托包全部数据
func (pc *pingbackCenter) GetExplorerDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.ExplorerDetailWeb, int64, error) {
	items, count, err := pc.explorer.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.ExplorerDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.explorer.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetMd5ChkCols ...前端显示的表头
func (pc *pingbackCenter) GetExplorerDetailCols() []map[string]string {
	return pc.explorer.Cols()
}

func (pc *pingbackCenter) GetExplorerDetailChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.explorer.GetAllChannels(pc.db)
	}
	chnList := strings.Split(item.Chn, ",")
	return chnList, nil
}

func (pc *pingbackCenter) GetExplorerDetailChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.explorer.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
