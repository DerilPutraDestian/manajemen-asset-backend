package repository

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type LoanRepository interface {
	Create(loan *models.AssetLoan) error
	FindActiveLoan(assetID int) (models.AssetLoan, error)
	Update(loan *models.AssetLoan) error
}

type loanRepo struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepo{db}
}

func (r *loanRepo) Create(loan *models.AssetLoan) error {
	return r.db.Create(loan).Error
}

func (r *loanRepo) FindActiveLoan(assetID int) (models.AssetLoan, error) {
	var loan models.AssetLoan
	err := r.db.Where("asset_id = ? AND status = ?", assetID, "borrowed").
		First(&loan).Error
	return loan, err
}

func (r *loanRepo) Update(loan *models.AssetLoan) error {
	return r.db.Save(loan).Error
}
