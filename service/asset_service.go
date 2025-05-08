package service

import (
	"errors"
	"fmt"

	"github.com/panjiallatief/GOAMS/models"
	"github.com/panjiallatief/GOAMS/repository"
)

var ErrSNAlreadyExists = errors.New("serial number already exists")

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

// CreateAsset checks SN uniqueness, generates barcode, and creates the asset.
func (s *assetService) CreateAsset(asset *models.Asset) error {
	// check serial number uniqueness
	exists, err := s.repo.ExistsBySN(asset.ItemSN)
	if err != nil {
		return err
	}
	if exists {
		return ErrSNAlreadyExists
	}
	// compute next no_urut for the year
	max, err := s.repo.MaxNoUrutForYear(asset.TahunBarangDatang)
	if err != nil {
		return err
	}
	asset.NoUrut = max + 1
	// append zero-padded counter to existing NoBarcode prefix
	asset.NoBarcode = asset.NoBarcode + fmt.Sprintf("%04d", asset.NoUrut)
	return s.repo.Create(asset)
}

func (s *assetService) GetAllAssets() ([]models.Asset, error) {
	return s.repo.GetAll()
}

func (s *assetService) GetAssetByID(id uint) (*models.Asset, error) {
	return s.repo.GetByID(id)
}

// UpdateAsset checks SN uniqueness and updates the asset.
func (s *assetService) UpdateAsset(asset *models.Asset) error {
	// fetch existing to compare serial numbers
	existing, err := s.repo.GetByID(asset.IDAsset)
	if err != nil {
		return err
	}
	// if SN changed, enforce uniqueness
	if asset.ItemSN != existing.ItemSN {
		exists, err := s.repo.ExistsBySN(asset.ItemSN)
		if err != nil {
			return err
		}
		if exists {
			return ErrSNAlreadyExists
		}
	}
	return s.repo.Update(asset)
}

func (s *assetService) DeleteAsset(id uint) error {
	asset, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(asset)
}
