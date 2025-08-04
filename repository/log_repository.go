package repository

import (
	"github.com/panjiallatief/GOAMS/models"
	"gorm.io/gorm"
)

// LogRepository defines CRUD operations for Log
type LogRepository interface {
	Create(log *models.Log) error
	GetByAssetID(assetID uint) ([]models.Log, error)
	GetAll() ([]models.Log, error)
}

type logRepository struct {
	db *gorm.DB
}

// NewLogRepository creates a new instance of LogRepository
func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{db: db}
}

func (r *logRepository) Create(log *models.Log) error {
	return r.db.Create(log).Error
}

func (r *logRepository) GetByAssetID(assetID uint) ([]models.Log, error) {
	var logs []models.Log
	if err := r.db.Where("id_asset = ?", assetID).Order("created_at DESC").Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *logRepository) GetAll() ([]models.Log, error) {
	var logs []models.Log
	if err := r.db.Order("created_at DESC").Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
