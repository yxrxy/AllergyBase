<template>
  <div class="followup-list">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>随访管理</span>
          <el-button type="primary" @click="showCreateDialog">创建随访计划</el-button>
        </div>
      </template>

      <el-form :model="searchForm" :inline="true" class="search-form">
        <el-form-item label="患者ID">
          <el-input v-model="searchForm.patientId" placeholder="请输入患者ID" clearable />
        </el-form-item>
        <el-form-item label="随访状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="计划中" value="planned" />
            <el-option label="进行中" value="ongoing" />
            <el-option label="已完成" value="completed" />
            <el-option label="已取消" value="cancelled" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadData">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="随访ID" width="100" />
        <el-table-column prop="patientId" label="患者ID" width="100" />
        <el-table-column prop="planName" label="计划名称" width="150" />
        <el-table-column prop="startDate" label="开始日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.startDate) }}
          </template>
        </el-table-column>
        <el-table-column prop="endDate" label="结束日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.endDate) }}
          </template>
        </el-table-column>
        <el-table-column prop="frequency" label="随访频率" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">{{ getStatusText(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="nextFollowupDate" label="下次随访" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.nextFollowupDate) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" @click="viewFollowup(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editFollowup(scope.row)">编辑</el-button>
            <el-button size="small" type="success" @click="addRecord(scope.row)">添加记录</el-button>
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

    <!-- 创建随访计划对话框 -->
    <el-dialog v-model="createDialogVisible" title="创建随访计划" width="600px">
      <el-form :model="createForm" :rules="rules" ref="createFormRef" label-width="100px">
        <el-form-item label="患者ID" prop="patientId">
          <el-input v-model="createForm.patientId" />
        </el-form-item>
        <el-form-item label="计划名称" prop="planName">
          <el-input v-model="createForm.planName" />
        </el-form-item>
        <el-form-item label="开始日期" prop="startDate">
          <el-date-picker v-model="createForm.startDate" type="date" placeholder="选择开始日期" />
        </el-form-item>
        <el-form-item label="结束日期" prop="endDate">
          <el-date-picker v-model="createForm.endDate" type="date" placeholder="选择结束日期" />
        </el-form-item>
        <el-form-item label="随访频率" prop="frequency">
          <el-select v-model="createForm.frequency" placeholder="请选择频率">
            <el-option label="每周" value="weekly" />
            <el-option label="每月" value="monthly" />
            <el-option label="每季度" value="quarterly" />
            <el-option label="每年" value="yearly" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="createForm.description" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createFollowup">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { followupApi } from '@/api/followup'

const router = useRouter()

const loading = ref(false)
const createDialogVisible = ref(false)
const tableData = ref([])

const searchForm = ref({
  patientId: '',
  status: ''
})

const createForm = ref({
  patientId: '',
  planName: '',
  startDate: '',
  endDate: '',
  frequency: '',
  description: ''
})

const pagination = ref({
  page: 1,
  size: 10,
  total: 0
})

const createFormRef = ref()

const rules = {
  patientId: [{ required: true, message: '请输入患者ID', trigger: 'blur' }],
  planName: [{ required: true, message: '请输入计划名称', trigger: 'blur' }],
  startDate: [{ required: true, message: '请选择开始日期', trigger: 'change' }],
  endDate: [{ required: true, message: '请选择结束日期', trigger: 'change' }],
  frequency: [{ required: true, message: '请选择随访频率', trigger: 'change' }]
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('zh-CN')
}

const getStatusType = (status) => {
  const statusMap = {
    'planned': 'info',
    'ongoing': 'warning',
    'completed': 'success',
    'cancelled': 'danger'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status) => {
  const statusText = {
    'planned': '计划中',
    'ongoing': '进行中',
    'completed': '已完成',
    'cancelled': '已取消'
  }
  return statusText[status] || '未知状态'
}

const loadData = async () => {
  loading.value = true
  try {
    const params = {
      ...searchForm.value,
      page: pagination.value.page,
      size: pagination.value.size
    }
    const response = await followupApi.getFollowupList(params)
    tableData.value = response.data.followups || []
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
    status: ''
  }
  pagination.value.page = 1
  loadData()
}

const showCreateDialog = () => {
  createForm.value = {
    patientId: '',
    planName: '',
    startDate: '',
    endDate: '',
    frequency: '',
    description: ''
  }
  createDialogVisible.value = true
}

const createFollowup = async () => {
  try {
    await createFormRef.value.validate()
    await followupApi.createFollowup(createForm.value)
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

const viewFollowup = (followup) => {
  router.push(`/followup/${followup.id}`)
}

const editFollowup = (followup) => {
  router.push(`/followup/${followup.id}/edit`)
}

const addRecord = (followup) => {
  router.push(`/followup/${followup.id}/add-record`)
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.followup-list {
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