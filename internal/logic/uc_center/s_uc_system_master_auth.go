package uc_center

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insUcSystemMasterAuth = sUcSystemMasterAuth{}

func UcSystemMasterAuth() *sUcSystemMasterAuth {
	return &insUcSystemMasterAuth
}

func init() {
	service.RegisterUcSystemMasterAuth(UcSystemMasterAuth())
}

type sUcSystemMasterAuth struct {
}

// CreateRole
//
// @Title 添加管理员角色信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 14:44:05
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterAuth) CreateRole(ctx context.Context, in system_master.CreateRoleInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.UcSystemMasterRole.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// ModifyRoleById
//
// @Title 按ID修改管理员角色信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 15:13:06
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterAuth) ModifyRoleById(ctx context.Context, in system_master.ModifyRoleByIdInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.UcSystemMasterRole.Ctx(ctx), utility.DBModifyByIdInput{
		Data: g.Map{
			"name":   in.Name,
			"remark": in.Remark,
		},
		Where: in.Id,
	})
}

// ListRole
//
// @Title 管理员角色信息列表
// @Description 管理员角色信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 15:49:15
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) ListRole(ctx context.Context, in system_master.ListRoleInput) (code int32, message string, output []*system_master.ListRoleOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.UcSystemMasterRole.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "id asc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &output)

	return code, message, output, err
}

// CreateRoleRelation
//
// @Title 绑定管理员与用户角色信息
// @Description 新建绑定管理员与用户角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 17:56:11
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterAuth) CreateRoleRelation(ctx context.Context, in system_master.CreateRoleRelationInput) (code int32, message string, err error) {
	// 查询关联信息是否存在
	code, message, _, err = utility.DBGetOneMapByWhere(dao.UcSystemMasterRoleRelation.Ctx(ctx), utility.DBGetOneByWhereInput{
		Field: "id",
		Where: "account_id=? and role_id=?",
		Args:  g.Slice{in.AccountId, in.RoleId},
	})

	if code != 1 {
		return code, message, err
	}

	return utility.DBInsert(dao.UcSystemMasterRoleRelation.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// ModifyRoleRelationById
//
// @Title 按ID修改管理员与用户角色绑定信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-08 11:09:17
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterAuth) ModifyRoleRelationById(ctx context.Context, in system_master.ModifyRoleRelationByIdInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.UcSystemMasterRoleRelation.Ctx(ctx), utility.DBModifyByIdInput{
		Data: g.Map{
			"account_id": in.AccountId,
			"role_id":    in.RoleId,
		},
		Where: in.Id,
	})
}

// ListRoleRelation
//
// @Title 管理员与用户角色绑定信息列表
// @Description 管理员与用户角色绑定信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-08 17:28:03
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) ListRoleRelation(ctx context.Context, in system_master.ListRoleRelationInput) (code int32, message string, output []*system_master.ListRoleRelationOutput, err error) {
	// 管理员与用户角色绑定信息列表
	code, message, out, err := utility.DBGetAllMapByWhere(dao.UcSystemMasterRoleRelation.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "id desc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	})

	if code != 0 {
		return code, message, nil, err
	}

	// 取出account id数组
	accountIds := utility.ArrayColumn(out, "account_id")
	// 查询管理员信息
	_, _, master, err := utility.DBGetAllMapByWhere(dao.UcSystemMaster.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "account_id, name",
		Where: "account_id in (?)",
		Args:  accountIds,
		Order: "id desc",
	})

	if err != nil {
		return -1, "", nil, err
	}

	// 管理员map信息
	masterNames := utility.MapsFromColumns(master, "account_id", "name")

	// 取出管理员角色id数组
	roleIds := utility.ArrayColumn(out, "role_id")
	// 查询管理员角色信息
	_, _, role, err := utility.DBGetAllMapByWhere(dao.UcSystemMasterRole.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "id, name",
		Where: "id in (?)",
		Args:  roleIds,
		Order: "id desc",
	})

	if err != nil {
		return -1, "", nil, err
	}
	// 管理员角色map信息
	roleNames := utility.MapsFromColumns(role, "id", "name")

	for _, val := range out {
		output = append(output, &system_master.ListRoleRelationOutput{
			Id:        gconv.Uint32(val["id"]),
			AccountId: gconv.Uint64(val["account_id"]),
			Name:      "",
			RoleId:    gconv.Uint16(val["role_id"]),
			RoleName:  "",
			CreatedAt: gtime.Time{},
			UpdatedAt: gtime.Time{},
		})
	}

}
func (s *sUcSystemMasterAuth) DeleteRoleRelation(ctx context.Context, in system_master.DeleteRoleRelationInput) (code int32, message string, err error) {
}
