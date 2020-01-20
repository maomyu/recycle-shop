package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuwe1/recycle-shop/basic/common"
	"github.com/yuwe1/recycle-shop/basic/logger"
	"github.com/yuwe1/recycle-shop/user-ser/model"
	"github.com/yuwe1/recycle-shop/user-ser/service"
)

type loginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
	if len(parm.Email) <= 0 || len(parm.Nickname) <= 0 || len(parm.Password) <= 0 {
		result := common.Result{
			Success:   0,
			Errorcode: 2,
			Message: &common.Data{
				Tip: "不能出现空值",
			},
		}
		c.JSON(200, result)
		return
	}
	if user, _ := usersrv.GetUserInfo(parm.Email); len(user.ID) > 0 {
		result := common.Result{
			Success:   0,
			Errorcode: 1,
			Message: &common.Data{
				Tip: "用户已经注册",
			},
		}
		c.JSON(200, result)
	} else {
		user := model.User{
			Nickname: parm.Nickname,
			Password: parm.Password,
			Email:    parm.Email,
			Birthday: time.Now(),
		}
		if r, ok := usersrv.Register(user); ok {
			result := common.Result{
				Success:   0,
				Errorcode: 0,
				Message:   &r,
			}
			c.JSON(200, result)

		} else {
			result := common.Result{
				Success:   0,
				Errorcode: 3,
				Message: &common.Data{
					Tip: "服务繁忙",
				},
			}
			c.JSON(200, result)
		}

	}
}
