package seeders

import (
	"asset-management/models"
	"fmt"
	"time"

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
	if err := db.Create(&category).Error; err != nil {
		return err
	}

	// ========================
	// 2. EMPLOYEE
	// ========================
	employee := models.Employee{
		Name:  "Deril",
		Email: "deril@mail.com",
		Phone: "08123456789",
	}
	if err := db.Create(&employee).Error; err != nil {
		return err
	}

	// ========================
	// 3. ASSET
	// ========================
	assets := []models.Asset{
		{
			Code:       "AST001",
			Name:       "Laptop Asus",
			Status:     "available",
			Condition:  "good",
			CategoryID: category.CategoryID,
		},
		{
			Code:       "AST002",
			Name:       "Printer Canon",
			Status:     "maintenance",
			Condition:  "fair",
			CategoryID: category.CategoryID,
		},
	}

	for i := range assets {
		if err := db.Create(&assets[i]).Error; err != nil {
			return err
		}
	}

	// ========================
	// 4. ASSET LOAN
	// ========================
	loan := models.AssetLoan{
		AssetID:    assets[0].AssetID,
		EmployeeID: employee.EmployeeID,
		LoanDate:   time.Now(),
		Status:     "borrowed",
	}
	if err := db.Create(&loan).Error; err != nil {
		return err
	}

	fmt.Println("Seeding success ✅")
	return nil
}
