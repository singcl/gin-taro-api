package core

import "github.com/singcl/gin-taro-api/internal/code"

func SuccessStandard(res interface{}) *code.Success {
	return &code.Success{
		Code:    0,
		Message: "success",
		Data:    res,
	}
}
