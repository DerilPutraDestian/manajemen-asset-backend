package service

import (
	"asset-management/models"
	"asset-management/repository"
	"fmt"
	"time"
)

type LoanService interface {
	Borrow(req models.LoanRequest) error
	Return(assetID int) error
}

type loanService struct {
	loanRepo    repository.LoanRepository
	assetRepo   repository.AssetRepository
	historyRepo repository.HistoryRepository
}

func NewLoanService(
	l repository.LoanRepository,
	a repository.AssetRepository,
	h repository.HistoryRepository,
) LoanService {
	return &loanService{
		loanRepo:    l,
		assetRepo:   a,
		historyRepo: h,
	}
}

// 🔥 BORROW (PINJAM ASSET)
func (s *loanService) Borrow(req models.LoanRequest) error {

	// 1. ambil asset
	asset, err := s.assetRepo.FindByID(req.AssetID)
	if err != nil {
		return fmt.Errorf("asset not found")
	}

	// 2. validasi tidak boleh dipinjam lagi
	if asset.Status == "borrowed" {
		return fmt.Errorf("asset already borrowed")
	}

	oldStatus := asset.Status

	// 3. update status asset
	asset.Status = "borrowed"
	if err := s.assetRepo.Update(&asset); err != nil {
		return err
	}

	// 4. create loan
	loan := models.AssetLoan{
		AssetID:    req.AssetID,
		EmployeeID: req.EmployeeID,
		LoanDate:   time.Now(),
		Status:     "borrowed",
	}

	if err := s.loanRepo.Create(&loan); err != nil {
		return err
	}

	// 5. create history
	history := models.AssetHistory{
		AssetID:   req.AssetID,
		OldStatus: oldStatus,
		NewStatus: "borrowed",
		ChangedAt: time.Now(),
	}

	return s.historyRepo.Create(&history)
}

// 🔥 RETURN (KEMBALIKAN ASSET)
func (s *loanService) Return(assetID int) error {

	// 1. ambil asset
	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		return fmt.Errorf("asset not found")
	}

	// 2. cari loan aktif
	loan, err := s.loanRepo.FindActiveLoan(assetID)
	if err != nil {
		return fmt.Errorf("active loan not found")
	}

	// 3. update loan
	loan.Status = "returned"
	loan.ReturnDate = time.Now()

	if err := s.loanRepo.Update(&loan); err != nil {
		return err
	}

	// 4. update asset
	oldStatus := asset.Status
	asset.Status = "available"

	if err := s.assetRepo.Update(&asset); err != nil {
		return err
	}

	// 5. history
	history := models.AssetHistory{
		AssetID:   assetID,
		OldStatus: oldStatus,
		NewStatus: "available",
		ChangedAt: time.Now(),
	}

	return s.historyRepo.Create(&history)
}
