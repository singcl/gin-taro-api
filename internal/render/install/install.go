package install

import (
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
	return func(c core.Context) {
		c.HTML("install_view", &struct{}{})
	}
}
