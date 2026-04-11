package service

import (
	models "asset-management/model"
	"asset-management/repository"
)

type MaintenanceService struct {
	repo *repository.MaintenanceRepository
}

func NewMaintenanceService(r *repository.MaintenanceRepository) *MaintenanceService {
	return &MaintenanceService{repo: r}
}

func (s *MaintenanceService) ListMaintenances(limit, offset int) ([]models.Maintenance, int64, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *MaintenanceService) CreateMaintenance(m *models.Maintenance) error {
	return s.repo.Create(m)
}

func (s *MaintenanceService) GetDetail(id string) (*models.Maintenance, error) {
	return s.repo.GetByID(id)
}

func (s *MaintenanceService) GetAssetHistory(assetID string) ([]models.Maintenance, error) {
	return s.repo.GetByAssetID(assetID)
}

func (s *MaintenanceService) UpdateMaintenance(m *models.Maintenance) error {
	return s.repo.Update(m)
}
