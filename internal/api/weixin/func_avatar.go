package weixin

import (
	"log"
	"net/http"

	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
)

func (h *handler) Avatar() core.HandlerFunc {
	return func(ctx core.Context) {
		// 单文件
		file, err := ctx.ShouldBindFile("file")
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinAvatarUploadError,
				code.Text(code.WeixinAvatarUploadError)).WithError(err),
			)
			return
		}

		log.Println(file.Filename)
		dst := configs.WeixinUploadFileDir + "/" + file.Filename
		// 上传文件至指定的完整文件路径
		ctx.SaveUploadedFile(file, dst)

		ctx.PayloadStandard(dst)
		// ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}
