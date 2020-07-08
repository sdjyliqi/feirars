package conf

// FeirarConfig ...
type FeirarConfig struct {
	DBMysql string `yaml:"db_mysql"`
	DBRedis string `yaml:"db_redis"`
	IPLOC   string `yaml:"ip_loc"`
}

// DefaultConfig .
var DefaultConfig = FeirarConfig{
	DBMysql: "root:Bit0123456789!@tcp(114.55.139.105:3306)/pingback?charset=utf8mb4",
	DBRedis: "redis://1:@127.0.0.1:6379",
	IPLOC:   "ip2region.db",
}
