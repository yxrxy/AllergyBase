package main

import (
	"context"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/yxrxy/AllergyBase/app/followup"
	"github.com/yxrxy/AllergyBase/config"
	"github.com/yxrxy/AllergyBase/kitex_gen/followup/followupservice"
	"github.com/yxrxy/AllergyBase/pkg/base"
	"github.com/yxrxy/AllergyBase/pkg/middleware"
)

func init() {
	config.Init("followup")
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		log.Fatalf("Followup: new etcd registry failed, err: %v", err)
	}
	listenAddr := config.Followup.RPCAddr
	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Followup: resolve tcp addr failed, err: %v", err)
	}
	p := base.TelemetryProvider(config.Server.Name, config.Otel.CollectorAddr)
	defer func() {
		if err := p.Shutdown(context.Background()); err != nil {
			log.Fatalf("Followup: shutdown telemetry provider failed, err: %v", err)
		}
	}()

	svr := followupservice.NewServer(
		// 注入依赖
		followup.InjectFollowupHandler(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			// 配置文件的值是在运行时才能获取到的，Kitex 无法在编译时验证它是否正确
			ServiceName: "FollowupService",
		}),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{
			MaxConnections: int(config.Server.MaxConnections),
			MaxQPS:         int(config.Server.MaxQPS),
		}),

		server.WithMiddleware(middleware.ErrorLog()),
		server.WithMiddleware(middleware.Respond()),
	)
	if err = svr.Run(); err != nil {
		log.Printf("Followup: run server failed, err: %v", err)
		return
	}
}
