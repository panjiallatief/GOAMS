package service

import (
	"time"

	"github.com/panjiallatief/GOAMS/models"
	"github.com/panjiallatief/GOAMS/repository"
)

// LogService defines business logic operations for Log
type LogService interface {
	CreateLog(assetID uint, action, user, details string) error
	GetAssetLogs(assetID uint) ([]models.Log, error)
	GetAllLogs() ([]models.Log, error)
}

type logService struct {
	repo repository.LogRepository
}

// NewLogService creates a new instance of LogService
func NewLogService(repo repository.LogRepository) LogService {
	return &logService{repo: repo}
}

func (s *logService) CreateLog(assetID uint, action, user, details string) error {
	log := &models.Log{
		AssetIDAsset: assetID,
		Action:       action,
		Nama:         user,
		Costudian:    user, // Assuming costudian is the same as user for simplicity
		// If costudian is different, you can pass it as a parameter
		Details:   details,
		CreatedAt: time.Now(),
	}
	return s.repo.Create(log)
}

func (s *logService) GetAssetLogs(assetID uint) ([]models.Log, error) {
	return s.repo.GetByAssetID(assetID)
}

func (s *logService) GetAllLogs() ([]models.Log, error) {
	return s.repo.GetAll()
}
