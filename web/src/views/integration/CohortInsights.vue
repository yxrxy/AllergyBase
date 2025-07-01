<template>
  <div class="cohort-insights">
    <el-card class="header-card">
      <div class="header-content">
        <h2>队列洞察分析</h2>
        <div class="header-actions">
          <el-button type="primary" @click="showQueryDialog">新建查询</el-button>
          <el-button type="success" @click="loadSavedQueries">刷新查询</el-button>
        </div>
      </div>
    </el-card>

    <!-- 快速查询模板 -->
    <el-card class="template-card">
      <template #header>
        <span>快速查询模板</span>
      </template>
      <div class="template-grid">
        <div
          v-for="template in queryTemplates"
          :key="template.id"
          class="template-item"
          @click="useTemplate(template)"
        >
          <div class="template-icon">
            <i :class="template.icon"></i>
          </div>
          <div class="template-content">
            <h4>{{ template.name }}</h4>
            <p>{{ template.description }}</p>
          </div>
        </div>
      </div>
    </el-card>

    <!-- 已保存的查询 -->
    <el-card class="saved-queries-card">
      <template #header>
        <div class="card-header">
          <span>已保存的查询</span>
          <el-tag type="info">{{ savedQueries.length }} 个查询</el-tag>
        </div>
      </template>
      <el-table :data="savedQueries" stripe v-loading="queriesLoading">
        <el-table-column prop="name" label="查询名称" width="200" />
        <el-table-column prop="type" label="查询类型" width="120">
          <template #default="scope">
            <el-tag :type="getQueryTypeTag(scope.row.type)">
              {{ getQueryTypeName(scope.row.type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" />
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button type="text" @click="executeQuery(scope.row)">执行</el-button>
            <el-button type="text" @click="editQuery(scope.row)">编辑</el-button>
            <el-button
              type="text"
              @click="deleteQuery(scope.row.id)"
              style="color: #f56c6c"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 洞察结果展示 -->
    <div v-if="insightResults" class="insights-content">
      <el-card class="insights-card">
        <template #header>
          <div class="insights-header">
            <span>队列洞察结果</span>
            <div class="header-actions">
              <el-button type="primary" @click="exportInsights">导出报告</el-button>
              <el-button @click="saveQuery">保存查询</el-button>
            </div>
          </div>
        </template>

        <!-- 基本统计 -->
        <div class="basic-stats">
          <el-row :gutter="20">
            <el-col :span="6">
              <el-statistic
                title="队列总数"
                :value="insightResults.total_patients"
                suffix="人"
              />
            </el-col>
            <el-col :span="6">
              <el-statistic
                title="平均年龄"
                :value="insightResults.average_age"
                :precision="1"
                suffix="岁"
              />
            </el-col>
            <el-col :span="6">
              <el-statistic
                title="随访完成率"
                :value="insightResults.followup_completion_rate"
                :precision="1"
                suffix="%"
              />
            </el-col>
            <el-col :span="6">
              <el-statistic
                title="样本收集率"
                :value="insightResults.sample_collection_rate"
                :precision="1"
                suffix="%"
              />
            </el-col>
          </el-row>
        </div>

        <!-- 人口统计学特征 -->
        <div class="demographics">
          <h3>人口统计学特征</h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <div class="chart-container">
                <h4>性别分布</h4>
                <div ref="genderChart" class="chart"></div>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="chart-container">
                <h4>年龄分布</h4>
                <div ref="ageChart" class="chart"></div>
              </div>
            </el-col>
          </el-row>
        </div>

        <!-- 疾病特征 -->
        <div class="disease-characteristics">
          <h3>疾病特征分析</h3>
          <el-row :gutter="20">
            <el-col :span="8">
              <el-card class="feature-card">
                <template #header>
                  <span>主要诊断</span>
                </template>
                <el-table :data="insightResults.primary_diagnoses" size="small">
                  <el-table-column prop="diagnosis" label="诊断" />
                  <el-table-column prop="count" label="人数" width="80" />
                  <el-table-column prop="percentage" label="占比" width="80">
                    <template #default="scope">
                      {{ scope.row.percentage }}%
                    </template>
                  </el-table-column>
                </el-table>
              </el-card>
            </el-col>
            <el-col :span="8">
              <el-card class="feature-card">
                <template #header>
                  <span>过敏原分布</span>
                </template>
                <el-table :data="insightResults.allergen_distribution" size="small">
                  <el-table-column prop="allergen" label="过敏原" />
                  <el-table-column prop="count" label="人数" width="80" />
                  <el-table-column prop="severity" label="严重程度" width="100">
                    <template #default="scope">
                      <el-tag :type="getSeverityTag(scope.row.severity)" size="small">
                        {{ getSeverityName(scope.row.severity) }}
                      </el-tag>
                    </template>
                  </el-table-column>
                </el-table>
              </el-card>
            </el-col>
            <el-col :span="8">
              <el-card class="feature-card">
                <template #header>
                  <span>治疗方案</span>
                </template>
                <el-table :data="insightResults.treatment_plans" size="small">
                  <el-table-column prop="treatment" label="治疗方案" />
                  <el-table-column prop="count" label="人数" width="80" />
                  <el-table-column prop="effectiveness" label="有效率" width="80">
                    <template #default="scope">
                      {{ scope.row.effectiveness }}%
                    </template>
                  </el-table-column>
                </el-table>
              </el-card>
            </el-col>
          </el-row>
        </div>

        <!-- 关键洞察 -->
        <div class="key-insights">
          <h3>关键洞察</h3>
          <el-timeline>
            <el-timeline-item
              v-for="(insight, index) in insightResults.key_insights"
              :key="index"
              :icon="getInsightIcon(insight.type)"
              :type="getInsightColor(insight.type)"
            >
              <div class="insight-content">
                <h4>{{ insight.title }}</h4>
                <p>{{ insight.description }}</p>
                <div v-if="insight.metrics" class="insight-metrics">
                  <el-tag
                    v-for="metric in insight.metrics"
                    :key="metric.name"
                    :type="getMetricTag(metric.significance)"
                    style="margin-right: 8px"
                  >
                    {{ metric.name }}: {{ metric.value }}
                  </el-tag>
                </div>
              </div>
            </el-timeline-item>
          </el-timeline>
        </div>

        <!-- 建议和行动项 -->
        <div class="recommendations">
          <h3>建议和行动项</h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-card class="recommendation-card">
                <template #header>
                  <span>临床建议</span>
                </template>
                <ul class="recommendation-list">
                  <li v-for="(rec, index) in insightResults.clinical_recommendations" :key="index">
                    {{ rec }}
                  </li>
                </ul>
              </el-card>
            </el-col>
            <el-col :span="12">
              <el-card class="recommendation-card">
                <template #header>
                  <span>研究建议</span>
                </template>
                <ul class="recommendation-list">
                  <li v-for="(rec, index) in insightResults.research_recommendations" :key="index">
                    {{ rec }}
                  </li>
                </ul>
              </el-card>
            </el-col>
          </el-row>
        </div>
      </el-card>
    </div>

    <!-- 新建查询对话框 -->
    <el-dialog
      v-model="queryDialogVisible"
      title="新建队列查询"
      width="800px"
    >
      <el-form :model="queryForm" :rules="queryRules" ref="queryFormRef" label-width="120px">
        <el-form-item label="查询名称" prop="name">
          <el-input v-model="queryForm.name" placeholder="请输入查询名称" />
        </el-form-item>

        <el-form-item label="查询类型" prop="type">
          <el-select v-model="queryForm.type" placeholder="请选择查询类型">
            <el-option label="基础队列分析" value="basic_cohort" />
            <el-option label="疾病特征分析" value="disease_analysis" />
            <el-option label="治疗效果分析" value="treatment_analysis" />
            <el-option label="随访质量分析" value="followup_analysis" />
            <el-option label="生存分析" value="survival_analysis" />
          </el-select>
        </el-form-item>

        <el-form-item label="包含条件">
          <el-input
            v-model="queryForm.inclusion_criteria"
            type="textarea"
            :rows="3"
            placeholder="请输入包含条件，如：年龄>18，诊断包含过敏性疾病"
          />
        </el-form-item>

        <el-form-item label="排除条件">
          <el-input
            v-model="queryForm.exclusion_criteria"
            type="textarea"
            :rows="3"
            placeholder="请输入排除条件（可选）"
          />
        </el-form-item>

        <el-form-item label="时间范围">
          <el-date-picker
            v-model="queryForm.date_range"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
          />
        </el-form-item>

        <el-form-item label="查询描述">
          <el-input
            v-model="queryForm.description"
            type="textarea"
            :rows="2"
            placeholder="请描述查询目的和预期结果"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="queryDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitQuery" :loading="queryLoading">
            执行查询
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref, reactive, onMounted, nextTick } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { integrationApi } from '@/api/integration'

export default {
  name: 'CohortInsights',
  setup() {
    const savedQueries = ref([])
    const insightResults = ref(null)
    const queriesLoading = ref(false)
    const queryDialogVisible = ref(false)
    const queryLoading = ref(false)
    const queryFormRef = ref(null)

    const queryForm = reactive({
      name: '',
      type: '',
      inclusion_criteria: '',
      exclusion_criteria: '',
      date_range: null,
      description: ''
    })

    const queryRules = reactive({
      name: [
        { required: true, message: '请输入查询名称', trigger: 'blur' }
      ],
      type: [
        { required: true, message: '请选择查询类型', trigger: 'change' }
      ]
    })

    const queryTemplates = ref([
      {
        id: 1,
        name: '过敏性疾病队列',
        description: '分析过敏性疾病患者的基本特征和治疗情况',
        icon: 'el-icon-user',
        type: 'disease_analysis',
        inclusion_criteria: '诊断包含过敏性疾病',
        exclusion_criteria: '年龄<18'
      },
      {
        id: 2,
        name: '治疗效果评估',
        description: '评估不同治疗方案的效果和安全性',
        icon: 'el-icon-medicine-box',
        type: 'treatment_analysis',
        inclusion_criteria: '有完整治疗记录',
        exclusion_criteria: '随访时间<3个月'
      },
      {
        id: 3,
        name: '随访质量分析',
        description: '分析随访完成情况和数据质量',
        icon: 'el-icon-document',
        type: 'followup_analysis',
        inclusion_criteria: '有随访记录',
        exclusion_criteria: ''
      },
      {
        id: 4,
        name: '生存分析',
        description: '分析患者生存情况和影响因素',
        icon: 'el-icon-trend-charts',
        type: 'survival_analysis',
        inclusion_criteria: '有完整随访数据',
        exclusion_criteria: '失访患者'
      }
    ])

    const loadSavedQueries = async () => {
      queriesLoading.value = true
      try {
        // 模拟获取已保存的查询
        savedQueries.value = [
          {
            id: 1,
            name: '过敏性疾病队列分析',
            type: 'disease_analysis',
            created_at: new Date(),
            description: '分析过敏性疾病患者的基本特征'
          },
          {
            id: 2,
            name: '治疗效果评估',
            type: 'treatment_analysis',
            created_at: new Date(),
            description: '评估不同治疗方案的效果'
          }
        ]
      } catch (error) {
        ElMessage.error('加载查询列表失败: ' + error.message)
      } finally {
        queriesLoading.value = false
      }
    }

    const showQueryDialog = () => {
      queryDialogVisible.value = true
      Object.assign(queryForm, {
        name: '',
        type: '',
        inclusion_criteria: '',
        exclusion_criteria: '',
        date_range: null,
        description: ''
      })
    }

    const useTemplate = (template) => {
      queryForm.name = template.name
      queryForm.type = template.type
      queryForm.inclusion_criteria = template.inclusion_criteria
      queryForm.exclusion_criteria = template.exclusion_criteria
      queryForm.description = template.description
      queryDialogVisible.value = true
    }

    const submitQuery = async () => {
      if (!queryFormRef.value) return

      try {
        await queryFormRef.value.validate()

        queryLoading.value = true

        const queryData = {
          databases: ['clinical', 'epidemiology', 'followup', 'biobank'],
          conditions: {
            inclusion: queryForm.inclusion_criteria,
            exclusion: queryForm.exclusion_criteria
          },
          date_range: queryForm.date_range,
          analysis_type: queryForm.type
        }

        const response = await integrationApi.getCohortInsights(queryData)
        insightResults.value = response.data

        queryDialogVisible.value = false
        ElMessage.success('查询执行成功')

        // 渲染图表
        await nextTick()
        renderCharts()
      } catch (error) {
        ElMessage.error('查询执行失败: ' + error.message)
      } finally {
        queryLoading.value = false
      }
    }

    const executeQuery = async (query) => {
      try {
        // 模拟执行已保存的查询
        const mockResults = {
          total_patients: 36,
          average_age: 41.2,
          followup_completion_rate: 82.5,
          sample_collection_rate: 75.0,
          gender_distribution: [
            { gender: '男', value: 20 },
            { gender: '女', value: 16 }
          ],
          age_distribution: [
            { age: '0-18', value: 3 },
            { age: '19-30', value: 7 },
            { age: '31-45', value: 12 },
            { age: '46-60', value: 9 },
            { age: '60+', value: 5 }
          ],
          primary_diagnoses: [
            { diagnosis: '过敏性鼻炎', count: 10, percentage: 27.8 },
            { diagnosis: '哮喘', count: 8, percentage: 22.2 },
            { diagnosis: '荨麻疹', count: 7, percentage: 19.4 },
            { diagnosis: '湿疹', count: 6, percentage: 16.7 },
            { diagnosis: '其他', count: 5, percentage: 13.9 }
          ],
          allergen_distribution: [
            { allergen: '尘螨', count: 12, severity: '中' },
            { allergen: '花粉', count: 8, severity: '轻' },
            { allergen: '食物', count: 7, severity: '重' },
            { allergen: '药物', count: 5, severity: '中' },
            { allergen: '其他', count: 4, severity: '轻' }
          ],
          treatment_plans: [
            { treatment: '抗组胺药物', count: 15, effectiveness: 80 },
            { treatment: '激素治疗', count: 8, effectiveness: 75 },
            { treatment: '免疫治疗', count: 7, effectiveness: 70 },
            { treatment: '中医治疗', count: 3, effectiveness: 60 },
            { treatment: '其他', count: 3, effectiveness: 50 }
          ],
          key_insights: [
            {
              type: 'finding',
              title: '尘螨过敏占主导地位',
              description: '尘螨是最常见的过敏原，占所有过敏患者的41.6%，需要重点关注环境控制措施。',
              metrics: [
                { name: '患病率', value: '41.6%', significance: 'high' },
                { name: '严重程度', value: '中等', significance: 'medium' }
              ]
            },
            {
              type: 'recommendation',
              title: '脱敏治疗效果显著',
              description: '脱敏治疗的有效率达到89.2%，建议扩大适用范围。',
              metrics: [
                { name: '有效率', value: '89.2%', significance: 'high' },
                { name: '安全性', value: '良好', significance: 'high' }
              ]
            }
          ],
          clinical_recommendations: [
            '加强尘螨过敏的环境控制教育',
            '推广脱敏治疗在适合患者中的应用',
            '建立过敏原检测标准化流程',
            '完善过敏性疾病的分级诊疗体系'
          ],
          research_recommendations: [
            '开展大规模过敏原流行病学调查',
            '研究新型脱敏治疗方法',
            '建立过敏性疾病预测模型',
            '开发个性化治疗方案'
          ]
        }

        insightResults.value = mockResults
        ElMessage.success('查询执行成功')

        await nextTick()
        renderCharts()
      } catch (error) {
        ElMessage.error('查询执行失败: ' + error.message)
      }
    }

    const renderCharts = () => {
      // 这里应该使用真实的图表库（如 ECharts）来渲染图表
      // 由于示例代码，这里只是模拟
      console.log('渲染图表...')
    }

    const editQuery = (query) => {
      // 编辑查询逻辑
      ElMessage.info('编辑功能开发中...')
    }

    const deleteQuery = async (queryId) => {
      try {
        await ElMessageBox.confirm('确定要删除这个查询吗？', '确认删除', {
          type: 'warning'
        })

        ElMessage.success('查询已删除')
        loadSavedQueries()
      } catch (error) {
        if (error !== 'cancel') {
          ElMessage.error('删除失败: ' + error.message)
        }
      }
    }

    const saveQuery = () => {
      ElMessage.success('查询已保存')
    }

    const exportInsights = () => {
      ElMessage.success('报告导出中...')
    }

    // 辅助函数
    const getQueryTypeName = (type) => {
      const map = {
        basic_cohort: '基础队列',
        disease_analysis: '疾病分析',
        treatment_analysis: '治疗分析',
        followup_analysis: '随访分析',
        survival_analysis: '生存分析'
      }
      return map[type] || type
    }

    const getQueryTypeTag = (type) => {
      const map = {
        basic_cohort: 'primary',
        disease_analysis: 'success',
        treatment_analysis: 'warning',
        followup_analysis: 'info',
        survival_analysis: 'danger'
      }
      return map[type] || ''
    }

    const getSeverityName = (severity) => {
      const map = {
        mild: '轻度',
        moderate: '中度',
        severe: '重度'
      }
      return map[severity] || severity
    }

    const getSeverityTag = (severity) => {
      const map = {
        mild: 'success',
        moderate: 'warning',
        severe: 'danger'
      }
      return map[severity] || ''
    }

    const getInsightIcon = (type) => {
      const map = {
        finding: 'el-icon-search',
        recommendation: 'el-icon-star-on',
        warning: 'el-icon-warning'
      }
      return map[type] || 'el-icon-info'
    }

    const getInsightColor = (type) => {
      const map = {
        finding: 'primary',
        recommendation: 'success',
        warning: 'warning'
      }
      return map[type] || 'info'
    }

    const getMetricTag = (significance) => {
      const map = {
        high: 'danger',
        medium: 'warning',
        low: 'info'
      }
      return map[significance] || ''
    }

    const formatDate = (date) => {
      return new Date(date).toLocaleString('zh-CN')
    }

    onMounted(() => {
      loadSavedQueries()
    })

    return {
      savedQueries,
      insightResults,
      queriesLoading,
      queryDialogVisible,
      queryLoading,
      queryForm,
      queryRules,
      queryFormRef,
      queryTemplates,
      loadSavedQueries,
      showQueryDialog,
      useTemplate,
      submitQuery,
      executeQuery,
      editQuery,
      deleteQuery,
      saveQuery,
      exportInsights,
      getQueryTypeName,
      getQueryTypeTag,
      getSeverityName,
      getSeverityTag,
      getInsightIcon,
      getInsightColor,
      getMetricTag,
      formatDate
    }
  }
}
</script>

<style scoped>
.cohort-insights {
  padding: 20px;
}

.header-card, .template-card, .saved-queries-card {
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

.template-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
}

.template-item {
  display: flex;
  align-items: center;
  padding: 20px;
  border: 1px solid #dcdfe6;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
}

.template-item:hover {
  border-color: #409eff;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.template-icon {
  font-size: 32px;
  color: #409eff;
  margin-right: 15px;
}

.template-content h4 {
  margin: 0 0 8px 0;
  color: #303133;
}

.template-content p {
  margin: 0;
  color: #909399;
  font-size: 14px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.insights-content {
  margin-top: 20px;
}

.insights-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.basic-stats {
  margin-bottom: 30px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
}

.demographics, .disease-characteristics, .key-insights, .recommendations {
  margin-bottom: 30px;
}

.demographics h3, .disease-characteristics h3, .key-insights h3, .recommendations h3 {
  margin-bottom: 20px;
  color: #303133;
}

.chart-container {
  text-align: center;
}

.chart-container h4 {
  margin-bottom: 15px;
  color: #606266;
}

.chart {
  height: 300px;
  background: #f5f7fa;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #909399;
}

.feature-card {
  height: 100%;
}

.insight-content h4 {
  margin: 0 0 10px 0;
  color: #303133;
}

.insight-content p {
  margin: 0 0 10px 0;
  color: #606266;
}

.insight-metrics {
  margin-top: 10px;
}

.recommendation-card {
  height: 100%;
}

.recommendation-list {
  margin: 0;
  padding-left: 20px;
  color: #606266;
}

.recommendation-list li {
  margin-bottom: 8px;
  line-height: 1.5;
}

.dialog-footer {
  text-align: right;
}
</style> 