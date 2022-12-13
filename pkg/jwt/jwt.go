package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecret = []byte("d4tg2dySDFhtdDFGG4w4yw4")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int64, userName string) (tokenString string, err error) {
	claim := MyClaims{
		UserID:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
			Issuer:    "yukeeok",
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(mySecret)
	return tokenString, err
}

func ParseToken(tokenStr string) (*MyClaims, error) {
	mc := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil // 这是我的secret
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("token invalid")
}
