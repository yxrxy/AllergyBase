<template>
  <div class="sample-list">
    <div class="header">
      <h2>样本管理</h2>
      <el-button type="primary" @click="showCreateDialog = true">
        <el-icon><Plus /></el-icon>
        新增样本
      </el-button>
    </div>

    <!-- 搜索区域 -->
    <div class="search-area">
      <el-form :model="searchForm" inline>
        <el-form-item label="样本编号">
          <el-input v-model="searchForm.sampleNo" placeholder="请输入样本编号" />
        </el-form-item>
        <el-form-item label="患者ID">
          <el-input v-model="searchForm.patientId" placeholder="请输入患者ID" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadSamples">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 样本列表 -->
    <el-table :data="sampleList" v-loading="loading">
      <el-table-column prop="sampleNo" label="样本编号" />
      <el-table-column prop="patientId" label="患者ID" />
      <el-table-column prop="sampleType" label="样本类型" />
      <el-table-column prop="collectionTime" label="采集时间" />
      <el-table-column prop="collectionMethod" label="采集方法" />
      <el-table-column prop="processor" label="处理人" />
      <el-table-column label="操作" width="300">
        <template #default="{ row }">
          <el-button-group>
            <el-button size="small" type="primary" @click="viewSample(row)">查看</el-button>
            <el-button size="small" type="info" @click="addStorage(row)">存储信息</el-button>
            <el-button size="small" type="success" @click="addGenomicData(row)">基因组数据</el-button>
            <el-button size="small" type="warning" @click="addProteomicsData(row)">蛋白组学</el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>

    <!-- 创建样本对话框 -->
    <el-dialog v-model="showCreateDialog" title="新增样本" width="600px">
      <el-form :model="sampleForm" :rules="sampleRules" ref="sampleFormRef" label-width="100px">
        <el-form-item label="患者ID" prop="patientId">
          <el-input v-model="sampleForm.patientId" />
        </el-form-item>
        <el-form-item label="样本编号" prop="sampleNo">
          <el-input v-model="sampleForm.sampleNo" />
        </el-form-item>
        <el-form-item label="样本类型">
          <el-input v-model="sampleForm.sampleType" />
        </el-form-item>
        <el-form-item label="采集时间">
          <el-date-picker v-model="sampleForm.collectionTime" type="datetime" format="YYYY-MM-DD HH:mm:ss" />
        </el-form-item>
        <el-form-item label="采集方法">
          <el-input v-model="sampleForm.collectionMethod" />
        </el-form-item>
        <el-form-item label="处理人">
          <el-input v-model="sampleForm.processor" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showCreateDialog = false">取消</el-button>
        <el-button type="primary" @click="createSample">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { biobankApi } from '@/api/biobank'

// 数据
const sampleList = ref([])
const loading = ref(false)
const showCreateDialog = ref(false)

// 搜索表单
const searchForm = reactive({
  sampleNo: '',
  patientId: ''
})

// 样本表单
const sampleForm = reactive({
  patientId: '',
  sampleNo: '',
  sampleType: '',
  collectionTime: '',
  collectionMethod: '',
  processor: ''
})

const sampleRules = {
  patientId: [{ required: true, message: '请输入患者ID', trigger: 'blur' }],
  sampleNo: [{ required: true, message: '请输入样本编号', trigger: 'blur' }]
}

const sampleFormRef = ref()

// 方法
const loadSamples = async () => {
  loading.value = true
  try {
    let response
    
    // 如果指定了患者ID且不为空，则获取特定患者的样本
    if (searchForm.patientId && searchForm.patientId.trim() !== '') {
      const patientId = parseInt(searchForm.patientId)
      if (!isNaN(patientId)) {
        response = await biobankApi.getPatientSamples(patientId, {
          sampleNo: searchForm.sampleNo
        })
      } else {
        ElMessage.error('请输入有效的患者ID')
        loading.value = false
        return
      }
    } else {
      // 否则获取所有样本列表
      const params = {}
      
      // 只有当搜索条件不为空时才添加到参数中
      if (searchForm.sampleNo && searchForm.sampleNo.trim() !== '') {
        params.sampleNo = searchForm.sampleNo
      }
      if (searchForm.patientId && searchForm.patientId.trim() !== '') {
        const patientId = parseInt(searchForm.patientId)
        if (!isNaN(patientId)) {
          params.patientId = patientId
        }
      }
      
      response = await biobankApi.listSamples(params)
    }
    
    sampleList.value = response.data.samples || []
  } catch (error) {
    ElMessage.error('加载样本列表失败')
    console.error('加载样本列表失败:', error)
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.sampleNo = ''
  searchForm.patientId = ''
  loadSamples()
}

const createSample = async () => {
  if (!sampleFormRef.value) return
  
  try {
    await sampleFormRef.value.validate()
    
    // 转换数据类型和处理时间格式
    const sampleData = {
      ...sampleForm,
      patientId: parseInt(sampleForm.patientId)
    }
    
    // 处理时间格式，如果为空则移除该字段
    if (!sampleData.collectionTime) {
      delete sampleData.collectionTime
    } else {
      // 确保时间格式正确
      sampleData.collectionTime = new Date(sampleData.collectionTime).toISOString()
    }
    
    await biobankApi.createSample(sampleData)
    ElMessage.success('创建样本成功')
    showCreateDialog.value = false
    loadSamples()
  } catch (error) {
    if (error !== false) {
      ElMessage.error('创建样本失败')
    }
  }
}

const viewSample = (sample) => {
  console.log('查看样本:', sample)
}

const addStorage = (sample) => {
  console.log('添加存储信息:', sample)
}

const addGenomicData = (sample) => {
  console.log('添加基因组数据:', sample)
}

const addProteomicsData = (sample) => {
  console.log('添加蛋白组学数据:', sample)
}

onMounted(() => {
  loadSamples()
})
</script>

<style scoped>
.sample-list {
  padding: 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.search-area {
  background: #f5f5f5;
  padding: 20px;
  border-radius: 4px;
  margin-bottom: 20px;
}
</style> 