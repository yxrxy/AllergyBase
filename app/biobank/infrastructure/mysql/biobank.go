package mysql

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/biobank/domain/model"
	"github.com/yxrxy/AllergyBase/app/biobank/domain/repository"
	"gorm.io/gorm"
)

type biobankDB struct {
	db *gorm.DB
}

func NewBiobankDB(db *gorm.DB) repository.BiobankDB {
	return &biobankDB{db: db}
}

// 样本管理
func (b *biobankDB) CreateSample(ctx context.Context, sample *model.Sample) (*model.Sample, error) {
	if err := b.db.WithContext(ctx).Create(sample).Error; err != nil {
		return nil, err
	}
	return sample, nil
}

func (b *biobankDB) GetSample(ctx context.Context, sampleID int64) (*model.Sample, error) {
	var sample model.Sample
	if err := b.db.WithContext(ctx).First(&sample, sampleID).Error; err != nil {
		return nil, err
	}
	return &sample, nil
}

func (b *biobankDB) ListSamples(ctx context.Context, offset, limit int) ([]*model.Sample, int64, error) {
	var samples []*model.Sample
	var total int64

	if err := b.db.WithContext(ctx).Model(&model.Sample{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := b.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&samples).Error; err != nil {
		return nil, 0, err
	}

	return samples, total, nil
}

func (b *biobankDB) UpdateSample(ctx context.Context, sample *model.Sample) error {
	return b.db.WithContext(ctx).Save(sample).Error
}

func (b *biobankDB) DeleteSample(ctx context.Context, sampleID int64) error {
	return b.db.WithContext(ctx).Delete(&model.Sample{}, sampleID).Error
}

// 存储信息管理
func (b *biobankDB) CreateStorageInfo(ctx context.Context, storageInfo *model.StorageInfo) (*model.StorageInfo, error) {
	if err := b.db.WithContext(ctx).Create(storageInfo).Error; err != nil {
		return nil, err
	}
	return storageInfo, nil
}

func (b *biobankDB) GetStorageInfo(ctx context.Context, storageID int64) (*model.StorageInfo, error) {
	var storageInfo model.StorageInfo
	if err := b.db.WithContext(ctx).First(&storageInfo, storageID).Error; err != nil {
		return nil, err
	}
	return &storageInfo, nil
}

func (b *biobankDB) ListStorageInfo(ctx context.Context, offset, limit int) ([]*model.StorageInfo, int64, error) {
	var storageInfos []*model.StorageInfo
	var total int64

	if err := b.db.WithContext(ctx).Model(&model.StorageInfo{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := b.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&storageInfos).Error; err != nil {
		return nil, 0, err
	}

	return storageInfos, total, nil
}

func (b *biobankDB) UpdateStorageInfo(ctx context.Context, storageInfo *model.StorageInfo) error {
	return b.db.WithContext(ctx).Save(storageInfo).Error
}

func (b *biobankDB) DeleteStorageInfo(ctx context.Context, storageID int64) error {
	return b.db.WithContext(ctx).Delete(&model.StorageInfo{}, storageID).Error
}

// 基因组数据管理
func (b *biobankDB) CreateGenomicData(ctx context.Context, genomicData *model.GenomicData) (*model.GenomicData, error) {
	if err := b.db.WithContext(ctx).Create(genomicData).Error; err != nil {
		return nil, err
	}
	return genomicData, nil
}

func (b *biobankDB) GetGenomicData(ctx context.Context, dataID int64) (*model.GenomicData, error) {
	var genomicData model.GenomicData
	if err := b.db.WithContext(ctx).First(&genomicData, dataID).Error; err != nil {
		return nil, err
	}
	return &genomicData, nil
}

func (b *biobankDB) ListGenomicData(ctx context.Context, offset, limit int) ([]*model.GenomicData, int64, error) {
	var genomicDataList []*model.GenomicData
	var total int64

	if err := b.db.WithContext(ctx).Model(&model.GenomicData{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := b.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&genomicDataList).Error; err != nil {
		return nil, 0, err
	}

	return genomicDataList, total, nil
}

func (b *biobankDB) UpdateGenomicData(ctx context.Context, genomicData *model.GenomicData) error {
	return b.db.WithContext(ctx).Save(genomicData).Error
}

func (b *biobankDB) DeleteGenomicData(ctx context.Context, dataID int64) error {
	return b.db.WithContext(ctx).Delete(&model.GenomicData{}, dataID).Error
}

// 蛋白质组数据管理
func (b *biobankDB) CreateProteomicsData(ctx context.Context, proteomicsData *model.ProteomicsData) (*model.ProteomicsData, error) {
	if err := b.db.WithContext(ctx).Create(proteomicsData).Error; err != nil {
		return nil, err
	}
	return proteomicsData, nil
}

func (b *biobankDB) GetProteomicsData(ctx context.Context, dataID int64) (*model.ProteomicsData, error) {
	var proteomicsData model.ProteomicsData
	if err := b.db.WithContext(ctx).First(&proteomicsData, dataID).Error; err != nil {
		return nil, err
	}
	return &proteomicsData, nil
}

func (b *biobankDB) ListProteomicsData(ctx context.Context, offset, limit int) ([]*model.ProteomicsData, int64, error) {
	var proteomicsDataList []*model.ProteomicsData
	var total int64

	if err := b.db.WithContext(ctx).Model(&model.ProteomicsData{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := b.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&proteomicsDataList).Error; err != nil {
		return nil, 0, err
	}

	return proteomicsDataList, total, nil
}

func (b *biobankDB) UpdateProteomicsData(ctx context.Context, proteomicsData *model.ProteomicsData) error {
	return b.db.WithContext(ctx).Save(proteomicsData).Error
}

func (b *biobankDB) DeleteProteomicsData(ctx context.Context, dataID int64) error {
	return b.db.WithContext(ctx).Delete(&model.ProteomicsData{}, dataID).Error
}
