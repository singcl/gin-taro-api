package core

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/proposal"
	"github.com/singcl/gin-taro-api/pkg/color"
	"github.com/singcl/gin-taro-api/pkg/env"
	"github.com/singcl/gin-taro-api/pkg/errors"
	"github.com/singcl/gin-taro-api/pkg/trace"
	"github.com/singcl/gin-taro-api/views"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

// https://patorjk.com/software/taag/#p=testall&f=ANSI%20Regular&t=gin-taro-api%0A
const _UI = `
██████╗ ██╗███╗   ██╗   ████████╗ █████╗ ██████╗  ██████╗        █████╗ ██████╗ ██╗
██╔════╝ ██║████╗  ██║   ╚══██╔══╝██╔══██╗██╔══██╗██╔═══██╗      ██╔══██╗██╔══██╗██║
██║  ███╗██║██╔██╗ ██║█████╗██║   ███████║██████╔╝██║   ██║█████╗███████║██████╔╝██║
██║   ██║██║██║╚██╗██║╚════╝██║   ██╔══██║██╔══██╗██║   ██║╚════╝██╔══██║██╔═══╝ ██║
╚██████╔╝██║██║ ╚████║      ██║   ██║  ██║██║  ██║╚██████╔╝      ██║  ██║██║     ██║
 ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝       ╚═╝  ╚═╝╚═╝     ╚═╝                                                                    
`

// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}
type Option func(*option)

type option struct {
	// disablePProf      bool
	// disableSwagger    bool
	// disablePrometheus bool
	enableCors bool
	// enableRate        bool
	enableOpenBrowser string
	alertNotify       proposal.NotifyHandler
	recordHandler     proposal.RecordHandler
}

// RouterGroup 包装gin的RouterGroup
type RouterGroup interface {
	Group(string, ...HandlerFunc) RouterGroup
	IRoutes
}

// ?
var _ IRoutes = (*router)(nil)

type IRoutes interface {
	Any(string, ...HandlerFunc)
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	OPTIONS(string, ...HandlerFunc)
	HEAD(string, ...HandlerFunc)
}

// DisableTraceLog 禁止记录日志
func DisableTraceLog(ctx Context) {
	ctx.disableTrace()
}

// DisableRecordMetrics 禁止记录指标
func DisableRecordMetrics(ctx Context) {
	ctx.disableRecordMetrics()
}

// WithEnableCors 设置支持跨域
func WithEnableCors() Option {
	return func(opt *option) {
		opt.enableCors = true
	}
}

// WithEnableOpenBrowser 启动后在浏览器中打开 uri
func WithEnableOpenBrowser(uri string) Option {
	return func(opt *option) {
		opt.enableOpenBrowser = uri
	}
}

// WithAlertNotify 设置告警通知
func WithAlertNotify(notifyHandler proposal.NotifyHandler) Option {
	return func(opt *option) {
		opt.alertNotify = notifyHandler
	}
}

// WithRecordMetrics 设置记录接口指标
func WithRecordMetrics(recordHandler proposal.RecordHandler) Option {
	return func(opt *option) {
		opt.recordHandler = recordHandler
	}
}

// WrapAuthHandler 用来处理 Auth 的入口
func WrapAuthHandler(handler func(Context) (sessionUserInfo proposal.SessionUserInfo, err BusinessError)) HandlerFunc {
	return func(ctx Context) {
		sessionUserInfo, err := handler(ctx)

		if err != nil {
			ctx.AbortWithError(err)
			return
		}
		ctx.setSessionUserInfo(sessionUserInfo)
	}
}

// WrapWeixinAuthHandler 用来处理 Auth 的入口
func WrapWeixinAuthHandler(handler func(Context) (sessionUserInfo proposal.WeixinSessionUserInfo, err BusinessError)) HandlerFunc {
	return func(ctx Context) {
		sessionUserInfo, err := handler(ctx)

		if err != nil {
			ctx.AbortWithError(err)
			return
		}
		ctx.setSessionWeixinUserInfo(sessionUserInfo)
	}
}

// AliasForRecordMetrics 对请求路径起个别名，用于记录指标。
// 如：Get /user/:username 这样的路径，因为 username 会有非常多的情况，这样记录指标非常不友好。
func AliasForRecordMetrics(path string) HandlerFunc {
	return func(ctx Context) {
		ctx.setAlias(path)
	}
}

type router struct {
	group *gin.RouterGroup
}

// router实现IRoutes接口
func (r *router) Any(relativePath string, handlers ...HandlerFunc) {
	r.group.Any(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) GET(relativePath string, handlers ...HandlerFunc) {
	r.group.GET(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) POST(relativePath string, handlers ...HandlerFunc) {
	r.group.POST(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) DELETE(relativePath string, handlers ...HandlerFunc) {
	r.group.DELETE(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PATCH(relativePath string, handlers ...HandlerFunc) {
	r.group.PATCH(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PUT(relativePath string, handlers ...HandlerFunc) {
	r.group.PUT(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	r.group.OPTIONS(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) HEAD(relativePath string, handlers ...HandlerFunc) {
	r.group.HEAD(relativePath, wrapHandlers(handlers...)...)
}

//
func (r *router) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	group := r.group.Group(relativePath, wrapHandlers(handlers...)...)
	return &router{group: group}
}

func wrapHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	funcs := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handler := handler
		funcs[i] = func(c *gin.Context) {
			ctx := newContext(c)
			defer releaseContext(ctx)

			handler(ctx)
		}
	}
	return funcs
}

type Kiko interface {
	http.Handler
	Group(relativePath string, handlers ...HandlerFunc) RouterGroup
}

type kiko struct {
	engine *gin.Engine
}

func (k *kiko) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	k.engine.ServeHTTP(w, req)
}

func (k *kiko) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	return &router{
		group: k.engine.Group(relativePath, wrapHandlers(handlers...)...),
	}
}

func New(logger *zap.Logger, options ...Option) (Kiko, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}
	// gin.SetMode(gin.ReleaseMode)

	kiko := &kiko{engine: gin.New()}

	fmt.Println(color.Blue(_UI))

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	kiko.engine.MaxMultipartMemory = 2 << 20 // 2 MiB

	// 静态资源服务
	// kiko.engine.Static("/uploads", "./uploads")
	kiko.engine.StaticFS("/uploads", http.Dir("./uploads"))
	kiko.engine.StaticFS("/views/static", http.Dir("./views/static"))
	kiko.engine.StaticFS("/views/templates", http.Dir("./views/templates"))
	kiko.engine.SetHTMLTemplate(template.Must(template.New("").ParseFS(views.Templates, "templates/**/*.tmpl")))

	// @DEBUG: DEBUG for fed live reload
	// 第一个参数静态资源前缀，第二参数静态资源目录
	// kiko.engine.Static("/uploads", "./uploads")
	// 和上面的方式功能一样，不过下面会启动一个静态资源文件系统
	// kiko.engine.StaticFS("/uploads", http.Dir("./uploads"))
	// kiko.engine.StaticFS("/views/static", http.Dir("./views/static"))
	// kiko.engine.StaticFS("/views/templates", http.Dir("./views/templates"))
	// kiko.engine.LoadHTMLGlob("views/templates/**/*.tmpl")

	// withoutTracePaths 这些请求，默认不记录日志
	withoutTracePaths := map[string]bool{
		"/metrics": true,

		"/debug/pprof/":             true,
		"/debug/pprof/cmdline":      true,
		"/debug/pprof/profile":      true,
		"/debug/pprof/symbol":       true,
		"/debug/pprof/trace":        true,
		"/debug/pprof/allocs":       true,
		"/debug/pprof/block":        true,
		"/debug/pprof/goroutine":    true,
		"/debug/pprof/heap":         true,
		"/debug/pprof/mutex":        true,
		"/debug/pprof/threadcreate": true,

		"/favicon.ico": true,

		"/system/health": true,
	}

	opt := new(option)
	for _, f := range options {
		f(opt)
	}

	if opt.enableCors {
		kiko.engine.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:     []string{"*"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
		}))
	}

	if opt.enableOpenBrowser != "" {
		_ = browser.OpenURL(opt.enableOpenBrowser)
	}

	// 暂时使用内置logger
	kiko.engine.Use(gin.Logger())

	// recover两次，防止处理时发生panic，尤其是在OnPanicNotify中。
	kiko.engine.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", string(debug.Stack())))
			}
		}()
		ctx.Next()
	})

	kiko.engine.Use(func(ctx *gin.Context) {
		if ctx.Writer.Status() == http.StatusNotFound {
			return
		}

		ts := time.Now()

		context := newContext(ctx)
		defer releaseContext(context)

		context.init()
		context.setLogger(logger)
		context.ableRecordMetrics()

		if !withoutTracePaths[ctx.Request.URL.Path] {
			if traceId := context.GetHeader(trace.Header); traceId != "" {
				context.setTrace(trace.New(traceId))
			} else {
				context.setTrace(trace.New(""))
			}
		}

		defer func() {
			var (
				response        interface{}
				traceId         string
				abortErr        error
				businessCode    int
				businessCodeMsg string
			)

			if ct := context.Trace(); ct != nil {
				context.SetHeader(trace.Header, ct.ID())
				traceId = ct.ID()
			}

			// region 发生 Panic 异常发送告警提醒
			if err := recover(); err != nil {
				stackInfo := string(debug.Stack())
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", stackInfo))
				context.AbortWithError(Error(
					http.StatusInternalServerError,
					code.ServerError,
					code.Text(code.ServerError)),
				)

				if notifyHandler := opt.alertNotify; notifyHandler != nil {
					notifyHandler(&proposal.AlertMessage{
						ProjectName:  configs.ProjectName,
						Env:          env.Active().Value(),
						TraceID:      traceId,
						HOST:         context.Host(),
						URI:          context.URI(),
						Method:       context.Method(),
						ErrorMessage: err,
						ErrorStack:   stackInfo,
						Timestamp:    time.Now(),
					})
				}
			}

			if ctx.IsAborted() {
				for i := range ctx.Errors {
					multierr.AppendInto(&abortErr, ctx.Errors[i])
				}

				if err := context.abortError(); err != nil { // customer err
					// 判断是否需要发送告警通知
					if err.IsAlert() {
						if notifyHandler := opt.alertNotify; notifyHandler != nil {
							notifyHandler(&proposal.AlertMessage{
								ProjectName:  configs.ProjectName,
								Env:          env.Active().Value(),
								TraceID:      traceId,
								HOST:         context.Host(),
								URI:          context.URI(),
								Method:       context.Method(),
								ErrorMessage: err.Message(),
								ErrorStack:   fmt.Sprintf("%+v", err.StackError()),
								Timestamp:    time.Now(),
							})
						}
					}

					multierr.AppendInto(&abortErr, err.StackError())
					businessCode = err.BusinessCode()
					businessCodeMsg = err.Message()
					response = &code.Failure{
						Code:    businessCode,
						Message: businessCodeMsg,
					}
					ctx.JSON(err.HTTPCode(), response)
				}
			}
			// 自定义ctx.Payload方式的返回
			response = context.getPayload()
			if response != nil {
				ctx.JSON(http.StatusOK, response)
			}

			// region 记录指标
			if opt.recordHandler != nil && context.isRecordMetrics() {
				path := context.Path()
				if alias := context.Alias(); alias != "" {
					path = alias
				}

				opt.recordHandler(&proposal.MetricsMessage{
					ProjectName:  configs.ProjectName,
					Env:          env.Active().Value(),
					TraceID:      traceId,
					HOST:         context.Host(),
					Path:         path,
					Method:       context.Method(),
					HTTPCode:     ctx.Writer.Status(),
					BusinessCode: businessCode,
					CostSeconds:  time.Since(ts).Seconds(),
					IsSuccess:    !ctx.IsAborted() && (ctx.Writer.Status() == http.StatusOK),
				})
			}
			// endregion

			// region 记录日志
			var t *trace.Trace
			if x := context.Trace(); x != nil {
				t = x.(*trace.Trace)
			} else {
				return
			}
			decodedURL, _ := url.QueryUnescape(ctx.Request.URL.RequestURI())
			// ctx.Request.Header，精简 Header 参数
			traceHeader := map[string]string{
				"Content-Type":              ctx.GetHeader("Content-Type"),
				configs.HeaderLoginToken:    ctx.GetHeader(configs.HeaderLoginToken),
				configs.HeaderSignToken:     ctx.GetHeader(configs.HeaderSignToken),
				configs.HeaderSignTokenDate: ctx.GetHeader(configs.HeaderSignTokenDate),
			}

			t.WithRequest(&trace.Request{
				TTL:        "un-limit",
				Method:     ctx.Request.Method,
				DecodedURL: decodedURL,
				Header:     traceHeader,
				Body:       string(context.RawData()),
			})

			var responseBody any

			if response != nil {
				responseBody = response
			}

			t.WithResponse(&trace.Response{
				Header:          ctx.Writer.Header(),
				HttpCode:        ctx.Writer.Status(),
				HttpCodeMsg:     http.StatusText(ctx.Writer.Status()),
				BusinessCode:    businessCode,
				BusinessCodeMsg: businessCodeMsg,
				Body:            responseBody,
				CostSeconds:     time.Since(ts).Seconds(),
			})

			t.Success = !ctx.IsAborted() && (ctx.Writer.Status() == http.StatusOK)
			t.CostSeconds = time.Since(ts).Seconds()

			logger.Info("trace-log",
				zap.Any("method", ctx.Request.Method),
				zap.Any("path", decodedURL),
				zap.Any("http_code", ctx.Writer.Status()),
				zap.Any("business_code", businessCode),
				zap.Any("success", t.Success),
				zap.Any("cost_seconds", t.CostSeconds),
				zap.Any("trace_id", t.Identifier),
				zap.Any("trace_info", t),
				zap.Error(abortErr),
			)
			// endregion
		}()

		ctx.Next()
	})

	system := kiko.Group("/system")
	{
		// 健康检查
		system.GET("health", func(ctx Context) {
			resp := &struct {
				Timestamp   time.Time `json:"timestamp"`
				Environment string    `json:"environment"`
				Host        string    `json:"host"`
				Status      string    `json:"status"`
			}{
				Timestamp:   time.Now(),
				Environment: env.Active().Value(),
				Host:        ctx.Host(),
				Status:      "ok",
			}
			ctx.Payload(resp)
		})
	}
	return kiko, nil
}
