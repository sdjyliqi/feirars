package control

import "github.com/sdjyliqi/feirars/models"

//GetNewsDetailItems ...获取客户端咨询弹窗相关接口
func (pc *pingbackCenter) GetNewsDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.NewsDetailWeb, int64, error) {
	items, count, err := pc.newsDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.NewsDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.newsDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetNewsDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetNewsDetailCols() []map[string]string {
	return pc.newsDetail.Cols()
}

func (pc *pingbackCenter) GetNewsChannel() ([]string, error) {
	return pc.newsDetail.GetAllChannels(pc.db)
}
