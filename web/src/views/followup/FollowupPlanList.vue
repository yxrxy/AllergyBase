<template>
  <div class="followup-plan-list">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>随访计划列表</span>
          <el-button type="primary" @click="showCreateDialog">创建计划</el-button>
        </div>
      </template>
      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="计划ID" width="100" />
        <el-table-column prop="patientId" label="患者ID" width="100" />
        <el-table-column prop="planName" label="计划名称" width="150" />
        <el-table-column prop="startDate" label="开始日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.startDate) }}
          </template>
        </el-table-column>
        <el-table-column prop="frequency" label="频率" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="240">
          <template #default="scope">
            <el-button size="small" @click="viewPlan(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editPlan(scope.row)">编辑</el-button>
            <el-button size="small" type="warning" @click="addRecord(scope.row)">添加记录</el-button>
          </template>
        </el-table-column>
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
    patientId: 'P001',
    planName: '术后随访',
    startDate: '2024-03-11',
    frequency: '每周',
    status: '进行中'
  },
  {
    id: 2,
    patientId: 'P002',
    planName: '慢病管理',
    startDate: '2024-01-02',
    frequency: '每月',
    status: '已完成'
  },
  {
    id: 3,
    patientId: 'P003',
    planName: '高血压随访',
    startDate: '2024-02-10',
    frequency: '每两周',
    status: '进行中'
  },
  {
    id: 4,
    patientId: 'P004',
    planName: '糖尿病随访',
    startDate: '2024-03-05',
    frequency: '每月',
    status: '已完成'
  }
])

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('zh-CN')
}

const getStatusType = (status) => {
  const statusMap = {
    '计划中': 'info',
    '进行中': 'warning',
    '已完成': 'success'
  }
  return statusMap[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    tableData.value = [
      {
        id: 1,
        patientId: 'P001',
        planName: '术后随访',
        startDate: '2024-03-11',
        frequency: '每周',
        status: '进行中'
      },
      {
        id: 2,
        patientId: 'P002',
        planName: '慢病管理',
        startDate: '2024-01-02',
        frequency: '每月',
        status: '已完成'
      },
      {
        id: 3,
        patientId: 'P003',
        planName: '高血压随访',
        startDate: '2024-02-10',
        frequency: '每两周',
        status: '进行中'
      },
      {
        id: 4,
        patientId: 'P004',
        planName: '糖尿病随访',
        startDate: '2024-03-05',
        frequency: '每月',
        status: '已完成'
      }
    ];
  } finally {
    loading.value = false
  }
}

const showCreateDialog = () => {
  ElMessage.info('创建功能开发中...')
}

function viewPlan(row) {
  ElMessage.info(`查看计划：${row.planName}`)
}

function editPlan(row) {
  ElMessage.info(`编辑计划：${row.planName}`)
}

function addRecord(row) {
  ElMessage.info(`为计划“${row.planName}”添加记录`)
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.followup-plan-list {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style> 