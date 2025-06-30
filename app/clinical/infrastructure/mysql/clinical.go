package mysql

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/clinical/domain/model"
	"github.com/yxrxy/AllergyBase/app/clinical/domain/repository"
	"gorm.io/gorm"
)

type clinicalDB struct {
	db *gorm.DB
}

func NewClinicalDB(db *gorm.DB) repository.ClinicalDB {
	return &clinicalDB{db: db}
}

// 患者管理
func (c *clinicalDB) CreatePatient(ctx context.Context, patient *model.Patient) (int64, error) {
	if err := c.db.WithContext(ctx).Create(patient).Error; err != nil {
		return 0, err
	}
	return patient.ID, nil
}

func (c *clinicalDB) GetPatient(ctx context.Context, patientID int64) (*model.Patient, error) {
	var patient model.Patient
	if err := c.db.WithContext(ctx).First(&patient, patientID).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

func (c *clinicalDB) GetPatientByNo(ctx context.Context, patientNo string) (*model.Patient, error) {
	var patient model.Patient
	if err := c.db.WithContext(ctx).Where("patient_no = ?", patientNo).First(&patient).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

func (c *clinicalDB) ListPatients(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Patient, int64, error) {
	var patients []*model.Patient
	var total int64

	query := c.db.WithContext(ctx).Model(&model.Patient{})

	// 添加搜索条件
	if patientNo, ok := filters["patientNo"]; ok && patientNo != "" {
		query = query.Where("patient_no LIKE ?", "%"+patientNo+"%")
	}
	if name, ok := filters["name"]; ok && name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if idCard, ok := filters["idCard"]; ok && idCard != "" {
		query = query.Where("id_card LIKE ?", "%"+idCard+"%")
	}
	if phone, ok := filters["phone"]; ok && phone != "" {
		query = query.Where("phone LIKE ?", "%"+phone+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&patients).Error; err != nil {
		return nil, 0, err
	}

	return patients, total, nil
}

func (c *clinicalDB) UpdatePatient(ctx context.Context, patient *model.Patient) error {
	return c.db.WithContext(ctx).Save(patient).Error
}

// 就诊记录管理
func (c *clinicalDB) CreateVisit(ctx context.Context, visit *model.Visit) (int64, error) {
	if err := c.db.WithContext(ctx).Create(visit).Error; err != nil {
		return 0, err
	}
	return visit.ID, nil
}

func (c *clinicalDB) GetVisit(ctx context.Context, visitID int64) (*model.Visit, error) {
	var visit model.Visit
	if err := c.db.WithContext(ctx).Preload("Patient").First(&visit, visitID).Error; err != nil {
		return nil, err
	}
	return &visit, nil
}

func (c *clinicalDB) GetPatientVisits(ctx context.Context, patientID int64, limit, offset int) ([]*model.Visit, error) {
	var visits []*model.Visit
	if err := c.db.WithContext(ctx).Where("patient_id = ?", patientID).
		Limit(limit).Offset(offset).Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func (c *clinicalDB) ListAllVisits(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Visit, int64, error) {
	var visits []*model.Visit
	var total int64

	query := c.db.WithContext(ctx).Model(&model.Visit{}).Preload("Patient")

	// 添加搜索条件
	if patientID, ok := filters["patientId"]; ok && patientID != "" {
		query = query.Where("patient_id = ?", patientID)
	}
	if visitNo, ok := filters["visitNo"]; ok && visitNo != "" {
		query = query.Where("visit_no LIKE ?", "%"+visitNo+"%")
	}
	if department, ok := filters["department"]; ok && department != "" {
		query = query.Where("department LIKE ?", "%"+department+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("visit_time DESC").Limit(limit).Offset(offset).Find(&visits).Error; err != nil {
		return nil, 0, err
	}

	return visits, total, nil
}

func (c *clinicalDB) UpdateVisit(ctx context.Context, visit *model.Visit) error {
	return c.db.WithContext(ctx).Save(visit).Error
}

// 诊断管理
func (c *clinicalDB) AddDiagnosis(ctx context.Context, diagnosis *model.Diagnosis) error {
	return c.db.WithContext(ctx).Create(diagnosis).Error
}

func (c *clinicalDB) GetVisitDiagnoses(ctx context.Context, visitID int64) ([]*model.Diagnosis, error) {
	var diagnoses []*model.Diagnosis
	if err := c.db.WithContext(ctx).Where("visit_id = ?", visitID).Find(&diagnoses).Error; err != nil {
		return nil, err
	}
	return diagnoses, nil
}

// 检查管理
func (c *clinicalDB) AddExamination(ctx context.Context, examination *model.Examination) error {
	return c.db.WithContext(ctx).Create(examination).Error
}

func (c *clinicalDB) GetVisitExaminations(ctx context.Context, visitID int64) ([]*model.Examination, error) {
	var examinations []*model.Examination
	if err := c.db.WithContext(ctx).Where("visit_id = ?", visitID).Find(&examinations).Error; err != nil {
		return nil, err
	}
	return examinations, nil
}

func (c *clinicalDB) ListAllDiagnoses(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Diagnosis, int64, error) {
	var diagnoses []*model.Diagnosis
	var total int64

	query := c.db.WithContext(ctx).Model(&model.Diagnosis{}).Preload("Visit").Preload("Visit.Patient")

	// 添加搜索条件
	if patientID, ok := filters["patientId"]; ok && patientID != "" {
		query = query.Joins("JOIN visits ON diagnoses.visit_id = visits.id").
			Where("visits.patient_id = ?", patientID)
	}
	if visitNo, ok := filters["visitNo"]; ok && visitNo != "" {
		query = query.Joins("JOIN visits ON diagnoses.visit_id = visits.id").
			Where("visits.visit_no LIKE ?", "%"+visitNo+"%")
	}
	if diagnosisCode, ok := filters["diagnosisCode"]; ok && diagnosisCode != "" {
		query = query.Where("diagnosis_code LIKE ?", "%"+diagnosisCode+"%")
	}
	if diagnosisName, ok := filters["diagnosisName"]; ok && diagnosisName != "" {
		query = query.Where("diagnosis_name LIKE ?", "%"+diagnosisName+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&diagnoses).Error; err != nil {
		return nil, 0, err
	}

	return diagnoses, total, nil
}

func (c *clinicalDB) ListAllExaminations(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Examination, int64, error) {
	var examinations []*model.Examination
	var total int64

	query := c.db.WithContext(ctx).Model(&model.Examination{}).Preload("Visit").Preload("Visit.Patient")

	// 添加搜索条件
	if patientID, ok := filters["patientId"]; ok && patientID != "" {
		query = query.Joins("JOIN visits ON examinations.visit_id = visits.id").
			Where("visits.patient_id = ?", patientID)
	}
	if visitNo, ok := filters["visitNo"]; ok && visitNo != "" {
		query = query.Joins("JOIN visits ON examinations.visit_id = visits.id").
			Where("visits.visit_no LIKE ?", "%"+visitNo+"%")
	}
	if examinationType, ok := filters["examinationType"]; ok && examinationType != "" {
		query = query.Where("examination_type LIKE ?", "%"+examinationType+"%")
	}
	if examinationName, ok := filters["examinationName"]; ok && examinationName != "" {
		query = query.Where("examination_name LIKE ?", "%"+examinationName+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("exam_time DESC").Limit(limit).Offset(offset).Find(&examinations).Error; err != nil {
		return nil, 0, err
	}

	return examinations, total, nil
}
