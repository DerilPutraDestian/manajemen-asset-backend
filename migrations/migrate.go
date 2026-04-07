package migrations

import (
	models "asset-management/model"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println("Running migrations...")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Asset{},
		&models.Maintenance{},
		&models.AssetLoan{},
		&models.AssetHistory{},
		&models.Employee{},
	)

	if err != nil {
		return err
	}

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	fmt.Println("Migration success")
	return nil
}

func DropAll(db *gorm.DB) error {
	fmt.Println("Dropping tables...")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	err := db.Migrator().DropTable(
		&models.AssetHistory{},
		&models.AssetLoan{},
		&models.Maintenance{},
		&models.Asset{},
		&models.Category{},
		&models.User{},
		&models.Employee{},
	)

	if err != nil {
		return err
	}

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	fmt.Println("Drop success")
	return nil
}
