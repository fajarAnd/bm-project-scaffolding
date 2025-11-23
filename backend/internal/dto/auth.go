package dto

import "github.com/baramulti/ticketing-system/backend/internal/models"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role,omitempty"` // defaults to "user" if empty
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}