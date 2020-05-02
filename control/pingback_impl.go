package control

import (
	"github.com/sdjyliqi/feirars/models"
)



func (pc *pingbackCenter) GetActiveDetailItems(pageID,pageCount int)([]models.ActiveDetailWeb,int64,error) {   //按照页面获取统计激活方式数据
	items,count,err :=  pc.activeDetail.GetItemsByPage(pc.db,pageID,pageCount)
    if err != nil {
		return nil, 0, nil
	}
	webTtems:= make([]models.ActiveDetailWeb,0,len(items))
    for _,v := range items {
    	wItem :=  pc.activeDetail.CovertWebItem(v)
		webTtems = append(webTtems,wItem)
	}
	return webTtems,count,nil
}

func (pc *pingbackCenter) GetFeirarDetailItems(pageID,pageCount int) ([]models.FeirarDetailWeb,int64,error) {
	items,count,err :=  pc.feirarDetail.GetItemsByPage(pc.db,pageID,pageCount)
	if err != nil {
		return nil, 0, nil
	}
	webTtems:= make([]models.FeirarDetailWeb,0,len(items))
	for _,v := range items {
		wItem :=  pc.feirarDetail.CovertWebItem(v)
		webTtems = append(webTtems,wItem)
	}
	return webTtems,count,nil
}
