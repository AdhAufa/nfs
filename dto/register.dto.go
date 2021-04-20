package dto

type RegisterRequest struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" form:"password" binding:"required" validate:"required"`
}
