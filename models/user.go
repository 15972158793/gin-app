package models

type ParamsUserSignUp struct {
	Name             string `json:"name" binding:"required"`
	OriginalPassword string `json:"original_password" binding:"required,min=6,max=20"`
	ConfirmPassword  string `json:"confirm_password" binding:"required,min=6,max=20,eqfield=OriginalPassword"`
}

type User struct {
	UserID   int64
	Name     string
	Password string
}
