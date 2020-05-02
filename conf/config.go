package conf

// BITConfig ...
type BITConfig struct {
	DBMysql string `yaml:"db_mysql"`
	DBRedis string `yaml:"db_redis"`
}

// DefaultConfig .
var DefaultConfig = BITConfig{
	DBMysql: "pingback-0001a:Pinback-123987!@tcp(127.0.0.1:3306)/pingback?charset=utf8mb4",
	DBRedis: "redis://1:@127.0.0.1:6379",
}
