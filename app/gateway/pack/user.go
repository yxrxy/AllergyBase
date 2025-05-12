package pack

import (
	"github.com/yxrxy/AllergyBase/app/gateway/model/model"
	rpcModel "github.com/yxrxy/AllergyBase/kitex_gen/model"
)

// BuildUserInfo 将 RPC 交流实体转换成 http 返回的实体
func BuildUserInfo(u *rpcModel.User) *model.User {
	return &model.User{
		ID:       u.Id,
		Username: u.Username,
		Avatar:   u.Avatar,
	}
}
