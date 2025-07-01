<template>
  <div class="correlation-analysis">
    <el-card class="header-card">
      <div class="header-content">
        <h2>关联性分析</h2>
        <div class="header-actions">
          <el-button type="primary" @click="showAnalysisDialog">新建分析</el-button>
          <el-button type="success" @click="loadAnalysisList">刷新列表</el-button>
        </div>
      </div>
    </el-card>

    <!-- 分析列表 -->
    <el-card class="list-card">
      <template #header>
        <span>分析历史</span>
      </template>
      <el-table :data="analysisList" stripe v-loading="listLoading">
        <el-table-column prop="id" label="分析ID" width="100" />
        <el-table-column prop="analysis_type" label="分析类型" width="120">
          <template #default="scope">
            <el-tag :type="getAnalysisTypeTag(scope.row.analysis_type)">
              {{ getAnalysisTypeName(scope.row.analysis_type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTag(scope.row.status)">
              {{ getStatusName(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column prop="progress" label="进度" width="120">
          <template #default="scope">
            <el-progress
              :percentage="scope.row.progress"
              :status="scope.row.status === 'failed' ? 'exception' : 
                      scope.row.status === 'completed' ? 'success' : ''"
            />
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" />
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button
              type="text"
              @click="viewAnalysisResult(scope.row)"
              :disabled="scope.row.status !== 'completed'"
            >
              查看结果
            </el-button>
            <el-button
              type="text"
              @click="deleteAnalysis(scope.row.id)"
              style="color: #f56c6c"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 分析结果展示 -->
    <div v-if="analysisResult" class="result-content">
      <el-card class="result-card">
        <template #header>
          <div class="result-header">
            <span>分析结果 - {{ getAnalysisTypeName(analysisResult.analysis_type) }}</span>
            <el-button type="primary" @click="exportResult">导出结果</el-button>
          </div>
        </template>
        
        <!-- 统计概览 -->
        <div class="result-overview">
          <el-row :gutter="20">
            <el-col :span="6">
              <el-statistic title="样本数量" :value="analysisResult.sample_count" />
            </el-col>
            <el-col :span="6">
              <el-statistic title="变量数量" :value="analysisResult.variable_count" />
            </el-col>
            <el-col :span="6">
              <el-statistic title="显著关联" :value="analysisResult.significant_count" />
            </el-col>
            <el-col :span="6">
              <el-statistic title="置信度" :value="analysisResult.confidence_level" suffix="%" />
            </el-col>
          </el-row>
        </div>

        <!-- 关联性矩阵 -->
        <div v-if="analysisResult.correlation_matrix" class="correlation-matrix">
          <h3>关联性矩阵</h3>
          <div class="matrix-container">
            <table class="matrix-table">
              <thead>
                <tr>
                  <th></th>
                  <th v-for="variable in analysisResult.variables" :key="variable">
                    {{ variable }}
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(row, rowIndex) in analysisResult.correlation_matrix" :key="rowIndex">
                  <th>{{ analysisResult.variables[rowIndex] }}</th>
                  <td
                    v-for="(value, colIndex) in row"
                    :key="colIndex"
                    :class="getCorrelationClass(value)"
                    :title="`相关系数: ${value.toFixed(4)}`"
                  >
                    {{ value.toFixed(3) }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- 显著关联列表 -->
        <div v-if="analysisResult.significant_correlations" class="significant-correlations">
          <h3>显著关联</h3>
          <el-table :data="analysisResult.significant_correlations" stripe>
            <el-table-column prop="variable1" label="变量1" width="150" />
            <el-table-column prop="variable2" label="变量2" width="150" />
            <el-table-column prop="correlation" label="相关系数" width="120">
              <template #default="scope">
                <span :class="getCorrelationClass(scope.row.correlation)">
                  {{ scope.row.correlation.toFixed(4) }}
                </span>
              </template>
            </el-table-column>
            <el-table-column prop="p_value" label="P值" width="120">
              <template #default="scope">
                {{ scope.row.p_value.toFixed(6) }}
              </template>
            </el-table-column>
            <el-table-column prop="strength" label="关联强度" width="100">
              <template #default="scope">
                <el-tag :type="getStrengthTag(scope.row.strength)">
                  {{ getStrengthName(scope.row.strength) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="interpretation" label="解释" />
          </el-table>
        </div>

        <!-- 分析结论 -->
        <div v-if="analysisResult.conclusions" class="conclusions">
          <h3>分析结论</h3>
          <el-timeline>
            <el-timeline-item
              v-for="(conclusion, index) in analysisResult.conclusions"
              :key="index"
              :icon="getConclusionIcon(conclusion.type)"
              :type="getConclusionColor(conclusion.type)"
            >
              <h4>{{ conclusion.title }}</h4>
              <p>{{ conclusion.description }}</p>
            </el-timeline-item>
          </el-timeline>
        </div>
      </el-card>
    </div>

    <!-- 新建分析对话框 -->
    <el-dialog
      v-model="analysisDialogVisible"
      title="新建关联性分析"
      width="600px"
    >
      <el-form :model="analysisForm" :rules="analysisRules" ref="analysisFormRef" label-width="120px">
        <el-form-item label="分析类型" prop="analysis_type">
          <el-select v-model="analysisForm.analysis_type" placeholder="请选择分析类型">
            <el-option label="临床-流调关联" value="clinical_epidemiology" />
            <el-option label="临床-随访关联" value="clinical_followup" />
            <el-option label="流调-生物样本关联" value="epidemiology_biobank" />
            <el-option label="全数据库关联" value="cross_database" />
            <el-option label="时间序列关联" value="temporal" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="分析描述" prop="description">
          <el-input
            v-model="analysisForm.description"
            type="textarea"
            :rows="3"
            placeholder="请描述分析目的和预期结果"
          />
        </el-form-item>

        <el-form-item label="目标变量" prop="target_variables">
          <el-input
            v-model="analysisForm.target_variables"
            placeholder="请输入目标变量，用逗号分隔"
          />
        </el-form-item>

        <el-form-item label="控制变量" prop="control_variables">
          <el-input
            v-model="analysisForm.control_variables"
            placeholder="请输入控制变量，用逗号分隔（可选）"
          />
        </el-form-item>

        <el-form-item label="置信水平">
          <el-slider
            v-model="analysisForm.confidence_level"
            :min="90"
            :max="99"
            :step="1"
            show-stops
            show-input
          />
        </el-form-item>

        <el-form-item label="最小样本量">
          <el-input-number
            v-model="analysisForm.min_sample_size"
            :min="10"
            :max="10000"
            :step="10"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="analysisDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="startAnalysis" :loading="analysisLoading">
            开始分析
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { integrationApi } from '@/api/integration'

export default {
  name: 'CorrelationAnalysis',
  setup() {
    const analysisList = ref([])
    const analysisResult = ref(null)
    const listLoading = ref(false)
    const analysisDialogVisible = ref(false)
    const analysisLoading = ref(false)
    const analysisFormRef = ref(null)

    const analysisForm = reactive({
      analysis_type: '',
      description: '',
      target_variables: '',
      control_variables: '',
      confidence_level: 95,
      min_sample_size: 100
    })

    const analysisRules = reactive({
      analysis_type: [
        { required: true, message: '请选择分析类型', trigger: 'change' }
      ],
      description: [
        { required: true, message: '请输入分析描述', trigger: 'blur' }
      ],
      target_variables: [
        { required: true, message: '请输入目标变量', trigger: 'blur' }
      ]
    })

    const loadAnalysisList = async () => {
      listLoading.value = true
      try {
        // 模拟获取分析列表
        analysisList.value = [
          {
            id: 1,
            analysis_type: 'clinical_epidemiology',
            status: 'completed',
            progress: 100,
            created_at: new Date(),
            description: '临床数据与流调数据关联性分析'
          },
          {
            id: 2,
            analysis_type: 'cross_database',
            status: 'running',
            progress: 65,
            created_at: new Date(),
            description: '全数据库关联性分析'
          }
        ]
      } catch (error) {
        ElMessage.error('加载分析列表失败: ' + error.message)
      } finally {
        listLoading.value = false
      }
    }

    const showAnalysisDialog = () => {
      analysisDialogVisible.value = true
      Object.assign(analysisForm, {
        analysis_type: '',
        description: '',
        target_variables: '',
        control_variables: '',
        confidence_level: 95,
        min_sample_size: 100
      })
    }

    const startAnalysis = async () => {
      if (!analysisFormRef.value) return
      
      try {
        await analysisFormRef.value.validate()
        
        analysisLoading.value = true
        
        const parameters = {
          target_variables: analysisForm.target_variables.split(',').map(v => v.trim()),
          control_variables: analysisForm.control_variables ? 
            analysisForm.control_variables.split(',').map(v => v.trim()) : [],
          confidence_level: analysisForm.confidence_level,
          min_sample_size: analysisForm.min_sample_size,
          description: analysisForm.description
        }

        await integrationApi.performCorrelationAnalysis(
          analysisForm.analysis_type,
          parameters
        )
        
        ElMessage.success('分析任务已提交，请稍后查看结果')
        analysisDialogVisible.value = false
        loadAnalysisList()
      } catch (error) {
        ElMessage.error('提交分析失败: ' + error.message)
      } finally {
        analysisLoading.value = false
      }
    }

    const viewAnalysisResult = async (analysis) => {
      try {
        const response = await integrationApi.getCorrelationAnalysis(analysis.id)
        analysisResult.value = response.data
      } catch (error) {
        ElMessage.error('获取分析结果失败: ' + error.message)
      }
    }

    const deleteAnalysis = async (analysisId) => {
      try {
        await ElMessageBox.confirm('确定要删除这个分析吗？', '确认删除', {
          type: 'warning'
        })
        
        // 这里应该调用删除API
        ElMessage.success('分析已删除')
        loadAnalysisList()
      } catch (error) {
        if (error !== 'cancel') {
          ElMessage.error('删除失败: ' + error.message)
        }
      }
    }

    const exportResult = () => {
      if (!analysisResult.value) return
      
      // 模拟导出功能
      ElMessage.success('结果导出中...')
    }

    // 辅助函数
    const getAnalysisTypeName = (type) => {
      const map = {
        clinical_epidemiology: '临床-流调',
        clinical_followup: '临床-随访',
        epidemiology_biobank: '流调-样本',
        cross_database: '全数据库',
        temporal: '时间序列'
      }
      return map[type] || type
    }

    const getAnalysisTypeTag = (type) => {
      const map = {
        clinical_epidemiology: 'primary',
        clinical_followup: 'success',
        epidemiology_biobank: 'warning',
        cross_database: 'info',
        temporal: 'danger'
      }
      return map[type] || ''
    }

    const getStatusName = (status) => {
      const map = {
        pending: '等待中',
        running: '运行中',
        completed: '已完成',
        failed: '失败'
      }
      return map[status] || status
    }

    const getStatusTag = (status) => {
      const map = {
        pending: 'info',
        running: 'warning',
        completed: 'success',
        failed: 'danger'
      }
      return map[status] || ''
    }

    const getCorrelationClass = (value) => {
      const abs = Math.abs(value)
      if (abs >= 0.7) return 'strong-correlation'
      if (abs >= 0.3) return 'moderate-correlation'
      return 'weak-correlation'
    }

    const getStrengthName = (strength) => {
      const map = {
        weak: '弱',
        moderate: '中等',
        strong: '强'
      }
      return map[strength] || strength
    }

    const getStrengthTag = (strength) => {
      const map = {
        weak: 'info',
        moderate: 'warning',
        strong: 'danger'
      }
      return map[strength] || ''
    }

    const getConclusionIcon = (type) => {
      const map = {
        finding: 'el-icon-search',
        recommendation: 'el-icon-star-on',
        warning: 'el-icon-warning'
      }
      return map[type] || 'el-icon-info'
    }

    const getConclusionColor = (type) => {
      const map = {
        finding: 'primary',
        recommendation: 'success',
        warning: 'warning'
      }
      return map[type] || 'info'
    }

    const formatDate = (date) => {
      return new Date(date).toLocaleString('zh-CN')
    }

    onMounted(() => {
      loadAnalysisList()
    })

    return {
      analysisList,
      analysisResult,
      listLoading,
      analysisDialogVisible,
      analysisLoading,
      analysisForm,
      analysisRules,
      analysisFormRef,
      loadAnalysisList,
      showAnalysisDialog,
      startAnalysis,
      viewAnalysisResult,
      deleteAnalysis,
      exportResult,
      getAnalysisTypeName,
      getAnalysisTypeTag,
      getStatusName,
      getStatusTag,
      getCorrelationClass,
      getStrengthName,
      getStrengthTag,
      getConclusionIcon,
      getConclusionColor,
      formatDate
    }
  }
}
</script>

<style scoped>
.correlation-analysis {
  padding: 20px;
}

.header-card, .list-card {
  margin-bottom: 20px;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header-content h2 {
  margin: 0;
  color: #303133;
}

.header-actions {
  display: flex;
  gap: 10px;
}

.result-content {
  margin-top: 20px;
}

.result-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.result-overview {
  margin-bottom: 30px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
}

.correlation-matrix {
  margin-bottom: 30px;
}

.matrix-container {
  overflow-x: auto;
  margin-top: 15px;
}

.matrix-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 12px;
}

.matrix-table th,
.matrix-table td {
  border: 1px solid #dcdfe6;
  padding: 8px;
  text-align: center;
  min-width: 80px;
}

.matrix-table th {
  background: #f5f7fa;
  font-weight: bold;
}

.strong-correlation {
  background: #f56c6c;
  color: white;
  font-weight: bold;
}

.moderate-correlation {
  background: #e6a23c;
  color: white;
}

.weak-correlation {
  background: #909399;
  color: white;
}

.significant-correlations,
.conclusions {
  margin-bottom: 30px;
}

.significant-correlations h3,
.conclusions h3 {
  margin-bottom: 15px;
  color: #303133;
}

.conclusions .el-timeline {
  padding-left: 20px;
}

.conclusions h4 {
  margin: 0 0 10px 0;
  color: #606266;
}

.conclusions p {
  margin: 0;
  color: #909399;
}

.dialog-footer {
  text-align: right;
}
</style> 