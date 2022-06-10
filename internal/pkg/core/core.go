package core

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/singcl/gin-taro-api/pkg/color"
	"github.com/singcl/gin-taro-api/pkg/env"
	"github.com/singcl/gin-taro-api/pkg/errors"
	"github.com/singcl/gin-taro-api/public"
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

type Option func(*option)

type option struct {
	// disablePProf      bool
	// disableSwagger    bool
	// disablePrometheus bool
	enableCors bool
	// enableRate        bool
	// enableOpenBrowser string
	// alertNotify       proposal.NotifyHandler
	// recordHandler     proposal.RecordHandler
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
	gin.SetMode(gin.ReleaseMode)

	kiko := &kiko{engine: gin.New()}

	fmt.Println(color.Blue(_UI))

	// 静态资源服务
	kiko.engine.StaticFS("public", http.FS(public.Public))
	kiko.engine.SetHTMLTemplate(template.Must(template.New("").ParseFS(views.Templates, "templates/**/*")))

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

		// ts := time.Now()

		context := newContext(ctx)
		defer releaseContext(context)

		context.init()
		context.setLogger(logger)

		defer func() {
			var (
				response interface{}
				abortErr error
			)
			if ctx.IsAborted() {
				for i := range ctx.Errors {
					multierr.AppendInto(&abortErr, ctx.Errors[i])
				}
			}
			// region 正确返回
			response = context.getPayload()
			if response != nil {
				ctx.JSON(http.StatusOK, response)
			}
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
