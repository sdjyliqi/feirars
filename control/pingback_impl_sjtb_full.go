package control

import (
	"github.com/sdjyliqi/feirars/models"
	"strings"
)

//GetSjtbFullItems ...获取升级托包全部数据
func (pc *pingbackCenter) GetSjtbFullItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.SjtbFullWeb, int64, error) {
	items, count, err := pc.sjtbFull.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.SjtbFullWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.sjtbFull.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetSjtbFullCols ...前端显示的表头
func (pc *pingbackCenter) GetSjtbFullCols() []map[string]string {
	return pc.sjtbFull.Cols()
}

func (pc *pingbackCenter) GetSjtbFullChannel(name string) ([]string, error) {
	item, err := pc.userBasic.GetUserBasic(pc.db, name)
	if err != nil {
		return nil, err
	}
	if item.Chn == "" {
		return pc.sjtbFull.GetAllChannels(pc.db)
	}
	chnList := strings.Split(item.Chn, ",")
	return chnList, nil
}
