package usecase

import (
	"context"

	"github.com/yxrxy/AllergyBase/app/biobank/domain/model"
	"github.com/yxrxy/AllergyBase/app/biobank/domain/repository"
	"github.com/yxrxy/AllergyBase/app/biobank/domain/service"
)

type BiobankUseCase interface {
	CreateSample(ctx context.Context, sample *model.Sample) (*model.Sample, error)
	GetSample(ctx context.Context, sampleID int64) (*model.Sample, error)
	ListSamples(ctx context.Context, offset, limit int) ([]*model.Sample, int64, error)
	CreateStorageInfo(ctx context.Context, storageInfo *model.StorageInfo) (*model.StorageInfo, error)
	GetStorageInfo(ctx context.Context, storageID int64) (*model.StorageInfo, error)
	CreateGenomicData(ctx context.Context, genomicData *model.GenomicData) (*model.GenomicData, error)
	GetGenomicData(ctx context.Context, dataID int64) (*model.GenomicData, error)
	CreateProteomicsData(ctx context.Context, proteomicsData *model.ProteomicsData) (*model.ProteomicsData, error)
	GetProteomicsData(ctx context.Context, dataID int64) (*model.ProteomicsData, error)
}

type useCase struct {
	db      repository.BiobankDB
	service *service.BiobankService
}

func NewBiobankCase(db repository.BiobankDB, service *service.BiobankService) BiobankUseCase {
	return &useCase{
		db:      db,
		service: service,
	}
}

func (u *useCase) CreateSample(ctx context.Context, sample *model.Sample) (*model.Sample, error) {
	return u.service.CreateSample(ctx, sample)
}

func (u *useCase) GetSample(ctx context.Context, sampleID int64) (*model.Sample, error) {
	return u.service.GetSample(ctx, sampleID)
}

func (u *useCase) ListSamples(ctx context.Context, offset, limit int) ([]*model.Sample, int64, error) {
	return u.service.ListSamples(ctx, offset, limit)
}

func (u *useCase) CreateStorageInfo(ctx context.Context, storageInfo *model.StorageInfo) (*model.StorageInfo, error) {
	return u.service.CreateStorageInfo(ctx, storageInfo)
}

func (u *useCase) GetStorageInfo(ctx context.Context, storageID int64) (*model.StorageInfo, error) {
	return u.service.GetStorageInfo(ctx, storageID)
}

func (u *useCase) CreateGenomicData(ctx context.Context, genomicData *model.GenomicData) (*model.GenomicData, error) {
	return u.service.CreateGenomicData(ctx, genomicData)
}

func (u *useCase) GetGenomicData(ctx context.Context, dataID int64) (*model.GenomicData, error) {
	return u.service.GetGenomicData(ctx, dataID)
}

func (u *useCase) CreateProteomicsData(ctx context.Context, proteomicsData *model.ProteomicsData) (*model.ProteomicsData, error) {
	return u.service.CreateProteomicsData(ctx, proteomicsData)
}

func (u *useCase) GetProteomicsData(ctx context.Context, dataID int64) (*model.ProteomicsData, error) {
	return u.service.GetProteomicsData(ctx, dataID)
}
