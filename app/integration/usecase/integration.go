package usecase

import (
	"context"
	"fmt"

	"github.com/yxrxy/AllergyBase/app/integration/domain/model"
	"github.com/yxrxy/AllergyBase/app/integration/domain/service"
)

// IntegrationUsecase 数据集成用例
type IntegrationUsecase struct {
	integrationService *service.IntegrationService
}

// NewIntegrationUsecase 创建数据集成用例
func NewIntegrationUsecase(integrationService *service.IntegrationService) *IntegrationUsecase {
	return &IntegrationUsecase{
		integrationService: integrationService,
	}
}

// GetPatientIntegratedView 获取患者综合视图
func (u *IntegrationUsecase) GetPatientIntegratedView(ctx context.Context, patientID int64) (*model.PatientIntegratedView, error) {
	if patientID <= 0 {
		return nil, fmt.Errorf("无效的患者ID")
	}

	return u.integrationService.GetPatientComprehensiveView(ctx, patientID)
}

// SearchPatients 搜索患者
func (u *IntegrationUsecase) SearchPatients(ctx context.Context, query *model.CrossDatabaseQuery) ([]*model.PatientIntegratedView, int64, error) {
	if query == nil {
		return nil, 0, fmt.Errorf("查询条件不能为空")
	}

	return u.integrationService.CrossDatabaseSearch(ctx, query)
}

// CheckPatientDataQuality 检查患者数据质量
func (u *IntegrationUsecase) CheckPatientDataQuality(ctx context.Context, patientID int64) (*model.DataQualityReport, error) {
	if patientID <= 0 {
		return nil, fmt.Errorf("无效的患者ID")
	}

	return u.integrationService.PerformDataQualityCheck(ctx, patientID)
}

// PerformCorrelationAnalysis 执行关联分析
func (u *IntegrationUsecase) PerformCorrelationAnalysis(ctx context.Context, analysisType string, parameters map[string]interface{}) (*model.CorrelationAnalysis, error) {
	if analysisType == "" {
		return nil, fmt.Errorf("分析类型不能为空")
	}

	if parameters == nil {
		parameters = make(map[string]interface{})
	}

	return u.integrationService.PerformCorrelationAnalysis(ctx, analysisType, parameters)
}

// GetCohortInsights 获取队列洞察
func (u *IntegrationUsecase) GetCohortInsights(ctx context.Context, query *model.CrossDatabaseQuery) (map[string]interface{}, error) {
	if query == nil {
		return nil, fmt.Errorf("查询条件不能为空")
	}

	return u.integrationService.GetCohortInsights(ctx, query)
}
