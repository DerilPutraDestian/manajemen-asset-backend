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
	// Gunakan FirstOrCreate agar seeder bisa dijalankan berkali-kali tanpa error duplicate
	if err := db.Where(models.Category{Name: "Electronics"}).FirstOrCreate(&category).Error; err != nil {
		return err
	}

	// ========================
	// 2. USER (DENGAN BCRYPT)
	// ========================
	// Password harus di-hash agar bisa login nanti
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
	// 3. ASSET
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
		// Cek berdasarkan AssetCode agar tidak double saat seeding ulang
		if err := db.Where(models.Asset{AssetCode: assets[i].AssetCode}).FirstOrCreate(&assets[i]).Error; err != nil {
			return err
		}
	}

	// ========================
	// 4. ASSET LOAN
	// ========================
	loan := models.AssetLoan{
		AssetID:  assets[0].ID,
		UserID:   user.ID,
		LoanDate: time.Now(),
		Status:   "borrowed",
	}

	// Hanya buat loan jika belum ada data loan untuk asset tersebut
	var count int64
	db.Model(&models.AssetLoan{}).Where("asset_id = ? AND status = ?", assets[0].ID, "borrowed").Count(&count)
	if count == 0 {
		if err := db.Create(&loan).Error; err != nil {
			return err
		}
	}

	now := time.Now()
	maintenance := models.Maintenance{
		AssetID:     assets[1].ID,
		ReportedBy:  user.ID,
		Description: "Tinta macet dan perlu pembersihan head",
		Status:      "pending",
		StartDate:   &now, // Mulai besok
	}

	db.Model(&models.Maintenance{}).Where("asset_id = ? AND status = ?", assets[1].ID, "pending").Count(&count)
	if count == 0 {
		db.Create(&maintenance)
	}

	fmt.Println("Seeding success!")
	return nil
}
