namespace go api.clinical

include "../clinical.thrift"

// API 服务
service ClinicalAPI {
    // 患者信息管理接口
    clinical.CreatePatientResponse CreatePatient(1: clinical.CreatePatientRequest request) (api.post="/api/v1/patient")
    clinical.GetPatientResponse GetPatient(1: clinical.GetPatientRequest request) (api.get="/api/v1/patient/:patientId")
    clinical.ListPatientsResponse ListPatients(1: clinical.ListPatientsRequest request) (api.get="/api/v1/patients")
    
    // 就诊记录管理接口
    clinical.CreateVisitResponse CreateVisit(1: clinical.CreateVisitRequest request) (api.post="/api/v1/visit")
    clinical.GetVisitResponse GetVisit(1: clinical.GetVisitRequest request) (api.get="/api/v1/visit/:visitId")
    clinical.GetPatientVisitsResponse GetPatientVisits(1: clinical.GetPatientVisitsRequest request) (api.get="/api/v1/patient/:patientId/visits")
    clinical.ListAllVisitsResponse ListAllVisits(1: clinical.ListAllVisitsRequest request) (api.get="/api/v1/visits")
    
    // 诊断管理接口
    clinical.AddDiagnosisResponse AddDiagnosis(1: clinical.AddDiagnosisRequest request) (api.post="/api/v1/diagnosis")
    clinical.GetVisitDiagnosesResponse GetVisitDiagnoses(1: clinical.GetVisitDiagnosesRequest request) (api.get="/api/v1/visit/:visitId/diagnoses")
    clinical.ListAllDiagnosesResponse ListAllDiagnoses(1: clinical.ListAllDiagnosesRequest request) (api.get="/api/v1/diagnoses")
    
    // 检查管理接口
    clinical.AddExaminationResponse AddExamination(1: clinical.AddExaminationRequest request) (api.post="/api/v1/examination")
    clinical.GetVisitExaminationsResponse GetVisitExaminations(1: clinical.GetVisitExaminationsRequest request) (api.get="/api/v1/visit/:visitId/examinations")
    clinical.ListAllExaminationsResponse ListAllExaminations(1: clinical.ListAllExaminationsRequest request) (api.get="/api/v1/examinations")
} 