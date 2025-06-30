<template>
  <div class="followup-plan-detail">
    <el-page-header @back="goBack" content="随访计划详情">
      <template #extra>
        <el-button type="primary" @click="editPlan">编辑计划</el-button>
      </template>
    </el-page-header>

    <el-card class="plan-info" v-loading="loading">
      <template #header>
        <div class="card-header">
          <span>计划信息</span>
        </div>
      </template>
      
      <el-descriptions :column="2" border v-if="plan">
        <el-descriptions-item label="计划ID">{{ plan.id }}</el-descriptions-item>
        <el-descriptions-item label="患者ID">{{ plan.patientId }}</el-descriptions-item>
        <el-descriptions-item label="计划名称">{{ plan.planName }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(plan.status)">{{ plan.status }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="开始日期">{{ formatDate(plan.startDate) }}</el-descriptions-item>
        <el-descriptions-item label="结束日期">{{ formatDate(plan.endDate) }}</el-descriptions-item>
        <el-descriptions-item label="随访频率">{{ plan.frequency }}</el-descriptions-item>
        <el-descriptions-item label="下次随访">{{ formatDate(plan.nextFollowupDate) }}</el-descriptions-item>
        <el-descriptions-item label="描述" :span="2">{{ plan.description || '-' }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ formatDate(plan.createdAt) }}</el-descriptions-item>
        <el-descriptions-item label="更新时间">{{ formatDate(plan.updatedAt) }}</el-descriptions-item>
      </el-descriptions>
    </el-card>

    <el-card class="records-section">
      <template #header>
        <div class="card-header">
          <span>随访记录</span>
          <el-button type="primary" size="small" @click="addRecord">添加记录</el-button>
        </div>
      </template>

      <el-table :data="records" v-loading="recordsLoading">
        <el-table-column prop="id" label="记录ID" width="100" />
        <el-table-column prop="followupDate" label="随访日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.followupDate) }}
          </template>
        </el-table-column>
        <el-table-column prop="followupType" label="随访类型" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getRecordStatusType(scope.row.status)">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="notes" label="备注" />
        <el-table-column label="操作" width="180">
          <template #default="scope">
            <el-button size="small" @click="viewRecord(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editRecord(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { followupApi } from '@/api/followup'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const recordsLoading = ref(false)
const plan = ref(null)
const records = ref([])

const goBack = () => {
  router.go(-1)
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleString('zh-CN')
}

const getStatusType = (status) => {
  const statusMap = {
    '计划中': 'info',
    '进行中': 'warning',
    '已完成': 'success',
    '已取消': 'danger'
  }
  return statusMap[status] || 'info'
}

const getRecordStatusType = (status) => {
  const statusMap = {
    '已完成': 'success',
    '已取消': 'danger',
    '待随访': 'warning'
  }
  return statusMap[status] || 'info'
}

const loadPlan = async () => {
  loading.value = true
  try {
    const response = await followupApi.getFollowupPlan(route.params.id)
    plan.value = response.data
  } catch (error) {
    ElMessage.error('加载计划详情失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const loadRecords = async () => {
  recordsLoading.value = true
  try {
    const response = await followupApi.getFollowupRecords({ planId: route.params.id })
    records.value = response.data.records || []
  } catch (error) {
    ElMessage.error('加载随访记录失败')
    console.error(error)
  } finally {
    recordsLoading.value = false
  }
}

const editPlan = () => {
  router.push(`/followup/plan/${route.params.id}/edit`)
}

const addRecord = () => {
  router.push(`/followup/plan/${route.params.id}/record/new`)
}

const viewRecord = (record) => {
  router.push(`/followup/record/${record.id}`)
}

const editRecord = (record) => {
  router.push(`/followup/record/${record.id}/edit`)
}

onMounted(() => {
  loadPlan()
  loadRecords()
})
</script>

<style scoped>
.followup-plan-detail {
  padding: 20px;
}

.plan-info {
  margin: 20px 0;
}

.records-section {
  margin: 20px 0;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style> 