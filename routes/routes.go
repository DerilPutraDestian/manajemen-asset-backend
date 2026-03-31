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
	// ASSET
	// =========================
	assetRepo := repository.NewAssetRepository(db)
	assetSvc := service.NewAssetService(assetRepo)
	assetHandler := handlers.NewAssetHandler(assetSvc)

	asset := app.Group("/api/assets")

	asset.Get("/", assetHandler.GetAll)
	asset.Get("/:id", assetHandler.GetByID)
	asset.Post("/", assetHandler.Create)
	asset.Put("/:id", assetHandler.Update)
	asset.Delete("/:id", assetHandler.Delete)

	catRepo := repository.NewCategoryRepository(db)
	catSvc := service.NewCategoryService(catRepo)
	catHandler := handlers.NewCategoryHandler(catSvc)

	category := app.Group("/api/categories")

	category.Get("/", catHandler.GetAll)
	category.Get("/:id", catHandler.GetByID)
	category.Post("/", catHandler.Create)
	category.Put("/:id", catHandler.Update)
	category.Delete("/:id", catHandler.Delete)

	// =========================
	// LOAN + HISTORY
	// =========================
	loanRepo := repository.NewLoanRepository(db)
	historyRepo := repository.NewHistoryRepository(db)

	loanSvc := service.NewLoanService(loanRepo, assetRepo, historyRepo)
	loanHandler := handlers.NewLoanHandler(loanSvc)

	loan := app.Group("/api/loans")

	loan.Post("/borrow", loanHandler.Borrow)
	loan.Post("/return/:asset_id", loanHandler.Return)
}
