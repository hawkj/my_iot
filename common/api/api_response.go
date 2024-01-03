package commonapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func ApiOk(ctx *gin.Context, data interface{}) {
	r := &CommonResponse{
		Code: 0,
		Data: data,
	}
	ctx.JSON(http.StatusOK, r)
}

func ApiOkWithCode(ctx *gin.Context, data interface{}, code int) {
	r := &CommonResponse{
		Code: code,
		Data: data,
	}
	ctx.JSON(http.StatusOK, r)
}

func ApiErrorWithCode(ctx *gin.Context, data interface{}, code int) {
	r := &CommonResponse{
		Code: code,
		Data: data,
	}
	ctx.JSON(http.StatusInternalServerError, r)
}

func ApiError(ctx *gin.Context, data interface{}) {
	r := &CommonResponse{
		Code: http.StatusInternalServerError,
		Data: data,
	}
	ctx.JSON(http.StatusInternalServerError, r)
}

func HttpErrorWithCode(ctx *gin.Context, Code int, data interface{}, httpError int) {
	r := &CommonResponse{
		Code: Code,
		Data: data,
	}
	ctx.JSON(httpError, r)
}
