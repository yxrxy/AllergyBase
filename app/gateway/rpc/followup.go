package rpc

import (
	"context"
	"log"

	"github.com/yxrxy/AllergyBase/kitex_gen/followup"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
	"github.com/yxrxy/AllergyBase/pkg/errno"
)

// InitFollowupRPC 初始化随访服务客户端
func InitFollowupRPC() {
	c, err := client.InitFollowupRPC()
	if err != nil {
		log.Fatalf("初始化随访服务客户端失败: %v", err)
	}
	followupClient = *c
}

// CreateFollowupPlanRPC 创建随访计划
func CreateFollowupPlanRPC(ctx context.Context, req *followup.CreateFollowupPlanRequest) (int64, error) {
	resp, err := followupClient.CreateFollowupPlan(ctx, req)
	if err != nil {
		log.Printf("创建随访计划RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.PlanId, nil
}

// GetFollowupPlanRPC 获取随访计划
func GetFollowupPlanRPC(ctx context.Context, req *followup.GetFollowupPlanRequest) (*followup.GetFollowupPlanResponse, error) {
	resp, err := followupClient.GetFollowupPlan(ctx, req)
	if err != nil {
		log.Printf("获取随访计划RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// GetFollowupPlansRPC 获取所有随访计划列表
func GetFollowupPlansRPC(ctx context.Context, req *followup.GetFollowupPlansRequest) (*followup.GetFollowupPlansResponse, error) {
	resp, err := followupClient.GetFollowupPlans(ctx, req)
	if err != nil {
		log.Printf("获取所有随访计划列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// GetPatientFollowupPlansRPC 获取患者随访计划列表
func GetPatientFollowupPlansRPC(ctx context.Context, req *followup.GetPatientFollowupPlansRequest) (*followup.GetPatientFollowupPlansResponse, error) {
	resp, err := followupClient.GetPatientFollowupPlans(ctx, req)
	if err != nil {
		log.Printf("获取患者随访计划列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// CreateFollowupRecordRPC 创建随访记录
func CreateFollowupRecordRPC(ctx context.Context, req *followup.CreateFollowupRecordRequest) (int64, error) {
	resp, err := followupClient.CreateFollowupRecord(ctx, req)
	if err != nil {
		log.Printf("创建随访记录RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.RecordId, nil
}

// GetFollowupRecordRPC 获取随访记录
func GetFollowupRecordRPC(ctx context.Context, req *followup.GetFollowupRecordRequest) (*followup.GetFollowupRecordResponse, error) {
	resp, err := followupClient.GetFollowupRecord(ctx, req)
	if err != nil {
		log.Printf("获取随访记录RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// GetPlanFollowupRecordsRPC 获取计划随访记录列表
func GetPlanFollowupRecordsRPC(ctx context.Context, req *followup.GetPlanFollowupRecordsRequest) (*followup.GetPlanFollowupRecordsResponse, error) {
	resp, err := followupClient.GetPlanFollowupRecords(ctx, req)
	if err != nil {
		log.Printf("获取计划随访记录列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// AddMedicationMonitorRPC 添加用药监测
func AddMedicationMonitorRPC(ctx context.Context, req *followup.AddMedicationMonitorRequest) (int64, error) {
	resp, err := followupClient.AddMedicationMonitor(ctx, req)
	if err != nil {
		log.Printf("添加用药监测RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.MonitorId, nil
}

// GetFollowupMedicationsRPC 获取随访用药信息
func GetFollowupMedicationsRPC(ctx context.Context, req *followup.GetFollowupMedicationsRequest) (*followup.GetFollowupMedicationsResponse, error) {
	resp, err := followupClient.GetFollowupMedications(ctx, req)
	if err != nil {
		log.Printf("获取随访用药信息RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}
