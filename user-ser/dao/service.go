package dao

import (
	"github.com/yuwe1/recycle-shop/basic/client/dbpool"
	"github.com/yuwe1/recycle-shop/basic/logger"
	"github.com/yuwe1/recycle-shop/user-ser/model"
)

func (d UserDao) UpdateUserInfo(user model.User) bool {
	session, err, p, c := dbpool.GetSession()
	defer func() {
		if session != nil {
			session.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	sql := "UPDATE USER set nickname = ?,passwd = ?," +
		"truename = ? ,sex = ?, email = ?,headerimage = ?," +
		"school = ?,signature = ?,birthday = ?,studentid = ? " +
		"WHERE id = ?"
	session.Begin()
	if _, err = session.Exec(sql, user.Nickname, user.Password, user.Truename, user.Sex,
		user.Email, user.HeaderImage, user.School, user.Signature,
		user.Birthday, user.StudentID, user.ID); err != nil {
		session.Rollback()
		return false
	} else {
		session.Commit()
	}

	return true
}
