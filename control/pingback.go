package control

import (
	"github.com/go-xorm/xorm"
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
)

type PingbackCenter interface {
	GetActiveDetailItems(pageID, pageCount int) ([]models.ActiveDetailWeb, int64, error) //获取客户的激活方式统计数据
	GetActiveDetailCols() []map[string]string
	GetFeirarDetailItems(pageID, pageCount int) ([]models.FeirarDetailWeb, int64, error) //获取feirar请求统计数据
	GetFeirarDetailCols() []map[string]string
	GetInstallDetailItems(pageID, pageCount int) ([]models.InstallDetailWeb, int64, error) //获取安装统计数据
	GetInstallDetailCols() []map[string]string
	GetUninstallDetailItems(pageID, pageCount int) ([]models.UninstallDetailWeb, int64, error) //获取卸载统计数据
	GetUninstallDetailCols() []map[string]string
	GetNewsDetailItems(pageID, pageCount int) ([]models.NewsDetailWeb, int64, error) //获取客户端弹窗统计数据
	GetNewsDetailCols() []map[string]string
	GetPreserveDetailItems(pageID, pageCount int) ([]models.PreserveDetailWeb, int64, error) //获取留存统计数据
	GetPreserveDetailCols() []map[string]string
}

type pingbackCenter struct {
	db              *xorm.Engine
	cfg             *conf.BITConfig
	activeDetail    models.ActiveDetail
	feirarDetail    models.FeirarDetail
	installDetail   models.InstallDetail
	uninstallDetail models.UninstallDetail
	newsDetail      models.NewsDetail
	preserveDetail  models.PreserveDetail
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
