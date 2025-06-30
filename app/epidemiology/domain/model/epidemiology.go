package model

import "time"

// EnvironmentExposure 环境暴露记录
type EnvironmentExposure struct {
	ID                   int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientID            int64     `gorm:"column:patient_id;index" json:"patient_id"`
	ResidenceType        string    `gorm:"column:residence_type;size:50" json:"residence_type"`
	BuildingMaterial     string    `gorm:"column:building_material;size:50" json:"building_material"`
	VentilationFrequency string    `gorm:"column:ventilation_frequency;size:50" json:"ventilation_frequency"`
	AirConditioningUsage string    `gorm:"column:air_conditioning_usage;type:text" json:"air_conditioning_usage"`
	PetExposure          bool      `gorm:"column:pet_exposure" json:"pet_exposure"`
	SmokingExposure      bool      `gorm:"column:smoking_exposure" json:"smoking_exposure"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (EnvironmentExposure) TableName() string {
	return "epi_environment_exposure"
}

// EnvironmentMonitor 环境监测数据
type EnvironmentMonitor struct {
	ID                  int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	LocationCode        string    `gorm:"column:location_code;size:50" json:"location_code"`
	MonitorTime         time.Time `gorm:"column:monitor_time" json:"monitor_time"`
	PM25                float64   `gorm:"column:pm25" json:"pm25"`
	Temperature         float64   `gorm:"column:temperature" json:"temperature"`
	Humidity            float64   `gorm:"column:humidity" json:"humidity"`
	PollenConcentration float64   `gorm:"column:pollen_concentration" json:"pollen_concentration"`
	AllergenLevel       string    `gorm:"column:allergen_level;type:text" json:"allergen_level"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (EnvironmentMonitor) TableName() string {
	return "epi_environment_monitor"
}

// LifestyleSurvey 生活方式调查
type LifestyleSurvey struct {
	ID                int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientID         int64     `gorm:"column:patient_id;index" json:"patient_id"`
	DietPattern       string    `gorm:"column:diet_pattern;size:50" json:"diet_pattern"`
	ExerciseFrequency string    `gorm:"column:exercise_frequency;size:50" json:"exercise_frequency"`
	SleepQuality      string    `gorm:"column:sleep_quality;size:50" json:"sleep_quality"`
	StressLevel       string    `gorm:"column:stress_level;size:50" json:"stress_level"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (LifestyleSurvey) TableName() string {
	return "epi_lifestyle_survey"
}
