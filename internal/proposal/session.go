package proposal

import "encoding/json"

// SessionUserInfo 当前用户会话信息
type SessionUserInfo struct {
	UserID   int32  `json:"user_id"`   // 用户ID
	UserName string `json:"user_name"` // 用户名
}

// Marshal 序列化到JSON
func (user *SessionUserInfo) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = json.Marshal(user)
	return
}

// weixinSessionUserInfo 当前用户会话信息
type WeixinSessionUserInfo struct {
	Openid     string `json:"openid"`      // 微信openid
	SessionKey string `json:"session_key"` // 微信session_key
}

// Marshal 序列化到JSON
func (wxUser *WeixinSessionUserInfo) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = json.Marshal(wxUser)
	return
}
