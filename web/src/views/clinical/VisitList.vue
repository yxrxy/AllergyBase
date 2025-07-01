<template>
  <div class="visit-list">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>就诊记录管理</span>
          <el-button type="primary" @click="showCreateDialog">新建就诊记录</el-button>
        </div>
      </template>

      <el-form :model="searchForm" :inline="true" class="search-form">
        <el-form-item label="患者ID">
          <el-input v-model="searchForm.patientId" placeholder="请输入患者ID" clearable />
        </el-form-item>
        <el-form-item label="科室">
          <el-input v-model="searchForm.department" placeholder="请输入科室" clearable />
        </el-form-item>
        <el-form-item label="就诊编号">
          <el-input v-model="searchForm.visitNo" placeholder="请输入就诊编号" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadData">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="就诊ID" width="100" />
        <el-table-column prop="patientId" label="患者ID" width="100" />
        <el-table-column prop="visitNo" label="就诊编号" width="120" />
        <el-table-column prop="visitTime" label="就诊时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.visitTime) }}
          </template>
        </el-table-column>
        <el-table-column prop="visitType" label="就诊类型" width="100">
          <template #default="scope">
            {{ getVisitTypeText(scope.row.visitType) }}
          </template>
        </el-table-column>
        <el-table-column prop="department" label="科室" width="120" />
        <el-table-column prop="doctorId" label="医生ID" width="100" />
        <el-table-column prop="chiefComplaint" label="主诉" min-width="200" show-overflow-tooltip />
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" @click="viewVisit(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editVisit(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteVisit(scope.row)">删除</el-button>
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

    <!-- 创建就诊记录对话框 -->
    <el-dialog v-model="createDialogVisible" title="新建就诊记录" width="600px">
      <el-form :model="createForm" :rules="rules" ref="createFormRef" label-width="100px">
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="患者ID" prop="patientId">
              <el-input v-model="createForm.patientId" type="number" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="就诊编号" prop="visitNo">
              <el-input v-model="createForm.visitNo" placeholder="留空自动生成" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="就诊时间" prop="visitTime">
              <el-date-picker
                v-model="createForm.visitTime"
                type="datetime"
                placeholder="选择就诊时间"
                format="YYYY-MM-DD HH:mm"
                value-format="YYYY-MM-DD HH:mm:ss"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="就诊类型" prop="visitType">
              <el-select v-model="createForm.visitType" placeholder="请选择就诊类型">
                <el-option label="门诊" :value="1" />
                <el-option label="急诊" :value="2" />
                <el-option label="住院" :value="3" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="科室" prop="department">
              <el-input v-model="createForm.department" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="医生ID">
              <el-input v-model="createForm.doctorId" type="number" placeholder="可选" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="主诉" prop="chiefComplaint">
          <el-input v-model="createForm.chiefComplaint" type="textarea" rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createVisit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { clinicalApi } from '@/api/clinical'

const router = useRouter()

const loading = ref(false)
const createDialogVisible = ref(false)
const tableData = ref([])

const searchForm = ref({
  patientId: '',
  department: '',
  visitNo: ''
})

const createForm = ref({
  patientId: '',
  visitNo: '',
  visitTime: '',
  visitType: 1,
  department: '',
  doctorId: '',
  chiefComplaint: ''
})

const pagination = ref({
  page: 1,
  size: 10,
  total: 0
})

const createFormRef = ref()

const rules = {
  patientId: [{ required: true, message: '请输入患者ID', trigger: 'blur' }],
  visitTime: [{ required: true, message: '请选择就诊时间', trigger: 'change' }],
  visitType: [{ required: true, message: '请选择就诊类型', trigger: 'change' }],
  department: [{ required: true, message: '请输入科室', trigger: 'blur' }],
  chiefComplaint: [{ required: true, message: '请输入主诉', trigger: 'blur' }]
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleString('zh-CN')
}

const getVisitTypeText = (visitType) => {
  const typeMap = {
    1: '门诊',
    2: '急诊',
    3: '住院'
  }
  return typeMap[visitType] || '未知'
}

const loadData = async () => {
  loading.value = true
  try {
    // 如果有患者ID搜索条件，使用患者就诊记录API
    if (searchForm.value.patientId) {
      const response = await clinicalApi.getPatientVisits(searchForm.value.patientId)
      tableData.value = response.data.visits || []
      pagination.value.total = tableData.value.length
    } else {
      // 否则使用通用列表API
      const params = {
        ...searchForm.value,
        offset: (pagination.value.page - 1) * pagination.value.size,
        limit: pagination.value.size
      }
      const response = await clinicalApi.listVisits(params)
      tableData.value = response.data.visits || []
      pagination.value.total = response.data.total || 0
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
    patientId: '',
    department: '',
    visitNo: ''
  }
  pagination.value.page = 1
  loadData()
}

const showCreateDialog = () => {
  createForm.value = {
    patientId: '',
    visitNo: '',
    visitTime: '',
    visitType: 1,
    department: '',
    doctorId: '',
    chiefComplaint: ''
  }
  createDialogVisible.value = true
}

const createVisit = async () => {
  try {
    await createFormRef.value.validate()
    
    // 构造符合后端API要求的请求数据
    const visitData = {
      id: 0, // 新建记录ID为0
      patientId: parseInt(createForm.value.patientId),
      visitTime: createForm.value.visitTime,
      visitType: createForm.value.visitType,
      visitNo: createForm.value.visitNo || `V${Date.now()}`, // 如果没有填写就诊编号，自动生成
      department: createForm.value.department,
      chiefComplaint: createForm.value.chiefComplaint
    }
    
    // 如果填写了医生ID，添加到请求中
    if (createForm.value.doctorId) {
      visitData.doctorId = parseInt(createForm.value.doctorId)
    }
    
    await clinicalApi.createVisit({
      visit: visitData
    })
    ElMessage.success('就诊记录创建成功')
    createDialogVisible.value = false
    await loadData()
  } catch (error) {
    console.error('创建就诊记录失败:', error)
    ElMessage.error('创建就诊记录失败')
  }
}

const viewVisit = (visit) => {
  router.push(`/clinical/visit/${visit.id}`)
}

const editVisit = (visit) => {
  router.push(`/clinical/visit/${visit.id}/edit`)
}

const deleteVisit = async (visit) => {
  try {
    await ElMessageBox.confirm('确定要删除这个就诊记录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await clinicalApi.deleteVisit(visit.id)
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
.visit-list {
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