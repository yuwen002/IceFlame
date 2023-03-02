package uc_center

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/internal/consts"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"
	"strings"
	"time"
)

var insUcSystemMaster = sUcSystemMaster{
	prefix:      "SA_",
	redisConfig: "uc_center",
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
	prefix      string
	redisConfig string
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
		Config: s.redisConfig,
		Key:    dao.UcAccount.Table() + ":exists_username",
		Value:  username,
	}

	code, message, err = utility.RCExistsSetData(data, func(condition interface{}) (code int32, message string, err error) {
		code, message, _, err = utility.DBGetOneMapByWhere(dao.UcAccount.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "id",
			Where: "username=?",
			Args:  condition,
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

// AccountIdCastTel
//
// @Title account id 换 管理员电话
// @Description account id 换 管理员电话
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-27 16:22:10
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMaster) AccountIdCastTel(ctx context.Context, accountId uint64) (code int32, message string, output interface{}, err error) {
	code, message, output, err = utility.RCCastHashData(utility.RedisCastData{
		Config: s.redisConfig,
		Key:    dao.UcSystemMaster.Table() + ":account_id_cast_tel",
		Field:  gconv.String(accountId),
	}, func(condition interface{}) (code int32, message string, output interface{}, err error) {
		code, message, out, err := utility.DBGetMapById(dao.UcSystemMaster.Ctx(ctx), utility.DBGetByIdInput{
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
// @Description 更新account id转换信息的电话
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 10:15:18
// @receiver s
// @param ctx
// @param accountId
// @param tel
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) UpdateAccountIdCastTel(ctx context.Context, accountId uint64, tel string) (code int32, message string, err error) {
	code, message, err = utility.RCUpdateCastHashData(utility.RedisCastData{
		Config: s.redisConfig,
		Key:    dao.UcSystemMaster.Table() + ":account_id_cast_tel",
		Field:  gconv.String(accountId),
		Data:   tel,
	})

	return code, message, err
}

// GetLoginCountRKey
//
// @Title 根据用户account ID 获取Redis key
// @Description 根据用户account ID 获取Redis key
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 18:18:29
// @receiver s
// @param accountId
// @return string
func (s *sUcSystemMaster) GetLoginCountRKey(accountId string) string {
	return dao.UcSystemMaster.Table() + "verify_login_count(" + time.Now().Format("2006-01-02") + "):account_id-" + accountId
}

// IncLoginCount
//
// @Title 密码输入错误,增加验证次数
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 18:16:29
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) IncLoginCount(ctx context.Context, accountId string) (code int32, message string, err error) {
	key := s.GetLoginCountRKey(accountId)

	redis := utility.InitRedis(s.redisConfig)
	num, err := redis.Incr(ctx, key)
	if err != nil {
		return -1, "", err
	}

	if num >= 3 {
		return 1, "登入次数超出限制", nil
	}

	return 0, "登入错误次数" + gconv.String(num), nil
}

// GetLoginCount
//
// @Title 获取登入错误次数
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 18:16:22
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) GetLoginCount(ctx context.Context, accountId string) (code int32, message string, err error) {
	key := s.GetLoginCountRKey(accountId)
	redis := utility.InitRedis(s.redisConfig)
	res, err := redis.Get(ctx, key)
	if err != nil {
		return -1, "", err
	}
	if res.Int8() >= 3 {
		return 1, "登入次数超出限制", nil
	}

	return 0, "登入错误次数" + res.String(), nil
}

// DelLoginCount
//
// @Title 删除登入错误次数
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-28 21:58:47
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) DelLoginCount(ctx context.Context, accountId string) (code int32, message string, err error) {
	key := s.GetLoginCountRKey(accountId)
	redis := utility.InitRedis(s.redisConfig)
	_, err = redis.Del(ctx, key)
	if err != nil {
		return -1, "", err
	}

	return 0, "删除登入次数成功", nil
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

	id := gconv.String(data["id"])
	// 获取登入次数，查看用户是否登入限制
	code, message, err = s.GetLoginCount(ctx, id)
	if code != 0 {
		return code, message, "", err
	}

	// 验证密码是否正确
	if utility.PasswordVerify(in.Password, gconv.String(data["password_hash"])) == false {
		code, message, err = s.IncLoginCount(ctx, id)
		if code != 0 {
			return code, message, "", err
		}
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

	// 登入成功，删除次数限制
	_, _, _ = s.DelLoginCount(ctx, id)
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

	id := gconv.String(data["id"])
	// 获取登入次数，查看用户是否登入限制
	code, message, err = s.GetLoginCount(ctx, id)
	if code != 0 {
		return code, message, "", err
	}

	// 验证密码是否正确
	if utility.PasswordVerify(in.Password, gconv.String(data["password_hash"])) == false {
		code, message, err = s.IncLoginCount(ctx, id)
		if code != 0 {
			return code, message, "", err
		}
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

	// 登入成功，删除次数限制
	_, _, _ = s.DelLoginCount(ctx, id)
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
func (s *sUcSystemMaster) ListSystemMaster(ctx context.Context, in system_master.ListSystemMasterInput) (code int32, message string, out []system_master.ListSystemMasterOutput, err error) {
	var systemMaster []*system_master.UcSystemMaster
	err = dao.UcSystemMaster.Ctx(ctx).With(system_master.UcSystemMaster{}.UcAccount).Where(
		"supper_master = 0 and account_id != ?",
		ctx.Value("master_id"),
	).Page(in.Page, in.Size).Order("id desc").Scan(&systemMaster)
	if err != nil {
		return -1, "", nil, err
	}

	for _, v := range systemMaster {

		out = append(out, system_master.ListSystemMasterOutput{
			Id:           v.Id,
			AccountId:    v.AccountId,
			Username:     strings.Replace(v.UcAccount.Username, s.prefix, "", -1),
			Tel:          v.Tel,
			Name:         v.Name,
			RealNameType: v.UcAccount.RealNameType,
			Status:       v.UcAccount.Status,
			CreatedAt:    v.UcAccount.CreatedAt,
			UpdatedAt:    v.UcAccount.UpdatedAt,
		})
	}

	return 0, "查询成功", out, nil
}

// ModifySystemMasterByAccountId
//
// @Title 修改管理员用户信息
// @Description 修改管理员用户信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-26 23:57:15
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) ModifySystemMasterByAccountId(ctx context.Context, in system_master.ModifySystemMasterByAccountIdInput) (code int32, message string, err error) {
	// 转换用户数据，用account id换取电话
	code, message, output, err := s.AccountIdCastTel(ctx, in.AccountId)
	if code != 0 {
		return code, message, err
	}

	outputTel := gconv.String(output)
	// 判断电话号信息是否变更，变更需要验证新电话号是否存在
	if outputTel != in.Tel {
		code, message, err = s.ExistsTel(ctx, in.Tel)
		fmt.Println(code)
		if code != 1 {
			return code, message, err
		}
	}

	//  修改用户信息
	err = dao.UcSystemMaster.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新管理员表信息
		code, _, err = utility.DBModifyById(dao.UcSystemMaster.Ctx(ctx), utility.DBModifyByIdInput{
			Data: g.Map{
				"tel":  in.Tel,
				"name": in.Name,
			},
			Where: "account_id=? and supper_master=0",
			Args:  in.AccountId,
		})

		if err != nil {
			return err
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
			return err
		}

		// 更新电话转换信息
		if outputTel != in.Tel {
			_, _, _ = s.UpdateAccountIdCastTel(ctx, in.AccountId, in.Tel)
		}

		return nil
	})

	if err != nil {
		return -1, "", err
	}

	return 0, "修改成功", nil
}

// ResetPasswordByAccountId
//
// @Title 重置管理员密码
// @Description 重置管理员，密码为空重置初始密码为123456
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 14:39:59
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) ResetPasswordByAccountId(ctx context.Context, in system_master.ResetPasswordByAccountIdInput) (code int32, message string, err error) {
	// 没设置新密码，初始化密码
	if in.Password == "" {
		in.Password = "123456"
	}

	// 新密码hash
	in.Password, err = utility.PasswordHash(in.Password)
	if err != nil {
		return -1, "", err
	}

	return utility.DBModifyById(dao.UcAccount.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  g.Map{"password_hash": in.Password},
		Where: in.AccountId,
	})
}

// ModifyStatusByAccountId
//
// @Title 修改管理用户状态
// @Description 修改管理用户状态
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 15:50:46
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMaster) ModifyStatusByAccountId(ctx context.Context, in system_master.ModifyStatusByAccountIdInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.UcAccount.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  g.Map{"status": in.Status},
		Where: in.AccountId,
	})
}
