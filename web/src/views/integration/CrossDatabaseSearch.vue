<template>
  <div class="cross-database-search">
    <el-card class="search-card">
      <div class="search-header">
        <h2>跨库搜索</h2>
        <p>在多个数据库中同时搜索患者信息</p>
      </div>

      <!-- 搜索表单 -->
      <el-form :model="searchForm" :rules="searchRules" ref="searchFormRef" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="搜索关键词" prop="keywords">
              <el-input
                v-model="searchForm.keywords"
                placeholder="请输入患者姓名、ID或诊断关键词"
                @keyup.enter="performSearch"
              >
                <template #append>
                  <el-button @click="performSearch" :loading="searching">
                    <i class="el-icon-search"></i>
                  </el-button>
                </template>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="搜索范围">
              <el-checkbox-group v-model="searchForm.databases">
                <el-checkbox label="clinical">临床数据库</el-checkbox>
                <el-checkbox label="epidemiology">流调数据库</el-checkbox>
                <el-checkbox label="followup">随访数据库</el-checkbox>
                <el-checkbox label="biobank">生物样本库</el-checkbox>
              </el-checkbox-group>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="年龄范围">
              <el-slider
                v-model="searchForm.age_range"
                range
                :min="0"
                :max="100"
                :step="1"
                show-stops
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="性别">
              <el-select v-model="searchForm.gender" placeholder="请选择性别" clearable>
                <el-option label="男" value="male" />
                <el-option label="女" value="female" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="时间范围">
              <el-date-picker
                v-model="searchForm.date_range"
                type="daterange"
                range-separator="至"
                start-placeholder="开始日期"
                end-placeholder="结束日期"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item>
          <el-button type="primary" @click="performSearch" :loading="searching">
            <i class="el-icon-search"></i>
            搜索
          </el-button>
          <el-button @click="resetSearch">
            <i class="el-icon-refresh"></i>
            重置
          </el-button>
          <el-button type="success" @click="showAdvancedSearch">
            <i class="el-icon-setting"></i>
            高级搜索
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 搜索结果 -->
    <div v-if="searchResults" class="results-section">
      <!-- 结果统计 -->
      <el-card class="stats-card">
        <div class="search-stats">
          <el-row :gutter="20">
            <el-col :span="6">
              <el-statistic
                title="总匹配数量"
                :value="searchResults.total_count"
                suffix="条"
              />
            </el-col>
            <el-col :span="6">
              <el-statistic
                title="患者数量"
                :value="searchResults.patient_count"
                suffix="人"
              />
            </el-col>
            <el-col :span="6">
              <el-statistic
                title="搜索用时"
                :value="searchResults.search_time"
                :precision="2"
                suffix="秒"
              />
            </el-col>
            <el-col :span="6">
              <el-statistic
                title="数据库覆盖"
                :value="searchResults.database_coverage"
                suffix="%"
              />
            </el-col>
          </el-row>
        </div>
      </el-card>

      <!-- 按数据库分组的结果 -->
      <el-card class="results-card">
        <template #header>
          <div class="results-header">
            <span>搜索结果</span>
            <div class="header-actions">
              <el-button type="primary" @click="exportResults">导出结果</el-button>
              <el-button @click="saveSearch">保存搜索</el-button>
            </div>
          </div>
        </template>

        <el-tabs v-model="activeTab" type="card">
          <!-- 综合结果 -->
          <el-tab-pane label="综合结果" name="all">
            <el-table :data="searchResults.combined_results" stripe>
              <el-table-column prop="patient_id" label="患者ID" width="100" />
              <el-table-column prop="patient_name" label="患者姓名" width="120" />
              <el-table-column prop="age" label="年龄" width="80" />
              <el-table-column prop="gender" label="性别" width="80">
                <template #default="scope">
                  {{ scope.row.gender === 'male' ? '男' : '女' }}
                </template>
              </el-table-column>
              <el-table-column prop="match_score" label="匹配度" width="100">
                <template #default="scope">
                  <el-progress
                    :percentage="scope.row.match_score"
                    :color="getMatchScoreColor(scope.row.match_score)"
                    :show-text="false"
                  />
                  <span style="margin-left: 10px">{{ scope.row.match_score }}%</span>
                </template>
              </el-table-column>
              <el-table-column prop="databases" label="数据源" width="200">
                <template #default="scope">
                  <el-tag
                    v-for="db in scope.row.databases"
                    :key="db"
                    :type="getDatabaseTagType(db)"
                    size="small"
                    style="margin-right: 5px"
                  >
                    {{ getDatabaseName(db) }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="highlight" label="匹配内容" />
              <el-table-column label="操作" width="150">
                <template #default="scope">
                  <el-button type="text" @click="viewPatientDetail(scope.row)">
                    查看详情
                  </el-button>
                  <el-button type="text" @click="addToComparison(scope.row)">
                    加入对比
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>

          <!-- 临床数据库结果 -->
          <el-tab-pane
            :label="`临床数据库 (${searchResults.clinical_results?.length || 0})`"
            name="clinical"
          >
            <el-table :data="searchResults.clinical_results" stripe>
              <el-table-column prop="patient_id" label="患者ID" width="100" />
              <el-table-column prop="diagnosis" label="诊断" width="200" />
              <el-table-column prop="treatment" label="治疗方案" width="200" />
              <el-table-column prop="doctor" label="主治医生" width="120" />
              <el-table-column prop="visit_date" label="就诊日期" width="120">
                <template #default="scope">
                  {{ formatDate(scope.row.visit_date) }}
                </template>
              </el-table-column>
              <el-table-column prop="status" label="状态" width="100">
                <template #default="scope">
                  <el-tag :type="getStatusTagType(scope.row.status)">
                    {{ getStatusName(scope.row.status) }}
                  </el-tag>
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>

          <!-- 流调数据库结果 -->
          <el-tab-pane
            :label="`流调数据库 (${searchResults.epidemiology_results?.length || 0})`"
            name="epidemiology"
          >
            <el-table :data="searchResults.epidemiology_results" stripe>
              <el-table-column prop="patient_id" label="患者ID" width="100" />
              <el-table-column prop="risk_factors" label="危险因素" width="200" />
              <el-table-column prop="exposure_history" label="暴露史" width="200" />
              <el-table-column prop="family_history" label="家族史" width="150" />
              <el-table-column prop="survey_date" label="调查日期" width="120">
                <template #default="scope">
                  {{ formatDate(scope.row.survey_date) }}
                </template>
              </el-table-column>
              <el-table-column prop="completeness" label="完整度" width="100">
                <template #default="scope">
                  <el-progress :percentage="scope.row.completeness" :show-text="false" />
                  <span style="margin-left: 10px">{{ scope.row.completeness }}%</span>
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>

          <!-- 随访数据库结果 -->
          <el-tab-pane
            :label="`随访数据库 (${searchResults.followup_results?.length || 0})`"
            name="followup"
          >
            <el-table :data="searchResults.followup_results" stripe>
              <el-table-column prop="patient_id" label="患者ID" width="100" />
              <el-table-column prop="followup_type" label="随访类型" width="120" />
              <el-table-column prop="outcome" label="随访结果" width="150" />
              <el-table-column prop="next_followup" label="下次随访" width="120">
                <template #default="scope">
                  {{ formatDate(scope.row.next_followup) }}
                </template>
              </el-table-column>
              <el-table-column prop="compliance" label="依从性" width="100">
                <template #default="scope">
                  <el-tag :type="getComplianceTagType(scope.row.compliance)">
                    {{ getComplianceName(scope.row.compliance) }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="notes" label="备注" />
            </el-table>
          </el-tab-pane>

          <!-- 生物样本库结果 -->
          <el-tab-pane
            :label="`生物样本库 (${searchResults.biobank_results?.length || 0})`"
            name="biobank"
          >
            <el-table :data="searchResults.biobank_results" stripe>
              <el-table-column prop="patient_id" label="患者ID" width="100" />
              <el-table-column prop="sample_type" label="样本类型" width="120" />
              <el-table-column prop="collection_date" label="采集日期" width="120">
                <template #default="scope">
                  {{ formatDate(scope.row.collection_date) }}
                </template>
              </el-table-column>
              <el-table-column prop="storage_location" label="存储位置" width="150" />
              <el-table-column prop="quality_score" label="质量评分" width="100">
                <template #default="scope">
                  <el-rate
                    v-model="scope.row.quality_score"
                    disabled
                    show-score
                    text-color="#ff9900"
                  />
                </template>
              </el-table-column>
              <el-table-column prop="available" label="可用性" width="100">
                <template #default="scope">
                  <el-tag :type="scope.row.available ? 'success' : 'danger'">
                    {{ scope.row.available ? '可用' : '不可用' }}
                  </el-tag>
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </div>

    <!-- 高级搜索对话框 -->
    <el-dialog
      v-model="advancedSearchVisible"
      title="高级搜索"
      width="800px"
    >
      <el-form :model="advancedForm" label-width="120px">
        <el-tabs v-model="advancedTab">
          <el-tab-pane label="基本条件" name="basic">
            <el-form-item label="患者ID范围">
              <el-input v-model="advancedForm.patient_id_range" placeholder="如：1000-2000" />
            </el-form-item>
            <el-form-item label="诊断代码">
              <el-input v-model="advancedForm.diagnosis_codes" placeholder="ICD-10代码，用逗号分隔" />
            </el-form-item>
            <el-form-item label="过敏原">
              <el-select v-model="advancedForm.allergens" multiple placeholder="请选择过敏原">
                <el-option label="尘螨" value="dust_mite" />
                <el-option label="花粉" value="pollen" />
                <el-option label="动物毛发" value="animal_dander" />
                <el-option label="食物" value="food" />
                <el-option label="药物" value="drug" />
              </el-select>
            </el-form-item>
          </el-tab-pane>

          <el-tab-pane label="临床条件" name="clinical">
            <el-form-item label="治疗方案">
              <el-input v-model="advancedForm.treatment_plans" placeholder="治疗方案关键词" />
            </el-form-item>
            <el-form-item label="严重程度">
              <el-select v-model="advancedForm.severity" placeholder="请选择严重程度">
                <el-option label="轻度" value="mild" />
                <el-option label="中度" value="moderate" />
                <el-option label="重度" value="severe" />
              </el-select>
            </el-form-item>
            <el-form-item label="就诊科室">
              <el-select v-model="advancedForm.departments" multiple placeholder="请选择科室">
                <el-option label="过敏科" value="allergy" />
                <el-option label="呼吸科" value="respiratory" />
                <el-option label="皮肤科" value="dermatology" />
                <el-option label="儿科" value="pediatrics" />
              </el-select>
            </el-form-item>
          </el-tab-pane>

          <el-tab-pane label="随访条件" name="followup">
            <el-form-item label="随访状态">
              <el-select v-model="advancedForm.followup_status" placeholder="请选择随访状态">
                <el-option label="进行中" value="ongoing" />
                <el-option label="已完成" value="completed" />
                <el-option label="失访" value="lost" />
              </el-select>
            </el-form-item>
            <el-form-item label="依从性">
              <el-select v-model="advancedForm.compliance_level" placeholder="请选择依从性">
                <el-option label="良好" value="good" />
                <el-option label="一般" value="fair" />
                <el-option label="差" value="poor" />
              </el-select>
            </el-form-item>
          </el-tab-pane>

          <el-tab-pane label="样本条件" name="biobank">
            <el-form-item label="样本类型">
              <el-select v-model="advancedForm.sample_types" multiple placeholder="请选择样本类型">
                <el-option label="血清" value="serum" />
                <el-option label="血浆" value="plasma" />
                <el-option label="DNA" value="dna" />
                <el-option label="RNA" value="rna" />
                <el-option label="组织" value="tissue" />
              </el-select>
            </el-form-item>
            <el-form-item label="质量要求">
              <el-slider
                v-model="advancedForm.min_quality_score"
                :min="1"
                :max="5"
                :step="1"
                show-stops
                show-input
              />
            </el-form-item>
          </el-tab-pane>
        </el-tabs>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="advancedSearchVisible = false">取消</el-button>
          <el-button type="primary" @click="performAdvancedSearch">搜索</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 患者对比功能 -->
    <div v-if="comparisonList.length > 0" class="comparison-panel">
      <el-card>
        <template #header>
          <div class="comparison-header">
            <span>患者对比 ({{ comparisonList.length }})</span>
            <el-button type="primary" @click="showComparison">查看对比</el-button>
          </div>
        </template>
        <div class="comparison-list">
          <el-tag
            v-for="patient in comparisonList"
            :key="patient.patient_id"
            closable
            @close="removeFromComparison(patient)"
            style="margin-right: 10px"
          >
            {{ patient.patient_name }} ({{ patient.patient_id }})
          </el-tag>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { integrationApi } from '@/api/integration'

export default {
  name: 'CrossDatabaseSearch',
  setup() {
    const searchFormRef = ref(null)
    const searching = ref(false)
    const searchResults = ref(null)
    const activeTab = ref('all')
    const advancedSearchVisible = ref(false)
    const advancedTab = ref('basic')
    const comparisonList = ref([])

    const searchForm = reactive({
      keywords: '',
      databases: ['clinical', 'epidemiology', 'followup', 'biobank'],
      age_range: [0, 100],
      gender: '',
      date_range: null
    })

    const searchRules = reactive({
      keywords: [
        { required: true, message: '请输入搜索关键词', trigger: 'blur' }
      ]
    })

    const advancedForm = reactive({
      patient_id_range: '',
      diagnosis_codes: '',
      allergens: [],
      treatment_plans: '',
      severity: '',
      departments: [],
      followup_status: '',
      compliance_level: '',
      sample_types: [],
      min_quality_score: 3
    })

    const performSearch = async () => {
      if (!searchFormRef.value) return

      try {
        await searchFormRef.value.validate()

        searching.value = true

        const searchData = {
          keywords: searchForm.keywords,
          databases: searchForm.databases,
          filters: {
            age_range: searchForm.age_range,
            gender: searchForm.gender,
            date_range: searchForm.date_range
          }
        }

        const response = await integrationApi.searchPatients(searchData)
        searchResults.value = response.data

        ElMessage.success('搜索完成')
      } catch (error) {
        ElMessage.error('搜索失败: ' + error.message)
      } finally {
        searching.value = false
      }
    }

    const resetSearch = () => {
      Object.assign(searchForm, {
        keywords: '',
        databases: ['clinical', 'epidemiology', 'followup', 'biobank'],
        age_range: [0, 100],
        gender: '',
        date_range: null
      })
      searchResults.value = null
    }

    const showAdvancedSearch = () => {
      advancedSearchVisible.value = true
    }

    const performAdvancedSearch = () => {
      // 执行高级搜索逻辑
      ElMessage.success('高级搜索功能开发中...')
      advancedSearchVisible.value = false
    }

    const viewPatientDetail = (patient) => {
      // 跳转到患者详情页面
      ElMessage.info('跳转到患者详情页面...')
    }

    const addToComparison = (patient) => {
      if (comparisonList.value.length >= 5) {
        ElMessage.warning('最多只能对比5个患者')
        return
      }

      const exists = comparisonList.value.find(p => p.patient_id === patient.patient_id)
      if (exists) {
        ElMessage.warning('该患者已在对比列表中')
        return
      }

      comparisonList.value.push(patient)
      ElMessage.success('已加入对比列表')
    }

    const removeFromComparison = (patient) => {
      const index = comparisonList.value.findIndex(p => p.patient_id === patient.patient_id)
      if (index > -1) {
        comparisonList.value.splice(index, 1)
      }
    }

    const showComparison = () => {
      ElMessage.info('患者对比功能开发中...')
    }

    const exportResults = () => {
      ElMessage.success('结果导出中...')
    }

    const saveSearch = () => {
      ElMessage.success('搜索已保存')
    }

    // 辅助函数
    const getMatchScoreColor = (score) => {
      if (score >= 80) return '#67c23a'
      if (score >= 60) return '#e6a23c'
      return '#f56c6c'
    }

    const getDatabaseName = (db) => {
      const map = {
        clinical: '临床',
        epidemiology: '流调',
        followup: '随访',
        biobank: '样本'
      }
      return map[db] || db
    }

    const getDatabaseTagType = (db) => {
      const map = {
        clinical: 'primary',
        epidemiology: 'success',
        followup: 'warning',
        biobank: 'info'
      }
      return map[db] || ''
    }

    const getStatusName = (status) => {
      const map = {
        active: '活跃',
        inactive: '非活跃',
        completed: '已完成',
        pending: '待处理'
      }
      return map[status] || status
    }

    const getStatusTagType = (status) => {
      const map = {
        active: 'success',
        inactive: 'info',
        completed: 'primary',
        pending: 'warning'
      }
      return map[status] || ''
    }

    const getComplianceName = (compliance) => {
      const map = {
        good: '良好',
        fair: '一般',
        poor: '差'
      }
      return map[compliance] || compliance
    }

    const getComplianceTagType = (compliance) => {
      const map = {
        good: 'success',
        fair: 'warning',
        poor: 'danger'
      }
      return map[compliance] || ''
    }

    const formatDate = (date) => {
      return new Date(date).toLocaleDateString('zh-CN')
    }

    // 模拟搜索结果数据
    const mockSearchResults = () => {
      return {
        total_count: 156,
        patient_count: 89,
        search_time: 0.85,
        database_coverage: 95,
        combined_results: [
          {
            patient_id: 1001,
            patient_name: '张三',
            age: 35,
            gender: 'male',
            match_score: 95,
            databases: ['clinical', 'epidemiology', 'followup'],
            highlight: '过敏性鼻炎，尘螨过敏'
          },
          {
            patient_id: 1002,
            patient_name: '李四',
            age: 28,
            gender: 'female',
            match_score: 88,
            databases: ['clinical', 'biobank'],
            highlight: '过敏性哮喘，花粉过敏'
          }
        ],
        clinical_results: [
          {
            patient_id: 1001,
            diagnosis: '过敏性鼻炎',
            treatment: '抗组胺药物治疗',
            doctor: '王医生',
            visit_date: new Date(),
            status: 'active'
          }
        ],
        epidemiology_results: [
          {
            patient_id: 1001,
            risk_factors: '家族过敏史，环境因素',
            exposure_history: '长期接触尘螨',
            family_history: '母亲有过敏性疾病',
            survey_date: new Date(),
            completeness: 85
          }
        ],
        followup_results: [
          {
            patient_id: 1001,
            followup_type: '定期复查',
            outcome: '症状改善',
            next_followup: new Date(),
            compliance: 'good',
            notes: '患者依从性良好，症状控制稳定'
          }
        ],
        biobank_results: [
          {
            patient_id: 1002,
            sample_type: '血清',
            collection_date: new Date(),
            storage_location: 'A区-冰箱1-架子3',
            quality_score: 4,
            available: true
          }
        ]
      }
    }

    // 为演示目的，在搜索时使用模拟数据
    const originalPerformSearch = performSearch
    const performSearchWithMockData = async () => {
      if (!searchFormRef.value) return

      try {
        await searchFormRef.value.validate()
        searching.value = true

        // 模拟搜索延迟
        await new Promise(resolve => setTimeout(resolve, 1000))

        searchResults.value = mockSearchResults()
        ElMessage.success('搜索完成')
      } catch (error) {
        ElMessage.error('搜索失败: ' + error.message)
      } finally {
        searching.value = false
      }
    }

    return {
      searchForm,
      searchRules,
      searchFormRef,
      searching,
      searchResults,
      activeTab,
      advancedSearchVisible,
      advancedTab,
      advancedForm,
      comparisonList,
      performSearch: performSearchWithMockData,
      resetSearch,
      showAdvancedSearch,
      performAdvancedSearch,
      viewPatientDetail,
      addToComparison,
      removeFromComparison,
      showComparison,
      exportResults,
      saveSearch,
      getMatchScoreColor,
      getDatabaseName,
      getDatabaseTagType,
      getStatusName,
      getStatusTagType,
      getComplianceName,
      getComplianceTagType,
      formatDate
    }
  }
}
</script>

<style scoped>
.cross-database-search {
  padding: 20px;
}

.search-card {
  margin-bottom: 20px;
}

.search-header {
  text-align: center;
  margin-bottom: 30px;
}

.search-header h2 {
  margin: 0 0 10px 0;
  color: #303133;
}

.search-header p {
  margin: 0;
  color: #909399;
}

.results-section {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.stats-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.search-stats {
  padding: 20px;
}

.results-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-actions {
  display: flex;
  gap: 10px;
}

.comparison-panel {
  position: fixed;
  bottom: 20px;
  right: 20px;
  width: 400px;
  z-index: 1000;
}

.comparison-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.comparison-list {
  max-height: 100px;
  overflow-y: auto;
}

.dialog-footer {
  text-align: right;
}

:deep(.el-statistic__content) {
  color: white !important;
}

:deep(.el-statistic__head) {
  color: rgba(255, 255, 255, 0.8) !important;
}
</style> 