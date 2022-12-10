package controller

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

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
	var params models.ParamsUserSignUp
	if err := c.ShouldBindJSON(&params); err != nil {
		zap.L().Error("ParamsUserSignUp invalid failed ... \n")

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans),
		})
		return
	}

	fmt.Println(params)

	//业务逻辑
	service.UserSignUp()

	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}

func UserLogin(c *gin.Context) {

}
