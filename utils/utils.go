package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

const (
	saltAvatar = "SKDedk49sjkc"
)

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	//如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

func Md5Avatar(name string) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s%s", name, saltAvatar)))
	return hex.EncodeToString(h.Sum(nil))
}

// 图片验证
func CheckImage(file *multipart.FileHeader) (err error) {
	fileExt := strings.ToLower(path.Ext(file.Filename))
	// TODO: 优化判断方式
	// 限制只能传指定格式图片
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		err = errors.New("只允许png,jpg,gif,jpeg文件")
		return
	}
	if file.Size > 200*1024 {
		err = errors.New("图片大小不得大于200KB")
		return
	}
	return nil
}
