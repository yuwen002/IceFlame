package uc_center

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"
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
	fullPath := "0"

	// 如果不是顶级ID，验证传入父ID是否存在
	if in.Fid != 0 {
		code, message, out, err := utility.DBGetOneMapByWhere(dao.UcPartner.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "account_id",
			Where: "account_id = ?",
			Args:  in.Fid,
		})
		if code != 0 {
			return code, message, err
		}

		path = gconv.String(out["path"])
		fullPath = gconv.String(out["full_path"])
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
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

		code, message, err = utility.DBInsert(dao.UcPartner.Ctx(ctx), utility.DBInsertInput{
			Data: g.Map{
				"account_id": id,
				"fid":        in.Fid,
				"level_id":   in.LevelId,
				"path":
			},
		})
	})

}
