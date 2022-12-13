package controller

import (
	"go.uber.org/zap"

	"github.com/15972158793/gin-app/models"

	"github.com/15972158793/gin-app/service"

	"github.com/gin-gonic/gin"
)

// UserSignUp
// @Summary 用户注册
// @Tags user
// @Description User sign up
// @Produce json
// @Param name formData string true "Name"
// @Param original_password formData string true "OriginalPassword"
// @Param confirm_password formData string true "ConfirmPassword"
// @Router /api/v1/user/sign_up [post]

func UserSignUp(c *gin.Context) {

	//参数校验
	params := new(models.ParamsUserSignUp)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error("ParamsUserSignUp invalid failed ... \n")
		ResponseError(c, CODE_INVALID_PARAMS)
		return
	}
	//业务逻辑
	if err := service.UserSignUp(params); err != nil {
		ResponseError(c, CODE_NORMAL_SQL_ERROR)
		return
	}
	//返回响应
	ResponseSuccess(c, "注册成功")
}

func UserLogin(c *gin.Context) {

	//参数校验
	params := new(models.ParamsUserLogin)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error("ParamsUserLogin invalid failed ... \n")
		ResponseError(c, CODE_INVALID_PARAMS)
		return
	}

	// 处理业务
	result, err := service.UserLogin(params)
	if err != nil {
		ResponseError(c, CODE_NORMAL_SQL_ERROR)
		return
	}

	//返回响应
	ResponseSuccess(c, result)
}

// UserInfo 根据token获取用户信息
func UserInfo(c *gin.Context) {

	ResponseSuccess(c, "获取成功")
}
