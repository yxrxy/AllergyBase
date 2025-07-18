package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yxrxy/AllergyBase/pkg/constants"
)

type Storage interface {
	Save(data []byte, filename string) (string, error)
}

type LocalStorage struct {
	basePath string
	baseURL  string
}

func NewLocalStorage(basePath, baseURL string) *LocalStorage {
	return &LocalStorage{
		basePath: basePath,
		baseURL:  baseURL,
	}
}

func (s *LocalStorage) Save(data []byte, filename string) (string, error) {
	fullPath := filepath.Join(s.basePath, filename)
	if err := os.WriteFile(fullPath, data, constants.FilePermission); err != nil {
		return "", fmt.Errorf("save file error: %w", err)
	}
	return fmt.Sprintf("%s/%s", s.baseURL, filename), nil
}
