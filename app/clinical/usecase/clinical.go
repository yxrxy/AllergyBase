package usecase

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/clinical/domain/model"
	"github.com/yxrxy/AllergyBase/app/clinical/domain/repository"
	"github.com/yxrxy/AllergyBase/app/clinical/domain/service"
)

type ClinicalCase struct {
	db  repository.ClinicalDB
	svc *service.ClinicalService
}

func NewClinicalCase(db repository.ClinicalDB, svc *service.ClinicalService) *ClinicalCase {
	return &ClinicalCase{
		db:  db,
		svc: svc,
	}
}

// 患者管理
func (uc *ClinicalCase) CreatePatient(ctx context.Context, patient *model.Patient) (int64, error) {
	return uc.svc.CreatePatient(ctx, patient)
}

func (uc *ClinicalCase) GetPatient(ctx context.Context, patientID int64) (*model.Patient, error) {
	return uc.svc.GetPatient(ctx, patientID)
}

func (uc *ClinicalCase) ListPatients(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Patient, int64, error) {
	return uc.svc.ListPatients(ctx, limit, offset, filters)
}

// 就诊记录管理
func (uc *ClinicalCase) CreateVisit(ctx context.Context, visit *model.Visit) (int64, error) {
	return uc.svc.CreateVisit(ctx, visit)
}

func (uc *ClinicalCase) GetVisit(ctx context.Context, visitID int64) (*model.Visit, error) {
	return uc.svc.GetVisit(ctx, visitID)
}

func (uc *ClinicalCase) GetPatientVisits(ctx context.Context, patientID int64, limit, offset int) ([]*model.Visit, error) {
	return uc.svc.GetPatientVisits(ctx, patientID, limit, offset)
}

func (uc *ClinicalCase) ListAllVisits(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Visit, int64, error) {
	return uc.svc.ListAllVisits(ctx, limit, offset, filters)
}

// 诊断管理
func (uc *ClinicalCase) AddDiagnosis(ctx context.Context, diagnosis *model.Diagnosis) error {
	return uc.svc.AddDiagnosis(ctx, diagnosis)
}

func (uc *ClinicalCase) GetVisitDiagnoses(ctx context.Context, visitID int64) ([]*model.Diagnosis, error) {
	return uc.svc.GetVisitDiagnoses(ctx, visitID)
}

// 检查管理
func (uc *ClinicalCase) AddExamination(ctx context.Context, examination *model.Examination) error {
	return uc.svc.AddExamination(ctx, examination)
}

func (uc *ClinicalCase) GetVisitExaminations(ctx context.Context, visitID int64) ([]*model.Examination, error) {
	return uc.svc.GetVisitExaminations(ctx, visitID)
}

func (uc *ClinicalCase) ListAllDiagnoses(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Diagnosis, int64, error) {
	return uc.svc.ListAllDiagnoses(ctx, limit, offset, filters)
}

func (uc *ClinicalCase) ListAllExaminations(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Examination, int64, error) {
	return uc.svc.ListAllExaminations(ctx, limit, offset, filters)
}
