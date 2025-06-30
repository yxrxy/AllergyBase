package repository

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/clinical/domain/model"
)

type ClinicalDB interface {
	// 患者管理
	CreatePatient(ctx context.Context, patient *model.Patient) (int64, error)
	GetPatient(ctx context.Context, patientID int64) (*model.Patient, error)
	GetPatientByNo(ctx context.Context, patientNo string) (*model.Patient, error)
	ListPatients(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Patient, int64, error)
	UpdatePatient(ctx context.Context, patient *model.Patient) error

	// 就诊记录管理
	CreateVisit(ctx context.Context, visit *model.Visit) (int64, error)
	GetVisit(ctx context.Context, visitID int64) (*model.Visit, error)
	GetPatientVisits(ctx context.Context, patientID int64, limit, offset int) ([]*model.Visit, error)
	ListAllVisits(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Visit, int64, error)
	UpdateVisit(ctx context.Context, visit *model.Visit) error

	// 诊断管理
	AddDiagnosis(ctx context.Context, diagnosis *model.Diagnosis) error
	GetVisitDiagnoses(ctx context.Context, visitID int64) ([]*model.Diagnosis, error)
	ListAllDiagnoses(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Diagnosis, int64, error)

	// 检查管理
	AddExamination(ctx context.Context, examination *model.Examination) error
	GetVisitExaminations(ctx context.Context, visitID int64) ([]*model.Examination, error)
	ListAllExaminations(ctx context.Context, limit, offset int, filters map[string]string) ([]*model.Examination, int64, error)
}

type ClinicalCache interface {
	SetPatientInfo(ctx context.Context, patientID int64, patient *model.Patient) error
	GetPatientInfo(ctx context.Context, patientID int64) (*model.Patient, error)
	DeletePatientInfo(ctx context.Context, patientID int64) error

	SetVisitInfo(ctx context.Context, visitID int64, visit *model.Visit) error
	GetVisitInfo(ctx context.Context, visitID int64) (*model.Visit, error)
	DeleteVisitInfo(ctx context.Context, visitID int64) error
}
