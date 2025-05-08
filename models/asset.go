package models

import "time"

// Asset represents the asset entity
type Asset struct {
	IDAsset           uint      `gorm:"column:id_asset;primaryKey;autoIncrement" json:"id_asset"`
<<<<<<< HEAD
	NoBarcode         string    `gorm:"column:no_barcode" json:"no_barcode"`
=======
	NoBarcode         string    `gorm:"column:No_Barcode" json:"no_barcode"`
>>>>>>> 894af313cb71edf4680f49e3bf702b0e96b13672
	Owner             string    `gorm:"column:owner" json:"owner"`
	Category          string    `gorm:"column:category" json:"category"`
	ItemName          string    `gorm:"column:item_name" json:"item_name"`
	NamaBrand         string    `gorm:"column:Nama_Brand" json:"nama_brand"`
	NamaItemModel     string    `gorm:"column:Nama_Item_Model" json:"nama_item_model"`
<<<<<<< HEAD
	ItemSN            string    `gorm:"column:item_serial_number" json:"item_sn"`
=======
	ItemSN            string    `gorm:"column:Item_Serial_Number" json:"item_sn"`
>>>>>>> 894af313cb71edf4680f49e3bf702b0e96b13672
	Vendor            string    `gorm:"column:vendor" json:"vendor"`
	Tag               string    `gorm:"column:tag;type:text" json:"tag"`
	UnitPrice         int       `gorm:"column:Unit_Price" json:"unit_price"`
	PR                string    `gorm:"column:PR" json:"pr"`
	PO                string    `gorm:"column:PO" json:"po"`
	BarangDatang      string    `gorm:"column:Barang_datang" json:"barang_datang"`
	QTY               int       `gorm:"column:QTY" json:"qty"`
	Satuan            string    `gorm:"column:Satuan" json:"satuan"`
	PICID             string    `gorm:"column:PIC_id" json:"pic_id"`
	Divisi            string    `gorm:"column:Divisi" json:"divisi"`
	DefaultLocation   string    `gorm:"column:Default_location" json:"default_location"`
	CurrentLocation   string    `gorm:"column:Current_location" json:"current_location"`
	Status            string    `gorm:"column:Status" json:"status"`
	PurchaseYear      string    `gorm:"column:Purchase_year" json:"purchase_year"`
	NoUrut            int       `gorm:"column:no_urut" json:"no_urut"`
	Action            string    `gorm:"column:action" json:"action"`
	NamaFile          string    `gorm:"column:namafile" json:"nama_file"`
	Costudian         string    `gorm:"column:Costudian" json:"costudian"`
	NamaTugas         string    `gorm:"column:namaTugas" json:"nama_tugas"`
	TglCekout         string    `gorm:"column:tgl_cekout" json:"tgl_cekout"`
	Kondisi           string    `gorm:"column:kondisi" json:"kondisi"`
	JabatanCostudian  string    `gorm:"column:jabatan_costudian" json:"jabatan_costudian"`
	DeptCostudian     string    `gorm:"column:dept_costudian" json:"dept_costudian"`
	KodeCekout        string    `gorm:"column:kode_cekout" json:"kode_cekout"`
	NoWaCostudian     string    `gorm:"column:no_wa_costudian" json:"no_wa_costudian"`
	TahunBarangDatang string    `gorm:"column:tahun_barang_datang" json:"tahun_barang_datang"`
	NoNavision        string    `gorm:"column:no_navision" json:"no_navision"`
	Note              string    `gorm:"column:note;type:text" json:"note"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (Asset) TableName() string {
	return "asset"
}
