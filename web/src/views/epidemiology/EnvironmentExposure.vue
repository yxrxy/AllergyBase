<template>
  <div class="environment-exposure">
    <div class="header">
      <h2>环境暴露调查管理</h2>
      <el-button type="primary" @click="showCreateDialog = true">
        <el-icon><Plus /></el-icon>
        新增环境暴露调查
      </el-button>
    </div>

    <!-- 搜索区域 -->
    <div class="search-area">
      <el-form :model="searchForm" inline>
        <el-form-item label="患者ID">
          <el-input v-model="searchForm.patientId" placeholder="请输入患者ID" />
        </el-form-item>
        <el-form-item label="居住类型">
          <el-select v-model="searchForm.residenceType" placeholder="请选择居住类型">
            <el-option label="公寓" value="apartment" />
            <el-option label="独栋房屋" value="house" />
            <el-option label="其他" value="other" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadData">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 环境暴露调查列表 -->
    <el-table :data="tableData" v-loading="loading">
      <el-table-column prop="id" label="调查ID" />
      <el-table-column prop="patientId" label="患者ID" />
      <el-table-column prop="residenceType" label="居住类型" />
      <el-table-column prop="buildingMaterial" label="建筑材料" />
      <el-table-column prop="petExposure" label="宠物暴露">
        <template #default="{ row }">
          <el-tag :type="row.petExposure ? 'success' : 'info'">
            {{ row.petExposure ? '有' : '无' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="smokingExposure" label="吸烟暴露">
        <template #default="{ row }">
          <el-tag :type="row.smokingExposure ? 'danger' : 'success'">
            {{ row.smokingExposure ? '有' : '无' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="createdAt" label="创建时间" />
      <el-table-column label="操作" width="200">
        <template #default="{ row }">
          <el-button size="small" @click="viewExposure(row)">查看</el-button>
          <el-button size="small" type="primary" @click="editExposure(row)">编辑</el-button>
          <el-button size="small" type="danger" @click="deleteExposure(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 创建/编辑环境暴露调查对话框 -->
    <el-dialog v-model="showCreateDialog" :title="isEdit ? '编辑环境暴露调查' : '新增环境暴露调查'" width="600px">
      <el-form :model="exposureForm" :rules="exposureRules" ref="exposureFormRef" label-width="100px">
        <el-form-item label="患者ID" prop="patientId">
          <el-input v-model="exposureForm.patientId" type="number" />
        </el-form-item>
        <el-form-item label="居住类型" prop="residenceType">
          <el-select v-model="exposureForm.residenceType" placeholder="请选择居住类型">
            <el-option label="公寓" value="apartment" />
            <el-option label="独栋房屋" value="house" />
            <el-option label="其他" value="other" />
          </el-select>
        </el-form-item>
        <el-form-item label="建筑材料">
          <el-input v-model="exposureForm.buildingMaterial" />
        </el-form-item>
        <el-form-item label="宠物暴露">
          <el-switch v-model="exposureForm.petExposure" />
        </el-form-item>
        <el-form-item label="吸烟暴露">
          <el-switch v-model="exposureForm.smokingExposure" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showCreateDialog = false">取消</el-button>
        <el-button type="primary" @click="saveExposure">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { epidemiologyApi } from '@/api/epidemiology'

// 数据
const tableData = ref([])
const loading = ref(false)
const showCreateDialog = ref(false)
const isEdit = ref(false)

// 搜索表单
const searchForm = reactive({
  patientId: '',
  residenceType: ''
})

// 环境暴露调查表单
const exposureForm = reactive({
  id: 0,
  patientId: '',
  residenceType: '',
  buildingMaterial: '',
  petExposure: false,
  smokingExposure: false
})

const exposureRules = {
  patientId: [{ required: true, message: '请输入患者ID', trigger: 'blur' }],
  residenceType: [{ required: true, message: '请选择居住类型', trigger: 'change' }]
}

const exposureFormRef = ref()

// 方法
const loadData = async () => {
  loading.value = true
  try {
    tableData.value = [
      {
        id: 1,
        patientId: 'P001',
        residenceType: '公寓',
        buildingMaterial: '钢筋混凝土',
        petExposure: true,
        smokingExposure: false,
        createdAt: '2024-06-01 10:00:00'
      },
      {
        id: 2,
        patientId: 'P002',
        residenceType: '独栋房屋',
        buildingMaterial: '砖混结构',
        petExposure: false,
        smokingExposure: true,
        createdAt: '2024-06-02 11:00:00'
      },
      {
        id: 3,
        patientId: 'P003',
        residenceType: '公寓',
        buildingMaterial: '钢结构',
        petExposure: true,
        smokingExposure: false,
        createdAt: '2024-06-03 09:30:00'
      }
    ];
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.patientId = ''
  searchForm.residenceType = ''
  loadData()
}

const resetForm = () => {
  exposureForm.id = 0
  exposureForm.patientId = ''
  exposureForm.residenceType = ''
  exposureForm.buildingMaterial = ''
  exposureForm.petExposure = false
  exposureForm.smokingExposure = false
}

const createExposure = async () => {
  if (!exposureFormRef.value) return
  
  try {
    await exposureFormRef.value.validate()
    
    const exposureData = {
      ...exposureForm,
      patientId: parseInt(exposureForm.patientId)
    }
    
    await epidemiologyApi.createEnvironmentExposure(exposureData)
    ElMessage.success('创建环境暴露调查成功')
    showCreateDialog.value = false
    loadData()
  } catch (error) {
    if (error !== false) {
      ElMessage.error('创建环境暴露调查失败')
    }
  }
}

const editExposure = (exposure) => {
  isEdit.value = true
  exposureForm.id = exposure.id
  exposureForm.patientId = exposure.patientId
  exposureForm.residenceType = exposure.residenceType || ''
  exposureForm.buildingMaterial = exposure.buildingMaterial || ''
  exposureForm.petExposure = exposure.petExposure || false
  exposureForm.smokingExposure = exposure.smokingExposure || false
  showCreateDialog.value = true
}

const saveExposure = async () => {
  if (isEdit.value) {
    // 编辑功能暂时显示提示
    ElMessage.info('编辑功能暂未实现')
  } else {
    await createExposure()
  }
}

const viewExposure = (exposure) => {
  console.log('查看环境暴露调查:', exposure)
  ElMessage.info('查看功能暂未实现')
}

const deleteExposure = async (exposure) => {
  try {
    await ElMessageBox.confirm('确定要删除这条环境暴露调查吗？', '提示', {
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
  loadData()
})
</script>

<style scoped>
.environment-exposure {
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