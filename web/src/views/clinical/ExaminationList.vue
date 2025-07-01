<template>
  <div class="examination-list">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>检查记录管理</span>
          <el-button type="primary" @click="showCreateDialog">新建检查记录</el-button>
        </div>
      </template>

      <el-form :model="searchForm" :inline="true" class="search-form">
        <el-form-item label="就诊ID">
          <el-input v-model="searchForm.visitId" placeholder="请输入就诊ID" clearable />
        </el-form-item>
        <el-form-item label="检查类型">
          <el-input v-model="searchForm.examType" placeholder="请输入检查类型" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadData">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="检查ID" width="100" />
        <el-table-column prop="visitId" label="就诊ID" width="100" />
        <el-table-column prop="examType" label="检查类型" width="150" />
        <el-table-column prop="examTime" label="检查时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.examTime) }}
          </template>
        </el-table-column>
        <el-table-column prop="examResult" label="检查结果" min-width="200" show-overflow-tooltip />
        <el-table-column prop="resultUnit" label="单位" width="100" />
        <el-table-column prop="referenceRange" label="参考范围" width="150" />
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" @click="viewExamination(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editExamination(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteExamination(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 创建检查记录对话框 -->
    <el-dialog v-model="createDialogVisible" title="新建检查记录" width="600px">
      <el-form :model="createForm" :rules="rules" ref="createFormRef" label-width="100px">
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="就诊ID" prop="visitId">
              <el-input v-model="createForm.visitId" type="number" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="检查类型" prop="examType">
              <el-input v-model="createForm.examType" placeholder="如：血常规、X光等" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="检查时间" prop="examTime">
          <el-date-picker
            v-model="createForm.examTime"
            type="datetime"
            placeholder="选择检查时间"
            format="YYYY-MM-DD HH:mm"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="检查结果" prop="examResult">
          <el-input v-model="createForm.examResult" type="textarea" rows="3" />
        </el-form-item>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="结果单位">
              <el-input v-model="createForm.resultUnit" placeholder="如：个/μL" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="参考范围">
              <el-input v-model="createForm.referenceRange" placeholder="如：4000-10000" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createExamination">确定</el-button>
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
  examType: ''
})

const createForm = ref({
  visitId: '',
  examType: '',
  examTime: '',
  examResult: '',
  resultUnit: '',
  referenceRange: ''
})

const createFormRef = ref()

const rules = {
  visitId: [{ required: true, message: '请输入就诊ID', trigger: 'blur' }],
  examType: [{ required: true, message: '请输入检查类型', trigger: 'blur' }],
  examTime: [{ required: true, message: '请选择检查时间', trigger: 'change' }],
  examResult: [{ required: true, message: '请输入检查结果', trigger: 'blur' }]
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleString('zh-CN')
}

const loadData = async () => {
  loading.value = true
  try {
    if (searchForm.value.visitId) {
      // 如果指定了就诊ID，获取特定就诊的检查记录
      const response = await clinicalApi.getVisitExaminations(searchForm.value.visitId)
      tableData.value = response.data.examinations || []
    } else {
      // 如果没有指定就诊ID，获取所有检查记录
      const response = await clinicalApi.listAllExaminations({
        examType: searchForm.value.examType
      })
      tableData.value = response.data.examinations || []
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
    examType: ''
  }
  loadData() // 重置后重新加载所有数据
}

const showCreateDialog = () => {
  createForm.value = {
    visitId: '',
    examType: '',
    examTime: '',
    examResult: '',
    resultUnit: '',
    referenceRange: ''
  }
  createDialogVisible.value = true
}

const createExamination = async () => {
  try {
    await createFormRef.value.validate()
    
    const examinationData = {
      id: 0,
      visitId: parseInt(createForm.value.visitId),
      examType: createForm.value.examType,
      examTime: createForm.value.examTime,
      examResult: createForm.value.examResult,
      resultUnit: createForm.value.resultUnit,
      referenceRange: createForm.value.referenceRange
    }
    
    await clinicalApi.addExamination({
      examination: examinationData
    })
    ElMessage.success('检查记录创建成功')
    createDialogVisible.value = false
    
    // 如果当前搜索的就诊ID与新建的相同，刷新数据
    if (searchForm.value.visitId === createForm.value.visitId) {
      await loadData()
    }
  } catch (error) {
    console.error('创建检查记录失败:', error)
    ElMessage.error('创建检查记录失败')
  }
}

const viewExamination = (examination) => {
  ElMessage.info('查看检查功能待实现')
}

const editExamination = (examination) => {
  ElMessage.info('编辑检查功能待实现')
}

const deleteExamination = async (examination) => {
  try {
    await ElMessageBox.confirm('确定要删除这个检查记录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await clinicalApi.deleteExamination(examination.id)
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
  loadData() // 页面加载时自动获取所有检查记录
})
</script>

<style scoped>
.examination-list {
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