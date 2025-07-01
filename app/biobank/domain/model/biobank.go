package model

import "time"

// Sample 样本信息
type Sample struct {
	ID               int64  `gorm:"primaryKey;autoIncrement"`
	PatientID        int64  `gorm:"index"`
	SampleNo         string `gorm:"uniqueIndex;size:50"`
	SampleType       string `gorm:"size:50"`
	CollectionTime   *time.Time
	CollectionMethod string `gorm:"size:100"`
	Processor        string `gorm:"size:50"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// TableName 指定表名
func (Sample) TableName() string {
	return "biobank_sample_info"
}

// StorageInfo 存储信息
type StorageInfo struct {
	ID                 int64   `gorm:"primaryKey;autoIncrement"`
	SampleID           int64   `gorm:"index"`
	StorageLocation    string  `gorm:"size:100"`
	StorageTemperature float64 // 存储温度
	StorageTime        *time.Time
	Status             string `gorm:"size:50"`
	CreatedAt          time.Time
	UpdatedAt          time.Time

	// 关联
	Sample *Sample `gorm:"foreignKey:SampleID"`
}

// TableName 指定表名
func (StorageInfo) TableName() string {
	return "biobank_storage_info"
}

// GenomicData 基因组数据
type GenomicData struct {
	ID               int64  `gorm:"primaryKey;autoIncrement"`
	SampleID         int64  `gorm:"index"`
	SequencePlatform string `gorm:"size:100"`
	SequenceType     string `gorm:"size:50"`
	GeneType         string `gorm:"size:50"`
	ResultFilePath   string `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	// 关联
	Sample *Sample `gorm:"foreignKey:SampleID"`
}

// TableName 指定表名
func (GenomicData) TableName() string {
	return "biobank_genomic_data"
}

// ProteomicsData 蛋白组学数据
type ProteomicsData struct {
	ID               int64   `gorm:"primaryKey;autoIncrement"`
	SampleID         int64   `gorm:"index"`
	AnalysisPlatform string  `gorm:"size:100"`
	ProteinMarker    string  `gorm:"size:100"`
	Concentration    float64 // 浓度
	CreatedAt        time.Time
	UpdatedAt        time.Time

	// 关联
	Sample *Sample `gorm:"foreignKey:SampleID"`
}

// TableName 指定表名
func (ProteomicsData) TableName() string {
	return "biobank_proteomics_data"
}
