package pack

import (
	"time"

	"github.com/yxrxy/AllergyBase/app/biobank/domain/model"
	"github.com/yxrxy/AllergyBase/kitex_gen/biobank"
	rpcmodel "github.com/yxrxy/AllergyBase/kitex_gen/model"
)

// Sample 将数据库样本实体转换为RPC响应实体
func Sample(sample *model.Sample) *rpcmodel.SampleInfo {
	if sample == nil {
		return nil
	}

	var collectionTime *string
	if sample.CollectionTime != nil && !sample.CollectionTime.IsZero() {
		t := sample.CollectionTime.Format("2006-01-02 15:04:05")
		collectionTime = &t
	}

	return &rpcmodel.SampleInfo{
		Id:               sample.ID,
		PatientId:        sample.PatientID,
		SampleNo:         sample.SampleNo,
		SampleType:       &sample.SampleType,
		CollectionTime:   collectionTime,
		CollectionMethod: &sample.CollectionMethod,
		Processor:        &sample.Processor,
	}
}

// StorageInfo 将数据库存储信息实体转换为RPC响应实体
func StorageInfo(storage *model.StorageInfo) *rpcmodel.StorageInfo {
	if storage == nil {
		return nil
	}

	var storageTime *string
	if storage.StorageTime != nil && !storage.StorageTime.IsZero() {
		t := storage.StorageTime.Format("2006-01-02 15:04:05")
		storageTime = &t
	}

	return &rpcmodel.StorageInfo{
		Id:                 storage.ID,
		SampleId:           storage.SampleID,
		StorageLocation:    &storage.StorageLocation,
		StorageTemperature: &storage.StorageTemperature,
		StorageTime:        storageTime,
		Status:             &storage.Status,
	}
}

// GenomicData 将数据库基因组数据实体转换为RPC响应实体
func GenomicData(genomic *model.GenomicData) *rpcmodel.GenomicData {
	if genomic == nil {
		return nil
	}

	return &rpcmodel.GenomicData{
		Id:               genomic.ID,
		SampleId:         genomic.SampleID,
		SequencePlatform: &genomic.SequencePlatform,
		SequenceType:     &genomic.SequenceType,
		GeneType:         &genomic.GeneType,
		ResultFilePath:   &genomic.ResultFilePath,
	}
}

// ProteomicsData 将数据库蛋白质组数据实体转换为RPC响应实体
func ProteomicsData(proteomics *model.ProteomicsData) *rpcmodel.ProteomicsData {
	if proteomics == nil {
		return nil
	}

	return &rpcmodel.ProteomicsData{
		Id:               proteomics.ID,
		SampleId:         proteomics.SampleID,
		AnalysisPlatform: &proteomics.AnalysisPlatform,
		ProteinMarker:    &proteomics.ProteinMarker,
		Concentration:    &proteomics.Concentration,
	}
}

// SampleList 将样本列表转换为RPC响应实体列表
func SampleList(samples []*model.Sample) []*rpcmodel.SampleInfo {
	if samples == nil {
		return nil
	}

	result := make([]*rpcmodel.SampleInfo, 0, len(samples))
	for _, sample := range samples {
		if sample != nil {
			result = append(result, Sample(sample))
		}
	}
	return result
}

// StorageInfoList 将存储信息列表转换为RPC响应实体列表
func StorageInfoList(storages []*model.StorageInfo) []*rpcmodel.StorageInfo {
	if storages == nil {
		return nil
	}

	result := make([]*rpcmodel.StorageInfo, len(storages))
	for i, storage := range storages {
		result[i] = StorageInfo(storage)
	}
	return result
}

// GenomicDataList 将基因组数据列表转换为RPC响应实体列表
func GenomicDataList(genomics []*model.GenomicData) []*rpcmodel.GenomicData {
	if genomics == nil {
		return nil
	}

	result := make([]*rpcmodel.GenomicData, len(genomics))
	for i, genomic := range genomics {
		result[i] = GenomicData(genomic)
	}
	return result
}

// ProteomicsDataList 将蛋白质组数据列表转换为RPC响应实体列表
func ProteomicsDataList(proteomics []*model.ProteomicsData) []*rpcmodel.ProteomicsData {
	if proteomics == nil {
		return nil
	}

	result := make([]*rpcmodel.ProteomicsData, len(proteomics))
	for i, protein := range proteomics {
		result[i] = ProteomicsData(protein)
	}
	return result
}

// BuildSampleFromRequest 将RPC请求转换为数据库实体
func BuildSampleFromRequest(req interface{}) *model.Sample {
	switch r := req.(type) {
	case *biobank.CreateSampleRequest:
		if r.Sample == nil {
			return &model.Sample{}
		}

		sample := &model.Sample{
			ID:        r.Sample.Id,
			PatientID: r.Sample.PatientId,
			SampleNo:  r.Sample.SampleNo,
		}

		// 处理可选字段
		if r.Sample.SampleType != nil {
			sample.SampleType = *r.Sample.SampleType
		}
		if r.Sample.CollectionMethod != nil {
			sample.CollectionMethod = *r.Sample.CollectionMethod
		}
		if r.Sample.Processor != nil {
			sample.Processor = *r.Sample.Processor
		}

		// 处理时间字段 - 只有当时间不为空且有效时才设置
		if r.Sample.CollectionTime != nil && *r.Sample.CollectionTime != "" {
			if t, err := time.Parse(time.RFC3339, *r.Sample.CollectionTime); err == nil && !t.IsZero() {
				sample.CollectionTime = &t
			}
		}

		return sample
	default:
		return &model.Sample{}
	}
}

// BuildStorageInfoFromRequest 将RPC请求转换为数据库实体
func BuildStorageInfoFromRequest(req interface{}) *model.StorageInfo {
	switch r := req.(type) {
	case *biobank.AddStorageInfoRequest:
		if r.Storage == nil {
			return &model.StorageInfo{}
		}

		storage := &model.StorageInfo{
			ID:       r.Storage.Id,
			SampleID: r.Storage.SampleId,
		}

		// 处理可选字段
		if r.Storage.StorageLocation != nil {
			storage.StorageLocation = *r.Storage.StorageLocation
		}
		if r.Storage.StorageTemperature != nil {
			storage.StorageTemperature = *r.Storage.StorageTemperature
		}
		if r.Storage.Status != nil {
			storage.Status = *r.Storage.Status
		}

		// 处理时间字段 - 只有当时间不为空且有效时才设置
		if r.Storage.StorageTime != nil && *r.Storage.StorageTime != "" {
			if t, err := time.Parse(time.RFC3339, *r.Storage.StorageTime); err == nil && !t.IsZero() {
				storage.StorageTime = &t
			}
		}

		return storage
	default:
		return &model.StorageInfo{}
	}
}

// BuildGenomicDataFromRequest 将RPC请求转换为数据库实体
func BuildGenomicDataFromRequest(req interface{}) *model.GenomicData {
	// 这里需要根据实际的请求结构进行转换
	// 临时返回空实体避免编译错误
	return &model.GenomicData{}
}

// BuildProteomicsDataFromRequest 将RPC请求转换为数据库实体
func BuildProteomicsDataFromRequest(req interface{}) *model.ProteomicsData {
	// 这里需要根据实际的请求结构进行转换
	// 临时返回空实体避免编译错误
	return &model.ProteomicsData{}
}
