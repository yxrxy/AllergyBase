package rpc

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/clinical/controllers/rpc/pack"
	"github.com/yxrxy/AllergyBase/app/clinical/domain/model"
	"github.com/yxrxy/AllergyBase/app/clinical/usecase"
	"github.com/yxrxy/AllergyBase/kitex_gen/clinical"
	"github.com/yxrxy/AllergyBase/pkg/base"
)

type ClinicalHandler struct {
	useCase *usecase.ClinicalCase
}

func NewClinicalHandler(useCase *usecase.ClinicalCase) *ClinicalHandler {
	return &ClinicalHandler{useCase: useCase}
}

// CreatePatient 创建患者
func (h *ClinicalHandler) CreatePatient(ctx context.Context, req *clinical.CreatePatientRequest) (r *clinical.CreatePatientResponse, err error) {
	r = new(clinical.CreatePatientResponse)

	patient := pack.BuildPatientFromRequest(req.Patient)

	var patientID int64
	if patientID, err = h.useCase.CreatePatient(ctx, patient); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.PatientId = patientID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetPatient 获取患者信息
func (h *ClinicalHandler) GetPatient(ctx context.Context, req *clinical.GetPatientRequest) (r *clinical.GetPatientResponse, err error) {
	r = new(clinical.GetPatientResponse)

	var patient *model.Patient
	if patient, err = h.useCase.GetPatient(ctx, req.PatientId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Patient = pack.Patient(patient)
	r.Base = base.BuildBaseResp(err)
	return
}

// ListPatients 获取患者列表
func (h *ClinicalHandler) ListPatients(ctx context.Context, req *clinical.ListPatientsRequest) (r *clinical.ListPatientsResponse, err error) {
	r = new(clinical.ListPatientsResponse)

	// 构建搜索过滤器
	filters := make(map[string]string)
	if req.PatientNo != nil {
		filters["patientNo"] = *req.PatientNo
	}
	if req.Name != nil {
		filters["name"] = *req.Name
	}
	if req.IdCard != nil {
		filters["idCard"] = *req.IdCard
	}
	if req.Phone != nil {
		filters["phone"] = *req.Phone
	}

	var patients []*model.Patient
	var total int64
	if patients, total, err = h.useCase.ListPatients(ctx, int(req.Limit), int(req.Offset), filters); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Patients = pack.PatientList(patients)
	r.Total = int32(total)
	r.Base = base.BuildBaseResp(err)
	return
}

// CreateVisit 创建就诊记录
func (h *ClinicalHandler) CreateVisit(ctx context.Context, req *clinical.CreateVisitRequest) (r *clinical.CreateVisitResponse, err error) {
	r = new(clinical.CreateVisitResponse)

	visit := pack.BuildVisitFromRequest(req.Visit)

	var visitID int64
	if visitID, err = h.useCase.CreateVisit(ctx, visit); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.VisitId = visitID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetVisit 获取就诊记录
func (h *ClinicalHandler) GetVisit(ctx context.Context, req *clinical.GetVisitRequest) (r *clinical.GetVisitResponse, err error) {
	r = new(clinical.GetVisitResponse)

	var visit *model.Visit
	if visit, err = h.useCase.GetVisit(ctx, req.VisitId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Visit = pack.Visit(visit)
	r.Base = base.BuildBaseResp(err)
	return
}

// GetPatientVisits 获取患者就诊记录列表
func (h *ClinicalHandler) GetPatientVisits(ctx context.Context, req *clinical.GetPatientVisitsRequest) (r *clinical.GetPatientVisitsResponse, err error) {
	r = new(clinical.GetPatientVisitsResponse)

	var visits []*model.Visit
	if visits, err = h.useCase.GetPatientVisits(ctx, req.PatientId, 0, 10); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Visits = pack.VisitList(visits)
	r.Base = base.BuildBaseResp(err)
	return
}

// AddDiagnosis 添加诊断
func (h *ClinicalHandler) AddDiagnosis(ctx context.Context, req *clinical.AddDiagnosisRequest) (r *clinical.AddDiagnosisResponse, err error) {
	r = new(clinical.AddDiagnosisResponse)

	diagnosis := pack.BuildDiagnosisFromRequest(req.Diagnosis)

	if err = h.useCase.AddDiagnosis(ctx, diagnosis); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Base = base.BuildBaseResp(err)
	return
}

// GetVisitDiagnoses 获取就诊诊断列表
func (h *ClinicalHandler) GetVisitDiagnoses(ctx context.Context, req *clinical.GetVisitDiagnosesRequest) (r *clinical.GetVisitDiagnosesResponse, err error) {
	r = new(clinical.GetVisitDiagnosesResponse)

	var diagnoses []*model.Diagnosis
	if diagnoses, err = h.useCase.GetVisitDiagnoses(ctx, req.VisitId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Diagnoses = pack.DiagnosisList(diagnoses)
	r.Base = base.BuildBaseResp(err)
	return
}

// AddExamination 添加检查
func (h *ClinicalHandler) AddExamination(ctx context.Context, req *clinical.AddExaminationRequest) (r *clinical.AddExaminationResponse, err error) {
	r = new(clinical.AddExaminationResponse)

	examination := pack.BuildExaminationFromRequest(req.Examination)

	if err = h.useCase.AddExamination(ctx, examination); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Base = base.BuildBaseResp(err)
	return
}

// GetVisitExaminations 获取就诊检查列表
func (h *ClinicalHandler) GetVisitExaminations(ctx context.Context, req *clinical.GetVisitExaminationsRequest) (r *clinical.GetVisitExaminationsResponse, err error) {
	r = new(clinical.GetVisitExaminationsResponse)

	var examinations []*model.Examination
	if examinations, err = h.useCase.GetVisitExaminations(ctx, req.VisitId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Examinations = pack.ExaminationList(examinations)
	r.Base = base.BuildBaseResp(err)
	return
}

// ListAllVisits 获取所有就诊记录列表
func (h *ClinicalHandler) ListAllVisits(ctx context.Context, req *clinical.ListAllVisitsRequest) (r *clinical.ListAllVisitsResponse, err error) {
	r = new(clinical.ListAllVisitsResponse)

	// 构建搜索过滤器
	filters := make(map[string]string)
	if req.PatientId != nil {
		filters["patientId"] = *req.PatientId
	}
	if req.VisitNo != nil {
		filters["visitNo"] = *req.VisitNo
	}
	if req.Department != nil {
		filters["department"] = *req.Department
	}

	var visits []*model.Visit
	var total int64
	if visits, total, err = h.useCase.ListAllVisits(ctx, int(req.Limit), int(req.Offset), filters); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Visits = pack.VisitList(visits)
	r.Total = int32(total)
	r.Base = base.BuildBaseResp(err)
	return
}

// ListAllDiagnoses 获取所有诊断记录列表
func (h *ClinicalHandler) ListAllDiagnoses(ctx context.Context, req *clinical.ListAllDiagnosesRequest) (r *clinical.ListAllDiagnosesResponse, err error) {
	r = new(clinical.ListAllDiagnosesResponse)

	// 构建搜索过滤器
	filters := make(map[string]string)
	if req.PatientId != nil {
		filters["patientId"] = *req.PatientId
	}
	if req.VisitNo != nil {
		filters["visitNo"] = *req.VisitNo
	}
	if req.DiagnosisCode != nil {
		filters["diagnosisCode"] = *req.DiagnosisCode
	}
	if req.DiagnosisName != nil {
		filters["diagnosisName"] = *req.DiagnosisName
	}

	var diagnoses []*model.Diagnosis
	var total int64
	if diagnoses, total, err = h.useCase.ListAllDiagnoses(ctx, int(req.Limit), int(req.Offset), filters); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Diagnoses = pack.DiagnosisList(diagnoses)
	r.Total = int32(total)
	r.Base = base.BuildBaseResp(err)
	return
}

// ListAllExaminations 获取所有检查记录列表
func (h *ClinicalHandler) ListAllExaminations(ctx context.Context, req *clinical.ListAllExaminationsRequest) (r *clinical.ListAllExaminationsResponse, err error) {
	r = new(clinical.ListAllExaminationsResponse)

	// 构建搜索过滤器
	filters := make(map[string]string)
	if req.PatientId != nil {
		filters["patientId"] = *req.PatientId
	}
	if req.VisitNo != nil {
		filters["visitNo"] = *req.VisitNo
	}
	if req.ExaminationType != nil {
		filters["examinationType"] = *req.ExaminationType
	}
	if req.ExaminationName != nil {
		filters["examinationName"] = *req.ExaminationName
	}

	var examinations []*model.Examination
	var total int64
	if examinations, total, err = h.useCase.ListAllExaminations(ctx, int(req.Limit), int(req.Offset), filters); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Examinations = pack.ExaminationList(examinations)
	r.Total = int32(total)
	r.Base = base.BuildBaseResp(err)
	return
}
