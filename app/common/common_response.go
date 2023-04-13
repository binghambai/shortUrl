package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Context interface{} `json:"context"`
}

type Wrapper struct {
	ctx *gin.Context
}

func WrapperContext(ctx *gin.Context) *Wrapper {
	return &Wrapper{ctx: ctx}
}

func newResponse() *Response {
	return &Response{
		Code:    200,
		Msg:     "success",
		Context: nil,
	}
}

func (w *Wrapper) Success(data interface{}) {
	resp := newResponse()
	resp.Context = data
	w.ctx.JSON(http.StatusOK, resp)
}

func (w *Wrapper) Error(statusCode int, errorMsg string) {
	resp := newResponse()
	resp.Msg = errorMsg
	resp.Code = statusCode
	w.ctx.JSON(statusCode, resp)
}
