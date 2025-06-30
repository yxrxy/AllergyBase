package rpc

import (
	"context"
	"log"

	"github.com/yxrxy/AllergyBase/kitex_gen/biobank"
	"github.com/yxrxy/AllergyBase/pkg/base/client"
	"github.com/yxrxy/AllergyBase/pkg/errno"
)

// InitBiobankRPC 初始化生物样本库服务客户端
func InitBiobankRPC() {
	c, err := client.InitBiobankRPC()
	if err != nil {
		log.Fatalf("初始化生物样本库服务客户端失败: %v", err)
	}
	biobankClient = *c
}

// CreateSampleRPC 创建样本
func CreateSampleRPC(ctx context.Context, req *biobank.CreateSampleRequest) (int64, error) {
	resp, err := biobankClient.CreateSample(ctx, req)
	if err != nil {
		log.Printf("创建样本RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.SampleId, nil
}

// GetSampleRPC 获取样本信息
func GetSampleRPC(ctx context.Context, req *biobank.GetSampleRequest) (*biobank.GetSampleResponse, error) {
	resp, err := biobankClient.GetSample(ctx, req)
	if err != nil {
		log.Printf("获取样本信息RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// GetPatientSamplesRPC 获取患者样本列表
func GetPatientSamplesRPC(ctx context.Context, req *biobank.GetPatientSamplesRequest) (*biobank.GetPatientSamplesResponse, error) {
	resp, err := biobankClient.GetPatientSamples(ctx, req)
	if err != nil {
		log.Printf("获取患者样本列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// AddStorageInfoRPC 添加存储信息
func AddStorageInfoRPC(ctx context.Context, req *biobank.AddStorageInfoRequest) (int64, error) {
	resp, err := biobankClient.AddStorageInfo(ctx, req)
	if err != nil {
		log.Printf("添加存储信息RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.StorageId, nil
}

// GetStorageInfoRPC 获取存储信息
func GetStorageInfoRPC(ctx context.Context, req *biobank.GetStorageInfoRequest) (*biobank.GetStorageInfoResponse, error) {
	resp, err := biobankClient.GetStorageInfo(ctx, req)
	if err != nil {
		log.Printf("获取存储信息RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// AddGenomicDataRPC 添加基因组数据
func AddGenomicDataRPC(ctx context.Context, req *biobank.AddGenomicDataRequest) (int64, error) {
	resp, err := biobankClient.AddGenomicData(ctx, req)
	if err != nil {
		log.Printf("添加基因组数据RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.DataId, nil
}

// GetSampleGenomicDataRPC 获取样本基因组数据
func GetSampleGenomicDataRPC(ctx context.Context, req *biobank.GetSampleGenomicDataRequest) (*biobank.GetSampleGenomicDataResponse, error) {
	resp, err := biobankClient.GetSampleGenomicData(ctx, req)
	if err != nil {
		log.Printf("获取样本基因组数据RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// AddProteomicsDataRPC 添加蛋白质组数据
func AddProteomicsDataRPC(ctx context.Context, req *biobank.AddProteomicsDataRequest) (int64, error) {
	resp, err := biobankClient.AddProteomicsData(ctx, req)
	if err != nil {
		log.Printf("添加蛋白质组数据RPC调用失败: %v", err)
		return 0, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp.DataId, nil
}

// GetSampleProteomicsDataRPC 获取样本蛋白质组数据
func GetSampleProteomicsDataRPC(ctx context.Context, req *biobank.GetSampleProteomicsDataRequest) (*biobank.GetSampleProteomicsDataResponse, error) {
	resp, err := biobankClient.GetSampleProteomicsData(ctx, req)
	if err != nil {
		log.Printf("获取样本蛋白质组数据RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}

// ListSamplesRPC 获取样本列表
func ListSamplesRPC(ctx context.Context, req *biobank.ListSamplesRequest) (*biobank.ListSamplesResponse, error) {
	resp, err := biobankClient.ListSamples(ctx, req)
	if err != nil {
		log.Printf("获取样本列表RPC调用失败: %v", err)
		return nil, errno.InternalServiceError.WithError(err)
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return resp, nil
}
