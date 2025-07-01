package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yxrxy/AllergyBase/app/followup/domain/model"
	"github.com/yxrxy/AllergyBase/app/followup/domain/repository"
)

type followupCache struct {
	rdb *redis.Client
}

func NewFollowupCache(rdb *redis.Client) repository.FollowupCache {
	return &followupCache{rdb: rdb}
}

const (
	planKeyPrefix               = "followup:plan:"
	recordKeyPrefix             = "followup:record:"
	medicationKeyPrefix         = "followup:medication:"
	patientLatestFollowupPrefix = "followup:patient_latest:"
	cacheExpiration             = 30 * time.Minute
)

// 随访计划缓存
func (f *followupCache) SetFollowupPlan(ctx context.Context, planID int64, plan *model.FollowupPlan) error {
	key := fmt.Sprintf("%s%d", planKeyPrefix, planID)
	data, err := json.Marshal(plan)
	if err != nil {
		return err
	}
	return f.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (f *followupCache) GetFollowupPlan(ctx context.Context, planID int64) (*model.FollowupPlan, error) {
	key := fmt.Sprintf("%s%d", planKeyPrefix, planID)
	data, err := f.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var plan model.FollowupPlan
	if err := json.Unmarshal([]byte(data), &plan); err != nil {
		return nil, err
	}

	return &plan, nil
}

func (f *followupCache) DeleteFollowupPlan(ctx context.Context, planID int64) error {
	key := fmt.Sprintf("%s%d", planKeyPrefix, planID)
	return f.rdb.Del(ctx, key).Err()
}

// 随访记录缓存
func (f *followupCache) SetFollowupRecord(ctx context.Context, recordID int64, record *model.FollowupRecord) error {
	key := fmt.Sprintf("%s%d", recordKeyPrefix, recordID)
	data, err := json.Marshal(record)
	if err != nil {
		return err
	}
	return f.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (f *followupCache) GetFollowupRecord(ctx context.Context, recordID int64) (*model.FollowupRecord, error) {
	key := fmt.Sprintf("%s%d", recordKeyPrefix, recordID)
	data, err := f.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var record model.FollowupRecord
	if err := json.Unmarshal([]byte(data), &record); err != nil {
		return nil, err
	}

	return &record, nil
}

func (f *followupCache) DeleteFollowupRecord(ctx context.Context, recordID int64) error {
	key := fmt.Sprintf("%s%d", recordKeyPrefix, recordID)
	return f.rdb.Del(ctx, key).Err()
}

// 用药监测缓存
func (f *followupCache) SetMedicationMonitor(ctx context.Context, monitorID int64, monitor *model.MedicationMonitor) error {
	key := fmt.Sprintf("%s%d", medicationKeyPrefix, monitorID)
	data, err := json.Marshal(monitor)
	if err != nil {
		return err
	}
	return f.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (f *followupCache) GetMedicationMonitor(ctx context.Context, monitorID int64) (*model.MedicationMonitor, error) {
	key := fmt.Sprintf("%s%d", medicationKeyPrefix, monitorID)
	data, err := f.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var monitor model.MedicationMonitor
	if err := json.Unmarshal([]byte(data), &monitor); err != nil {
		return nil, err
	}

	return &monitor, nil
}

func (f *followupCache) DeleteMedicationMonitor(ctx context.Context, monitorID int64) error {
	key := fmt.Sprintf("%s%d", medicationKeyPrefix, monitorID)
	return f.rdb.Del(ctx, key).Err()
}

// 患者最新随访记录缓存
func (f *followupCache) SetPatientLatestFollowup(ctx context.Context, patientID int64, record *model.FollowupRecord) error {
	key := fmt.Sprintf("%s%d", patientLatestFollowupPrefix, patientID)
	data, err := json.Marshal(record)
	if err != nil {
		return err
	}
	return f.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (f *followupCache) GetPatientLatestFollowup(ctx context.Context, patientID int64) (*model.FollowupRecord, error) {
	key := fmt.Sprintf("%s%d", patientLatestFollowupPrefix, patientID)
	data, err := f.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var record model.FollowupRecord
	if err := json.Unmarshal([]byte(data), &record); err != nil {
		return nil, err
	}

	return &record, nil
}

func (f *followupCache) DeletePatientLatestFollowup(ctx context.Context, patientID int64) error {
	key := fmt.Sprintf("%s%d", patientLatestFollowupPrefix, patientID)
	return f.rdb.Del(ctx, key).Err()
}
