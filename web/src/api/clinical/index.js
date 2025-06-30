import request from '@/utils/request'

export const clinicalApi = {
  // 患者信息管理
  // 创建患者
  createPatient(data) {
    return request({
      url: '/api/v1/patient',
      method: 'post',
      data
    })
  },

  // 获取患者信息
  getPatient(patientId) {
    return request({
      url: `/api/v1/patient/${patientId}`,
      method: 'get',
      params: { patientId: patientId }
    })
  },

  // 获取患者列表
  listPatients(params) {
    return request({
      url: '/api/v1/patients',
      method: 'get',
      params
    })
  },

  // 更新患者信息
  updatePatient(data) {
    return request({
      url: `/api/v1/patient/${data.patient.id}`,
      method: 'put',
      data
    })
  },

  // 就诊记录管理
  // 创建就诊记录
  createVisit(data) {
    return request({
      url: '/api/v1/visit',
      method: 'post',
      data
    })
  },

  // 获取就诊记录
  getVisit(visitId) {
    return request({
      url: `/api/v1/visit/${visitId}`,
      method: 'get',
      params: { visitId: visitId }
    })
  },

  // 获取患者就诊记录
  getPatientVisits(patientId, params) {
    return request({
      url: `/api/v1/patient/${patientId}/visits`,
      method: 'get',
      params: { patientId: patientId, ...params }
    })
  },

  // 获取就诊记录列表
  listVisits(params) {
    return request({
      url: '/api/v1/visits',
      method: 'get',
      params
    })
  },

  // 更新就诊记录
  updateVisit(data) {
    return request({
      url: `/api/v1/visit/${data.visit.id}`,
      method: 'put',
      data
    })
  },

  // 删除就诊记录
  deleteVisit(visitId) {
    return request({
      url: `/api/v1/visit/${visitId}`,
      method: 'delete',
      params: { visitId: visitId }
    })
  },

  // 诊断管理
  // 添加诊断
  addDiagnosis(data) {
    return request({
      url: '/api/v1/diagnosis',
      method: 'post',
      data
    })
  },

  // 获取就诊诊断
  getVisitDiagnoses(visitId) {
    return request({
      url: `/api/v1/visit/${visitId}/diagnoses`,
      method: 'get',
      params: { visitId: visitId }
    })
  },

  // 更新诊断
  updateDiagnosis(data) {
    return request({
      url: `/api/v1/diagnosis/${data.diagnosis.id}`,
      method: 'put',
      data
    })
  },

  // 删除诊断
  deleteDiagnosis(diagnosisId) {
    return request({
      url: `/api/v1/diagnosis/${diagnosisId}`,
      method: 'delete',
      params: { diagnosisId: diagnosisId }
    })
  },

  // 获取所有诊断记录
  listAllDiagnoses(params = {}) {
    return request({
      url: '/api/v1/diagnoses',
      method: 'get',
      params: {
        offset: params.offset || 0,
        limit: params.limit || 100,
        ...params
      }
    })
  },

  // 检查管理
  // 添加检查
  addExamination(data) {
    return request({
      url: '/api/v1/examination',
      method: 'post',
      data
    })
  },

  // 获取就诊检查
  getVisitExaminations(visitId) {
    return request({
      url: `/api/v1/visit/${visitId}/examinations`,
      method: 'get',
      params: { visitId: visitId }
    })
  },

  // 更新检查
  updateExamination(data) {
    return request({
      url: `/api/v1/examination/${data.examination.id}`,
      method: 'put',
      data
    })
  },

  // 删除检查
  deleteExamination(examinationId) {
    return request({
      url: `/api/v1/examination/${examinationId}`,
      method: 'delete',
      params: { examinationId: examinationId }
    })
  },

  // 获取所有检查记录
  listAllExaminations(params = {}) {
    return request({
      url: '/api/v1/examinations',
      method: 'get',
      params: {
        offset: params.offset || 0,
        limit: params.limit || 100,
        ...params
      }
    })
  }
} 