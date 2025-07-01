package rpc

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/followup/controllers/rpc/pack"
	"github.com/yxrxy/AllergyBase/app/followup/domain/model"
	"github.com/yxrxy/AllergyBase/app/followup/usecase"
	"github.com/yxrxy/AllergyBase/kitex_gen/followup"
	"github.com/yxrxy/AllergyBase/pkg/base"
)

type FollowupHandler struct {
	useCase usecase.FollowupUseCase
}

func NewFollowupHandler(useCase usecase.FollowupUseCase) *FollowupHandler {
	return &FollowupHandler{
		useCase: useCase,
	}
}

// CreateFollowupPlan 创建随访计划
func (h *FollowupHandler) CreateFollowupPlan(ctx context.Context, req *followup.CreateFollowupPlanRequest) (r *followup.CreateFollowupPlanResponse, err error) {
	r = new(followup.CreateFollowupPlanResponse)

	plan := pack.BuildFollowupPlanFromRequest(req.Plan)

	var createdPlan *model.FollowupPlan
	if createdPlan, err = h.useCase.CreateFollowupPlan(ctx, plan); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.PlanId = createdPlan.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetFollowupPlan 获取随访计划
func (h *FollowupHandler) GetFollowupPlan(ctx context.Context, req *followup.GetFollowupPlanRequest) (r *followup.GetFollowupPlanResponse, err error) {
	r = new(followup.GetFollowupPlanResponse)

	var plan *model.FollowupPlan
	if plan, err = h.useCase.GetFollowupPlan(ctx, req.PlanId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Plan = pack.FollowupPlan(plan)
	r.Base = base.BuildBaseResp(err)
	return
}

// GetFollowupPlans 获取所有随访计划列表
func (h *FollowupHandler) GetFollowupPlans(ctx context.Context, req *followup.GetFollowupPlansRequest) (r *followup.GetFollowupPlansResponse, err error) {
	r = new(followup.GetFollowupPlansResponse)

	// 获取分页参数
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())

	// 设置默认值
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 调用 usecase 获取随访计划列表
	plans, total, err := h.useCase.GetActiveFollowupPlans(ctx, offset, pageSize)
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Plans = pack.FollowupPlanList(plans)
	r.Total = int32(total)
	r.Base = base.BuildBaseResp(err)
	return
}

// GetPatientFollowupPlans 获取患者的随访计划列表
func (h *FollowupHandler) GetPatientFollowupPlans(ctx context.Context, req *followup.GetPatientFollowupPlansRequest) (r *followup.GetPatientFollowupPlansResponse, err error) {
	r = new(followup.GetPatientFollowupPlansResponse)

	var plans []*model.FollowupPlan
	if plans, _, err = h.useCase.GetPatientFollowupPlans(ctx, req.PatientId, 0, 10); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Plans = pack.FollowupPlanList(plans)
	r.Base = base.BuildBaseResp(err)
	return
}

// CreateFollowupRecord 创建随访记录
func (h *FollowupHandler) CreateFollowupRecord(ctx context.Context, req *followup.CreateFollowupRecordRequest) (r *followup.CreateFollowupRecordResponse, err error) {
	r = new(followup.CreateFollowupRecordResponse)

	record := pack.BuildFollowupRecordFromRequest(req.Record)

	var createdRecord *model.FollowupRecord
	if createdRecord, err = h.useCase.CreateFollowupRecord(ctx, record); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.RecordId = createdRecord.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetFollowupRecord 获取单个随访记录
func (h *FollowupHandler) GetFollowupRecord(ctx context.Context, req *followup.GetFollowupRecordRequest) (r *followup.GetFollowupRecordResponse, err error) {
	r = new(followup.GetFollowupRecordResponse)

	var record *model.FollowupRecord
	if record, err = h.useCase.GetFollowupRecord(ctx, req.RecordId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Record = pack.FollowupRecord(record)
	r.Base = base.BuildBaseResp(err)
	return
}

// GetPlanFollowupRecords 获取计划的随访记录列表
func (h *FollowupHandler) GetPlanFollowupRecords(ctx context.Context, req *followup.GetPlanFollowupRecordsRequest) (r *followup.GetPlanFollowupRecordsResponse, err error) {
	r = new(followup.GetPlanFollowupRecordsResponse)

	var records []*model.FollowupRecord
	if records, _, err = h.useCase.GetPlanFollowupRecords(ctx, req.PlanId, 0, 10); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Records = pack.FollowupRecordList(records)
	r.Base = base.BuildBaseResp(err)
	return
}

// AddMedicationMonitor 添加用药监测
func (h *FollowupHandler) AddMedicationMonitor(ctx context.Context, req *followup.AddMedicationMonitorRequest) (r *followup.AddMedicationMonitorResponse, err error) {
	r = new(followup.AddMedicationMonitorResponse)

	monitor := pack.BuildMedicationMonitorFromRequest(req.Monitor)

	var createdMonitor *model.MedicationMonitor
	if createdMonitor, err = h.useCase.CreateMedicationMonitor(ctx, monitor); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.MonitorId = createdMonitor.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetFollowupMedications 获取随访的用药监测列表
func (h *FollowupHandler) GetFollowupMedications(ctx context.Context, req *followup.GetFollowupMedicationsRequest) (r *followup.GetFollowupMedicationsResponse, err error) {
	r = new(followup.GetFollowupMedicationsResponse)

	var monitors []*model.MedicationMonitor
	if monitors, err = h.useCase.GetFollowupMedicationMonitors(ctx, req.FollowupId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Medications = pack.MedicationMonitorList(monitors)
	r.Base = base.BuildBaseResp(err)
	return
}
