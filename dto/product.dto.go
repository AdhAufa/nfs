package dto

type ProductRequest struct {
	Name  string `json:"name" form:"name" binding:"required" validate:"required"`
	Price uint64 `json:"price" form:"email" binding:"required" validate:"required"`
}
