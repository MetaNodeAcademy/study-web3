package handler

import (
	"github.com/gin-gonic/gin"
	"task4/internal/logic"
	"task4/pkg/response"
)

type UserRegReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	user := UserRegReq{}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(c, 400, "无效的参数")
		return
	}
	if err := logic.UserLogic.Register(user.Username, user.Password); err != nil {
		response.Fail(c, 500, "注册失败")
		return
	}

	response.Success(c, nil, "注册成功")
}
