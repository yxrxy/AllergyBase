namespace go biobank

include "model.thrift"

// 创建样本请求
struct CreateSampleRequest {
    1: required model.SampleInfo sample
}

// 创建样本响应
struct CreateSampleResponse {
    1: required model.BaseResp Base
    2: required i64 sampleId
}

// 获取样本请求
struct GetSampleRequest {
    1: required i64 sampleId
}

// 获取样本响应
struct GetSampleResponse {
    1: required model.BaseResp Base
    2: required model.SampleInfo sample
}

// 获取患者样本请求
struct GetPatientSamplesRequest {
    1: required i64 patientId
}

// 获取患者样本响应
struct GetPatientSamplesResponse {
    1: required model.BaseResp Base
    2: required list<model.SampleInfo> samples
}

// 获取样本列表请求
struct ListSamplesRequest {
    1: optional i32 offset
    2: optional i32 limit
    3: optional string sampleNo
    4: optional i64 patientId
}

// 获取样本列表响应
struct ListSamplesResponse {
    1: required model.BaseResp Base
    2: required list<model.SampleInfo> samples
    3: required i64 total
}

// 添加存储信息请求
struct AddStorageInfoRequest {
    1: required model.StorageInfo storage
}

// 添加存储信息响应
struct AddStorageInfoResponse {
    1: required model.BaseResp Base
    2: required i64 storageId
}

// 获取存储信息请求
struct GetStorageInfoRequest {
    1: required i64 sampleId
}

// 获取存储信息响应
struct GetStorageInfoResponse {
    1: required model.BaseResp Base
    2: required model.StorageInfo storage
}

// 添加基因组数据请求
struct AddGenomicDataRequest {
    1: required model.GenomicData data
}

// 添加基因组数据响应
struct AddGenomicDataResponse {
    1: required model.BaseResp Base
    2: required i64 dataId
}

// 获取样本基因组数据请求
struct GetSampleGenomicDataRequest {
    1: required i64 sampleId
}

// 获取样本基因组数据响应
struct GetSampleGenomicDataResponse {
    1: required model.BaseResp Base
    2: required list<model.GenomicData> data
}

// 添加蛋白组学数据请求
struct AddProteomicsDataRequest {
    1: required model.ProteomicsData data
}

// 添加蛋白组学数据响应
struct AddProteomicsDataResponse {
    1: required model.BaseResp Base
    2: required i64 dataId
}

// 获取样本蛋白组学数据请求
struct GetSampleProteomicsDataRequest {
    1: required i64 sampleId
}

// 获取样本蛋白组学数据响应
struct GetSampleProteomicsDataResponse {
    1: required model.BaseResp Base
    2: required list<model.ProteomicsData> data
}

service BiobankService {
    // 样本信息管理
    CreateSampleResponse CreateSample(1: CreateSampleRequest req)
    GetSampleResponse GetSample(1: GetSampleRequest req)
    ListSamplesResponse ListSamples(1: ListSamplesRequest req)
    GetPatientSamplesResponse GetPatientSamples(1: GetPatientSamplesRequest req)
    
    // 存储信息管理
    AddStorageInfoResponse AddStorageInfo(1: AddStorageInfoRequest req)
    GetStorageInfoResponse GetStorageInfo(1: GetStorageInfoRequest req)
    
    // 基因组数据管理
    AddGenomicDataResponse AddGenomicData(1: AddGenomicDataRequest req)
    GetSampleGenomicDataResponse GetSampleGenomicData(1: GetSampleGenomicDataRequest req)
    
    // 蛋白组学数据管理
    AddProteomicsDataResponse AddProteomicsData(1: AddProteomicsDataRequest req)
    GetSampleProteomicsDataResponse GetSampleProteomicsData(1: GetSampleProteomicsDataRequest req)
} 