<template>
  <div class="patient-integrated-view">
    <el-card class="header-card">
      <div class="header-content">
        <h2>患者综合视图</h2>
        <el-input
          v-model="patientNo"
          placeholder="请输入患者编号"
          style="width: 200px; margin-right: 10px"
        />
        <el-button type="primary" @click="loadPatientData">查询</el-button>
      </div>
    </el-card>

    <div v-if="patientData" class="content-area">
      <!-- 患者基本信息 -->
      <el-card class="info-card">
        <template #header>
          <div class="card-header">
            <span>基本信息</span>
            <el-tag :type="getQualityTagType(patientData.statistics?.data_quality_score)">
              数据质量: {{ patientData.statistics?.data_quality_score || 0 }}%
            </el-tag>
          </div>
        </template>
        <el-descriptions :column="4" border>
          <el-descriptions-item label="患者编号">{{ patientData.patient?.patient_no }}</el-descriptions-item>
          <el-descriptions-item label="姓名">{{ patientData.patient?.name }}</el-descriptions-item>
          <el-descriptions-item label="性别">{{ getGenderText(patientData.patient?.gender) }}</el-descriptions-item>
          <el-descriptions-item label="年龄">{{ calculateAge(patientData.patient?.birth_date) }}</el-descriptions-item>
          <el-descriptions-item label="联系电话">{{ patientData.patient?.phone }}</el-descriptions-item>
          <el-descriptions-item label="身份证号">{{ patientData.patient?.id_card }}</el-descriptions-item>
          <el-descriptions-item label="地址" :span="2">{{ patientData.patient?.address }}</el-descriptions-item>
        </el-descriptions>
      </el-card>

      <!-- 统计概览 -->
      <el-card class="stats-card">
        <template #header>
          <span>数据统计</span>
        </template>
        <el-row :gutter="20">
          <el-col :span="6">
            <div class="stat-item">
              <div class="stat-number">{{ patientData.statistics?.total_visits || 0 }}</div>
              <div class="stat-label">就诊次数</div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="stat-item">
              <div class="stat-number">{{ patientData.statistics?.total_diagnoses || 0 }}</div>
              <div class="stat-label">诊断记录</div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="stat-item">
              <div class="stat-number">{{ patientData.statistics?.total_followups || 0 }}</div>
              <div class="stat-label">随访记录</div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="stat-item">
              <div class="stat-number">{{ patientData.statistics?.total_samples || 0 }}</div>
              <div class="stat-label">生物样本</div>
            </div>
          </el-col>
        </el-row>
      </el-card>

      <!-- 详细数据标签页 -->
      <el-card class="detail-card" v-loading="loading">
        <el-tabs v-model="activeTab" type="card">
          <!-- 临床数据 -->
          <el-tab-pane label="临床数据" name="clinical">
            <div class="tab-content">
              <h3>就诊记录</h3>
              <el-table :data="patientData.visits" stripe>
                <el-table-column prop="visit_no" label="就诊号" width="120" />
                <el-table-column prop="visit_time" label="就诊时间" width="150">
                  <template #default="scope">
                    {{ formatDate(scope.row.visit_time) }}
                  </template>
                </el-table-column>
                <el-table-column prop="department" label="科室" width="100" />
                <el-table-column prop="doctor_id" label="医生ID" width="100" />
                <el-table-column prop="chief_complaint" label="主诉" />
              </el-table>

              <h3 style="margin-top: 20px">诊断记录</h3>
              <el-table :data="patientData.diagnoses" stripe>
                <el-table-column prop="icd_code" label="ICD编码" width="120" />
                <el-table-column prop="diagnosis_type" label="诊断类型" width="100" />
                <el-table-column prop="created_at" label="诊断时间" width="150">
                  <template #default="scope">
                    {{ formatDate(scope.row.created_at) }}
                  </template>
                </el-table-column>
              </el-table>

              <h3 style="margin-top: 20px">检查记录</h3>
              <el-table :data="patientData.examinations" stripe>
                <el-table-column prop="exam_type" label="检查类型" />
                <el-table-column prop="exam_time" label="检查时间" width="150">
                  <template #default="scope">
                    {{ formatDate(scope.row.exam_time) }}
                  </template>
                </el-table-column>
                <el-table-column prop="exam_result" label="检查结果" />
                <el-table-column prop="result_unit" label="结果单位" />
                <el-table-column prop="reference_range" label="参考范围" />
                <el-table-column prop="created_at" label="录入时间" width="150">
                  <template #default="scope">
                    {{ formatDate(scope.row.created_at) }}
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </el-tab-pane>

          <!-- 流调数据 -->
          <el-tab-pane label="流调数据" name="epidemiology">
            <div class="tab-content">
              <h3>环境暴露</h3>
              <el-table :data="patientData.environment_exposures" stripe>
                <el-table-column prop="residence_type" label="居住类型" />
                <el-table-column prop="building_material" label="建筑材料" />
                <el-table-column prop="pet_exposure" label="宠物暴露">
                  <template #default="scope">
                    {{ scope.row.pet_exposure ? '是' : '否' }}
                  </template>
                </el-table-column>
                <el-table-column prop="smoking_exposure" label="吸烟暴露">
                  <template #default="scope">
                    {{ scope.row.smoking_exposure ? '是' : '否' }}
                  </template>
                </el-table-column>
              </el-table>

              <h3 style="margin-top: 20px">生活方式</h3>
              <el-table :data="patientData.lifestyle_surveys" stripe>
                <el-table-column prop="diet_pattern" label="饮食模式" />
                <el-table-column prop="exercise_frequency" label="运动频率" />
                <el-table-column prop="sleep_quality" label="睡眠质量" />
                <el-table-column prop="stress_level" label="压力水平" />
              </el-table>
            </div>
          </el-tab-pane>

          <!-- 随访数据 -->
          <el-tab-pane label="随访数据" name="followup">
            <div class="tab-content">
              <h3>随访计划</h3>
              <el-table :data="patientData.followup_plans" stripe>
                <el-table-column prop="plan_name" label="计划名称" />
                <el-table-column prop="plan_type" label="计划类型" />
                <el-table-column prop="start_date" label="开始日期" width="120">
                  <template #default="scope">
                    {{ formatDate(scope.row.start_date) }}
                  </template>
                </el-table-column>
                <el-table-column prop="end_date" label="结束日期" width="120">
                  <template #default="scope">
                    {{ formatDate(scope.row.end_date) }}
                  </template>
                </el-table-column>
                <el-table-column prop="frequency" label="频率" />
                <el-table-column prop="status" label="状态" />
              </el-table>

              <h3 style="margin-top: 20px">随访记录</h3>
              <el-table :data="patientData.followup_records" stripe>
                <el-table-column prop="followup_date" label="随访日期" width="150">
                  <template #default="scope">
                    {{ formatDate(scope.row.followup_date) }}
                  </template>
                </el-table-column>
                <el-table-column prop="followup_method" label="随访方式" width="100" />
                <el-table-column prop="symptoms" label="症状" />
                <el-table-column prop="treatment_effect" label="治疗效果" />
                <el-table-column prop="followup_doctor" label="随访医生" width="100" />
              </el-table>

              <h3 style="margin-top: 20px">用药监测</h3>
              <el-table :data="patientData.medications" stripe>
                <el-table-column prop="medication_name" label="药物名称" />
                <el-table-column prop="dosage" label="剂量" width="100" />
                <el-table-column prop="frequency" label="频次" width="100" />
                <el-table-column prop="compliance" label="依从性" width="100" />
                <el-table-column prop="side_effects" label="不良反应" />
              </el-table>
            </div>
          </el-tab-pane>

          <!-- 生物样本 -->
          <el-tab-pane label="生物样本" name="biobank">
            <div class="tab-content">
              <h3>样本信息</h3>
              <el-table :data="patientData.samples" stripe>
                <el-table-column prop="sample_no" label="样本编号" width="120" />
                <el-table-column prop="sample_type" label="样本类型" width="100" />
                <el-table-column prop="collection_time" label="采集时间" width="150">
                  <template #default="scope">
                    {{ formatDate(scope.row.collection_time) }}
                  </template>
                </el-table-column>
                <el-table-column prop="collection_method" label="采集方法" />
                <el-table-column prop="processor" label="处理人" width="100" />
              </el-table>

              <h3 style="margin-top: 20px">基因组数据</h3>
              <el-table :data="patientData.genomic_data" stripe>
                <el-table-column prop="sequence_platform" label="测序平台" />
                <el-table-column prop="sequence_type" label="测序类型" width="100" />
                <el-table-column prop="gene_type" label="基因类型" width="100" />
                <el-table-column prop="result_file_path" label="结果文件" />
              </el-table>

              <h3 style="margin-top: 20px">蛋白组学数据</h3>
              <el-table :data="patientData.proteomics_data" stripe>
                <el-table-column prop="analysis_platform" label="分析平台" />
                <el-table-column prop="protein_marker" label="蛋白标志物" />
                <el-table-column prop="concentration" label="浓度" />
              </el-table>
            </div>
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { integrationApi } from '@/api/integration'

export default {
  name: 'PatientIntegratedView',
  setup() {
    const patientNo = ref('')
    const patientData = ref(null)
    const loading = ref(false)
    const activeTab = ref('clinical')

    const mockData = {
      patient: {
        patient_no: 'P001',
        name: '张三',
        gender: '1',
        birth_date: '1985-05-15',
        phone: '13800138001',
        id_card: '110101198505150011',
        address: '北京市海淀区中关村大街123号',
        height: 175.5,
        weight: 70.2,
        birth_weight: 3.2,
        medical_insurance_type: 1,
        medical_insurance_no: 'YB001234567',
        emergency_contact: '张小三',
        emergency_phone: '13900139001'
      },
      statistics: {
        total_visits: 12,
        total_diagnoses: 5,
        total_followups: 3,
        total_samples: 4,
        data_quality_score: 100,
        active_medications: 2,
        first_visit_date: '2024-01-01',
        last_visit_date: '2024-03-15',
        first_diagnosis_date: '2024-01-01',
        last_diagnosis_date: '2024-03-10',
        total_examinations: 8,
        first_exam_date: '2024-01-01',
        last_exam_date: '2024-03-12',
        last_followup_date: '2024-03-18',
        next_followup_date: '2024-04-01',
        first_sample_date: '2024-01-01',
        last_sample_date: '2024-03-15',
        total_genomic: 2,
        first_genomic_date: '2024-01-15',
        last_genomic_date: '2024-02-20',
        total_proteomics: 3,
        first_proteomics_date: '2024-01-20',
        last_proteomics_date: '2024-03-01'
      },
      visits: [
        { visit_no: 'V001', visit_time: '2024-01-01', department: '内科', doctor_id: 101, chief_complaint: '咳嗽' },
        { visit_no: 'V002', visit_time: '2024-02-15', department: '外科', doctor_id: 102, chief_complaint: '腹痛' },
        { visit_no: 'V003', visit_time: '2024-03-10', department: '儿科', doctor_id: 103, chief_complaint: '发热' }
      ],
      diagnoses: [
        { icd_code: 'J20', diagnosis_type: '急性支气管炎', created_at: '2024-01-01' },
        { icd_code: 'K35', diagnosis_type: '急性阑尾炎', created_at: '2024-02-15' },
        { icd_code: 'J45', diagnosis_type: '支气管哮喘', created_at: '2024-03-10' }
      ],
      examinations: [
        { exam_type: '血常规', exam_time: '2024-01-01', exam_result: '正常', result_unit: '', reference_range: '4-10', created_at: '2024-01-01' },
        { exam_type: '腹部B超', exam_time: '2024-02-15', exam_result: '阑尾肿大', result_unit: '', reference_range: '', created_at: '2024-02-15' },
        { exam_type: '胸部X线', exam_time: '2024-03-10', exam_result: '肺部感染', result_unit: '', reference_range: '', created_at: '2024-03-10' }
      ],
      followup_plans: [
        {
          plan_name: '术后随访',
          plan_type: '电话',
          start_date: '2024-03-11',
          end_date: '2024-04-11',
          frequency: '每周',
          status: '进行中'
        },
        {
          plan_name: '慢病管理',
          plan_type: '门诊',
          start_date: '2024-01-02',
          end_date: '2024-06-02',
          frequency: '每月',
          status: '已完成'
        }
      ],
      followup_records: [
        {
          followup_date: '2024-03-18',
          followup_method: '电话',
          symptoms: '恢复良好，无明显不适',
          treatment_effect: '效果良好',
          followup_doctor: '王医生'
        },
        {
          followup_date: '2024-04-01',
          followup_method: '门诊',
          symptoms: '偶有咳嗽',
          treatment_effect: '基本控制',
          followup_doctor: '李医生'
        }
      ],
      medications: [
        {
          medication_name: '氯雷他定',
          dosage: '10mg',
          frequency: '每日一次',
          compliance: '良好',
          side_effects: '轻微嗜睡'
        },
        {
          medication_name: '布地奈德鼻喷剂',
          dosage: '64μg',
          frequency: '每日两次',
          compliance: '良好',
          side_effects: '无'
        }
      ],
      environment_exposures: [
        {
          residence_type: '公寓',
          building_material: '钢筋混凝土',
          pet_exposure: true,
          smoking_exposure: false
        }
      ],
      lifestyle_surveys: [
        {
          diet_pattern: '均衡饮食',
          exercise_frequency: '每周3-4次',
          sleep_quality: '良好',
          stress_level: '中等'
        }
      ],
      samples: [
        {
          sample_no: 'S001',
          sample_type: '血液',
          collection_time: '2024-01-01',
          collection_method: '静脉采血',
          processor: '护士A'
        },
        {
          sample_no: 'S002',
          sample_type: '血清',
          collection_time: '2024-01-02',
          collection_method: '离心分离',
          processor: '技师B'
        },
        {
          sample_no: 'S003',
          sample_type: '尿液',
          collection_time: '2024-01-03',
          collection_method: '自取',
          processor: '护士A'
        },
        {
          sample_no: 'S004',
          sample_type: '唾液',
          collection_time: '2024-01-04',
          collection_method: '自取',
          processor: '护士A'
        }
      ],
      genomic_data: [
        {
          sequence_platform: 'Illumina NovaSeq 6000',
          sequence_type: '全基因组测序',
          gene_type: '过敏相关基因',
          result_file_path: '/data/genomics/S001_WGS.vcf'
        },
        {
          sequence_platform: 'BGI MGISEQ-2000',
          sequence_type: '外显子测序',
          gene_type: '免疫相关基因',
          result_file_path: '/data/genomics/S002_WES.vcf'
        },
        {
          sequence_platform: 'Thermo Fisher Ion Proton',
          sequence_type: '靶向测序',
          gene_type: '药物代谢基因',
          result_file_path: '/data/genomics/S003_Target.vcf'
        }
      ],
      proteomics_data: [
        {
          analysis_platform: '质谱仪',
          protein_marker: 'IgE',
          concentration: 120.5
        },
        {
          analysis_platform: '质谱仪',
          protein_marker: 'IL-4',
          concentration: 25.3
        },
        {
          analysis_platform: '质谱仪',
          protein_marker: 'IL-13',
          concentration: 18.7
        }
      ]
    }

    // 多组mock数据
    const mockDataList = [
      {
        patient: {
          patient_no: 'P001',
          name: '张三',
          gender: '1',
          birth_date: '1985-05-15',
          phone: '13800138001',
          id_card: '110101198505150011',
          address: '北京市海淀区中关村大街123号',
          height: 175.5,
          weight: 70.2,
          birth_weight: 3.2,
          medical_insurance_type: 1,
          medical_insurance_no: 'YB001234567',
          emergency_contact: '张小三',
          emergency_phone: '13900139001'
        },
        statistics: {
          total_visits: 12,
          total_diagnoses: 5,
          total_followups: 3,
          total_samples: 4,
          data_quality_score: 100,
          active_medications: 2,
          first_visit_date: '2024-01-01',
          last_visit_date: '2024-03-15',
          first_diagnosis_date: '2024-01-01',
          last_diagnosis_date: '2024-03-10',
          total_examinations: 8,
          first_exam_date: '2024-01-01',
          last_exam_date: '2024-03-12',
          last_followup_date: '2024-03-18',
          next_followup_date: '2024-04-01',
          first_sample_date: '2024-01-01',
          last_sample_date: '2024-03-15',
          total_genomic: 2,
          first_genomic_date: '2024-01-15',
          last_genomic_date: '2024-02-20',
          total_proteomics: 3,
          first_proteomics_date: '2024-01-20',
          last_proteomics_date: '2024-03-01'
        },
        visits: [
          { visit_no: 'V001', visit_time: '2024-01-01', department: '内科', doctor_id: 101, chief_complaint: '咳嗽' },
          { visit_no: 'V002', visit_time: '2024-02-15', department: '外科', doctor_id: 102, chief_complaint: '腹痛' },
          { visit_no: 'V003', visit_time: '2024-03-10', department: '儿科', doctor_id: 103, chief_complaint: '发热' }
        ],
        diagnoses: [
          { icd_code: 'J20', diagnosis_type: '急性支气管炎', created_at: '2024-01-01' },
          { icd_code: 'K35', diagnosis_type: '急性阑尾炎', created_at: '2024-02-15' },
          { icd_code: 'J45', diagnosis_type: '支气管哮喘', created_at: '2024-03-10' }
        ],
        examinations: [
          { exam_type: '血常规', exam_time: '2024-01-01', exam_result: '正常', result_unit: '', reference_range: '4-10', created_at: '2024-01-01' },
          { exam_type: '腹部B超', exam_time: '2024-02-15', exam_result: '阑尾肿大', result_unit: '', reference_range: '', created_at: '2024-02-15' },
          { exam_type: '胸部X线', exam_time: '2024-03-10', exam_result: '肺部感染', result_unit: '', reference_range: '', created_at: '2024-03-10' }
        ],
        followup_plans: [
          {
            plan_name: '术后随访',
            plan_type: '电话',
            start_date: '2024-03-11',
            end_date: '2024-04-11',
            frequency: '每周',
            status: '进行中'
          },
          {
            plan_name: '慢病管理',
            plan_type: '门诊',
            start_date: '2024-01-02',
            end_date: '2024-06-02',
            frequency: '每月',
            status: '已完成'
          }
        ],
        followup_records: [
          {
            followup_date: '2024-03-18',
            followup_method: '电话',
            symptoms: '恢复良好，无明显不适',
            treatment_effect: '效果良好',
            followup_doctor: '王医生'
          },
          {
            followup_date: '2024-04-01',
            followup_method: '门诊',
            symptoms: '偶有咳嗽',
            treatment_effect: '基本控制',
            followup_doctor: '李医生'
          }
        ],
        medications: [
          {
            medication_name: '氯雷他定',
            dosage: '10mg',
            frequency: '每日一次',
            compliance: '良好',
            side_effects: '轻微嗜睡'
          },
          {
            medication_name: '布地奈德鼻喷剂',
            dosage: '64μg',
            frequency: '每日两次',
            compliance: '良好',
            side_effects: '无'
          }
        ],
        environment_exposures: [
          {
            residence_type: '公寓',
            building_material: '钢筋混凝土',
            pet_exposure: true,
            smoking_exposure: false
          }
        ],
        lifestyle_surveys: [
          {
            diet_pattern: '均衡饮食',
            exercise_frequency: '每周3-4次',
            sleep_quality: '良好',
            stress_level: '中等'
          }
        ],
        samples: [
          {
            sample_no: 'S001',
            sample_type: '血液',
            collection_time: '2024-01-01',
            collection_method: '静脉采血',
            processor: '护士A'
          },
          {
            sample_no: 'S002',
            sample_type: '血清',
            collection_time: '2024-01-02',
            collection_method: '离心分离',
            processor: '技师B'
          },
          {
            sample_no: 'S003',
            sample_type: '尿液',
            collection_time: '2024-01-03',
            collection_method: '自取',
            processor: '护士A'
          },
          {
            sample_no: 'S004',
            sample_type: '唾液',
            collection_time: '2024-01-04',
            collection_method: '自取',
            processor: '护士A'
          }
        ],
        genomic_data: [
          {
            sequence_platform: 'Illumina NovaSeq 6000',
            sequence_type: '全基因组测序',
            gene_type: '过敏相关基因',
            result_file_path: '/data/genomics/S001_WGS.vcf'
          },
          {
            sequence_platform: 'BGI MGISEQ-2000',
            sequence_type: '外显子测序',
            gene_type: '免疫相关基因',
            result_file_path: '/data/genomics/S002_WES.vcf'
          },
          {
            sequence_platform: 'Thermo Fisher Ion Proton',
            sequence_type: '靶向测序',
            gene_type: '药物代谢基因',
            result_file_path: '/data/genomics/S003_Target.vcf'
          }
        ],
        proteomics_data: [
          {
            analysis_platform: '质谱仪',
            protein_marker: 'IgE',
            concentration: 120.5
          },
          {
            analysis_platform: '质谱仪',
            protein_marker: 'IL-4',
            concentration: 25.3
          },
          {
            analysis_platform: '质谱仪',
            protein_marker: 'IL-13',
            concentration: 18.7
          }
        ]
      },
      {
        patient: {
          patient_no: 'P002',
          name: '李四',
          gender: '0',
          birth_date: '1990-08-22',
          phone: '13800138002',
          id_card: '110101199008220022',
          address: '北京市朝阳区建国路456号',
          height: 165.0,
          weight: 55.8,
          birth_weight: 2.8,
          medical_insurance_type: 2,
          medical_insurance_no: 'YB002345678',
          emergency_contact: '李大四',
          emergency_phone: '13900139002'
        },
        statistics: {
          total_visits: 8,
          total_diagnoses: 3,
          total_followups: 2,
          total_samples: 2,
          data_quality_score: 95,
          active_medications: 1,
          first_visit_date: '2024-02-01',
          last_visit_date: '2024-04-10',
          first_diagnosis_date: '2024-02-01',
          last_diagnosis_date: '2024-04-05',
          total_examinations: 5,
          first_exam_date: '2024-02-01',
          last_exam_date: '2024-04-08',
          last_followup_date: '2024-04-15',
          next_followup_date: '2024-05-01',
          first_sample_date: '2024-02-15',
          last_sample_date: '2024-04-10',
          total_genomic: 1,
          first_genomic_date: '2024-03-01',
          last_genomic_date: '2024-03-01',
          total_proteomics: 2,
          first_proteomics_date: '2024-03-15',
          last_proteomics_date: '2024-04-01'
        },
        visits: [
          { visit_no: 'V004', visit_time: '2024-02-01', department: '皮肤科', doctor_id: 104, chief_complaint: '皮疹' },
          { visit_no: 'V005', visit_time: '2024-03-15', department: '内科', doctor_id: 105, chief_complaint: '头痛' },
          { visit_no: 'V006', visit_time: '2024-04-10', department: '皮肤科', doctor_id: 104, chief_complaint: '瘙痒' }
        ],
        diagnoses: [
          { icd_code: 'L23', diagnosis_type: '过敏性皮炎', created_at: '2024-02-01' },
          { icd_code: 'G44', diagnosis_type: '偏头痛', created_at: '2024-03-15' },
          { icd_code: 'L50', diagnosis_type: '荨麻疹', created_at: '2024-04-10' }
        ],
        examinations: [
          { exam_type: '皮肤镜检查', exam_time: '2024-02-01', exam_result: '皮肤炎症', result_unit: '', reference_range: '', created_at: '2024-02-01' },
          { exam_type: '血常规', exam_time: '2024-03-15', exam_result: '正常', result_unit: '', reference_range: '4-10', created_at: '2024-03-15' },
          { exam_type: '过敏原检测', exam_time: '2024-04-10', exam_result: '尘螨阳性', result_unit: '', reference_range: '', created_at: '2024-04-10' }
        ],
        followup_plans: [
          {
            plan_name: '皮肤科随访',
            plan_type: '门诊',
            start_date: '2024-04-15',
            end_date: '2024-07-15',
            frequency: '每两周',
            status: '进行中'
          }
        ],
        followup_records: [
          {
            followup_date: '2024-04-15',
            followup_method: '门诊',
            symptoms: '皮疹有所改善',
            treatment_effect: '效果良好',
            followup_doctor: '张医生'
          }
        ],
        medications: [
          {
            medication_name: '地塞米松软膏',
            dosage: '外用',
            frequency: '每日两次',
            compliance: '良好',
            side_effects: '无'
          }
        ],
        environment_exposures: [
          {
            residence_type: '别墅',
            building_material: '砖混结构',
            pet_exposure: false,
            smoking_exposure: true
          }
        ],
        lifestyle_surveys: [
          {
            diet_pattern: '清淡饮食',
            exercise_frequency: '每周1-2次',
            sleep_quality: '一般',
            stress_level: '高'
          }
        ],
        samples: [
          {
            sample_no: 'S005',
            sample_type: '血液',
            collection_time: '2024-02-15',
            collection_method: '静脉采血',
            processor: '护士C'
          },
          {
            sample_no: 'S006',
            sample_type: '皮肤组织',
            collection_time: '2024-04-10',
            collection_method: '活检',
            processor: '技师D'
          }
        ],
        genomic_data: [
          {
            sequence_platform: 'Illumina HiSeq 2500',
            sequence_type: '外显子测序',
            gene_type: '皮肤相关基因',
            result_file_path: '/data/genomics/S005_WES.vcf'
          }
        ],
        proteomics_data: [
          {
            analysis_platform: '质谱仪',
            protein_marker: 'IgE',
            concentration: 85.2
          },
          {
            analysis_platform: '质谱仪',
            protein_marker: 'IL-17',
            concentration: 15.8
          }
        ]
      },
      {
        patient: {
          patient_no: 'P003',
          name: '王五',
          gender: '1',
          birth_date: '1988-12-10',
          phone: '13800138003',
          id_card: '110101198812100033',
          address: '北京市西城区西单街789号',
          height: 180.0,
          weight: 75.5,
          birth_weight: 3.5,
          medical_insurance_type: 1,
          medical_insurance_no: 'YB003456789',
          emergency_contact: '王老五',
          emergency_phone: '13900139003'
        },
        statistics: {
          total_visits: 15,
          total_diagnoses: 7,
          total_followups: 5,
          total_samples: 6,
          data_quality_score: 98,
          active_medications: 3,
          first_visit_date: '2023-12-01',
          last_visit_date: '2024-04-20',
          first_diagnosis_date: '2023-12-01',
          last_diagnosis_date: '2024-04-15',
          total_examinations: 12,
          first_exam_date: '2023-12-01',
          last_exam_date: '2024-04-18',
          last_followup_date: '2024-04-20',
          next_followup_date: '2024-05-05',
          first_sample_date: '2023-12-15',
          last_sample_date: '2024-04-20',
          total_genomic: 3,
          first_genomic_date: '2024-01-01',
          last_genomic_date: '2024-03-15',
          total_proteomics: 4,
          first_proteomics_date: '2024-01-15',
          last_proteomics_date: '2024-04-01'
        },
        visits: [
          { visit_no: 'V007', visit_time: '2023-12-01', department: '呼吸科', doctor_id: 106, chief_complaint: '呼吸困难' },
          { visit_no: 'V008', visit_time: '2024-01-15', department: '呼吸科', doctor_id: 106, chief_complaint: '哮喘发作' },
          { visit_no: 'V009', visit_time: '2024-02-20', department: '内科', doctor_id: 107, chief_complaint: '胸闷' },
          { visit_no: 'V010', visit_time: '2024-03-10', department: '呼吸科', doctor_id: 106, chief_complaint: '咳嗽' },
          { visit_no: 'V011', visit_time: '2024-04-20', department: '呼吸科', doctor_id: 106, chief_complaint: '定期复查' }
        ],
        diagnoses: [
          { icd_code: 'J45', diagnosis_type: '支气管哮喘', created_at: '2023-12-01' },
          { icd_code: 'J44', diagnosis_type: '慢性阻塞性肺疾病', created_at: '2024-01-15' },
          { icd_code: 'J20', diagnosis_type: '急性支气管炎', created_at: '2024-02-20' },
          { icd_code: 'J47', diagnosis_type: '支气管扩张', created_at: '2024-03-10' }
        ],
        examinations: [
          { exam_type: '肺功能检查', exam_time: '2023-12-01', exam_result: 'FEV1降低', result_unit: 'L', reference_range: '3.0-4.5', created_at: '2023-12-01' },
          { exam_type: '胸部CT', exam_time: '2024-01-15', exam_result: '肺气肿', result_unit: '', reference_range: '', created_at: '2024-01-15' },
          { exam_type: '血常规', exam_time: '2024-02-20', exam_result: '白细胞升高', result_unit: '10^9/L', reference_range: '3.5-9.5', created_at: '2024-02-20' },
          { exam_type: '痰培养', exam_time: '2024-03-10', exam_result: '阴性', result_unit: '', reference_range: '', created_at: '2024-03-10' }
        ],
        followup_plans: [
          {
            plan_name: '哮喘管理',
            plan_type: '门诊',
            start_date: '2023-12-15',
            end_date: '2024-12-15',
            frequency: '每月',
            status: '进行中'
          },
          {
            plan_name: '肺功能监测',
            plan_type: '检查',
            start_date: '2024-01-01',
            end_date: '2024-06-01',
            frequency: '每三个月',
            status: '进行中'
          }
        ],
        followup_records: [
          {
            followup_date: '2024-01-15',
            followup_method: '门诊',
            symptoms: '哮喘控制良好',
            treatment_effect: '效果良好',
            followup_doctor: '刘医生'
          },
          {
            followup_date: '2024-02-20',
            followup_method: '门诊',
            symptoms: '偶有胸闷',
            treatment_effect: '基本控制',
            followup_doctor: '刘医生'
          },
          {
            followup_date: '2024-03-10',
            followup_method: '门诊',
            symptoms: '症状稳定',
            treatment_effect: '效果良好',
            followup_doctor: '刘医生'
          },
          {
            followup_date: '2024-04-20',
            followup_method: '门诊',
            symptoms: '控制良好',
            treatment_effect: '效果良好',
            followup_doctor: '刘医生'
          }
        ],
        medications: [
          {
            medication_name: '沙丁胺醇',
            dosage: '100μg',
            frequency: '按需使用',
            compliance: '良好',
            side_effects: '心悸'
          },
          {
            medication_name: '布地奈德',
            dosage: '200μg',
            frequency: '每日两次',
            compliance: '良好',
            side_effects: '无'
          },
          {
            medication_name: '孟鲁司特',
            dosage: '10mg',
            frequency: '每日一次',
            compliance: '良好',
            side_effects: '无'
          }
        ],
        environment_exposures: [
          {
            residence_type: '平房',
            building_material: '砖木结构',
            pet_exposure: true,
            smoking_exposure: false
          }
        ],
        lifestyle_surveys: [
          {
            diet_pattern: '清淡饮食',
            exercise_frequency: '每周1-2次',
            sleep_quality: '一般',
            stress_level: '中等'
          }
        ],
        samples: [
          {
            sample_no: 'S007',
            sample_type: '血液',
            collection_time: '2023-12-15',
            collection_method: '静脉采血',
            processor: '护士E'
          },
          {
            sample_no: 'S008',
            sample_type: '痰液',
            collection_time: '2024-01-15',
            collection_method: '自取',
            processor: '护士E'
          },
          {
            sample_no: 'S009',
            sample_type: '血清',
            collection_time: '2024-02-20',
            collection_method: '离心分离',
            processor: '技师F'
          },
          {
            sample_no: 'S010',
            sample_type: '血液',
            collection_time: '2024-03-15',
            collection_method: '静脉采血',
            processor: '护士E'
          },
          {
            sample_no: 'S011',
            sample_type: '尿液',
            collection_time: '2024-04-01',
            collection_method: '自取',
            processor: '护士E'
          },
          {
            sample_no: 'S012',
            sample_type: '唾液',
            collection_time: '2024-04-20',
            collection_method: '自取',
            processor: '护士E'
          }
        ],
        genomic_data: [
          {
            sequence_platform: 'Illumina NovaSeq 6000',
            sequence_type: '全基因组测序',
            gene_type: '呼吸系统基因',
            result_file_path: '/data/genomics/S007_WGS.vcf'
          },
          {
            sequence_platform: 'BGI MGISEQ-2000',
            sequence_type: '外显子测序',
            gene_type: '免疫相关基因',
            result_file_path: '/data/genomics/S008_WES.vcf'
          },
          {
            sequence_platform: 'Thermo Fisher Ion Proton',
            sequence_type: '靶向测序',
            gene_type: '药物代谢基因',
            result_file_path: '/data/genomics/S009_Target.vcf'
          }
        ],
        proteomics_data: [
          {
            analysis_platform: '质谱仪',
            protein_marker: 'IgE',
            concentration: 180.3
          },
          {
            analysis_platform: '质谱仪',
            protein_marker: 'IL-4',
            concentration: 35.7
          },
          {
            analysis_platform: '质谱仪',
            protein_marker: 'IL-13',
            concentration: 28.9
          },
          {
            analysis_platform: '质谱仪',
            protein_marker: 'IL-5',
            concentration: 12.4
          }
        ]
      }
    ]

    const loadPatientData = async () => {
      if (!patientNo.value) {
        ElMessage.warning('请输入患者编号')
        return
      }
      loading.value = true
      try {
        // 随机选择一组mock数据
        const randomIndex = Math.floor(Math.random() * mockDataList.length)
        patientData.value = mockDataList[randomIndex]
        ElMessage.success('数据加载成功')
      } catch (error) {
        ElMessage.error('数据加载失败: ' + error.message)
        patientData.value = null
      } finally {
        loading.value = false
      }
    }

    const getGenderText = (gender) => {
      if (gender === 1 || gender === '1' || gender === 'M') return '男'
      if (gender === 0 || gender === '0' || gender === 'F') return '女'
      return '未知'
    }

    const calculateAge = (birthDate) => {
      if (!birthDate) return '未知'
      const birth = new Date(birthDate)
      const today = new Date()
      let age = today.getFullYear() - birth.getFullYear()
      const m = today.getMonth() - birth.getMonth()
      if (m < 0 || (m === 0 && today.getDate() < birth.getDate())) {
        age--
      }
      return age
    }

    const formatDate = (dateString) => {
      if (!dateString) return ''
      return new Date(dateString).toLocaleDateString('zh-CN')
    }

    const getQualityTagType = (score) => {
      if (score >= 90) return 'success'
      if (score >= 70) return 'warning'
      return 'danger'
    }

    return {
      patientNo,
      patientData,
      loading,
      activeTab,
      loadPatientData,
      getGenderText,
      calculateAge,
      formatDate,
      getQualityTagType
    }
  }
}
</script>

<style scoped>
.patient-integrated-view {
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

.content-area {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-card, .stats-card, .detail-card {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.stats-card .el-row {
  margin: 0;
}

.stat-item {
  text-align: center;
  padding: 20px;
  border-right: 1px solid #ebeef5;
}

.stat-item:last-child {
  border-right: none;
}

.stat-number {
  font-size: 32px;
  font-weight: bold;
  color: #409eff;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.detail-card {
  min-height: 500px;
}

.tab-content {
  padding: 20px;
}

.tab-content h3 {
  margin: 0 0 15px 0;
  color: #303133;
  font-size: 16px;
  font-weight: 600;
}

.el-table {
  margin-bottom: 20px;
}
</style> 