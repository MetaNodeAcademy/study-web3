package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	error2 "task4/pkg/error"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: msg,
		Data:    data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
func Error(c *gin.Context, appErr *error2.AppError) {
	c.JSON(http.StatusOK, Response{
		Code:    appErr.Code,
		Message: appErr.Message,
		Data:    nil,
	})
}

func FailStop(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
