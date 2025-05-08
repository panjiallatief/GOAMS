package repository

import (
	"github.com/panjiallatief/GOAMS/models"
	"gorm.io/gorm"
)

// AssetRepository defines CRUD operations for Asset
type AssetRepository interface {
	Create(asset *models.Asset) error
	GetAll() ([]models.Asset, error)
	GetByID(id uint) (*models.Asset, error)
	Update(asset *models.Asset) error
	Delete(asset *models.Asset) error
}

type assetRepository struct {
	db *gorm.DB
}

// NewAssetRepository creates a new instance of AssetRepository
func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) Create(asset *models.Asset) error {
	return r.db.Create(asset).Error
}

func (r *assetRepository) GetAll() ([]models.Asset, error) {
	var assets []models.Asset
	if err := r.db.Find(&assets).Error; err != nil {
		return nil, err
	}
	return assets, nil
}

func (r *assetRepository) GetByID(id uint) (*models.Asset, error) {
	var asset models.Asset
	if err := r.db.First(&asset, "id_asset = ?", id).Error; err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *assetRepository) Update(asset *models.Asset) error {
	return r.db.Save(asset).Error
}

func (r *assetRepository) Delete(asset *models.Asset) error {
	return r.db.Delete(asset).Error
}
