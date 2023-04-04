package uc_center

import (
	"context"
	"errors"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var insUcPartner = sUcPartner{
	redisConfig: "uc_center",
}

func UcPartner() *sUcPartner {
	return &insUcPartner
}

func init() {
	service.RegisterUcPartner(UcPartner())
}

type sUcPartner struct {
	redisConfig string
}

// ExistsAccountId
//
// @Title 验证合伙人是否存在
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-31 11:22:56
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return err
func (s *sUcPartner) ExistsAccountId(ctx context.Context, accountId uint64) (code int32, message string, err error) {
	data := utility.RedisExistsData{
		Config: s.redisConfig,
		Key:    dao.UcPartner.Table() + ":exists_account_id",
		Value:  accountId,
	}

	code, message, err = utility.RCExistsSetData(data, func(condition interface{}) (code int32, message string, err error) {
		code, message, _, err = utility.DBGetOneMapByWhere(dao.UcPartner.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "account_id",
			Where: "account_id = ?",
			Args:  accountId,
		})

		return code, message, err
	})

	return code, message, err
}

// CreatePartnerLevel
//
// @Title 新建合伙人等级
// @Description 新建合伙人等级
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 15:13:39
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcPartner) CreatePartnerLevel(ctx context.Context, in system_master.CreatePartnerLevelInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.UcPartnerLevel.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetPartnerLevelById
//
// @Title 按ID获取合伙人等级信息
// @Description 按ID获取合伙人等级信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 15:53:00
// @receiver s
func (s *sUcPartner) GetPartnerLevelById(ctx context.Context, id uint16) (code int32, message string, output *system_master.PartnerLevelOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.UcPartnerLevel.Ctx(ctx), utility.DBGetByIdInput{Where: id}, &output)
	return code, message, output, err
}

// ModifyPartnerLevelById
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 16:34:08
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcPartner) ModifyPartnerLevelById(ctx context.Context, in system_master.ModifyPartnerLevelInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.UcPartnerLevel.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

// ListPartnerLevel
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 17:05:05
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sUcPartner) ListPartnerLevel(ctx context.Context, in system_master.ListPartnerLevelInput) (code int32, message string, output []*system_master.PartnerLevelOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.UcPartnerLevel.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "sort desc, id asc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &output)

	return code, message, output, err
}

// GetMapPartnerLevelAll
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-04-04 16:23:48
// @receiver s
// @param ctx
// @return code
// @return message
// @return out
// @return err
func (s *sUcPartner) GetMapPartnerLevelAll(ctx context.Context) (code int32, message string, out map[string]string, err error) {
	code, message, level, err := utility.DBGetAllMapByWhere(dao.UcPartnerLevel.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "id, name",
	})

	if code != 0 {
		return code, message, nil, err
	}

	out = utility.MapStrStr(utility.MapsFromColumns(level, "id", "name"))

	return 0, "查询成功", out, nil
}

// CreatePartner
//
// @Title 新建合伙人
// @Description 新建合伙人
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-31 11:33:18
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcPartner) CreatePartner(ctx context.Context, in system_master.CreatePartnerInput) (code int32, message string, err error) {
	// 路径初始化
	path := "0"
	// 深度初始化
	var depth uint8 = 1
	//
	var team2ndId string = ""
	// 如果不是顶级ID，验证传入父ID是否存在
	if in.Fid != 0 {
		code, _, out, err := utility.DBGetOneMapByWhere(dao.UcPartner.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "account_id",
			Where: "account_id = ?",
			Args:  in.Fid,
		})
		if err != nil {
			return -1, "", err
		}

		if code == 1 {
			return code, "合伙人上级信息错误", err
		}

		// 层级路径
		path = gconv.String(out["path"]) + "," + gconv.String(out["id"])
		// 层级深度
		depth = gconv.Uint8(out["depth"]) + 1

		if gconv.Uint8(out["depth"]) > 2 {
			arrPath := strings.Split(gconv.String(out["path"]), ",")
			team2ndId = arrPath[len(arrPath)-2]
		}
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 注册员工信息
		code, message, id, err := service.UcEmployee().CreateEmployee(ctx, system_master.CreateEmployeeInput{
			Password: in.Password,
			Name:     in.Name,
			Tel:      in.Tel,
		})

		if err != nil {
			return err
		}

		if code != 0 {
			err = errors.New(message)
			return err
		}

		// 员工身份绑定合伙人
		code, message, err = service.UcEmployee().CreateEmployeeRoleRelation(ctx, system_master.CreateEmployeeRoleRelationInput{
			AccountId: gconv.Uint64(id),
			RoleId:    in.RoleId,
		})

		if err != nil {
			return err
		}

		if code != 0 {
			err = errors.New(message)
			return err
		}

		// 新建员工钱包
		code, message, err = service.UcEmployeeEwallet().CreateEmployeeEwallet(ctx, gconv.Uint64(id))
		if err != nil {
			return err
		}

		if code != 0 {
			err = errors.New(message)
			return err
		}

		// 合伙人关系写入
		code, message, err = utility.DBInsert(dao.UcPartner.Ctx(ctx), utility.DBInsertInput{
			Data: g.Map{
				"account_id":     id,
				"fid":            in.Fid,
				"level_id":       in.LevelId,
				"depth":          depth,
				"path":           path,
				"full_path":      path + "," + gconv.String(id),
				"promotion_type": in.PromotionType,
			},
		})

		if err != nil {
			return err
		}

		if code != 0 {
			err = errors.New(message)
			return err
		}

		if in.Fid != 0 {
			// 更新上级团队人数
			code, message, err = utility.DBModifyIncByWhere(dao.UcPartner.Ctx(ctx), utility.DBModifyIncByWhereInput{
				Field:  "team_1st",
				Where:  "account_id = ?",
				Args:   in.Fid,
				Amount: 1,
			})

			if err != nil {
				return err
			}

			if code != 0 {
				err = errors.New(message)
				return err
			}

			if team2ndId != "" {
				// 更新上上级团队人数
				code, message, err = utility.DBModifyIncByWhere(dao.UcPartner.Ctx(ctx), utility.DBModifyIncByWhereInput{
					Field:  "team_1st",
					Where:  "account_id = ?",
					Args:   team2ndId,
					Amount: 1,
				})

				if err != nil {
					return err
				}

				if code != 0 {
					err = errors.New(message)
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return -1, "", err
	}

	return 0, "添加合伙人成功", nil
}

// GetPartnerByAccountId
//
// @Title 按account id获取合伙人信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-04-03 17:37:02
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return out
// @return err
func (s *sUcPartner) GetPartnerByAccountId(ctx context.Context, accountId uint64) (code int32, message string, out *system_master.PartnerOutput, err error) {
	// 查询合伙人信息
	code, message, partner, err := utility.DBGetOneMapByWhere(dao.UcPartner.Ctx(ctx), utility.DBGetOneByWhereInput{
		Field: "fid, level_id, team_1st, team_2nd, created_at, updated_at",
		Where: "account_id = ?",
		Args:  accountId,
	})

	if err != nil {
		return -1, "", nil, err
	}

	if code == 1 {
		return code, message, nil, nil
	}

	out = new(system_master.PartnerOutput)
	out.AccountId = accountId
	out.Team1st = gconv.Uint32(partner["team_1st"])
	out.Team2nd = gconv.Uint32(partner["team_2nd"])
	out.Fid = gconv.Uint64(partner["fid"])
	out.LevelId = gconv.Uint16(partner["level_id"])
	out.PromotionType = gconv.Uint8(partner["promotion_type"])
	out.CreatedAt = gconv.GTime(partner["created_at"])
	out.UpdatedAt = gconv.GTime(partner["updated_at"])

	// 所有合伙人级别
	code, message, level, err := s.GetMapPartnerLevelAll(ctx)
	if code != 0 {
		return code, message, nil, err
	}
	out.LevelName = level[gconv.String(partner["level_id"])]

	// 查询上级信息
	if out.Fid != 0 {
		code, message, fPartner, err := utility.DBGetOneMapByWhere(dao.UcEmployee.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "name, tel",
			Where: "account_id = ?",
			Args:  out.Fid,
		})

		if err != nil {
			return -1, "", nil, err
		}

		if code == 1 {
			return code, message, nil, nil
		}

		out.FName = gconv.String(fPartner["name"])
		out.FTel = gconv.String(fPartner["tel"])
	}

	// 合伙人员工表基本信息
	code, message, employee, err := utility.DBGetOneMapByWhere(dao.UcEmployee.Ctx(ctx), utility.DBGetOneByWhereInput{
		Field: "name, tel, invite_code, real_name_type",
		Where: "account_id = ?",
		Args:  accountId,
	})

	if err != nil {
		return -1, "", nil, err
	}

	if code == 1 {
		return code, message, nil, nil
	}

	out.Name = gconv.String(employee["name"])
	out.Tel = gconv.String(employee["tel"])
	out.InviteCode = gconv.String(employee["invite_code"])

	return 0, "查询成功", out, nil
}

// ModifyPartnerByAccountId
//
// @Title 按account id修改合伙人信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-04-04 17:33:51
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcPartner) ModifyPartnerByAccountId(ctx context.Context, in system_master.ModifyPartnerInput) (code int32, message string, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, _, err = utility.DBModifyByWhere(dao.UcPartner.Ctx(ctx), utility.DBModifyByWhereInput{
			Data: g.Map{
				"level_id":       in.LevelId,
				"promotion_type": in.PromotionType,
			},
			Where: "account_id = ?",
			Args:  in.AccountId,
		})

		if err != nil {
			return err
		}

		_, _, err = service.UcEmployee().ModifyEmployeeByAccountId(ctx, system_master.ModifyEmployeeInput{
			AccountId: in.AccountId,
			Name:      in.Name,
			Tel:       in.Tel,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return -1, "", nil
	}

	return 0, "修改合伙人成功", nil
}
