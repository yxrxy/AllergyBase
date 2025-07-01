<template>
  <div class="storage-list">
    <div class="header">
      <h2>存储信息管理</h2>
      <el-button type="primary" @click="showCreateDialog = true">
        <el-icon><Plus /></el-icon>
        新增存储信息
      </el-button>
    </div>

    <!-- 搜索区域 -->
    <div class="search-area">
      <el-form :model="searchForm" inline>
        <el-form-item label="样本ID">
          <el-input v-model="searchForm.sampleId" placeholder="请输入样本ID" />
        </el-form-item>
        <el-form-item label="存储位置">
          <el-input v-model="searchForm.storageLocation" placeholder="请输入存储位置" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadStorage">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 存储信息列表 -->
    <el-table :data="storageList" v-loading="loading">
      <el-table-column prop="id" label="存储ID" />
      <el-table-column prop="sampleId" label="样本ID" />
      <el-table-column prop="storageLocation" label="存储位置" />
      <el-table-column prop="storageTemperature" label="存储温度" />
      <el-table-column prop="storageTime" label="入库时间" />
      <el-table-column prop="status" label="状态" />
      <el-table-column label="操作" width="200">
        <template #default="{ row }">
          <el-button size="small" @click="viewStorage(row)">查看</el-button>
          <el-button size="small" type="primary" @click="editStorage(row)">编辑</el-button>
          <el-button size="small" type="danger" @click="deleteStorage(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 创建/编辑存储信息对话框 -->
    <el-dialog v-model="showCreateDialog" :title="isEdit ? '编辑存储信息' : '新增存储信息'" width="600px">
      <el-form :model="storageForm" :rules="storageRules" ref="storageFormRef" label-width="100px">
        <el-form-item label="样本ID" prop="sampleId">
          <el-input v-model="storageForm.sampleId" type="number" />
        </el-form-item>
        <el-form-item label="存储位置" prop="storageLocation">
          <el-input v-model="storageForm.storageLocation" />
        </el-form-item>
        <el-form-item label="存储温度">
          <el-input v-model="storageForm.storageTemperature" type="number" />
        </el-form-item>
        <el-form-item label="入库时间">
          <el-date-picker v-model="storageForm.storageTime" type="datetime" format="YYYY-MM-DD HH:mm:ss" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="storageForm.status" placeholder="请选择状态">
            <el-option label="正常" value="正常" />
            <el-option label="异常" value="异常" />
            <el-option label="已销毁" value="已销毁" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showCreateDialog = false">取消</el-button>
        <el-button type="primary" @click="saveStorage">确定</el-button>
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
const storageList = ref([])
const loading = ref(false)
const showCreateDialog = ref(false)
const isEdit = ref(false)

// 搜索表单
const searchForm = reactive({
  sampleId: '',
  storageLocation: ''
})

// 存储信息表单
const storageForm = reactive({
  id: 0,
  sampleId: '',
  storageLocation: '',
  storageTemperature: '',
  storageTime: '',
  status: ''
})

const storageRules = {
  sampleId: [{ required: true, message: '请输入样本ID', trigger: 'blur' }],
  storageLocation: [{ required: true, message: '请输入存储位置', trigger: 'blur' }]
}

const storageFormRef = ref()

// 方法
const loadStorage = async () => {
  loading.value = true
  try {
    // 暂时显示所有样本的存储信息（通过循环查询）
    const allStorage = []
    const sampleIds = [1, 2, 3] // 这里可以根据实际需要调整
    
    for (const sampleId of sampleIds) {
      try {
        const resp = await biobankApi.getStorageInfo(sampleId)
        if (resp.data.storage) {
          allStorage.push(resp.data.storage)
        }
      } catch (error) {
        // 忽略没有存储信息的样本
      }
    }
    
    // 根据搜索条件过滤
    let filteredStorage = allStorage
    if (searchForm.sampleId) {
      filteredStorage = filteredStorage.filter(item => 
        item.sampleId.toString().includes(searchForm.sampleId)
      )
    }
    if (searchForm.storageLocation) {
      filteredStorage = filteredStorage.filter(item => 
        item.storageLocation && item.storageLocation.includes(searchForm.storageLocation)
      )
    }
    
    storageList.value = filteredStorage
  } catch (error) {
    ElMessage.error('加载存储信息失败')
    console.error('加载存储信息失败:', error)
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  searchForm.sampleId = ''
  searchForm.storageLocation = ''
  loadStorage()
}

const resetForm = () => {
  storageForm.id = 0
  storageForm.sampleId = ''
  storageForm.storageLocation = ''
  storageForm.storageTemperature = ''
  storageForm.storageTime = ''
  storageForm.status = ''
}

const createStorage = async () => {
  if (!storageFormRef.value) return
  
  try {
    await storageFormRef.value.validate()
    
    const storageData = {
      ...storageForm,
      sampleId: parseInt(storageForm.sampleId),
      storageTemperature: parseFloat(storageForm.storageTemperature) || 0
    }
    
    // 处理时间格式
    if (storageData.storageTime) {
      storageData.storageTime = new Date(storageData.storageTime).toISOString()
    }
    
    await biobankApi.addStorageInfo(storageData)
    ElMessage.success('创建存储信息成功')
    showCreateDialog.value = false
    loadStorage()
  } catch (error) {
    if (error !== false) {
      ElMessage.error('创建存储信息失败')
    }
  }
}

const editStorage = (storage) => {
  isEdit.value = true
  storageForm.id = storage.id
  storageForm.sampleId = storage.sampleId
  storageForm.storageLocation = storage.storageLocation || ''
  storageForm.storageTemperature = storage.storageTemperature || ''
  storageForm.storageTime = storage.storageTime ? new Date(storage.storageTime) : ''
  storageForm.status = storage.status || ''
  showCreateDialog.value = true
}

const saveStorage = async () => {
  if (isEdit.value) {
    // 编辑功能暂时显示提示
    ElMessage.info('编辑功能暂未实现')
  } else {
    await createStorage()
  }
}

const viewStorage = (storage) => {
  console.log('查看存储信息:', storage)
  ElMessage.info('查看功能暂未实现')
}

const deleteStorage = async (storage) => {
  try {
    await ElMessageBox.confirm('确定要删除这条存储信息吗？', '提示', {
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
  loadStorage()
})
</script>

<style scoped>
.storage-list {
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