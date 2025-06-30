package service

import (
	"context"
	"fmt"
	"time"

	"github.com/yxrxy/AllergyBase/app/biobank/domain/model"
	"github.com/yxrxy/AllergyBase/app/biobank/domain/repository"
)

type BiobankService struct {
	db    repository.BiobankDB
	cache repository.BiobankCache
}

func NewBiobankService(db repository.BiobankDB, cache repository.BiobankCache) *BiobankService {
	return &BiobankService{
		db:    db,
		cache: cache,
	}
}

// 样本管理相关方法
func (s *BiobankService) CreateSample(ctx context.Context, sample *model.Sample) (*model.Sample, error) {
	// 生成样本编号
	if sample.SampleNo == "" {
		sample.SampleNo = s.generateSampleNo()
	}

	// 设置创建时间
	sample.CreatedAt = time.Now()
	sample.UpdatedAt = time.Now()

	createdSample, err := s.db.CreateSample(ctx, sample)
	if err != nil {
		return nil, err
	}

	// 缓存样本信息
	_ = s.cache.SetSampleInfo(ctx, createdSample.ID, createdSample)

	return createdSample, nil
}

func (s *BiobankService) GetSample(ctx context.Context, sampleID int64) (*model.Sample, error) {
	// 先从缓存获取
	sample, err := s.cache.GetSampleInfo(ctx, sampleID)
	if err == nil && sample != nil {
		return sample, nil
	}

	// 从数据库获取
	sample, err = s.db.GetSample(ctx, sampleID)
	if err != nil {
		return nil, err
	}

	// 缓存样本信息
	_ = s.cache.SetSampleInfo(ctx, sampleID, sample)

	return sample, nil
}

func (s *BiobankService) ListSamples(ctx context.Context, offset, limit int) ([]*model.Sample, int64, error) {
	return s.db.ListSamples(ctx, offset, limit)
}

// 存储信息管理相关方法
func (s *BiobankService) CreateStorageInfo(ctx context.Context, storageInfo *model.StorageInfo) (*model.StorageInfo, error) {
	// 设置创建时间
	storageInfo.CreatedAt = time.Now()
	storageInfo.UpdatedAt = time.Now()

	createdStorageInfo, err := s.db.CreateStorageInfo(ctx, storageInfo)
	if err != nil {
		return nil, err
	}

	// 缓存存储信息
	_ = s.cache.SetStorageInfo(ctx, createdStorageInfo.ID, createdStorageInfo)

	return createdStorageInfo, nil
}

func (s *BiobankService) GetStorageInfo(ctx context.Context, storageID int64) (*model.StorageInfo, error) {
	// 先从缓存获取
	storageInfo, err := s.cache.GetStorageInfo(ctx, storageID)
	if err == nil && storageInfo != nil {
		return storageInfo, nil
	}

	// 从数据库获取
	storageInfo, err = s.db.GetStorageInfo(ctx, storageID)
	if err != nil {
		return nil, err
	}

	// 缓存存储信息
	_ = s.cache.SetStorageInfo(ctx, storageID, storageInfo)

	return storageInfo, nil
}

// 基因组数据管理相关方法
func (s *BiobankService) CreateGenomicData(ctx context.Context, genomicData *model.GenomicData) (*model.GenomicData, error) {
	// 设置创建时间
	genomicData.CreatedAt = time.Now()
	genomicData.UpdatedAt = time.Now()

	createdGenomicData, err := s.db.CreateGenomicData(ctx, genomicData)
	if err != nil {
		return nil, err
	}

	// 缓存基因组数据
	_ = s.cache.SetGenomicData(ctx, createdGenomicData.ID, createdGenomicData)

	return createdGenomicData, nil
}

func (s *BiobankService) GetGenomicData(ctx context.Context, dataID int64) (*model.GenomicData, error) {
	// 先从缓存获取
	genomicData, err := s.cache.GetGenomicData(ctx, dataID)
	if err == nil && genomicData != nil {
		return genomicData, nil
	}

	// 从数据库获取
	genomicData, err = s.db.GetGenomicData(ctx, dataID)
	if err != nil {
		return nil, err
	}

	// 缓存基因组数据
	_ = s.cache.SetGenomicData(ctx, dataID, genomicData)

	return genomicData, nil
}

// 蛋白质组数据管理相关方法
func (s *BiobankService) CreateProteomicsData(ctx context.Context, proteomicsData *model.ProteomicsData) (*model.ProteomicsData, error) {
	// 设置创建时间
	proteomicsData.CreatedAt = time.Now()
	proteomicsData.UpdatedAt = time.Now()

	createdProteomicsData, err := s.db.CreateProteomicsData(ctx, proteomicsData)
	if err != nil {
		return nil, err
	}

	// 缓存蛋白质组数据
	_ = s.cache.SetProteomicsData(ctx, createdProteomicsData.ID, createdProteomicsData)

	return createdProteomicsData, nil
}

func (s *BiobankService) GetProteomicsData(ctx context.Context, dataID int64) (*model.ProteomicsData, error) {
	// 先从缓存获取
	proteomicsData, err := s.cache.GetProteomicsData(ctx, dataID)
	if err == nil && proteomicsData != nil {
		return proteomicsData, nil
	}

	// 从数据库获取
	proteomicsData, err = s.db.GetProteomicsData(ctx, dataID)
	if err != nil {
		return nil, err
	}

	// 缓存蛋白质组数据
	_ = s.cache.SetProteomicsData(ctx, dataID, proteomicsData)

	return proteomicsData, nil
}

// 生成样本编号
func (s *BiobankService) generateSampleNo() string {
	return fmt.Sprintf("S%s%06d", time.Now().Format("20060102"), time.Now().Unix()%1000000)
}
