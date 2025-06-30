package service

import (
	"context"
	"fmt"
	"time"

	"github.com/yxrxy/AllergyBase/app/clinical/domain/model"
	"github.com/yxrxy/AllergyBase/app/clinical/domain/repository"
)

type ClinicalService struct {
	db    repository.ClinicalDB
	cache repository.ClinicalCache
}

func NewClinicalService(db repository.ClinicalDB, cache repository.ClinicalCache) *ClinicalService {
	return &ClinicalService{
		db:    db,
		cache: cache,
	}
}

// 患者管理相关方法
func (s *ClinicalService) CreatePatient(ctx context.Context, patient *model.Patient) (int64, error) {
	// 生成患者编号
	if patient.PatientNo == "" {
		patient.PatientNo = s.generatePatientNo()
	}

	patientID, err := s.db.CreatePatient(ctx, patient)
	if err != nil {
		return 0, err
	}

	patient.ID = patientID
	// 缓存患者信息
	_ = s.cache.SetPatientInfo(ctx, patientID, patient)

	return patientID, nil
}

func (s *ClinicalService) GetPatient(ctx context.Context, patientID int64) (*model.Patient, error) {
	// 先从缓存获取
	patient, err := s.cache.GetPatientInfo(ctx, patientID)
	if err == nil && patient != nil {
		return patient, nil
	}

	// 从数据库获取
	patient, err = s.db.GetPatient(ctx, patientID)
	if err != nil {
		return nil, err
	}

	// 缓存患者信息
	_ = s.cache.SetPatientInfo(ctx, patientID, patient)

	return patient, nil
}

func (s *ClinicalService) ListPatients(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Patient, int64, error) {
	return s.db.ListPatients(ctx, limit, offset, filters)
}

// 就诊记录相关方法
func (s *ClinicalService) CreateVisit(ctx context.Context, visit *model.Visit) (int64, error) {
	// 生成就诊编号
	if visit.VisitNo == "" {
		visit.VisitNo = s.generateVisitNo()
	}

	visitID, err := s.db.CreateVisit(ctx, visit)
	if err != nil {
		return 0, err
	}

	visit.ID = visitID
	// 缓存就诊信息
	_ = s.cache.SetVisitInfo(ctx, visitID, visit)

	return visitID, nil
}

func (s *ClinicalService) GetVisit(ctx context.Context, visitID int64) (*model.Visit, error) {
	// 先从缓存获取
	visit, err := s.cache.GetVisitInfo(ctx, visitID)
	if err == nil && visit != nil {
		return visit, nil
	}

	// 从数据库获取
	visit, err = s.db.GetVisit(ctx, visitID)
	if err != nil {
		return nil, err
	}

	// 缓存就诊信息
	_ = s.cache.SetVisitInfo(ctx, visitID, visit)

	return visit, nil
}

func (s *ClinicalService) GetPatientVisits(ctx context.Context, patientID int64, limit, offset int) ([]*model.Visit, error) {
	return s.db.GetPatientVisits(ctx, patientID, limit, offset)
}

func (s *ClinicalService) ListAllVisits(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Visit, int64, error) {
	return s.db.ListAllVisits(ctx, limit, offset, filters)
}

// 诊断相关方法
func (s *ClinicalService) AddDiagnosis(ctx context.Context, diagnosis *model.Diagnosis) error {
	return s.db.AddDiagnosis(ctx, diagnosis)
}

func (s *ClinicalService) GetVisitDiagnoses(ctx context.Context, visitID int64) ([]*model.Diagnosis, error) {
	return s.db.GetVisitDiagnoses(ctx, visitID)
}

// 检查相关方法
func (s *ClinicalService) AddExamination(ctx context.Context, examination *model.Examination) error {
	return s.db.AddExamination(ctx, examination)
}

func (s *ClinicalService) GetVisitExaminations(ctx context.Context, visitID int64) ([]*model.Examination, error) {
	return s.db.GetVisitExaminations(ctx, visitID)
}

func (s *ClinicalService) ListAllDiagnoses(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Diagnosis, int64, error) {
	return s.db.ListAllDiagnoses(ctx, limit, offset, filters)
}

func (s *ClinicalService) ListAllExaminations(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Examination, int64, error) {
	return s.db.ListAllExaminations(ctx, limit, offset, filters)
}

// 生成患者编号
func (s *ClinicalService) generatePatientNo() string {
	return fmt.Sprintf("P%s%06d", time.Now().Format("20060102"), time.Now().Unix()%1000000)
}

// 生成就诊编号
func (s *ClinicalService) generateVisitNo() string {
	return fmt.Sprintf("V%s%06d", time.Now().Format("20060102"), time.Now().Unix()%1000000)
}
