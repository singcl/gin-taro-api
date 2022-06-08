package core

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	_BodyName       = "_body_"
	_PayloadName    = "_payload_"
	_AbortErrorName = "_abort_error_"
)

type HandlerFunc func(c Context)

var _ Context = (*context)(nil)

/************************ Context接口************************/
type Context interface {
	init()
	// Payload 正确返回
	Payload(payload interface{})
	getPayload() interface{}
	// Host 获取 Request.Host
	Host() string
	// ShouldBindURI 反序列化 path 参数(如路由路径为 /user/:name)
	// tag: `uri:"xxx"`
	ShouldBindURI(obj interface{}) error
	// AbortWithError 错误返回
	AbortWithError(err BusinessError)
}

type context struct {
	ctx *gin.Context
}

/******************* context实现Context接口*******************/
func (c *context) init() {
	body, err := c.ctx.GetRawData()
	if err != nil {
		panic(err)
	}
	c.ctx.Set(_BodyName, body)                                   // cache body是为了trace使用
	c.ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // re-construct req body
}

func (c *context) Payload(payload interface{}) {
	c.ctx.Set(_PayloadName, payload)
}

func (c *context) getPayload() interface{} {
	if payload, ok := c.ctx.Get(_PayloadName); ok != false {
		return payload
	}
	return nil
}

// Host 请求的host
func (c *context) Host() string {
	return c.ctx.Request.Host
}

// ShouldBindURI 反序列化path参数(如路由路径为 /user/:name)
// tag: `uri:"xxx"`
func (c *context) ShouldBindURI(obj interface{}) error {
	return c.ctx.ShouldBindUri(obj)
}

func (c *context) AbortWithError(err BusinessError) {
	if err != nil {
		httpCode := err.HTTPCode()
		if httpCode == 0 {
			httpCode = http.StatusInternalServerError
		}

		c.ctx.AbortWithStatus(httpCode)
		c.ctx.Set(_AbortErrorName, err)
	}
}

var contextPool = &sync.Pool{
	New: func() any {
		return new(context)
	},
}

func newContext(ctx *gin.Context) Context {
	context := contextPool.Get().(*context)
	context.ctx = ctx
	return context
}

func releaseContext(ctx Context) {
	c := ctx.(*context)
	c.ctx = nil
	contextPool.Put(c)
}
