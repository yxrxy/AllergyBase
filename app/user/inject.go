package user

import (
	"github.com/yxrxy/AllergyBase/app/user/controllers/rpc"
	"github.com/yxrxy/AllergyBase/app/user/domain/service"
	"github.com/yxrxy/AllergyBase/app/user/infrastructure/cache"
	"github.com/yxrxy/AllergyBase/app/user/infrastructure/mysql"
	"github.com/yxrxy/AllergyBase/app/user/usecase"
	"github.com/yxrxy/AllergyBase/config"
	"github.com/yxrxy/AllergyBase/kitex_gen/user"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
)

func InjectUserHandler() user.UserService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}

	re, err := client.NewRedisClient(config.Redis.DB.User)
	if err != nil {
		panic(err)
	}

	db := mysql.NewUserDB(gormDB)
	redisCache := cache.NewUserCache(re)
	svc := service.NewUserService(db, redisCache)
	uc := usecase.NewUserCase(db, svc)

	return rpc.NewUserHandler(uc)
}
