package conf

// BITConfig ...
type BITConfig struct {
	DBMysql string `yaml:"db_mysql"`
	DBRedis string `yaml:"db_redis"`
}

// DefaultConfig .
var DefaultConfig = BITConfig{
	DBMysql: "test:windtest@tcp(localhost:3306)/abcd?charset=utf8mb4",
	DBRedis: "redis://1:@127.0.0.1:6379",
}
