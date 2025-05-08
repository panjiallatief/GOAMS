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
	// ExistsBySN returns true if an asset with the given serial number exists
	ExistsBySN(itemSN string) (bool, error)
	// MaxNoUrutForYear returns the current maximum no_urut for assets arriving in the given year
	MaxNoUrutForYear(year string) (int, error)
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

// ExistsBySN returns true if an asset with the given serial number exists
func (r *assetRepository) ExistsBySN(itemSN string) (bool, error) {
	var count int64
	if err := r.db.Model(&models.Asset{}).
		Where("item_serial_number = ?", itemSN).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// MaxNoUrutForYear returns the current maximum no_urut for assets arriving in the given year
func (r *assetRepository) MaxNoUrutForYear(year string) (int, error) {
	var max int
	if err := r.db.Model(&models.Asset{}).
		Select("COALESCE(MAX(no_urut), 0)").
		Where("tahun_barang_datang = ?", year).
		Scan(&max).Error; err != nil {
		return 0, err
	}
	return max, nil
}
