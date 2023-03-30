package uc_center

import (
	"context"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insUcPartner = sUcPartner{}

func UcPartner() *sUcPartner {
	return &insUcPartner
}

func init() {
	service.RegisterUcPartner(UcPartner())
}

type sUcPartner struct {
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
