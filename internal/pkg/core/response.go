package core

import (
	"github.com/gin-gonic/gin"
	"time"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Ts      int64       `json:"ts"`
	Data    interface{} `json:"data"`
}

type Option func(data *ResponseData)

func Code(code int) Option {
	return func(data *ResponseData) {
		data.Code = code
	}
}

func Message(msg string) Option {
	return func(data *ResponseData) {
		data.Message = msg
	}
}

func Response(data interface{}, opts ...Option) *ResponseData {
	r := &ResponseData{
		Code: 200,
		Ts:   time.Now().Unix(),
		Data: data,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r ResponseData) Json(ctx *gin.Context) {
	ctx.JSON(r.Code, r)
}
