-- 使用数据库
USE allergybase;

CREATE TABLE IF NOT EXISTS `clinical_patient_info` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '患者ID',
    `patient_no` varchar(50) NOT NULL COMMENT '患者编号',
    `name` varchar(50) DEFAULT NULL COMMENT '姓名',
    `gender` tinyint(4) DEFAULT NULL COMMENT '性别:0-女,1-男',
    `birth_date` date DEFAULT NULL COMMENT '出生日期',
    `height` decimal(5,2) DEFAULT NULL COMMENT '身高(cm)',
    `weight` decimal(5,2) DEFAULT NULL COMMENT '体重(kg)',
    `birth_weight` decimal(5,2) DEFAULT NULL COMMENT '出生体重(kg)',
    `medical_insurance_type` tinyint(4) DEFAULT NULL COMMENT '医保类型',
    `medical_insurance_no` varchar(50) DEFAULT NULL COMMENT '医保号',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_patient_no` (`patient_no`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='患者信息表';

CREATE TABLE IF NOT EXISTS `clinical_visit_record` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '就诊记录ID',
    `patient_id` bigint(20) NOT NULL COMMENT '患者ID',
    `visit_no` varchar(50) NOT NULL COMMENT '就诊号',
    `visit_time` timestamp NOT NULL COMMENT '就诊时间',
    `department` varchar(50) DEFAULT NULL COMMENT '就诊科室',
    `visit_type` tinyint(4) DEFAULT NULL COMMENT '就诊类型:1-门诊,2-急诊,3-住院',
    `doctor_id` bigint(20) DEFAULT NULL COMMENT '医生ID',
    `chief_complaint` text COMMENT '主诉',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_patient_id` (`patient_id`),
    KEY `idx_visit_no` (`visit_no`),
    KEY `idx_deleted_at` (`deleted_at`),
    CONSTRAINT `fk_visit_patient` FOREIGN KEY (`patient_id`) REFERENCES `clinical_patient_info` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='就诊记录表';

CREATE TABLE IF NOT EXISTS `clinical_diagnosis` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '诊断ID',
    `visit_id` bigint(20) NOT NULL COMMENT '就诊记录ID',
    `diagnosis_type` varchar(50) DEFAULT NULL COMMENT '诊断类型',
    `icd_code` varchar(50) DEFAULT NULL COMMENT 'ICD-11编码',
    `severity` varchar(20) DEFAULT NULL COMMENT '严重程度',
    `diagnosis_time` timestamp NULL DEFAULT NULL COMMENT '诊断时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_visit_id` (`visit_id`),
    KEY `idx_deleted_at` (`deleted_at`),
    CONSTRAINT `fk_diagnosis_visit` FOREIGN KEY (`visit_id`) REFERENCES `clinical_visit_record` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='诊断表';

CREATE TABLE IF NOT EXISTS `clinical_examination` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '检查ID',
    `visit_id` bigint(20) NOT NULL COMMENT '就诊记录ID',
    `exam_type` varchar(50) DEFAULT NULL COMMENT '检查类型',
    `exam_time` timestamp NULL DEFAULT NULL COMMENT '检查时间',
    `exam_result` text COMMENT '检查结果',
    `result_unit` varchar(20) DEFAULT NULL COMMENT '结果单位',
    `reference_range` varchar(50) DEFAULT NULL COMMENT '参考范围',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_visit_id` (`visit_id`),
    KEY `idx_deleted_at` (`deleted_at`),
    CONSTRAINT `fk_examination_visit` FOREIGN KEY (`visit_id`) REFERENCES `clinical_visit_record` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='检查表';