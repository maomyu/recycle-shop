package service

import (
	"encoding/json"
	"github.com/yuwe1/recycle-shop/basic/common"
	"github.com/yuwe1/recycle-shop/basic/mq"
	"github.com/yuwe1/recycle-shop/user-ser/dao"
	"github.com/yuwe1/recycle-shop/user-ser/model"
)

func (u Service) GetUserInfo(email string) (model.User, error) {

	userdao := dao.UserDao{}
	user := userdao.QueryUserByEmail(email)
	return user, nil
}
func (u Service) UpdateUserInfo(user model.User) bool {
	dao := dao.UserDao{}
	// 获得用户的基本信息
	school := user.School
	user = dao.QueryUserByEmail(user.Email)
	user.School = school
	return dao.UpdateUserInfo(user)
}

// 注册
func (u RegisterService) Register(user model.User) (model.RegisterResult, bool) {
	userdao := dao.UserDao{}
	if ok, _ := userdao.InsertUser(user); !ok {
		return model.RegisterResult{}, false
	}
	user = userdao.QueryUserByEmail(user.Email)
	// 设置该用户关注的官方账号
	// 获得所有的官方账号id
	ids := userdao.GetOfficialID()

	isrealname := 0
	if len(user.StudentID) > 0 {
		isrealname = 1
	}
	follow := "user:follow:" + user.ID
	userdao.UpdateFollow(ids, follow)
	// 使用rabbitmq消息队列通知redis增加某个用户的信用值，因为信用值在
	// 程序的更新比较频繁，故采用消息队列的形式
	byt, _ := json.Marshal(&common.CrediteScore{
		ID:    user.ID,
		Score: 10,
	})
	client := mq.GetRabbitMQ()
	client.PublishOnQueue(byt, "creditscore", "updatecreditscore", "creditscore:")
	result := model.RegisterResult{
		ID:          user.ID,
		Nickname:    user.Nickname,
		Sex:         user.Sex,
		Email:       user.Email,
		HeaderImage: user.HeaderImage,
		School:      user.School,
		Signature:   user.Signature,
		Birthday:    user.Birthday.Format("2006-01-02"),
		StudentID:   user.StudentID,
		Isrealname:  isrealname,
		Follow:      len(ids),
		Fans:        0,
		Creditscore: 10,
	}
	return result, true
}
