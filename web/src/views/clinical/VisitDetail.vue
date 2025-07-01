<template>
  <div class="visit-detail">
    <el-page-header @back="goBack" content="就诊详情">
      <template #extra>
        <el-button type="primary" @click="editVisit">编辑就诊记录</el-button>
      </template>
    </el-page-header>

    <el-card class="visit-info" v-loading="loading">
      <template #header>
        <div class="card-header">
          <span>就诊信息</span>
        </div>
      </template>
      
      <el-descriptions :column="2" border v-if="visit">
        <el-descriptions-item label="就诊ID">{{ visit.id }}</el-descriptions-item>
        <el-descriptions-item label="患者ID">{{ visit.patientId }}</el-descriptions-item>
        <el-descriptions-item label="就诊编号">{{ visit.visitNo }}</el-descriptions-item>
        <el-descriptions-item label="就诊时间">{{ formatDate(visit.visitTime) }}</el-descriptions-item>
        <el-descriptions-item label="就诊类型">{{ getVisitTypeText(visit.visitType) }}</el-descriptions-item>
        <el-descriptions-item label="科室">{{ visit.department }}</el-descriptions-item>
        <el-descriptions-item label="医生ID">{{ visit.doctorId || '-' }}</el-descriptions-item>
        <el-descriptions-item label="主诉" :span="2">{{ visit.chiefComplaint }}</el-descriptions-item>
      </el-descriptions>
    </el-card>

    <el-card class="diagnoses-section">
      <template #header>
        <div class="card-header">
          <span>诊断记录</span>
          <el-button type="primary" size="small" @click="addDiagnosis">添加诊断</el-button>
        </div>
      </template>

      <el-table :data="diagnoses" v-loading="diagnosesLoading">
        <el-table-column prop="id" label="诊断ID" width="100" />
        <el-table-column prop="diagnosisType" label="诊断类型" width="120">
          <template #default="scope">
            <el-tag>{{ scope.row.diagnosisType }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="icdCode" label="ICD编码" width="150" />
        <el-table-column label="操作" width="180">
          <template #default="scope">
            <el-button size="small" @click="viewDiagnosis(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editDiagnosis(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteDiagnosis(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-card class="examinations-section">
      <template #header>
        <div class="card-header">
          <span>检查记录</span>
          <el-button type="primary" size="small" @click="addExamination">添加检查</el-button>
        </div>
      </template>

      <el-table :data="examinations" v-loading="examinationsLoading">
        <el-table-column prop="id" label="检查ID" width="100" />
        <el-table-column prop="examType" label="检查类型" width="150" />
        <el-table-column prop="examTime" label="检查时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.examTime) }}
          </template>
        </el-table-column>
        <el-table-column prop="examResult" label="检查结果" min-width="200" show-overflow-tooltip />
        <el-table-column prop="resultUnit" label="单位" width="100" />
        <el-table-column prop="referenceRange" label="参考范围" width="150" />
        <el-table-column label="操作" width="180">
          <template #default="scope">
            <el-button size="small" @click="viewExamination(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editExamination(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteExamination(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { clinicalApi } from '@/api/clinical'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const diagnosesLoading = ref(false)
const examinationsLoading = ref(false)
const visit = ref(null)
const diagnoses = ref([])
const examinations = ref([])

const goBack = () => {
  router.go(-1)
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleString('zh-CN')
}

const getVisitTypeText = (visitType) => {
  const typeMap = {
    1: '门诊',
    2: '急诊',
    3: '住院'
  }
  return typeMap[visitType] || '未知'
}

const loadData = async () => {
  const visitId = route.params.id
  
  // 加载就诊信息
  loading.value = true
  try {
    const response = await clinicalApi.getVisit(visitId)
    visit.value = response.data.visit
  } catch (error) {
    ElMessage.error('加载就诊信息失败')
    console.error(error)
  } finally {
    loading.value = false
  }

  // 加载诊断记录
  diagnosesLoading.value = true
  try {
    const response = await clinicalApi.getVisitDiagnoses(visitId)
    diagnoses.value = response.data.diagnoses || []
  } catch (error) {
    ElMessage.error('加载诊断记录失败')
    console.error(error)
  } finally {
    diagnosesLoading.value = false
  }

  // 加载检查记录
  examinationsLoading.value = true
  try {
    const response = await clinicalApi.getVisitExaminations(visitId)
    examinations.value = response.data.examinations || []
  } catch (error) {
    ElMessage.error('加载检查记录失败')
    console.error(error)
  } finally {
    examinationsLoading.value = false
  }
}

const editVisit = () => {
  router.push(`/clinical/visit/${route.params.id}/edit`)
}

const addDiagnosis = () => {
  router.push('/clinical/diagnosis')
}

const viewDiagnosis = (diagnosis) => {
  ElMessage.info('查看诊断功能待实现')
}

const editDiagnosis = (diagnosis) => {
  ElMessage.info('编辑诊断功能待实现')
}

const deleteDiagnosis = async (diagnosis) => {
  try {
    await ElMessageBox.confirm('确定要删除这个诊断记录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await clinicalApi.deleteDiagnosis(diagnosis.id)
    ElMessage.success('删除成功')
    // 重新加载诊断记录
    diagnosesLoading.value = true
    try {
      const response = await clinicalApi.getVisitDiagnoses(route.params.id)
      diagnoses.value = response.data.diagnoses || []
    } finally {
      diagnosesLoading.value = false
    }
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
      console.error(error)
    }
  }
}

const addExamination = () => {
  router.push('/clinical/examination')
}

const viewExamination = (examination) => {
  ElMessage.info('查看检查功能待实现')
}

const editExamination = (examination) => {
  ElMessage.info('编辑检查功能待实现')
}

const deleteExamination = async (examination) => {
  try {
    await ElMessageBox.confirm('确定要删除这个检查记录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await clinicalApi.deleteExamination(examination.id)
    ElMessage.success('删除成功')
    // 重新加载检查记录
    examinationsLoading.value = true
    try {
      const response = await clinicalApi.getVisitExaminations(route.params.id)
      examinations.value = response.data.examinations || []
    } finally {
      examinationsLoading.value = false
    }
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
      console.error(error)
    }
  }
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.visit-detail {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.visit-info {
  margin-bottom: 20px;
}

.diagnoses-section {
  margin-bottom: 20px;
}

.examinations-section {
  margin-bottom: 20px;
}
</style> 