package logic

import (
	"errors"
	"task4/internal/model"
	"task4/pkg/db"
)

type commentLogic struct{}

var CommentLogic = new(commentLogic)

func (comm *commentLogic) CreateComment(comment *model.Comment, userId *uint) error {
	var count int64
	if err := db.DB.Model(model.Post{}).Where("id = ?", comment.PostID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("文章不存在")
	}
	comment.UserID = *userId
	if err := db.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (comm *commentLogic) CommentByPostId(postId *string) (*[]model.Comment, error) {
	var commentList []model.Comment
	if err := db.DB.Model(model.Comment{}).Where("post_id = ?", postId).Find(&commentList).Error; err != nil {
		return nil, err
	}
	return &commentList, nil
}
