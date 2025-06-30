const integrationRoutes = [
  {
    path: '/integration',
    name: 'IntegrationDashboard',
    component: () => import('@/views/integration/Dashboard.vue'),
    meta: {
      title: '集成概览',
      icon: 'el-icon-data-analysis'
    }
  },
  {
    path: '/integration/patient-view',
    name: 'PatientIntegratedView',
    component: () => import('@/views/integration/PatientIntegratedView.vue'),
    meta: {
      title: '患者综合视图',
      icon: 'el-icon-user'
    }
  },
  {
    path: '/integration/data-quality',
    name: 'DataQualityReport',
    component: () => import('@/views/integration/DataQualityReport.vue'),
    meta: {
      title: '数据质量报告',
      icon: 'el-icon-document-checked'
    }
  },
  {
    path: '/integration/correlation-analysis',
    name: 'CorrelationAnalysis',
    component: () => import('@/views/integration/CorrelationAnalysis.vue'),
    meta: {
      title: '关联性分析',
      icon: 'el-icon-share'
    }
  },
  {
    path: '/integration/cohort-insights',
    name: 'CohortInsights',
    component: () => import('@/views/integration/CohortInsights.vue'),
    meta: {
      title: '队列洞察',
      icon: 'el-icon-pie-chart'
    }
  },
  {
    path: '/integration/cross-database-search',
    name: 'CrossDatabaseSearch',
    component: () => import('@/views/integration/CrossDatabaseSearch.vue'),
    meta: {
      title: '跨库搜索',
      icon: 'el-icon-search'
    }
  }
]

export default integrationRoutes 