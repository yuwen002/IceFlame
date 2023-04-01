package uc_center

import (
	"context"
	"errors"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"
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

// CreateEmployee
//
// @Title 新建员工
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-30 00:11:48
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcEmployee) CreateEmployee(ctx context.Context, in system_master.CreateEmployeeInput) (code int32, message string, lastInsertId int64, err error) {
	// 验证用户手机号是否存在
	code, message, err = s.ExistsTel(ctx, in.Tel)
	if err != nil {
		return code, message, 0, err
	}
	if code == 0 {
		return 1, "用户名手机号已存在", 0, err
	}

	// 密码hash
	if in.Password == "" {
		in.Password = "123456"
	}
	in.Password, err = utility.PasswordHash(in.Password)
	if err != nil {
		return code, message, 0, err
	}

	err = dao.UcAccount.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 写入主表数据uc_account
		code, message, lastInsertId, err := utility.DBInsertAndGetId(dao.UcAccount.Ctx(ctx), utility.DBInsertInput{Data: g.Map{
			"username":      s.prefix + in.Tel,
			"password_hash": in.Password,
			"tel":           s.prefix + in.Tel,
		}})

		if err != nil {
			return err
		}

		if code == 1 {
			err = errors.New(message)
			return err
		}

		// 写入员工数据
		code, message, err = utility.DBInsert(dao.UcEmployee.Ctx(ctx), utility.DBInsertInput{Data: g.Map{
			"account_id": lastInsertId,
			"name":       in.Name,
			"tel":        in.Tel,
		}})
		if err != nil {
			return err
		}

		if code == 1 {
			err = errors.New(message)
			return err
		}

		return nil
	})

	if err != nil {
		return -1, "", 0, err
	}

	return 0, "数据写入成功", lastInsertId, nil
}
