package v1

import (
	"github.com/15972158793/gin-app/controller"
	"go.uber.org/zap"

	"github.com/15972158793/gin-app/models"

	"github.com/15972158793/gin-app/service"

	"github.com/gin-gonic/gin"
)

// UserSignUp
// @Summary 用户注册
// @Tags user
// @Description 用户注册路由
// @Accept json
// @Produce json
// @Param name formData string true "昵称"
// @Param original_password formData string true "输入密码"
// @Param confirm_password formData string true "再次输入密码"
// @Router /api/v1/user/sign_up [post]
func UserSignUp(c *gin.Context) {

	//参数校验
	params := new(models.ParamsUserSignUp)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error("ParamsUserSignUp invalid failed ... \n")
		controller.ResponseError(c, controller.CODE_INVALID_PARAMS)
		return
	}
	//业务逻辑
	if err := service.UserSignUp(params); err != nil {
		controller.ResponseError(c, controller.CODE_NORMAL_SQL_ERROR)
		return
	}
	//返回响应
	controller.ResponseSuccess(c, "注册成功")
}

func UserLogin(c *gin.Context) {

	//参数校验
	params := new(models.ParamsUserLogin)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error("ParamsUserLogin invalid failed ... \n")
		controller.ResponseError(c, controller.CODE_INVALID_PARAMS)
		return
	}

	// 处理业务
	result, err := service.UserLogin(params)
	if err != nil {
		controller.ResponseError(c, controller.CODE_NORMAL_SQL_ERROR)
		return
	}

	//返回响应
	controller.ResponseSuccess(c, result)
}

// UserInfo 根据token获取用户信息
func UserInfo(c *gin.Context) {

	controller.ResponseSuccess(c, "获取成功")
}
