package routes

import (
	"asset-management/handlers"
	"asset-management/repository"
	"asset-management/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// =========================
	// USER & AUTH
	// =========================
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userSvc)

	// Public Route untuk Login
	app.Post("/api/login", userHandler.Login)

	user := app.Group("/api/users")
	user.Get("/", userHandler.Index)
	user.Post("/", userHandler.Store)
	user.Get("/:id", userHandler.Show)
	user.Put("/:id", userHandler.Update)
	user.Delete("/:id", userHandler.Delete)

	assetRepo := repository.NewAssetRepository(db)
	assetSvc := service.NewAssetService(assetRepo)
	assetHandler := handlers.NewAssetHandler(assetSvc)

	asset := app.Group("/api/assets")
	asset.Get("/", assetHandler.Index)
	asset.Get("/:id", assetHandler.Show)
	asset.Post("/", assetHandler.Store)
	asset.Put("/:id", assetHandler.Update)
	asset.Delete("/:id", assetHandler.Delete)

	catRepo := repository.NewCategoryRepository(db)
	catSvc := service.NewCategoryService(catRepo)
	catHandler := handlers.NewCategoryHandler(catSvc)

	category := app.Group("/api/categories")
	category.Get("/", catHandler.Index)
	category.Get("/:id", catHandler.Show)
	category.Post("/", catHandler.Store)
	category.Put("/:id", catHandler.Update)
	category.Delete("/:id", catHandler.Delete)

	// =========================
	// LOAN (PEMINJAMAN)
	// =========================
	loanRepo := repository.NewAssetLoanRepository(db)
	loanSvc := service.NewAssetLoanService(loanRepo)
	loanHandler := handlers.NewLoanHandler(loanSvc)

	loan := app.Group("/api/loans")
	loan.Get("/", loanHandler.Index)
	loan.Post("/", loanHandler.Store)
	loan.Put("/:id", loanHandler.Update)

	// =========================
	// MAINTENANCE
	// =========================
	mtRepo := repository.NewMaintenanceRepository(db)
	mtSvc := service.NewMaintenanceService(mtRepo)
	mtHandler := handlers.NewMaintenanceHandler(mtSvc)

	maintenance := app.Group("/api/maintenances")
	maintenance.Get("/", mtHandler.Index)
	maintenance.Post("/", mtHandler.Store)
	maintenance.Get("/:id", mtHandler.Update) // Bisa digunakan untuk detail
	maintenance.Put("/:id", mtHandler.Update)
	// Delete DIHAPUS sesuai prinsip Audit Trail (Histori Permanen)
}
