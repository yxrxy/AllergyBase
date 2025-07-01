<template>
  <div class="epidemiology-list">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>流行病学调查</span>
          <el-button type="primary" @click="showCreateDialog">新建调查</el-button>
        </div>
      </template>

      <el-form :model="searchForm" :inline="true" class="search-form">
        <el-form-item label="患者ID">
          <el-input v-model="searchForm.patientId" placeholder="请输入患者ID" clearable />
        </el-form-item>
        <el-form-item label="调查类型">
          <el-select v-model="searchForm.surveyType" placeholder="请选择调查类型" clearable>
            <el-option label="环境暴露" value="environment" />
            <el-option label="生活方式" value="lifestyle" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadData">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="调查ID" width="100" />
        <el-table-column prop="patientId" label="患者ID" width="100" />
        <el-table-column prop="surveyType" label="调查类型" width="120">
          <template #default="scope">
            <el-tag :type="scope.row.surveyType === 'environment' ? 'success' : 'info'">
              {{ scope.row.surveyType === 'environment' ? '环境暴露' : '生活方式' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="surveyDate" label="调查日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.surveyDate) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="150">
          <template #default="scope">
            {{ formatDate(scope.row.createdAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" @click="viewSurvey(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editSurvey(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteSurvey(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.size"
        :total="pagination.total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="loadData"
        @current-change="loadData"
      />
    </el-card>

    <!-- 创建调查对话框 -->
    <el-dialog v-model="createDialogVisible" title="新建流行病学调查" width="600px">
      <el-form :model="createForm" :rules="rules" ref="createFormRef" label-width="100px">
        <el-form-item label="患者ID" prop="patientId">
          <el-input v-model="createForm.patientId" />
        </el-form-item>
        <el-form-item label="调查类型" prop="surveyType">
          <el-select v-model="createForm.surveyType" placeholder="请选择调查类型">
            <el-option label="环境暴露" value="environment" />
            <el-option label="生活方式" value="lifestyle" />
          </el-select>
        </el-form-item>
        <el-form-item label="调查日期" prop="surveyDate">
          <el-date-picker v-model="createForm.surveyDate" type="date" placeholder="选择日期" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="createForm.notes" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createSurvey">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { epidemiologyApi } from '@/api/epidemiology'

const router = useRouter()

const loading = ref(false)
const createDialogVisible = ref(false)
const tableData = ref([])

const searchForm = ref({
  patientId: '',
  surveyType: ''
})

const createForm = ref({
  patientId: '',
  surveyType: '',
  surveyDate: '',
  notes: ''
})

const pagination = ref({
  page: 1,
  size: 10,
  total: 0
})

const createFormRef = ref()

const rules = {
  patientId: [{ required: true, message: '请输入患者ID', trigger: 'blur' }],
  surveyType: [{ required: true, message: '请选择调查类型', trigger: 'change' }],
  surveyDate: [{ required: true, message: '请选择调查日期', trigger: 'change' }]
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('zh-CN')
}

const getStatusType = (status) => {
  const statusMap = {
    '进行中': 'warning',
    '已完成': 'success',
    '已取消': 'danger'
  }
  return statusMap[status] || 'info'
}

const loadData = async () => {
  loading.value = true
  try {
    const params = {
      ...searchForm.value,
      page: pagination.value.page,
      size: pagination.value.size
    }
    const response = await epidemiologyApi.getSurveyList(params)
    tableData.value = response.data.surveys || []
    pagination.value.total = response.data.total || 0
  } catch (error) {
    ElMessage.error('加载数据失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.value = {
    patientId: '',
    surveyType: ''
  }
  pagination.value.page = 1
  loadData()
}

const showCreateDialog = () => {
  createForm.value = {
    patientId: '',
    surveyType: '',
    surveyDate: '',
    notes: ''
  }
  createDialogVisible.value = true
}

const createSurvey = async () => {
  try {
    await createFormRef.value.validate()
    await epidemiologyApi.createSurvey(createForm.value)
    createDialogVisible.value = false
    ElMessage.success('创建成功')
    loadData()
  } catch (error) {
    if (error !== false) {
      ElMessage.error('创建失败')
      console.error(error)
    }
  }
}

const viewSurvey = (survey) => {
  router.push(`/epidemiology/surveys/${survey.id}`)
}

const editSurvey = (survey) => {
  router.push(`/epidemiology/surveys/${survey.id}/edit`)
}

const deleteSurvey = async (survey) => {
  try {
    await ElMessageBox.confirm('确定要删除这个调查吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await epidemiologyApi.deleteSurvey(survey.id)
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
  loadData()
})
</script>

<style scoped>
.epidemiology-list {
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