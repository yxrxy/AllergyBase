package repository

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/biobank/domain/model"
)

type BiobankDB interface {
	// 样本管理
	CreateSample(ctx context.Context, sample *model.Sample) (*model.Sample, error)
	GetSample(ctx context.Context, sampleID int64) (*model.Sample, error)
	ListSamples(ctx context.Context, offset, limit int) ([]*model.Sample, int64, error)
	UpdateSample(ctx context.Context, sample *model.Sample) error
	DeleteSample(ctx context.Context, sampleID int64) error

	// 存储信息管理
	CreateStorageInfo(ctx context.Context, storageInfo *model.StorageInfo) (*model.StorageInfo, error)
	GetStorageInfo(ctx context.Context, storageID int64) (*model.StorageInfo, error)
	ListStorageInfo(ctx context.Context, offset, limit int) ([]*model.StorageInfo, int64, error)
	UpdateStorageInfo(ctx context.Context, storageInfo *model.StorageInfo) error
	DeleteStorageInfo(ctx context.Context, storageID int64) error

	// 基因组数据管理
	CreateGenomicData(ctx context.Context, genomicData *model.GenomicData) (*model.GenomicData, error)
	GetGenomicData(ctx context.Context, dataID int64) (*model.GenomicData, error)
	ListGenomicData(ctx context.Context, offset, limit int) ([]*model.GenomicData, int64, error)
	UpdateGenomicData(ctx context.Context, genomicData *model.GenomicData) error
	DeleteGenomicData(ctx context.Context, dataID int64) error

	// 蛋白质组数据管理
	CreateProteomicsData(ctx context.Context, proteomicsData *model.ProteomicsData) (*model.ProteomicsData, error)
	GetProteomicsData(ctx context.Context, dataID int64) (*model.ProteomicsData, error)
	ListProteomicsData(ctx context.Context, offset, limit int) ([]*model.ProteomicsData, int64, error)
	UpdateProteomicsData(ctx context.Context, proteomicsData *model.ProteomicsData) error
	DeleteProteomicsData(ctx context.Context, dataID int64) error
}

type BiobankCache interface {
	// 样本缓存
	SetSampleInfo(ctx context.Context, sampleID int64, sample *model.Sample) error
	GetSampleInfo(ctx context.Context, sampleID int64) (*model.Sample, error)
	DeleteSampleInfo(ctx context.Context, sampleID int64) error

	// 存储信息缓存
	SetStorageInfo(ctx context.Context, storageID int64, storageInfo *model.StorageInfo) error
	GetStorageInfo(ctx context.Context, storageID int64) (*model.StorageInfo, error)
	DeleteStorageInfo(ctx context.Context, storageID int64) error

	// 基因组数据缓存
	SetGenomicData(ctx context.Context, dataID int64, genomicData *model.GenomicData) error
	GetGenomicData(ctx context.Context, dataID int64) (*model.GenomicData, error)
	DeleteGenomicData(ctx context.Context, dataID int64) error

	// 蛋白质组数据缓存
	SetProteomicsData(ctx context.Context, dataID int64, proteomicsData *model.ProteomicsData) error
	GetProteomicsData(ctx context.Context, dataID int64) (*model.ProteomicsData, error)
	DeleteProteomicsData(ctx context.Context, dataID int64) error
}
