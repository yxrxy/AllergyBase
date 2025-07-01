package repository

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/integration/domain/model"
)

// IntegrationRepository 数据集成存储库接口
type IntegrationRepository interface {
	// 患者综合视图
	GetPatientIntegratedView(ctx context.Context, patientID int64) (*model.PatientIntegratedView, error)
	GetPatientIntegratedViewBatch(ctx context.Context, patientIDs []int64) ([]*model.PatientIntegratedView, error)

	// 跨数据库查询
	CrossDatabaseQuery(ctx context.Context, query *model.CrossDatabaseQuery) ([]*model.PatientIntegratedView, int64, error)

	// 数据质量检查
	CheckDataQuality(ctx context.Context, patientID int64) (*model.DataQualityReport, error)
	BatchCheckDataQuality(ctx context.Context, patientIDs []int64) ([]*model.DataQualityReport, error)

	// 关联分析
	PerformCorrelationAnalysis(ctx context.Context, analysisType string, parameters map[string]interface{}) (*model.CorrelationAnalysis, error)
	GetCorrelationAnalysis(ctx context.Context, analysisID string) (*model.CorrelationAnalysis, error)
	ListCorrelationAnalyses(ctx context.Context, page, pageSize int) ([]*model.CorrelationAnalysis, int64, error)

	// 统计查询
	GetPatientStatistics(ctx context.Context, patientID int64) (*model.PatientStatistics, error)
	GetCohortStatistics(ctx context.Context, query *model.CrossDatabaseQuery) (map[string]interface{}, error)
}
