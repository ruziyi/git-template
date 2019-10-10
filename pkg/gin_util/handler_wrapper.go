package ginUtil

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"sync"
)

type handlerFunc func(ctx *gin.Context) error

func WrapHandler(handler handlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := handler(ctx)
		if err == nil {
			return
		}
		clientError, ok := err.(ClientError)

		if !ok {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		code, headers := clientError.ResponseHeaders()
		for k, v := range headers {
			ctx.Header(k, v)
		}
		ctx.JSON(code, clientError.ResponseBody())
	}
}

var errType = reflect.TypeOf((*error)(nil)).Elem()
var ctxType = reflect.TypeOf((*gin.Context)(nil))

var funcParamsPool = sync.Pool{
	New: func() interface{} {
		return make([]reflect.Value, 0)
	},
}

func WrapRequestHandler(fn interface{}) gin.HandlerFunc {
	fnv := reflect.ValueOf(fn)
	t := fnv.Type()
	if t.Kind() != reflect.Func {
		panic("must be function")
	}

	if t.Kind() != reflect.Func {
		panic("invalid parameter")
	}
	if t.NumIn() != 2 {
		panic("invalid parameter")
	}
	if t.NumOut() != 2 {
		panic("invalid parameter")
	}
	if !t.Out(0).Implements(errType) {
		panic("invalid parameter")
	}
	if t.In(0) != ctxType {
		panic("invalid parameter")
	}
	if t.In(1).Kind() != reflect.Ptr {
		panic("invalid parameter")
	}

	return WrapHandler(func(ctx *gin.Context) error {
		value := reflect.New(t.In(1).Elem())
		err := Bind(value, ctx)
		if err != nil {
			return NewDefaultHTTPError(err, err.Error())
		}
		if former, ok := value.Interface().(validation.ValidFormer); ok {
			valid := validation.Validation{}
			b, err := valid.Valid(former)
			if err != nil {
				return NewDefaultHTTPError(err, err.Error())
			}
			if !b {
				return NewDefaultHTTPError(nil, valid.Errors[0].Error())
			}
		}

		params := funcParamsPool.Get().([]reflect.Value)
		params = append(params, reflect.ValueOf(ctx))
		params = append(params, value)

		resp := fnv.Call(params)

		funcParamsPool.Put(params[:0])
		ret0 := resp[0].Interface()
		if ret0 != nil {
			return ret0.(error)
		}
		ctx.JSON(200, resp[1].Interface())
		return nil
	})
}
