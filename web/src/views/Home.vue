<template>
  <div class="home">
    <!-- 简洁现代导航栏 -->
    <header class="navbar">
      <div class="navbar-container">
        <div class="navbar-brand">
          <h1>过敏疾病管理系统</h1>
          <span class="version">数据库大作业</span>
        </div>
        
        <nav class="navbar-nav">
          <el-dropdown class="nav-item" @command="handleNavigation">
            <span class="nav-link">
              <el-icon><User /></el-icon>
              用户管理
              <el-icon class="arrow"><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="/login">用户登录</el-dropdown-item>
                <el-dropdown-item command="/register">用户注册</el-dropdown-item>
                <el-dropdown-item command="/profile">个人信息</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <el-dropdown class="nav-item" @command="handleNavigation">
            <span class="nav-link">
              <el-icon><OfficeBuilding /></el-icon>
              临床管理
              <el-icon class="arrow"><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="/clinical/patients">患者管理</el-dropdown-item>
                <el-dropdown-item command="/clinical/visits">就诊记录</el-dropdown-item>
                <el-dropdown-item command="/clinical/diagnosis">诊断管理</el-dropdown-item>
                <el-dropdown-item command="/clinical/examination">检查管理</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <el-dropdown class="nav-item" @command="handleNavigation">
            <span class="nav-link">
              <el-icon><Cpu /></el-icon>
              生物样本库
              <el-icon class="arrow"><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="/biobank/samples">样本管理</el-dropdown-item>
                <el-dropdown-item command="/biobank/storage">存储管理</el-dropdown-item>
                <el-dropdown-item command="/biobank/genomics">基因组学</el-dropdown-item>
                <el-dropdown-item command="/biobank/proteomics">蛋白组学</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <el-dropdown class="nav-item" @command="handleNavigation">
            <span class="nav-link">
              <el-icon><MapLocation /></el-icon>
              流行病学
              <el-icon class="arrow"><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="/epidemiology/exposure">环境暴露</el-dropdown-item>
                <el-dropdown-item command="/epidemiology/monitoring">环境监测</el-dropdown-item>
                <el-dropdown-item command="/epidemiology/lifestyle">生活方式</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <el-dropdown class="nav-item" @command="handleNavigation">
            <span class="nav-link">
              <el-icon><Document /></el-icon>
              随访管理
              <el-icon class="arrow"><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="/followup/plans">随访计划</el-dropdown-item>
                <el-dropdown-item command="/followup/records">随访记录</el-dropdown-item>
                <el-dropdown-item command="/followup/medications">用药监测</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <el-dropdown class="nav-item" @command="handleNavigation">
            <span class="nav-link">
              <el-icon><Connection /></el-icon>
              数据集成
              <el-icon class="arrow"><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="/integration">集成仪表板</el-dropdown-item>
                <el-dropdown-item command="/integration/patient-view">患者综合视图</el-dropdown-item>
                <el-dropdown-item command="/integration/cross-database-search">跨库搜索</el-dropdown-item>
                <el-dropdown-item command="/integration/data-quality">数据质量报告</el-dropdown-item>
                <el-dropdown-item command="/integration/correlation-analysis">关联分析</el-dropdown-item>
                <el-dropdown-item command="/integration/cohort-insights">队列洞察</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </nav>
      </div>
    </header>

    <!-- 主要内容区域 -->
    <main class="main-content">
      <!-- 数据概览 -->
      <section class="stats-section">
        <div class="stats-grid">
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><UserFilled /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-value">{{ systemStats.totalPatients }}</div>
              <div class="stat-label">总患者数</div>
              <div class="stat-change" :class="systemStats.patientGrowth.includes('+') ? 'positive' : 'negative'">
                <el-icon>
                  <CaretTop v-if="systemStats.patientGrowth.includes('+')" />
                  <CaretBottom v-else />
                </el-icon>
                {{ systemStats.patientGrowth }}
              </div>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><DataBoard /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-value">{{ systemStats.totalRecords }}</div>
              <div class="stat-label">数据记录</div>
              <div class="stat-change" :class="systemStats.recordsGrowth.includes('+') ? 'positive' : 'negative'">
                <el-icon>
                  <CaretTop v-if="systemStats.recordsGrowth.includes('+')" />
                  <CaretBottom v-else />
                </el-icon>
                {{ systemStats.recordsGrowth }}
              </div>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><Checked /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-value">{{ formatQuality(systemStats.dataQuality) }}%</div>
              <div class="stat-label">数据质量</div>
              <div class="stat-change" :class="systemStats.qualityTrend.includes('+') ? 'positive' : 'negative'">
                <el-icon>
                  <CaretTop v-if="systemStats.qualityTrend.includes('+')" />
                  <CaretBottom v-else />
                </el-icon>
                {{ systemStats.qualityTrend }}
              </div>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><User /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-value">{{ systemStats.activeUsers }}</div>
              <div class="stat-label">活跃用户</div>
              <div class="stat-change" :class="systemStats.usersGrowth.includes('+') ? 'positive' : 'negative'">
                <el-icon>
                  <CaretTop v-if="systemStats.usersGrowth.includes('+')" />
                  <CaretBottom v-else />
                </el-icon>
                {{ systemStats.usersGrowth }}
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- 快速操作 -->
      <section class="actions-section">
        <h2>快速操作</h2>
        <div class="actions-grid">
          <div class="action-card" @click="$router.push('/integration')">
            <el-icon class="action-icon"><DataBoard /></el-icon>
            <h3>数据集成仪表板</h3>
            <p>跨数据库统一视图和数据质量监控</p>
          </div>

          <div class="action-card" @click="$router.push('/integration/patient-view')">
            <el-icon class="action-icon"><UserFilled /></el-icon>
            <h3>患者综合视图</h3>
            <p>全维度患者信息整合展示</p>
          </div>

          <div class="action-card" @click="$router.push('/integration/cross-database-search')">
            <el-icon class="action-icon"><Search /></el-icon>
            <h3>跨库搜索</h3>
            <p>多数据库联合查询和分析</p>
          </div>

          <div class="action-card" @click="$router.push('/integration/data-quality')">
            <el-icon class="action-icon"><Checked /></el-icon>
            <h3>数据质量报告</h3>
            <p>数据完整性和准确性评估</p>
          </div>

          <div class="action-card" @click="$router.push('/integration/correlation-analysis')">
            <el-icon class="action-icon"><TrendCharts /></el-icon>
            <h3>关联分析</h3>
            <p>多维度数据关联性分析</p>
          </div>

          <div class="action-card" @click="$router.push('/integration/cohort-insights')">
            <el-icon class="action-icon"><Histogram /></el-icon>
            <h3>队列洞察</h3>
            <p>患者群体特征和趋势分析</p>
          </div>
        </div>
      </section>

      <!-- 最近活动 -->
      <section class="activity-section">
        <h2>最近活动</h2>
        <div class="activity-list" v-loading="loadingActivities">
          <div class="activity-item" v-for="activity in recentActivities" :key="activity.id">
            <div class="activity-time">{{ formatTime(activity.created_at) }}</div>
            <div class="activity-content">
              <div class="activity-title">{{ activity.title }}</div>
              <div class="activity-desc">{{ activity.description }}</div>
            </div>
          </div>
          <div v-if="recentActivities.length === 0" class="empty-state">
            暂无最近活动
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { 
  User, OfficeBuilding, Cpu, MapLocation, Document, Connection, DataBoard, 
  UserFilled, Search, Checked, TrendCharts, Histogram, ArrowRight,
  CaretTop, CaretBottom, ArrowDown
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { integrationApi } from '@/api/integration'

const router = useRouter()

// 系统统计数据
const systemStats = reactive({
  totalPatients: 0,
  patientGrowth: '+0%',
  totalRecords: 0, 
  recordsGrowth: '+0%',
  dataQuality: 0,
  qualityTrend: '+0%',
  activeUsers: 0,
  usersGrowth: '+0%'
})

// 最近活动数据
const recentActivities = ref([])
const loadingActivities = ref(false)

// 导航处理
const handleNavigation = (command) => {
  router.push(command)
}

// 格式化时间
const formatTime = (timestamp) => {
  if (!timestamp) return ''
  const date = new Date(timestamp)
  const now = new Date()
  const diff = now - date
  
  if (diff < 60000) return '刚刚'
  if (diff < 3600000) return `${Math.floor(diff / 60000)}分钟前`
  if (diff < 86400000) return `${Math.floor(diff / 3600000)}小时前`
  if (diff < 604800000) return `${Math.floor(diff / 86400000)}天前`
  
  return date.toLocaleDateString()
}

// 格式化数据质量
const formatQuality = (quality) => {
  return parseFloat(quality).toFixed(2)
}

// 加载系统统计数据
const loadSystemStats = async () => {
  try {
    console.log('🔄 开始加载后端统计数据...')
    
    // 只调用主要的统计API，它已经包含了所有需要的数据
    const statsRes = await integrationApi.getDashboardStats()
    
    console.log('📊 后端返回的统计数据:', statsRes)

    // 更新统计数据
    if (statsRes.code === '10000' && statsRes.data) {
      const data = statsRes.data
      
      // 使用后端返回的真实数据
      systemStats.totalPatients = data.total_patients || 0
      systemStats.patientGrowth = data.patient_growth || '+0%'
      systemStats.totalRecords = data.total_records || 0
      systemStats.recordsGrowth = data.records_growth || '+0%'
      systemStats.dataQuality = data.data_quality || 0
      systemStats.qualityTrend = data.quality_trend || '+0%'
      systemStats.activeUsers = data.active_users || 0
      systemStats.usersGrowth = data.users_growth || '+0%'
      
      console.log('✅ 统计数据加载成功，来自后端真实数据!')
      console.log('📈 当前显示数据:', {
        patients: systemStats.totalPatients,
        records: systemStats.totalRecords,
        quality: systemStats.dataQuality,
        users: systemStats.activeUsers
      })
    } else {
      console.warn('⚠️ 后端返回数据格式异常:', statsRes)
      throw new Error('数据格式异常')
    }
  } catch (error) {
    console.error('❌ 加载统计数据失败:', error)
    ElMessage.error('无法连接到后端服务，请检查网络连接')
    
    // 设置默认值，明确标识为备用数据
    systemStats.totalPatients = 0
    systemStats.patientGrowth = '+0%'
    systemStats.totalRecords = 0
    systemStats.recordsGrowth = '+0%'
    systemStats.dataQuality = 0
    systemStats.qualityTrend = '+0%'
    systemStats.activeUsers = 0
    systemStats.usersGrowth = '+0%'
    
    console.log('🔄 已设置为默认值，等待后端服务恢复')
  }
}

// 加载最近活动
const loadRecentActivities = async () => {
  loadingActivities.value = true
  try {
    console.log('🔄 开始加载最近活动数据...')
    
    const response = await integrationApi.getRecentActivities(5)
    
    console.log('📋 后端返回的活动数据:', response)
    
    if (response.code === '10000' && response.data) {
      // 检查数据结构，适配不同的返回格式
      const activities = response.data.activities || response.data || []
      recentActivities.value = activities
      
      console.log('✅ 活动数据加载成功，来自后端真实数据!')
      console.log('📝 活动列表:', activities)
    } else {
      console.warn('⚠️ 后端返回活动数据格式异常:', response)
      throw new Error('活动数据格式异常')
    }
  } catch (error) {
    console.error('❌ 加载活动数据失败:', error)
    ElMessage.warning('无法加载最近活动数据')
    
    // 清空活动列表，不使用模拟数据
    recentActivities.value = []
    console.log('🔄 已清空活动列表，等待后端服务恢复')
  } finally {
    loadingActivities.value = false
  }
}

// 页面加载时获取数据
onMounted(async () => {
  await Promise.all([
    loadSystemStats(),
    loadRecentActivities()
  ])
})
</script>

<style scoped>
.home {
  min-height: 100vh;
  background: #f5f7fa;
}

/* 导航栏 */
.navbar {
  background: #fff;
  border-bottom: 1px solid #e4e7ed;
  position: sticky;
  top: 0;
  z-index: 1000;
}

.navbar-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 64px;
}

.navbar-brand h1 {
  color: #303133;
  font-size: 20px;
  font-weight: 600;
  margin: 0;
}

.version {
  background: #409eff;
  color: white;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
  margin-left: 8px;
}

.navbar-nav {
  display: flex;
  align-items: center;
  gap: 32px;
}

.nav-item {
  cursor: pointer;
}

.nav-item:focus,
.nav-item:focus-visible {
  outline: none !important;
  box-shadow: none !important;
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 6px;
  color: #606266;
  font-size: 14px;
  padding: 8px 0;
  transition: color 0.3s;
  outline: none !important;
}

.nav-link:hover {
  color: #409eff;
}

.nav-link:focus,
.nav-link:focus-visible {
  outline: none !important;
  box-shadow: none !important;
}

.nav-link .arrow {
  font-size: 12px;
  transition: transform 0.3s;
}

/* 去除Element Plus下拉框的默认focus样式 */
.el-dropdown {
  outline: none !important;
}

.el-dropdown:focus,
.el-dropdown:focus-visible {
  outline: none !important;
  box-shadow: none !important;
}

.el-dropdown-link {
  outline: none !important;
}

.el-dropdown-link:focus,
.el-dropdown-link:focus-visible {
  outline: none !important;
  box-shadow: none !important;
}

/* 主要内容 */
.main-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 32px 24px;
}

/* 统计卡片 */
.stats-section {
  margin-bottom: 32px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 20px;
}

.stat-card {
  background: #fff;
  border-radius: 8px;
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 16px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.04);
  transition: transform 0.3s, box-shadow 0.3s;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  background: #f0f9ff;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #409eff;
  font-size: 24px;
}

.stat-content {
  flex: 1;
}

.stat-value {
  font-size: 28px;
  font-weight: 600;
  color: #303133;
  line-height: 1;
}

.stat-label {
  color: #909399;
  font-size: 14px;
  margin: 4px 0;
}

.stat-change {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  font-weight: 500;
}

.stat-change.positive {
  color: #67c23a;
}

.stat-change.negative {
  color: #f56c6c;
}

/* 快速操作 */
.actions-section {
  margin-bottom: 32px;
}

.actions-section h2 {
  color: #303133;
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 16px;
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 16px;
}

.action-card {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.04);
}

.action-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
}

.action-icon {
  font-size: 32px;
  color: #409eff;
  margin-bottom: 12px;
}

.action-card h3 {
  color: #303133;
  font-size: 16px;
  font-weight: 600;
  margin: 0 0 8px 0;
}

.action-card p {
  color: #909399;
  font-size: 14px;
  margin: 0;
  line-height: 1.5;
}

/* 最近活动 */
.activity-section h2 {
  color: #303133;
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 16px;
}

.activity-list {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.04);
  min-height: 200px;
}

.activity-item {
  display: flex;
  gap: 16px;
  padding: 16px 20px;
  border-bottom: 1px solid #f0f2f5;
}

.activity-item:last-child {
  border-bottom: none;
}

.activity-time {
  color: #909399;
  font-size: 12px;
  white-space: nowrap;
  min-width: 80px;
}

.activity-content {
  flex: 1;
}

.activity-title {
  color: #303133;
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 4px;
}

.activity-desc {
  color: #606266;
  font-size: 13px;
  line-height: 1.4;
}

.empty-state {
  text-align: center;
  color: #909399;
  padding: 40px;
  font-size: 14px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .navbar-container {
    flex-direction: column;
    height: auto;
    padding: 16px;
  }
  
  .navbar-nav {
    flex-wrap: wrap;
    justify-content: center;
    gap: 16px;
    margin-top: 16px;
  }
  
  .main-content {
    padding: 16px;
  }
  
  .stats-grid {
    grid-template-columns: 1fr;
  }
  
  .actions-grid {
    grid-template-columns: 1fr;
  }
}
</style> 