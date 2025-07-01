<template>
  <div class="diagnosis-list">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>诊断记录管理</span>
          <el-button type="primary" @click="showCreateDialog">新建诊断记录</el-button>
        </div>
      </template>

      <el-form :model="searchForm" :inline="true" class="search-form">
        <el-form-item label="就诊ID">
          <el-input v-model="searchForm.visitId" placeholder="请输入就诊ID" clearable />
        </el-form-item>
        <el-form-item label="诊断类型">
          <el-input v-model="searchForm.diagnosisType" placeholder="请输入诊断类型" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadData">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="诊断ID" width="100" />
        <el-table-column prop="visitId" label="就诊ID" width="100" />
        <el-table-column prop="diagnosisType" label="诊断类型" width="150" />
        <el-table-column prop="icdCode" label="ICD编码" width="150" />
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" @click="viewDiagnosis(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editDiagnosis(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteDiagnosis(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 创建诊断记录对话框 -->
    <el-dialog v-model="createDialogVisible" title="新建诊断记录" width="500px">
      <el-form :model="createForm" :rules="rules" ref="createFormRef" label-width="100px">
        <el-form-item label="就诊ID" prop="visitId">
          <el-input v-model="createForm.visitId" type="number" />
        </el-form-item>
        <el-form-item label="诊断类型" prop="diagnosisType">
          <el-select v-model="createForm.diagnosisType" placeholder="请选择诊断类型">
            <el-option label="主要诊断" value="主要诊断" />
            <el-option label="次要诊断" value="次要诊断" />
            <el-option label="疑似诊断" value="疑似诊断" />
            <el-option label="排除诊断" value="排除诊断" />
          </el-select>
        </el-form-item>
        <el-form-item label="ICD编码" prop="icdCode">
          <el-input v-model="createForm.icdCode" placeholder="请输入ICD编码" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createDiagnosis">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { clinicalApi } from '@/api/clinical'

const loading = ref(false)
const createDialogVisible = ref(false)
const tableData = ref([])

const searchForm = ref({
  visitId: '',
  diagnosisType: ''
})

const createForm = ref({
  visitId: '',
  diagnosisType: '',
  icdCode: ''
})

const createFormRef = ref()

const rules = {
  visitId: [{ required: true, message: '请输入就诊ID', trigger: 'blur' }],
  diagnosisType: [{ required: true, message: '请选择诊断类型', trigger: 'change' }],
  icdCode: [{ required: true, message: '请输入ICD编码', trigger: 'blur' }]
}

const loadData = async () => {
  loading.value = true
  try {
    if (searchForm.value.visitId) {
      const response = await clinicalApi.getVisitDiagnoses(searchForm.value.visitId)
      tableData.value = response.data.diagnoses || []
    } else {
      // 如果没有指定就诊ID，获取所有诊断记录
      const response = await clinicalApi.listAllDiagnoses({
        offset: 0,
        limit: 100
      })
      tableData.value = response.data.diagnoses || []
    }
  } catch (error) {
    ElMessage.error('加载数据失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.value = {
    visitId: '',
    diagnosisType: ''
  }
  tableData.value = []
}

const showCreateDialog = () => {
  createForm.value = {
    visitId: '',
    diagnosisType: '',
    icdCode: ''
  }
  createDialogVisible.value = true
}

const createDiagnosis = async () => {
  try {
    await createFormRef.value.validate()
    
    const diagnosisData = {
      id: 0,
      visitId: parseInt(createForm.value.visitId),
      diagnosisType: createForm.value.diagnosisType,
      icdCode: createForm.value.icdCode
    }
    
    await clinicalApi.addDiagnosis({
      diagnosis: diagnosisData
    })
    ElMessage.success('诊断记录创建成功')
    createDialogVisible.value = false
    
    // 如果当前搜索的就诊ID与新建的相同，刷新数据
    if (searchForm.value.visitId === createForm.value.visitId) {
      await loadData()
    }
  } catch (error) {
    console.error('创建诊断记录失败:', error)
    ElMessage.error('创建诊断记录失败')
  }
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
    loadData()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
      console.error(error)
    }
  }
}

onMounted(() => {
  // 页面加载时自动加载所有诊断记录
  loadData()
})
</script>

<style scoped>
.diagnosis-list {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-form {
  margin-bottom: 20px;
}
</style> 