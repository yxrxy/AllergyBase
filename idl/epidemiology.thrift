namespace go epidemiology

include "model.thrift"

// 创建环境暴露记录请求
struct CreateEnvironmentExposureRequest {
    1: required model.EnvironmentExposure exposure
}

// 创建环境暴露记录响应
struct CreateEnvironmentExposureResponse {
    1: required model.BaseResp Base
    2: required i64 exposureId
}

// 获取环境暴露记录请求
struct GetEnvironmentExposureRequest {
    1: required i64 patientId
}

// 获取环境暴露记录响应
struct GetEnvironmentExposureResponse {
    1: required model.BaseResp Base
    2: required model.EnvironmentExposure exposure
}

// 获取患者环境暴露记录列表请求
struct GetPatientEnvironmentExposuresRequest {
    1: required i64 patientId
    2: required i32 offset
    3: required i32 limit
}

// 获取患者环境暴露记录列表响应
struct GetPatientEnvironmentExposuresResponse {
    1: required model.BaseResp Base
    2: required list<model.EnvironmentExposure> exposures
    3: required i32 total
}

// 添加环境监测数据请求
struct AddEnvironmentMonitorRequest {
    1: required model.EnvironmentMonitor monitor
}

// 添加环境监测数据响应
struct AddEnvironmentMonitorResponse {
    1: required model.BaseResp Base
    2: required i64 monitorId
}

// 获取环境监测数据请求
struct GetEnvironmentMonitorsRequest {
    1: required string locationCode
    2: required string startTime
    3: required string endTime
}

// 获取环境监测数据响应
struct GetEnvironmentMonitorsResponse {
    1: required model.BaseResp Base
    2: required list<model.EnvironmentMonitor> monitors
}

// 创建生活方式调查请求
struct CreateLifestyleSurveyRequest {
    1: required model.LifestyleSurvey survey
}

// 创建生活方式调查响应
struct CreateLifestyleSurveyResponse {
    1: required model.BaseResp Base
    2: required i64 surveyId
}

// 获取生活方式调查请求
struct GetLifestyleSurveyRequest {
    1: required i64 patientId
}

// 获取生活方式调查响应
struct GetLifestyleSurveyResponse {
    1: required model.BaseResp Base
    2: required model.LifestyleSurvey survey
}

service EpidemiologyService {
    // 环境暴露管理
    CreateEnvironmentExposureResponse CreateEnvironmentExposure(1: CreateEnvironmentExposureRequest req)
    GetEnvironmentExposureResponse GetEnvironmentExposure(1: GetEnvironmentExposureRequest req)
    GetPatientEnvironmentExposuresResponse GetPatientEnvironmentExposures(1: GetPatientEnvironmentExposuresRequest req)
    
    // 环境监测数据管理
    AddEnvironmentMonitorResponse AddEnvironmentMonitor(1: AddEnvironmentMonitorRequest req)
    GetEnvironmentMonitorsResponse GetEnvironmentMonitors(1: GetEnvironmentMonitorsRequest req)
    
    // 生活方式调查管理
    CreateLifestyleSurveyResponse CreateLifestyleSurvey(1: CreateLifestyleSurveyRequest req)
    GetLifestyleSurveyResponse GetLifestyleSurvey(1: GetLifestyleSurveyRequest req)
} 