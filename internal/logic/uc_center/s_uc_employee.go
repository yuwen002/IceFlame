package uc_center

import (
	"context"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insUcEmployee = sUcEmployee{
	prefix:      "EMP_",
	redisConfig: "uc_center",
}

func UcEmployee() *sUcEmployee {
	return &insUcEmployee
}

func init() {
	service.RegisterUcEmployee(UcEmployee())
}

type sUcEmployee struct {
	prefix      string
	redisConfig string
}

// ExistsTel
//
// @Title 验证员工手机号是否存在
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-29 17:46:42
// @receiver s
// @param ctx
// @param tel
// @return code
// @return message
// @return err
func (s *sUcEmployee) ExistsTel(ctx context.Context, tel string) (code int32, message string, err error) {
	tel = s.prefix + tel
	data := utility.RedisExistsData{
		Config: s.redisConfig,
		Key:    dao.UcAccount.Table() + ":exists_tel",
		Value:  tel,
	}

	code, message, err = utility.RCExistsSetData(data, func(condition interface{}) (code int32, message string, err error) {
		code, message, _, err = utility.DBGetOneMapByWhere(dao.UcAccount.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "id",
			Where: "tel=?",
			Args:  condition,
		})

		return code, message, err
	})

	return code, message, err
}

func (s *sUcEmployee) CreateEmployee(ctx context.Context, in system_master.CreateEmployeeInput) (code int32, message string, err error) {
	// 验证用户手机号是否存在
	code, message, err = s.ExistsTel(ctx, in.Tel)
	if err != nil {
		return code, message, err
	}
	if code == 0 {
		return 1, "用户名手机号已存在", err
	}
}
