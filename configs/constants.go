package configs

import "time"

const (
	// MinGoVersion 最小 Go 版本
	MinGoVersion = 1.16

	// ProjectName 项目名称
	ProjectName = "gin-taro-api"

	// ProjectVersion 项目版本
	ProjectVersion = "v0.0.1"

	// ProjectDomain 项目域名
	ProjectDomain = "http://127.0.0.1"

	// ProjectPort 项目端口
	ProjectPort = ":9000"

	// ProjectInstallMark 项目安装完成标识
	ProjectInstallMark = "INSTALL.lock"

	// ProjectAccessLogFile 项目访问日志存放文件
	ProjectAccessLogFile = "./logs/" + ProjectName + "-access.log"

	// ProjectCronLogFile 项目后台任务日志存放文件
	ProjectCronLogFile = "./logs/" + ProjectName + "-cron.log"

	// HeaderSignTokenTimeout 签名有效期为 2 分钟
	HeaderSignTokenTimeout = time.Minute * 2

	// HeaderLoginToken 登录验证 Token，Header 中传递的参数
	HeaderLoginToken = "Token"

	// HeaderWeixinLoginToken 登录验证 WxToken，Header 中传递的参数
	HeaderWeixinLoginToken = "WxToken"

	// HeaderSignToken 签名验证 Authorization，Header 中传递的参数
	HeaderSignToken = "Authorization"

	// HeaderSignTokenDate 签名验证 Authorization-Date，Header 中传递的参数
	HeaderSignTokenDate = "Authorization-Date"

	// RedisKeyPrefixLoginUser Redis Key 前缀 - 登录用户信息
	RedisKeyPrefixLoginUser = ProjectName + ":login-user:"

	// RedisKeyPrefixWeixinLoginUser Redis Key 前缀 - 微信登录用户信息
	RedisKeyPrefixWeixinLoginUser = ProjectName + ":weixin-login-user:"

	// RedisKeyPrefixSignature Redis Key 前缀 - 签名验证信息
	RedisKeyPrefixSignature = ProjectName + ":signature:"

	// ZhCN 简体中文 - 中国
	ZhCN = "zh-cn"

	// EnUS 英文 - 美国
	EnUS = "en-us"

	// LoginSessionTTL 登录有效期为 24 小时
	LoginSessionTTL = time.Hour * 24
)
