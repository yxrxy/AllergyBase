package followup

import (
	"github.com/yxrxy/AllergyBase/app/followup/controllers/rpc"
	"github.com/yxrxy/AllergyBase/app/followup/domain/service"
	"github.com/yxrxy/AllergyBase/app/followup/infrastructure/cache"
	"github.com/yxrxy/AllergyBase/app/followup/infrastructure/mysql"
	"github.com/yxrxy/AllergyBase/app/followup/usecase"
	"github.com/yxrxy/AllergyBase/kitex_gen/followup"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
)

func InjectFollowupHandler() followup.FollowupService {
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
	followupDB := mysql.NewFollowupDB(mysqlClient)
	followupCache := cache.NewFollowupCache(redisClient)

	// 创建领域服务层
	followupService := service.NewFollowupService(followupDB, followupCache)

	// 创建用例层
	followupUseCase := usecase.NewFollowupUseCase(followupDB, followupService)

	// 返回控制器
	return rpc.NewFollowupHandler(followupUseCase)
}
