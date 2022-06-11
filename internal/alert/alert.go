package alert

import (
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/proposal"
	"go.uber.org/zap"
)

// NotifyHandler 告警通知
func NotifyHandler(logger *zap.Logger) func(msg *proposal.AlertMessage) {
	if logger == nil {
		panic("logger required")
	}

	return func(msg *proposal.AlertMessage) {
		cfg := configs.Get().Mail
		if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Pass == "" || cfg.To == "" {
			logger.Error("Mail config error")
			return
		}
	}
}
