package usecase

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/model"
	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/repository"
	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/service"
)

type EpidemiologyUseCase interface {
	// 环境暴露相关
	CreateEnvironmentExposure(ctx context.Context, exposure *model.EnvironmentExposure) (*model.EnvironmentExposure, error)
	GetEnvironmentExposure(ctx context.Context, exposureID int64) (*model.EnvironmentExposure, error)
	GetPatientEnvironmentExposures(ctx context.Context, patientID int64, offset, limit int) ([]*model.EnvironmentExposure, int64, error)
	UpdateEnvironmentExposure(ctx context.Context, exposure *model.EnvironmentExposure) error
	DeleteEnvironmentExposure(ctx context.Context, exposureID int64) error

	// 环境监测相关
	CreateEnvironmentMonitor(ctx context.Context, monitor *model.EnvironmentMonitor) (*model.EnvironmentMonitor, error)
	GetEnvironmentMonitor(ctx context.Context, monitorID int64) (*model.EnvironmentMonitor, error)
	GetEnvironmentMonitorsByLocation(ctx context.Context, location, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error)
	GetEnvironmentMonitorsByTimeRange(ctx context.Context, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error)
	UpdateEnvironmentMonitor(ctx context.Context, monitor *model.EnvironmentMonitor) error
	DeleteEnvironmentMonitor(ctx context.Context, monitorID int64) error

	// 生活方式调查相关
	CreateLifestyleSurvey(ctx context.Context, survey *model.LifestyleSurvey) (*model.LifestyleSurvey, error)
	GetLifestyleSurvey(ctx context.Context, surveyID int64) (*model.LifestyleSurvey, error)
	GetPatientLifestyleSurveys(ctx context.Context, patientID int64, offset, limit int) ([]*model.LifestyleSurvey, int64, error)
	GetLatestLifestyleSurvey(ctx context.Context, patientID int64) (*model.LifestyleSurvey, error)
	UpdateLifestyleSurvey(ctx context.Context, survey *model.LifestyleSurvey) error
	DeleteLifestyleSurvey(ctx context.Context, surveyID int64) error
}

type useCase struct {
	db      repository.EpidemiologyDB
	service *service.EpidemiologyService
}

func NewEpidemiologyUseCase(db repository.EpidemiologyDB, service *service.EpidemiologyService) EpidemiologyUseCase {
	return &useCase{
		db:      db,
		service: service,
	}
}

// 环境暴露相关实现
func (u *useCase) CreateEnvironmentExposure(ctx context.Context, exposure *model.EnvironmentExposure) (*model.EnvironmentExposure, error) {
	return u.service.CreateEnvironmentExposure(ctx, exposure)
}

func (u *useCase) GetEnvironmentExposure(ctx context.Context, exposureID int64) (*model.EnvironmentExposure, error) {
	return u.service.GetEnvironmentExposure(ctx, exposureID)
}

func (u *useCase) GetPatientEnvironmentExposures(ctx context.Context, patientID int64, offset, limit int) ([]*model.EnvironmentExposure, int64, error) {
	return u.service.GetPatientEnvironmentExposures(ctx, patientID, offset, limit)
}

func (u *useCase) UpdateEnvironmentExposure(ctx context.Context, exposure *model.EnvironmentExposure) error {
	return u.db.UpdateEnvironmentExposure(ctx, exposure)
}

func (u *useCase) DeleteEnvironmentExposure(ctx context.Context, exposureID int64) error {
	return u.db.DeleteEnvironmentExposure(ctx, exposureID)
}

// 环境监测相关实现
func (u *useCase) CreateEnvironmentMonitor(ctx context.Context, monitor *model.EnvironmentMonitor) (*model.EnvironmentMonitor, error) {
	return u.service.CreateEnvironmentMonitor(ctx, monitor)
}

func (u *useCase) GetEnvironmentMonitor(ctx context.Context, monitorID int64) (*model.EnvironmentMonitor, error) {
	return u.service.GetEnvironmentMonitor(ctx, monitorID)
}

func (u *useCase) GetEnvironmentMonitorsByLocation(ctx context.Context, location, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error) {
	return u.service.GetEnvironmentMonitorsByLocation(ctx, location, startTime, endTime, offset, limit)
}

func (u *useCase) GetEnvironmentMonitorsByTimeRange(ctx context.Context, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error) {
	return u.service.GetEnvironmentMonitorsByTimeRange(ctx, startTime, endTime, offset, limit)
}

func (u *useCase) UpdateEnvironmentMonitor(ctx context.Context, monitor *model.EnvironmentMonitor) error {
	return u.db.UpdateEnvironmentMonitor(ctx, monitor)
}

func (u *useCase) DeleteEnvironmentMonitor(ctx context.Context, monitorID int64) error {
	return u.db.DeleteEnvironmentMonitor(ctx, monitorID)
}

// 生活方式调查相关实现
func (u *useCase) CreateLifestyleSurvey(ctx context.Context, survey *model.LifestyleSurvey) (*model.LifestyleSurvey, error) {
	return u.service.CreateLifestyleSurvey(ctx, survey)
}

func (u *useCase) GetLifestyleSurvey(ctx context.Context, surveyID int64) (*model.LifestyleSurvey, error) {
	return u.service.GetLifestyleSurvey(ctx, surveyID)
}

func (u *useCase) GetPatientLifestyleSurveys(ctx context.Context, patientID int64, offset, limit int) ([]*model.LifestyleSurvey, int64, error) {
	return u.service.GetPatientLifestyleSurveys(ctx, patientID, offset, limit)
}

func (u *useCase) GetLatestLifestyleSurvey(ctx context.Context, patientID int64) (*model.LifestyleSurvey, error) {
	return u.service.GetLatestLifestyleSurvey(ctx, patientID)
}

func (u *useCase) UpdateLifestyleSurvey(ctx context.Context, survey *model.LifestyleSurvey) error {
	return u.db.UpdateLifestyleSurvey(ctx, survey)
}

func (u *useCase) DeleteLifestyleSurvey(ctx context.Context, surveyID int64) error {
	return u.db.DeleteLifestyleSurvey(ctx, surveyID)
}
