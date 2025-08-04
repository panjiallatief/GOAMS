package repository

import (
	"fmt"

	"gorm.io/gorm"
)

// DropdownRepository provides access to dropdown data
type DropdownRepository interface {
	GetDistinctPICs() ([]map[string]string, error)
	GetDistinctDivisions() ([]map[string]string, error)
	GetDistinctOwners() ([]map[string]string, error)
	GetDistinctVendors() ([]map[string]string, error)
	GetDistinctCategories() ([]map[string]string, error)
	GetDistinctBrands() ([]map[string]string, error)
	GetDistinctModels() ([]map[string]string, error)
}

type dropdownRepository struct {
	db *gorm.DB
}

// NewDropdownRepository creates a new instance of DropdownRepository
func NewDropdownRepository(db *gorm.DB) DropdownRepository {
	return &dropdownRepository{db: db}
}

func (r *dropdownRepository) GetDistinctPICs() ([]map[string]string, error) {
	var pics []map[string]string
	// Query from pic_id table - the column is actually called 'divisi' not 'nama_pic'
	rows, err := r.db.Table("pic_id").
		Select("id_pic_id, nama_pic").
		Order("nama_pic").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var namaPic string
		if err := rows.Scan(&id, &namaPic); err != nil {
			return nil, err
		}
		pics = append(pics, map[string]string{
			"id":       fmt.Sprintf("%d", id),
			"nama_pic": namaPic,
		})
	}
	return pics, nil
}

func (r *dropdownRepository) GetDistinctDivisions() ([]map[string]string, error) {
	var divisions []map[string]string
	// Query from divisi table
	rows, err := r.db.Table("divisi").
		Select("id_divisi, nama_divisi, barcode_nama_divisi, kode_divisi").
		Order("nama_divisi").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var namaDivisi string
		var barcodeNameDivisi *string
		var kodeDivisi *int

		if err := rows.Scan(&id, &namaDivisi, &barcodeNameDivisi, &kodeDivisi); err != nil {
			// If columns don't exist, try simpler query
			rows.Close()
			return r.getSimpleDivisions()
		}

		division := map[string]string{
			"id":           fmt.Sprintf("%d", id),
			"nama_divisi":  namaDivisi,
			"barcode_name": "",
			"kode_divisi":  "",
		}

		// Handle barcode name
		if barcodeNameDivisi != nil {
			division["barcode_name"] = *barcodeNameDivisi
		} else {
			division["barcode_name"] = fmt.Sprintf("%d", id)
		}

		// Handle kode divisi
		if kodeDivisi != nil {
			division["kode_divisi"] = fmt.Sprintf("%d", *kodeDivisi)
		} else {
			division["kode_divisi"] = fmt.Sprintf("%d", id)
		}

		divisions = append(divisions, division)
	}
	return divisions, nil
}

// Fallback for simpler divisi table structure
func (r *dropdownRepository) getSimpleDivisions() ([]map[string]string, error) {
	var divisions []map[string]string
	rows, err := r.db.Table("divisi").
		Select("id_divisi, nama_divisi").
		Order("nama_divisi").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var namaDivisi string
		if err := rows.Scan(&id, &namaDivisi); err != nil {
			return nil, err
		}
		divisions = append(divisions, map[string]string{
			"id":          fmt.Sprintf("%d", id),
			"nama_divisi": namaDivisi,
			"kode_divisi": fmt.Sprintf("%d", id), // Use ID as code if not available
		})
	}
	return divisions, nil
}

func (r *dropdownRepository) GetDistinctOwners() ([]map[string]string, error) {
	var owners []map[string]string
	// Query from perusahaan table
	rows, err := r.db.Table("perusahaan").
		Select("id_perusahaan, nama_perusahaan, barcode_nama_perusahaan").
		Order("nama_perusahaan").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var namaPerusahaan string
		var barcodeNamePerusahaan *string

		if err := rows.Scan(&id, &namaPerusahaan, &barcodeNamePerusahaan); err != nil {
			// If barcodename_perusahaan doesn't exist
			rows.Close()
			return r.getSimpleOwners()
		}

		owner := map[string]string{
			"id":              fmt.Sprintf("%d", id),
			"nama_perusahaan": namaPerusahaan,
		}

		if barcodeNamePerusahaan != nil {
			owner["barcode_name"] = *barcodeNamePerusahaan
		} else {
			// Generate barcode from name if not available
			if len(namaPerusahaan) >= 3 {
				owner["barcode_name"] = namaPerusahaan[:3]
			} else {
				owner["barcode_name"] = namaPerusahaan
			}
		}

		owners = append(owners, owner)
	}
	return owners, nil
}

// Fallback for simpler perusahaan table structure
func (r *dropdownRepository) getSimpleOwners() ([]map[string]string, error) {
	var owners []map[string]string
	rows, err := r.db.Table("perusahaan").
		Select("id_perusahaan, nama_perusahaan, barcode_nama_perusahaan").
		Order("nama_perusahaan").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var namaPerusahaan string
		if err := rows.Scan(&id, &namaPerusahaan); err != nil {
			return nil, err
		}

		var barcode_nama_perusahaan string
		if err := rows.Scan(&barcode_nama_perusahaan); err != nil {
			return nil, err
		}

		owners = append(owners, map[string]string{
			"id":              fmt.Sprintf("%d", id),
			"nama_perusahaan": namaPerusahaan,
			"barcode_name":    barcode_nama_perusahaan,
		})
	}
	return owners, nil
}

func (r *dropdownRepository) GetDistinctVendors() ([]map[string]string, error) {
	var vendors []map[string]string
	// Query from vendor table
	rows, err := r.db.Table("vendor").
		Select("id_vendor, nama_vendor").
		Order("nama_vendor").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var namaVendor string
		if err := rows.Scan(&id, &namaVendor); err != nil {
			return nil, err
		}
		vendors = append(vendors, map[string]string{
			"id":          fmt.Sprintf("%d", id),
			"nama_vendor": namaVendor,
		})
	}
	return vendors, nil
}

func (r *dropdownRepository) GetDistinctCategories() ([]map[string]string, error) {
	var categories []map[string]string
	// Query from category table
	rows, err := r.db.Table("category").
		Select("id_category, nama_category").
		Order("nama_category").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var namaCategory string
		if err := rows.Scan(&id, &namaCategory); err != nil {
			return nil, err
		}
		categories = append(categories, map[string]string{
			"id":            fmt.Sprintf("%d", id),
			"nama_category": namaCategory,
		})
	}
	return categories, nil
}

func (r *dropdownRepository) GetDistinctBrands() ([]map[string]string, error) {
	var brands []map[string]string
	// Query from brand table
	rows, err := r.db.Table("brand").
		Select("id_brand, nama_brand").
		Order("nama_brand").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var namaBrand string
		if err := rows.Scan(&id, &namaBrand); err != nil {
			return nil, err
		}
		brands = append(brands, map[string]string{
			"id":         fmt.Sprintf("%d", id),
			"nama_brand": namaBrand,
		})
	}
	return brands, nil
}

func (r *dropdownRepository) GetDistinctModels() ([]map[string]string, error) {
	var models []map[string]string
	// Query from item_model table
	rows, err := r.db.Table("item_model").
		Select("id_item, nama_item_model").
		Order("nama_item_model").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var namaItemModel string
		if err := rows.Scan(&id, &namaItemModel); err != nil {
			return nil, err
		}
		models = append(models, map[string]string{
			"id":              fmt.Sprintf("%d", id),
			"nama_item_model": namaItemModel,
		})
	}
	return models, nil
}
