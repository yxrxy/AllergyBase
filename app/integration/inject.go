package integration

import (
	"gorm.io/gorm"

	"github.com/yxrxy/AllergyBase/app/integration/controllers/rpc"
	"github.com/yxrxy/AllergyBase/app/integration/domain/repository"
	"github.com/yxrxy/AllergyBase/app/integration/domain/service"
	"github.com/yxrxy/AllergyBase/app/integration/infrastructure/mysql"
	"github.com/yxrxy/AllergyBase/app/integration/usecase"
)

// IntegrationSet 数据集成模块依赖注入集合
type IntegrationSet struct {
	IntegrationRepo    repository.IntegrationRepository
	IntegrationService *service.IntegrationService
	IntegrationUsecase *usecase.IntegrationUsecase
	IntegrationRPC     *rpc.IntegrationRPCController
}

// NewIntegrationSet 创建数据集成模块依赖注入集合
func NewIntegrationSet(db *gorm.DB) *IntegrationSet {
	// 基础设施层
	integrationRepo := mysql.NewIntegrationMysqlRepository(db)

	// 领域服务层
	integrationService := service.NewIntegrationService(integrationRepo)

	// 应用用例层
	integrationUsecase := usecase.NewIntegrationUsecase(integrationService)

	// 控制器层
	integrationRPC := rpc.NewIntegrationRPCController(integrationUsecase)

	return &IntegrationSet{
		IntegrationRepo:    integrationRepo,
		IntegrationService: integrationService,
		IntegrationUsecase: integrationUsecase,
		IntegrationRPC:     integrationRPC,
	}
}
