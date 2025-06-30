package rpc

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/integration/domain/model"
	"github.com/yxrxy/AllergyBase/app/integration/usecase"
)

// IntegrationRPCController 数据集成RPC控制器
type IntegrationRPCController struct {
	integrationUsecase *usecase.IntegrationUsecase
}

// NewIntegrationRPCController 创建数据集成RPC控制器
func NewIntegrationRPCController(integrationUsecase *usecase.IntegrationUsecase) *IntegrationRPCController {
	return &IntegrationRPCController{
		integrationUsecase: integrationUsecase,
	}
}

// GetPatientIntegratedView 获取患者综合视图
func (c *IntegrationRPCController) GetPatientIntegratedView(ctx context.Context, patientID int64) (*model.PatientIntegratedView, error) {
	return c.integrationUsecase.GetPatientIntegratedView(ctx, patientID)
}

// SearchPatients 搜索患者
func (c *IntegrationRPCController) SearchPatients(ctx context.Context, query *model.CrossDatabaseQuery) ([]*model.PatientIntegratedView, int64, error) {
	return c.integrationUsecase.SearchPatients(ctx, query)
}

// CheckPatientDataQuality 检查患者数据质量
func (c *IntegrationRPCController) CheckPatientDataQuality(ctx context.Context, patientID int64) (*model.DataQualityReport, error) {
	return c.integrationUsecase.CheckPatientDataQuality(ctx, patientID)
}

// PerformCorrelationAnalysis 执行关联分析
func (c *IntegrationRPCController) PerformCorrelationAnalysis(ctx context.Context, analysisType string, parameters map[string]interface{}) (*model.CorrelationAnalysis, error) {
	return c.integrationUsecase.PerformCorrelationAnalysis(ctx, analysisType, parameters)
}

// GetCohortInsights 获取队列洞察
func (c *IntegrationRPCController) GetCohortInsights(ctx context.Context, query *model.CrossDatabaseQuery) (map[string]interface{}, error) {
	return c.integrationUsecase.GetCohortInsights(ctx, query)
}
