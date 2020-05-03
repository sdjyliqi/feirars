package control

//
//type UserCenter interface {
//	Login(userID, passport string) error    //用户登录
//	Register(userID, passport string) error //用户注册
//	Logout(userID string) error             //用户登录
//}
//
//type userCenter struct {
//	db    *xorm.Engine
//	cfg   *conf.BITConfig
//}
//
//func CreateUserCenter(cfg *conf.BITConfig) UserCenter {
//	utils.InitMySQL(cfg.DBMysql, false)
//	return &userCenter{
//		db:    utils.GetMysqlClient(),
//		cfg:   cfg,
//	}
//}
