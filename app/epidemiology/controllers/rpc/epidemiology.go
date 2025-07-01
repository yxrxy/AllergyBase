package rpc

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/epidemiology/controllers/rpc/pack"
	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/model"
	"github.com/yxrxy/AllergyBase/app/epidemiology/usecase"
	"github.com/yxrxy/AllergyBase/kitex_gen/epidemiology"
	"github.com/yxrxy/AllergyBase/pkg/base"
)

type EpidemiologyHandler struct {
	useCase usecase.EpidemiologyUseCase
}

func NewEpidemiologyHandler(useCase usecase.EpidemiologyUseCase) *EpidemiologyHandler {
	return &EpidemiologyHandler{useCase: useCase}
}

// CreateEnvironmentExposure 创建环境暴露记录
func (h *EpidemiologyHandler) CreateEnvironmentExposure(ctx context.Context, req *epidemiology.CreateEnvironmentExposureRequest) (r *epidemiology.CreateEnvironmentExposureResponse, err error) {
	r = new(epidemiology.CreateEnvironmentExposureResponse)

	exposure := pack.BuildEnvironmentExposureFromRequest(req.Exposure)

	var createdExposure *model.EnvironmentExposure
	if createdExposure, err = h.useCase.CreateEnvironmentExposure(ctx, exposure); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.ExposureId = createdExposure.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetEnvironmentExposure 获取环境暴露记录
func (h *EpidemiologyHandler) GetEnvironmentExposure(ctx context.Context, req *epidemiology.GetEnvironmentExposureRequest) (r *epidemiology.GetEnvironmentExposureResponse, err error) {
	r = new(epidemiology.GetEnvironmentExposureResponse)

	var exposure *model.EnvironmentExposure
	if exposure, err = h.useCase.GetEnvironmentExposure(ctx, req.PatientId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Exposure = pack.EnvironmentExposure(exposure)
	r.Base = base.BuildBaseResp(err)
	return
}

// GetPatientEnvironmentExposures 获取患者环境暴露记录列表
func (h *EpidemiologyHandler) GetPatientEnvironmentExposures(ctx context.Context, req *epidemiology.GetPatientEnvironmentExposuresRequest) (r *epidemiology.GetPatientEnvironmentExposuresResponse, err error) {
	r = new(epidemiology.GetPatientEnvironmentExposuresResponse)

	var exposures []*model.EnvironmentExposure
	var total int64
	if exposures, total, err = h.useCase.GetPatientEnvironmentExposures(ctx, req.PatientId, int(req.Offset), int(req.Limit)); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Exposures = pack.EnvironmentExposureList(exposures)
	r.Total = int32(total)
	r.Base = base.BuildBaseResp(err)
	return
}

// AddEnvironmentMonitor 添加环境监测数据
func (h *EpidemiologyHandler) AddEnvironmentMonitor(ctx context.Context, req *epidemiology.AddEnvironmentMonitorRequest) (r *epidemiology.AddEnvironmentMonitorResponse, err error) {
	r = new(epidemiology.AddEnvironmentMonitorResponse)

	monitor := pack.BuildEnvironmentMonitorFromRequest(req.Monitor)

	var createdMonitor *model.EnvironmentMonitor
	if createdMonitor, err = h.useCase.CreateEnvironmentMonitor(ctx, monitor); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.MonitorId = createdMonitor.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetEnvironmentMonitors 获取环境监测数据
func (h *EpidemiologyHandler) GetEnvironmentMonitors(ctx context.Context, req *epidemiology.GetEnvironmentMonitorsRequest) (r *epidemiology.GetEnvironmentMonitorsResponse, err error) {
	r = new(epidemiology.GetEnvironmentMonitorsResponse)

	var monitors []*model.EnvironmentMonitor
	if monitors, _, err = h.useCase.GetEnvironmentMonitorsByLocation(ctx, req.LocationCode, req.StartTime, req.EndTime, 0, 100); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Monitors = pack.EnvironmentMonitorList(monitors)
	r.Base = base.BuildBaseResp(err)
	return
}

// CreateLifestyleSurvey 创建生活方式调查
func (h *EpidemiologyHandler) CreateLifestyleSurvey(ctx context.Context, req *epidemiology.CreateLifestyleSurveyRequest) (r *epidemiology.CreateLifestyleSurveyResponse, err error) {
	r = new(epidemiology.CreateLifestyleSurveyResponse)

	survey := pack.BuildLifestyleSurveyFromRequest(req.Survey)

	var createdSurvey *model.LifestyleSurvey
	if createdSurvey, err = h.useCase.CreateLifestyleSurvey(ctx, survey); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.SurveyId = createdSurvey.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetLifestyleSurvey 获取生活方式调查
func (h *EpidemiologyHandler) GetLifestyleSurvey(ctx context.Context, req *epidemiology.GetLifestyleSurveyRequest) (r *epidemiology.GetLifestyleSurveyResponse, err error) {
	r = new(epidemiology.GetLifestyleSurveyResponse)

	var survey *model.LifestyleSurvey
	if survey, err = h.useCase.GetLifestyleSurvey(ctx, req.PatientId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Survey = pack.LifestyleSurvey(survey)
	r.Base = base.BuildBaseResp(err)
	return
}
