 -- 添加患者表缺失字段的迁移脚本
USE allergybase;

-- 添加缺失的字段到 clinical_patient_info 表
ALTER TABLE `clinical_patient_info` 
ADD COLUMN `id_card` varchar(18) DEFAULT NULL COMMENT '身份证号' AFTER `birth_date`,
ADD COLUMN `phone` varchar(20) DEFAULT NULL COMMENT '电话号码' AFTER `id_card`,
ADD COLUMN `address` varchar(200) DEFAULT NULL COMMENT '地址' AFTER `phone`,
ADD COLUMN `emergency_contact` varchar(100) DEFAULT NULL COMMENT '紧急联系人' AFTER `address`,
ADD COLUMN `emergency_phone` varchar(20) DEFAULT NULL COMMENT '紧急联系人电话' AFTER `emergency_contact`;

-- 验证表结构
DESCRIBE `clinical_patient_info`;