package rpc

import (
	"context"
	"strings"

	"github.com/yxrxy/AllergyBase/app/biobank/controllers/rpc/pack"
	"github.com/yxrxy/AllergyBase/app/biobank/domain/model"
	"github.com/yxrxy/AllergyBase/app/biobank/usecase"
	"github.com/yxrxy/AllergyBase/kitex_gen/biobank"
	"github.com/yxrxy/AllergyBase/pkg/base"
)

// BiobankHandler Biobank RPC服务处理器
type BiobankHandler struct {
	BiobankUseCase usecase.BiobankUseCase
}

// NewBiobankHandler 创建新的Biobank处理器
func NewBiobankHandler(biobankUseCase usecase.BiobankUseCase) *BiobankHandler {
	return &BiobankHandler{
		BiobankUseCase: biobankUseCase,
	}
}

// CreateSample 创建样本
func (h *BiobankHandler) CreateSample(ctx context.Context, req *biobank.CreateSampleRequest) (r *biobank.CreateSampleResponse, err error) {
	r = new(biobank.CreateSampleResponse)

	sample := pack.BuildSampleFromRequest(req)
	var createdSample *model.Sample
	if createdSample, err = h.BiobankUseCase.CreateSample(ctx, sample); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.SampleId = createdSample.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetSample 获取样本信息
func (h *BiobankHandler) GetSample(ctx context.Context, req *biobank.GetSampleRequest) (r *biobank.GetSampleResponse, err error) {
	r = new(biobank.GetSampleResponse)

	var sample *model.Sample
	if sample, err = h.BiobankUseCase.GetSample(ctx, req.SampleId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Sample = pack.Sample(sample)
	r.Base = base.BuildBaseResp(err)
	return
}

// ListSamples 获取样本列表
func (h *BiobankHandler) ListSamples(ctx context.Context, req *biobank.ListSamplesRequest) (r *biobank.ListSamplesResponse, err error) {
	r = new(biobank.ListSamplesResponse)

	// 设置默认分页参数
	offset := 0
	limit := 20
	if req.Offset != nil {
		offset = int(*req.Offset)
	}
	if req.Limit != nil {
		limit = int(*req.Limit)
	}

	var samples []*model.Sample
	if samples, _, err = h.BiobankUseCase.ListSamples(ctx, offset, limit); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	// 根据搜索条件过滤样本
	var filteredSamples []*model.Sample
	for _, sample := range samples {
		// 按样本编号过滤
		if req.SampleNo != nil && *req.SampleNo != "" {
			if !strings.Contains(sample.SampleNo, *req.SampleNo) {
				continue
			}
		}
		// 按患者ID过滤
		if req.PatientId != nil && *req.PatientId > 0 {
			if sample.PatientID != *req.PatientId {
				continue
			}
		}
		filteredSamples = append(filteredSamples, sample)
	}

	r.Samples = pack.SampleList(filteredSamples)
	r.Total = int64(len(filteredSamples))
	r.Base = base.BuildBaseResp(err)
	return
}

// GetPatientSamples 获取患者样本列表
func (h *BiobankHandler) GetPatientSamples(ctx context.Context, req *biobank.GetPatientSamplesRequest) (r *biobank.GetPatientSamplesResponse, err error) {
	r = new(biobank.GetPatientSamplesResponse)

	var samples []*model.Sample
	if samples, _, err = h.BiobankUseCase.ListSamples(ctx, 0, 100); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	// 过滤出指定患者的样本
	var patientSamples []*model.Sample
	for _, sample := range samples {
		if sample.PatientID == req.PatientId {
			patientSamples = append(patientSamples, sample)
		}
	}

	r.Samples = pack.SampleList(patientSamples)
	r.Base = base.BuildBaseResp(err)
	return
}

// AddStorageInfo 添加存储信息
func (h *BiobankHandler) AddStorageInfo(ctx context.Context, req *biobank.AddStorageInfoRequest) (r *biobank.AddStorageInfoResponse, err error) {
	r = new(biobank.AddStorageInfoResponse)

	storageInfo := pack.BuildStorageInfoFromRequest(req)
	var createdStorageInfo *model.StorageInfo
	if createdStorageInfo, err = h.BiobankUseCase.CreateStorageInfo(ctx, storageInfo); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.StorageId = createdStorageInfo.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetStorageInfo 获取存储信息
func (h *BiobankHandler) GetStorageInfo(ctx context.Context, req *biobank.GetStorageInfoRequest) (r *biobank.GetStorageInfoResponse, err error) {
	r = new(biobank.GetStorageInfoResponse)

	var storageInfo *model.StorageInfo
	if storageInfo, err = h.BiobankUseCase.GetStorageInfo(ctx, req.SampleId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.Storage = pack.StorageInfo(storageInfo)
	r.Base = base.BuildBaseResp(err)
	return
}

// AddGenomicData 添加基因组数据
func (h *BiobankHandler) AddGenomicData(ctx context.Context, req *biobank.AddGenomicDataRequest) (r *biobank.AddGenomicDataResponse, err error) {
	r = new(biobank.AddGenomicDataResponse)

	genomicData := pack.BuildGenomicDataFromRequest(req)
	var createdGenomicData *model.GenomicData
	if createdGenomicData, err = h.BiobankUseCase.CreateGenomicData(ctx, genomicData); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.DataId = createdGenomicData.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetSampleGenomicData 获取样本基因组数据
func (h *BiobankHandler) GetSampleGenomicData(ctx context.Context, req *biobank.GetSampleGenomicDataRequest) (r *biobank.GetSampleGenomicDataResponse, err error) {
	r = new(biobank.GetSampleGenomicDataResponse)

	// 这里暂时返回空数据，实际应该根据sampleId查询相关的基因组数据
	var genomicData []*model.GenomicData
	genomicData = []*model.GenomicData{}

	r.Data = pack.GenomicDataList(genomicData)
	r.Base = base.BuildBaseResp(err)
	return
}

// AddProteomicsData 添加蛋白质组数据
func (h *BiobankHandler) AddProteomicsData(ctx context.Context, req *biobank.AddProteomicsDataRequest) (r *biobank.AddProteomicsDataResponse, err error) {
	r = new(biobank.AddProteomicsDataResponse)

	proteomicsData := pack.BuildProteomicsDataFromRequest(req)
	var createdProteomicsData *model.ProteomicsData
	if createdProteomicsData, err = h.BiobankUseCase.CreateProteomicsData(ctx, proteomicsData); err != nil {
		r.Base = base.BuildBaseResp(err)
		return
	}

	r.DataId = createdProteomicsData.ID
	r.Base = base.BuildBaseResp(err)
	return
}

// GetSampleProteomicsData 获取样本蛋白质组数据
func (h *BiobankHandler) GetSampleProteomicsData(ctx context.Context, req *biobank.GetSampleProteomicsDataRequest) (r *biobank.GetSampleProteomicsDataResponse, err error) {
	r = new(biobank.GetSampleProteomicsDataResponse)

	// 这里暂时返回空数据，实际应该根据sampleId查询相关的蛋白质组数据
	var proteomicsData []*model.ProteomicsData
	proteomicsData = []*model.ProteomicsData{}

	r.Data = pack.ProteomicsDataList(proteomicsData)
	r.Base = base.BuildBaseResp(err)
	return
}
