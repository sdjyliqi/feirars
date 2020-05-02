package control

import (
	"github.com/go-xorm/xorm"
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
)

type PingbackCenter interface {
	GetActiveDetailItems(pageID,pageCount int) ([]models.ActiveDetailWeb,int64,error)    //按照页面获取统计激活方式数据
	GetFeirarDetailItems(pageID,pageCount int) ([]models.FeirarDetailWeb,int64,error)    //按照页面获取统计激活方式数据
}

type pingbackCenter struct {
	db    *xorm.Engine
	cfg   *conf.BITConfig
	activeDetail models.ActiveDetail
	feirarDetail models.FeirarDetail
	installDetail models.InstallDetail
	uninstallDetail models.UninstallDetail
	newsDetail models.NewsDetail
	preserveDetail models.PreserveDetail
	
}

func CreatePingbackCenter(cfg *conf.BITConfig) PingbackCenter {
	 utils.InitMySQL(cfg.DBMysql, false)
	return &pingbackCenter{
		db:              utils.GetMysqlClient(),
		cfg:             cfg,
		activeDetail:    models.ActiveDetail{},
		feirarDetail:    models.FeirarDetail{},
		installDetail:   models.InstallDetail{},
		uninstallDetail: models.UninstallDetail{},
		newsDetail:      models.NewsDetail{},
		preserveDetail:  models.PreserveDetail{},
	}
}
