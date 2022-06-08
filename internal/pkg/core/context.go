package core

import (
	"bytes"
	"io/ioutil"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	_BodyName    = "_body_"
	_PayloadName = "_payload_"
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
