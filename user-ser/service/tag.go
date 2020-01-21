package service

import "github.com/yuwe1/recycle-shop/user-ser/dao"

type TagService struct {
}

func (t TagService) UpdateTages(id string, ids []string) bool {
	// 更新标签
	tagdao := dao.UserDao{}
	return tagdao.UpdateUserTags(ids, id)
}
