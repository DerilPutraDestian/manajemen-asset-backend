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
		Preload("Employee").       // Load data Employee yang meminjam
		Limit(limit).
		Offset(offset).
		Find(&assetsLoan).Error

	return assetsLoan, count, err
}

func (r *AssetLoanRepository) GetByID(id string) (*models.AssetLoan, error) {
	var data models.AssetLoan
	// Di sini juga sebaiknya tambahkan Preload("Asset.Category") agar data lengkap saat ambil detail
	err := r.db.Preload("Asset").Preload("Asset.Category").Preload("Employee").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *AssetLoanRepository) Create(loan *models.AssetLoan) error {
	return r.db.Omit("Asset", "Employee").Create(loan).Error
}

func (r *AssetLoanRepository) Update(loan *models.AssetLoan) error {
	return r.db.Save(loan).Error
}
