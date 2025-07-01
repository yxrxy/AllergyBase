package rpc

import (
	"context"
	"log"

	"github.com/yxrxy/AllergyBase/kitex_gen/epidemiology"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
	"github.com/yxrxy/AllergyBase/pkg/errno"
)

// InitEpidemiologyRPC 初始化流行病学服务客户端
func InitEpidemiologyRPC() {
	c, err := client.InitEpidemiologyRPC()
	if err != nil {
		log.Fatalf("初始化流行病学服务客户端失败: %v", err)
	}
	epidemiologyClient = *c
}

// CreateEnvironmentExposureRPC 创建环境暴露记录
func CreateEnvironmentExposureRPC(ctx context.Context, req *epidemiology.CreateEnvironmentExposureRequest) (int64, error) {
	resp, err := epidemiologyClient.CreateEnvironmentExposure(ctx, req)
	if err != nil {
		log.Printf("创建环境暴露记录RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.ExposureId, nil
}

// GetEnvironmentExposureRPC 获取环境暴露记录
func GetEnvironmentExposureRPC(ctx context.Context, req *epidemiology.GetEnvironmentExposureRequest) (*epidemiology.GetEnvironmentExposureResponse, error) {
	resp, err := epidemiologyClient.GetEnvironmentExposure(ctx, req)
	if err != nil {
		log.Printf("获取环境暴露记录RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// GetPatientEnvironmentExposuresRPC 获取患者环境暴露记录列表
func GetPatientEnvironmentExposuresRPC(ctx context.Context, req *epidemiology.GetPatientEnvironmentExposuresRequest) (*epidemiology.GetPatientEnvironmentExposuresResponse, error) {
	resp, err := epidemiologyClient.GetPatientEnvironmentExposures(ctx, req)
	if err != nil {
		log.Printf("获取患者环境暴露记录列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// AddEnvironmentMonitorRPC 添加环境监测
func AddEnvironmentMonitorRPC(ctx context.Context, req *epidemiology.AddEnvironmentMonitorRequest) (int64, error) {
	resp, err := epidemiologyClient.AddEnvironmentMonitor(ctx, req)
	if err != nil {
		log.Printf("添加环境监测RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.MonitorId, nil
}

// GetEnvironmentMonitorsRPC 获取环境监测列表
func GetEnvironmentMonitorsRPC(ctx context.Context, req *epidemiology.GetEnvironmentMonitorsRequest) (*epidemiology.GetEnvironmentMonitorsResponse, error) {
	resp, err := epidemiologyClient.GetEnvironmentMonitors(ctx, req)
	if err != nil {
		log.Printf("获取环境监测列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// CreateLifestyleSurveyRPC 创建生活方式调查
func CreateLifestyleSurveyRPC(ctx context.Context, req *epidemiology.CreateLifestyleSurveyRequest) (int64, error) {
	resp, err := epidemiologyClient.CreateLifestyleSurvey(ctx, req)
	if err != nil {
		log.Printf("创建生活方式调查RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.SurveyId, nil
}

// GetLifestyleSurveyRPC 获取生活方式调查
func GetLifestyleSurveyRPC(ctx context.Context, req *epidemiology.GetLifestyleSurveyRequest) (*epidemiology.GetLifestyleSurveyResponse, error) {
	resp, err := epidemiologyClient.GetLifestyleSurvey(ctx, req)
	if err != nil {
		log.Printf("获取生活方式调查RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}
