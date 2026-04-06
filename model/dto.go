package models

type CreateAssetRequest struct {
	Code       string `json:"asset_code" validate:"required"`
	Name       string `json:"asset_name" validate:"required"`
	Status     string `json:"status" validate:"required"`
	Condition  string `json:"condition" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}
type CreateCategoryRequest struct {
	Name string `json:"category_name" validate:"required,min=3"`
}
type LoanRequest struct {
	AssetID    int `json:"asset_id" validate:"required"`
	EmployeeID int `json:"employee_id" validate:"required"`
}
