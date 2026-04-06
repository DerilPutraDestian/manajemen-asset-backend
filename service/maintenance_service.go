package service

import (
	models "asset-management/model"
	"asset-management/repository"
	"errors"
)

type MaintenanceService struct {
	repo *repository.MaintenanceRepository
}

func NewMaintenanceService(repo *repository.MaintenanceRepository) *MaintenanceService {
	return &MaintenanceService{repo: repo}
}

func (s *MaintenanceService) ListByAsset(assetID string, limit, offset int) ([]models.Maintenance, int64, error) {
	return s.repo.GetByAsset(assetID, limit, offset)
}

func (s *MaintenanceService) GetMaintenance(id string) (*models.Maintenance, error) {
	return s.repo.GetByID(id)
}

func (s *MaintenanceService) CreateMaintenance(m *models.Maintenance) error {
	// Default status saat pembuatan laporan
	if m.Status == "" {
		m.Status = "pending"
	}
	return s.repo.Create(m)
}

func (s *MaintenanceService) UpdateMaintenance(m *models.Maintenance) error {
	// Contoh Logika Bisnis: Jika status 'completed', pastikan EndDate terisi
	if m.Status == "completed" && m.EndDate.IsZero() {
		return errors.New("end date is required when maintenance is completed")
	}

	return s.repo.Update(m)
}
