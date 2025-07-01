CREATE TABLE epi_environment_exposure (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    patient_id BIGINT NOT NULL,
    residence_type VARCHAR(50) COMMENT '居住类型',
    building_material VARCHAR(50) COMMENT '建筑材料',
    ventilation_frequency VARCHAR(50) COMMENT '通风频率',
    air_conditioning_usage TEXT COMMENT '空调使用情况',
    pet_exposure BOOLEAN COMMENT '宠物接触史',
    smoking_exposure BOOLEAN COMMENT '吸烟暴露',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (patient_id) REFERENCES clinical_patient_info(id)
);

CREATE TABLE epi_environment_monitor (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    location_code VARCHAR(50) COMMENT '监测点位置编码',
    monitor_time TIMESTAMP NOT NULL COMMENT '监测时间',
    pm25 DECIMAL(6,2) COMMENT 'PM2.5浓度',
    temperature DECIMAL(5,2) COMMENT '温度',
    humidity DECIMAL(5,2) COMMENT '湿度',
    pollen_concentration DECIMAL(6,2) COMMENT '花粉浓度',
    allergen_level TEXT COMMENT '过敏原水平',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);

CREATE TABLE epi_lifestyle_survey (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    patient_id BIGINT NOT NULL,
    diet_pattern VARCHAR(50) COMMENT '饮食模式',
    exercise_frequency VARCHAR(50) COMMENT '运动频率',
    sleep_quality VARCHAR(50) COMMENT '睡眠质量',
    stress_level VARCHAR(50) COMMENT '压力水平',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (patient_id) REFERENCES clinical_patient_info(id)
);