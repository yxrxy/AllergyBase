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
	"github.com/yxrxy/AllergyBase/app/biobank"
	"github.com/yxrxy/AllergyBase/config"
	"github.com/yxrxy/AllergyBase/kitex_gen/biobank/biobankservice"
	"github.com/yxrxy/AllergyBase/pkg/base"
	"github.com/yxrxy/AllergyBase/pkg/middleware"
)

func init() {
	config.Init("biobank")
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		log.Fatalf("Biobank: new etcd registry failed, err: %v", err)
	}
	listenAddr := config.Biobank.RPCAddr
	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Biobank: resolve tcp addr failed, err: %v", err)
	}
	p := base.TelemetryProvider(config.Server.Name, config.Otel.CollectorAddr)
	defer func() {
		if err := p.Shutdown(context.Background()); err != nil {
			log.Fatalf("Biobank: shutdown telemetry provider failed, err: %v", err)
		}
	}()

	svr := biobankservice.NewServer(
		// 注入依赖
		biobank.InjectBiobankHandler(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			// 配置文件的值是在运行时才能获取到的，Kitex 无法在编译时验证它是否正确
			ServiceName: "BiobankService",
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
		log.Printf("Biobank: run server failed, err: %v", err)
		return
	}
}
