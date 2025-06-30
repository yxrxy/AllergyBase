<template>
  <div class="patient-detail">
    <div class="page-header">
      <el-button @click="$router.back()" icon="ArrowLeft">返回</el-button>
      <h1>患者详情</h1>
      <el-button type="primary" @click="editPatient">编辑</el-button>
    </div>

    <el-card v-loading="loading">
      <template #header>
        <span>基本信息</span>
      </template>
        <el-descriptions :column="2" border>
        <el-descriptions-item label="患者编号">{{ patient.patientNo }}</el-descriptions-item>
        <el-descriptions-item label="姓名">{{ patient.name }}</el-descriptions-item>
        <el-descriptions-item label="性别">{{ patient.gender === 1 ? '男' : '女' }}</el-descriptions-item>
        <el-descriptions-item label="出生日期">{{ patient.birthDate }}</el-descriptions-item>
        <el-descriptions-item label="身高(cm)">{{ patient.height || '-' }}</el-descriptions-item>
        <el-descriptions-item label="体重(kg)">{{ patient.weight || '-' }}</el-descriptions-item>
        <el-descriptions-item label="出生体重(kg)">{{ patient.birthWeight || '-' }}</el-descriptions-item>
        <el-descriptions-item label="医保类型">{{ patient.medicalInsuranceType || '-' }}</el-descriptions-item>
        <el-descriptions-item label="医保号">{{ patient.medicalInsuranceNo || '-' }}</el-descriptions-item>
      </el-descriptions>
    </el-card>

    <!-- 就诊记录 -->
    <el-card class="mt-4">
      <template #header>
        <span>就诊记录</span>
        <el-button type="primary" size="small" @click="addVisit">新增就诊</el-button>
      </template>
      
      <el-table :data="visits" v-loading="visitsLoading">
        <el-table-column prop="visit_no" label="就诊号" />
        <el-table-column prop="visit_time" label="就诊时间" />
        <el-table-column prop="department" label="科室" />
        <el-table-column prop="doctor" label="医生" />
        <el-table-column prop="chief_complaint" label="主诉" show-overflow-tooltip />
        <el-table-column label="操作" width="150">
          <template #default="{ row }">
            <el-button size="small" @click="viewVisit(row.id)">查看</el-button>
            <el-button size="small" type="primary" @click="editVisit(row.id)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { clinicalApi } from '@/api/clinical'

const route = useRoute()
const loading = ref(false)
const visitsLoading = ref(false)
const patient = ref({})
const visits = ref([])

const loadPatient = async () => {
  loading.value = true
  try {
    const response = await clinicalApi.getPatient(route.params.id)
    patient.value = response.data.patient || {}
  } catch (error) {
    console.error('加载患者信息失败:', error)
  } finally {
    loading.value = false
  }
}

const loadVisits = async () => {
  visitsLoading.value = true
  try {
    const response = await clinicalApi.getPatientVisits(route.params.id)
    visits.value = response.data.visits || []
  } catch (error) {
    console.error('加载就诊记录失败:', error)
  } finally {
    visitsLoading.value = false
  }
}

const editPatient = () => {
  // 编辑患者逻辑
  console.log('编辑患者')
}

const addVisit = () => {
  // 新增就诊记录逻辑
  console.log('新增就诊')
}

const viewVisit = (visitId) => {
  // 查看就诊记录详情
  console.log('查看就诊记录:', visitId)
}

const editVisit = (visitId) => {
  // 编辑就诊记录
  console.log('编辑就诊记录:', visitId)
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleString('zh-CN')
}

onMounted(() => {
  loadPatient()
  loadVisits()
})
</script>

<style scoped>
.patient-detail {
  padding: 20px;
}

.page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.page-header h1 {
  margin: 0;
  flex: 1;
  text-align: center;
}

.mt-4 {
  margin-top: 16px;
}
</style> 