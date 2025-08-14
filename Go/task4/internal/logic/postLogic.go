package logic

import (
	"task4/internal/model"
	"task4/pkg/db"
	error2 "task4/pkg/error"
)

type postLogic struct{}

var PostLogic = new(postLogic)

// 创建文章
func (p *postLogic) CreatePost(post *model.Post) error {
	return db.DB.Create(post).Error
}

// 分页查询文章
func (p *postLogic) PostPage(c *db.QueryParams, userId uint) (*db.PagedResult, error) {
	var posts []model.Post
	return db.Paginate(db.DB.Where("user_id = ?", userId), *c, &posts)
}

// 查询文章详情
func (p *postLogic) PostById(postId string, userId uint) (*model.Post, error) {
	post := model.Post{}
	err := db.DB.Where("user_id = ?", userId).First(&post, postId).Error
	return &post, err
}

// 修改文章
func (p *postLogic) EditPost(post *model.Post, userId *uint) error {
	var count int64
	if err := db.DB.Model(model.Post{}).Where("id = ? and user_id = ?", post.ID, userId).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return error2.ErrUnauthorized
	}
	if err := db.DB.Updates(post).Error; err != nil {
		return err
	}
	return nil
}

// 删除文章
func (p *postLogic) DelPost(postId *string, userId *uint) error {
	var count int64
	if err := db.DB.Model(model.Post{}).Where("id = ? and user_id = ?", postId, userId).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return error2.ErrUnauthorized
	}
	if err := db.DB.Delete(&model.Post{}, postId).Error; err != nil {
		return err
	}
	return nil
}
