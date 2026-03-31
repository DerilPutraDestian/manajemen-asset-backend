package migrations

import (
	"asset-management/models"
	"fmt"

	"gorm.io/gorm"
)

// =======================
// MIGRATE (CREATE TABLE)
// =======================
func Migrate(db *gorm.DB) error {
	fmt.Println("Running migrations...")

	// 🔥 Disable FK sementara (biar aman)
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	err := db.AutoMigrate(
		// ✅ PARENT
		&models.Employee{},
		&models.Category{},

		// ✅ CHILD
		&models.Asset{},
		&models.AssetLoan{},
		&models.AssetHistory{},
	)

	if err != nil {
		return err
	}

	// 🔥 Enable lagi FK
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	fmt.Println("Migration success")
	return nil
}

// =======================
// DROP ALL TABLES
// =======================
func DropAll(db *gorm.DB) error {
	fmt.Println("Dropping tables...")

	// 🔥 Disable FK dulu
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	err := db.Migrator().DropTable(
		// 🔻 CHILD dulu (biar tidak error FK)
		&models.AssetHistory{},
		&models.AssetLoan{},
		&models.Asset{},

		// 🔻 PARENT
		&models.Category{},
		&models.Employee{},
	)

	if err != nil {
		return err
	}

	// 🔥 Enable lagi FK
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	fmt.Println("Drop success")
	return nil
}
