package dao

import (
	"github.com/yuwe1/recycle-shop/basic/client/dbpool"
	"github.com/yuwe1/recycle-shop/basic/logger"
	"github.com/yuwe1/recycle-shop/user-ser/model"
)

type UserDao struct {
}

func (u UserDao) QueryUserByEmail(email string) (user model.User) {
	session, err, p, c := dbpool.GetSession()
	defer func() {
		if session != nil {
			session.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()

	sql := "select * from user where email = ?"
	row := session.QueryRow(sql, email)
	row.Scan(&user.ID, &user.Nickname, &user.Password, &user.Truename,
		&user.Sex, &user.Email, &user.HeaderImage, &user.School, &user.Signature,
		&user.Birthday, &user.StudentID,
	)
	return user
}
