package dto

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" form:"password" binding:"required" validate:"required"`
}
