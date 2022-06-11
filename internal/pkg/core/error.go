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
	// Message 获取错误描述
	Message() string
	// IsAlert 是否开启告警通知
	IsAlert() bool
	// StackError 获取带堆栈的错误信息
	StackError() error
	// BusinessCode 获取业务码
	BusinessCode() int
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

func (e *businessError) IsAlert() bool {
	return e.isAlert
}

func (e *businessError) Message() string {
	return e.message
}

func (e *businessError) StackError() error {
	return e.stackError
}

func (e *businessError) BusinessCode() int {
	return e.businessCode
}
