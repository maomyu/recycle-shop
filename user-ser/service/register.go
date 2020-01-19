package service

import (
	"github.com/yuwe1/recycle-shop/user-ser/dao"
	"github.com/yuwe1/recycle-shop/user-ser/model"
)

type UserService struct {
}

func (u UserService) GetUserInfo(email string) (model.User, error) {

	userdao := dao.UserDao{}
	user := userdao.QueryUserByEmail(email)
	return user, nil
}
