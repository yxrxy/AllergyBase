package repository

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/followup/domain/model"
)

type FollowupDB interface {
	// 随访计划管理
	CreateFollowupPlan(ctx context.Context, plan *model.FollowupPlan) (*model.FollowupPlan, error)
	GetFollowupPlan(ctx context.Context, planID int64) (*model.FollowupPlan, error)
	GetPatientFollowupPlans(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupPlan, int64, error)
	UpdateFollowupPlan(ctx context.Context, plan *model.FollowupPlan) error
	DeleteFollowupPlan(ctx context.Context, planID int64) error
	GetActiveFollowupPlans(ctx context.Context, offset, limit int) ([]*model.FollowupPlan, int64, error)

	// 随访记录管理
	CreateFollowupRecord(ctx context.Context, record *model.FollowupRecord) (*model.FollowupRecord, error)
	GetFollowupRecord(ctx context.Context, recordID int64) (*model.FollowupRecord, error)
	GetPlanFollowupRecords(ctx context.Context, planID int64, offset, limit int) ([]*model.FollowupRecord, int64, error)
	GetPatientFollowupRecords(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupRecord, int64, error)
	UpdateFollowupRecord(ctx context.Context, record *model.FollowupRecord) error
	DeleteFollowupRecord(ctx context.Context, recordID int64) error
	GetLatestFollowupRecord(ctx context.Context, patientID int64) (*model.FollowupRecord, error)

	// 用药监测管理
	CreateMedicationMonitor(ctx context.Context, monitor *model.MedicationMonitor) (*model.MedicationMonitor, error)
	GetMedicationMonitor(ctx context.Context, monitorID int64) (*model.MedicationMonitor, error)
	GetFollowupMedicationMonitors(ctx context.Context, followupID int64) ([]*model.MedicationMonitor, error)
	GetPatientMedicationMonitors(ctx context.Context, patientID int64, offset, limit int) ([]*model.MedicationMonitor, int64, error)
	UpdateMedicationMonitor(ctx context.Context, monitor *model.MedicationMonitor) error
	DeleteMedicationMonitor(ctx context.Context, monitorID int64) error
	GetActiveMedicationMonitors(ctx context.Context, patientID int64) ([]*model.MedicationMonitor, error)
}

type FollowupCache interface {
	// 随访计划缓存
	SetFollowupPlan(ctx context.Context, planID int64, plan *model.FollowupPlan) error
	GetFollowupPlan(ctx context.Context, planID int64) (*model.FollowupPlan, error)
	DeleteFollowupPlan(ctx context.Context, planID int64) error

	// 随访记录缓存
	SetFollowupRecord(ctx context.Context, recordID int64, record *model.FollowupRecord) error
	GetFollowupRecord(ctx context.Context, recordID int64) (*model.FollowupRecord, error)
	DeleteFollowupRecord(ctx context.Context, recordID int64) error

	// 用药监测缓存
	SetMedicationMonitor(ctx context.Context, monitorID int64, monitor *model.MedicationMonitor) error
	GetMedicationMonitor(ctx context.Context, monitorID int64) (*model.MedicationMonitor, error)
	DeleteMedicationMonitor(ctx context.Context, monitorID int64) error

	// 患者最新随访记录缓存
	SetPatientLatestFollowup(ctx context.Context, patientID int64, record *model.FollowupRecord) error
	GetPatientLatestFollowup(ctx context.Context, patientID int64) (*model.FollowupRecord, error)
	DeletePatientLatestFollowup(ctx context.Context, patientID int64) error
}
