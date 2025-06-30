package repository

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/model"
)

type EpidemiologyDB interface {
	// 环境暴露管理
	CreateEnvironmentExposure(ctx context.Context, exposure *model.EnvironmentExposure) (*model.EnvironmentExposure, error)
	GetEnvironmentExposure(ctx context.Context, exposureID int64) (*model.EnvironmentExposure, error)
	GetPatientEnvironmentExposures(ctx context.Context, patientID int64, offset, limit int) ([]*model.EnvironmentExposure, int64, error)
	UpdateEnvironmentExposure(ctx context.Context, exposure *model.EnvironmentExposure) error
	DeleteEnvironmentExposure(ctx context.Context, exposureID int64) error

	// 环境监测管理
	CreateEnvironmentMonitor(ctx context.Context, monitor *model.EnvironmentMonitor) (*model.EnvironmentMonitor, error)
	GetEnvironmentMonitor(ctx context.Context, monitorID int64) (*model.EnvironmentMonitor, error)
	GetEnvironmentMonitorsByLocation(ctx context.Context, location string, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error)
	GetEnvironmentMonitorsByTimeRange(ctx context.Context, startTime, endTime string, offset, limit int) ([]*model.EnvironmentMonitor, int64, error)
	UpdateEnvironmentMonitor(ctx context.Context, monitor *model.EnvironmentMonitor) error
	DeleteEnvironmentMonitor(ctx context.Context, monitorID int64) error

	// 生活方式调查管理
	CreateLifestyleSurvey(ctx context.Context, survey *model.LifestyleSurvey) (*model.LifestyleSurvey, error)
	GetLifestyleSurvey(ctx context.Context, surveyID int64) (*model.LifestyleSurvey, error)
	GetPatientLifestyleSurveys(ctx context.Context, patientID int64, offset, limit int) ([]*model.LifestyleSurvey, int64, error)
	GetLatestLifestyleSurvey(ctx context.Context, patientID int64) (*model.LifestyleSurvey, error)
	UpdateLifestyleSurvey(ctx context.Context, survey *model.LifestyleSurvey) error
	DeleteLifestyleSurvey(ctx context.Context, surveyID int64) error
}

type EpidemiologyCache interface {
	// 环境暴露缓存
	SetEnvironmentExposure(ctx context.Context, exposureID int64, exposure *model.EnvironmentExposure) error
	GetEnvironmentExposure(ctx context.Context, exposureID int64) (*model.EnvironmentExposure, error)
	DeleteEnvironmentExposure(ctx context.Context, exposureID int64) error

	// 环境监测缓存
	SetEnvironmentMonitor(ctx context.Context, monitorID int64, monitor *model.EnvironmentMonitor) error
	GetEnvironmentMonitor(ctx context.Context, monitorID int64) (*model.EnvironmentMonitor, error)
	DeleteEnvironmentMonitor(ctx context.Context, monitorID int64) error

	// 生活方式调查缓存
	SetLifestyleSurvey(ctx context.Context, surveyID int64, survey *model.LifestyleSurvey) error
	GetLifestyleSurvey(ctx context.Context, surveyID int64) (*model.LifestyleSurvey, error)
	DeleteLifestyleSurvey(ctx context.Context, surveyID int64) error

	// 患者最新生活方式调查缓存
	SetPatientLatestSurvey(ctx context.Context, patientID int64, survey *model.LifestyleSurvey) error
	GetPatientLatestSurvey(ctx context.Context, patientID int64) (*model.LifestyleSurvey, error)
	DeletePatientLatestSurvey(ctx context.Context, patientID int64) error
}
