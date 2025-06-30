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
	"github.com/yxrxy/AllergyBase/app/clinical"
	"github.com/yxrxy/AllergyBase/config"
	"github.com/yxrxy/AllergyBase/kitex_gen/clinical/clinicalservice"
	"github.com/yxrxy/AllergyBase/pkg/base"
	"github.com/yxrxy/AllergyBase/pkg/middleware"
)

func init() {
	config.Init("clinical")
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		log.Fatalf("Clinical: new etcd registry failed, err: %v", err)
	}
	listenAddr := config.Clinical.RPCAddr
	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Clinical: resolve tcp addr failed, err: %v", err)
	}
	p := base.TelemetryProvider(config.Server.Name, config.Otel.CollectorAddr)
	defer func() {
		if err := p.Shutdown(context.Background()); err != nil {
			log.Fatalf("Clinical: shutdown telemetry provider failed, err: %v", err)
		}
	}()

	svr := clinicalservice.NewServer(
		// 注入依赖
		clinical.InjectClinicalHandler(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			// 配置文件的值是在运行时才能获取到的，Kitex 无法在编译时验证它是否正确
			ServiceName: "ClinicalService",
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
		log.Printf("Clinical: run server failed, err: %v", err)
		return
	}
}
