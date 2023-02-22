package uc_center

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insUcSystemMaster = sUcSystemMaster{
	prefix: "SA_",
}

func UcSystemMaster() *sUcSystemMaster {
	return &insUcSystemMaster
}

func init() {
	service.RegisterUcSystemMaster(UcSystemMaster())
}

// sUcSystemMaster
// @Description: 后台管理员管理结构体
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 16:06:45
type sUcSystemMaster struct {
	prefix string
}

// ExistsUsername
//
// @Title 验证用户名是否存在
// @Description 验证用户名是否存在
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 17:16:27
// @receiver s
// @param ctx
// @param username
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) ExistsUsername(ctx context.Context, username string) (code int32, message string, err error) {
	username = s.prefix + username
	data := utility.RedisExistsData{
		Config: "uc_center",
		Key:    dao.UcAccount.Table() + ":exists_username",
		Value:  username,
	}

	code, message, err = utility.RCExistsSetData(data, func(condition interface{}) (code int32, message string, err error) {
		code, message, _, err = utility.DBGetOneMapByWhere(dao.UcAccount.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "id",
			Where: "username=?",
			Args:  username,
		})

		return code, message, err
	})

	return code, message, err
}

// CreateSystemMaster
//
// @Title 新建管理员用户
// @Description 新建管理员用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 17:16:46
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) CreateSystemMaster(ctx context.Context, in uc_center.CreateSystemMasterInput) (code int32, message string, err error) {
	// 验证用户名是否存在
	code, message, err = s.ExistsUsername(ctx, in.Tel)
	if err != nil {
		return code, message, err
	}
	if code == 1 {
		return 1, "用户名已存在", err
	}

	// 无密码，默认密码123456
	if in.Password == "" {
		in.Password, err = utility.PasswordHash("123456")
		if err != nil {
			return code, message, err
		}
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

		// 写入管理员表数据
		code, message, err = utility.DBInsert(dao.UcSystemMaster.Ctx(ctx), utility.DBInsertInput{
			Data: g.Map{
				"account_id": lastInsertId,
				"name":       in.Name,
				"tel":        in.Tel,
			},
		})

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
		return -1, "", err
	}

	return 0, "数据写入成功", nil
}