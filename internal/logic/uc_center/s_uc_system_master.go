package uc_center

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/internal/consts"
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

// ExistsTel
//
// @Title 验证用户手机号是否存在
// @Description 验证用户电话号是否存在
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 16:41:58
// @receiver s
// @param ctx
// @param tel
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) ExistsTel(ctx context.Context, tel string) (code int32, message string, err error) {
	tel = s.prefix + tel
	data := utility.RedisExistsData{
		Config: "uc_center",
		Key:    dao.UcAccount.Table() + ":exists_tel",
		Value:  tel,
	}

	code, message, err = utility.RCExistsSetData(data, func(condition interface{}) (code int32, message string, err error) {
		code, message, _, err = utility.DBGetOneMapByWhere(dao.UcAccount.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "id",
			Where: "tel=?",
			Args:  tel,
		})

		return code, message, err
	})

	return code, message, err
}

// Register
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
func (s *sUcSystemMaster) Register(ctx context.Context, in system_master.RegisterInput) (code int32, message string, err error) {
	// 验证用户名是否存在
	code, message, err = s.ExistsUsername(ctx, in.Tel)
	if err != nil {
		return code, message, err
	}
	if code == 0 {
		return 1, "用户名已存在", err
	}

	// 验证用户手机号是否存在
	code, message, err = s.ExistsTel(ctx, in.Tel)
	if err != nil {
		return code, message, err
	}
	if code == 0 {
		return 1, "用户名手机号已存在", err
	}

	// 密码hash
	in.Password, err = utility.PasswordHash(in.Password)
	if err != nil {
		return code, message, err
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

// LoginTelPassword
//
// @Title 管理员用户电话，密码登入
// @Description 管理员用户电话，密码登入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 11:08:30
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return token
// @return err
func (s *sUcSystemMaster) LoginTelPassword(ctx context.Context, in system_master.LoginTelPasswordInput) (code int32, message string, token string, err error) {
	// 查询登入信息
	code, message, data, err := utility.DBGetOneMapByWhere(dao.UcAccount.Ctx(ctx), utility.DBGetOneByWhereInput{
		Where: "tel = ? and status = 0",
		Args:  s.prefix + in.Tel,
	})

	if err != nil {
		return -1, "", "", err
	}

	if code == 1 {
		return 1, "登入信息不存在", "", nil
	}

	// 验证密码是否正确
	if utility.PasswordVerify(in.Password, gconv.String(data["password_hash"])) == false {
		return 1, "用户名密码错误", "", nil
	}

	// 生成token
	token, err = utility.CreateToken(utility.CustomClaimsInput{
		Id: gconv.Uint64(data["id"]),
		UserInfo: g.Map{
			"username": data["username"],
		},
		Expire:  60 * 60 * 24,
		Issuer:  "系统管理员:" + gconv.String(data["name"]),
		Subject: "后台管理",
	}, consts.MasterSecretKey)
	if err != nil {
		return -1, "", "", err
	}

	return 0, "登入成功", token, nil
}

// LoginUsernamePassword
//
// @Title 管理员用户名密码登入
// @Description 管理员用户名密码登入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 16:23:48
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return token
// @return err
func (s *sUcSystemMaster) LoginUsernamePassword(ctx context.Context, in system_master.LoginUsernamePasswordInput) (code int32, message string, token string, err error) {
	// 查询登入信息
	code, message, data, err := utility.DBGetOneMapByWhere(dao.UcAccount.Ctx(ctx), utility.DBGetOneByWhereInput{
		Where: "username = ? and status = 0",
		Args:  s.prefix + in.Username,
	})

	if err != nil {
		return -1, "", "", err
	}

	if code == 1 {
		return 1, "登入信息不存在", "", nil
	}

	// 验证密码是否正确
	if utility.PasswordVerify(in.Password, gconv.String(data["password_hash"])) == false {
		return 1, "用户名密码错误", "", nil
	}

	// 生成token
	token, err = utility.CreateToken(utility.CustomClaimsInput{
		Id: gconv.Uint64(data["id"]),
		UserInfo: g.Map{
			"username": data["username"],
		},
		Expire:  60 * 60 * 24,
		Issuer:  "系统管理员:" + gconv.String(data["name"]),
		Subject: "后台管理",
	}, consts.MasterSecretKey)
	if err != nil {
		return -1, "", "", err
	}

	return 0, "登入成功", token, nil
}

// CreateSystemMaster
//
// @Title 后台管理员新建用户
// @Description 后台管理员新建用户，后台管理新建用户，可自定义用户名
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 16:58:13
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) CreateSystemMaster(ctx context.Context, in system_master.CreateSystemMasterInput) (code int32, message string, err error) {
	// 验证用户名是否存在
	code, message, err = s.ExistsUsername(ctx, in.Username)
	if err != nil {
		return code, message, err
	}
	if code == 0 {
		return 1, "用户名已存在", err
	}

	// 验证用户手机号是否存在
	code, message, err = s.ExistsTel(ctx, in.Tel)
	if err != nil {
		return code, message, err
	}
	if code == 0 {
		return 1, "用户名手机号已存在", err
	}

	// 无密码，默认密码123456
	if in.Password == "" {
		in.Password, err = utility.PasswordHash("123456")
		if err != nil {
			return code, message, err
		}
	} else {
		in.Password, err = utility.PasswordHash(in.Password)
		if err != nil {
			return code, message, err
		}
	}

	err = dao.UcAccount.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 写入主表数据uc_account
		code, message, lastInsertId, err := utility.DBInsertAndGetId(dao.UcAccount.Ctx(ctx), utility.DBInsertInput{Data: g.Map{
			"username":      s.prefix + in.Username,
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

// ModifyPasswordSelfById
//
// @Title 修改管理员密码
// @Description 修改管理员密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-24 13:35:22
// @receiver s
// @param ctx
// @param password
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) ModifyPasswordSelfById(ctx context.Context, in system_master.ModifyPasswordInput) (code int32, message string, err error) {
	// 查询旧密码信息
	id := ctx.Value("master_id")
	code, message, out, err := utility.DBGetMapById(dao.UcAccount.Ctx(ctx), utility.DBGetByIdInput{
		Field: "password_hash",
		Where: id,
	})

	if code != 0 {
		return code, message, err
	}

	// 验证旧密码
	if !utility.PasswordVerify(in.OldPassword, gconv.String(out["password_hash"])) {
		return 1, "旧密码不正确", nil
	}

	// 新密码hash
	in.NewPassword, err = utility.PasswordHash(in.NewPassword)
	if err != nil {
		return -1, "", err
	}

	// 更新密码
	code, message, err = utility.DBModifyById(dao.UcAccount.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  g.Map{"password_hash": in.NewPassword},
		Where: id,
	})

	return code, message, err
}

// ListSystemMaster
//
// @Title 管理员信息列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-24 23:47:38
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return out
// @return err
func (s *sUcSystemMaster) ListSystemMaster(ctx context.Context, in system_master.ListSystemMasterInput) (code int32, message string, out []map[string]interface{}, err error) {
	var systemMaster []*system_master.UcSystemMaster
	err = dao.UcSystemMaster.Ctx(ctx).With(system_master.UcSystemMaster{}).Where("supper_master=0").Page(in.Page, in.Size).Scan(&systemMaster)
	if err != nil {
		return -1, "", nil, err
	}

	return 0, "查询成功", gconv.Maps(systemMaster), nil
}
