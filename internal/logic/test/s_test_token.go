package test

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"ice_flame/internal/consts"
	"ice_flame/internal/model"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insJwtTokenTest = sJwtTokenTest{}

func NewJwtTokenTest() *sJwtTokenTest {
	return &insJwtTokenTest
}

func init() {
	service.RegisterJwtTokenTest(NewJwtTokenTest())
}

type sJwtTokenTest struct{}

// TestSetToken
//
// @Title 生成token
// @Description 根据条件生成token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-21 17:37:00
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sJwtTokenTest) TestSetToken(ctx context.Context, in model.JwtClaimsInput) (code int32, message string, output string, err error) {
	token, err := utility.CreateToken(utility.CustomClaims{
		Id:       in.Id,
		UserInfo: in.UserInfo,
		Expire:   in.Expire,
		Before:   in.Before,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  in.Issuer,
			Subject: in.Subject,
		},
	}, consts.USER_SECRET_KEY)

	if err != nil {
		return -1, "", "", err
	}

	return 0, "生成token成功", token, nil
}

// TestGetToken
//
// @Title 解析Token
// @Description 解析Token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-21 18:03:05
// @receiver s
// @param ctx
// @param token
// @return code
// @return message
// @return output
// @return err
func (s *sJwtTokenTest) TestGetToken(ctx context.Context, token string) (code int32, message string, output interface{}, err error) {
	claims, err := utility.ParseToken(token, consts.USER_SECRET_KEY)
	if err != nil {
		return -1, "", "", err
	}

	return 0, "生成token成功", claims, nil
}
