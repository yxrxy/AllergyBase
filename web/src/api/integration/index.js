import { integrationRequest } from '@/utils/request'

// 数据集成API
export const integrationApi = {
  // 获取仪表板统计数据
  getDashboardStats() {
    return integrationRequest({
      url: '/api/v1/integration/dashboard/stats',
      method: 'get'
    })
  },

  // 获取最近活动
  getRecentActivities(limit = 10) {
    return integrationRequest({
      url: '/api/v1/integration/activities/recent',
      method: 'get',
      params: { limit }
    })
  },

  // 获取患者总数统计
  getPatientsCount() {
    return integrationRequest({
      url: '/api/v1/integration/stats/patients/count',
      method: 'get'
    })
  },

  // 获取数据记录总数统计
  getRecordsCount() {
    return integrationRequest({
      url: '/api/v1/integration/stats/records/count',
      method: 'get'
    })
  },

  // 获取数据质量评分
  getQualityScore() {
    return integrationRequest({
      url: '/api/v1/integration/stats/quality/score',
      method: 'get'
    })
  },

  // 获取活跃用户数统计
  getActiveUsersCount() {
    return integrationRequest({
      url: '/api/v1/integration/stats/users/active/count',
      method: 'get'
    })
  },

  // 获取患者综合视图
  getPatientIntegratedView(patientNo) {
    return integrationRequest({
      url: `/api/v1/integration/patient/integrated-view`,
      method: 'get',
      params: { patient_no: patientNo }
    })
  },

  // 搜索患者
  searchPatients(keyword) {
    return integrationRequest({
      url: '/api/v1/integration/search/patients',
      method: 'post',
      params: { keyword }
    })
  },

  // 检查患者数据质量
  checkPatientDataQuality(patientId) {
    return integrationRequest({
      url: `/api/v1/integration/patient/${patientId}/data-quality`,
      method: 'get',
      params: { patientId: patientId }
    })
  },

  // 批量检查数据质量
  batchCheckDataQuality(patientIds) {
    return integrationRequest({
      url: '/api/v1/integration/batch/data-quality',
      method: 'post',
      data: { patient_ids: patientIds }
    })
  },

  // 执行关联分析
  performCorrelationAnalysis(analysisType, parameters) {
    return integrationRequest({
      url: '/api/v1/integration/analysis/correlation/perform',
      method: 'post',
      data: {
        analysis_type: analysisType,
        parameters: parameters
      }
    })
  },

  // 获取关联分析结果
  getCorrelationAnalysis(analysisId) {
    return integrationRequest({
      url: `/api/v1/integration/analysis/correlation/${analysisId}`,
      method: 'get',
      params: { analysisId: analysisId }
    })
  },

  // 获取队列洞察
  getCohortInsights(query) {
    return integrationRequest({
      url: '/api/v1/integration/cohort/insights',
      method: 'get',
      params: query
    })
  },

  // 获取患者统计信息
  getPatientStatistics(patientId) {
    return integrationRequest({
      url: `/api/v1/integration/patient/${patientId}/statistics`,
      method: 'get',
      params: { patientId: patientId }
    })
  },

  // 导出患者数据
  exportPatientData(patientIds, format = 'excel') {
    return integrationRequest({
      url: '/api/v1/integration/export/patient-data',
      method: 'post',
      data: {
        patient_ids: patientIds,
        format: format
      },
      responseType: 'blob'
    })
  }
} 