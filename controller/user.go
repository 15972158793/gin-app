package controller

import (
	"net/http"

	"github.com/15972158793/gin-app/service"

	"github.com/gin-gonic/gin"
)

// @Summary 用户注册
// @Tags tag
// @Description User sign up
// @Produce json
// @Param name formData string true "Name"
// @Param orignal_password formData string "OrignalPassword"
// @Param confirm_password formData string "ConfirmPassword"
// @Router /api/v1/user/sign_up [post]
func UserSignUp(c *gin.Context) {

	//参数校验

	//业务逻辑
	service.UserSignUp()

	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}

func UserLogin(c *gin.Context) {

}
