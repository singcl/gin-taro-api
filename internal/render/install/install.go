package install

import (
	"runtime"

	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) View() core.HandlerFunc {
	type viewResponse struct {
		Config       configs.Config
		MinGoVersion float64
		GoVersion    string
	}

	return func(c core.Context) {
		obj := new(viewResponse)
		obj.Config = configs.Get()
		obj.MinGoVersion = configs.MinGoVersion
		obj.GoVersion = runtime.Version()
		c.HTML("install_view", obj)
	}
}
