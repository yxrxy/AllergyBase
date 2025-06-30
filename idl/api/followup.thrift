 namespace go api.followup

include "../followup.thrift"

// API 服务
service FollowupAPI {
    // 随访计划管理接口
    followup.CreateFollowupPlanResponse CreateFollowupPlan(1: followup.CreateFollowupPlanRequest request) (api.post="/api/v1/followup/plan")
    followup.GetFollowupPlanResponse GetFollowupPlan(1: followup.GetFollowupPlanRequest request) (api.get="/api/v1/followup/plan/:planId")
    followup.GetFollowupPlansResponse GetFollowupPlans(1: followup.GetFollowupPlansRequest request) (api.get="/api/v1/followup/plans")
    followup.GetPatientFollowupPlansResponse GetPatientFollowupPlans(1: followup.GetPatientFollowupPlansRequest request) (api.get="/api/v1/patient/:patientId/followup/plans")
    
    // 随访记录管理接口
    followup.CreateFollowupRecordResponse CreateFollowupRecord(1: followup.CreateFollowupRecordRequest request) (api.post="/api/v1/followup/record")
    followup.GetFollowupRecordResponse GetFollowupRecord(1: followup.GetFollowupRecordRequest request) (api.get="/api/v1/followup/record/:recordId")
    followup.GetPlanFollowupRecordsResponse GetPlanFollowupRecords(1: followup.GetPlanFollowupRecordsRequest request) (api.get="/api/v1/followup/plan/:planId/records")
    
    // 用药监测管理接口
    followup.AddMedicationMonitorResponse AddMedicationMonitor(1: followup.AddMedicationMonitorRequest request) (api.post="/api/v1/followup/medication")
    followup.GetFollowupMedicationsResponse GetFollowupMedications(1: followup.GetFollowupMedicationsRequest request) (api.get="/api/v1/followup/:followupId/medications")
}