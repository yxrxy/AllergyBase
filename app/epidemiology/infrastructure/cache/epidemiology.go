package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/model"
	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/repository"
)

type epidemiologyCache struct {
	rdb *redis.Client
}

func NewEpidemiologyCache(rdb *redis.Client) repository.EpidemiologyCache {
	return &epidemiologyCache{rdb: rdb}
}

const (
	exposureKeyPrefix         = "epidemiology:exposure:"
	monitorKeyPrefix          = "epidemiology:monitor:"
	surveyKeyPrefix           = "epidemiology:survey:"
	patientLatestSurveyPrefix = "epidemiology:patient_latest_survey:"
	cacheExpiration           = 30 * time.Minute
)

// 环境暴露缓存
func (e *epidemiologyCache) SetEnvironmentExposure(ctx context.Context, exposureID int64, exposure *model.EnvironmentExposure) error {
	key := fmt.Sprintf("%s%d", exposureKeyPrefix, exposureID)
	data, err := json.Marshal(exposure)
	if err != nil {
		return err
	}
	return e.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (e *epidemiologyCache) GetEnvironmentExposure(ctx context.Context, exposureID int64) (*model.EnvironmentExposure, error) {
	key := fmt.Sprintf("%s%d", exposureKeyPrefix, exposureID)
	data, err := e.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var exposure model.EnvironmentExposure
	if err := json.Unmarshal([]byte(data), &exposure); err != nil {
		return nil, err
	}

	return &exposure, nil
}

func (e *epidemiologyCache) DeleteEnvironmentExposure(ctx context.Context, exposureID int64) error {
	key := fmt.Sprintf("%s%d", exposureKeyPrefix, exposureID)
	return e.rdb.Del(ctx, key).Err()
}

// 环境监测缓存
func (e *epidemiologyCache) SetEnvironmentMonitor(ctx context.Context, monitorID int64, monitor *model.EnvironmentMonitor) error {
	key := fmt.Sprintf("%s%d", monitorKeyPrefix, monitorID)
	data, err := json.Marshal(monitor)
	if err != nil {
		return err
	}
	return e.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (e *epidemiologyCache) GetEnvironmentMonitor(ctx context.Context, monitorID int64) (*model.EnvironmentMonitor, error) {
	key := fmt.Sprintf("%s%d", monitorKeyPrefix, monitorID)
	data, err := e.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var monitor model.EnvironmentMonitor
	if err := json.Unmarshal([]byte(data), &monitor); err != nil {
		return nil, err
	}

	return &monitor, nil
}

func (e *epidemiologyCache) DeleteEnvironmentMonitor(ctx context.Context, monitorID int64) error {
	key := fmt.Sprintf("%s%d", monitorKeyPrefix, monitorID)
	return e.rdb.Del(ctx, key).Err()
}

// 生活方式调查缓存
func (e *epidemiologyCache) SetLifestyleSurvey(ctx context.Context, surveyID int64, survey *model.LifestyleSurvey) error {
	key := fmt.Sprintf("%s%d", surveyKeyPrefix, surveyID)
	data, err := json.Marshal(survey)
	if err != nil {
		return err
	}
	return e.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (e *epidemiologyCache) GetLifestyleSurvey(ctx context.Context, surveyID int64) (*model.LifestyleSurvey, error) {
	key := fmt.Sprintf("%s%d", surveyKeyPrefix, surveyID)
	data, err := e.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var survey model.LifestyleSurvey
	if err := json.Unmarshal([]byte(data), &survey); err != nil {
		return nil, err
	}

	return &survey, nil
}

func (e *epidemiologyCache) DeleteLifestyleSurvey(ctx context.Context, surveyID int64) error {
	key := fmt.Sprintf("%s%d", surveyKeyPrefix, surveyID)
	return e.rdb.Del(ctx, key).Err()
}

// 患者最新生活方式调查缓存
func (e *epidemiologyCache) SetPatientLatestSurvey(ctx context.Context, patientID int64, survey *model.LifestyleSurvey) error {
	key := fmt.Sprintf("%s%d", patientLatestSurveyPrefix, patientID)
	data, err := json.Marshal(survey)
	if err != nil {
		return err
	}
	return e.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (e *epidemiologyCache) GetPatientLatestSurvey(ctx context.Context, patientID int64) (*model.LifestyleSurvey, error) {
	key := fmt.Sprintf("%s%d", patientLatestSurveyPrefix, patientID)
	data, err := e.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var survey model.LifestyleSurvey
	if err := json.Unmarshal([]byte(data), &survey); err != nil {
		return nil, err
	}

	return &survey, nil
}

func (e *epidemiologyCache) DeletePatientLatestSurvey(ctx context.Context, patientID int64) error {
	key := fmt.Sprintf("%s%d", patientLatestSurveyPrefix, patientID)
	return e.rdb.Del(ctx, key).Err()
}
