package utility

import (
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// CustomClaimsInput
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-21 23:39:02
type CustomClaimsInput struct {
	Id            uint64                 // 用户ID
	UserInfo      map[string]interface{} // 用户信息
	Expire        int64                  // 过期时间
	Before        int64                  // 生效时间
	Issuer        string                 // 签发人
	Subject       string                 // 签发主题
	SigningMethod jwt.SigningMethod
}

// CustomClaims
// @Description: 自定义输入信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-21 16:21:09
type CustomClaims struct {
	Id       uint64                 `json:"user_id"`
	UserInfo map[string]interface{} `json:"user_info"`
	jwt.RegisteredClaims
}

// JwtClaims
// @Description:Jwt  用户签名验证
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-21 10:32:32
type JwtClaims struct {
	SecretKey string
}

// VerifySecretKey
//
// @Title 验证SecretKey
// @Description 验证SecretKey
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-21 16:53:10
// @receiver jc
// @return error
func (jc *JwtClaims) VerifySecretKey() error {
	if jc.SecretKey == "" {
		err := errors.New("密钥信息错误")
		return err
	}
	return nil
}

// CreateToken
//
// @Title 创建token
// @Description 创建Token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-21 15:13:27
// @receiver jc
// @param in
// @return string
// @return error
func (jc *JwtClaims) CreateToken(in CustomClaimsInput) (string, error) {
	err := jc.VerifySecretKey()
	if err != nil {
		return "", err
	}

	if in.Id == 0 {
		err := errors.New("ID信息错误")
		return "", err
	}

	// 判断过期时间
	var expiresAt *jwt.NumericDate
	if in.Expire > 0 {
		expiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(gconv.Int64(in.Expire)) * 24 * time.Hour))
	} else {
		expiresAt = jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour * time.Duration(1)))
	}

	// 签发时间
	now := jwt.NewNumericDate(time.Now())

	// 判断生效时间
	var notBefore *jwt.NumericDate
	if in.Before > 0 {
		notBefore = jwt.NewNumericDate(time.Now().Add(time.Duration(in.Before) * 24 * time.Hour))
	} else {
		notBefore = now
	}

	var signingMethod jwt.SigningMethod
	if in.SigningMethod == nil {
		signingMethod = jwt.SigningMethodHS256
	}

	customClaims := CustomClaims{
		Id:       in.Id,
		UserInfo: in.UserInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    in.Issuer,
			Subject:   in.Subject,
			Audience:  nil,
			ExpiresAt: expiresAt,
			NotBefore: notBefore,
			IssuedAt:  now,
		},
	}

	claims := jwt.NewWithClaims(signingMethod, customClaims)

	return claims.SignedString([]byte(jc.SecretKey))
}

// ParseToken
//
// @Title 解析Token
// @Description 解析Token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-21 15:15:26
// @receiver jc
// @param token
// @return jwt.MapClaims
// @return error
func (jc *JwtClaims) ParseToken(token string) (map[string]interface{}, error) {
	err := jc.VerifySecretKey()
	if err != nil {
		return nil, err
	}

	withClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jc.SecretKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token不正确")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已经过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token未激活")
			} else {
				return nil, errors.New("token错误")
			}
		}
	}

	if claims, ok := withClaims.Claims.(*CustomClaims); ok && withClaims.Valid {
		return gconv.Map(claims), nil
	}

	return nil, errors.New("无法处理此token")
}

// CreateToken
//
// @Title 创建token
// @Description 创建token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-21 17:31:05
// @param in
// @param secretKey
// @return string
// @return error
func CreateToken(in CustomClaimsInput, secretKey string) (string, error) {
	jwt := JwtClaims{SecretKey: secretKey}
	return jwt.CreateToken(in)
}

// ParseToken
//
// @Title 解析Token
// @Description 解析Token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-21 18:01:45
// @param token
// @param secretKey
// @return map[string]interface{}
// @return error
func ParseToken(token string, secretKey string) (map[string]interface{}, error) {
	jwt := JwtClaims{SecretKey: secretKey, Sig}
	return jwt.ParseToken(token)
}
