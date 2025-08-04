package models

import "time"

// Asset represents the asset entity matching the actual database schema
type Asset struct {
	IDAsset           int64      `gorm:"primaryKey;column:id_asset" json:"id_asset"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
	BarangDatang      string     `gorm:"column:barang_datang" json:"barang_datang"`
	Costudian         string     `gorm:"column:costudian" json:"costudian"`
	CurrentLocation   string     `gorm:"column:current_location" json:"current_location"`
	DefaultLocation   string     `gorm:"column:default_location" json:"default_location"`
	ItemSerialNumber  string     `gorm:"column:item_serial_number" json:"item_serial_number"`
	ItemName          string     `gorm:"column:item_name" json:"item_name"`
	NoBarcode         string     `gorm:"column:no_barcode" json:"no_barcode"`
	PicID             string     `gorm:"column:pic_id" json:"pic_id"`
	PO                string     `gorm:"column:po" json:"po"`
	PR                string     `gorm:"column:pr" json:"pr"`
	PurchaseYear      string     `gorm:"column:purchase_year" json:"purchase_year"`
	Qty               int        `gorm:"column:qty" json:"qty"`
	Satuan            string     `gorm:"column:satuan" json:"satuan"`
	Status            string     `gorm:"column:status" json:"status"`
	UnitPrice         int        `gorm:"column:unit_price" json:"unit_price"`
	Action            string     `gorm:"column:action" json:"action"`
	DeptCostudian     string     `gorm:"column:dept_costudian" json:"dept_costudian"`
	JabatanCostudian  string     `gorm:"column:jabatan_costudian" json:"jabatan_costudian"`
	KodeCekout        string     `gorm:"column:kode_cekout" json:"kode_cekout"`
	Kondisi           string     `gorm:"column:kondisi" json:"kondisi"`
	NamaTugas         string     `gorm:"column:nama_tugas" json:"nama_tugas"`
	NamaFile          string     `gorm:"column:namafile" json:"namafile"`
	NoUrut            int64      `gorm:"column:no_urut" json:"no_urut"`
	NoWaCostudian     string     `gorm:"column:no_wa_costudian" json:"no_wa_costudian"`
	TahunBarangDatang string     `gorm:"column:tahun_barang_datang" json:"tahun_barang_datang"`
	TglCekout         string     `gorm:"column:tgl_cekout" json:"tgl_cekout"`
	Owner             string     `gorm:"column:owner" json:"owner"`
	Vendor            string     `gorm:"column:vendor" json:"vendor"`
	Category          string     `gorm:"column:category" json:"category"`
	NamaBrand         string     `gorm:"column:nama_brand" json:"nama_brand"`
	NamaItemModel     string     `gorm:"column:nama_item_model" json:"nama_item_model"`
	Divisi            string     `gorm:"column:divisi" json:"divisi"`
	Tag               string     `gorm:"column:tag" json:"tag"`
	NoNavision        string     `gorm:"column:no_navision" json:"no_navision"`
	Note              string     `gorm:"column:note" json:"note"`
}

// TableName sets the insert table name for this struct type
func (Asset) TableName() string {
	return "asset"
}
