<template>
  <div class="medication-monitor">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>药物监测管理</span>
          <el-button type="primary" @click="showCreateDialog">添加监测记录</el-button>
        </div>
      </template>

      <el-form :model="searchForm" :inline="true" class="search-form">
        <el-form-item label="患者ID">
          <el-input v-model="searchForm.patientId" placeholder="请输入患者ID" clearable />
        </el-form-item>
        <el-form-item label="药物名称">
          <el-input v-model="searchForm.medicationName" placeholder="请输入药物名称" clearable />
        </el-form-item>
        <el-form-item label="监测状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="进行中" value="ongoing" />
            <el-option label="已完成" value="completed" />
            <el-option label="已停止" value="stopped" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadData">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <el-table :data="tableData" v-loading="loading">
        <el-table-column prop="id" label="监测ID" width="100" />
        <el-table-column prop="patientId" label="患者ID" width="100" />
        <el-table-column prop="medicationName" label="药物名称" width="150" />
        <el-table-column prop="dosage" label="剂量" width="100" />
        <el-table-column prop="frequency" label="用药频率" width="120" />
        <el-table-column prop="startDate" label="开始日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.startDate) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">{{ getStatusText(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sideEffects" label="副作用" />
        <el-table-column label="操作" width="240">
          <template #default="scope">
            <el-button size="small" @click="viewMonitor(scope.row)">查看</el-button>
            <el-button size="small" type="primary" @click="editMonitor(scope.row)">编辑</el-button>
            <el-button size="small" type="warning" @click="addRecord(scope.row)">添加记录</el-button>
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

    <!-- 创建监测记录对话框 -->
    <el-dialog v-model="createDialogVisible" title="添加药物监测记录" width="600px">
      <el-form :model="createForm" :rules="rules" ref="createFormRef" label-width="100px">
        <el-form-item label="患者ID" prop="patientId">
          <el-input v-model="createForm.patientId" />
        </el-form-item>
        <el-form-item label="药物名称" prop="medicationName">
          <el-input v-model="createForm.medicationName" />
        </el-form-item>
        <el-form-item label="剂量" prop="dosage">
          <el-input v-model="createForm.dosage" />
        </el-form-item>
        <el-form-item label="用药频率" prop="frequency">
          <el-select v-model="createForm.frequency" placeholder="请选择频率">
            <el-option label="每日一次" value="once_daily" />
            <el-option label="每日两次" value="twice_daily" />
            <el-option label="每日三次" value="three_times_daily" />
            <el-option label="每周一次" value="once_weekly" />
          </el-select>
        </el-form-item>
        <el-form-item label="开始日期" prop="startDate">
          <el-date-picker v-model="createForm.startDate" type="date" placeholder="选择开始日期" />
        </el-form-item>
        <el-form-item label="预计结束日期">
          <el-date-picker v-model="createForm.endDate" type="date" placeholder="选择结束日期" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="createForm.notes" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createMonitor">确定</el-button>
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
  medicationName: '',
  status: ''
})

const createForm = ref({
  patientId: '',
  medicationName: '',
  dosage: '',
  frequency: '',
  startDate: '',
  endDate: '',
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
  medicationName: [{ required: true, message: '请输入药物名称', trigger: 'blur' }],
  dosage: [{ required: true, message: '请输入剂量', trigger: 'blur' }],
  frequency: [{ required: true, message: '请选择用药频率', trigger: 'change' }],
  startDate: [{ required: true, message: '请选择开始日期', trigger: 'change' }]
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('zh-CN')
}

const getStatusType = (status) => {
  const statusMap = {
    'ongoing': 'warning',
    'completed': 'success',
    'stopped': 'danger'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status) => {
  const statusMap = {
    'ongoing': '进行中',
    'completed': '已完成',
    'stopped': '已停止'
  }
  return statusMap[status] || status
}

const loadData = async () => {
  loading.value = true
  try {
    tableData.value = [
      {
        id: 1,
        patientId: 'P001',
        medicationName: '氯雷他定',
        dosage: '10mg',
        frequency: '每日一次',
        startDate: '2024-06-01',
        status: 'ongoing',
        sideEffects: '轻微嗜睡'
      },
      {
        id: 2,
        patientId: 'P002',
        medicationName: '布地奈德鼻喷剂',
        dosage: '64μg',
        frequency: '每日两次',
        startDate: '2024-06-02',
        status: 'completed',
        sideEffects: '无'
      },
      {
        id: 3,
        patientId: 'P003',
        medicationName: '孟鲁司特',
        dosage: '10mg',
        frequency: '每日一次',
        startDate: '2024-06-03',
        status: 'stopped',
        sideEffects: '偶有头痛'
      }
    ];
    pagination.value.total = 3;
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.value = {
    patientId: '',
    medicationName: '',
    status: ''
  }
  pagination.value.page = 1
  loadData()
}

const showCreateDialog = () => {
  createForm.value = {
    patientId: '',
    medicationName: '',
    dosage: '',
    frequency: '',
    startDate: '',
    endDate: '',
    notes: ''
  }
  createDialogVisible.value = true
}

const createMonitor = async () => {
  try {
    await createFormRef.value.validate()
    await followupApi.createMedicationMonitor(createForm.value)
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

const viewMonitor = (monitor) => {
  router.push(`/followup/medication/${monitor.id}`)
}

const editMonitor = (monitor) => {
  router.push(`/followup/medication/${monitor.id}/edit`)
}

const addRecord = (monitor) => {
  router.push(`/followup/medication/${monitor.id}/record/new`)
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.medication-monitor {
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