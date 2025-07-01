<template>
  <div class="followup-record-list">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>随访记录列表</span>
          <el-button type="primary" @click="showCreateDialog">添加记录</el-button>
        </div>
      </template>
      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="记录ID" width="100" />
        <el-table-column prop="planId" label="计划ID" width="100" />
        <el-table-column prop="followupDate" label="随访日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.followupDate) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100" />
        <el-table-column prop="notes" label="备注" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { followupApi } from '@/api/followup'

const loading = ref(false)
const tableData = ref([
  {
    id: 1,
    planId: 1,
    followupDate: '2024-03-18',
    status: '正常',
    notes: '恢复良好，无明显不适'
  },
  {
    id: 2,
    planId: 2,
    followupDate: '2024-04-01',
    status: '随访中',
    notes: '偶有咳嗽'
  }
])

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('zh-CN')
}

const loadData = async () => {
  loading.value = true
  try {
    const response = await followupApi.getFollowupRecords({})
    tableData.value = response.data.records || []
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const showCreateDialog = () => {
  ElMessage.info('创建功能开发中...')
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.followup-record-list {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style> 