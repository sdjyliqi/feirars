package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
)

//GetActiveDetailItems ...获取客户的激活方式统计数据
func (pc *pingbackCenter) GetActiveDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.ActiveDetailWeb, int64, error) { //按照页面获取统计激活方式数据
	items, count, err := pc.activeDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.ActiveDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.activeDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

func (pc *pingbackCenter) GetActiveDetailCols() []map[string]string {
	return pc.activeDetail.Cols()
}

func (pc *pingbackCenter) GetActiveChannel() ([]string, error) {
	return pc.activeDetail.GetAllChannels(pc.db)
}

func (pc *pingbackCenter) GetActiveChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.activeDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}

//GetFeirarDetailItems ...获取feirar 接口统计数据
func (pc *pingbackCenter) GetFeirarDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.FeirarDetailWeb, int64, error) {
	items, count, err := pc.feirarDetail.GetItemsByPage(pc.db, chn, pageID, pageCount, tsStart, tsEnd)
	if err != nil {
		return nil, 0, nil
	}
	webItems := make([]models.FeirarDetailWeb, 0, len(items))
	for _, v := range items {
		wItem := pc.feirarDetail.CovertWebItem(v)
		webItems = append(webItems, wItem)
	}
	return webItems, count, nil
}

//GetFeirarDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetFeirarDetailCols() []map[string]string {
	return pc.feirarDetail.Cols()
}

func (pc *pingbackCenter) GetFeirarChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error) {
	return pc.feirarDetail.GetChartItems(pc.db, chn, tsStart, tsEnd)
}
func (pc *pingbackCenter) GetFeirarChannel() ([]string, error) {
	return pc.feirarDetail.GetAllChannels(pc.db)
}

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

//GetFeirarDetailCols ...前端显示的表头
func (pc *pingbackCenter) GetInstallDetailCols() []map[string]string {
	return pc.installDetail.Cols()
}
func (pc *pingbackCenter) GetInstallChannel() ([]string, error) {
	return pc.installDetail.GetAllChannels(pc.db)
}

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

func (pc *pingbackCenter) GetUninstallChannel() ([]string, error) {
	return pc.uninstallDetail.GetAllChannels(pc.db)
}

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

func (pc *pingbackCenter) GetPreserveChannel() ([]string, error) {
	return pc.preserveDetail.GetAllChannels(pc.db)
}
