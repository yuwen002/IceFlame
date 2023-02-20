package utility

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtClaims struct {
	Id        uint64                 `json:"id" :"id"`
	UserInfo  map[string]interface{} `json:"user_info"`
	SecretKey string
}

func (jc *JwtClaims) CreateToken(in map[string]interface{}) {
	now := jwt.NewNumericDate(time.Now())
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":       jc.Id,
		"UserInfo": jc.UserInfo,
		"RegisteredClaims": jwt.RegisteredClaims{
			Issuer:    "", // 签发人
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour * time.Duration(1))),
			NotBefore: now, // 在此之前不可用
			IssuedAt:  now, //签发时间
		},
	})

}
