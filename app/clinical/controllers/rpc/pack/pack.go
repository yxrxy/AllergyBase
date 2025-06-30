package pack

import (
	"time"

	"github.com/yxrxy/AllergyBase/app/clinical/domain/model"
	rpcmodel "github.com/yxrxy/AllergyBase/kitex_gen/model"
)

// Patient 将数据库患者实体转换为RPC响应实体
func Patient(p *model.Patient) *rpcmodel.Patient {
	if p == nil {
		return nil
	}

	var gender rpcmodel.Gender
	if p.Gender == "1" {
		gender = rpcmodel.Gender_MALE
	} else {
		gender = rpcmodel.Gender_FEMALE
	}

	var birthDate string
	if p.BirthDate != nil {
		birthDate = p.BirthDate.Format("2006-01-02")
	}

	medicalInsuranceType := int16(p.MedicalInsuranceType)

	return &rpcmodel.Patient{
		Id:                   p.ID,
		PatientNo:            p.PatientNo,
		Name:                 p.Name,
		Gender:               gender,
		BirthDate:            birthDate,
		IdCard:               &p.IdCard,
		Phone:                &p.Phone,
		Address:              &p.Address,
		EmergencyContact:     &p.EmergencyContact,
		EmergencyPhone:       &p.EmergencyPhone,
		Height:               &p.Height,
		Weight:               &p.Weight,
		BirthWeight:          &p.BirthWeight,
		MedicalInsuranceType: &medicalInsuranceType,
		MedicalInsuranceNo:   &p.MedicalInsuranceNo,
	}
}

// Visit 将数据库就诊记录实体转换为RPC响应实体
func Visit(v *model.Visit) *rpcmodel.Visit {
	if v == nil {
		return nil
	}

	// 转换就诊类型
	var visitType rpcmodel.VisitType
	switch v.VisitType {
	case 1:
		visitType = rpcmodel.VisitType_OUTPATIENT
	case 2:
		visitType = rpcmodel.VisitType_EMERGENCY
	case 3:
		visitType = rpcmodel.VisitType_INPATIENT
	default:
		visitType = rpcmodel.VisitType_OUTPATIENT
	}

	visit := &rpcmodel.Visit{
		Id:             v.ID,
		PatientId:      v.PatientID,
		VisitNo:        v.VisitNo,
		VisitTime:      v.VisitTime.Format("2006-01-02 15:04:05"),
		VisitType:      visitType,
		Department:     &v.Department,
		ChiefComplaint: &v.ChiefComplaint,
	}

	// 设置医生ID（如果有）
	if v.DoctorID != nil {
		visit.DoctorId = v.DoctorID
	}

	return visit
}

// Diagnosis 将数据库诊断记录实体转换为RPC响应实体
func Diagnosis(d *model.Diagnosis) *rpcmodel.Diagnosis {
	if d == nil {
		return nil
	}
	return &rpcmodel.Diagnosis{
		Id:            d.ID,
		VisitId:       d.VisitID,
		DiagnosisType: &d.DiagnosisType,
		IcdCode:       &d.DiagnosisCode,
	}
}

// Examination 将数据库检查记录实体转换为RPC响应实体
func Examination(e *model.Examination) *rpcmodel.Examination {
	if e == nil {
		return nil
	}

	examTime := e.ExamTime.Format("2006-01-02 15:04:05")

	return &rpcmodel.Examination{
		Id:             e.ID,
		VisitId:        e.VisitID,
		ExamType:       &e.ExamType,
		ExamTime:       &examTime,
		ExamResult_:    &e.ExamResult,
		ResultUnit:     &e.Unit,
		ReferenceRange: &e.ReferenceValue,
	}
}

// PatientList 将患者列表转换为RPC响应实体
func PatientList(patients []*model.Patient) []*rpcmodel.Patient {
	if len(patients) == 0 {
		return nil
	}
	result := make([]*rpcmodel.Patient, 0, len(patients))
	for _, p := range patients {
		result = append(result, Patient(p))
	}
	return result
}

// VisitList 将就诊记录列表转换为RPC响应实体
func VisitList(visits []*model.Visit) []*rpcmodel.Visit {
	if len(visits) == 0 {
		return nil
	}
	result := make([]*rpcmodel.Visit, 0, len(visits))
	for _, v := range visits {
		result = append(result, Visit(v))
	}
	return result
}

// DiagnosisList 将诊断记录列表转换为RPC响应实体
func DiagnosisList(diagnoses []*model.Diagnosis) []*rpcmodel.Diagnosis {
	if len(diagnoses) == 0 {
		return nil
	}
	result := make([]*rpcmodel.Diagnosis, 0, len(diagnoses))
	for _, d := range diagnoses {
		result = append(result, Diagnosis(d))
	}
	return result
}

// ExaminationList 将检查记录列表转换为RPC响应实体
func ExaminationList(examinations []*model.Examination) []*rpcmodel.Examination {
	if len(examinations) == 0 {
		return nil
	}
	result := make([]*rpcmodel.Examination, 0, len(examinations))
	for _, e := range examinations {
		result = append(result, Examination(e))
	}
	return result
}

// BuildPatientFromRequest 将RPC请求转换为数据库实体
func BuildPatientFromRequest(req *rpcmodel.Patient) *model.Patient {
	if req == nil {
		return nil
	}

	var gender string
	if req.Gender == rpcmodel.Gender_MALE {
		gender = "1" // 男性对应数字1
	} else {
		gender = "0" // 女性对应数字0
	}

	var birthDate *time.Time
	if req.BirthDate != "" {
		if t, err := time.Parse("2006-01-02", req.BirthDate); err == nil {
			birthDate = &t
		}
	}

	var height, weight, birthWeight float64
	if req.Height != nil {
		height = *req.Height
	}
	if req.Weight != nil {
		weight = *req.Weight
	}
	if req.BirthWeight != nil {
		birthWeight = *req.BirthWeight
	}

	var medicalInsuranceType int
	if req.MedicalInsuranceType != nil {
		medicalInsuranceType = int(*req.MedicalInsuranceType)
	}

	var medicalInsuranceNo string
	if req.MedicalInsuranceNo != nil {
		medicalInsuranceNo = *req.MedicalInsuranceNo
	}

	return &model.Patient{
		ID:                   req.Id,
		PatientNo:            req.PatientNo,
		Name:                 req.Name,
		Gender:               gender,
		BirthDate:            birthDate,
		IdCard:               getStringValue(req.IdCard),
		Phone:                getStringValue(req.Phone),
		Address:              getStringValue(req.Address),
		EmergencyContact:     getStringValue(req.EmergencyContact),
		EmergencyPhone:       getStringValue(req.EmergencyPhone),
		Height:               height,
		Weight:               weight,
		BirthWeight:          birthWeight,
		MedicalInsuranceType: medicalInsuranceType,
		MedicalInsuranceNo:   medicalInsuranceNo,
	}
}

// BuildVisitFromRequest 将RPC请求转换为数据库实体
func BuildVisitFromRequest(req *rpcmodel.Visit) *model.Visit {
	if req == nil {
		return nil
	}

	var visitTime time.Time
	if req.VisitTime != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", req.VisitTime); err == nil {
			visitTime = t
		} else {
			visitTime = time.Now()
		}
	} else {
		visitTime = time.Now()
	}

	visit := &model.Visit{
		PatientID:      req.PatientId,
		VisitNo:        req.VisitNo,
		VisitTime:      visitTime,
		VisitType:      int(req.VisitType),
		Department:     getStringValue(req.Department),
		ChiefComplaint: getStringValue(req.ChiefComplaint),
	}

	// 设置医生ID（如果提供）
	if req.DoctorId != nil {
		visit.DoctorID = req.DoctorId
	}

	// 只有在更新记录时才设置ID
	if req.Id > 0 {
		visit.ID = req.Id
	}

	return visit
}

// BuildDiagnosisFromRequest 将RPC请求转换为数据库实体
func BuildDiagnosisFromRequest(req *rpcmodel.Diagnosis) *model.Diagnosis {
	if req == nil {
		return nil
	}

	diagnosis := &model.Diagnosis{
		VisitID:       req.VisitId,
		DiagnosisCode: getStringValue(req.IcdCode),
		DiagnosisType: getStringValue(req.DiagnosisType),
	}

	// 只有在更新记录时才设置ID
	if req.Id > 0 {
		diagnosis.ID = req.Id
	}

	return diagnosis
}

// BuildExaminationFromRequest 将RPC请求转换为数据库实体
func BuildExaminationFromRequest(req *rpcmodel.Examination) *model.Examination {
	if req == nil {
		return nil
	}

	var examTime time.Time
	if req.ExamTime != nil && *req.ExamTime != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", *req.ExamTime); err == nil {
			examTime = t
		} else {
			examTime = time.Now()
		}
	} else {
		examTime = time.Now()
	}

	examination := &model.Examination{
		VisitID:        req.VisitId,
		ExamType:       getStringValue(req.ExamType),
		ExamTime:       examTime,
		ExamResult:     getStringValue(req.ExamResult_),
		Unit:           getStringValue(req.ResultUnit),
		ReferenceValue: getStringValue(req.ReferenceRange),
	}

	// 只有在更新记录时才设置ID
	if req.Id > 0 {
		examination.ID = req.Id
	}

	return examination
}

// Helper function to safely extract string from pointer
func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
