package mysql

import (
	"context"
	"fmt"
	"math"
	"time"

	biobankModel "github.com/yxrxy/AllergyBase/app/biobank/domain/model"
	clinicalModel "github.com/yxrxy/AllergyBase/app/clinical/domain/model"
	epidemiologyModel "github.com/yxrxy/AllergyBase/app/epidemiology/domain/model"
	followupModel "github.com/yxrxy/AllergyBase/app/followup/domain/model"
	"github.com/yxrxy/AllergyBase/app/integration/domain/model"
	"github.com/yxrxy/AllergyBase/app/integration/domain/repository"
	"gorm.io/gorm"
)

// IntegrationMysqlRepository MySQL数据集成存储库实现
type IntegrationMysqlRepository struct {
	db *gorm.DB
}

// NewIntegrationMysqlRepository 创建MySQL数据集成存储库
func NewIntegrationMysqlRepository(db *gorm.DB) repository.IntegrationRepository {
	return &IntegrationMysqlRepository{
		db: db,
	}
}

// GetPatientIntegratedView 获取患者综合视图
func (r *IntegrationMysqlRepository) GetPatientIntegratedView(ctx context.Context, patientID int64) (*model.PatientIntegratedView, error) {
	view := &model.PatientIntegratedView{}

	// 获取患者基本信息
	var patient clinicalModel.Patient
	if err := r.db.WithContext(ctx).First(&patient, patientID).Error; err != nil {
		return nil, fmt.Errorf("获取患者信息失败: %w", err)
	}
	view.Patient = &patient

	// 获取临床数据
	if err := r.loadClinicalData(ctx, patientID, view); err != nil {
		return nil, fmt.Errorf("加载临床数据失败: %w", err)
	}

	// 获取流调数据
	if err := r.loadEpidemiologyData(ctx, patientID, view); err != nil {
		return nil, fmt.Errorf("加载流调数据失败: %w", err)
	}

	// 获取随访数据
	if err := r.loadFollowupData(ctx, patientID, view); err != nil {
		return nil, fmt.Errorf("加载随访数据失败: %w", err)
	}

	// 获取生物样本数据
	if err := r.loadBiobankData(ctx, patientID, view); err != nil {
		return nil, fmt.Errorf("加载生物样本数据失败: %w", err)
	}

	return view, nil
}

// GetPatientIntegratedViewBatch 批量获取患者综合视图
func (r *IntegrationMysqlRepository) GetPatientIntegratedViewBatch(ctx context.Context, patientIDs []int64) ([]*model.PatientIntegratedView, error) {
	var views []*model.PatientIntegratedView

	for _, id := range patientIDs {
		view, err := r.GetPatientIntegratedView(ctx, id)
		if err != nil {
			continue // 跳过错误的记录，继续处理其他记录
		}
		views = append(views, view)
	}

	return views, nil
}

// CrossDatabaseQuery 跨数据库查询
func (r *IntegrationMysqlRepository) CrossDatabaseQuery(ctx context.Context, query *model.CrossDatabaseQuery) ([]*model.PatientIntegratedView, int64, error) {
	// 构建基础查询
	baseQuery := r.db.WithContext(ctx).Model(&clinicalModel.Patient{})

	// 应用筛选条件
	baseQuery = r.applyQueryFilters(baseQuery, query)

	// 获取总数
	var total int64
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("统计查询结果失败: %w", err)
	}

	// 获取患者ID列表
	var patientIDs []int64
	offset := (query.Page - 1) * query.PageSize
	if err := baseQuery.Select("id").Offset(offset).Limit(query.PageSize).Pluck("id", &patientIDs).Error; err != nil {
		return nil, 0, fmt.Errorf("查询患者ID失败: %w", err)
	}

	// 批量获取综合视图
	views, err := r.GetPatientIntegratedViewBatch(ctx, patientIDs)
	if err != nil {
		return nil, 0, fmt.Errorf("批量获取患者视图失败: %w", err)
	}

	return views, total, nil
}

// CheckDataQuality 检查数据质量
func (r *IntegrationMysqlRepository) CheckDataQuality(ctx context.Context, patientID int64) (*model.DataQualityReport, error) {
	report := &model.DataQualityReport{
		PatientID: patientID,
		CheckTime: time.Now(),
		Issues:    []model.DataQualityIssue{},
	}

	// 检查临床数据质量
	clinicalScore, clinicalIssues := r.checkClinicalDataQuality(ctx, patientID)
	report.ClinicalScore = clinicalScore
	report.Issues = append(report.Issues, clinicalIssues...)

	// 检查流调数据质量
	epiScore, epiIssues := r.checkEpidemiologyDataQuality(ctx, patientID)
	report.EpidemiologyScore = epiScore
	report.Issues = append(report.Issues, epiIssues...)

	// 检查随访数据质量
	followupScore, followupIssues := r.checkFollowupDataQuality(ctx, patientID)
	report.FollowupScore = followupScore
	report.Issues = append(report.Issues, followupIssues...)

	// 检查生物样本数据质量
	biobankScore, biobankIssues := r.checkBiobankDataQuality(ctx, patientID)
	report.BiobankScore = biobankScore
	report.Issues = append(report.Issues, biobankIssues...)

	// 计算总体评分
	report.OverallScore = (clinicalScore + epiScore + followupScore + biobankScore) / 4

	return report, nil
}

// BatchCheckDataQuality 批量检查数据质量
func (r *IntegrationMysqlRepository) BatchCheckDataQuality(ctx context.Context, patientIDs []int64) ([]*model.DataQualityReport, error) {
	var reports []*model.DataQualityReport

	for _, id := range patientIDs {
		report, err := r.CheckDataQuality(ctx, id)
		if err != nil {
			continue // 跳过错误的记录
		}
		reports = append(reports, report)
	}

	return reports, nil
}

// PerformCorrelationAnalysis 执行关联分析
func (r *IntegrationMysqlRepository) PerformCorrelationAnalysis(ctx context.Context, analysisType string, parameters map[string]interface{}) (*model.CorrelationAnalysis, error) {
	analysis := &model.CorrelationAnalysis{
		AnalysisID:   fmt.Sprintf("%s_%d", analysisType, time.Now().Unix()),
		AnalysisType: analysisType,
		CreatedAt:    time.Now(),
		Parameters:   parameters,
		Results:      []model.CorrelationResult{},
	}

	// 根据分析类型执行不同的分析
	switch analysisType {
	case "environmental_disease":
		results, err := r.analyzeEnvironmentalDiseaseCorrelation(ctx, parameters)
		if err != nil {
			return nil, err
		}
		analysis.Results = results
	case "lifestyle_outcome":
		results, err := r.analyzeLifestyleOutcomeCorrelation(ctx, parameters)
		if err != nil {
			return nil, err
		}
		analysis.Results = results
	default:
		return nil, fmt.Errorf("不支持的分析类型: %s", analysisType)
	}

	return analysis, nil
}

// GetCorrelationAnalysis 获取关联分析结果
func (r *IntegrationMysqlRepository) GetCorrelationAnalysis(ctx context.Context, analysisID string) (*model.CorrelationAnalysis, error) {
	// 这里应该从缓存或数据库中获取分析结果
	// 简化实现，返回空结果
	return nil, fmt.Errorf("分析结果不存在: %s", analysisID)
}

// ListCorrelationAnalyses 列出关联分析
func (r *IntegrationMysqlRepository) ListCorrelationAnalyses(ctx context.Context, page, pageSize int) ([]*model.CorrelationAnalysis, int64, error) {
	// 简化实现，返回空列表
	return []*model.CorrelationAnalysis{}, 0, nil
}

// GetPatientStatistics 获取患者统计信息
func (r *IntegrationMysqlRepository) GetPatientStatistics(ctx context.Context, patientID int64) (*model.PatientStatistics, error) {
	stats := &model.PatientStatistics{}

	// 统计就诊次数
	var totalVisits int64
	r.db.WithContext(ctx).Model(&clinicalModel.Visit{}).Where("patient_id = ?", patientID).Count(&totalVisits)
	stats.TotalVisits = int(totalVisits)

	// 统计诊断次数
	var totalDiagnoses int64
	r.db.WithContext(ctx).Model(&clinicalModel.Diagnosis{}).
		Joins("JOIN clinical_visit_record ON clinical_diagnosis.visit_id = clinical_visit_record.id").
		Where("clinical_visit_record.patient_id = ?", patientID).Count(&totalDiagnoses)
	stats.TotalDiagnoses = int(totalDiagnoses)

	// 统计检查次数
	var totalExaminations int64
	r.db.WithContext(ctx).Model(&clinicalModel.Examination{}).
		Joins("JOIN clinical_visit_record ON clinical_examination.visit_id = clinical_visit_record.id").
		Where("clinical_visit_record.patient_id = ?", patientID).Count(&totalExaminations)
	stats.TotalExaminations = int(totalExaminations)

	// 统计随访次数
	var totalFollowups int64
	r.db.WithContext(ctx).Model(&followupModel.FollowupRecord{}).Where("patient_id = ?", patientID).Count(&totalFollowups)
	stats.TotalFollowups = int(totalFollowups)

	// 统计样本数量
	var totalSamples int64
	r.db.WithContext(ctx).Model(&biobankModel.Sample{}).Where("patient_id = ?", patientID).Count(&totalSamples)
	stats.TotalSamples = int(totalSamples)

	return stats, nil
}

// GetCohortStatistics 获取队列统计信息
func (r *IntegrationMysqlRepository) GetCohortStatistics(ctx context.Context, query *model.CrossDatabaseQuery) (map[string]interface{}, error) {
	statistics := make(map[string]interface{})

	baseQuery := r.db.WithContext(ctx).Model(&clinicalModel.Patient{})
	baseQuery = r.applyQueryFilters(baseQuery, query)

	// 总患者数
	var totalPatients int64
	baseQuery.Count(&totalPatients)
	statistics["total_patients"] = totalPatients

	// 性别分布
	var genderStats []struct {
		Gender string
		Count  int64
	}
	baseQuery.Select("gender, COUNT(*) as count").Group("gender").Scan(&genderStats)
	for _, stat := range genderStats {
		if stat.Gender == "M" {
			statistics["male_count"] = stat.Count
		} else if stat.Gender == "F" {
			statistics["female_count"] = stat.Count
		}
	}

	// 年龄分布
	var ageStats struct {
		AvgAge float64
		MinAge int
		MaxAge int
	}
	baseQuery.Select("AVG(YEAR(CURDATE()) - YEAR(birth_date)) as avg_age, MIN(YEAR(CURDATE()) - YEAR(birth_date)) as min_age, MAX(YEAR(CURDATE()) - YEAR(birth_date)) as max_age").
		Scan(&ageStats)
	statistics["average_age"] = ageStats.AvgAge
	statistics["min_age"] = ageStats.MinAge
	statistics["max_age"] = ageStats.MaxAge

	return statistics, nil
}

// 辅助方法：加载临床数据
func (r *IntegrationMysqlRepository) loadClinicalData(ctx context.Context, patientID int64, view *model.PatientIntegratedView) error {
	// 加载就诊记录
	if err := r.db.WithContext(ctx).Where("patient_id = ?", patientID).Find(&view.Visits).Error; err != nil {
		return err
	}

	// 加载诊断记录
	if err := r.db.WithContext(ctx).
		Joins("JOIN clinical_visit_record ON clinical_diagnosis.visit_id = clinical_visit_record.id").
		Where("clinical_visit_record.patient_id = ?", patientID).
		Find(&view.Diagnoses).Error; err != nil {
		return err
	}

	// 加载检查记录
	if err := r.db.WithContext(ctx).
		Joins("JOIN clinical_visit_record ON clinical_examination.visit_id = clinical_visit_record.id").
		Where("clinical_visit_record.patient_id = ?", patientID).
		Find(&view.Examinations).Error; err != nil {
		return err
	}

	return nil
}

// 辅助方法：加载流调数据
func (r *IntegrationMysqlRepository) loadEpidemiologyData(ctx context.Context, patientID int64, view *model.PatientIntegratedView) error {
	// 加载环境暴露记录
	if err := r.db.WithContext(ctx).Where("patient_id = ?", patientID).Find(&view.EnvironmentExposures).Error; err != nil {
		return err
	}

	// 加载生活方式调查
	if err := r.db.WithContext(ctx).Where("patient_id = ?", patientID).Find(&view.LifestyleSurveys).Error; err != nil {
		return err
	}

	return nil
}

// 辅助方法：加载随访数据
func (r *IntegrationMysqlRepository) loadFollowupData(ctx context.Context, patientID int64, view *model.PatientIntegratedView) error {
	// 加载随访计划
	if err := r.db.WithContext(ctx).Where("patient_id = ?", patientID).Find(&view.FollowupPlans).Error; err != nil {
		return err
	}

	// 加载随访记录
	if err := r.db.WithContext(ctx).Where("patient_id = ?", patientID).Find(&view.FollowupRecords).Error; err != nil {
		return err
	}

	// 加载用药监测
	if err := r.db.WithContext(ctx).
		Joins("JOIN followup_record ON medication_monitor.followup_id = followup_record.id").
		Where("followup_record.patient_id = ?", patientID).
		Find(&view.Medications).Error; err != nil {
		return err
	}

	return nil
}

// 辅助方法：加载生物样本数据
func (r *IntegrationMysqlRepository) loadBiobankData(ctx context.Context, patientID int64, view *model.PatientIntegratedView) error {
	// 加载样本信息
	if err := r.db.WithContext(ctx).Where("patient_id = ?", patientID).Find(&view.Samples).Error; err != nil {
		return err
	}

	// 加载基因组数据
	if err := r.db.WithContext(ctx).
		Joins("JOIN sample ON genomic_data.sample_id = sample.id").
		Where("sample.patient_id = ?", patientID).
		Find(&view.GenomicData).Error; err != nil {
		return err
	}

	// 加载蛋白组学数据
	if err := r.db.WithContext(ctx).
		Joins("JOIN sample ON proteomics_data.sample_id = sample.id").
		Where("sample.patient_id = ?", patientID).
		Find(&view.ProteomicsData).Error; err != nil {
		return err
	}

	return nil
}

// 辅助方法：应用查询过滤器
func (r *IntegrationMysqlRepository) applyQueryFilters(query *gorm.DB, filter *model.CrossDatabaseQuery) *gorm.DB {
	// 患者ID过滤
	if len(filter.PatientIDs) > 0 {
		query = query.Where("id IN ?", filter.PatientIDs)
	}

	// 性别过滤
	if filter.Gender != nil {
		query = query.Where("gender = ?", *filter.Gender)
	}

	// 年龄范围过滤
	if filter.AgeRange != nil {
		if filter.AgeRange.Min != nil {
			query = query.Where("YEAR(CURDATE()) - YEAR(birth_date) >= ?", *filter.AgeRange.Min)
		}
		if filter.AgeRange.Max != nil {
			query = query.Where("YEAR(CURDATE()) - YEAR(birth_date) <= ?", *filter.AgeRange.Max)
		}
	}

	// 日期范围过滤
	if filter.DateRange != nil {
		if filter.DateRange.StartDate != nil {
			query = query.Where("created_at >= ?", *filter.DateRange.StartDate)
		}
		if filter.DateRange.EndDate != nil {
			query = query.Where("created_at <= ?", *filter.DateRange.EndDate)
		}
	}

	return query
}

// 辅助方法：检查临床数据质量
func (r *IntegrationMysqlRepository) checkClinicalDataQuality(ctx context.Context, patientID int64) (float64, []model.DataQualityIssue) {
	var issues []model.DataQualityIssue
	score := 100.0

	// 检查患者基本信息完整性
	var patient clinicalModel.Patient
	if err := r.db.WithContext(ctx).First(&patient, patientID).Error; err != nil {
		issues = append(issues, model.DataQualityIssue{
			Category:    "clinical",
			Severity:    "critical",
			Field:       "patient_info",
			Description: "患者基本信息缺失",
			Suggestion:  "补充患者基本信息",
		})
		score -= 50
	} else {
		// 检查关键字段
		if patient.Name == "" {
			issues = append(issues, model.DataQualityIssue{
				Category:    "clinical",
				Severity:    "high",
				Field:       "name",
				Description: "患者姓名缺失",
				Suggestion:  "补充患者姓名",
			})
			score -= 10
		}
		if patient.BirthDate == nil {
			issues = append(issues, model.DataQualityIssue{
				Category:    "clinical",
				Severity:    "medium",
				Field:       "birth_date",
				Description: "出生日期缺失",
				Suggestion:  "补充出生日期",
			})
			score -= 5
		}
	}

	return math.Max(0, score), issues
}

// 辅助方法：检查流调数据质量
func (r *IntegrationMysqlRepository) checkEpidemiologyDataQuality(ctx context.Context, patientID int64) (float64, []model.DataQualityIssue) {
	var issues []model.DataQualityIssue
	score := 100.0

	// 检查环境暴露数据
	var exposureCount int64
	r.db.WithContext(ctx).Model(&epidemiologyModel.EnvironmentExposure{}).Where("patient_id = ?", patientID).Count(&exposureCount)
	if exposureCount == 0 {
		issues = append(issues, model.DataQualityIssue{
			Category:    "epidemiology",
			Severity:    "medium",
			Field:       "environment_exposure",
			Description: "环境暴露数据缺失",
			Suggestion:  "收集环境暴露信息",
		})
		score -= 30
	}

	// 检查生活方式调查
	var lifestyleCount int64
	r.db.WithContext(ctx).Model(&epidemiologyModel.LifestyleSurvey{}).Where("patient_id = ?", patientID).Count(&lifestyleCount)
	if lifestyleCount == 0 {
		issues = append(issues, model.DataQualityIssue{
			Category:    "epidemiology",
			Severity:    "medium",
			Field:       "lifestyle_survey",
			Description: "生活方式调查数据缺失",
			Suggestion:  "完成生活方式调查",
		})
		score -= 30
	}

	return math.Max(0, score), issues
}

// 辅助方法：检查随访数据质量
func (r *IntegrationMysqlRepository) checkFollowupDataQuality(ctx context.Context, patientID int64) (float64, []model.DataQualityIssue) {
	var issues []model.DataQualityIssue
	score := 100.0

	// 检查随访计划
	var planCount int64
	r.db.WithContext(ctx).Model(&followupModel.FollowupPlan{}).Where("patient_id = ?", patientID).Count(&planCount)
	if planCount == 0 {
		issues = append(issues, model.DataQualityIssue{
			Category:    "followup",
			Severity:    "high",
			Field:       "followup_plan",
			Description: "随访计划缺失",
			Suggestion:  "制定随访计划",
		})
		score -= 40
	}

	return math.Max(0, score), issues
}

// 辅助方法：检查生物样本数据质量
func (r *IntegrationMysqlRepository) checkBiobankDataQuality(ctx context.Context, patientID int64) (float64, []model.DataQualityIssue) {
	var issues []model.DataQualityIssue
	score := 100.0

	// 检查样本数据
	var sampleCount int64
	r.db.WithContext(ctx).Model(&biobankModel.Sample{}).Where("patient_id = ?", patientID).Count(&sampleCount)
	if sampleCount == 0 {
		issues = append(issues, model.DataQualityIssue{
			Category:    "biobank",
			Severity:    "low",
			Field:       "sample",
			Description: "生物样本数据缺失",
			Suggestion:  "收集生物样本",
		})
		score -= 50
	}

	return math.Max(0, score), issues
}

// 辅助方法：分析环境-疾病关联
func (r *IntegrationMysqlRepository) analyzeEnvironmentalDiseaseCorrelation(ctx context.Context, parameters map[string]interface{}) ([]model.CorrelationResult, error) {
	// 简化实现，返回模拟结果
	results := []model.CorrelationResult{
		{
			Factor1:         "PM2.5暴露",
			Factor2:         "过敏性哮喘",
			CorrelationType: "pearson",
			Coefficient:     0.65,
			PValue:          0.001,
			SampleSize:      150,
		},
		{
			Factor1:         "花粉浓度",
			Factor2:         "过敏性鼻炎",
			CorrelationType: "pearson",
			Coefficient:     0.78,
			PValue:          0.0001,
			SampleSize:      120,
		},
	}

	return results, nil
}

// 辅助方法：分析生活方式-结局关联
func (r *IntegrationMysqlRepository) analyzeLifestyleOutcomeCorrelation(ctx context.Context, parameters map[string]interface{}) ([]model.CorrelationResult, error) {
	// 简化实现，返回模拟结果
	results := []model.CorrelationResult{
		{
			Factor1:         "运动频率",
			Factor2:         "症状严重程度",
			CorrelationType: "spearman",
			Coefficient:     -0.45,
			PValue:          0.01,
			SampleSize:      200,
		},
		{
			Factor1:         "睡眠质量",
			Factor2:         "治疗效果",
			CorrelationType: "pearson",
			Coefficient:     0.52,
			PValue:          0.005,
			SampleSize:      180,
		},
	}

	return results, nil
}
