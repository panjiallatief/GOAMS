package service

import (
	"github.com/panjiallatief/GOAMS/repository"
)

// DropdownService provides dropdown data
type DropdownService interface {
	GetAllPICs() ([]map[string]string, error)
	GetAllDivisions() ([]map[string]string, error)
	GetAllOwners() ([]map[string]string, error)
	GetAllVendors() ([]map[string]string, error)
	GetAllCategories() ([]map[string]string, error)
	GetAllBrands() ([]map[string]string, error)
	GetAllModels() ([]map[string]string, error)
}

type dropdownService struct {
	repo repository.DropdownRepository
}

// NewDropdownService creates a new instance of DropdownService
func NewDropdownService(repo repository.DropdownRepository) DropdownService {
	return &dropdownService{repo: repo}
}

func (s *dropdownService) GetAllPICs() ([]map[string]string, error) {
	return s.repo.GetDistinctPICs()
}

func (s *dropdownService) GetAllDivisions() ([]map[string]string, error) {
	return s.repo.GetDistinctDivisions()
}

func (s *dropdownService) GetAllOwners() ([]map[string]string, error) {
	return s.repo.GetDistinctOwners()
}

func (s *dropdownService) GetAllVendors() ([]map[string]string, error) {
	return s.repo.GetDistinctVendors()
}

func (s *dropdownService) GetAllCategories() ([]map[string]string, error) {
	return s.repo.GetDistinctCategories()
}

func (s *dropdownService) GetAllBrands() ([]map[string]string, error) {
	return s.repo.GetDistinctBrands()
}

func (s *dropdownService) GetAllModels() ([]map[string]string, error) {
	return s.repo.GetDistinctModels()
}
