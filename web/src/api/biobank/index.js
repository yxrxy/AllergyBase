import request from '@/utils/request'

export const biobankApi = {
  // 样本信息管理
  // 创建样本
  createSample(data) {
    return request({
      url: '/api/v1/sample',
      method: 'post',
      data: {
        sample: {
          id: 0, // 临时ID，后端会忽略并自动生成
          ...data
        }
      }
    })
  },

  // 获取样本信息
  getSample(sampleId) {
    return request({
      url: `/api/v1/sample/${sampleId}`,
      method: 'get',
      params: { sampleId: sampleId }
    })
  },

  // 获取所有样本列表（用于样本管理页面）
  listSamples(params) {
    return request({
      url: '/api/v1/samples',
      method: 'get',
      params
    })
  },

  // 获取患者样本
  getPatientSamples(patientId, params) {
    return request({
      url: `/api/v1/patient/${patientId}/samples`,
      method: 'get',
      params: { patientId: patientId, ...params }
    })
  },

  // 存储信息管理
  // 添加存储信息
  addStorageInfo(data) {
    return request({
      url: '/api/v1/sample/storage',
      method: 'post',
      data
    })
  },

  // 获取存储信息
  getStorageInfo(sampleId) {
    return request({
      url: `/api/v1/sample/${sampleId}/storage`,
      method: 'get'
    })
  },

  // 基因组数据管理
  // 添加基因组数据
  addGenomicData(data) {
    return request({
      url: '/api/v1/sample/genomic',
      method: 'post',
      data
    })
  },

  // 获取样本基因组数据
  getSampleGenomicData(sampleId) {
    return request({
      url: `/api/v1/sample/${sampleId}/genomic`,
      method: 'get',
      params: { sampleId: sampleId }
    })
  },

  // 蛋白组学数据管理
  // 添加蛋白组学数据
  addProteomicsData(data) {
    return request({
      url: '/api/v1/sample/proteomics',
      method: 'post',
      data
    })
  },

  // 获取样本蛋白组学数据
  getSampleProteomicsData(sampleId) {
    return request({
      url: `/api/v1/sample/${sampleId}/proteomics`,
      method: 'get',
      params: { sampleId: sampleId }
    })
  }
} 