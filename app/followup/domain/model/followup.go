package model

import "time"

// FollowupPlan 随访计划
type FollowupPlan struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	PatientID int64  `gorm:"index"`
	PlanName  string `gorm:"size:200"`
	PlanType  string `gorm:"size:50"`
	StartDate time.Time
	EndDate   time.Time
	Frequency string `gorm:"size:100"`
	Status    string `gorm:"size:50"`
	CreatedBy string `gorm:"size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FollowupRecord 随访记录
type FollowupRecord struct {
	ID              int64 `gorm:"primaryKey;autoIncrement"`
	PlanID          int64 `gorm:"index"`
	PatientID       int64 `gorm:"index"`
	FollowupDate    time.Time
	FollowupMethod  string `gorm:"size:50"`
	Symptoms        string `gorm:"type:text"`
	TreatmentEffect string `gorm:"type:text"`
	NextPlan        string `gorm:"type:text"`
	FollowupDoctor  string `gorm:"size:50"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// 关联
	Plan *FollowupPlan `gorm:"foreignKey:PlanID"`
}

// MedicationMonitor 用药监测
type MedicationMonitor struct {
	ID             int64  `gorm:"primaryKey;autoIncrement"`
	FollowupID     int64  `gorm:"index"`
	MedicationName string `gorm:"size:200"`
	Dosage         string `gorm:"size:100"`
	Frequency      string `gorm:"size:100"`
	StartDate      time.Time
	EndDate        *time.Time
	SideEffects    string `gorm:"type:text"`
	Compliance     string `gorm:"size:50"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	// 关联
	Followup *FollowupRecord `gorm:"foreignKey:FollowupID"`
}
