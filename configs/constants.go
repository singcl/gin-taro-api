package configs

const (
	// ProjectName 项目名称
	ProjectName = "gin-taro-api"

	// ProjectPort 项目端口
	ProjectPort = ":9000"

	// RedisKeyPrefixLoginUser Redis Key 前缀 - 登录用户信息
	RedisKeyPrefixLoginUser = ProjectName + ":login-user:"
)
