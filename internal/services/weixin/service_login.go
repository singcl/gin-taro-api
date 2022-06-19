package weixin

import "github.com/singcl/gin-taro-api/internal/pkg/core"

type SearchCode2SessionData struct {
	Appid     string `json: "appid"`      // 小程序 appId
	Secret    string `json: "select"`     // 小程序 appSecret
	JsCode    string `json: "js_code"`    // 登录时获取的 code
	GrantType string `json: "grant_type"` // 授权类型，此处只需填写 authorization_code
}

type Code2SessionData struct {
	Openid     string `json: "openid"`      // 用户唯一标识
	SessionKey string `json: "session_key"` // 会话密钥
	Unionid    string `json: "unionid"`     // 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见 UnionID 机制说明。
	Errcode    int32  `json: "errcode"`     // 错误码
	Errmsg     string `json: "errmsg"`      // 错误信息
}

func (s *service) Login(ctx core.Context, searchCode2Session *SearchCode2SessionData) (info *Code2SessionData, err error) {
	// TODO
	return
}
