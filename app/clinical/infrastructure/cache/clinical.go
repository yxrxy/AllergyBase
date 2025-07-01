package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yxrxy/AllergyBase/app/clinical/domain/model"
	"github.com/yxrxy/AllergyBase/app/clinical/domain/repository"
)

type clinicalCache struct {
	rdb *redis.Client
}

func NewClinicalCache(rdb *redis.Client) repository.ClinicalCache {
	return &clinicalCache{rdb: rdb}
}

const (
	patientKeyPrefix = "clinical:patient:"
	visitKeyPrefix   = "clinical:visit:"
	cacheExpiration  = 30 * time.Minute
)

// 患者信息缓存
func (c *clinicalCache) SetPatientInfo(ctx context.Context, patientID int64, patient *model.Patient) error {
	key := fmt.Sprintf("%s%d", patientKeyPrefix, patientID)
	data, err := json.Marshal(patient)
	if err != nil {
		return err
	}
	return c.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (c *clinicalCache) GetPatientInfo(ctx context.Context, patientID int64) (*model.Patient, error) {
	key := fmt.Sprintf("%s%d", patientKeyPrefix, patientID)
	data, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var patient model.Patient
	if err := json.Unmarshal([]byte(data), &patient); err != nil {
		return nil, err
	}

	return &patient, nil
}

func (c *clinicalCache) DeletePatientInfo(ctx context.Context, patientID int64) error {
	key := fmt.Sprintf("%s%d", patientKeyPrefix, patientID)
	return c.rdb.Del(ctx, key).Err()
}

// 就诊信息缓存
func (c *clinicalCache) SetVisitInfo(ctx context.Context, visitID int64, visit *model.Visit) error {
	key := fmt.Sprintf("%s%d", visitKeyPrefix, visitID)
	data, err := json.Marshal(visit)
	if err != nil {
		return err
	}
	return c.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (c *clinicalCache) GetVisitInfo(ctx context.Context, visitID int64) (*model.Visit, error) {
	key := fmt.Sprintf("%s%d", visitKeyPrefix, visitID)
	data, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var visit model.Visit
	if err := json.Unmarshal([]byte(data), &visit); err != nil {
		return nil, err
	}

	return &visit, nil
}

func (c *clinicalCache) DeleteVisitInfo(ctx context.Context, visitID int64) error {
	key := fmt.Sprintf("%s%d", visitKeyPrefix, visitID)
	return c.rdb.Del(ctx, key).Err()
}
