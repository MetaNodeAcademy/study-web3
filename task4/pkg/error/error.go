package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppError struct {
	Code    int    `json:"-"`
	ErrCode string `json:"error_code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func ThrowErr(c *gin.Context, appErr *AppError, message string) {
	if message != "" {
		appErr.Message = message
	}
	c.Error(appErr)
	c.Abort()
}

var (
	ErrSystem             = &AppError{http.StatusInternalServerError, "SYSTEM_ERROR", "系统异常,请稍后再试"}
	ErrUserNotFound       = &AppError{http.StatusNotFound, "USER_NOT_FOUND", "用户不存在"}
	ErrInvalidCredentials = &AppError{http.StatusUnauthorized, "INVALID_CREDENTIALS", "认证失败"}
	ErrUnauthorized       = &AppError{http.StatusForbidden, "UNAUTHORIZED", "权限不足"}
	ErrInvalidParams      = &AppError{http.StatusBadRequest, "INVALID_REQUEST", "请求参数错误"}
)
