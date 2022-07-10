package weixin

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

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
		// TODO: 优化判断方式
		// 限制只能传指定格式图片
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinAvatarUploadError,
				code.Text(code.WeixinAvatarUploadError)).WithError(errors.New("只允许png,jpg,gif,jpeg文件")),
			)
			return
		}

		// 文件重命名
		h := md5.New()
		h.Write([]byte(fmt.Sprintf("%s%s", file.Filename, time.Now().String())))
		fileName := hex.EncodeToString(h.Sum(nil))

		// 上传目录
		fileDir := fmt.Sprintf("%s/%d%s/", configs.WeixinUploadFileDir, time.Now().Year(), time.Now().Month().String())
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
		ctx.SaveUploadedFile(file, dst)

		ctx.PayloadStandard(dst[2:])
		// ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}
