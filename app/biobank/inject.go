package biobank

import (
	"github.com/yxrxy/AllergyBase/app/biobank/controllers/rpc"
	"github.com/yxrxy/AllergyBase/app/biobank/domain/service"
	"github.com/yxrxy/AllergyBase/app/biobank/infrastructure/cache"
	"github.com/yxrxy/AllergyBase/app/biobank/infrastructure/mysql"
	"github.com/yxrxy/AllergyBase/app/biobank/usecase"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
)

func InjectBiobankHandler() *rpc.BiobankHandler {
	// 初始化 MySQL 和 Redis 客户端
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}

	redisClient, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}

	// 创建基础设施层
	biobankDB := mysql.NewBiobankDB(gormDB)
	biobankCache := cache.NewBiobankCache(redisClient)

	// 创建 domain service
	biobankService := service.NewBiobankService(biobankDB, biobankCache)

	// 创建 usecase
	biobankUseCase := usecase.NewBiobankCase(biobankDB, biobankService)

	// 创建并返回 handler
	return rpc.NewBiobankHandler(biobankUseCase)
}
