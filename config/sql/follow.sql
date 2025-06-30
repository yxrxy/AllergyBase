-- 随访计划表
CREATE TABLE IF NOT EXISTS `followup_plan` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `patient_id` BIGINT NOT NULL COMMENT '患者ID',
    `plan_name` VARCHAR(200) COMMENT '计划名称',
    `plan_type` VARCHAR(50) COMMENT '计划类型',
    `start_date` DATE COMMENT '计划开始日期',
    `end_date` DATE COMMENT '计划结束日期',
    `frequency` VARCHAR(100) COMMENT '随访频率',
    `status` VARCHAR(50) DEFAULT 'active' COMMENT '状态',
    `created_by` VARCHAR(50) COMMENT '创建人',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_patient_id` (`patient_id`),
    INDEX `idx_status` (`status`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='随访计划表';

-- 随访记录表
CREATE TABLE IF NOT EXISTS `followup_record` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `plan_id` BIGINT NOT NULL COMMENT '计划ID',
    `patient_id` BIGINT NOT NULL COMMENT '患者ID',
    `followup_date` TIMESTAMP NOT NULL COMMENT '随访日期',
    `followup_method` VARCHAR(50) COMMENT '随访方式',
    `symptoms` TEXT COMMENT '症状描述',
    `treatment_effect` TEXT COMMENT '治疗效果',
    `next_plan` TEXT COMMENT '下次计划',
    `followup_doctor` VARCHAR(50) COMMENT '随访医生',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_plan_id` (`plan_id`),
    INDEX `idx_patient_id` (`patient_id`),
    INDEX `idx_followup_date` (`followup_date`),
    FOREIGN KEY (`plan_id`) REFERENCES `followup_plan`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='随访记录表';

-- 用药监测表
CREATE TABLE IF NOT EXISTS `medication_monitor` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `followup_id` BIGINT NOT NULL COMMENT '随访记录ID',
    `medication_name` VARCHAR(200) COMMENT '药品名称',
    `dosage` VARCHAR(100) COMMENT '剂量',
    `frequency` VARCHAR(100) COMMENT '使用频率',
    `start_date` DATE COMMENT '开始日期',
    `end_date` DATE COMMENT '结束日期',
    `side_effects` TEXT COMMENT '不良反应',
    `compliance` VARCHAR(50) COMMENT '依从性',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_followup_id` (`followup_id`),
    FOREIGN KEY (`followup_id`) REFERENCES `followup_record`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用药监测表';