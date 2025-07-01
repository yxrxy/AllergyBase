package model

import "time"

// Patient 患者基本信息
type Patient struct {
	ID                   int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientNo            string     `gorm:"uniqueIndex;size:50;not null" json:"patient_no"`
	Name                 string     `gorm:"size:50" json:"name"`
	Gender               string     `gorm:"type:tinyint" json:"gender"`
	BirthDate            *time.Time `gorm:"type:date" json:"birth_date"`
	IdCard               string     `gorm:"size:18" json:"id_card"`
	Phone                string     `gorm:"size:20" json:"phone"`
	Address              string     `gorm:"size:200" json:"address"`
	EmergencyContact     string     `gorm:"size:100" json:"emergency_contact"`
	EmergencyPhone       string     `gorm:"size:20" json:"emergency_phone"`
	Height               float64    `gorm:"type:decimal(5,2)" json:"height"`
	Weight               float64    `gorm:"type:decimal(5,2)" json:"weight"`
	BirthWeight          float64    `gorm:"type:decimal(5,2)" json:"birth_weight"`
	MedicalInsuranceType int        `gorm:"type:tinyint" json:"medical_insurance_type"`
	MedicalInsuranceNo   string     `gorm:"size:50" json:"medical_insurance_no"`
	CreatedAt            time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt            *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName 指定患者表名
func (Patient) TableName() string {
	return "clinical_patient_info"
}

// Visit 就诊记录
type Visit struct {
	ID             int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientID      int64      `gorm:"index" json:"patient_id"`
	VisitNo        string     `gorm:"uniqueIndex;size:50" json:"visit_no"`
	VisitTime      time.Time  `json:"visit_time"`
	Department     string     `gorm:"size:50" json:"department"`
	VisitType      int        `gorm:"type:tinyint" json:"visit_type"`
	DoctorID       *int64     `gorm:"column:doctor_id" json:"doctor_id"`
	ChiefComplaint string     `gorm:"type:text" json:"chief_complaint"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Patient        *Patient   `gorm:"foreignKey:PatientID" json:"patient"`
}

// TableName 指定就诊记录表名
func (Visit) TableName() string {
	return "clinical_visit_record"
}

// Diagnosis 诊断记录
type Diagnosis struct {
	ID            int64  `gorm:"primaryKey;autoIncrement"`
	VisitID       int64  `gorm:"index"`
	DiagnosisCode string `gorm:"column:icd_code;size:50"`
	DiagnosisType string `gorm:"size:50"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// 关联
	Visit *Visit `gorm:"foreignKey:VisitID"`
}

// TableName 指定诊断记录表名
func (Diagnosis) TableName() string {
	return "clinical_diagnosis"
}

// Examination 检查记录
type Examination struct {
	ID             int64  `gorm:"primaryKey;autoIncrement"`
	VisitID        int64  `gorm:"index"`
	ExamType       string `gorm:"size:100"`
	ExamResult     string `gorm:"type:text"`
	ReferenceValue string `gorm:"column:reference_range;size:200"`
	Unit           string `gorm:"column:result_unit;size:50"`
	ExamTime       time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time

	// 关联
	Visit *Visit `gorm:"foreignKey:VisitID"`
}

// TableName 指定检查记录表名
func (Examination) TableName() string {
	return "clinical_examination"
}
