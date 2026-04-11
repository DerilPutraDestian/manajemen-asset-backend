package seeders

import (
	models "asset-management/model"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	fmt.Println("Seeding database...")

	// ========================
	// 1. CATEGORY
	// ========================
	category := models.Category{
		Name: "Electronics",
	}
	if err := db.Where(models.Category{Name: "Electronics"}).FirstOrCreate(&category).Error; err != nil {
		return err
	}

	// ========================
	// 2. USER (ADMIN)
	// ========================
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	user := models.User{
		Name:     "Deril Admin",
		Email:    "deril@mail.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}
	if err := db.Where(models.User{Email: "deril@mail.com"}).FirstOrCreate(&user).Error; err != nil {
		return err
	}

	// ========================
	// 3. EMPLOYEE (Penting: Untuk relasi AssetLoan)
	// ========================
	employee := models.Employee{
		Name:  "Budi Santoso",
		Email: "budi@mail.com",
	}
	// Pastikan employee dibuat dulu agar ID-nya tersedia untuk AssetLoan
	if err := db.Where(models.Employee{Email: "budi@mail.com"}).FirstOrCreate(&employee).Error; err != nil {
		return err
	}

	// ========================
	// 4. ASSET
	// ========================
	assets := []models.Asset{
		{
			AssetCode:  "AST001",
			Name:       "Laptop Asus ExpertBook",
			Status:     "active",
			CategoryID: category.ID,
		},
		{
			AssetCode:  "AST002",
			Name:       "Printer Canon Pixma",
			Status:     "active",
			CategoryID: category.ID,
		},
	}

	for i := range assets {
		if err := db.Where(models.Asset{AssetCode: assets[i].AssetCode}).FirstOrCreate(&assets[i]).Error; err != nil {
			return err
		}
	}

	// ========================
	// 5. ASSET LOAN
	// ========================
	loan := models.AssetLoan{
		AssetID:    assets[0].ID,
		EmployeeID: employee.ID, // ID dari employee yang dibuat di langkah 3
		LoanDate:   time.Now(),
		Status:     "borrowed",
	}

	var count int64
	db.Model(&models.AssetLoan{}).Where("asset_id = ? AND status = ?", assets[0].ID, "borrowed").Count(&count)
	if count == 0 {
		if err := db.Create(&loan).Error; err != nil {
			return err
		}
	}

	// ========================
	// 6. MAINTENANCE
	// ========================
	now := time.Now()
	maintenance := models.Maintenance{
		AssetID:           assets[1].ID,
		IssueDescription:  "Tinta macet dan perlu pembersihan head",
		MaintenanceStatus: "pending",
		StartDate:         &now,
	}

	db.Model(&models.Maintenance{}).Where("asset_id = ? AND maintenance_status = ?", assets[1].ID, "pending").Count(&count)
	if count == 0 {
		if err := db.Create(&maintenance).Error; err != nil {
			return err
		}
	}

	fmt.Println("Seeding success!")
	return nil
}
