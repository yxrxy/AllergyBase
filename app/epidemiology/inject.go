package epidemiology

import (
	"github.com/yxrxy/AllergyBase/app/epidemiology/controllers/rpc"
	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/service"
	"github.com/yxrxy/AllergyBase/app/epidemiology/infrastructure/cache"
	"github.com/yxrxy/AllergyBase/app/epidemiology/infrastructure/mysql"
	"github.com/yxrxy/AllergyBase/app/epidemiology/usecase"
	"github.com/yxrxy/AllergyBase/kitex_gen/epidemiology"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
)

func InjectEpidemiologyHandler() epidemiology.EpidemiologyService {
	// 初始化数据库连接
	mysqlClient, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}

	// 初始化Redis连接
	redisClient, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}

	// 创建基础设施层
	epidemiologyDB := mysql.NewEpidemiologyDB(mysqlClient)
	epidemiologyCache := cache.NewEpidemiologyCache(redisClient)

	// 创建领域服务层
	epidemiologyService := service.NewEpidemiologyService(epidemiologyDB, epidemiologyCache)

	// 创建用例层
	epidemiologyUseCase := usecase.NewEpidemiologyUseCase(epidemiologyDB, epidemiologyService)

	// 返回控制器
	return rpc.NewEpidemiologyHandler(epidemiologyUseCase)
}
