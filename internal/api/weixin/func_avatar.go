package weixin

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/utils"
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

		fileExt := strings.ToLower(path.Ext(file.Filename))
		// 限制只能传指定格式图片
		if err := utils.CheckImage(file); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinAvatarUploadError,
				code.Text(code.WeixinAvatarUploadError)).WithError(err),
			)
			return
		}

		// 文件重命名
		fileName := utils.Md5Avatar(file.Filename)
		// 上传目录
		fileDir := fmt.Sprintf("%s/%s/", configs.WeixinUploadFileDir, utils.Md5Avatar(ctx.SessionWeixinUserInfo().Openid))
		isExist, err := utils.PathExists(fileDir)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinAvatarUploadError,
				code.Text(code.WeixinAvatarUploadError)).WithError(err),
			)
			return
		}
		if !isExist {
			// 递归创建目录，区别于os.MKdir
			os.MkdirAll(fileDir, os.ModePerm)
		}
		dst := fmt.Sprintf("%s%s%s", fileDir, fileName, fileExt)
		// log.Println(dst)
		// 上传文件至指定的完整文件路径
		err = ctx.SaveUploadedFile(file, dst)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinAvatarUploadError,
				code.Text(code.WeixinAvatarUploadError)).WithError(err),
			)
			return
		}

		dstShort := dst[2:]
		err = h.weixinService.Avatar(ctx, dstShort, ctx.SessionWeixinUserInfo().Openid)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinAvatarUploadError,
				code.Text(code.WeixinAvatarUploadError)).WithError(err),
			)
			return
		}

		ctx.PayloadStandard(dstShort)
		// ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}
