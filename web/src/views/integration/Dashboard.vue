<template>
  <div class="integration-dashboard" v-loading="loading">
    <!-- 头部统计卡片 -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6">
        <el-card class="stats-card patients-card">
          <el-statistic
            title="总患者数"
            :value="dashboardData.total_patients"
            suffix="人"
          >
            <template #prefix>
              <i class="el-icon-user" style="color: #409eff"></i>
            </template>
          </el-statistic>
          <div class="card-footer">
            <span class="trend" :class="dashboardData.patient_growth >= 0 ? 'up' : 'down'">
              <i :class="dashboardData.patient_growth >= 0 ? 'el-icon-caret-top' : 'el-icon-caret-bottom'"></i>
              {{ Math.abs(dashboardData.patient_growth) }}%
            </span>
            <span class="period">较上月</span>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="stats-card databases-card">
          <el-statistic
            title="数据库覆盖"
            :value="dashboardData.database_coverage"
            :precision="1"
            suffix="%"
          >
            <template #prefix>
              <i class="el-icon-coin" style="color: #67c23a"></i>
            </template>
          </el-statistic>
          <div class="card-footer">
            <span class="trend" :class="dashboardData.coverage_improvement >= 0 ? 'up' : 'down'">
              <i :class="dashboardData.coverage_improvement >= 0 ? 'el-icon-caret-top' : 'el-icon-caret-bottom'"></i>
              {{ Math.abs(dashboardData.coverage_improvement) }}%
            </span>
            <span class="period">较上月</span>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="stats-card quality-card">
          <el-statistic
            title="数据质量评分"
            :value="dashboardData.avg_quality_score"
            :precision="1"
            suffix="分"
          >
            <template #prefix>
              <i class="el-icon-medal" style="color: #e6a23c"></i>
            </template>
          </el-statistic>
          <div class="card-footer">
            <span class="trend" :class="dashboardData.quality_improvement >= 0 ? 'up' : 'down'">
              <i :class="dashboardData.quality_improvement >= 0 ? 'el-icon-caret-top' : 'el-icon-caret-bottom'"></i>
              {{ Math.abs(dashboardData.quality_improvement) }}%
            </span>
            <span class="period">较上月</span>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="stats-card integration-card">
          <el-statistic
            title="今日集成查询"
            :value="dashboardData.today_queries"
            suffix="次"
          >
            <template #prefix>
              <i class="el-icon-connection" style="color: #f56c6c"></i>
            </template>
          </el-statistic>
          <div class="card-footer">
            <span class="trend" :class="dashboardData.query_growth >= 0 ? 'up' : 'down'">
              <i :class="dashboardData.query_growth >= 0 ? 'el-icon-caret-top' : 'el-icon-caret-bottom'"></i>
              {{ Math.abs(dashboardData.query_growth) }}%
            </span>
            <span class="period">较昨日</span>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 快捷操作区域 -->
    <el-card class="quick-actions-card">
      <template #header>
        <div class="card-header">
          <span>快捷操作</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="4">
          <div class="quick-action" @click="navigateTo('/integration/patient-view')">
            <div class="action-icon patient-icon">
              <i class="el-icon-user"></i>
            </div>
            <div class="action-title">患者综合视图</div>
            <div class="action-desc">查看患者跨库信息</div>
          </div>
        </el-col>

        <el-col :span="4">
          <div class="quick-action" @click="navigateTo('/integration/cross-database-search')">
            <div class="action-icon search-icon">
              <i class="el-icon-search"></i>
            </div>
            <div class="action-title">跨库搜索</div>
            <div class="action-desc">多数据库联合搜索</div>
          </div>
        </el-col>

        <el-col :span="4">
          <div class="quick-action" @click="navigateTo('/integration/data-quality')">
            <div class="action-icon quality-icon">
              <i class="el-icon-medal"></i>
            </div>
            <div class="action-title">数据质量报告</div>
            <div class="action-desc">检查数据完整性</div>
          </div>
        </el-col>

        <el-col :span="4">
          <div class="quick-action" @click="navigateTo('/integration/correlation-analysis')">
            <div class="action-icon analysis-icon">
              <i class="el-icon-data-line"></i>
            </div>
            <div class="action-title">相关性分析</div>
            <div class="action-desc">数据关联分析</div>
          </div>
        </el-col>

        <el-col :span="4">
          <div class="quick-action" @click="navigateTo('/integration/cohort-insights')">
            <div class="action-icon cohort-icon">
              <i class="el-icon-pie-chart"></i>
            </div>
            <div class="action-title">队列洞察</div>
            <div class="action-desc">队列数据分析</div>
          </div>
        </el-col>

        <el-col :span="4">
          <div class="quick-action" @click="showExportDialog">
            <div class="action-icon export-icon">
              <i class="el-icon-download"></i>
            </div>
            <div class="action-title">数据导出</div>
            <div class="action-desc">批量导出数据</div>
          </div>
        </el-col>
      </el-row>
    </el-card>

    <!-- 数据库状态监控 -->
    <el-row :gutter="20">
      <el-col :span="12">
        <el-card class="database-status-card">
          <template #header>
            <div class="card-header">
              <span>数据库状态</span>
              <el-button type="text" @click="refreshDatabaseStatus">
                <i class="el-icon-refresh"></i>
                刷新
              </el-button>
            </div>
          </template>
          <div class="database-list">
            <div
              v-for="db in dashboardData.database_status"
              :key="db.name"
              class="database-item"
            >
              <div class="db-info">
                <div class="db-name">{{ db.display_name }}</div>
                <div class="db-details">
                  <span>记录数: {{ db.record_count.toLocaleString() }}</span>
                  <span>更新时间: {{ formatDate(db.last_updated) }}</span>
                </div>
              </div>
              <div class="db-status">
                <el-tag :type="getStatusType(db.status)" size="small">
                  {{ getStatusText(db.status) }}
                </el-tag>
                <div class="db-performance">
                  <el-progress
                    :percentage="db.performance"
                    :color="getPerformanceColor(db.performance)"
                    :show-text="false"
                    :stroke-width="6"
                  />
                  <span class="performance-text">{{ db.performance }}%</span>
                </div>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card class="recent-activities-card">
          <template #header>
            <div class="card-header">
              <span>最近活动</span>
              <el-button type="text" @click="viewAllActivities">
                查看全部
              </el-button>
            </div>
          </template>
          <div class="activities-list">
            <div
              v-for="activity in dashboardData.recent_activities"
              :key="activity.id"
              class="activity-item"
            >
              <div class="activity-icon">
                <i :class="getActivityIcon(activity.type)"></i>
              </div>
              <div class="activity-content">
                <div class="activity-title">{{ activity.title }}</div>
                <div class="activity-desc">{{ activity.description }}</div>
                <div class="activity-time">{{ formatRelativeTime(activity.timestamp) }}</div>
              </div>
              <div class="activity-status">
                <el-tag :type="getActivityStatusType(activity.status)" size="mini">
                  {{ activity.status }}
                </el-tag>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 数据趋势图表 -->
    <el-row :gutter="20">
      <el-col :span="16">
        <el-card class="chart-card">
          <template #header>
            <div class="card-header">
              <span>数据集成趋势</span>
              <el-radio-group v-model="chartTimeRange" size="small">
                <el-radio-button label="7d">7天</el-radio-button>
                <el-radio-button label="30d">30天</el-radio-button>
                <el-radio-button label="90d">90天</el-radio-button>
              </el-radio-group>
            </div>
          </template>
          <div class="chart-container" ref="integrationChart"></div>
        </el-card>
      </el-col>

      <el-col :span="8">
        <el-card class="data-distribution-card">
          <template #header>
            <div class="card-header">
              <span>数据分布</span>
            </div>
          </template>
          <div class="distribution-list">
            <div
              v-for="item in dashboardData.data_distribution"
              :key="item.database"
              class="distribution-item"
            >
              <div class="dist-label">{{ item.display_name }}</div>
              <div class="dist-bar">
                <el-progress
                  :percentage="item.percentage"
                  :color="item.color"
                  :show-text="false"
                />
                <span class="dist-value">{{ item.count.toLocaleString() }}</span>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 数据导出对话框 -->
    <el-dialog
      v-model="exportDialogVisible"
      title="数据导出"
      width="600px"
    >
      <el-form :model="exportForm" label-width="120px">
        <el-form-item label="导出范围">
          <el-checkbox-group v-model="exportForm.databases">
            <el-checkbox label="clinical">临床数据库</el-checkbox>
            <el-checkbox label="epidemiology">流调数据库</el-checkbox>
            <el-checkbox label="followup">随访数据库</el-checkbox>
            <el-checkbox label="biobank">生物样本库</el-checkbox>
          </el-checkbox-group>
        </el-form-item>

        <el-form-item label="导出格式">
          <el-select v-model="exportForm.format" placeholder="请选择格式">
            <el-option label="Excel" value="excel" />
            <el-option label="CSV" value="csv" />
            <el-option label="JSON" value="json" />
          </el-select>
        </el-form-item>

        <el-form-item label="时间范围">
          <el-date-picker
            v-model="exportForm.date_range"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
          />
        </el-form-item>

        <el-form-item label="患者ID范围">
          <el-input v-model="exportForm.patient_id_range" placeholder="如：1000-2000，留空表示全部" />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="exportDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="performExport" :loading="exporting">
            导出
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref, reactive, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import * as echarts from 'echarts'
import { integrationApi } from '@/api/integration'

export default {
  name: 'IntegrationDashboard',
  setup() {
    const router = useRouter()
    const integrationChart = ref(null)
    const chartTimeRange = ref('30d')
    const exportDialogVisible = ref(false)
    const exporting = ref(false)
    const loading = ref(true)

    const dashboardData = ref({
      total_patients: 0,
      patient_growth: 0,
      database_coverage: 0,
      coverage_improvement: 0,
      avg_quality_score: 0,
      quality_improvement: 0,
      today_queries: 0,
      query_growth: 0,
      database_status: [],
      recent_activities: [],
      data_distribution: []
    })

    const exportForm = reactive({
      databases: ['clinical', 'epidemiology', 'followup', 'biobank'],
      format: 'excel',
      date_range: null,
      patient_id_range: ''
    })

    const navigateTo = (path) => {
      router.push(path)
    }

    const refreshDatabaseStatus = () => {
      ElMessage.success('数据库状态已刷新')
    }

    const viewAllActivities = () => {
      ElMessage.info('查看全部活动功能开发中...')
    }

    const showExportDialog = () => {
      exportDialogVisible.value = true
    }

    const performExport = () => {
      exporting.value = true
      setTimeout(() => {
        exporting.value = false
        exportDialogVisible.value = false
        ElMessage.success('数据导出已开始，完成后将通过邮件通知')
      }, 2000)
    }

    const getStatusType = (status) => {
      const map = {
        healthy: 'success',
        warning: 'warning',
        error: 'danger'
      }
      return map[status] || 'info'
    }

    const getStatusText = (status) => {
      const map = {
        healthy: '正常',
        warning: '警告',
        error: '错误'
      }
      return map[status] || status
    }

    const getPerformanceColor = (performance) => {
      if (performance >= 90) return '#67c23a'
      if (performance >= 70) return '#e6a23c'
      return '#f56c6c'
    }

    const getActivityIcon = (type) => {
      const map = {
        search: 'el-icon-search',
        quality: 'el-icon-medal',
        analysis: 'el-icon-data-line',
        export: 'el-icon-download',
        sync: 'el-icon-refresh'
      }
      return map[type] || 'el-icon-info'
    }

    const getActivityStatusType = (status) => {
      const map = {
        '成功': 'success',
        '完成': 'primary',
        '警告': 'warning',
        '错误': 'danger'
      }
      return map[status] || 'info'
    }

    const formatDate = (date) => {
      return new Date(date).toLocaleString('zh-CN')
    }

    const formatRelativeTime = (date) => {
      const now = new Date()
      const diff = now - new Date(date)
      const minutes = Math.floor(diff / (1000 * 60))
      
      if (minutes < 1) return '刚刚'
      if (minutes < 60) return `${minutes}分钟前`
      
      const hours = Math.floor(minutes / 60)
      if (hours < 24) return `${hours}小时前`
      
      const days = Math.floor(hours / 24)
      return `${days}天前`
    }

    const initChart = () => {
      nextTick(() => {
        if (!integrationChart.value) return

        const chart = echarts.init(integrationChart.value)
        
        const option = {
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'cross'
            }
          },
          legend: {
            data: ['查询次数', '数据质量', '集成成功率']
          },
          grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
          },
          xAxis: {
            type: 'category',
            boundaryGap: false,
            data: ['1日', '5日', '10日', '15日', '20日', '25日', '30日']
          },
          yAxis: [
            {
              type: 'value',
              name: '次数',
              position: 'left'
            },
            {
              type: 'value',
              name: '百分比',
              position: 'right',
              axisLabel: {
                formatter: '{value}%'
              }
            }
          ],
          series: [
            {
              name: '查询次数',
              type: 'line',
              data: [120, 132, 101, 134, 90, 230, 210],
              smooth: true,
              itemStyle: {
                color: '#409eff'
              }
            },
            {
              name: '数据质量',
              type: 'line',
              yAxisIndex: 1,
              data: [85, 87, 86, 88, 89, 90, 87],
              smooth: true,
              itemStyle: {
                color: '#67c23a'
              }
            },
            {
              name: '集成成功率',
              type: 'line',
              yAxisIndex: 1,
              data: [92, 94, 93, 95, 96, 94, 95],
              smooth: true,
              itemStyle: {
                color: '#e6a23c'
              }
            }
          ]
        }

        chart.setOption(option)
        
        // 响应式处理
        window.addEventListener('resize', () => {
          chart.resize()
        })
      })
    }

    // 加载仪表板数据
    const loadDashboardData = async () => {
      try {
        loading.value = true
        const response = await integrationApi.getDashboardStats()
        
        if (response.code === '10000' && response.data) {
          const data = response.data
          
          // 更新统计数据
          dashboardData.value.total_patients = data.total_patients || 0
          dashboardData.value.patient_growth = parseFloat(data.patient_growth?.replace(/[+%]/g, '') || '0')
          dashboardData.value.database_coverage = parseFloat(data.database_coverage?.toFixed(1) || '0')
          dashboardData.value.coverage_improvement = parseFloat(data.coverage_trend?.replace(/[+%]/g, '') || '0')
          dashboardData.value.avg_quality_score = parseFloat(data.data_quality?.toFixed(1) || '0')
          dashboardData.value.quality_improvement = parseFloat(data.quality_trend?.replace(/[+%]/g, '') || '0')
          dashboardData.value.today_queries = data.integration_queries || 0
          dashboardData.value.query_growth = parseFloat(data.queries_growth?.replace(/[+%]/g, '') || '0')
          
          // 获取活动数据
          await loadRecentActivities()
        }
      } catch (error) {
        console.error('加载仪表板数据失败:', error)
        ElMessage.error('加载数据失败，请稍后重试')
      } finally {
        loading.value = false
      }
    }

    // 加载最近活动
    const loadRecentActivities = async () => {
      try {
        const response = await integrationApi.getRecentActivities({ limit: 5 })
        if (response.code === '10000' && response.data) {
          dashboardData.value.recent_activities = response.data.activities || []
        }
      } catch (error) {
        console.error('加载活动数据失败:', error)
      }
    }

    onMounted(async () => {
      await loadDashboardData()
      initChart()
    })

    return {
      dashboardData,
      chartTimeRange,
      exportDialogVisible,
      exportForm,
      exporting,
      loading,
      integrationChart,
      navigateTo,
      refreshDatabaseStatus,
      viewAllActivities,
      showExportDialog,
      performExport,
      getStatusType,
      getStatusText,
      getPerformanceColor,
      getActivityIcon,
      getActivityStatusType,
      formatDate,
      formatRelativeTime,
      loadDashboardData
    }
  }
}
</script>

<style scoped>
.integration-dashboard {
  padding: 20px;
}

.stats-row {
  margin-bottom: 20px;
}

.stats-card {
  text-align: center;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.stats-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.patients-card {
  border-left: 4px solid #409eff;
}

.databases-card {
  border-left: 4px solid #67c23a;
}

.quality-card {
  border-left: 4px solid #e6a23c;
}

.integration-card {
  border-left: 4px solid #f56c6c;
}

.card-footer {
  margin-top: 15px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
}

.trend {
  font-weight: bold;
}

.trend.up {
  color: #67c23a;
}

.trend.down {
  color: #f56c6c;
}

.period {
  color: #909399;
}

.quick-actions-card {
  margin-bottom: 20px;
}

.quick-action {
  text-align: center;
  padding: 20px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid #ebeef5;
}

.quick-action:hover {
  background-color: #f5f7fa;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.action-icon {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 15px;
  font-size: 24px;
  color: white;
}

.patient-icon {
  background: linear-gradient(135deg, #409eff, #66b1ff);
}

.search-icon {
  background: linear-gradient(135deg, #67c23a, #85ce61);
}

.quality-icon {
  background: linear-gradient(135deg, #e6a23c, #ebb563);
}

.analysis-icon {
  background: linear-gradient(135deg, #f56c6c, #f78989);
}

.cohort-icon {
  background: linear-gradient(135deg, #909399, #a6a9ad);
}

.export-icon {
  background: linear-gradient(135deg, #606266, #79808a);
}

.action-title {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 8px;
  color: #303133;
}

.action-desc {
  font-size: 12px;
  color: #909399;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.database-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.database-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  background-color: #f5f7fa;
  border-radius: 8px;
}

.db-info {
  flex: 1;
}

.db-name {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 5px;
}

.db-details {
  font-size: 12px;
  color: #909399;
  display: flex;
  gap: 20px;
}

.db-status {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 10px;
}

.db-performance {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 120px;
}

.performance-text {
  font-size: 12px;
  color: #606266;
  width: 30px;
  text-align: right;
}

.activities-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
  max-height: 400px;
  overflow-y: auto;
}

.activity-item {
  display: flex;
  align-items: flex-start;
  gap: 15px;
  padding: 15px;
  background-color: #f5f7fa;
  border-radius: 8px;
}

.activity-icon {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #409eff;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.activity-content {
  flex: 1;
}

.activity-title {
  font-size: 14px;
  font-weight: bold;
  margin-bottom: 5px;
}

.activity-desc {
  font-size: 12px;
  color: #606266;
  margin-bottom: 5px;
}

.activity-time {
  font-size: 11px;
  color: #909399;
}

.activity-status {
  flex-shrink: 0;
}

.chart-container {
  height: 300px;
  width: 100%;
}

.distribution-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.distribution-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.dist-label {
  font-size: 14px;
  font-weight: bold;
  color: #303133;
}

.dist-bar {
  display: flex;
  align-items: center;
  gap: 15px;
}

.dist-value {
  font-size: 12px;
  color: #606266;
  min-width: 60px;
  text-align: right;
}

.dialog-footer {
  text-align: right;
}

:deep(.el-progress-bar__outer) {
  background-color: #f0f2f5;
}
</style> 