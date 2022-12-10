package models

type ParamsUserSignUp struct {
	Name             string `json:"name" binding:"required"`
	OriginalPassword string `json:"original_password" binding:"required"`
	ConfirmPassword  string `json:"confirm_password" binding:"required"`
}
