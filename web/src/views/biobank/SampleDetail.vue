<template>
  <div class="sample-detail">
    <el-page-header @back="goBack" content="样本详情">
      <template #extra>
        <el-button type="primary" @click="editSample">编辑样本</el-button>
      </template>
    </el-page-header>

    <el-card class="sample-info" v-loading="loading">
      <template #header>
        <div class="card-header">
          <span>样本信息</span>
        </div>
      </template>
      
      <el-descriptions :column="2" border v-if="sample">
        <el-descriptions-item label="样本ID">{{ sample.id }}</el-descriptions-item>
        <el-descriptions-item label="患者ID">{{ sample.patientId }}</el-descriptions-item>
        <el-descriptions-item label="样本类型">{{ sample.sampleType }}</el-descriptions-item>
        <el-descriptions-item label="采集日期">{{ formatDate(sample.collectionDate) }}</el-descriptions-item>
        <el-descriptions-item label="采集部位">{{ sample.collectionSite }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(sample.status)">{{ sample.status }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="质量评估">{{ sample.qualityAssessment || '-' }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ sample.notes || '-' }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ formatDate(sample.createdAt) }}</el-descriptions-item>
        <el-descriptions-item label="更新时间">{{ formatDate(sample.updatedAt) }}</el-descriptions-item>
      </el-descriptions>
    </el-card>

    <el-card class="storage-section">
      <template #header>
        <div class="card-header">
          <span>存储信息</span>
          <el-button type="primary" size="small" @click="addStorage">添加存储信息</el-button>
        </div>
      </template>

      <el-table :data="storageData" v-loading="storageLoading">
        <el-table-column prop="id" label="存储ID" width="100" />
        <el-table-column prop="storageLocation" label="存储位置" width="150" />
        <el-table-column prop="temperature" label="温度" width="100" />
        <el-table-column prop="humidity" label="湿度" width="100" />
        <el-table-column prop="storageDate" label="存储日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.storageDate) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStorageStatusType(scope.row.status)">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180">
          <template #default="scope">
            <el-button size="small" @click="viewStorage(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editStorage(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-card class="genomic-section">
      <template #header>
        <div class="card-header">
          <span>基因组数据</span>
          <el-button type="primary" size="small" @click="addGenomicData">添加基因组数据</el-button>
        </div>
      </template>

      <el-table :data="genomicData" v-loading="genomicLoading">
        <el-table-column prop="id" label="数据ID" width="100" />
        <el-table-column prop="dataType" label="数据类型" width="120" />
        <el-table-column prop="filePath" label="文件路径" />
        <el-table-column prop="fileSize" label="文件大小" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getDataStatusType(scope.row.status)">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" @click="downloadData(scope.row)">下载</el-button>
            <el-button size="small" type="primary" @click="viewGenomicData(scope.row)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-card class="proteomics-section">
      <template #header>
        <div class="card-header">
          <span>蛋白质组数据</span>
          <el-button type="primary" size="small" @click="addProteomicsData">添加蛋白质组数据</el-button>
        </div>
      </template>

      <el-table :data="proteomicsData" v-loading="proteomicsLoading">
        <el-table-column prop="id" label="数据ID" width="100" />
        <el-table-column prop="analysisType" label="分析类型" width="120" />
        <el-table-column prop="proteinCount" label="蛋白质数量" width="120" />
        <el-table-column prop="filePath" label="文件路径" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getDataStatusType(scope.row.status)">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" @click="downloadData(scope.row)">下载</el-button>
            <el-button size="small" type="primary" @click="viewProteomicsData(scope.row)">查看</el-button>
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
import { biobankApi } from '@/api/biobank'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const storageLoading = ref(false)
const genomicLoading = ref(false)
const proteomicsLoading = ref(false)

const sample = ref(null)
const storageData = ref([])
const genomicData = ref([])
const proteomicsData = ref([])

const goBack = () => {
  router.go(-1)
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleString('zh-CN')
}

const getStatusType = (status) => {
  const statusMap = {
    '正常': 'success',
    '异常': 'danger',
    '处理中': 'warning'
  }
  return statusMap[status] || 'info'
}

const getStorageStatusType = (status) => {
  const statusMap = {
    '存储中': 'success',
    '已取出': 'warning',
    '已销毁': 'danger'
  }
  return statusMap[status] || 'info'
}

const getDataStatusType = (status) => {
  const statusMap = {
    '处理中': 'warning',
    '已完成': 'success',
    '失败': 'danger'
  }
  return statusMap[status] || 'info'
}

const loadSample = async () => {
  loading.value = true
  try {
    const response = await biobankApi.getSample(route.params.id)
    sample.value = response.data
  } catch (error) {
    ElMessage.error('加载样本详情失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const loadStorageData = async () => {
  storageLoading.value = true
  try {
    const response = await biobankApi.getStorageInfo({ sampleId: route.params.id })
    storageData.value = response.data.storageList || []
  } catch (error) {
    ElMessage.error('加载存储信息失败')
    console.error(error)
  } finally {
    storageLoading.value = false
  }
}

const loadGenomicData = async () => {
  genomicLoading.value = true
  try {
    const response = await biobankApi.getGenomicData({ sampleId: route.params.id })
    genomicData.value = response.data.genomicList || []
  } catch (error) {
    ElMessage.error('加载基因组数据失败')
    console.error(error)
  } finally {
    genomicLoading.value = false
  }
}

const loadProteomicsData = async () => {
  proteomicsLoading.value = true
  try {
    const response = await biobankApi.getProteomicsData({ sampleId: route.params.id })
    proteomicsData.value = response.data.proteomicsList || []
  } catch (error) {
    ElMessage.error('加载蛋白质组数据失败')
    console.error(error)
  } finally {
    proteomicsLoading.value = false
  }
}

const editSample = () => {
  router.push(`/biobank/sample/${route.params.id}/edit`)
}

const addStorage = () => {
  ElMessage.info('添加存储信息功能开发中...')
}

const viewStorage = (storage) => {
  ElMessage.info('查看存储信息功能开发中...')
}

const editStorage = (storage) => {
  ElMessage.info('编辑存储信息功能开发中...')
}

const addGenomicData = () => {
  ElMessage.info('添加基因组数据功能开发中...')
}

const viewGenomicData = (data) => {
  ElMessage.info('查看基因组数据功能开发中...')
}

const addProteomicsData = () => {
  ElMessage.info('添加蛋白质组数据功能开发中...')
}

const viewProteomicsData = (data) => {
  ElMessage.info('查看蛋白质组数据功能开发中...')
}

const downloadData = (data) => {
  ElMessage.info('下载功能开发中...')
}

onMounted(() => {
  loadSample()
  loadStorageData()
  loadGenomicData()
  loadProteomicsData()
})
</script>

<style scoped>
.sample-detail {
  padding: 20px;
}

.sample-info,
.storage-section,
.genomic-section,
.proteomics-section {
  margin: 20px 0;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style> 