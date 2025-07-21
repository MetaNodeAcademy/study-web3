package logic

import "fmt"

type userLogic struct{}

var UserLogic = new(userLogic)

// 用户注册
func (s *userLogic) Register(username, password string) error {
	// db.DB.Create(&model.User{})
	fmt.Println("userLogic===>Register")
	return nil
}
