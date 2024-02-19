package util

import (
	"errors"
	"fuxiaochen-api-with-go/global"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int64 `json:"userID"`
	jwt.RegisteredClaims       // 内嵌标准的声明
}

func NewToken(id int64) (string, error) {
	claims := CustomClaims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(global.Conf.JWTConfig.ExpireDuration * time.Hour)),
			Issuer:    "a123",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 注意：这里需要传 []byte 不能直接传 string，否则报错：key is of invalid type
	return token.SignedString([]byte(global.Conf.JWTConfig.Secret))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 这里需要返回 []byte 不能直接传 string
		return []byte(global.Conf.JWTConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的 token")
}
