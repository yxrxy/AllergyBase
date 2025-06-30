namespace go model

// 基本响应结构
struct BaseResp {
    1: i64 code,    // 错误码，0表示成功
    2: string msg,  // 错误信息
}

// 用户模型
struct User {
    1: required i64 id,              // 用户ID
    2: required string username,     // 用户名
    4: optional string avatar,       // 头像URL
    6: optional i64 followCount,     // 关注数
    7: optional i64 followerCount,   // 粉丝数
    8: optional bool isFollow,       // 是否已关注
    9: optional i64 likeCount,       // 获赞数量
    10: optional i64 videoCount,     // 视频数量
}

// 性别枚举
enum Gender {
    FEMALE = 0
    MALE = 1
}

// 就诊类型
enum VisitType {
    OUTPATIENT = 1    // 门诊
    EMERGENCY = 2     // 急诊
    INPATIENT = 3     // 住院
}

// 患者信息
struct Patient {
    1: required i64 id
    2: required string patientNo      // 患者编号
    3: required string name           // 姓名
    4: required Gender gender         // 性别
    5: required string birthDate      // 出生日期
    6: optional string idCard         // 身份证号
    7: optional string phone          // 电话号码
    8: optional string address        // 地址
    9: optional string emergencyContact     // 紧急联系人
    10: optional string emergencyPhone      // 紧急联系人电话
    11: optional double height         // 身高(cm)
    12: optional double weight         // 体重(kg)
    13: optional double birthWeight    // 出生体重(kg)
    14: optional i16 medicalInsuranceType  // 医保类型
    15: optional string medicalInsuranceNo // 医保号
}

// 就诊记录
struct Visit {
    1: required i64 id
    2: required i64 patientId
    3: required string visitNo        // 就诊号
    4: required string visitTime      // 就诊时间
    5: optional string department     // 就诊科室
    6: required VisitType visitType   // 就诊类型
    7: optional i64 doctorId         // 医生ID
    8: optional string chiefComplaint // 主诉
}

// 诊断信息
struct Diagnosis {
    1: required i64 id
    2: required i64 visitId
    3: optional string diagnosisType  // 诊断类型
    4: optional string icdCode        // ICD-11编码
    5: optional string severity       // 严重程度
    6: optional string diagnosisTime  // 诊断时间
}

// 检查信息
struct Examination {
    1: required i64 id
    2: required i64 visitId
    3: optional string examType       // 检查类型
    4: optional string examTime       // 检查时间
    5: optional string examResult     // 检查结果
    6: optional string resultUnit     // 结果单位
    7: optional string referenceRange // 参考范围
}

// 环境暴露信息
struct EnvironmentExposure {
    1: required i64 id
    2: required i64 patientId
    3: optional string residenceType        // 居住类型
    4: optional string buildingMaterial     // 建筑材料
    5: optional string ventilationFrequency // 通风频率
    6: optional string airConditioningUsage // 空调使用情况
    7: optional bool petExposure           // 宠物接触史
    8: optional bool smokingExposure       // 吸烟暴露
}

// 环境监测数据
struct EnvironmentMonitor {
    1: required i64 id
    2: required string locationCode     // 监测点位置编码
    3: required string monitorTime      // 监测时间
    4: optional double pm25            // PM2.5浓度
    5: optional double temperature     // 温度
    6: optional double humidity        // 湿度
    7: optional double pollenConcentration // 花粉浓度
    8: optional string allergenLevel    // 过敏原水平
}

// 生活方式调查
struct LifestyleSurvey {
    1: required i64 id
    2: required i64 patientId
    3: optional string dietPattern      // 饮食模式
    4: optional string exerciseFrequency // 运动频率
    5: optional string sleepQuality     // 睡眠质量
    6: optional string stressLevel      // 压力水平
}

// 随访计划
struct FollowupPlan {
    1: required i64 id
    2: required i64 patientId
    3: required string planStartDate     // 计划开始日期
    4: optional string followupFrequency // 随访频率
    5: optional string followupMethod    // 随访方式
    6: optional i64 responsibleDoctorId  // 负责医生ID
}

// 随访记录
struct FollowupRecord {
    1: required i64 id
    2: required i64 planId
    3: required string followupDate      // 随访日期
    4: optional i32 symptomScore        // 症状评分
    5: optional string medicationCompliance // 用药依从性
    6: optional string sideEffects      // 不良反应
    7: optional i32 qualityOfLifeScore  // 生活质量评分
}

// 用药监测
struct MedicationMonitor {
    1: required i64 id
    2: required i64 followupId
    3: required string drugName         // 药品名称
    4: optional string usageFrequency   // 使用频率
    5: optional string complianceLevel  // 依从性级别
    6: optional string effectiveness    // 效果评估
}

// 生物样本信息
struct SampleInfo {
    1: required i64 id
    2: required i64 patientId
    3: required string sampleNo        // 样本编号
    4: optional string sampleType      // 样本类型
    5: optional string collectionTime  // 采集时间
    6: optional string collectionMethod // 采集方法
    7: optional string processor       // 处理人
}

// 样本存储信息
struct StorageInfo {
    1: required i64 id
    2: required i64 sampleId
    3: optional string storageLocation  // 存储位置
    4: optional double storageTemperature // 存储温度
    5: optional string storageTime      // 入库时间
    6: optional string status           // 样本状态
}

// 基因组数据
struct GenomicData {
    1: required i64 id
    2: required i64 sampleId
    3: optional string sequencePlatform // 测序平台
    4: optional string sequenceType     // 测序类型
    5: optional string geneType         // 基因类型
    6: optional string resultFilePath   // 结果文件路径
}

// 蛋白组学数据
struct ProteomicsData {
    1: required i64 id
    2: required i64 sampleId
    3: optional string analysisPlatform // 分析平台
    4: optional string proteinMarker    // 蛋白标志物
    5: optional double concentration    // 浓度
}