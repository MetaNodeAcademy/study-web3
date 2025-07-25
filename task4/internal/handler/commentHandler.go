package handler

import (
	"github.com/gin-gonic/gin"
	"task4/internal/logic"
	"task4/internal/model"
	error2 "task4/pkg/error"
	"task4/pkg/response"
)

// 创建评论
func CreateComment(c *gin.Context) {
	comment := model.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	userId := c.MustGet("userID").(uint)
	err := logic.CommentLogic.CreateComment(&comment, &userId)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "创建评论失败")
		return
	}
	response.Success(c, nil, "创建评论成功")
}

// 查询某篇文章的所有评论
func CommentByPostId(c *gin.Context) {
	postId, exist := c.GetQuery("postId")
	if !exist {
		response.Error(c, error2.ErrInvalidParams)
	}
	list, err := logic.CommentLogic.CommentByPostId(&postId)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "查询文章评论失败")
		return
	}
	response.Success(c, list, "查询文章评论成功")
}
