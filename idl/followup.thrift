namespace go followup

include "model.thrift"

// 创建随访计划请求
struct CreateFollowupPlanRequest {
    1: required model.FollowupPlan plan
}

// 创建随访计划响应
struct CreateFollowupPlanResponse {
    1: required model.BaseResp Base
    2: required i64 planId
}

// 获取随访计划请求
struct GetFollowupPlanRequest {
    1: required i64 planId
}

// 获取随访计划响应
struct GetFollowupPlanResponse {
    1: required model.BaseResp Base
    2: required model.FollowupPlan plan
}

// 获取所有随访计划请求
struct GetFollowupPlansRequest {
    1: optional i32 page
    2: optional i32 pageSize
    3: optional string status
    4: optional string planName
}

// 获取所有随访计划响应
struct GetFollowupPlansResponse {
    1: required model.BaseResp Base
    2: required list<model.FollowupPlan> plans
    3: required i32 total
}

// 获取患者随访计划请求
struct GetPatientFollowupPlansRequest {
    1: required i64 patientId
}

// 获取患者随访计划响应
struct GetPatientFollowupPlansResponse {
    1: required model.BaseResp Base
    2: required list<model.FollowupPlan> plans
}

// 创建随访记录请求
struct CreateFollowupRecordRequest {
    1: required model.FollowupRecord record
}

// 创建随访记录响应
struct CreateFollowupRecordResponse {
    1: required model.BaseResp Base
    2: required i64 recordId
}

// 获取随访记录请求
struct GetFollowupRecordRequest {
    1: required i64 recordId
}

// 获取随访记录响应
struct GetFollowupRecordResponse {
    1: required model.BaseResp Base
    2: required model.FollowupRecord record
}

// 获取计划随访记录请求
struct GetPlanFollowupRecordsRequest {
    1: required i64 planId
}

// 获取计划随访记录响应
struct GetPlanFollowupRecordsResponse {
    1: required model.BaseResp Base
    2: required list<model.FollowupRecord> records
}

// 添加用药监测请求
struct AddMedicationMonitorRequest {
    1: required model.MedicationMonitor monitor
}

// 添加用药监测响应
struct AddMedicationMonitorResponse {
    1: required model.BaseResp Base
    2: required i64 monitorId
}

// 获取随访用药监测请求
struct GetFollowupMedicationsRequest {
    1: required i64 followupId
}

// 获取随访用药监测响应
struct GetFollowupMedicationsResponse {
    1: required model.BaseResp Base
    2: required list<model.MedicationMonitor> medications
}

service FollowupService {
    // 随访计划管理
    CreateFollowupPlanResponse CreateFollowupPlan(1: CreateFollowupPlanRequest req)
    GetFollowupPlanResponse GetFollowupPlan(1: GetFollowupPlanRequest req)
    GetFollowupPlansResponse GetFollowupPlans(1: GetFollowupPlansRequest req)
    GetPatientFollowupPlansResponse GetPatientFollowupPlans(1: GetPatientFollowupPlansRequest req)
    
    // 随访记录管理
    CreateFollowupRecordResponse CreateFollowupRecord(1: CreateFollowupRecordRequest req)
    GetFollowupRecordResponse GetFollowupRecord(1: GetFollowupRecordRequest req)
    GetPlanFollowupRecordsResponse GetPlanFollowupRecords(1: GetPlanFollowupRecordsRequest req)
    
    // 用药监测管理
    AddMedicationMonitorResponse AddMedicationMonitor(1: AddMedicationMonitorRequest req)
    GetFollowupMedicationsResponse GetFollowupMedications(1: GetFollowupMedicationsRequest req)
} 