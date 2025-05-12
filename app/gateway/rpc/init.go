package rpc

import (
	"github.com/yxrxy/AllergyBase/kitex_gen/user/userservice"
)

var (
	userClient userservice.Client
)

// Init 初始化所有RPC客户端
func Init() {
	InitUserRPC()
}
