<template>
  <div class="environment-monitor">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>环境监测数据</span>
          <el-button type="primary" @click="showCreateDialog">添加监测数据</el-button>
        </div>
      </template>
      <!-- 搜索区域 -->
      <div class="search-area">
        <el-form :model="searchForm" inline>
          <el-form-item label="位置编码">
            <el-input v-model="searchForm.locationCode" placeholder="请输入监测点位置编码" />
          </el-form-item>
          <el-form-item label="开始时间">
            <el-date-picker
              v-model="searchForm.startTime"
              type="datetime"
              placeholder="选择开始时间"
              format="YYYY-MM-DD HH:mm:ss"
              value-format="YYYY-MM-DD HH:mm:ss"
            />
          </el-form-item>
          <el-form-item label="结束时间">
            <el-date-picker
              v-model="searchForm.endTime"
              type="datetime"
              placeholder="选择结束时间"
              format="YYYY-MM-DD HH:mm:ss"
              value-format="YYYY-MM-DD HH:mm:ss"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="loadData">搜索</el-button>
            <el-button @click="resetSearch">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="监测ID" width="100" />
        <el-table-column prop="locationCode" label="位置编码" width="120" />
        <el-table-column prop="pm25" label="PM2.5" width="100" />
        <el-table-column prop="temperature" label="温度(°C)" width="100" />
        <el-table-column prop="humidity" label="湿度(%)" width="100" />
        <el-table-column prop="monitorTime" label="监测时间" width="150">
          <template #default="scope">
            {{ formatDate(scope.row.monitorTime) }}
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { epidemiologyApi } from '@/api/epidemiology'

const loading = ref(false)
const tableData = ref([])

const searchForm = reactive({
  locationCode: '',
  startTime: '',
  endTime: ''
})

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
        locationCode: 'LOC001',
        pm25: 35.2,
        temperature: 15.5,
        humidity: 65.0,
        monitorTime: '2024-06-01 08:00:00'
      },
      {
        id: 2,
        locationCode: 'LOC002',
        pm25: 42.8,
        temperature: 18.2,
        humidity: 58.0,
        monitorTime: '2024-06-01 12:00:00'
      },
      {
        id: 3,
        locationCode: 'LOC003',
        pm25: 52.1,
        temperature: 16.8,
        humidity: 62.0,
        monitorTime: '2024-06-01 16:00:00'
      }
    ];
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.locationCode = ''
  searchForm.startTime = ''
  searchForm.endTime = ''
  loadData()
}

const showCreateDialog = () => {
  ElMessage.info('添加功能开发中...')
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.environment-monitor {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.search-area {
  background: #f5f5f5;
  padding: 20px;
  border-radius: 4px;
  margin-bottom: 20px;
}
</style> 