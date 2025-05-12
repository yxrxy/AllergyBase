import request from '@/utils/request'

export const allergyApi = {
  // 获取最近的过敏记录
  getRecentRecords() {
    return request({
      url: '/api/v1/allergy/records/recent',
      method: 'get'
    })
  },

  // 获取所有过敏记录
  getAllRecords(params) {
    return request({
      url: '/api/v1/allergy/records',
      method: 'get',
      params
    })
  },

  // 创建新的过敏记录
  createRecord(data) {
    return request({
      url: '/api/v1/allergy/records',
      method: 'post',
      data
    })
  },

  // 更新过敏记录
  updateRecord(id, data) {
    return request({
      url: `/api/v1/allergy/records/${id}`,
      method: 'put',
      data
    })
  },

  // 删除过敏记录
  deleteRecord(id) {
    return request({
      url: `/api/v1/allergy/records/${id}`,
      method: 'delete'
    })
  },

  // 获取过敏分析报告
  getAnalysisReport() {
    return request({
      url: '/api/v1/allergy/analysis',
      method: 'get'
    })
  },

  // 获取预防建议
  getPreventionAdvice() {
    return request({
      url: '/api/v1/allergy/prevention',
      method: 'get'
    })
  }
} 