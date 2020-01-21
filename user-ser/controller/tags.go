package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuwe1/recycle-shop/basic/common"
	"github.com/yuwe1/recycle-shop/user-ser/model"
	"github.com/yuwe1/recycle-shop/user-ser/service"
)

type TagRequest struct {
	Email        string   `json:"email"`
	School       string   `json:"school"`
	Interesttags []string `json:"interesttags"`
}

// 用户已经注册成功后提交标签
func UpdateTags(c *gin.Context) {
	id := c.Param("id")
	var tagrequest TagRequest
	c.BindJSON(&tagrequest)
	// 检测是否有空值

	if len(id) < 0 || len(tagrequest.Interesttags) <= 3 || len(tagrequest.Email) <= 0 {
		result := common.Result{
			Success:   0,
			Errorcode: 1,
			Message: &common.Data{
				Tip: "标签数量不足",
			},
		}
		c.JSON(200, result)
		return
	}
	tagservice := service.TagService{}
	ser := service.Service{}

	if tagservice.UpdateTages(id, tagrequest.Interesttags) {
		if len(tagrequest.School) > 0 {
			if ser.UpdateUserInfo(model.User{
				Email:  tagrequest.Email,
				School: tagrequest.School,
			}) {

				result := common.Result{
					Success:   0,
					Errorcode: 0,
				}
				c.JSON(200, result)
			} else {
				result := common.Result{
					Success:   0,
					Errorcode: 2,
					Message: &common.Data{
						Tip: "服务繁忙",
					},
				}
				c.JSON(200, result)
			}
		}
	} else {
		result := common.Result{
			Success:   0,
			Errorcode: 2,
			Message: &common.Data{
				Tip: "服务繁忙",
			},
		}
		c.JSON(200, result)
	}

}
