package control

import (
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/utils"
)

type UserCenter interface {
	Login(userID, passport string) error    //用户登录
	Register(userID, passport string) error //用户注册
	Logout(userID string) error             //用户登录
}

type userCenter struct {
	db    *xorm.Engine
	cfg   *conf.BITConfig
	redis *redis.Client
}

func CreateUserCenter(cfg *conf.BITConfig) UserCenter {
	utils.InitMySQL(cfg.DBMysql, false)
	utils.InitRedisClients(cfg.DBRedis)
	return &userCenter{
		db:    utils.GetMysqlClient(),
		cfg:   cfg,
		redis: utils.GetRedisClient(),
	}
}
