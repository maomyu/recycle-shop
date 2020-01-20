package service

import (
	"github.com/yuwe1/recycle-shop/user-ser/dao"
	"github.com/yuwe1/recycle-shop/user-ser/model"
)

func (l LoginService) Login(email string, password string) (b bool, r model.LoginResult) {
	// 根据用户id查出用户的信息
	userService := Service{}
	user, _ := userService.GetUserInfo(email)
	if len(user.ID) <= 0 {
		return false, r
	}
	userdao := dao.UserDao{}
	// 设置用户的在线状态status
	userdao.SetOnlineStatus(user.ID, 1)
	// 封装结果
	isrealname := 0
	if len(user.Truename) > 0 {
		isrealname = 1
	}
	// 获得所关注的人的数量
	follownum := userdao.GetFollowNum(user.ID)
	// 获得粉丝的数量
	fansnum := userdao.GetFansNum(user.ID)
	// 获得信用值
	score := userdao.GetCreditScore(user.ID)
	// 转换日期
	t := user.Birthday

	r = model.LoginResult{
		ID:          user.ID,
		Nickname:    user.Nickname,
		Sex:         user.Sex,
		Email:       user.Email,
		HeaderImage: user.Email,
		School:      user.School,
		Signature:   user.Signature,
		Birthday:    t.Format("2006-01-02"),
		StudentID:   user.StudentID,
		Isrealname:  isrealname,
		Follow:      follownum,
		Fans:        fansnum,
		Creditscore: score,
	}
	//
	return true, r
}
