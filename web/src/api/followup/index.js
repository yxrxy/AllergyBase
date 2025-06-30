import request from '@/utils/request'

export const followupApi = {
  // 随访计划管理
  // 创建随访计划
  createFollowupPlan(data) {
    return request({
      url: '/api/v1/followup/plan',
      method: 'post',
      data
    })
  },

  // 获取随访计划
  getFollowupPlan(planId) {
    return request({
      url: `/api/v1/followup/plan/${planId}`,
      method: 'get',
      params: { planId: planId }
    })
  },

  // 获取所有随访计划列表
  getFollowupPlans(params) {
    return request({
      url: '/api/v1/followup/plans',
      method: 'get',
      params
    })
  },

  // 获取患者随访计划
  getPatientFollowupPlans(patientId, params) {
    return request({
      url: `/api/v1/patient/${patientId}/followup/plans`,
      method: 'get',
      params: { patientId: patientId, ...params }
    })
  },

  // 随访记录管理
  // 创建随访记录
  createFollowupRecord(data) {
    return request({
      url: '/api/v1/followup/record',
      method: 'post',
      data
    })
  },

  // 获取随访记录
  getFollowupRecord(recordId) {
    return request({
      url: `/api/v1/followup/record/${recordId}`,
      method: 'get',
      params: { recordId: recordId }
    })
  },

  // 获取计划随访记录
  getPlanFollowupRecords(planId, params) {
    return request({
      url: `/api/v1/followup/plan/${planId}/records`,
      method: 'get',
      params: { planId: planId, ...params }
    })
  },

  // 用药监测管理
  // 添加用药监测
  addMedicationMonitor(data) {
    return request({
      url: '/api/v1/followup/medication',
      method: 'post',
      data
    })
  },

  // 获取随访用药监测
  getFollowupMedications(followupId, params) {
    return request({
      url: `/api/v1/followup/${followupId}/medications`,
      method: 'get',
      params: { followupId: followupId, ...params }
    })
  }
} 