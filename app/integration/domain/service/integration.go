package service

import (
	"context"
	"fmt"
	"time"

	"github.com/yxrxy/AllergyBase/app/integration/domain/model"
	"github.com/yxrxy/AllergyBase/app/integration/domain/repository"
)

// IntegrationService 数据集成服务
type IntegrationService struct {
	integrationRepo repository.IntegrationRepository
}

// NewIntegrationService 创建数据集成服务
func NewIntegrationService(integrationRepo repository.IntegrationRepository) *IntegrationService {
	return &IntegrationService{
		integrationRepo: integrationRepo,
	}
}

// GetPatientComprehensiveView 获取患者综合视图
func (s *IntegrationService) GetPatientComprehensiveView(ctx context.Context, patientID int64) (*model.PatientIntegratedView, error) {
	view, err := s.integrationRepo.GetPatientIntegratedView(ctx, patientID)
	if err != nil {
		return nil, fmt.Errorf("获取患者综合视图失败: %w", err)
	}

	// 计算统计信息
	if view != nil {
		view.Statistics = s.calculatePatientStatistics(view)
	}

	return view, nil
}

// CrossDatabaseSearch 跨数据库搜索
func (s *IntegrationService) CrossDatabaseSearch(ctx context.Context, query *model.CrossDatabaseQuery) ([]*model.PatientIntegratedView, int64, error) {
	// 验证查询参数
	if err := s.validateQuery(query); err != nil {
		return nil, 0, fmt.Errorf("查询参数验证失败: %w", err)
	}

	// 执行跨数据库查询
	results, total, err := s.integrationRepo.CrossDatabaseQuery(ctx, query)
	if err != nil {
		return nil, 0, fmt.Errorf("跨数据库查询失败: %w", err)
	}

	// 为每个结果计算统计信息
	for _, result := range results {
		if result != nil {
			result.Statistics = s.calculatePatientStatistics(result)
		}
	}

	return results, total, nil
}

// PerformDataQualityCheck 执行数据质量检查
func (s *IntegrationService) PerformDataQualityCheck(ctx context.Context, patientID int64) (*model.DataQualityReport, error) {
	report, err := s.integrationRepo.CheckDataQuality(ctx, patientID)
	if err != nil {
		return nil, fmt.Errorf("数据质量检查失败: %w", err)
	}

	// 生成改进建议
	if report != nil {
		report.Recommendations = s.generateQualityRecommendations(report)
	}

	return report, nil
}

// PerformCorrelationAnalysis 执行关联分析
func (s *IntegrationService) PerformCorrelationAnalysis(ctx context.Context, analysisType string, parameters map[string]interface{}) (*model.CorrelationAnalysis, error) {
	// 验证分析参数
	if err := s.validateAnalysisParameters(analysisType, parameters); err != nil {
		return nil, fmt.Errorf("分析参数验证失败: %w", err)
	}

	// 执行关联分析
	analysis, err := s.integrationRepo.PerformCorrelationAnalysis(ctx, analysisType, parameters)
	if err != nil {
		return nil, fmt.Errorf("关联分析执行失败: %w", err)
	}

	// 解释分析结果
	if analysis != nil {
		s.interpretCorrelationResults(analysis)
	}

	return analysis, nil
}

// GetCohortInsights 获取队列洞察
func (s *IntegrationService) GetCohortInsights(ctx context.Context, query *model.CrossDatabaseQuery) (map[string]interface{}, error) {
	statistics, err := s.integrationRepo.GetCohortStatistics(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("获取队列统计失败: %w", err)
	}

	// 增强统计信息
	insights := s.enhanceStatistics(statistics)

	return insights, nil
}

// calculatePatientStatistics 计算患者统计信息
func (s *IntegrationService) calculatePatientStatistics(view *model.PatientIntegratedView) *model.PatientStatistics {
	stats := &model.PatientStatistics{
		TotalVisits:       len(view.Visits),
		TotalDiagnoses:    len(view.Diagnoses),
		TotalExaminations: len(view.Examinations),
		TotalFollowups:    len(view.FollowupRecords),
		TotalSamples:      len(view.Samples),
	}

	// 计算就诊日期范围
	if len(view.Visits) > 0 {
		stats.FirstVisitDate = view.Visits[0].VisitTime
		stats.LastVisitDate = view.Visits[len(view.Visits)-1].VisitTime
	}

	// 计算最后随访日期和下次随访日期
	if len(view.FollowupRecords) > 0 {
		lastRecord := view.FollowupRecords[len(view.FollowupRecords)-1]
		stats.LastFollowupDate = &lastRecord.FollowupDate
	}

	// 计算活跃用药数量
	activeCount := 0
	for _, med := range view.Medications {
		if med.EndDate == nil || med.EndDate.After(time.Now()) {
			activeCount++
		}
	}
	stats.ActiveMedications = activeCount

	return stats
}

// validateQuery 验证查询参数
func (s *IntegrationService) validateQuery(query *model.CrossDatabaseQuery) error {
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 || query.PageSize > 1000 {
		query.PageSize = 20
	}

	// 验证年龄范围
	if query.AgeRange != nil {
		if query.AgeRange.Min != nil && *query.AgeRange.Min < 0 {
			return fmt.Errorf("年龄范围最小值不能为负数")
		}
		if query.AgeRange.Max != nil && *query.AgeRange.Max < 0 {
			return fmt.Errorf("年龄范围最大值不能为负数")
		}
		if query.AgeRange.Min != nil && query.AgeRange.Max != nil && *query.AgeRange.Min > *query.AgeRange.Max {
			return fmt.Errorf("年龄范围最小值不能大于最大值")
		}
	}

	// 验证日期范围
	if query.DateRange != nil {
		if query.DateRange.StartDate != nil && query.DateRange.EndDate != nil && query.DateRange.StartDate.After(*query.DateRange.EndDate) {
			return fmt.Errorf("开始日期不能晚于结束日期")
		}
	}

	return nil
}

// generateQualityRecommendations 生成数据质量改进建议
func (s *IntegrationService) generateQualityRecommendations(report *model.DataQualityReport) []string {
	var recommendations []string

	// 基于整体评分生成建议
	if report.OverallScore < 60 {
		recommendations = append(recommendations, "数据质量较低，建议进行全面的数据清理和补充")
	} else if report.OverallScore < 80 {
		recommendations = append(recommendations, "数据质量中等，建议重点关注缺失字段的补充")
	}

	// 基于各模块评分生成具体建议
	if report.ClinicalScore < 70 {
		recommendations = append(recommendations, "临床数据不完整，建议补充患者基本信息和诊疗记录")
	}
	if report.EpidemiologyScore < 70 {
		recommendations = append(recommendations, "流调数据缺失，建议完善环境暴露和生活方式调查")
	}
	if report.FollowupScore < 70 {
		recommendations = append(recommendations, "随访数据不足，建议建立规范的随访计划和记录")
	}
	if report.BiobankScore < 70 {
		recommendations = append(recommendations, "生物样本数据缺失，建议收集相关样本和检测数据")
	}

	// 基于具体问题生成建议
	for _, issue := range report.Issues {
		if issue.Severity == "critical" {
			recommendations = append(recommendations, fmt.Sprintf("紧急处理：%s", issue.Suggestion))
		}
	}

	return recommendations
}

// validateAnalysisParameters 验证分析参数
func (s *IntegrationService) validateAnalysisParameters(analysisType string, parameters map[string]interface{}) error {
	validTypes := map[string]bool{
		"environmental_disease": true,
		"lifestyle_outcome":     true,
		"medication_effect":     true,
		"genetic_phenotype":     true,
	}

	if !validTypes[analysisType] {
		return fmt.Errorf("不支持的分析类型: %s", analysisType)
	}

	// 根据分析类型验证参数
	switch analysisType {
	case "environmental_disease":
		if _, ok := parameters["environmental_factors"]; !ok {
			return fmt.Errorf("环境-疾病关联分析需要环境因子参数")
		}
	case "lifestyle_outcome":
		if _, ok := parameters["lifestyle_factors"]; !ok {
			return fmt.Errorf("生活方式-结局关联分析需要生活方式因子参数")
		}
	}

	return nil
}

// interpretCorrelationResults 解释关联分析结果
func (s *IntegrationService) interpretCorrelationResults(analysis *model.CorrelationAnalysis) {
	for i := range analysis.Results {
		result := &analysis.Results[i]

		// 判断显著性
		if result.PValue < 0.001 {
			result.Significance = "highly_significant"
		} else if result.PValue < 0.01 {
			result.Significance = "very_significant"
		} else if result.PValue < 0.05 {
			result.Significance = "significant"
		} else {
			result.Significance = "not_significant"
		}
	}
}

// enhanceStatistics 增强统计信息
func (s *IntegrationService) enhanceStatistics(statistics map[string]interface{}) map[string]interface{} {
	enhanced := make(map[string]interface{})

	// 复制原始统计信息
	for k, v := range statistics {
		enhanced[k] = v
	}

	// 添加计算字段
	if totalPatients, ok := statistics["total_patients"].(int64); ok && totalPatients > 0 {
		if maleCount, ok := statistics["male_count"].(int64); ok {
			enhanced["male_percentage"] = float64(maleCount) / float64(totalPatients) * 100
		}
		if femaleCount, ok := statistics["female_count"].(int64); ok {
			enhanced["female_percentage"] = float64(femaleCount) / float64(totalPatients) * 100
		}
	}

	// 添加数据完整性指标
	if totalRecords, ok := statistics["total_records"].(int64); ok && totalRecords > 0 {
		if completeRecords, ok := statistics["complete_records"].(int64); ok {
			enhanced["completeness_rate"] = float64(completeRecords) / float64(totalRecords) * 100
		}
	}

	return enhanced
}
