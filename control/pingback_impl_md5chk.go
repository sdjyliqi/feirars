package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strings"
)

//GetMd5ChkItems ...获取升级托包全部数据
func (pc *pingbackCenter) GetMd5ChkItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.Md5CheckWeb, int64, error) {
	items, count, err := pc.md5chk.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.Md5CheckWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.md5chk.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetMd5ChkCols ...前端显示的表头
func (pc *pingbackCenter) GetMd5ChkCols() []map[string]string {
	return pc.md5chk.Cols()
}

func (pc *pingbackCenter) GetMd5ChkChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.md5chk.GetAllChannels(pc.db)
	}
	chnList := strings.Split(item.Chn, ",")
	return chnList, nil
}

func (pc *pingbackCenter) GetMd5ChkChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.md5chk.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
