package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

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
	// Validate foreign key references
	validationErrors := []string{}

	if asset.PicID != "" {
		if valid, err := s.repo.ValidatePIC(asset.PicID); err != nil {
			return fmt.Errorf("error validating PIC: %w", err)
		} else if !valid {
			validationErrors = append(validationErrors, fmt.Sprintf("PIC '%s' tidak ditemukan", asset.PicID))
		}
	}

	if asset.Owner != "" {
		if valid, err := s.repo.ValidateOwner(asset.Owner); err != nil {
			return fmt.Errorf("error validating Owner: %w", err)
		} else if !valid {
			validationErrors = append(validationErrors, fmt.Sprintf("Owner '%s' tidak ditemukan", asset.Owner))
		}
	}

	if asset.Vendor != "" {
		if valid, err := s.repo.ValidateVendor(asset.Vendor); err != nil {
			return fmt.Errorf("error validating Vendor: %w", err)
		} else if !valid {
			validationErrors = append(validationErrors, fmt.Sprintf("Vendor '%s' tidak ditemukan", asset.Vendor))
		}
	}

	if asset.Category != "" {
		if valid, err := s.repo.ValidateCategory(asset.Category); err != nil {
			return fmt.Errorf("error validating Category: %w", err)
		} else if !valid {
			validationErrors = append(validationErrors, fmt.Sprintf("Category '%s' tidak ditemukan", asset.Category))
		}
	}

	if asset.NamaBrand != "" {
		if valid, err := s.repo.ValidateBrand(asset.NamaBrand); err != nil {
			return fmt.Errorf("error validating Brand: %w", err)
		} else if !valid {
			validationErrors = append(validationErrors, fmt.Sprintf("Brand '%s' tidak ditemukan", asset.NamaBrand))
		}
	}

	if asset.NamaItemModel != "" {
		if valid, err := s.repo.ValidateModel(asset.NamaItemModel); err != nil {
			return fmt.Errorf("error validating Model: %w", err)
		} else if !valid {
			validationErrors = append(validationErrors, fmt.Sprintf("Model '%s' tidak ditemukan", asset.NamaItemModel))
		}
	}

	if asset.Divisi != "" {
		if valid, err := s.repo.ValidateDivisi(asset.Divisi); err != nil {
			return fmt.Errorf("error validating Divisi: %w", err)
		} else if !valid {
			validationErrors = append(validationErrors, fmt.Sprintf("Divisi '%s' tidak ditemukan", asset.Divisi))
		}
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("validation errors: %s", strings.Join(validationErrors, ", "))
	}

	// Set default values
	if asset.Action == "" {
		asset.Action = "Cek In"
	}
	if asset.Kondisi == "" {
		asset.Kondisi = "Bagus"
	}
	if asset.NamaTugas == "" {
		asset.NamaTugas = "System"
	}

	// Set nama file
	if asset.NamaFile == "" && asset.PO != "" && asset.PR != "" {
		asset.NamaFile = asset.PO + "_" + asset.PR
	}

	// Convert date format from YYYY-MM-DD to DD/MM/YYYY for barang_datang
	if asset.BarangDatang != "" && strings.Contains(asset.BarangDatang, "-") {
		parts := strings.Split(asset.BarangDatang, "-")
		if len(parts) == 3 {
			asset.BarangDatang = parts[2] + "/" + parts[1] + "/" + parts[0]
			// Set tahun_barang_datang from year (last 2 digits)
			if asset.TahunBarangDatang == "" && len(parts[0]) >= 2 {
				asset.TahunBarangDatang = parts[0][2:]
			}
		}
	}

	// Check serial number uniqueness
	if asset.ItemSerialNumber != "" &&
		!strings.EqualFold(asset.ItemSerialNumber, "NO SN") &&
		!strings.EqualFold(asset.ItemSerialNumber, "no sn") &&
		strings.TrimSpace(asset.ItemSerialNumber) != "" {
		exists, err := s.repo.ExistsBySN(asset.ItemSerialNumber)
		if err != nil {
			return err
		}
		if exists {
			return ErrSNAlreadyExists
		}
	}

	// Set timestamps
	now := time.Now()
	asset.CreatedAt = &now
	asset.UpdatedAt = &now

	// Get max no_urut for the year
	max, err := s.repo.MaxNoUrutForYear(asset.TahunBarangDatang)
	if err != nil {
		return err
	}
	asset.NoUrut = int64(max + 1)

	// Generate barcode
	asset.NoBarcode = s.generateBarcode(asset)

	return s.repo.Create(asset)
}

func (s *assetService) generateBarcode(asset *models.Asset) string {
	// Extract date code from tahun_barang_datang or barang_datang
	dateCode := asset.TahunBarangDatang
	if dateCode == "" && asset.BarangDatang != "" {
		// Try to extract from barang_datang (format: DD/MM/YYYY)
		parts := strings.Split(asset.BarangDatang, "/")
		if len(parts) == 3 && len(parts[2]) >= 2 {
			dateCode = parts[2][2:] // Last 2 digits of year
		}
	}

	// For the barcode, we need to extract the proper codes from the database values
	// Based on your data, it seems like the barcode format should be:
	// OWNER_CODE + DIVISI_CODE + YEAR + NO_URUT

	// Example from your data: "FMN10222100136"
	// FMN = First Media News (owner code)
	// 10 = Division code
	// 22 = Year (2022)
	// 2100136 = Sequential number

	ownerCode := "FMN" // Default for First Media News
	if asset.Owner == "BERITA SATU MEDIA HOLDING" {
		ownerCode = "BER"
	}

	// Extract division code - you might need a lookup table for this
	divisiCode := "10" // Default
	if strings.Contains(asset.Divisi, "8,2") {
		divisiCode = "8,2"
	}

	// Format no_urut with leading zeros
	noUrutStr := fmt.Sprintf("%04d", asset.NoUrut)

	// Combine to create barcode
	barcode := ownerCode + divisiCode + dateCode + noUrutStr

	return barcode
}

func getCodeFromString(s string, length int) string {
	// Simple function to get a code from a string
	// You might want to implement a more sophisticated mapping
	code := strings.ToUpper(strings.ReplaceAll(s, " ", ""))
	if len(code) > length {
		return code[:length]
	}
	// Pad with zeros if shorter
	for len(code) < length {
		code += "0"
	}
	return code
}

func (s *assetService) GetAllAssets() ([]models.Asset, error) {
	assets, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	
	// Convert []*models.Asset to []models.Asset
	result := make([]models.Asset, len(assets))
	for i, asset := range assets {
		result[i] = *asset
	}
	return result, nil
}

func (s *assetService) GetAssetByID(id uint) (*models.Asset, error) {
	return s.repo.GetByID(int64(id))
}

// UpdateAsset checks SN uniqueness and updates the asset.
func (s *assetService) UpdateAsset(asset *models.Asset) error {
	// fetch existing to compare serial numbers
	existing, err := s.repo.GetByID(asset.IDAsset)
	if err != nil {
		return err
	}

	// if SN changed, enforce uniqueness
	if asset.ItemSerialNumber != existing.ItemSerialNumber &&
		asset.ItemSerialNumber != "" &&
		!strings.EqualFold(asset.ItemSerialNumber, "NO SN") {
		exists, err := s.repo.ExistsBySN(asset.ItemSerialNumber)
		if err != nil {
			return err
		}
		if exists {
			return ErrSNAlreadyExists
		}
	}

	// Update timestamp
	now := time.Now()
	asset.UpdatedAt = &now

	// Preserve created_at
	asset.CreatedAt = existing.CreatedAt

	return s.repo.Update(asset)
}

func (s *assetService) DeleteAsset(id uint) error {
	asset, err := s.repo.GetByID(int64(id))
	if err != nil {
		return err
	}
	return s.repo.Delete(asset.IDAsset)
}

// You'll need to pass the barcode values from the frontend
// For now, let's create a request structure that includes the barcode data
type CreateAssetRequest struct {
	*models.Asset
	DivisiCode     string `json:"divisi_code"`
	DivisiBarcode  string `json:"divisi_barcode"`
	OwnerBarcodeID string `json:"owner_barcode_id"`
}
