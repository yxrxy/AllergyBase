CREATE TABLE IF NOT EXISTS `biobank_sample_info` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `patient_id` BIGINT NOT NULL,
    `sample_no` VARCHAR(50) UNIQUE NOT NULL COMMENT '样本编号',
    `sample_type` VARCHAR(50) COMMENT '样本类型',
    `collection_time` TIMESTAMP NULL COMMENT '采集时间',
    `collection_method` VARCHAR(100) COMMENT '采集方法',
    `processor` VARCHAR(50) COMMENT '处理人',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (`patient_id`) REFERENCES `clinical_patient_info`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='样本信息表';

CREATE TABLE IF NOT EXISTS `biobank_storage_info` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `sample_id` BIGINT NOT NULL,
    `storage_location` VARCHAR(100) COMMENT '存储位置',
    `storage_temperature` DECIMAL(5,2) COMMENT '存储温度',
    `storage_time` TIMESTAMP NULL COMMENT '入库时间',
    `status` VARCHAR(50) COMMENT '样本状态',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (`sample_id`) REFERENCES `biobank_sample_info`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储信息表';

CREATE TABLE IF NOT EXISTS `biobank_genomic_data` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `sample_id` BIGINT NOT NULL,
    `sequence_platform` VARCHAR(100) COMMENT '测序平台',
    `sequence_type` VARCHAR(50) COMMENT '测序类型',
    `gene_type` VARCHAR(50) COMMENT '基因类型',
    `result_file_path` TEXT COMMENT '结果文件路径',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (`sample_id`) REFERENCES `biobank_sample_info`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='基因组数据表';

CREATE TABLE IF NOT EXISTS `biobank_proteomics_data` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `sample_id` BIGINT NOT NULL,
    `analysis_platform` VARCHAR(100) COMMENT '分析平台',
    `protein_marker` VARCHAR(100) COMMENT '蛋白标志物',
    `concentration` DECIMAL(10,4) COMMENT '浓度',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (`sample_id`) REFERENCES `biobank_sample_info`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='蛋白组学数据表';