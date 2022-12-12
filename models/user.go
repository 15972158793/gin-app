package models

// 负责处理参数验证

type ParamsUserSignUp struct {
	UserName         string `json:"user_name" binding:"required,min=6"`
	OriginalPassword string `json:"original_password" binding:"required,min=6,max=20"`
	ConfirmPassword  string `json:"confirm_password" binding:"required,min=6,max=20,eqfield=OriginalPassword"`
}

type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
}

type ParamsUserLogin struct {
	UserName string `json:"user_name" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}
