package service

import (
	"github.com/panjiallatief/GOAMS/models"
	"github.com/panjiallatief/GOAMS/repository"
)

// AssetService defines business logic operations for Asset
type AssetService interface {
	CreateAsset(asset *models.Asset) error
	GetAllAssets() ([]models.Asset, error)
	GetAssetByID(id uint) (*models.Asset, error)
	UpdateAsset(asset *models.Asset) error
	DeleteAsset(id uint) error
}

type assetService struct {
	repo repository.AssetRepository
}

// NewAssetService creates a new instance of AssetService
func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{repo: repo}
}

func (s *assetService) CreateAsset(asset *models.Asset) error {
	return s.repo.Create(asset)
}

func (s *assetService) GetAllAssets() ([]models.Asset, error) {
	return s.repo.GetAll()
}

func (s *assetService) GetAssetByID(id uint) (*models.Asset, error) {
	return s.repo.GetByID(id)
}

func (s *assetService) UpdateAsset(asset *models.Asset) error {
	return s.repo.Update(asset)
}

func (s *assetService) DeleteAsset(id uint) error {
	asset, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(asset)
}
