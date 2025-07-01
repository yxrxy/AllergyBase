package usecase

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/followup/domain/model"
	"github.com/yxrxy/AllergyBase/app/followup/domain/repository"
	"github.com/yxrxy/AllergyBase/app/followup/domain/service"
)

type FollowupUseCase interface {
	// 随访计划相关
	CreateFollowupPlan(ctx context.Context, plan *model.FollowupPlan) (*model.FollowupPlan, error)
	GetFollowupPlan(ctx context.Context, planID int64) (*model.FollowupPlan, error)
	GetPatientFollowupPlans(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupPlan, int64, error)
	UpdateFollowupPlan(ctx context.Context, plan *model.FollowupPlan) error
	DeleteFollowupPlan(ctx context.Context, planID int64) error
	GetActiveFollowupPlans(ctx context.Context, offset, limit int) ([]*model.FollowupPlan, int64, error)

	// 随访记录相关
	CreateFollowupRecord(ctx context.Context, record *model.FollowupRecord) (*model.FollowupRecord, error)
	GetFollowupRecord(ctx context.Context, recordID int64) (*model.FollowupRecord, error)
	GetPlanFollowupRecords(ctx context.Context, planID int64, offset, limit int) ([]*model.FollowupRecord, int64, error)
	GetPatientFollowupRecords(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupRecord, int64, error)
	UpdateFollowupRecord(ctx context.Context, record *model.FollowupRecord) error
	DeleteFollowupRecord(ctx context.Context, recordID int64) error
	GetLatestFollowupRecord(ctx context.Context, patientID int64) (*model.FollowupRecord, error)

	// 用药监测相关
	CreateMedicationMonitor(ctx context.Context, monitor *model.MedicationMonitor) (*model.MedicationMonitor, error)
	GetMedicationMonitor(ctx context.Context, monitorID int64) (*model.MedicationMonitor, error)
	GetFollowupMedicationMonitors(ctx context.Context, followupID int64) ([]*model.MedicationMonitor, error)
	GetPatientMedicationMonitors(ctx context.Context, patientID int64, offset, limit int) ([]*model.MedicationMonitor, int64, error)
	UpdateMedicationMonitor(ctx context.Context, monitor *model.MedicationMonitor) error
	DeleteMedicationMonitor(ctx context.Context, monitorID int64) error
	GetActiveMedicationMonitors(ctx context.Context, patientID int64) ([]*model.MedicationMonitor, error)
}

type useCase struct {
	db      repository.FollowupDB
	service *service.FollowupService
}

func NewFollowupUseCase(db repository.FollowupDB, service *service.FollowupService) FollowupUseCase {
	return &useCase{
		db:      db,
		service: service,
	}
}

// 随访计划相关实现
func (u *useCase) CreateFollowupPlan(ctx context.Context, plan *model.FollowupPlan) (*model.FollowupPlan, error) {
	return u.service.CreateFollowupPlan(ctx, plan)
}

func (u *useCase) GetFollowupPlan(ctx context.Context, planID int64) (*model.FollowupPlan, error) {
	return u.service.GetFollowupPlan(ctx, planID)
}

func (u *useCase) GetPatientFollowupPlans(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupPlan, int64, error) {
	return u.service.GetPatientFollowupPlans(ctx, patientID, offset, limit)
}

func (u *useCase) UpdateFollowupPlan(ctx context.Context, plan *model.FollowupPlan) error {
	return u.db.UpdateFollowupPlan(ctx, plan)
}

func (u *useCase) DeleteFollowupPlan(ctx context.Context, planID int64) error {
	return u.db.DeleteFollowupPlan(ctx, planID)
}

func (u *useCase) GetActiveFollowupPlans(ctx context.Context, offset, limit int) ([]*model.FollowupPlan, int64, error) {
	return u.service.GetActiveFollowupPlans(ctx, offset, limit)
}

// 随访记录相关实现
func (u *useCase) CreateFollowupRecord(ctx context.Context, record *model.FollowupRecord) (*model.FollowupRecord, error) {
	return u.service.CreateFollowupRecord(ctx, record)
}

func (u *useCase) GetFollowupRecord(ctx context.Context, recordID int64) (*model.FollowupRecord, error) {
	return u.service.GetFollowupRecord(ctx, recordID)
}

func (u *useCase) GetPlanFollowupRecords(ctx context.Context, planID int64, offset, limit int) ([]*model.FollowupRecord, int64, error) {
	return u.service.GetPlanFollowupRecords(ctx, planID, offset, limit)
}

func (u *useCase) GetPatientFollowupRecords(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupRecord, int64, error) {
	return u.service.GetPatientFollowupRecords(ctx, patientID, offset, limit)
}

func (u *useCase) UpdateFollowupRecord(ctx context.Context, record *model.FollowupRecord) error {
	return u.db.UpdateFollowupRecord(ctx, record)
}

func (u *useCase) DeleteFollowupRecord(ctx context.Context, recordID int64) error {
	return u.db.DeleteFollowupRecord(ctx, recordID)
}

func (u *useCase) GetLatestFollowupRecord(ctx context.Context, patientID int64) (*model.FollowupRecord, error) {
	return u.service.GetLatestFollowupRecord(ctx, patientID)
}

// 用药监测相关实现
func (u *useCase) CreateMedicationMonitor(ctx context.Context, monitor *model.MedicationMonitor) (*model.MedicationMonitor, error) {
	return u.service.CreateMedicationMonitor(ctx, monitor)
}

func (u *useCase) GetMedicationMonitor(ctx context.Context, monitorID int64) (*model.MedicationMonitor, error) {
	return u.service.GetMedicationMonitor(ctx, monitorID)
}

func (u *useCase) GetFollowupMedicationMonitors(ctx context.Context, followupID int64) ([]*model.MedicationMonitor, error) {
	return u.service.GetFollowupMedicationMonitors(ctx, followupID)
}

func (u *useCase) GetPatientMedicationMonitors(ctx context.Context, patientID int64, offset, limit int) ([]*model.MedicationMonitor, int64, error) {
	return u.service.GetPatientMedicationMonitors(ctx, patientID, offset, limit)
}

func (u *useCase) UpdateMedicationMonitor(ctx context.Context, monitor *model.MedicationMonitor) error {
	return u.db.UpdateMedicationMonitor(ctx, monitor)
}

func (u *useCase) DeleteMedicationMonitor(ctx context.Context, monitorID int64) error {
	return u.db.DeleteMedicationMonitor(ctx, monitorID)
}

func (u *useCase) GetActiveMedicationMonitors(ctx context.Context, patientID int64) ([]*model.MedicationMonitor, error) {
	return u.service.GetActiveMedicationMonitors(ctx, patientID)
}
