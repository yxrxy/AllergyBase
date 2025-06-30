namespace go api.epidemiology

include "../epidemiology.thrift"

// API 服务
service EpidemiologyAPI {
    // 环境暴露管理接口
    epidemiology.CreateEnvironmentExposureResponse CreateEnvironmentExposure(1: epidemiology.CreateEnvironmentExposureRequest request) (api.post="/api/v1/environment/exposure")
    epidemiology.GetEnvironmentExposureResponse GetEnvironmentExposure(1: epidemiology.GetEnvironmentExposureRequest request) (api.get="/api/v1/patient/:patientId/environment/exposure")
    epidemiology.GetPatientEnvironmentExposuresResponse GetPatientEnvironmentExposures(1: epidemiology.GetPatientEnvironmentExposuresRequest request) (api.get="/api/v1/environment/exposures")
    
    // 环境监测数据管理接口
    epidemiology.AddEnvironmentMonitorResponse AddEnvironmentMonitor(1: epidemiology.AddEnvironmentMonitorRequest request) (api.post="/api/v1/environment/monitor")
    epidemiology.GetEnvironmentMonitorsResponse GetEnvironmentMonitors(1: epidemiology.GetEnvironmentMonitorsRequest request) (api.get="/api/v1/environment/monitors")
    
    // 生活方式调查管理接口
    epidemiology.CreateLifestyleSurveyResponse CreateLifestyleSurvey(1: epidemiology.CreateLifestyleSurveyRequest request) (api.post="/api/v1/lifestyle/survey")
    epidemiology.GetLifestyleSurveyResponse GetLifestyleSurvey(1: epidemiology.GetLifestyleSurveyRequest request) (api.get="/api/v1/patient/:patientId/lifestyle/survey")
} 