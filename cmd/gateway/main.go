package main

import (
	"log"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/yxrxy/AllergyBase/app/gateway/handler/api/integration"
	"github.com/yxrxy/AllergyBase/app/gateway/mw"
	"github.com/yxrxy/AllergyBase/app/gateway/router"
	"github.com/yxrxy/AllergyBase/app/gateway/rpc"
	"github.com/yxrxy/AllergyBase/config"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
)

func init() {
	config.Init("gateway")
	rpc.Init()

	// 初始化数据库连接
	db, err := client.InitMySQL()
	if err != nil {
		log.Printf("数据库连接初始化失败: %v", err)
	} else {
		// 初始化集成服务的数据库连接
		integration.InitDB(db)
		log.Println("集成服务数据库连接初始化成功")
	}
}

func main() {
	listenAddr := config.Gateway.Addr

	h := server.New(
		server.WithHostPorts(listenAddr),
		server.WithHandleMethodNotAllowed(true),
	)

	h.NoHijackConnPool = true

	h.Use(
		mw.CORS(),
	)

	router.RegisterStaticRoutes(h)
	router.GeneratedRegister(h)

	h.Spin()
}
