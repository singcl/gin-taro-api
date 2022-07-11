package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
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

// TODO: 优化判断方式
// 限制只能传指定格式图片
func CheckImage(fileExt string) bool {
	return fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg"
}
