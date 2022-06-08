package configs

type Config struct {
	Redis struct {
		Addr         string `toml:"addr"`
		Pass         string `toml:"pass"`
		Db           int    `toml:"db"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	} `toml:"redis"`
}

// 生产配置
var config = new(Config)

// 获取配置
func Get() Config {
	return *config
}
