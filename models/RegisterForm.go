package models

type RegisterForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"username" binding:"required"`
	Email    string `json:"username" binding:"required"`
}
