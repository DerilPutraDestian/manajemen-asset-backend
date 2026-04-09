package repository

import (
	models "asset-management/model"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) GetAll(search string, limit, offset int) ([]models.Employee, int64, error) {
	var employees []models.Employee
	var total int64

	query := r.db.Model(&models.Employee{})

	if search != "" {
		// Menggunakan kolom employee_name sesuai database kamu
		query = query.Where("employee_name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.Count(&total).Limit(limit).Offset(offset).Find(&employees).Error
	return employees, total, err
}

func (r *EmployeeRepository) GetByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	// Mencari berdasarkan primary key (employee_id)
	err := r.db.First(&employee, id).Error
	return &employee, err
}

func (r *EmployeeRepository) Create(e *models.Employee) error {
	return r.db.Create(e).Error
}

func (r *EmployeeRepository) Update(e *models.Employee) error {
	return r.db.Save(e).Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Employee{}, id).Error
}
