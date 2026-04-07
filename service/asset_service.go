package service

import (
	models "asset-management/model"
	"asset-management/repository"
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

type AssetService struct {
	repo        *repository.AssetRepository
	historyRepo repository.HistoryRepository
}

// 🔥 FIX: tambahkan historyRepo di constructor
func NewAssetService(repo *repository.AssetRepository, historyRepo repository.HistoryRepository) *AssetService {
	return &AssetService{
		repo:        repo,
		historyRepo: historyRepo,
	}
}

// =========================
// GET
// =========================
func (s *AssetService) ListAssets(assetCode, search string, limit, offset int) ([]models.Asset, int64, error) {
	return s.repo.GetAll(assetCode, search, limit, offset)
}

func (s *AssetService) GetAsset(id string) (*models.Asset, error) {
	return s.repo.GetByID(id)
}

// =========================
// CREATE
// =========================
func (s *AssetService) CreateAsset(asset *models.Asset) error {
	fileName := fmt.Sprintf("%s.png", asset.AssetCode)
	uploadDir := "./public/qrcodes/"
	filePath := uploadDir + fileName

	// pastikan folder ada
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		_ = os.MkdirAll(uploadDir, os.ModePerm)
	}

	// generate QR
	err := qrcode.WriteFile(asset.AssetCode, qrcode.Medium, 256, filePath)
	if err != nil {
		return fmt.Errorf("gagal generate qr code: %v", err)
	}

	asset.QRCode = filePath

	// 🔥 simpan asset dulu
	if err := s.repo.Create(asset); err != nil {
		return err
	}

	// 🔥 baru simpan history (SETELAH berhasil)
	_ = s.historyRepo.Create(&models.AssetHistory{
		AssetID:     asset.ID,
		Action:      "create",
		Description: "Asset created",
	})

	return nil
}

// =========================
// UPDATE
// =========================
func (s *AssetService) UpdateAsset(asset *models.Asset) error {
	if err := s.repo.Update(asset); err != nil {
		return err
	}

	// 🔥 history
	_ = s.historyRepo.Create(&models.AssetHistory{
		AssetID:     asset.ID,
		Action:      "update",
		Description: "Asset updated",
	})

	return nil
}

// =========================
// DELETE
// =========================
func (s *AssetService) DeleteAsset(asset *models.Asset) error {
	fileName := fmt.Sprintf("./public/qrcodes/%s.png", asset.AssetCode)
	_ = os.Remove(fileName)

	if err := s.repo.Delete(asset); err != nil {
		return err
	}

	// 🔥 history
	_ = s.historyRepo.Create(&models.AssetHistory{
		AssetID:     asset.ID,
		Action:      "delete",
		Description: "Asset deleted",
	})

	return nil
}
