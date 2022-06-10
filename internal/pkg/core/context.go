package core

import (
	"bytes"
	stdctx "context"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/singcl/gin-taro-api/pkg/trace"
	"go.uber.org/zap"
)

const (
	_BodyName        = "_body_"
	_PayloadName     = "_payload_"
	_AbortErrorName  = "_abort_error_"
	_LoggerName      = "_logger_"
	_TraceName       = "_trace_"
	_IsRecordMetrics = "_is_record_metrics_"
)

type HandlerFunc func(c Context)
type Trace = trace.T

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
	// ShouldBindForm 同时反序列化 querystring 和 postform;
	// 当 querystring 和 postform 存在相同字段时，postform 优先使用。
	// tag: `form:"xxx"`
	ShouldBindForm(obj interface{}) error
	// RequestContext 获取请求的 context (当 client 关闭后，会自动 canceled)
	RequestContext() StdContext

	// Trace 获取 Trace 对象
	Trace() Trace
	setTrace(trace Trace)
	disableTrace()

	// disableRecordMetrics 设置禁止记录指标
	disableRecordMetrics()
	ableRecordMetrics()
	isRecordMetrics() bool

	// Logger 获取 Logger 对象
	Logger() *zap.Logger
	setLogger(logger *zap.Logger)

	// HTML 返回界面
	HTML(name string, obj interface{})
}

type context struct {
	ctx *gin.Context
}

type StdContext struct {
	stdctx.Context
	Trace
	*zap.Logger
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

// ShouldBindForm 同时反序列化querystring和postform;
// 当querystring和postform存在相同字段时，postform优先使用。
// tag: `form:"xxx"`
func (c *context) ShouldBindForm(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.Form)
}

func (c *context) Trace() Trace {
	t, ok := c.ctx.Get(_TraceName)
	if !ok || t == nil {
		return nil
	}

	return t.(Trace)
}

func (c *context) setTrace(trace Trace) {
	c.ctx.Set(_TraceName, trace)
}

func (c *context) disableTrace() {
	c.setTrace(nil)
}

func (c *context) isRecordMetrics() bool {
	isRecordMetrics, ok := c.ctx.Get(_IsRecordMetrics)
	if !ok {
		return false
	}

	return isRecordMetrics.(bool)
}

func (c *context) ableRecordMetrics() {
	c.ctx.Set(_IsRecordMetrics, true)
}

func (c *context) disableRecordMetrics() {
	c.ctx.Set(_IsRecordMetrics, false)
}

func (c *context) Logger() *zap.Logger {
	logger, ok := c.ctx.Get(_LoggerName)
	if !ok {
		return nil
	}

	return logger.(*zap.Logger)
}

func (c *context) setLogger(logger *zap.Logger) {
	c.ctx.Set(_LoggerName, logger)
}

// RequestContext (包装 Trace + Logger) 获取请求的 context (当client关闭后，会自动canceled)
func (c *context) RequestContext() StdContext {
	return StdContext{
		//c.ctx.Request.Context(),
		stdctx.Background(),
		c.Trace(),
		c.Logger(),
	}
}

func (c *context) HTML(name string, obj interface{}) {
	c.ctx.HTML(http.StatusOK, name+".tmpl", obj)
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
