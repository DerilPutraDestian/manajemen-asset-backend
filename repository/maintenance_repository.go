package repository

import (
	models "asset-management/model"

	"gorm.io/gorm"
)

type MaintenanceRepository struct {
	db *gorm.DB
}

func NewMaintenanceRepository(db *gorm.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}

// Menampilkan histori maintenance berdasarkan Asset tertentu
func (r *MaintenanceRepository) GetByAsset(assetID string, limit, offset int) ([]models.Maintenance, int64, error) {
	var data []models.Maintenance
	var total int64

	query := r.db.Model(&models.Maintenance{}).Where("asset_id = ?", assetID)

	err := query.Count(&total).
		Preload("Asset").
		Preload("User"). // Asumsi ReportedBy merujuk ke tabel User
		Limit(limit).
		Offset(offset).
		Order("created_at DESC"). // Histori terbaru di atas
		Find(&data).Error

	return data, total, err
}

// Mengambil satu data maintenance untuk kebutuhan Update/Detail
func (r *MaintenanceRepository) GetByID(id string) (*models.Maintenance, error) {
	var m models.Maintenance
	err := r.db.Preload("Asset").First(&m, "id = ?", id).Error
	return &m, err
}

func (r *MaintenanceRepository) Create(m *models.Maintenance) error {
	return r.db.Create(m).Error
}

func (r *MaintenanceRepository) Update(m *models.Maintenance) error {
	// Menggunakan .Select untuk menentukan kolom mana saja yang BOLEH diubah.
	// Kolom asset_id dan reported_by TIDAK dimasukkan agar tidak kena cek Foreign Key.
	return r.db.Model(m).Select("description", "status", "start_date", "end_date", "updated_at").Updates(m).Error
}
