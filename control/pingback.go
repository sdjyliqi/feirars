package control

import (
	"github.com/go-xorm/xorm"
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
)

type PingbackCenter interface {
	GetActiveDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.ActiveDetailWeb, int64, error) //获取客户的激活方式统计数据
	GetActiveDetailCols() []map[string]string
	GetActiveChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)
	GetActiveChannel() ([]string, error)

	GetFeirarDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.FeirarDetailWeb, int64, error) //获取feirar请求统计数据
	GetFeirarDetailCols() []map[string]string
	GetFeirarChart(chn string, tsStart, tsEnd int64) (*utils.ChartDetail, error)
	GetFeirarChannel() ([]string, error)

	GetInstallDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.InstallDetailWeb, int64, error) //获取安装统计数据
	GetInstallDetailCols() []map[string]string
	GetInstallChannel() ([]string, error)

	GetUninstallDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.UninstallDetailWeb, int64, error) //获取卸载统计数据
	GetUninstallDetailCols() []map[string]string
	GetUninstallChannel() ([]string, error)

	GetNewsDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.NewsDetailWeb, int64, error) //获取客户端弹窗统计数据
	GetNewsDetailCols() []map[string]string
	GetNewsChannel() ([]string, error)

	GetPreserveDetailItems(chn string, pageID, pageCount int, tsStart, tsEnd int64) ([]models.PreserveDetailWeb, int64, error) //获取留存统计数据
	GetPreserveDetailCols() []map[string]string
	GetPreserveChannel() ([]string, error)
}

type pingbackCenter struct {
	db              *xorm.Engine
	cfg             *conf.FeirarConfig
	activeDetail    models.ActiveDetail
	feirarDetail    models.FeirarDetail
	installDetail   models.InstallDetail
	uninstallDetail models.UninstallDetail
	newsDetail      models.NewsDetail
	preserveDetail  models.PreserveDetail
}

func CreatePingbackCenter(cfg *conf.FeirarConfig) PingbackCenter {
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
