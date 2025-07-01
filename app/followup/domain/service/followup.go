package service

import (
	"context"
	"time"

	"github.com/yxrxy/AllergyBase/app/followup/domain/model"
	"github.com/yxrxy/AllergyBase/app/followup/domain/repository"
)

type FollowupService struct {
	db    repository.FollowupDB
	cache repository.FollowupCache
}

func NewFollowupService(db repository.FollowupDB, cache repository.FollowupCache) *FollowupService {
	return &FollowupService{
		db:    db,
		cache: cache,
	}
}

// 随访计划相关方法
func (s *FollowupService) CreateFollowupPlan(ctx context.Context, plan *model.FollowupPlan) (*model.FollowupPlan, error) {
	// 设置创建时间
	plan.CreatedAt = time.Now()
	plan.UpdatedAt = time.Now()

	createdPlan, err := s.db.CreateFollowupPlan(ctx, plan)
	if err != nil {
		return nil, err
	}

	// 缓存随访计划
	_ = s.cache.SetFollowupPlan(ctx, createdPlan.ID, createdPlan)

	return createdPlan, nil
}

func (s *FollowupService) GetFollowupPlan(ctx context.Context, planID int64) (*model.FollowupPlan, error) {
	// 先从缓存获取
	plan, err := s.cache.GetFollowupPlan(ctx, planID)
	if err == nil && plan != nil {
		return plan, nil
	}

	// 从数据库获取
	plan, err = s.db.GetFollowupPlan(ctx, planID)
	if err != nil {
		return nil, err
	}

	// 缓存随访计划
	_ = s.cache.SetFollowupPlan(ctx, planID, plan)

	return plan, nil
}

func (s *FollowupService) GetPatientFollowupPlans(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupPlan, int64, error) {
	return s.db.GetPatientFollowupPlans(ctx, patientID, offset, limit)
}

func (s *FollowupService) GetActiveFollowupPlans(ctx context.Context, offset, limit int) ([]*model.FollowupPlan, int64, error) {
	return s.db.GetActiveFollowupPlans(ctx, offset, limit)
}

// 随访记录相关方法
func (s *FollowupService) CreateFollowupRecord(ctx context.Context, record *model.FollowupRecord) (*model.FollowupRecord, error) {
	// 设置创建时间
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()

	createdRecord, err := s.db.CreateFollowupRecord(ctx, record)
	if err != nil {
		return nil, err
	}

	// 缓存随访记录
	_ = s.cache.SetFollowupRecord(ctx, createdRecord.ID, createdRecord)

	// 更新患者最新随访记录缓存
	_ = s.cache.SetPatientLatestFollowup(ctx, record.PatientID, createdRecord)

	return createdRecord, nil
}

func (s *FollowupService) GetFollowupRecord(ctx context.Context, recordID int64) (*model.FollowupRecord, error) {
	// 先从缓存获取
	record, err := s.cache.GetFollowupRecord(ctx, recordID)
	if err == nil && record != nil {
		return record, nil
	}

	// 从数据库获取
	record, err = s.db.GetFollowupRecord(ctx, recordID)
	if err != nil {
		return nil, err
	}

	// 缓存随访记录
	_ = s.cache.SetFollowupRecord(ctx, recordID, record)

	return record, nil
}

func (s *FollowupService) GetPlanFollowupRecords(ctx context.Context, planID int64, offset, limit int) ([]*model.FollowupRecord, int64, error) {
	return s.db.GetPlanFollowupRecords(ctx, planID, offset, limit)
}

func (s *FollowupService) GetPatientFollowupRecords(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupRecord, int64, error) {
	return s.db.GetPatientFollowupRecords(ctx, patientID, offset, limit)
}

func (s *FollowupService) GetLatestFollowupRecord(ctx context.Context, patientID int64) (*model.FollowupRecord, error) {
	// 先从缓存获取最新随访记录
	record, err := s.cache.GetPatientLatestFollowup(ctx, patientID)
	if err == nil && record != nil {
		return record, nil
	}

	// 从数据库获取
	record, err = s.db.GetLatestFollowupRecord(ctx, patientID)
	if err != nil {
		return nil, err
	}

	// 缓存患者最新随访记录
	_ = s.cache.SetPatientLatestFollowup(ctx, patientID, record)

	return record, nil
}

// 用药监测相关方法
func (s *FollowupService) CreateMedicationMonitor(ctx context.Context, monitor *model.MedicationMonitor) (*model.MedicationMonitor, error) {
	// 设置创建时间
	monitor.CreatedAt = time.Now()
	monitor.UpdatedAt = time.Now()

	createdMonitor, err := s.db.CreateMedicationMonitor(ctx, monitor)
	if err != nil {
		return nil, err
	}

	// 缓存用药监测信息
	_ = s.cache.SetMedicationMonitor(ctx, createdMonitor.ID, createdMonitor)

	return createdMonitor, nil
}

func (s *FollowupService) GetMedicationMonitor(ctx context.Context, monitorID int64) (*model.MedicationMonitor, error) {
	// 先从缓存获取
	monitor, err := s.cache.GetMedicationMonitor(ctx, monitorID)
	if err == nil && monitor != nil {
		return monitor, nil
	}

	// 从数据库获取
	monitor, err = s.db.GetMedicationMonitor(ctx, monitorID)
	if err != nil {
		return nil, err
	}

	// 缓存用药监测信息
	_ = s.cache.SetMedicationMonitor(ctx, monitorID, monitor)

	return monitor, nil
}

func (s *FollowupService) GetFollowupMedicationMonitors(ctx context.Context, followupID int64) ([]*model.MedicationMonitor, error) {
	return s.db.GetFollowupMedicationMonitors(ctx, followupID)
}

func (s *FollowupService) GetPatientMedicationMonitors(ctx context.Context, patientID int64, offset, limit int) ([]*model.MedicationMonitor, int64, error) {
	return s.db.GetPatientMedicationMonitors(ctx, patientID, offset, limit)
}

func (s *FollowupService) GetActiveMedicationMonitors(ctx context.Context, patientID int64) ([]*model.MedicationMonitor, error) {
	return s.db.GetActiveMedicationMonitors(ctx, patientID)
}
