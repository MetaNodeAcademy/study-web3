package handler

import (
	"github.com/gin-gonic/gin"
	"task4/internal/logic"
	"task4/internal/model"
	"task4/pkg/response"
)

func Page(c *gin.Context) {
	params := model.UserPageReq{}
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Fail(c, 400, "无效的参数")
		return
	}
	userList, err := logic.UserLogic.Page(&params)
	if err != nil {
		response.Fail(c, 500, "分页查询用户失败")
	}
	response.Success(c, userList, "查询成功")
}

// 注册
func Register(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(c, 400, "无效的参数")
		return
	}
	if err := logic.UserLogic.Register(&user); err != nil {
		response.Fail(c, 500, "注册失败")
		return
	}
	response.Success(c, nil, "注册成功")
}

// 登录
func Login(c *gin.Context) {
	req := model.UserLoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "无效的参数")
		return
	}
	resp, err := logic.UserLogic.Login(&req)
	if err != nil {
		response.Fail(c, 500, "登录失败")
		return
	}
	response.Success(c, resp, "登录成功")

}
