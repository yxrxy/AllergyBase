package rpc

import (
	"context"
	"log"

	"github.com/yxrxy/AllergyBase/kitex_gen/clinical"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
	"github.com/yxrxy/AllergyBase/pkg/errno"
)

// InitClinicalRPC 初始化临床服务客户端
func InitClinicalRPC() {
	c, err := client.InitClinicalRPC()
	if err != nil {
		log.Fatalf("初始化临床服务客户端失败: %v", err)
	}
	clinicalClient = *c
}

// CreatePatientRPC 创建患者
func CreatePatientRPC(ctx context.Context, req *clinical.CreatePatientRequest) (int64, error) {
	resp, err := clinicalClient.CreatePatient(ctx, req)
	if err != nil {
		log.Printf("创建患者RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.PatientId, nil
}

// GetPatientRPC 获取患者信息
func GetPatientRPC(ctx context.Context, req *clinical.GetPatientRequest) (*clinical.GetPatientResponse, error) {
	resp, err := clinicalClient.GetPatient(ctx, req)
	if err != nil {
		log.Printf("获取患者信息RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// ListPatientsRPC 获取患者列表
func ListPatientsRPC(ctx context.Context, req *clinical.ListPatientsRequest) (*clinical.ListPatientsResponse, error) {
	resp, err := clinicalClient.ListPatients(ctx, req)
	if err != nil {
		log.Printf("获取患者列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// CreateVisitRPC 创建就诊记录
func CreateVisitRPC(ctx context.Context, req *clinical.CreateVisitRequest) (int64, error) {
	resp, err := clinicalClient.CreateVisit(ctx, req)
	if err != nil {
		log.Printf("创建就诊记录RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.VisitId, nil
}

// GetVisitRPC 获取就诊记录
func GetVisitRPC(ctx context.Context, req *clinical.GetVisitRequest) (*clinical.GetVisitResponse, error) {
	resp, err := clinicalClient.GetVisit(ctx, req)
	if err != nil {
		log.Printf("获取就诊记录RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// GetPatientVisitsRPC 获取患者就诊记录列表
func GetPatientVisitsRPC(ctx context.Context, req *clinical.GetPatientVisitsRequest) (*clinical.GetPatientVisitsResponse, error) {
	resp, err := clinicalClient.GetPatientVisits(ctx, req)
	if err != nil {
		log.Printf("获取患者就诊记录列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// ListAllVisitsRPC 获取所有就诊记录列表
func ListAllVisitsRPC(ctx context.Context, limit, offset int, filters map[string]string) (interface{}, int64, error) {
	// 创建请求参数
	req := &clinical.ListAllVisitsRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	// 设置过滤条件
	if visitNo, ok := filters["visitNo"]; ok && visitNo != "" {
		req.VisitNo = &visitNo
	}
	if department, ok := filters["department"]; ok && department != "" {
		req.Department = &department
	}
	if patientId, ok := filters["patientId"]; ok && patientId != "" {
		req.PatientId = &patientId
	}

	resp, err := clinicalClient.ListAllVisits(ctx, req)
	if err != nil {
		log.Printf("获取所有就诊记录列表RPC调用失败: %v", err)
		return nil, 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.Visits, int64(resp.Total), nil
}

// AddDiagnosisRPC 添加诊断
func AddDiagnosisRPC(ctx context.Context, req *clinical.AddDiagnosisRequest) (int64, error) {
	resp, err := clinicalClient.AddDiagnosis(ctx, req)
	if err != nil {
		log.Printf("添加诊断RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.DiagnosisId, nil
}

// GetVisitDiagnosesRPC 获取就诊诊断列表
func GetVisitDiagnosesRPC(ctx context.Context, req *clinical.GetVisitDiagnosesRequest) (*clinical.GetVisitDiagnosesResponse, error) {
	resp, err := clinicalClient.GetVisitDiagnoses(ctx, req)
	if err != nil {
		log.Printf("获取就诊诊断列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// AddExaminationRPC 添加检查记录
func AddExaminationRPC(ctx context.Context, req *clinical.AddExaminationRequest) (int64, error) {
	resp, err := clinicalClient.AddExamination(ctx, req)
	if err != nil {
		log.Printf("添加检查记录RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.ExaminationId, nil
}

// GetVisitExaminationsRPC 获取就诊检查记录列表
func GetVisitExaminationsRPC(ctx context.Context, req *clinical.GetVisitExaminationsRequest) (*clinical.GetVisitExaminationsResponse, error) {
	resp, err := clinicalClient.GetVisitExaminations(ctx, req)
	if err != nil {
		log.Printf("获取就诊检查记录列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// ListAllDiagnosesRPC 获取所有诊断记录列表
func ListAllDiagnosesRPC(ctx context.Context, limit, offset int, filters map[string]string) (interface{}, int64, error) {
	// 创建请求参数
	req := &clinical.ListAllDiagnosesRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	// 设置过滤条件
	if visitNo, ok := filters["visitNo"]; ok && visitNo != "" {
		req.VisitNo = &visitNo
	}
	if diagnosisCode, ok := filters["diagnosisCode"]; ok && diagnosisCode != "" {
		req.DiagnosisCode = &diagnosisCode
	}
	if diagnosisName, ok := filters["diagnosisName"]; ok && diagnosisName != "" {
		req.DiagnosisName = &diagnosisName
	}
	if patientId, ok := filters["patientId"]; ok && patientId != "" {
		req.PatientId = &patientId
	}

	resp, err := clinicalClient.ListAllDiagnoses(ctx, req)
	if err != nil {
		log.Printf("获取所有诊断记录列表RPC调用失败: %v", err)
		return nil, 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.Diagnoses, int64(resp.Total), nil
}

// ListAllExaminationsRPC 获取所有检查记录列表
func ListAllExaminationsRPC(ctx context.Context, limit, offset int, filters map[string]string) (interface{}, int64, error) {
	// 创建请求参数
	req := &clinical.ListAllExaminationsRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	// 设置过滤条件
	if visitNo, ok := filters["visitNo"]; ok && visitNo != "" {
		req.VisitNo = &visitNo
	}
	if examinationType, ok := filters["examinationType"]; ok && examinationType != "" {
		req.ExaminationType = &examinationType
	}
	if examinationName, ok := filters["examinationName"]; ok && examinationName != "" {
		req.ExaminationName = &examinationName
	}
	if patientId, ok := filters["patientId"]; ok && patientId != "" {
		req.PatientId = &patientId
	}

	resp, err := clinicalClient.ListAllExaminations(ctx, req)
	if err != nil {
		log.Printf("获取所有检查记录列表RPC调用失败: %v", err)
		return nil, 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.Examinations, int64(resp.Total), nil
}
