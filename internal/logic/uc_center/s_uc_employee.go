package uc_center

import (
	"context"
	"errors"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/util/grand"

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
		Key:    dao.UcAccount.Table() + ":exists_employee_tel",
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

// ExistsInviteCode
//
// @Title 验证邀请码是否存在
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 23:53:05
// @receiver s
// @param inviteCode
// @return code
// @return message
// @return err
func (s *sUcEmployee) ExistsInviteCode(ctx context.Context, inviteCode string) (code int32, message string, err error) {
	data := utility.RedisExistsData{
		Config: s.redisConfig,
		Key:    dao.UcEmployee.Table() + ":exists_invite_code",
		Value:  inviteCode,
	}

	code, message, err = utility.RCExistsSetData(data, func(condition interface{}) (code int32, message string, err error) {
		code, message, _, err = utility.DBGetOneMapByWhere(dao.UcEmployee.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "invite_code",
			Where: "invite_code = ?",
			Args:  condition,
		})

		return code, message, err
	})

	return code, message, err
}

// AccountIdCastTel
//
// @Title account id转换员工注册电话
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-05 00:35:28
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return output
// @return err
func (s *sUcEmployee) AccountIdCastTel(ctx context.Context, accountId uint64) (code int32, message string, output interface{}, err error) {
	code, message, output, err = utility.RCCastHashData(utility.RedisCastData{
		Config: s.redisConfig,
		Key:    dao.UcEmployee.Table() + ":account_id_cast_tel",
		Field:  gconv.String(accountId),
	}, func(condition interface{}) (code int32, message string, output interface{}, err error) {
		code, message, out, err := utility.DBGetMapById(dao.UcEmployee.Ctx(ctx), utility.DBGetByIdInput{
			Field: "tel",
			Where: "account_id=?",
			Args:  condition,
		})

		return code, message, out["tel"], err
	})

	return code, message, output, err
}

// UpdateAccountIdCastTel
//
// @Title 更新account id转换信息的电话
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-05 00:46:16
// @receiver s
// @param ctx
// @param accountId
// @param tel
// @return code
// @return message
// @return err
func (s *sUcEmployee) UpdateAccountIdCastTel(ctx context.Context, accountId uint64, tel string) (code int32, message string, err error) {
	code, message, err = utility.RCUpdateCastHashData(utility.RedisCastData{
		Config: s.redisConfig,
		Key:    dao.UcEmployee.Table() + ":account_id_cast_tel",
		Field:  gconv.String(accountId),
		Data:   tel,
	})

	return code, message, err
}

// CreateEmployeeRole
//
// @Title 新建员工角色
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-01 23:51:42
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcEmployee) CreateEmployeeRole(ctx context.Context, in system_master.CreateEmployeeRoleInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.UcEmployeeRole.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetEmployeeRoleById
//
// @Title 按ID获取员工角色
// @Description 按ID获取员工角色
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-01 23:58:55
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return out
// @return err
func (s *sUcEmployee) GetEmployeeRoleById(ctx context.Context, id uint16) (code int32, message string, out *system_master.GetEmployeeRoleOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.UcEmployeeRole.Ctx(ctx), utility.DBGetByIdInput{Where: id}, &out)
	return code, message, out, err
}

// ModifyEmployeeRoleById
//
// @Title 按ID修改员工角色
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:05:28
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcEmployee) ModifyEmployeeRoleById(ctx context.Context, in system_master.ModifyRoleByIdInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.UcEmployeeRole.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

// ListEmployeeRole
//
// @Title 员工角色列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:12:14
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return out
// @return err
func (s *sUcEmployee) ListEmployeeRole(ctx context.Context, in system_master.ListEmployeeRoleInput) (code int32, message string, out []*system_master.GetEmployeeRoleOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.UcEmployeeRole.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "id asc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &out)

	return code, message, out, err
}

// CreateEmployeeRoleRelation
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 01:18:56
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcEmployee) CreateEmployeeRoleRelation(ctx context.Context, in system_master.CreateEmployeeRoleRelationInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.UcEmployeeRoleRelation.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// CreateEmployee
//
// @Title 新建员工角色绑定员工ID
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
	inviteCode := grand.Digits(6)
	code, message, err = s.ExistsInviteCode(ctx, inviteCode)
	if code == 0 || code >= 2 {
		return s.CreateEmployee(ctx, in)
	}

	// 验证用户手机号是否存在
	code, message, err = s.ExistsTel(ctx, in.Tel)
	if err != nil {
		return code, message, 0, err
	}
	if code == 0 || code >= 2 {
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
		code, message, accountId, err := utility.DBInsertAndGetId(dao.UcAccount.Ctx(ctx), utility.DBInsertInput{Data: g.Map{
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
			"account_id":  accountId,
			"name":        in.Name,
			"tel":         in.Tel,
			"invite_code": inviteCode,
		}})
		if err != nil {
			return err
		}

		if code == 1 {
			err = errors.New(message)
			return err
		}

		lastInsertId = accountId
		return nil
	})

	if err != nil {
		return -1, "", 0, err
	}

	return 0, "数据写入成功", lastInsertId, nil
}

// ModifyEmployeeByAccountId
//
// @Title 按ID修改员工信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-04-04 17:47:43
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcEmployee) ModifyEmployeeByAccountId(ctx context.Context, in system_master.ModifyEmployeeInput) (code int32, message string, err error) {
	// 转换用户数据，用account id换取电话
	code, message, output, err := s.AccountIdCastTel(ctx, in.AccountId)
	if code != 0 {
		return code, message, err
	}

	tel := gconv.String(output)
	// 判断电话号信息是否变更，变更需要验证新电话号是否存在
	if tel != in.Tel {
		// 验证修改电话是否存在
		code, message, err = s.ExistsTel(ctx, in.Tel)
		if err != nil {
			return -1, "", err
		}

		if code == 0 || code >= 2 {
			return 1, "修改电话已存在", nil
		}
	}

	// tel 未更新，主账户不更新电话
	data := g.Map{"status": in.Status}
	if code == 1 {
		data["tel"] = s.prefix + in.Tel
	}
	// 修改主表信息
	_, _, err = utility.DBModifyById(dao.UcAccount.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  data,
		Where: in.AccountId,
	})

	if err != nil {
		return -1, "", err
	}

	_, _, err = utility.DBModifyByWhere(dao.UcEmployee.Ctx(ctx), utility.DBModifyByWhereInput{
		Data: g.Map{
			"name": in.Name,
			"tel":  in.Tel,
		},
		Where: "account_id = ?",
		Args:  in.AccountId,
	})
	if err != nil {
		return -1, "", err
	}

	// 更新电话转换信息
	if tel != in.Tel {
		_, _, _ = s.UpdateAccountIdCastTel(ctx, in.AccountId, in.Tel)
	}

	return 0, "修改成功", nil
}
