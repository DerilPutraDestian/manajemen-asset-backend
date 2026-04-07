package routes

import (
	"asset-management/handlers"
	"asset-management/middleware" // Pastikan folder middleware sudah ada
	"asset-management/repository"
	"asset-management/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userSvc)

	// Asset & Category
	assetRepo := repository.NewAssetRepository(db)
	historyRepo := repository.NewHistoryRepository(db)
	assetSvc := service.NewAssetService(assetRepo, historyRepo)
	assetHandler := handlers.NewAssetHandler(assetSvc)

	catRepo := repository.NewCategoryRepository(db)
	catSvc := service.NewCategoryService(catRepo)
	catHandler := handlers.NewCategoryHandler(catSvc)

	// Employee (DATA BARU)
	empRepo := repository.NewEmployeeRepository(db)
	empSvc := service.NewEmployeeService(empRepo)
	empHandler := handlers.NewEmployeeHandler(empSvc)

	// Loan & Maintenance
	loanRepo := repository.NewAssetLoanRepository(db)
	loanSvc := service.NewAssetLoanService(loanRepo)
	loanHandler := handlers.NewLoanHandler(loanSvc)

	mtRepo := repository.NewMaintenanceRepository(db)
	mtSvc := service.NewMaintenanceService(mtRepo)
	mtHandler := handlers.NewMaintenanceHandler(mtSvc)

	api := app.Group("/api")
	api.Post("/login", userHandler.Login)
	api.Use(middleware.JWTMiddleware())

	employees := api.Group("/employees")
	employees.Get("/", empHandler.Index)
	employees.Post("/", empHandler.Store)

	// --- ASSETS ---
	assets := api.Group("/assets")
	assets.Get("/", assetHandler.Index)
	assets.Get("/:id", assetHandler.Show)
	assets.Post("/", assetHandler.Store)    // Biasanya Admin
	assets.Put("/:id", assetHandler.Update) // Biasanya Admin
	assets.Delete("/:id", assetHandler.Delete)

	// --- CATEGORIES ---
	categories := api.Group("/categories")
	categories.Get("/", catHandler.Index)
	categories.Get("/:id", catHandler.Show)
	categories.Post("/", catHandler.Store)
	categories.Put("/:id", catHandler.Update)
	categories.Delete("/:id", catHandler.Delete)

	// --- LOANS (PEMINJAMAN) ---
	loans := api.Group("/loans")
	loans.Get("/", loanHandler.Index)
	loans.Post("/", loanHandler.Store)
	loans.Put("/:id", loanHandler.Update) // Digunakan untuk Return Asset

	// --- MAINTENANCE ---
	maintenances := api.Group("/maintenances")
	maintenances.Get("/", mtHandler.Index)
	maintenances.Post("/", mtHandler.Store)
	maintenances.Put("/:id", mtHandler.Update)
}
