package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuwe1/recycle-shop/basic/logger"
	"github.com/yuwe1/recycle-shop/user-ser/service"
)

type loginRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 注册
func Register(c *gin.Context) {

	parm := new(loginRequest)
	err := c.BindJSON(parm)
	defer func() {
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	if err != nil {
		fmt.Errorf("[user-ser]:[register]:[%w]", err)
	}

	//判读该账户是否可以注册
	usersrv := &service.UserService{}
	if user, _ := usersrv.GetUserInfo(parm.Email); len(user.ID) > 0 {
		fmt.Println(user)
	} else {
		fmt.Println("可以注册")
	}
}
