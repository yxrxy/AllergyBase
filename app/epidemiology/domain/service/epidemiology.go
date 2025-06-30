package service

import (
	"context"
	"time"

	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/model"
	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/repository"
)

type EpidemiologyService struct {
	db    repository.EpidemiologyDB
	cache repository.EpidemiologyCache
}

func NewEpidemiologyService(db repository.EpidemiologyDB, cache repository.EpidemiologyCache) *EpidemiologyService {
	return &EpidemiologyService{
		db:    db,
		cache: cache,
	}
}

// 环境暴露相关方法
func (s *EpidemiologyService) CreateEnvironmentExposure(ctx context.Context, exposure *model.EnvironmentExposure) (*model.EnvironmentExposure, error) {
	// 设置创建时间
	exposure.CreatedAt = time.Now()
	exposure.UpdatedAt = time.Now()

	createdExposure, err := s.db.CreateEnvironmentExposure(ctx, exposure)
	if err != nil {
		return nil, err
	}

	// 缓存环境暴露信息
	_ = s.cache.SetEnvironmentExposure(ctx, createdExposure.ID, createdExposure)

	return createdExposure, nil
}

func (s *EpidemiologyService) GetEnvironmentExposure(ctx context.Context, exposureID int64) (*model.EnvironmentExposure, error) {
	// 先从缓存获取
	exposure, err := s.cache.GetEnvironmentExposure(ctx, exposureID)
	if err == nil && exposure != nil {
		return exposure, nil
	}

	// 从数据库获取
	exposure, err = s.db.GetEnvironmentExposure(ctx, exposureID)
	if err != nil {
		return nil, err
	}

	// 缓存环境暴露信息
	_ = s.cache.SetEnvironmentExposure(ctx, exposureID, exposure)

	return exposure, nil
}

func (s *EpidemiologyService) GetPatientEnvironmentExposures(ctx context.Context, patientID int64, offset, limit int) ([]*model.EnvironmentExposure, int64, error) {
	return s.db.GetPatientEnvironmentExposures(ctx, patientID, offset, limit)
}

// 环境监测相关方法
func (s *EpidemiologyService) CreateEnvironmentMonitor(ctx context.Context, monitor *model.EnvironmentMonitor) (*model.EnvironmentMonitor, error) {
	// 设置创建时间
	monitor.CreatedAt = time.Now()
	monitor.UpdatedAt = time.Now()

	createdMonitor, err := s.db.CreateEnvironmentMonitor(ctx, monitor)
	if err != nil {
		return nil, err
	}

	// 缓存环境监测信息
	_ = s.cache.SetEnvironmentMonitor(ctx, createdMonitor.ID, createdMonitor)

	return createdMonitor, nil
}

func (s *EpidemiologyService) GetEnvironmentMonitor(ctx context.Context, monitorID int64) (*model.EnvironmentMonitor, error) {
	// 先从缓存获取
	monitor, err := s.cache.GetEnvironmentMonitor(ctx, monitorID)
	if err == nil && monitor != nil {
		return monitor, nil
	}

	// 从数据库获取
	monitor, err = s.db.GetEnvironmentMonitor(ctx, monitorID)
	if err != nil {
		return nil, err
	}

	// 缓存环境监测信息
	_ = s.cache.SetEnvironmentMonitor(ctx, monitorID, monitor)

	return monitor, nil
}

func (s *EpidemiologyService) GetEnvironmentMonitorsByLocation(ctx context.Context, location, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error) {
	return s.db.GetEnvironmentMonitorsByLocation(ctx, location, startTime, endTime, offset, limit)
}

func (s *EpidemiologyService) GetEnvironmentMonitorsByTimeRange(ctx context.Context, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error) {
	return s.db.GetEnvironmentMonitorsByTimeRange(ctx, startTime, endTime, offset, limit)
}

// 生活方式调查相关方法
func (s *EpidemiologyService) CreateLifestyleSurvey(ctx context.Context, survey *model.LifestyleSurvey) (*model.LifestyleSurvey, error) {
	// 设置创建时间
	survey.CreatedAt = time.Now()
	survey.UpdatedAt = time.Now()

	createdSurvey, err := s.db.CreateLifestyleSurvey(ctx, survey)
	if err != nil {
		return nil, err
	}

	// 缓存生活方式调查信息
	_ = s.cache.SetLifestyleSurvey(ctx, createdSurvey.ID, createdSurvey)

	// 更新患者最新调查缓存
	_ = s.cache.SetPatientLatestSurvey(ctx, survey.PatientID, createdSurvey)

	return createdSurvey, nil
}

func (s *EpidemiologyService) GetLifestyleSurvey(ctx context.Context, surveyID int64) (*model.LifestyleSurvey, error) {
	// 先从缓存获取
	survey, err := s.cache.GetLifestyleSurvey(ctx, surveyID)
	if err == nil && survey != nil {
		return survey, nil
	}

	// 从数据库获取
	survey, err = s.db.GetLifestyleSurvey(ctx, surveyID)
	if err != nil {
		return nil, err
	}

	// 缓存生活方式调查信息
	_ = s.cache.SetLifestyleSurvey(ctx, surveyID, survey)

	return survey, nil
}

func (s *EpidemiologyService) GetPatientLifestyleSurveys(ctx context.Context, patientID int64, offset, limit int) ([]*model.LifestyleSurvey, int64, error) {
	return s.db.GetPatientLifestyleSurveys(ctx, patientID, offset, limit)
}

func (s *EpidemiologyService) GetLatestLifestyleSurvey(ctx context.Context, patientID int64) (*model.LifestyleSurvey, error) {
	// 先从缓存获取最新调查
	survey, err := s.cache.GetPatientLatestSurvey(ctx, patientID)
	if err == nil && survey != nil {
		return survey, nil
	}

	// 从数据库获取
	survey, err = s.db.GetLatestLifestyleSurvey(ctx, patientID)
	if err != nil {
		return nil, err
	}

	// 缓存患者最新调查
	_ = s.cache.SetPatientLatestSurvey(ctx, patientID, survey)

	return survey, nil
}
