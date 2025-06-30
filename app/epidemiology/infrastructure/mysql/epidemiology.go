package mysql

import (
	"context"
	"time"

	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/model"
	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/repository"
	"gorm.io/gorm"
)

type epidemiologyDB struct {
	db *gorm.DB
}

func NewEpidemiologyDB(db *gorm.DB) repository.EpidemiologyDB {
	return &epidemiologyDB{db: db}
}

// 环境暴露管理
func (e *epidemiologyDB) CreateEnvironmentExposure(ctx context.Context, exposure *model.EnvironmentExposure) (*model.EnvironmentExposure, error) {
	if err := e.db.WithContext(ctx).Create(exposure).Error; err != nil {
		return nil, err
	}
	return exposure, nil
}

func (e *epidemiologyDB) GetEnvironmentExposure(ctx context.Context, exposureID int64) (*model.EnvironmentExposure, error) {
	var exposure model.EnvironmentExposure
	if err := e.db.WithContext(ctx).First(&exposure, exposureID).Error; err != nil {
		return nil, err
	}
	return &exposure, nil
}

func (e *epidemiologyDB) GetPatientEnvironmentExposures(ctx context.Context, patientID int64, offset, limit int) ([]*model.EnvironmentExposure, int64, error) {
	var exposures []*model.EnvironmentExposure
	var total int64

	if err := e.db.WithContext(ctx).Model(&model.EnvironmentExposure{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := e.db.WithContext(ctx).Where("patient_id = ?", patientID).
		Order("created_at DESC").Limit(limit).Offset(offset).Find(&exposures).Error; err != nil {
		return nil, 0, err
	}

	return exposures, total, nil
}

func (e *epidemiologyDB) UpdateEnvironmentExposure(ctx context.Context, exposure *model.EnvironmentExposure) error {
	return e.db.WithContext(ctx).Save(exposure).Error
}

func (e *epidemiologyDB) DeleteEnvironmentExposure(ctx context.Context, exposureID int64) error {
	return e.db.WithContext(ctx).Delete(&model.EnvironmentExposure{}, exposureID).Error
}

// 环境监测管理
func (e *epidemiologyDB) CreateEnvironmentMonitor(ctx context.Context, monitor *model.EnvironmentMonitor) (*model.EnvironmentMonitor, error) {
	if err := e.db.WithContext(ctx).Create(monitor).Error; err != nil {
		return nil, err
	}
	return monitor, nil
}

func (e *epidemiologyDB) GetEnvironmentMonitor(ctx context.Context, monitorID int64) (*model.EnvironmentMonitor, error) {
	var monitor model.EnvironmentMonitor
	if err := e.db.WithContext(ctx).First(&monitor, monitorID).Error; err != nil {
		return nil, err
	}
	return &monitor, nil
}

func (e *epidemiologyDB) GetEnvironmentMonitorsByLocation(ctx context.Context, location, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error) {
	var monitors []*model.EnvironmentMonitor
	var total int64

	query := e.db.WithContext(ctx).Model(&model.EnvironmentMonitor{}).Where("location_code = ?", location)

	if startTime != "" && endTime != "" {
		start, _ := time.Parse("2006-01-02 15:04:05", startTime)
		end, _ := time.Parse("2006-01-02 15:04:05", endTime)
		query = query.Where("monitor_time BETWEEN ? AND ?", start, end)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("monitor_time DESC").Limit(limit).Offset(offset).Find(&monitors).Error; err != nil {
		return nil, 0, err
	}

	return monitors, total, nil
}

func (e *epidemiologyDB) GetEnvironmentMonitorsByTimeRange(ctx context.Context, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error) {
	var monitors []*model.EnvironmentMonitor
	var total int64

	query := e.db.WithContext(ctx).Model(&model.EnvironmentMonitor{})

	if startTime != "" && endTime != "" {
		start, _ := time.Parse("2006-01-02 15:04:05", startTime)
		end, _ := time.Parse("2006-01-02 15:04:05", endTime)
		query = query.Where("monitor_time BETWEEN ? AND ?", start, end)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("monitor_time DESC").Limit(limit).Offset(offset).Find(&monitors).Error; err != nil {
		return nil, 0, err
	}

	return monitors, total, nil
}

func (e *epidemiologyDB) UpdateEnvironmentMonitor(ctx context.Context, monitor *model.EnvironmentMonitor) error {
	return e.db.WithContext(ctx).Save(monitor).Error
}

func (e *epidemiologyDB) DeleteEnvironmentMonitor(ctx context.Context, monitorID int64) error {
	return e.db.WithContext(ctx).Delete(&model.EnvironmentMonitor{}, monitorID).Error
}

// 生活方式调查管理
func (e *epidemiologyDB) CreateLifestyleSurvey(ctx context.Context, survey *model.LifestyleSurvey) (*model.LifestyleSurvey, error) {
	if err := e.db.WithContext(ctx).Create(survey).Error; err != nil {
		return nil, err
	}
	return survey, nil
}

func (e *epidemiologyDB) GetLifestyleSurvey(ctx context.Context, surveyID int64) (*model.LifestyleSurvey, error) {
	var survey model.LifestyleSurvey
	if err := e.db.WithContext(ctx).First(&survey, surveyID).Error; err != nil {
		return nil, err
	}
	return &survey, nil
}

func (e *epidemiologyDB) GetPatientLifestyleSurveys(ctx context.Context, patientID int64, offset, limit int) ([]*model.LifestyleSurvey, int64, error) {
	var surveys []*model.LifestyleSurvey
	var total int64

	if err := e.db.WithContext(ctx).Model(&model.LifestyleSurvey{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := e.db.WithContext(ctx).Where("patient_id = ?", patientID).
		Order("created_at DESC").Limit(limit).Offset(offset).Find(&surveys).Error; err != nil {
		return nil, 0, err
	}

	return surveys, total, nil
}

func (e *epidemiologyDB) GetLatestLifestyleSurvey(ctx context.Context, patientID int64) (*model.LifestyleSurvey, error) {
	var survey model.LifestyleSurvey
	if err := e.db.WithContext(ctx).Where("patient_id = ?", patientID).
		Order("created_at DESC").First(&survey).Error; err != nil {
		return nil, err
	}
	return &survey, nil
}

func (e *epidemiologyDB) UpdateLifestyleSurvey(ctx context.Context, survey *model.LifestyleSurvey) error {
	return e.db.WithContext(ctx).Save(survey).Error
}

func (e *epidemiologyDB) DeleteLifestyleSurvey(ctx context.Context, surveyID int64) error {
	return e.db.WithContext(ctx).Delete(&model.LifestyleSurvey{}, surveyID).Error
}
