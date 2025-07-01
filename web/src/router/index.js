import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import integrationRoutes from './modules/integration'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/user/Login.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/user/Register.vue')
  },
  // 临床管理路由
  {
    path: '/clinical',
    name: 'Clinical',
    children: [
      {
        path: 'patients',
        name: 'PatientList',
        component: () => import('@/views/clinical/PatientList.vue')
      },
      {
        path: 'patient/:id',
        name: 'PatientDetail',
        component: () => import('@/views/clinical/PatientDetail.vue')
      },
      {
        path: 'visits',
        name: 'VisitList',
        component: () => import('@/views/clinical/VisitList.vue')
      },
      {
        path: 'visit/:id',
        name: 'VisitDetail',
        component: () => import('@/views/clinical/VisitDetail.vue')
      },
      {
        path: 'diagnosis',
        name: 'DiagnosisList',
        component: () => import('@/views/clinical/DiagnosisList.vue')
      },
      {
        path: 'examination',
        name: 'ExaminationList',
        component: () => import('@/views/clinical/ExaminationList.vue')
      }
    ]
  },
  // 生物样本库路由
  {
    path: '/biobank',
    name: 'Biobank',
    children: [
      {
        path: 'samples',
        name: 'SampleList',
        component: () => import('@/views/biobank/SampleList.vue')
      },
      {
        path: 'sample/:id',
        name: 'SampleDetail',
        component: () => import('@/views/biobank/SampleDetail.vue')
      },
      {
        path: 'storage',
        name: 'Storage',
        component: () => import('@/views/biobank/Storage.vue')
      },
      {
        path: 'genomics',
        name: 'GenomicData',
        component: () => import('@/views/biobank/GenomicData.vue')
      },
      {
        path: 'proteomics',
        name: 'ProteomicsData',
        component: () => import('@/views/biobank/ProteomicsData.vue')
      }
    ]
  },
  // 流行病学路由
  {
    path: '/epidemiology',
    name: 'Epidemiology',
    children: [
      {
        path: 'exposure',
        name: 'EnvironmentExposure',
        component: () => import('@/views/epidemiology/EnvironmentExposure.vue')
      },
      {
        path: 'monitor',
        name: 'EnvironmentMonitor',
        component: () => import('@/views/epidemiology/EnvironmentMonitor.vue')
      },
      {
        path: 'monitoring',
        name: 'EnvironmentMonitoring',
        component: () => import('@/views/epidemiology/EnvironmentMonitor.vue')
      },
      {
        path: 'lifestyle',
        name: 'LifestyleSurvey',
        component: () => import('@/views/epidemiology/LifestyleSurvey.vue')
      }
    ]
  },
  // 随访管理路由
  {
    path: '/followup',
    name: 'Followup',
    children: [
      {
        path: 'plans',
        name: 'FollowupPlanList',
        component: () => import('@/views/followup/FollowupPlanList.vue')
      },
      {
        path: 'plan/:id',
        name: 'FollowupPlanDetail',
        component: () => import('@/views/followup/FollowupPlanDetail.vue')
      },
      {
        path: 'records',
        name: 'FollowupRecordList',
        component: () => import('@/views/followup/FollowupRecordList.vue')
      },
      {
        path: 'medication',
        name: 'MedicationMonitor',
        component: () => import('@/views/followup/MedicationMonitor.vue')
      },
      {
        path: 'medications',
        name: 'MedicationMonitorAlias',
        component: () => import('@/views/followup/MedicationMonitor.vue')
      }
    ]
  },
  // 数据集成路由
  ...integrationRoutes
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 