package weixin

import "time"

// 微信用户表
//go:generate gormgen -structs Weixin -input .
type Weixin struct {
	Id          int32     // 主键
	Openid      string    // 微信openid
	Unionid     string    // 微信unionid
	SessionKey  string    // 微信session_key
	Nickname    string    // 昵称
	AvatarUrl   string    // 头像
	Mobile      string    // 手机号
	IsUsed      int32     // 是否启用 1:是  -1:否
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
