<template>
  <div class="genomic-data">
    <div class="header">
      <h2>基因组数据管理</h2>
      <el-button type="primary" @click="showCreateDialog = true">
        <el-icon><Plus /></el-icon>
        新增基因组数据
      </el-button>
    </div>

    <!-- 搜索区域 -->
    <div class="search-area">
      <el-form :model="searchForm" inline>
        <el-form-item label="样本ID">
          <el-input v-model="searchForm.sampleId" placeholder="请输入样本ID" />
        </el-form-item>
        <el-form-item label="测序平台">
          <el-input v-model="searchForm.sequencePlatform" placeholder="请输入测序平台" />
        </el-form-item>
        <el-form-item label="基因类型">
          <el-input v-model="searchForm.geneType" placeholder="请输入基因类型" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadGenomicData">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 基因组数据列表 -->
    <el-table :data="genomicDataList" v-loading="loading">
      <el-table-column prop="id" label="数据ID" />
      <el-table-column prop="sampleId" label="样本ID" />
      <el-table-column prop="sequencePlatform" label="测序平台" />
      <el-table-column prop="sequenceType" label="测序类型" />
      <el-table-column prop="geneType" label="基因类型" />
      <el-table-column prop="resultFilePath" label="结果文件路径" />
      <el-table-column label="操作" width="200">
        <template #default="{ row }">
          <el-button size="small" @click="viewGenomicData(row)">查看</el-button>
          <el-button size="small" type="primary" @click="editGenomicData(row)">编辑</el-button>
          <el-button size="small" type="danger" @click="deleteGenomicData(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 创建/编辑基因组数据对话框 -->
    <el-dialog v-model="showCreateDialog" :title="isEdit ? '编辑基因组数据' : '新增基因组数据'" width="600px">
      <el-form :model="genomicDataForm" :rules="genomicDataRules" ref="genomicDataFormRef" label-width="100px">
        <el-form-item label="样本ID" prop="sampleId">
          <el-input v-model="genomicDataForm.sampleId" type="number" />
        </el-form-item>
        <el-form-item label="测序平台" prop="sequencePlatform">
          <el-input v-model="genomicDataForm.sequencePlatform" />
        </el-form-item>
        <el-form-item label="测序类型">
          <el-select v-model="genomicDataForm.sequenceType" placeholder="请选择测序类型">
            <el-option label="全基因组测序" value="全基因组测序" />
            <el-option label="外显子测序" value="外显子测序" />
            <el-option label="转录组测序" value="转录组测序" />
            <el-option label="单细胞测序" value="单细胞测序" />
          </el-select>
        </el-form-item>
        <el-form-item label="基因类型">
          <el-input v-model="genomicDataForm.geneType" />
        </el-form-item>
        <el-form-item label="结果文件路径">
          <el-input v-model="genomicDataForm.resultFilePath" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showCreateDialog = false">取消</el-button>
        <el-button type="primary" @click="saveGenomicData">确定</el-button>
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
const genomicDataList = ref([
  {
    id: 1,
    sampleId: 'S001',
    sequencePlatform: 'Illumina NovaSeq 6000',
    sequenceType: '全基因组测序',
    geneType: '过敏相关基因',
    resultFilePath: '/data/genomics/S001_WGS.vcf'
  },
  {
    id: 2,
    sampleId: 'S002',
    sequencePlatform: 'BGI MGISEQ-2000',
    sequenceType: '外显子测序',
    geneType: '免疫相关基因',
    resultFilePath: '/data/genomics/S002_WES.vcf'
  },
  {
    id: 3,
    sampleId: 'S003',
    sequencePlatform: 'Thermo Fisher Ion Proton',
    sequenceType: '靶向测序',
    geneType: '药物代谢基因',
    resultFilePath: '/data/genomics/S003_Target.vcf'
  }
])
const loading = ref(false)
const showCreateDialog = ref(false)
const isEdit = ref(false)

// 搜索表单
const searchForm = reactive({
  sampleId: '',
  sequencePlatform: '',
  geneType: ''
})

// 基因组数据表单
const genomicDataForm = reactive({
  id: 0,
  sampleId: '',
  sequencePlatform: '',
  sequenceType: '',
  geneType: '',
  resultFilePath: ''
})

const genomicDataRules = {
  sampleId: [{ required: true, message: '请输入样本ID', trigger: 'blur' }],
  sequencePlatform: [{ required: true, message: '请输入测序平台', trigger: 'blur' }]
}

const genomicDataFormRef = ref()

// 方法
const loadGenomicData = async () => {
  loading.value = true
  try {
    // 暂时显示所有样本的基因组数据（通过循环查询）
    const allGenomicData = []
    const sampleIds = [1, 2, 3] // 这里可以根据实际需要调整
    
    for (const sampleId of sampleIds) {
      try {
        const resp = await biobankApi.getSampleGenomicData(sampleId)
        if (resp.data.data && resp.data.data.length > 0) {
          allGenomicData.push(...resp.data.data)
        }
      } catch (error) {
        // 忽略没有基因组数据的样本
      }
    }
    
    // 根据搜索条件过滤
    let filteredData = allGenomicData
    if (searchForm.sampleId) {
      filteredData = filteredData.filter(item => 
        item.sampleId.toString().includes(searchForm.sampleId)
      )
    }
    if (searchForm.sequencePlatform) {
      filteredData = filteredData.filter(item => 
        item.sequencePlatform && item.sequencePlatform.includes(searchForm.sequencePlatform)
      )
    }
    if (searchForm.geneType) {
      filteredData = filteredData.filter(item => 
        item.geneType && item.geneType.includes(searchForm.geneType)
      )
    }
    
    genomicDataList.value = filteredData
  } catch (error) {
    ElMessage.error('加载基因组数据失败')
    console.error('加载基因组数据失败:', error)
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.sampleId = ''
  searchForm.sequencePlatform = ''
  searchForm.geneType = ''
  loadGenomicData()
}

const resetForm = () => {
  genomicDataForm.id = 0
  genomicDataForm.sampleId = ''
  genomicDataForm.sequencePlatform = ''
  genomicDataForm.sequenceType = ''
  genomicDataForm.geneType = ''
  genomicDataForm.resultFilePath = ''
}

const createGenomicData = async () => {
  if (!genomicDataFormRef.value) return
  
  try {
    await genomicDataFormRef.value.validate()
    
    const genomicData = {
      ...genomicDataForm,
      sampleId: parseInt(genomicDataForm.sampleId)
    }
    
    await biobankApi.addGenomicData(genomicData)
    ElMessage.success('创建基因组数据成功')
    showCreateDialog.value = false
    loadGenomicData()
  } catch (error) {
    if (error !== false) {
      ElMessage.error('创建基因组数据失败')
    }
  }
}

const editGenomicData = (genomicData) => {
  isEdit.value = true
  genomicDataForm.id = genomicData.id
  genomicDataForm.sampleId = genomicData.sampleId
  genomicDataForm.sequencePlatform = genomicData.sequencePlatform || ''
  genomicDataForm.sequenceType = genomicData.sequenceType || ''
  genomicDataForm.geneType = genomicData.geneType || ''
  genomicDataForm.resultFilePath = genomicData.resultFilePath || ''
  showCreateDialog.value = true
}

const saveGenomicData = async () => {
  if (isEdit.value) {
    // 编辑功能暂时显示提示
    ElMessage.info('编辑功能暂未实现')
  } else {
    await createGenomicData()
  }
}

const viewGenomicData = (genomicData) => {
  console.log('查看基因组数据:', genomicData)
  ElMessage.info('查看功能暂未实现')
}

const deleteGenomicData = async (genomicData) => {
  try {
    await ElMessageBox.confirm('确定要删除这条基因组数据吗？', '提示', {
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
  genomicDataList.value = [
    {
      id: 1,
      sampleId: 'S001',
      sequencePlatform: 'Illumina NovaSeq 6000',
      sequenceType: '全基因组测序',
      geneType: '过敏相关基因',
      resultFilePath: '/data/genomics/S001_WGS.vcf'
    },
    {
      id: 2,
      sampleId: 'S002',
      sequencePlatform: 'BGI MGISEQ-2000',
      sequenceType: '外显子测序',
      geneType: '免疫相关基因',
      resultFilePath: '/data/genomics/S002_WES.vcf'
    },
    {
      id: 3,
      sampleId: 'S003',
      sequencePlatform: 'Thermo Fisher Ion Proton',
      sequenceType: '靶向测序',
      geneType: '药物代谢基因',
      resultFilePath: '/data/genomics/S003_Target.vcf'
    }
  ];
})
</script>

<style scoped>
.genomic-data {
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