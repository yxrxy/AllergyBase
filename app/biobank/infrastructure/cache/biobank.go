package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yxrxy/AllergyBase/app/biobank/domain/model"
	"github.com/yxrxy/AllergyBase/app/biobank/domain/repository"
)

type biobankCache struct {
	rdb *redis.Client
}

func NewBiobankCache(rdb *redis.Client) repository.BiobankCache {
	return &biobankCache{rdb: rdb}
}

const (
	sampleKeyPrefix     = "biobank:sample:"
	storageKeyPrefix    = "biobank:storage:"
	genomicKeyPrefix    = "biobank:genomic:"
	proteomicsKeyPrefix = "biobank:proteomics:"
	cacheExpiration     = 30 * time.Minute
)

// 样本缓存
func (b *biobankCache) SetSampleInfo(ctx context.Context, sampleID int64, sample *model.Sample) error {
	key := fmt.Sprintf("%s%d", sampleKeyPrefix, sampleID)
	data, err := json.Marshal(sample)
	if err != nil {
		return err
	}
	return b.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (b *biobankCache) GetSampleInfo(ctx context.Context, sampleID int64) (*model.Sample, error) {
	key := fmt.Sprintf("%s%d", sampleKeyPrefix, sampleID)
	data, err := b.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var sample model.Sample
	if err := json.Unmarshal([]byte(data), &sample); err != nil {
		return nil, err
	}

	return &sample, nil
}

func (b *biobankCache) DeleteSampleInfo(ctx context.Context, sampleID int64) error {
	key := fmt.Sprintf("%s%d", sampleKeyPrefix, sampleID)
	return b.rdb.Del(ctx, key).Err()
}

// 存储信息缓存
func (b *biobankCache) SetStorageInfo(ctx context.Context, storageID int64, storageInfo *model.StorageInfo) error {
	key := fmt.Sprintf("%s%d", storageKeyPrefix, storageID)
	data, err := json.Marshal(storageInfo)
	if err != nil {
		return err
	}
	return b.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (b *biobankCache) GetStorageInfo(ctx context.Context, storageID int64) (*model.StorageInfo, error) {
	key := fmt.Sprintf("%s%d", storageKeyPrefix, storageID)
	data, err := b.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var storageInfo model.StorageInfo
	if err := json.Unmarshal([]byte(data), &storageInfo); err != nil {
		return nil, err
	}

	return &storageInfo, nil
}

func (b *biobankCache) DeleteStorageInfo(ctx context.Context, storageID int64) error {
	key := fmt.Sprintf("%s%d", storageKeyPrefix, storageID)
	return b.rdb.Del(ctx, key).Err()
}

// 基因组数据缓存
func (b *biobankCache) SetGenomicData(ctx context.Context, dataID int64, genomicData *model.GenomicData) error {
	key := fmt.Sprintf("%s%d", genomicKeyPrefix, dataID)
	data, err := json.Marshal(genomicData)
	if err != nil {
		return err
	}
	return b.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (b *biobankCache) GetGenomicData(ctx context.Context, dataID int64) (*model.GenomicData, error) {
	key := fmt.Sprintf("%s%d", genomicKeyPrefix, dataID)
	data, err := b.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var genomicData model.GenomicData
	if err := json.Unmarshal([]byte(data), &genomicData); err != nil {
		return nil, err
	}

	return &genomicData, nil
}

func (b *biobankCache) DeleteGenomicData(ctx context.Context, dataID int64) error {
	key := fmt.Sprintf("%s%d", genomicKeyPrefix, dataID)
	return b.rdb.Del(ctx, key).Err()
}

// 蛋白质组数据缓存
func (b *biobankCache) SetProteomicsData(ctx context.Context, dataID int64, proteomicsData *model.ProteomicsData) error {
	key := fmt.Sprintf("%s%d", proteomicsKeyPrefix, dataID)
	data, err := json.Marshal(proteomicsData)
	if err != nil {
		return err
	}
	return b.rdb.Set(ctx, key, data, cacheExpiration).Err()
}

func (b *biobankCache) GetProteomicsData(ctx context.Context, dataID int64) (*model.ProteomicsData, error) {
	key := fmt.Sprintf("%s%d", proteomicsKeyPrefix, dataID)
	data, err := b.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var proteomicsData model.ProteomicsData
	if err := json.Unmarshal([]byte(data), &proteomicsData); err != nil {
		return nil, err
	}

	return &proteomicsData, nil
}

func (b *biobankCache) DeleteProteomicsData(ctx context.Context, dataID int64) error {
	key := fmt.Sprintf("%s%d", proteomicsKeyPrefix, dataID)
	return b.rdb.Del(ctx, key).Err()
}
