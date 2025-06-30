<template>
  <div class="proteomics-data">
    <div class="header">
      <h2>蛋白组学数据管理</h2>
      <el-button type="primary" @click="showCreateDialog = true">
        <el-icon><Plus /></el-icon>
        新增蛋白组学数据
      </el-button>
    </div>

    <!-- 搜索区域 -->
    <div class="search-area">
      <el-form :model="searchForm" inline>
        <el-form-item label="样本ID">
          <el-input v-model="searchForm.sampleId" placeholder="请输入样本ID" />
        </el-form-item>
        <el-form-item label="分析平台">
          <el-input v-model="searchForm.analysisPlatform" placeholder="请输入分析平台" />
        </el-form-item>
        <el-form-item label="蛋白标志物">
          <el-input v-model="searchForm.proteinMarker" placeholder="请输入蛋白标志物" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadProteomicsData">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 蛋白组学数据列表 -->
    <el-table :data="proteomicsDataList" v-loading="loading">
      <el-table-column prop="id" label="数据ID" />
      <el-table-column prop="sampleId" label="样本ID" />
      <el-table-column prop="analysisPlatform" label="分析平台" />
      <el-table-column prop="proteinMarker" label="蛋白标志物" />
      <el-table-column prop="concentration" label="浓度" />
      <el-table-column label="操作" width="200">
        <template #default="{ row }">
          <el-button size="small" @click="viewProteomicsData(row)">查看</el-button>
          <el-button size="small" type="primary" @click="editProteomicsData(row)">编辑</el-button>
          <el-button size="small" type="danger" @click="deleteProteomicsData(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 创建/编辑蛋白组学数据对话框 -->
    <el-dialog v-model="showCreateDialog" :title="isEdit ? '编辑蛋白组学数据' : '新增蛋白组学数据'" width="600px">
      <el-form :model="proteomicsDataForm" :rules="proteomicsDataRules" ref="proteomicsDataFormRef" label-width="100px">
        <el-form-item label="样本ID" prop="sampleId">
          <el-input v-model="proteomicsDataForm.sampleId" type="number" />
        </el-form-item>
        <el-form-item label="分析平台" prop="analysisPlatform">
          <el-input v-model="proteomicsDataForm.analysisPlatform" />
        </el-form-item>
        <el-form-item label="蛋白标志物">
          <el-input v-model="proteomicsDataForm.proteinMarker" />
        </el-form-item>
        <el-form-item label="浓度">
          <el-input v-model="proteomicsDataForm.concentration" type="number" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showCreateDialog = false">取消</el-button>
        <el-button type="primary" @click="saveProteomicsData">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { biobankApi } from '@/api/biobank'

// 数据
const proteomicsDataList = ref([
  {
    id: 1,
    sampleId: 'S001',
    analysisPlatform: '质谱仪',
    proteinMarker: 'IgE',
    concentration: 120.5
  },
  {
    id: 2,
    sampleId: 'S002',
    analysisPlatform: '质谱仪',
    proteinMarker: 'IL-4',
    concentration: 25.3
  },
  {
    id: 3,
    sampleId: 'S003',
    analysisPlatform: '质谱仪',
    proteinMarker: 'IL-13',
    concentration: 18.7
  }
])
const loading = ref(false)
const showCreateDialog = ref(false)
const isEdit = ref(false)

// 搜索表单
const searchForm = reactive({
  sampleId: '',
  analysisPlatform: '',
  proteinMarker: ''
})

// 蛋白组学数据表单
const proteomicsDataForm = reactive({
  id: 0,
  sampleId: '',
  analysisPlatform: '',
  proteinMarker: '',
  concentration: ''
})

const proteomicsDataRules = {
  sampleId: [{ required: true, message: '请输入样本ID', trigger: 'blur' }],
  analysisPlatform: [{ required: true, message: '请输入分析平台', trigger: 'blur' }]
}

const proteomicsDataFormRef = ref()

// 方法
const loadProteomicsData = async () => {
  loading.value = true
  try {
    proteomicsDataList.value = [
      {
        id: 1,
        sampleId: 'S001',
        analysisPlatform: '质谱仪',
        proteinMarker: 'IgE',
        concentration: 120.5
      },
      {
        id: 2,
        sampleId: 'S002',
        analysisPlatform: '质谱仪',
        proteinMarker: 'IL-4',
        concentration: 25.3
      },
      {
        id: 3,
        sampleId: 'S003',
        analysisPlatform: '质谱仪',
        proteinMarker: 'IL-13',
        concentration: 18.7
      }
    ];
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.sampleId = ''
  searchForm.analysisPlatform = ''
  searchForm.proteinMarker = ''
  loadProteomicsData()
}

const resetForm = () => {
  proteomicsDataForm.id = 0
  proteomicsDataForm.sampleId = ''
  proteomicsDataForm.analysisPlatform = ''
  proteomicsDataForm.proteinMarker = ''
  proteomicsDataForm.concentration = ''
}

const createProteomicsData = async () => {
  if (!proteomicsDataFormRef.value) return
  
  try {
    await proteomicsDataFormRef.value.validate()
    
    const proteomicsData = {
      ...proteomicsDataForm,
      sampleId: parseInt(proteomicsDataForm.sampleId),
      concentration: parseFloat(proteomicsDataForm.concentration) || 0
    }
    
    await biobankApi.addProteomicsData(proteomicsData)
    ElMessage.success('创建蛋白组学数据成功')
    showCreateDialog.value = false
    loadProteomicsData()
  } catch (error) {
    if (error !== false) {
      ElMessage.error('创建蛋白组学数据失败')
    }
  }
}

const editProteomicsData = (proteomicsData) => {
  isEdit.value = true
  proteomicsDataForm.id = proteomicsData.id
  proteomicsDataForm.sampleId = proteomicsData.sampleId
  proteomicsDataForm.analysisPlatform = proteomicsData.analysisPlatform || ''
  proteomicsDataForm.proteinMarker = proteomicsData.proteinMarker || ''
  proteomicsDataForm.concentration = proteomicsData.concentration || ''
  showCreateDialog.value = true
}

const saveProteomicsData = async () => {
  if (isEdit.value) {
    // 编辑功能暂时显示提示
    ElMessage.info('编辑功能暂未实现')
  } else {
    await createProteomicsData()
  }
}

const viewProteomicsData = (proteomicsData) => {
  console.log('查看蛋白组学数据:', proteomicsData)
  ElMessage.info('查看功能暂未实现')
}

const deleteProteomicsData = async (proteomicsData) => {
  try {
    await ElMessageBox.confirm('确定要删除这条蛋白组学数据吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    // 删除功能暂时显示提示
    ElMessage.info('删除功能暂未实现')
  } catch (error) {
    // 用户取消删除
  }
}

onMounted(() => {
  proteomicsDataList.value = [
    {
      id: 1,
      sampleId: 'S001',
      analysisPlatform: '质谱仪',
      proteinMarker: 'IgE',
      concentration: 120.5
    },
    {
      id: 2,
      sampleId: 'S002',
      analysisPlatform: '质谱仪',
      proteinMarker: 'IL-4',
      concentration: 25.3
    },
    {
      id: 3,
      sampleId: 'S003',
      analysisPlatform: '质谱仪',
      proteinMarker: 'IL-13',
      concentration: 18.7
    }
  ];
  loadProteomicsData()
})
</script>

<style scoped>
.proteomics-data {
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