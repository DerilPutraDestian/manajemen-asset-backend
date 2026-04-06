package service

import (
	models "asset-management/model"
	"asset-management/repository"
)

type AssetLoanService struct {
	repo *repository.AssetLoanRepository
}

func NewAssetLoanService(repo *repository.AssetLoanRepository) *AssetLoanService {
	return &AssetLoanService{repo: repo}
}
func (s *AssetLoanService) ListLoans(assetLoanCode, search string, limit, offset int) ([]models.AssetLoan, int64, error) {
	return s.repo.GetAll(assetLoanCode, search, limit, offset)
}
func (s *AssetLoanService) GetLoanByID(id string) (*models.AssetLoan, error) {
	return s.repo.GetByID(id)
}
func (s *AssetLoanService) CreateLoan(l *models.AssetLoan) error {
	return s.repo.Create(l)
}

func (s *AssetLoanService) UpdateLoan(l *models.AssetLoan) error {
	return s.repo.Update(l)
}
