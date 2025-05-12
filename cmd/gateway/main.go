package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/yxrxy/AllergyBase/app/gateway/mw"
	"github.com/yxrxy/AllergyBase/app/gateway/router"
	"github.com/yxrxy/AllergyBase/app/gateway/rpc"
	"github.com/yxrxy/AllergyBase/config"
)

func init() {
	config.Init("gateway")
	rpc.Init()
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
