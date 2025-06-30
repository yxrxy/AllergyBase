<template>
  <div class="data-quality-report">
    <el-card class="header-card">
      <div class="header-content">
        <h2>数据质量报告</h2>
        <div class="header-actions">
          <el-input
            v-model="patientId"
            placeholder="请输入患者ID"
            style="width: 200px; margin-right: 10px"
          />
          <el-button type="primary" @click="checkSinglePatient">检查单个患者</el-button>
          <el-button type="success" @click="showBatchDialog">批量检查</el-button>
        </div>
      </div>
    </el-card>

    <!-- 质量报告展示 -->
    <div v-if="qualityReport" class="report-content">
      <!-- 总体评分 -->
      <el-card class="score-card">
        <template #header>
          <span>数据质量总体评分</span>
        </template>
        <div class="score-container">
          <div class="score-circle">
            <el-progress
              type="circle"
              :percentage="qualityReport.overall_score"
              :color="getScoreColor(qualityReport.overall_score)"
              :width="120"
            >
              <template #default="{ percentage }">
                <span class="score-text">{{ percentage }}分</span>
              </template>
            </el-progress>
          </div>
          <div class="score-details">
            <div class="score-item">
              <span class="label">临床数据质量:</span>
              <el-progress
                :percentage="qualityReport.clinical_score"
                :color="getScoreColor(qualityReport.clinical_score)"
                :show-text="false"
                style="width: 200px"
              />
              <span class="score">{{ qualityReport.clinical_score }}分</span>
            </div>
            <div class="score-item">
              <span class="label">流调数据质量:</span>
              <el-progress
                :percentage="qualityReport.epidemiology_score"
                :color="getScoreColor(qualityReport.epidemiology_score)"
                :show-text="false"
                style="width: 200px"
              />
              <span class="score">{{ qualityReport.epidemiology_score }}分</span>
            </div>
            <div class="score-item">
              <span class="label">随访数据质量:</span>
              <el-progress
                :percentage="qualityReport.followup_score"
                :color="getScoreColor(qualityReport.followup_score)"
                :show-text="false"
                style="width: 200px"
              />
              <span class="score">{{ qualityReport.followup_score }}分</span>
            </div>
            <div class="score-item">
              <span class="label">生物样本质量:</span>
              <el-progress
                :percentage="qualityReport.biobank_score"
                :color="getScoreColor(qualityReport.biobank_score)"
                :show-text="false"
                style="width: 200px"
              />
              <span class="score">{{ qualityReport.biobank_score }}分</span>
            </div>
          </div>
        </div>
      </el-card>

      <!-- 质量问题列表 -->
      <el-card class="issues-card">
        <template #header>
          <div class="card-header">
            <span>数据质量问题</span>
            <el-tag :type="getIssueCountTagType(qualityReport.issues?.length || 0)">
              共 {{ qualityReport.issues?.length || 0 }} 个问题
            </el-tag>
          </div>
        </template>
        <el-table :data="qualityReport.issues" stripe>
          <el-table-column prop="category" label="类别" width="100">
            <template #default="scope">
              <el-tag :type="getCategoryTagType(scope.row.category)" size="small">
                {{ getCategoryText(scope.row.category) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="severity" label="严重程度" width="100">
            <template #default="scope">
              <el-tag :type="getSeverityTagType(scope.row.severity)" size="small">
                {{ getSeverityText(scope.row.severity) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="field" label="字段" width="120" />
          <el-table-column prop="description" label="问题描述" />
          <el-table-column prop="suggestion" label="改进建议" />
        </el-table>
      </el-card>

      <!-- 改进建议 -->
      <el-card class="recommendations-card">
        <template #header>
          <span>改进建议</span>
        </template>
        <el-timeline>
          <el-timeline-item
            v-for="(recommendation, index) in qualityReport.recommendations"
            :key="index"
            :icon="getRecommendationIcon(index)"
            :type="getRecommendationColor(index)"
          >
            {{ recommendation }}
          </el-timeline-item>
        </el-timeline>
      </el-card>
    </div>

    <!-- 批量检查对话框 -->
    <el-dialog
      v-model="batchDialogVisible"
      title="批量数据质量检查"
      width="600px"
    >
      <div class="batch-content">
        <el-input
          v-model="batchPatientIds"
          type="textarea"
          :rows="6"
          placeholder="请输入患者ID，每行一个"
        />
        <div class="batch-actions">
          <el-button @click="batchDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="checkBatchPatients" :loading="batchLoading">
            开始检查
          </el-button>
        </div>
      </div>
    </el-dialog>

    <!-- 批量结果对话框 -->
    <el-dialog
      v-model="batchResultVisible"
      title="批量检查结果"
      width="80%"
    >
      <el-table :data="batchResults" stripe>
        <el-table-column prop="patient_id" label="患者ID" width="100" />
        <el-table-column prop="overall_score" label="总体评分" width="100">
          <template #default="scope">
            <el-progress
              type="circle"
              :percentage="scope.row.overall_score"
              :color="getScoreColor(scope.row.overall_score)"
              :width="50"
            />
          </template>
        </el-table-column>
        <el-table-column prop="clinical_score" label="临床" width="80" />
        <el-table-column prop="epidemiology_score" label="流调" width="80" />
        <el-table-column prop="followup_score" label="随访" width="80" />
        <el-table-column prop="biobank_score" label="样本" width="80" />
        <el-table-column label="问题数量" width="100">
          <template #default="scope">
            <el-tag :type="getIssueCountTagType(scope.row.issues?.length || 0)">
              {{ scope.row.issues?.length || 0 }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100">
          <template #default="scope">
            <el-button type="text" @click="viewPatientReport(scope.row)">
              查看详情
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <div v-if="loading" class="loading-container">
      <el-loading text="检查中..." />
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { integrationApi } from '@/api/integration'

export default {
  name: 'DataQualityReport',
  setup() {
    const patientId = ref('')
    const qualityReport = ref(null)
    const loading = ref(false)
    const batchDialogVisible = ref(false)
    const batchResultVisible = ref(false)
    const batchPatientIds = ref('')
    const batchLoading = ref(false)
    const batchResults = ref([])

    const checkSinglePatient = async () => {
      if (!patientId.value) {
        ElMessage.warning('请输入患者ID')
        return
      }

      loading.value = true
      try {
        const response = await integrationApi.checkPatientDataQuality(patientId.value)
        qualityReport.value = response.data
        ElMessage.success('数据质量检查完成')
      } catch (error) {
        ElMessage.error('检查失败: ' + error.message)
      } finally {
        loading.value = false
      }
    }

    const showBatchDialog = () => {
      batchDialogVisible.value = true
      batchPatientIds.value = ''
    }

    const checkBatchPatients = async () => {
      if (!batchPatientIds.value.trim()) {
        ElMessage.warning('请输入患者ID')
        return
      }

      const ids = batchPatientIds.value.split('\n')
        .map(id => id.trim())
        .filter(id => id)
        .map(id => parseInt(id))
        .filter(id => !isNaN(id))

      if (ids.length === 0) {
        ElMessage.warning('请输入有效的患者ID')
        return
      }

      batchLoading.value = true
      try {
        const response = await integrationApi.batchCheckDataQuality(ids)
        batchResults.value = response.data
        batchDialogVisible.value = false
        batchResultVisible.value = true
        ElMessage.success(`批量检查完成，共检查 ${ids.length} 个患者`)
      } catch (error) {
        ElMessage.error('批量检查失败: ' + error.message)
      } finally {
        batchLoading.value = false
      }
    }

    const viewPatientReport = (report) => {
      qualityReport.value = report
      batchResultVisible.value = false
    }

    const getScoreColor = (score) => {
      if (score >= 90) return '#67c23a'
      if (score >= 70) return '#e6a23c'
      return '#f56c6c'
    }

    const getCategoryText = (category) => {
      const map = {
        clinical: '临床',
        epidemiology: '流调',
        followup: '随访',
        biobank: '样本'
      }
      return map[category] || category
    }

    const getCategoryTagType = (category) => {
      const map = {
        clinical: 'primary',
        epidemiology: 'success',
        followup: 'warning',
        biobank: 'info'
      }
      return map[category] || ''
    }

    const getSeverityText = (severity) => {
      const map = {
        low: '低',
        medium: '中',
        high: '高',
        critical: '严重'
      }
      return map[severity] || severity
    }

    const getSeverityTagType = (severity) => {
      const map = {
        low: 'info',
        medium: 'warning',
        high: 'danger',
        critical: 'danger'
      }
      return map[severity] || ''
    }

    const getIssueCountTagType = (count) => {
      if (count === 0) return 'success'
      if (count <= 3) return 'warning'
      return 'danger'
    }

    const getRecommendationIcon = (index) => {
      return index === 0 ? 'el-icon-star-on' : 'el-icon-check'
    }

    const getRecommendationColor = (index) => {
      return index === 0 ? 'primary' : 'success'
    }

    return {
      patientId,
      qualityReport,
      loading,
      batchDialogVisible,
      batchResultVisible,
      batchPatientIds,
      batchLoading,
      batchResults,
      checkSinglePatient,
      showBatchDialog,
      checkBatchPatients,
      viewPatientReport,
      getScoreColor,
      getCategoryText,
      getCategoryTagType,
      getSeverityText,
      getSeverityTagType,
      getIssueCountTagType,
      getRecommendationIcon,
      getRecommendationColor
    }
  }
}
</script>

<style scoped>
.data-quality-report {
  padding: 20px;
}

.header-card {
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
  align-items: center;
}

.report-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.score-card, .issues-card, .recommendations-card {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.score-container {
  display: flex;
  align-items: center;
  gap: 40px;
}

.score-circle {
  flex-shrink: 0;
}

.score-text {
  font-size: 16px;
  font-weight: bold;
}

.score-details {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.score-item {
  display: flex;
  align-items: center;
  gap: 15px;
}

.score-item .label {
  width: 120px;
  color: #606266;
  font-size: 14px;
}

.score-item .score {
  width: 50px;
  text-align: right;
  font-weight: bold;
  color: #303133;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.batch-content {
  padding: 20px 0;
}

.batch-actions {
  margin-top: 20px;
  text-align: right;
}

.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 200px;
}

.el-timeline {
  padding-left: 20px;
}
</style> 