package ginUtil

import (
	"github.com/bytedance/go-tagexpr/binding"
	"github.com/gin-gonic/gin"
)

var binder = binding.New(nil)

func BindAndValidate(structPointer interface{}, ctx *gin.Context) error {
	return binder.BindAndValidate(structPointer, ctx.Request, ctx.Params)
}
func Bind(structPointer interface{}, ctx *gin.Context) error {
	return binder.Bind(structPointer, ctx.Request, ctx.Params)
}
