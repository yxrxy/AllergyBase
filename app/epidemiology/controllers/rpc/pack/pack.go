package pack

import (
	"time"

	"github.com/yxrxy/AllergyBase/app/epidemiology/domain/model"
	rpcmodel "github.com/yxrxy/AllergyBase/kitex_gen/model"
)

// EnvironmentExposure 将数据库环境暴露实体转换为RPC响应实体
func EnvironmentExposure(e *model.EnvironmentExposure) *rpcmodel.EnvironmentExposure {
	if e == nil {
		return nil
	}

	return &rpcmodel.EnvironmentExposure{
		Id:                   e.ID,
		PatientId:            e.PatientID,
		ResidenceType:        &e.ResidenceType,
		BuildingMaterial:     &e.BuildingMaterial,
		VentilationFrequency: &e.VentilationFrequency,
		AirConditioningUsage: &e.AirConditioningUsage,
		PetExposure:          &e.PetExposure,
		SmokingExposure:      &e.SmokingExposure,
	}
}

// EnvironmentMonitor 将数据库环境监测实体转换为RPC响应实体
func EnvironmentMonitor(m *model.EnvironmentMonitor) *rpcmodel.EnvironmentMonitor {
	if m == nil {
		return nil
	}
	return &rpcmodel.EnvironmentMonitor{
		Id:                  m.ID,
		LocationCode:        m.LocationCode,
		MonitorTime:         m.MonitorTime.Format("2006-01-02 15:04:05"),
		Pm25:                &m.PM25,
		Temperature:         &m.Temperature,
		Humidity:            &m.Humidity,
		PollenConcentration: &m.PollenConcentration,
		AllergenLevel:       &m.AllergenLevel,
	}
}

// LifestyleSurvey 将数据库生活方式调查实体转换为RPC响应实体
func LifestyleSurvey(l *model.LifestyleSurvey) *rpcmodel.LifestyleSurvey {
	if l == nil {
		return nil
	}
	return &rpcmodel.LifestyleSurvey{
		Id:                l.ID,
		PatientId:         l.PatientID,
		DietPattern:       &l.DietPattern,
		ExerciseFrequency: &l.ExerciseFrequency,
		SleepQuality:      &l.SleepQuality,
		StressLevel:       &l.StressLevel,
	}
}

// EnvironmentExposureList 将环境暴露列表转换为RPC响应实体
func EnvironmentExposureList(exposures []*model.EnvironmentExposure) []*rpcmodel.EnvironmentExposure {
	if len(exposures) == 0 {
		return nil
	}
	result := make([]*rpcmodel.EnvironmentExposure, 0, len(exposures))
	for _, e := range exposures {
		result = append(result, EnvironmentExposure(e))
	}
	return result
}

// EnvironmentMonitorList 将环境监测列表转换为RPC响应实体
func EnvironmentMonitorList(monitors []*model.EnvironmentMonitor) []*rpcmodel.EnvironmentMonitor {
	if len(monitors) == 0 {
		return nil
	}
	result := make([]*rpcmodel.EnvironmentMonitor, 0, len(monitors))
	for _, m := range monitors {
		result = append(result, EnvironmentMonitor(m))
	}
	return result
}

// LifestyleSurveyList 将生活方式调查列表转换为RPC响应实体
func LifestyleSurveyList(surveys []*model.LifestyleSurvey) []*rpcmodel.LifestyleSurvey {
	if len(surveys) == 0 {
		return nil
	}
	result := make([]*rpcmodel.LifestyleSurvey, 0, len(surveys))
	for _, s := range surveys {
		result = append(result, LifestyleSurvey(s))
	}
	return result
}

// BuildEnvironmentExposureFromRequest 将RPC请求转换为数据库实体
func BuildEnvironmentExposureFromRequest(req *rpcmodel.EnvironmentExposure) *model.EnvironmentExposure {
	if req == nil {
		return nil
	}

	exposure := &model.EnvironmentExposure{
		PatientID:            req.PatientId,
		ResidenceType:        req.GetResidenceType(),
		BuildingMaterial:     req.GetBuildingMaterial(),
		VentilationFrequency: req.GetVentilationFrequency(),
		AirConditioningUsage: req.GetAirConditioningUsage(),
		PetExposure:          req.GetPetExposure(),
		SmokingExposure:      req.GetSmokingExposure(),
	}

	// 只有在ID不为0时才设置ID（用于更新操作）
	if req.Id != 0 {
		exposure.ID = req.Id
	}

	return exposure
}

// BuildEnvironmentMonitorFromRequest 将RPC请求转换为数据库实体
func BuildEnvironmentMonitorFromRequest(req *rpcmodel.EnvironmentMonitor) *model.EnvironmentMonitor {
	if req == nil {
		return nil
	}

	var monitorTime time.Time
	if req.MonitorTime != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", req.MonitorTime); err == nil {
			monitorTime = t
		}
	}

	return &model.EnvironmentMonitor{
		ID:                  req.Id,
		LocationCode:        req.LocationCode,
		MonitorTime:         monitorTime,
		PM25:                req.GetPm25(),
		Temperature:         req.GetTemperature(),
		Humidity:            req.GetHumidity(),
		PollenConcentration: req.GetPollenConcentration(),
		AllergenLevel:       req.GetAllergenLevel(),
	}
}

// BuildLifestyleSurveyFromRequest 将RPC请求转换为数据库实体
func BuildLifestyleSurveyFromRequest(req *rpcmodel.LifestyleSurvey) *model.LifestyleSurvey {
	if req == nil {
		return nil
	}
	return &model.LifestyleSurvey{
		ID:                req.Id,
		PatientID:         req.PatientId,
		DietPattern:       req.GetDietPattern(),
		ExerciseFrequency: req.GetExerciseFrequency(),
		SleepQuality:      req.GetSleepQuality(),
		StressLevel:       req.GetStressLevel(),
	}
}
