namespace go clinical

include "model.thrift"

// 创建患者请求
struct CreatePatientRequest {
    1: required model.Patient patient
}

// 创建患者响应
struct CreatePatientResponse {
    1: required model.BaseResp Base
    2: required i64 patientId
}

// 获取患者请求
struct GetPatientRequest {
    1: required i64 patientId
}

// 获取患者响应
struct GetPatientResponse {
    1: required model.BaseResp Base
    2: required model.Patient patient
}

// 获取患者列表请求
struct ListPatientsRequest {
    1: required i32 offset
    2: required i32 limit
    3: optional string patientNo
    4: optional string name
    5: optional string idCard
    6: optional string phone
}

// 获取患者列表响应
struct ListPatientsResponse {
    1: required model.BaseResp Base
    2: required list<model.Patient> patients
    3: required i32 total
}

// 创建就诊记录请求
struct CreateVisitRequest {
    1: required model.Visit visit
}

// 创建就诊记录响应
struct CreateVisitResponse {
    1: required model.BaseResp Base
    2: required i64 visitId
}

// 获取就诊记录请求
struct GetVisitRequest {
    1: required i64 visitId
}

// 获取就诊记录响应
struct GetVisitResponse {
    1: required model.BaseResp Base
    2: required model.Visit visit
}

// 获取患者就诊记录请求
struct GetPatientVisitsRequest {
    1: required i64 patientId
}

// 获取患者就诊记录响应
struct GetPatientVisitsResponse {
    1: required model.BaseResp Base
    2: required list<model.Visit> visits
}

// 获取所有就诊记录请求
struct ListAllVisitsRequest {
    1: required i32 offset
    2: required i32 limit
    3: optional string patientId
    4: optional string visitNo
    5: optional string department
}

// 获取所有就诊记录响应
struct ListAllVisitsResponse {
    1: required model.BaseResp Base
    2: required list<model.Visit> visits
    3: required i32 total
}

// 添加诊断请求
struct AddDiagnosisRequest {
    1: required model.Diagnosis diagnosis
}

// 添加诊断响应
struct AddDiagnosisResponse {
    1: required model.BaseResp Base
    2: required i64 diagnosisId
}

// 获取就诊诊断请求
struct GetVisitDiagnosesRequest {
    1: required i64 visitId
}

// 获取就诊诊断响应
struct GetVisitDiagnosesResponse {
    1: required model.BaseResp Base
    2: required list<model.Diagnosis> diagnoses
}

// 添加检查请求
struct AddExaminationRequest {
    1: required model.Examination examination
}

// 添加检查响应
struct AddExaminationResponse {
    1: required model.BaseResp Base
    2: required i64 examinationId
}

// 获取就诊检查请求
struct GetVisitExaminationsRequest {
    1: required i64 visitId
}

// 获取就诊检查响应
struct GetVisitExaminationsResponse {
    1: required model.BaseResp Base
    2: required list<model.Examination> examinations
}

// 获取所有诊断请求
struct ListAllDiagnosesRequest {
    1: required i32 offset
    2: required i32 limit
    3: optional string patientId
    4: optional string visitNo
    5: optional string diagnosisCode
    6: optional string diagnosisName
}

// 获取所有诊断响应
struct ListAllDiagnosesResponse {
    1: required model.BaseResp Base
    2: required list<model.Diagnosis> diagnoses
    3: required i32 total
}

// 获取所有检查请求
struct ListAllExaminationsRequest {
    1: required i32 offset
    2: required i32 limit
    3: optional string patientId
    4: optional string visitNo
    5: optional string examinationType
    6: optional string examinationName
}

// 获取所有检查响应
struct ListAllExaminationsResponse {
    1: required model.BaseResp Base
    2: required list<model.Examination> examinations
    3: required i32 total
}

service ClinicalService {
    // 患者信息管理
    CreatePatientResponse CreatePatient(1: CreatePatientRequest req)
    GetPatientResponse GetPatient(1: GetPatientRequest req)
    ListPatientsResponse ListPatients(1: ListPatientsRequest req)
    
    // 就诊记录管理
    CreateVisitResponse CreateVisit(1: CreateVisitRequest req)
    GetVisitResponse GetVisit(1: GetVisitRequest req)
    GetPatientVisitsResponse GetPatientVisits(1: GetPatientVisitsRequest req)
    ListAllVisitsResponse ListAllVisits(1: ListAllVisitsRequest req)
    
    // 诊断管理
    AddDiagnosisResponse AddDiagnosis(1: AddDiagnosisRequest req)
    GetVisitDiagnosesResponse GetVisitDiagnoses(1: GetVisitDiagnosesRequest req)
    ListAllDiagnosesResponse ListAllDiagnoses(1: ListAllDiagnosesRequest req)
    
    // 检查管理
    AddExaminationResponse AddExamination(1: AddExaminationRequest req)
    GetVisitExaminationsResponse GetVisitExaminations(1: GetVisitExaminationsRequest req)
    ListAllExaminationsResponse ListAllExaminations(1: ListAllExaminationsRequest req)
}