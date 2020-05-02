package control

import (
	"github.com/sdjyliqi/feirars/models"
)

//GetActiveDetailItems ...获取客户的激活方式统计数据
func (pc *pingbackCenter) GetActiveDetailItems(pageID, pageCount int) ([]models.ActiveDetailWeb, int64, error) { //按照页面获取统计激活方式数据
	items, count, err := pc.activeDetail.GetItemsByPage(pc.db, pageID, pageCount)
	if err != nil {
		return nil, 0, nil
	}
	webTtems := make([]models.ActiveDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.activeDetail.CovertWebItem(v)
		webTtems = append(webTtems, wItem)
	}
	return webTtems, count, nil
}

//GetFeirarDetailItems ...获取feirar 接口统计数据
func (pc *pingbackCenter) GetFeirarDetailItems(pageID, pageCount int) ([]models.FeirarDetailWeb, int64, error) {
	items, count, err := pc.feirarDetail.GetItemsByPage(pc.db, pageID, pageCount)
	if err != nil {
		return nil, 0, nil
	}
	webTtems := make([]models.FeirarDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.feirarDetail.CovertWebItem(v)
		webTtems = append(webTtems, wItem)
	}
	return webTtems, count, nil
}

//获取安装统计数据
func (pc *pingbackCenter) GetInstallDetailItems(pageID, pageCount int) ([]models.InstallDetailWeb, int64, error) {
	items, count, err := pc.installDetail.GetItemsByPage(pc.db, pageID, pageCount)
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

func (pc *pingbackCenter) GetUninstallDetailItems(pageID, pageCount int) ([]models.UninstallDetailWeb, int64, error) {
	items, count, err := pc.uninstallDetail.GetItemsByPage(pc.db, pageID, pageCount)
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

//GetNewsDetailItems ...获取客户端咨询弹窗相关接口
func (pc *pingbackCenter) GetNewsDetailItems(pageID, pageCount int) ([]models.NewsDetailWeb, int64, error) {
	items, count, err := pc.newsDetail.GetItemsByPage(pc.db, pageID, pageCount)
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

//GetPreserveDetailItems ...获取留存统计相关数据
func (pc *pingbackCenter) GetPreserveDetailItems(pageID, pageCount int) ([]models.PreserveDetailWeb, int64, error) {
	items, count, err := pc.preserveDetail.GetItemsByPage(pc.db, pageID, pageCount)
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
