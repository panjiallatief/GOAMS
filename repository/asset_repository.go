package repository

import (
	"github.com/panjiallatief/GOAMS/models"
	"gorm.io/gorm"
)

// AssetRepository defines CRUD operations for Asset
type AssetRepository interface {
	Create(asset *models.Asset) error
	GetAll() ([]*models.Asset, error)
	GetByID(id int64) (*models.Asset, error)
	Update(asset *models.Asset) error
	Delete(id int64) error
	// ExistsBySN returns true if an asset with the given serial number exists
	ExistsBySN(itemSN string) (bool, error)
	// MaxNoUrutForYear returns the current maximum no_urut for assets arriving in the given year
	MaxNoUrutForYear(year string) (int, error)
	// Add validation methods
	ValidatePIC(picName string) (bool, error)
	ValidateOwner(ownerName string) (bool, error)
	ValidateVendor(vendorName string) (bool, error)
	ValidateCategory(categoryName string) (bool, error)
	ValidateBrand(brandName string) (bool, error)
	ValidateModel(modelName string) (bool, error)
	ValidateDivisi(divisiName string) (bool, error)
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

func (r *assetRepository) GetAll() ([]*models.Asset, error) {
	var assets []*models.Asset
	// Select all fields explicitly to ensure they're loaded
	if err := r.db.Order("id_asset DESC").Find(&assets).Error; err != nil {
		return nil, err
	}
	return assets, nil
}

func (r *assetRepository) GetByID(id int64) (*models.Asset, error) {
	var asset models.Asset
	if err := r.db.First(&asset, "id_asset = ?", id).Error; err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *assetRepository) Update(asset *models.Asset) error {
	// Use Updates to update all fields including zero values
	return r.db.Model(&models.Asset{}).Where("id_asset = ?", asset.IDAsset).Updates(asset).Error
}

func (r *assetRepository) Delete(id int64) error {
	return r.db.Delete(&models.Asset{}, "id_asset = ?", id).Error
}

// ExistsBySN returns true if an asset with the given serial number exists
func (r *assetRepository) ExistsBySN(itemSN string) (bool, error) {
	var count int64
	if err := r.db.Model(&models.Asset{}).
		Where("item_serial_number = ? AND item_serial_number != 'NO SN'", itemSN).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// MaxNoUrutForYear returns the current maximum no_urut for assets arriving in the given year
func (r *assetRepository) MaxNoUrutForYear(year string) (int, error) {
	var max int64
	if err := r.db.Model(&models.Asset{}).
		Select("COALESCE(MAX(no_urut), 0)").
		Where("tahun_barang_datang = ?", year).
		Scan(&max).Error; err != nil {
		return 0, err
	}
	return int(max), nil
}

// Implement validation methods
func (r *assetRepository) ValidatePIC(picName string) (bool, error) {
	var count int64
	err := r.db.Table("pic_id").Where("nama_pic = ?", picName).Count(&count).Error
	return count > 0, err
}

func (r *assetRepository) ValidateOwner(ownerName string) (bool, error) {
	var count int64
	err := r.db.Table("perusahaan").Where("nama_perusahaan = ?", ownerName).Count(&count).Error
	return count > 0, err
}

func (r *assetRepository) ValidateVendor(vendorName string) (bool, error) {
	var count int64
	err := r.db.Table("vendor").Where("nama_vendor = ?", vendorName).Count(&count).Error
	return count > 0, err
}

func (r *assetRepository) ValidateCategory(categoryName string) (bool, error) {
	var count int64
	err := r.db.Table("category").Where("nama_category = ?", categoryName).Count(&count).Error
	return count > 0, err
}

func (r *assetRepository) ValidateBrand(brandName string) (bool, error) {
	var count int64
	err := r.db.Table("brand").Where("nama_brand = ?", brandName).Count(&count).Error
	return count > 0, err
}

func (r *assetRepository) ValidateModel(modelName string) (bool, error) {
	var count int64
	err := r.db.Table("item_model").Where("nama_item_model = ?", modelName).Count(&count).Error
	return count > 0, err
}

func (r *assetRepository) ValidateDivisi(divisiName string) (bool, error) {
	var count int64
	err := r.db.Table("divisi").Where("nama_divisi = ?", divisiName).Count(&count).Error
	return count > 0, err
}
