package core

import "github.com/singcl/gin-taro-api/pkg/errors"

var _ BusinessError = (*businessError)(nil)

type BusinessError interface {
	// i 为了避免被其他包实现
	i()
	// HTTPCode 获取 HTTP 状态码
	HTTPCode() int
	// WithError 设置错误信息
	WithError(err error) BusinessError
}

type businessError struct {
	httpCode     int    // HTTP 状态码
	businessCode int    // 业务码
	message      string // 错误描述
	stackError   error  // 含有堆栈信息的错误
	isAlert      bool   // 是否告警通知
}

func Error(httpCode, businessCode int, message string) BusinessError {
	return &businessError{
		httpCode:     httpCode,
		businessCode: businessCode,
		message:      message,
		isAlert:      false,
	}
}

func (e *businessError) i() {}

func (e *businessError) HTTPCode() int {
	return e.httpCode
}

func (e *businessError) WithError(err error) BusinessError {
	e.stackError = errors.WithStack(err)
	return e
}
