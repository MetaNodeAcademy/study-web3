package logic

import (
	"golang.org/x/crypto/bcrypt"
	"task4/internal/model"
	"task4/pkg/auth"
	"task4/pkg/db"
)

type userLogic struct{}

var UserLogic = new(userLogic)

func (s *userLogic) Page(req *model.UserPageReq) (*db.PagedResult, error) {
	var users []model.User
	return db.Paginate(
		db.DB.Omit("password"),
		db.QueryParams{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		&users,
	)
}

// 用户注册
func (s *userLogic) Register(req *model.User) error {
	cryptPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(cryptPwd)
	if err := db.DB.Create(&req).Error; err != nil {
		return err
	}
	return nil
}

// 用户登录
func (s *userLogic) Login(req *model.UserLoginReq) (*model.UserLoginResp, error) {
	var user model.User
	//根据username查询用户信息
	if err := db.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}
	//生成token
	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &model.UserLoginResp{Token: token}, nil
}
