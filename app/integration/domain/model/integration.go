package model

import (
	"time"

	biobankModel "github.com/yxrxy/AllergyBase/app/biobank/domain/model"
	clinicalModel "github.com/yxrxy/AllergyBase/app/clinical/domain/model"
	epidemiologyModel "github.com/yxrxy/AllergyBase/app/epidemiology/domain/model"
	followupModel "github.com/yxrxy/AllergyBase/app/followup/domain/model"
)

// PatientIntegratedView 患者综合视图
type PatientIntegratedView struct {
	// 基础信息
	Patient *clinicalModel.Patient `json:"patient"`

	// 临床数据
	Visits       []clinicalModel.Visit       `json:"visits"`
	Diagnoses    []clinicalModel.Diagnosis   `json:"diagnoses"`
	Examinations []clinicalModel.Examination `json:"examinations"`

	// 流调数据
	EnvironmentExposures []epidemiologyModel.EnvironmentExposure `json:"environment_exposures"`
	LifestyleSurveys     []epidemiologyModel.LifestyleSurvey     `json:"lifestyle_surveys"`

	// 随访数据
	FollowupPlans   []followupModel.FollowupPlan      `json:"followup_plans"`
	FollowupRecords []followupModel.FollowupRecord    `json:"followup_records"`
	Medications     []followupModel.MedicationMonitor `json:"medications"`

	// 生物样本数据
	Samples        []biobankModel.Sample         `json:"samples"`
	GenomicData    []biobankModel.GenomicData    `json:"genomic_data"`
	ProteomicsData []biobankModel.ProteomicsData `json:"proteomics_data"`

	// 统计信息
	Statistics *PatientStatistics `json:"statistics"`
}

// PatientStatistics 患者统计信息
type PatientStatistics struct {
	TotalVisits       int        `json:"total_visits"`
	FirstVisitDate    time.Time  `json:"first_visit_date"`
	LastVisitDate     time.Time  `json:"last_visit_date"`
	TotalDiagnoses    int        `json:"total_diagnoses"`
	TotalExaminations int        `json:"total_examinations"`
	TotalFollowups    int        `json:"total_followups"`
	TotalSamples      int        `json:"total_samples"`
	LastFollowupDate  *time.Time `json:"last_followup_date"`
	NextFollowupDate  *time.Time `json:"next_followup_date"`
	ActiveMedications int        `json:"active_medications"`
}

// DataQualityReport 数据质量报告
type DataQualityReport struct {
	PatientID         int64              `json:"patient_id"`
	CheckTime         time.Time          `json:"check_time"`
	OverallScore      float64            `json:"overall_score"`
	ClinicalScore     float64            `json:"clinical_score"`
	EpidemiologyScore float64            `json:"epidemiology_score"`
	FollowupScore     float64            `json:"followup_score"`
	BiobankScore      float64            `json:"biobank_score"`
	Issues            []DataQualityIssue `json:"issues"`
	Recommendations   []string           `json:"recommendations"`
}

// DataQualityIssue 数据质量问题
type DataQualityIssue struct {
	Category    string `json:"category"` // clinical, epidemiology, followup, biobank
	Severity    string `json:"severity"` // low, medium, high, critical
	Field       string `json:"field"`
	Description string `json:"description"`
	Suggestion  string `json:"suggestion"`
}

// CrossDatabaseQuery 跨数据库查询条件
type CrossDatabaseQuery struct {
	// 基础筛选条件
	PatientIDs []int64    `json:"patient_ids"`
	Gender     *string    `json:"gender"`
	AgeRange   *AgeRange  `json:"age_range"`
	DateRange  *DateRange `json:"date_range"`

	// 临床条件
	DiagnosisCodes []string `json:"diagnosis_codes"`
	Departments    []string `json:"departments"`

	// 流调条件
	EnvironmentFactors []string `json:"environment_factors"`
	LifestyleFactors   []string `json:"lifestyle_factors"`

	// 随访条件
	FollowupStatus  []string `json:"followup_status"`
	MedicationNames []string `json:"medication_names"`

	// 生物样本条件
	SampleTypes    []string `json:"sample_types"`
	HasGenomicData *bool    `json:"has_genomic_data"`

	// 分页参数
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// AgeRange 年龄范围
type AgeRange struct {
	Min *int `json:"min"`
	Max *int `json:"max"`
}

// DateRange 日期范围
type DateRange struct {
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

// CorrelationAnalysis 关联分析结果
type CorrelationAnalysis struct {
	AnalysisID    string                 `json:"analysis_id"`
	AnalysisType  string                 `json:"analysis_type"` // environmental_disease, lifestyle_outcome, etc.
	CreatedAt     time.Time              `json:"created_at"`
	Parameters    map[string]interface{} `json:"parameters"`
	Results       []CorrelationResult    `json:"results"`
	Visualization *VisualizationConfig   `json:"visualization"`
}

// CorrelationResult 关联分析单个结果
type CorrelationResult struct {
	Factor1         string  `json:"factor1"`
	Factor2         string  `json:"factor2"`
	CorrelationType string  `json:"correlation_type"` // pearson, spearman, etc.
	Coefficient     float64 `json:"coefficient"`
	PValue          float64 `json:"p_value"`
	Significance    string  `json:"significance"` // significant, not_significant
	SampleSize      int     `json:"sample_size"`
}

// VisualizationConfig 可视化配置
type VisualizationConfig struct {
	ChartType   string                 `json:"chart_type"`
	Title       string                 `json:"title"`
	XAxis       string                 `json:"x_axis"`
	YAxis       string                 `json:"y_axis"`
	ColorScheme string                 `json:"color_scheme"`
	Options     map[string]interface{} `json:"options"`
}
