package models

import "time"

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required,min=2,max=50"`
	LastName  string `json:"last_name" binding:"required,min=2,max=50"`
	Email     string `json:"email" binding:"required,email,min=15,max=100"`
	Password  string `json:"password" binding:"required,min=6,max=16"`
}

type VerifyRequest struct {
	Code  string `json:"code" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type AuthResponse struct {
	ID          int64
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	AccessToken string    `json:"accesss_token"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email,min=15,max=100"`
	Password string `json:"password" binding:"required"`
}
