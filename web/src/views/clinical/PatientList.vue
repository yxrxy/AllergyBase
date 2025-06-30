<template>
  <div class="patient-list">
    <div class="header">
      <h2>患者管理</h2>
      <el-button type="primary" @click="showCreateDialog = true">
        <el-icon><Plus /></el-icon>
        新增患者
      </el-button>
    </div>

    <!-- 搜索区域 -->
    <div class="search-area">
      <el-form :model="searchForm" inline>
        <el-form-item label="患者编号">
          <el-input v-model="searchForm.patientNo" placeholder="请输入患者编号" clearable />
        </el-form-item>
        <el-form-item label="姓名">
          <el-input v-model="searchForm.name" placeholder="请输入患者姓名" clearable />
        </el-form-item>
        <el-form-item label="身份证号">
          <el-input v-model="searchForm.idCard" placeholder="请输入身份证号" clearable />
        </el-form-item>
        <el-form-item label="电话号码">
          <el-input v-model="searchForm.phone" placeholder="请输入电话号码" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="searchPatients">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 患者列表 -->
    <el-table :data="patientList" v-loading="loading" stripe>
      <el-table-column prop="patientNo" label="患者编号" width="120" />
      <el-table-column prop="name" label="姓名" width="100" />
      <el-table-column prop="gender" label="性别" width="80">
        <template #default="{ row }">
          {{ row.gender === 1 ? '男' : '女' }}
        </template>
      </el-table-column>
      <el-table-column prop="birthDate" label="出生日期" width="120" />
      <el-table-column prop="phone" label="电话号码" width="130" />
      <el-table-column prop="idCard" label="身份证号" width="150" />
      <el-table-column prop="height" label="身高(cm)" width="90" />
      <el-table-column prop="weight" label="体重(kg)" width="90" />
      <el-table-column label="操作" width="220" fixed="right">
        <template #default="{ row }">
          <div class="action-buttons">
            <el-button size="small" @click="viewPatient(row)">查看</el-button>
            <el-button size="small" type="primary" @click="editPatient(row)">编辑</el-button>
            <el-button size="small" @click="viewVisits(row)">就诊记录</el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="pagination">
      <el-pagination
        v-model:current-page="pagination.currentPage"
        v-model:page-size="pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

    <!-- 创建/编辑患者对话框 -->
    <el-dialog v-model="showCreateDialog" :title="patientForm.id ? '编辑患者' : '新增患者'" width="800px">
      <el-form :model="patientForm" :rules="patientRules" ref="patientFormRef" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="患者编号" prop="patientNo">
              <el-input v-model="patientForm.patientNo" placeholder="请输入患者编号" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="姓名" prop="name">
              <el-input v-model="patientForm.name" placeholder="请输入患者姓名" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="性别" prop="gender">
              <el-select v-model="patientForm.gender" placeholder="请选择性别">
                <el-option label="男" :value="1" />
                <el-option label="女" :value="0" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="出生日期" prop="birthDate">
              <el-date-picker 
                v-model="patientForm.birthDate" 
                type="date" 
                format="YYYY-MM-DD"
                value-format="YYYY-MM-DD"
                placeholder="请选择出生日期" 
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="身份证号">
              <el-input v-model="patientForm.idCard" placeholder="请输入身份证号" maxlength="18" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="电话号码">
              <el-input v-model="patientForm.phone" placeholder="请输入电话号码" maxlength="11" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="地址">
          <el-input v-model="patientForm.address" placeholder="请输入详细地址" />
        </el-form-item>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="紧急联系人">
              <el-input v-model="patientForm.emergencyContact" placeholder="请输入紧急联系人姓名" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="紧急联系电话">
              <el-input v-model="patientForm.emergencyPhone" placeholder="请输入紧急联系人电话" maxlength="11" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="身高(cm)">
              <el-input-number v-model="patientForm.height" :min="0" :max="300" :precision="1" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="体重(kg)">
              <el-input-number v-model="patientForm.weight" :min="0" :max="500" :precision="1" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="出生体重(kg)">
              <el-input-number v-model="patientForm.birthWeight" :min="0" :max="10" :precision="2" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="医保类型">
              <el-select v-model="patientForm.medicalInsuranceType" placeholder="请选择医保类型">
                <el-option label="城镇职工医保" :value="1" />
                <el-option label="城镇居民医保" :value="2" />
                <el-option label="新农合" :value="3" />
                <el-option label="商业保险" :value="4" />
                <el-option label="自费" :value="5" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="医保号">
              <el-input v-model="patientForm.medicalInsuranceNo" placeholder="请输入医保号" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <template #footer>
        <el-button @click="cancelEdit">取消</el-button>
        <el-button type="primary" @click="savePatient">{{ patientForm.id ? '更新' : '创建' }}</el-button>
      </template>
    </el-dialog>

    <!-- 患者详情对话框 -->
    <el-dialog v-model="showDetailDialog" title="患者详细信息" width="900px">
      <div v-if="patientDetail" class="patient-detail">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="患者编号">
            {{ patientDetail.patientNo }}
          </el-descriptions-item>
          <el-descriptions-item label="姓名">
            {{ patientDetail.name }}
          </el-descriptions-item>
          <el-descriptions-item label="性别">
            {{ patientDetail.gender === 1 ? '男' : '女' }}
          </el-descriptions-item>
          <el-descriptions-item label="出生日期">
            {{ patientDetail.birthDate }}
          </el-descriptions-item>
          <el-descriptions-item label="年龄">
            {{ calculateAge(patientDetail.birthDate) }}岁
          </el-descriptions-item>
          <el-descriptions-item label="身份证号">
            {{ patientDetail.idCard || '未填写' }}
          </el-descriptions-item>
          <el-descriptions-item label="电话号码">
            {{ patientDetail.phone || '未填写' }}
          </el-descriptions-item>
          <el-descriptions-item label="身高">
            {{ patientDetail.height ? patientDetail.height + 'cm' : '未填写' }}
          </el-descriptions-item>
          <el-descriptions-item label="体重">
            {{ patientDetail.weight ? patientDetail.weight + 'kg' : '未填写' }}
          </el-descriptions-item>
          <el-descriptions-item label="出生体重">
            {{ patientDetail.birthWeight ? patientDetail.birthWeight + 'kg' : '未填写' }}
          </el-descriptions-item>
          <el-descriptions-item label="医保类型">
            {{ getMedicalInsuranceTypeName(patientDetail.medicalInsuranceType) }}
          </el-descriptions-item>
          <el-descriptions-item label="医保号">
            {{ patientDetail.medicalInsuranceNo || '未填写' }}
          </el-descriptions-item>
          <el-descriptions-item label="地址" :span="2">
            {{ patientDetail.address || '未填写' }}
          </el-descriptions-item>
          <el-descriptions-item label="紧急联系人">
            {{ patientDetail.emergencyContact || '未填写' }}
          </el-descriptions-item>
          <el-descriptions-item label="紧急联系电话">
            {{ patientDetail.emergencyPhone || '未填写' }}
          </el-descriptions-item>
          <el-descriptions-item label="创建时间" :span="2">
            {{ formatDateTime(patientDetail.createdAt) }}
          </el-descriptions-item>
        </el-descriptions>
      </div>
      <template #footer>
        <el-button @click="showDetailDialog = false">关闭</el-button>
        <el-button type="primary" @click="editPatientFromDetail">编辑患者</el-button>
      </template>
    </el-dialog>

    <!-- 就诊记录对话框 -->
    <el-dialog v-model="showVisitsDialog" title="就诊记录" width="1000px">
      <div v-if="patientVisits.length > 0">
        <el-table :data="patientVisits" stripe>
          <el-table-column prop="visitDate" label="就诊日期" width="120" />
          <el-table-column prop="department" label="科室" width="120" />
          <el-table-column prop="doctor" label="医生" width="100" />
          <el-table-column prop="chiefComplaint" label="主诉" />
          <el-table-column prop="diagnosis" label="诊断" width="150" />
          <el-table-column label="操作" width="120">
            <template #default="{ row }">
              <el-button size="small" @click="viewVisitDetail(row)">查看详情</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div v-else class="empty-state">
        <el-empty description="暂无就诊记录" />
      </div>
      <template #footer>
        <el-button @click="showVisitsDialog = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { clinicalApi } from '@/api/clinical'

// 数据
const patientList = ref([])
const loading = ref(false)
const showCreateDialog = ref(false)
const showDetailDialog = ref(false)
const showVisitsDialog = ref(false)
const patientDetail = ref(null)
const patientVisits = ref([])

// 搜索表单
const searchForm = reactive({
  patientNo: '',
  name: '',
  idCard: '',
  phone: ''
})

// 分页
const pagination = reactive({
  currentPage: 1,
  pageSize: 20,
  total: 0
})

// 患者表单
const patientForm = reactive({
  id: 0,
  patientNo: '',
  name: '',
  gender: null,
  birthDate: '',
  idCard: '',
  phone: '',
  address: '',
  emergencyContact: '',
  emergencyPhone: '',
  height: null,
  weight: null,
  birthWeight: null,
  medicalInsuranceType: null,
  medicalInsuranceNo: ''
})

const patientRules = {
  patientNo: [{ required: true, message: '请输入患者编号', trigger: 'blur' }],
  name: [{ required: true, message: '请输入患者姓名', trigger: 'blur' }],
  gender: [{ required: true, message: '请选择性别', trigger: 'change' }],
  birthDate: [{ required: true, message: '请选择出生日期', trigger: 'change' }]
}

const patientFormRef = ref()

// 方法
const loadPatients = async () => {
  loading.value = true
  try {
    const params = {
      offset: (pagination.currentPage - 1) * pagination.pageSize,
      limit: pagination.pageSize,
      ...searchForm
    }
    
    // 过滤掉空值
    Object.keys(params).forEach(key => {
      if (params[key] === '' || params[key] === null) {
        delete params[key]
      }
    })
    
    const response = await clinicalApi.listPatients(params)
    patientList.value = response.data.patients || []
    pagination.total = response.data.total || 0
  } catch (error) {
    ElMessage.error('加载患者列表失败')
    console.error('加载患者列表失败:', error)
  } finally {
    loading.value = false
  }
}

const searchPatients = () => {
  pagination.currentPage = 1
  loadPatients()
}

const resetSearch = () => {
  Object.assign(searchForm, {
    patientNo: '',
    name: '',
    idCard: '',
    phone: ''
  })
  pagination.currentPage = 1
  loadPatients()
}

const handleSizeChange = (val) => {
  pagination.pageSize = val
  pagination.currentPage = 1
  loadPatients()
}

const handleCurrentChange = (val) => {
  pagination.currentPage = val
  loadPatients()
}

const viewPatient = async (patient) => {
  try {
    const response = await clinicalApi.getPatient(patient.id)
    patientDetail.value = response.data.patient
    showDetailDialog.value = true
    ElMessage.success('获取患者信息成功')
  } catch (error) {
    ElMessage.error('获取患者信息失败')
    console.error('获取患者信息失败:', error)
  }
}

const editPatient = (patient) => {
  // 填充编辑表单
  Object.assign(patientForm, {
    id: patient.id,
    patientNo: patient.patientNo,
    name: patient.name,
    gender: patient.gender,
    birthDate: patient.birthDate,
    idCard: patient.idCard || '',
    phone: patient.phone || '',
    address: patient.address || '',
    emergencyContact: patient.emergencyContact || '',
    emergencyPhone: patient.emergencyPhone || '',
    height: patient.height || null,
    weight: patient.weight || null,
    birthWeight: patient.birthWeight || null,
    medicalInsuranceType: patient.medicalInsuranceType || null,
    medicalInsuranceNo: patient.medicalInsuranceNo || ''
  })
  showCreateDialog.value = true
}

const editPatientFromDetail = () => {
  showDetailDialog.value = false
  editPatient(patientDetail.value)
}

const viewVisits = async (patient) => {
  try {
    const response = await clinicalApi.getPatientVisits(patient.id)
    patientVisits.value = response.data.visits || []
    showVisitsDialog.value = true
    ElMessage.success(`获取到 ${patientVisits.value.length} 条就诊记录`)
  } catch (error) {
    ElMessage.error('获取就诊记录失败')
    console.error('获取就诊记录失败:', error)
  }
}

const viewVisitDetail = (visit) => {
  // TODO: 实现就诊记录详情查看
  ElMessage.info('就诊记录详情功能待实现')
}

const cancelEdit = () => {
  showCreateDialog.value = false
}

const savePatient = async () => {
  if (!patientFormRef.value) return
  
  try {
    await patientFormRef.value.validate()
    
    const patientData = {
      patient: {
        id: patientForm.id || 0,
        patientNo: patientForm.patientNo,
        name: patientForm.name,
        gender: Number(patientForm.gender),
        birthDate: patientForm.birthDate,
        idCard: patientForm.idCard,
        phone: patientForm.phone,
        address: patientForm.address,
        emergencyContact: patientForm.emergencyContact,
        emergencyPhone: patientForm.emergencyPhone,
        height: patientForm.height || 0,
        weight: patientForm.weight || 0,
        birthWeight: patientForm.birthWeight || 0,
        medicalInsuranceType: patientForm.medicalInsuranceType,
        medicalInsuranceNo: patientForm.medicalInsuranceNo
      }
    }
    
    if (patientForm.id) {
      // 更新患者
      await clinicalApi.updatePatient(patientData)
      ElMessage.success('更新患者成功')
    } else {
      // 创建患者
      await clinicalApi.createPatient(patientData)
      ElMessage.success('创建患者成功')
    }
    
    showCreateDialog.value = false
    
    // 重置表单
    Object.assign(patientForm, {
      id: 0,
      patientNo: '',
      name: '',
      gender: null,
      birthDate: '',
      idCard: '',
      phone: '',
      address: '',
      emergencyContact: '',
      emergencyPhone: '',
      height: null,
      weight: null,
      birthWeight: null,
      medicalInsuranceType: null,
      medicalInsuranceNo: ''
    })
    patientFormRef.value?.resetFields()
    
    // 重新加载患者列表
    loadPatients()
  } catch (error) {
    if (error !== false) {
      ElMessage.error(patientForm.id ? '更新患者失败' : '创建患者失败')
      console.error(patientForm.id ? '更新患者失败:' : '创建患者失败:', error)
    }
  }
}

// 工具函数
const calculateAge = (birthDate) => {
  if (!birthDate) return 0
  const today = new Date()
  const birth = new Date(birthDate)
  let age = today.getFullYear() - birth.getFullYear()
  const monthDiff = today.getMonth() - birth.getMonth()
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birth.getDate())) {
    age--
  }
  return age
}

const getMedicalInsuranceTypeName = (type) => {
  const typeMap = {
    1: '城镇职工医保',
    2: '城镇居民医保',
    3: '新农合',
    4: '商业保险',
    5: '自费'
  }
  return typeMap[type] || '未填写'
}

const formatDateTime = (dateTime) => {
  if (!dateTime) return '未知'
  return new Date(dateTime).toLocaleString('zh-CN')
}

onMounted(() => {
  loadPatients()
})
</script>

<style scoped>
.patient-list {
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

.pagination {
  margin-top: 20px;
  text-align: right;
}

/* 操作按钮样式 */
.action-buttons {
  display: inline-flex;
  gap: 6px;
  align-items: center;
}

.action-buttons .el-button {
  margin: 0;
}

/* 患者详情样式 */
.patient-detail {
  margin-bottom: 20px;
}

.empty-state {
  text-align: center;
  padding: 40px 0;
}
</style> 