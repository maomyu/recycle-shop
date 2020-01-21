package dao

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/yuwe1/recycle-shop/basic/client/dbpool"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
	"github.com/yuwe1/recycle-shop/basic/common"
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

	sql := "SELECT * FROM USER WHERE email = ?"

	session.Begin()
	row := session.DB.QueryRow(sql, email)

	row.Scan(&user.ID, &user.Nickname, &user.Password,
		&user.Truename, &user.Sex,
		&user.Email,
		&user.HeaderImage, &user.School, &user.Signature,
		&user.Birthday, &user.StudentID, &user.Role,
	)
	session.Commit()
	return user
}

// 注册
func (d UserDao) InsertUser(user model.User) (bool, error) {
	session, err, p, c := dbpool.GetSession()
	defer func() {
		if session != nil {
			session.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()

	id := common.GetRandomID()

	sql := "INSERT INTO USER (id,nickname,passwd, email,birthday) VALUES(?,?,?,?,?)"
	session.Begin()
	if _, err := session.Exec(sql, id, user.Nickname, user.Password,
		user.Email, user.Birthday,
	); err != nil {
		logger.Sugar.Errorf("[user-ser:dao:insertuser]:[%w]", err)
		return false, err
	}
	session.Commit()
	return true, nil
}

// 获得所有的官方账号的ids
func (d UserDao) GetOfficialID() []string {
	session, err, p, c := dbpool.GetSession()
	defer func() {
		if session != nil {
			session.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()

	sql := "SELECT id FROM USER WHERE role = ?"
	ids := []string{}
	if rows, err := session.Query(sql, 1); err != nil {
		logger.Sugar.Error(err)
	} else {
		for rows.Next() {
			temp := ""
			rows.Scan(&temp)
			ids = append(ids, temp)
		}
	}
	return ids
}

// 更新关注的人
func (d UserDao) UpdateFollow(ids []string, key string) {
	f, err, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	conn := f.GetConn()
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("UpdateFollow [选择数据库]: [%w]", err)
	}
	conn.Do("MULTI")
	for _, v := range ids {
		if reply, _ := redis.Int(conn.Do("SADD", key, v)); reply <= 0 {
			fmt.Errorf("[user-ser]:[UpdateFollow]:[%w]", err)
		}
	}
	conn.Do("EXEC")
}
