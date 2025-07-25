package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"task4/internal/logic"
	"task4/internal/model"
	"task4/pkg/db"
	error2 "task4/pkg/error"
	"task4/pkg/response"
)

// 创建文章
func CreatePost(c *gin.Context) {
	params := model.Post{}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	userId := c.MustGet("userID").(uint)
	params.UserID = userId
	err := logic.PostLogic.CreatePost(&params)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "创建文章失败")
		return
	}
	response.Success(c, nil, "创建文章成功")
}

// 分页查询文章
func PostPage(c *gin.Context) {
	params := db.QueryParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	currentUserId := c.MustGet("userID").(uint)
	result, err := logic.PostLogic.PostPage(&params, currentUserId)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "分页查询文章失败")
		return
	}
	response.Success(c, result, "查询文章成功")
}

// 查询文章详情
func PostById(c *gin.Context) {
	postId, exist := c.GetQuery("postId")
	if !exist {
		response.Error(c, error2.ErrInvalidParams)
	}
	currentUserId := c.MustGet("userID").(uint)
	post, err := logic.PostLogic.PostById(postId, currentUserId)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "查询文章失败")
		return
	}
	response.Success(c, post, "查询文章成功")
}

// 修改文章
func EditPost(c *gin.Context) {
	params := model.Post{}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	currentUserId := c.MustGet("userID").(uint)
	err := logic.PostLogic.EditPost(&params, &currentUserId)
	if err != nil {
		if errors.Is(err, error2.ErrUnauthorized) {
			response.Error(c, error2.ErrUnauthorized)
		} else {
			response.Fail(c, error2.ErrSystem.Code, "更新文章失败")
		}
		return
	}
	response.Success(c, nil, "更新文章成功")
}

// 删除文章
func DelPost(c *gin.Context) {
	postId, exist := c.GetQuery("postId")
	if !exist {
		response.Error(c, error2.ErrInvalidParams)
	}
	currentUserId := c.MustGet("userID").(uint)
	err := logic.PostLogic.DelPost(&postId, &currentUserId)
	if err != nil {
		if errors.Is(err, error2.ErrUnauthorized) {
			response.Error(c, error2.ErrUnauthorized)
		} else {
			response.Fail(c, error2.ErrSystem.Code, "删除文章失败")
		}
		return
	}
	response.Success(c, nil, "删除文章成功")
}
