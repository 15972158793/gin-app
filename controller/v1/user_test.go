package v1

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/15972158793/gin-app/dao/redis"
	"github.com/15972158793/gin-app/pkg"
	"github.com/15972158793/gin-app/setting"

	"github.com/15972158793/gin-app/dao/mysql"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func _initConfig() {
	if err := setting.SetUp(); err != nil {
		fmt.Println("setting.SetUp() failed ...")
		return
	}

	pkg.SetUp()

	if err := mysql.SetUp(); err != nil {
		fmt.Println("mysql.SetUp() failed ...")
		return
	}
	defer mysql.Close()

	if err := redis.SetUp(); err != nil {
		fmt.Println("redis.SetUp() failed ...")
		return
	}
	defer redis.Close()
}

func TestUserLogin(t *testing.T) {

	// 涉及到各种库的
	_initConfig()

	gin.SetMode(gin.TestMode)
	url := "/api/v1/user/login"
	r := gin.New()
	r.POST(url, UserLogin)

	body := `{
        "user_name":"用户昵称6个字符",
        "password":"123456"
    }`

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
}

func TestUserSignUp(t *testing.T) {

}
