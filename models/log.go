package models

import "time"

// Log represents the activity log for assets
type Log struct {
	IDLog        uint      `gorm:"column:id_log;primaryKey;autoIncrement" json:"id"`
	AssetIDAsset uint      `gorm:"column:asset_id_asset" json:"asset_id"`
	Action       string    `gorm:"column:action" json:"action"`
	Nama         string    `gorm:"column:nama" json:"nama"`
	Costudian    string    `gorm:"column:costudian" json:"costudian"`
	Details      string    `gorm:"column:details" json:"details"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName sets the insert table name for this struct type
func (Log) TableName() string {
	return "log"
}
