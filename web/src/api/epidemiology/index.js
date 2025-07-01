import request from '@/utils/request'

export const epidemiologyApi = {
  // 环境暴露管理
  // 创建环境暴露记录
  createEnvironmentExposure(data) {
    return request({
      url: '/api/v1/environment/exposure',
      method: 'post',
      data
    })
  },

  // 获取环境暴露记录
  getEnvironmentExposure(patientId) {
    return request({
      url: `/api/v1/patient/${patientId}/environment/exposure`,
      method: 'get',
      params: { patientId: patientId }
    })
  },

  // 获取环境暴露记录列表
  getEnvironmentExposures(params) {
    return request({
      url: '/api/v1/environment/exposures',
      method: 'get',
      params
    })
  },

  // 环境监测数据管理
  // 添加环境监测数据
  addEnvironmentMonitor(data) {
    return request({
      url: '/api/v1/environment/monitor',
      method: 'post',
      data
    })
  },

  // 获取环境监测数据
  getEnvironmentMonitors(params) {
    return request({
      url: '/api/v1/environment/monitors',
      method: 'get',
      params
    })
  },

  // 生活方式调查管理
  // 创建生活方式调查
  createLifestyleSurvey(data) {
    return request({
      url: '/api/v1/lifestyle/survey',
      method: 'post',
      data
    })
  },

  // 获取生活方式调查
  getLifestyleSurvey(patientId) {
    return request({
      url: `/api/v1/patient/${patientId}/lifestyle/survey`,
      method: 'get',
      params: { patientId: patientId }
    })
  },

  // 获取生活方式调查列表
  getLifestyleSurveys(params) {
    return request({
      url: '/api/v1/lifestyle/surveys',
      method: 'get',
      params
    })
  }
} 