package pack

import (
	"time"

	"github.com/yxrxy/AllergyBase/app/followup/domain/model"
	rpcmodel "github.com/yxrxy/AllergyBase/kitex_gen/model"
)

// FollowupPlan 将数据库随访计划实体转换为RPC响应实体
func FollowupPlan(p *model.FollowupPlan) *rpcmodel.FollowupPlan {
	if p == nil {
		return nil
	}
	return &rpcmodel.FollowupPlan{
		Id:                  p.ID,
		PatientId:           p.PatientID,
		PlanStartDate:       p.StartDate.Format("2006-01-02"),
		FollowupFrequency:   &p.Frequency,
		FollowupMethod:      &p.PlanType,
		ResponsibleDoctorId: func() *int64 { id := int64(0); return &id }(),
	}
}

// FollowupRecord 将数据库随访记录实体转换为RPC响应实体
func FollowupRecord(r *model.FollowupRecord) *rpcmodel.FollowupRecord {
	if r == nil {
		return nil
	}

	// 默认值
	symptomScore := int32(0)
	qualityScore := int32(0)

	return &rpcmodel.FollowupRecord{
		Id:                   r.ID,
		PlanId:               r.PlanID,
		FollowupDate:         r.FollowupDate.Format("2006-01-02"),
		SymptomScore:         &symptomScore,
		MedicationCompliance: &r.TreatmentEffect,
		SideEffects:          &r.Symptoms,
		QualityOfLifeScore:   &qualityScore,
	}
}

// MedicationMonitor 将数据库用药监测实体转换为RPC响应实体
func MedicationMonitor(m *model.MedicationMonitor) *rpcmodel.MedicationMonitor {
	if m == nil {
		return nil
	}
	return &rpcmodel.MedicationMonitor{
		Id:              m.ID,
		FollowupId:      m.FollowupID,
		DrugName:        m.MedicationName,
		UsageFrequency:  &m.Frequency,
		ComplianceLevel: &m.Compliance,
		Effectiveness:   &m.SideEffects,
	}
}

// FollowupPlanList 将随访计划列表转换为RPC响应实体
func FollowupPlanList(plans []*model.FollowupPlan) []*rpcmodel.FollowupPlan {
	if len(plans) == 0 {
		return nil
	}
	result := make([]*rpcmodel.FollowupPlan, 0, len(plans))
	for _, p := range plans {
		result = append(result, FollowupPlan(p))
	}
	return result
}

// FollowupRecordList 将随访记录列表转换为RPC响应实体
func FollowupRecordList(records []*model.FollowupRecord) []*rpcmodel.FollowupRecord {
	if len(records) == 0 {
		return nil
	}
	result := make([]*rpcmodel.FollowupRecord, 0, len(records))
	for _, r := range records {
		result = append(result, FollowupRecord(r))
	}
	return result
}

// MedicationMonitorList 将用药监测列表转换为RPC响应实体
func MedicationMonitorList(monitors []*model.MedicationMonitor) []*rpcmodel.MedicationMonitor {
	if len(monitors) == 0 {
		return nil
	}
	result := make([]*rpcmodel.MedicationMonitor, 0, len(monitors))
	for _, m := range monitors {
		result = append(result, MedicationMonitor(m))
	}
	return result
}

// BuildFollowupPlanFromRequest 将RPC请求转换为数据库实体
func BuildFollowupPlanFromRequest(req *rpcmodel.FollowupPlan) *model.FollowupPlan {
	if req == nil {
		return nil
	}

	var startDate time.Time
	if req.PlanStartDate != "" {
		if t, err := time.Parse("2006-01-02", req.PlanStartDate); err == nil {
			startDate = t
		}
	}

	return &model.FollowupPlan{
		ID:        req.Id,
		PatientID: req.PatientId,
		StartDate: startDate,
		Frequency: req.GetFollowupFrequency(),
		PlanType:  req.GetFollowupMethod(),
		CreatedBy: "system",
		Status:    "active",
	}
}

// BuildFollowupRecordFromRequest 将RPC请求转换为数据库实体
func BuildFollowupRecordFromRequest(req *rpcmodel.FollowupRecord) *model.FollowupRecord {
	if req == nil {
		return nil
	}

	var followupDate time.Time
	if req.FollowupDate != "" {
		if t, err := time.Parse("2006-01-02", req.FollowupDate); err == nil {
			followupDate = t
		}
	}

	return &model.FollowupRecord{
		ID:              req.Id,
		PlanID:          req.PlanId,
		FollowupDate:    followupDate,
		Symptoms:        req.GetSideEffects(),
		TreatmentEffect: req.GetMedicationCompliance(),
		FollowupMethod:  "phone", // 默认值
		FollowupDoctor:  "system",
	}
}

// BuildMedicationMonitorFromRequest 将RPC请求转换为数据库实体
func BuildMedicationMonitorFromRequest(req *rpcmodel.MedicationMonitor) *model.MedicationMonitor {
	if req == nil {
		return nil
	}
	return &model.MedicationMonitor{
		ID:             req.Id,
		FollowupID:     req.FollowupId,
		MedicationName: req.DrugName,
		Frequency:      req.GetUsageFrequency(),
		Compliance:     req.GetComplianceLevel(),
		SideEffects:    req.GetEffectiveness(),
		Dosage:         "default", // 默认值
	}
}
