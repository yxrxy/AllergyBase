package mysql

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/followup/domain/model"
	"github.com/yxrxy/AllergyBase/app/followup/domain/repository"
	"gorm.io/gorm"
)

type followupDB struct {
	db *gorm.DB
}

func NewFollowupDB(db *gorm.DB) repository.FollowupDB {
	return &followupDB{db: db}
}

// 随访计划管理
func (f *followupDB) CreateFollowupPlan(ctx context.Context, plan *model.FollowupPlan) (*model.FollowupPlan, error) {
	if err := f.db.WithContext(ctx).Create(plan).Error; err != nil {
		return nil, err
	}
	return plan, nil
}

func (f *followupDB) GetFollowupPlan(ctx context.Context, planID int64) (*model.FollowupPlan, error) {
	var plan model.FollowupPlan
	if err := f.db.WithContext(ctx).First(&plan, planID).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

func (f *followupDB) GetPatientFollowupPlans(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupPlan, int64, error) {
	var plans []*model.FollowupPlan
	var total int64

	if err := f.db.WithContext(ctx).Model(&model.FollowupPlan{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := f.db.WithContext(ctx).Where("patient_id = ?", patientID).
		Order("created_at DESC").Limit(limit).Offset(offset).Find(&plans).Error; err != nil {
		return nil, 0, err
	}

	return plans, total, nil
}

func (f *followupDB) UpdateFollowupPlan(ctx context.Context, plan *model.FollowupPlan) error {
	return f.db.WithContext(ctx).Save(plan).Error
}

func (f *followupDB) DeleteFollowupPlan(ctx context.Context, planID int64) error {
	return f.db.WithContext(ctx).Delete(&model.FollowupPlan{}, planID).Error
}

func (f *followupDB) GetActiveFollowupPlans(ctx context.Context, offset, limit int) ([]*model.FollowupPlan, int64, error) {
	var plans []*model.FollowupPlan
	var total int64

	query := f.db.WithContext(ctx).Model(&model.FollowupPlan{}).Where("status = ?", "active")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&plans).Error; err != nil {
		return nil, 0, err
	}

	return plans, total, nil
}

// 随访记录管理
func (f *followupDB) CreateFollowupRecord(ctx context.Context, record *model.FollowupRecord) (*model.FollowupRecord, error) {
	if err := f.db.WithContext(ctx).Create(record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (f *followupDB) GetFollowupRecord(ctx context.Context, recordID int64) (*model.FollowupRecord, error) {
	var record model.FollowupRecord
	if err := f.db.WithContext(ctx).Preload("Plan").First(&record, recordID).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (f *followupDB) GetPlanFollowupRecords(ctx context.Context, planID int64, offset, limit int) ([]*model.FollowupRecord, int64, error) {
	var records []*model.FollowupRecord
	var total int64

	if err := f.db.WithContext(ctx).Model(&model.FollowupRecord{}).Where("plan_id = ?", planID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := f.db.WithContext(ctx).Where("plan_id = ?", planID).
		Preload("Plan").Order("followup_date DESC").Limit(limit).Offset(offset).Find(&records).Error; err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

func (f *followupDB) GetPatientFollowupRecords(ctx context.Context, patientID int64, offset, limit int) ([]*model.FollowupRecord, int64, error) {
	var records []*model.FollowupRecord
	var total int64

	if err := f.db.WithContext(ctx).Model(&model.FollowupRecord{}).Where("patient_id = ?", patientID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := f.db.WithContext(ctx).Where("patient_id = ?", patientID).
		Preload("Plan").Order("followup_date DESC").Limit(limit).Offset(offset).Find(&records).Error; err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

func (f *followupDB) UpdateFollowupRecord(ctx context.Context, record *model.FollowupRecord) error {
	return f.db.WithContext(ctx).Save(record).Error
}

func (f *followupDB) DeleteFollowupRecord(ctx context.Context, recordID int64) error {
	return f.db.WithContext(ctx).Delete(&model.FollowupRecord{}, recordID).Error
}

func (f *followupDB) GetLatestFollowupRecord(ctx context.Context, patientID int64) (*model.FollowupRecord, error) {
	var record model.FollowupRecord
	if err := f.db.WithContext(ctx).Where("patient_id = ?", patientID).
		Preload("Plan").Order("followup_date DESC").First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// 用药监测管理
func (f *followupDB) CreateMedicationMonitor(ctx context.Context, monitor *model.MedicationMonitor) (*model.MedicationMonitor, error) {
	if err := f.db.WithContext(ctx).Create(monitor).Error; err != nil {
		return nil, err
	}
	return monitor, nil
}

func (f *followupDB) GetMedicationMonitor(ctx context.Context, monitorID int64) (*model.MedicationMonitor, error) {
	var monitor model.MedicationMonitor
	if err := f.db.WithContext(ctx).Preload("Followup").First(&monitor, monitorID).Error; err != nil {
		return nil, err
	}
	return &monitor, nil
}

func (f *followupDB) GetFollowupMedicationMonitors(ctx context.Context, followupID int64) ([]*model.MedicationMonitor, error) {
	var monitors []*model.MedicationMonitor
	if err := f.db.WithContext(ctx).Where("followup_id = ?", followupID).
		Preload("Followup").Find(&monitors).Error; err != nil {
		return nil, err
	}
	return monitors, nil
}

func (f *followupDB) GetPatientMedicationMonitors(ctx context.Context, patientID int64, offset, limit int) ([]*model.MedicationMonitor, int64, error) {
	var monitors []*model.MedicationMonitor
	var total int64

	// 通过关联随访记录查询患者的用药监测
	subQuery := f.db.Select("id").Where("patient_id = ?", patientID).Table("followup_records")

	if err := f.db.WithContext(ctx).Model(&model.MedicationMonitor{}).
		Where("followup_id IN (?)", subQuery).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := f.db.WithContext(ctx).Where("followup_id IN (?)", subQuery).
		Preload("Followup").Order("created_at DESC").Limit(limit).Offset(offset).Find(&monitors).Error; err != nil {
		return nil, 0, err
	}

	return monitors, total, nil
}

func (f *followupDB) UpdateMedicationMonitor(ctx context.Context, monitor *model.MedicationMonitor) error {
	return f.db.WithContext(ctx).Save(monitor).Error
}

func (f *followupDB) DeleteMedicationMonitor(ctx context.Context, monitorID int64) error {
	return f.db.WithContext(ctx).Delete(&model.MedicationMonitor{}, monitorID).Error
}

func (f *followupDB) GetActiveMedicationMonitors(ctx context.Context, patientID int64) ([]*model.MedicationMonitor, error) {
	var monitors []*model.MedicationMonitor

	// 通过关联随访记录查询患者的活跃用药监测
	subQuery := f.db.Select("id").Where("patient_id = ?", patientID).Table("followup_records")

	if err := f.db.WithContext(ctx).Where("followup_id IN (?) AND (end_date IS NULL OR end_date > NOW())", subQuery).
		Preload("Followup").Order("start_date DESC").Find(&monitors).Error; err != nil {
		return nil, err
	}

	return monitors, nil
}
