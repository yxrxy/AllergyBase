package clinical

import (
	"github.com/yxrxy/AllergyBase/app/clinical/controllers/rpc"
	"github.com/yxrxy/AllergyBase/app/clinical/domain/service"
	"github.com/yxrxy/AllergyBase/app/clinical/infrastructure/cache"
	"github.com/yxrxy/AllergyBase/app/clinical/infrastructure/mysql"
	"github.com/yxrxy/AllergyBase/app/clinical/usecase"
	"github.com/yxrxy/AllergyBase/kitex_gen/clinical"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
)

func InjectClinicalHandler() clinical.ClinicalService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}

	re, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}

	db := mysql.NewClinicalDB(gormDB)
	redisCache := cache.NewClinicalCache(re)
	svc := service.NewClinicalService(db, redisCache)
	uc := usecase.NewClinicalCase(db, svc)

	return rpc.NewClinicalHandler(uc)
}
