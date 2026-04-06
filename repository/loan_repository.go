package repository

import (
	models "asset-management/model"

	"gorm.io/gorm"
)

type AssetLoanRepository struct {
	db *gorm.DB
}

func NewAssetLoanRepository(db *gorm.DB) *AssetLoanRepository {
	return &AssetLoanRepository{db: db}
}

func (r *AssetLoanRepository) GetAll(assetLoanCode, search string, limit, offset int) ([]models.AssetLoan, int64, error) {
	var assetsLoan []models.AssetLoan
	var count int64

	query := r.db.Model(&models.AssetLoan{})

	if assetLoanCode != "" {
		query = query.Where("asset_loan_code = ?", assetLoanCode)
	}

	// Gunakan search jika ada (opsional, tergantung kebutuhanmu)
	if search != "" {
		// Contoh: mencari berdasarkan nama asset melalui join atau search langsung jika ada fieldnya
	}

	err := query.Count(&count).
		Preload("Asset").          // Load data Asset
		Preload("Asset.Category"). // Load data Category MILIK Asset (Nested Preload)
		Preload("User").           // Load data User yang meminjam
		Limit(limit).
		Offset(offset).
		Find(&assetsLoan).Error

	return assetsLoan, count, err
}

func (r *AssetLoanRepository) GetByID(id string) (*models.AssetLoan, error) {
	var data models.AssetLoan
	// Di sini juga sebaiknya tambahkan Preload("Asset.Category") agar data lengkap saat ambil detail
	err := r.db.Preload("Asset").Preload("Asset.Category").Preload("User").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *AssetLoanRepository) Create(l *models.AssetLoan) error {
	if err := r.db.Create(l).Error; err != nil {
		return err
	}
	// Ambil ulang agar respons JSON langsung berisi data Asset, Category, dan User
	return r.db.Preload("Asset").Preload("Asset.Category").Preload("User").First(l, "id = ?", l.ID).Error
}

func (r *AssetLoanRepository) Update(l *models.AssetLoan) error {
	return r.db.Save(l).Error
}
