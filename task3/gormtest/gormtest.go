package gormtest

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
type User struct {
	gorm.Model
	ID       uint
	Name     string
	PostSize uint
	Posts    []Post
}
type Post struct {
	gorm.Model
	ID            uint
	Title         string
	Content       string
	UserID        uint
	Comments      []Comment
	CommentStatus string
}
type Comment struct {
	gorm.Model
	ID      uint
	Content string
	PostID  uint
}

var db *gorm.DB

func init() {
	const url = "root:Xm@123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("connect DB failed")
	}
}

func InitTables() {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
func InitData() {
	user := User{Name: "张三"}
	post1 := Post{Title: "第一篇文章", Content: "内容A", CommentStatus: "有评论"}
	post2 := Post{Title: "第二篇文章", Content: "内容B", CommentStatus: "有评论"}
	comment1 := Comment{Content: "评论1"}
	comment2 := Comment{Content: "评论2"}
	comment3 := Comment{Content: "评论3"}

	post1.Comments = append(post1.Comments, comment1, comment2)
	post2.Comments = append(post2.Comments, comment3)
	user.Posts = append(user.Posts, post1, post2)
	db.Debug().Create(&user)
}

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func GetPostByUserId(userId uint) User {
	var user User
	db.Debug().
		Model(&User{ID: userId}).
		Preload("Posts").
		Preload("Posts.Comments").
		First(&user)
	return user
}

func GetPostWithMostComments() Post {
	var post Post
	//找到最多评论文章id
	most := db.Debug().
		Model(&Comment{}).
		Select("post_id, count(*) as comment_count").
		Group("post_id").
		Order("comment_count desc").
		Limit(1)
	//关联子查询，查询出文章信息
	db.Debug().
		Model(&Post{}).
		Select("posts.*").
		Joins("JOIN (?) most ON posts.id = most.post_id", most).
		First(&post)
	db.Debug().Model(&Comment{PostID: post.ID}).Find(&post.Comments)
	return post
}

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
