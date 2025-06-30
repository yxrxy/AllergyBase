<template>
  <div class="lifestyle-survey">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>生活方式调查</span>
          <el-button type="primary" @click="showCreateDialog">新建调查</el-button>
        </div>
      </template>
      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="调查ID" width="100" />
        <el-table-column prop="patientId" label="患者ID" width="100" />
        <el-table-column prop="dietPattern" label="饮食模式" width="120" />
        <el-table-column prop="exerciseFrequency" label="运动频率" width="120" />
        <el-table-column prop="sleepQuality" label="睡眠质量" width="100" />
        <el-table-column prop="stressLevel" label="压力水平" width="100" />
        <el-table-column prop="createdAt" label="创建时间" width="150">
          <template #default="scope">
            {{ formatDate(scope.row.createdAt) }}
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { epidemiologyApi } from '@/api/epidemiology'

const loading = ref(false)
const tableData = ref([])

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleString('zh-CN')
}

const loadData = async () => {
  loading.value = true
  try {
    tableData.value = [
      {
        id: 1,
        patientId: 'P001',
        dietPattern: '均衡饮食',
        exerciseFrequency: '每周3-4次',
        sleepQuality: '良好',
        stressLevel: '中等',
        createdAt: '2024-06-01 10:00:00'
      },
      {
        id: 2,
        patientId: 'P002',
        dietPattern: '高蛋白饮食',
        exerciseFrequency: '每周1-2次',
        sleepQuality: '一般',
        stressLevel: '高',
        createdAt: '2024-06-02 11:00:00'
      },
      {
        id: 3,
        patientId: 'P003',
        dietPattern: '低碳水饮食',
        exerciseFrequency: '每周5-6次',
        sleepQuality: '良好',
        stressLevel: '低',
        createdAt: '2024-06-03 09:30:00'
      }
    ];
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
.lifestyle-survey {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style> 