package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 0 // 成功代码
	FAILED  int = 1 // 失败代码
)

// 普通成功返回
func Success(ctx *gin.Context,v interface{}) {
	ctx.JSON(http.StatusOK,gin.H{
		"code":SUCCESS,
		"msg":"OK",
		"data":v,
	})
}

// 普通失败返回
func Failed(ctx *gin.Context,v interface{}) {
	ctx.JSON(http.StatusOK,gin.H{
		"code":FAILED,
		"msg":v,
	})
}

