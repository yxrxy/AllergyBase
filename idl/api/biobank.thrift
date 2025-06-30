namespace go api.biobank

include "../biobank.thrift"

// API 服务
service BiobankAPI {
    // 样本信息管理接口
    biobank.CreateSampleResponse CreateSample(1: biobank.CreateSampleRequest request) (api.post="/api/v1/sample")
    biobank.GetSampleResponse GetSample(1: biobank.GetSampleRequest request) (api.get="/api/v1/sample/:sampleId")
    biobank.ListSamplesResponse ListSamples(1: biobank.ListSamplesRequest request) (api.get="/api/v1/samples")
    biobank.GetPatientSamplesResponse GetPatientSamples(1: biobank.GetPatientSamplesRequest request) (api.get="/api/v1/patient/:patientId/samples")
    
    // 存储信息管理接口
    biobank.AddStorageInfoResponse AddStorageInfo(1: biobank.AddStorageInfoRequest request) (api.post="/api/v1/sample/storage")
    biobank.GetStorageInfoResponse GetStorageInfo(1: biobank.GetStorageInfoRequest request) (api.get="/api/v1/sample/:sampleId/storage")
    
    // 基因组数据管理接口
    biobank.AddGenomicDataResponse AddGenomicData(1: biobank.AddGenomicDataRequest request) (api.post="/api/v1/sample/genomic")
    biobank.GetSampleGenomicDataResponse GetSampleGenomicData(1: biobank.GetSampleGenomicDataRequest request) (api.get="/api/v1/sample/:sampleId/genomic")
    
    // 蛋白组学数据管理接口
    biobank.AddProteomicsDataResponse AddProteomicsData(1: biobank.AddProteomicsDataRequest request) (api.post="/api/v1/sample/proteomics")
    biobank.GetSampleProteomicsDataResponse GetSampleProteomicsData(1: biobank.GetSampleProteomicsDataRequest request) (api.get="/api/v1/sample/:sampleId/proteomics")
} 