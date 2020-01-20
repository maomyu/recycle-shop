package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuwe1/recycle-shop/basic/common"
	"github.com/yuwe1/recycle-shop/user-ser/service"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	id := c.Param("id")

	loginr := &loginRequest{}
	c.BindJSON(&loginr)

	// 判断id是否为null
	if len(id) <= 0 || len(loginr.Email) <= 0 || len(loginr.Password) < 0 {
		res := common.Result{
			Success:   0,
			Errorcode: 1,
			Message: common.Data{
				Tip: "不能出现空值",
			},
		}
		c.JSON(200, res)
	}
	// 根据id获取用户的基本信息
	// 如果没有该用户，说明未注册
	// 如果有用户,返回用户的基本信息
	loginService := service.LoginService{}
	if ok, data := loginService.Login(loginr.Email, loginr.Password); ok {
		res := common.Result{
			Success:   0,
			Errorcode: 0,
			Message:   &data,
		}
		c.JSON(200, res)
	} else {
		res := common.Result{
			Success:   0,
			Errorcode: 2,
			Message: common.Data{
				Tip: "未找到该账号",
			},
		}
		c.JSON(200, res)
	}

}
