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

func NewAssetService(repo *repository.AssetRepository, historyRepo repository.HistoryRepository) *AssetService {
	return &AssetService{
		repo:        repo,
		historyRepo: historyRepo,
	}
}

// =========================
// CREATE
// =========================
func (s *AssetService) CreateAsset(asset *models.Asset) error {
	// 1. Generate QR Code
	fileName := fmt.Sprintf("%s.png", asset.AssetCode)
	uploadDir := "./public/qrcodes/"
	filePath := uploadDir + fileName

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		_ = os.MkdirAll(uploadDir, os.ModePerm)
	}

	if err := qrcode.WriteFile(asset.AssetCode, qrcode.Medium, 256, filePath); err != nil {
		return fmt.Errorf("gagal generate qr code: %v", err)
	}
	asset.QRCode = filePath

	// 2. Simpan Asset
	if err := s.repo.Create(asset); err != nil {
		return err
	}

	// 3. Simpan History (Initial Status)
	_ = s.historyRepo.Create(&models.AssetHistory{
		AssetID:   asset.ID,
		OldStatus: "", // Tidak ada status lama karena baru dibuat
		NewStatus: asset.Status,
		Note:      "Initial creation",
	})

	return nil
}

// =========================
// UPDATE
// =========================
func (s *AssetService) UpdateAsset(asset *models.Asset) error {
	// 1. Ambil data lama untuk mencatat OldStatus di history
	oldAsset, err := s.repo.GetByID(asset.ID)
	if err != nil {
		return fmt.Errorf("asset tidak ditemukan: %v", err)
	}
	oldStatus := oldAsset.Status

	// 2. Update data asset
	if err := s.repo.Update(asset); err != nil {
		return err
	}

	// 3. Simpan history jika ada perubahan status
	_ = s.historyRepo.Create(&models.AssetHistory{
		AssetID:   asset.ID,
		OldStatus: oldStatus,
		NewStatus: asset.Status,
		Note:      "Asset information updated",
	})

	return nil
}

// =========================
// DELETE
// =========================
func (s *AssetService) Delete(asset *models.Asset) error {
	// 1. HAPUS HISTORY DULU (Untuk menghindari Error 1451)
	// Pastikan repository history kamu punya method DeleteByAssetID
	if err := s.historyRepo.Delete(asset.ID); err != nil {
		return fmt.Errorf("gagal menghapus history asset: %v", err)
	}

	// 2. Hapus File QR Code
	fileName := fmt.Sprintf("./public/qrcodes/%s.png", asset.AssetCode)
	_ = os.Remove(fileName)

	// 3. Hapus Asset
	if err := s.repo.Delete(asset); err != nil {
		return err
	}

	return nil
}

// GET Methods tetap sama
func (s *AssetService) ListAssets(assetCode, search string, limit, offset int) ([]models.Asset, int64, error) {
	return s.repo.GetAll(assetCode, search, limit, offset)
}

func (s *AssetService) GetAsset(id string) (*models.Asset, error) {
	return s.repo.GetByID(id)
}
