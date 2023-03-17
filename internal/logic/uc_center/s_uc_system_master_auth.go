package uc_center

import (
	"context"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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

func (s *sUcSystemMasterAuth) GetRoleById(ctx context.Context, id uint16) (code int32, message string, output system_master.GetRoleByIdOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.UcSystemMasterRole.Ctx(ctx), utility.DBGetByIdInput{
		Field: "id, name, remark",
		Where: id,
	}, &output)

	return code, message, output, err
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

// GetRoleAll
//
// @Title 获取管理员角色信息
// @Description 获取管理员角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 18:17:26
// @receiver s
// @param ctx
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) GetRoleAll(ctx context.Context) (code int32, message string, output []*system_master.GetRoleAllOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.UcSystemMasterRole.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "id,name",
		Order: "id asc",
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

	if err != nil {
		return -1, "", err
	}

	if code == 0 {
		return 1, "添加数据失败，数据已存在", err
	}

	return utility.DBInsert(dao.UcSystemMasterRoleRelation.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetRoleRelationById
//
// @Title 按ID获取管理员与用户角色绑定信息
// @Description 按ID获取管理员与用户角色绑定信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 11:44:45
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) GetRoleRelationById(ctx context.Context, id uint32) (code int32, message string, output system_master.GetRoleRelationByIdOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.UcSystemMasterRoleRelation.Ctx(ctx), utility.DBGetByIdInput{
		Field: "id, account_id, role_id",
		Where: id,
	}, &output)

	return code, message, output, err
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
		// 管理员姓名查找
		masterName := ""
		if v, ok := masterNames[gconv.String(val["account_id"])]; ok {
			masterName = gconv.String(v)
		}

		// 角色名称查找
		roleName := ""
		if v, ok := roleNames[gconv.String(val["role_id"])]; ok {
			roleName = gconv.String(v)
		}

		output = append(output, &system_master.ListRoleRelationOutput{
			Id:        gconv.Uint32(val["id"]),
			AccountId: gconv.Uint64(val["account_id"]),
			Name:      masterName,
			RoleId:    gconv.Uint16(val["role_id"]),
			RoleName:  roleName,
			CreatedAt: gtime.New(val["created_at"]),
			UpdatedAt: gtime.New(val["updated_at"]),
		})
	}

	return 0, "查询数据成功", output, nil
}

// DeleteRoleRelationById
//
// @Title 删除管理员与用户角色绑定信息
// @Description 删除管理员与用户角色绑定信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-09 00:02:55
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterAuth) DeleteRoleRelationById(ctx context.Context, in system_master.DeleteRoleRelationInput) (code int32, message string, err error) {
	return utility.DBDelById(dao.UcSystemMasterRoleRelation.Ctx(ctx), utility.DBDelByIdInput{Where: in.Id})
}

// CreatePermission
//
// @Title 新建权限
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 16:35:23
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterAuth) CreatePermission(ctx context.Context, in system_master.CreateListPermissionInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.UcSystemPermission.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetPermissionById
//
// @Title 按ID获取权限菜单信息
// @Description 按ID获取权限菜单信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 16:19:47
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) GetPermissionById(ctx context.Context, id uint32) (code int32, message string, output system_master.GetPermissionByIdOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.UcSystemPermission.Ctx(ctx), utility.DBGetByIdInput{
		Field: "id, fid, name, module, type, status, sort, remark",
		Where: id,
	}, &output)

	return code, message, output, err
}

// ModifyPermissionById
//
// @Title 按ID修改权限信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 17:14:54
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterAuth) ModifyPermissionById(ctx context.Context, in system_master.ModifyPermissionByIdInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.UcSystemPermission.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

// ListPermission
//
// @Title 权限信息列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-10 15:57:58
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) ListPermission(ctx context.Context, in system_master.ListPermissionInput) (code int32, message string, output []*system_master.ListPermissionOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.UcSystemPermission.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "id asc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &output)

	return code, message, output, err
}

// ModifyPermissionRelation
//
// @Title 角色权限信息关联
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-13 16:51:56
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterAuth) ModifyPermissionRelation(ctx context.Context, in system_master.ModifyPermissionRelationInput) (code int32, message string, err error) {
	// 查询已有权限列表
	code, message, out, err := utility.DBGetAllMapByWhere(dao.UcSystemPermissionRelation.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "permission_id",
		Where: "role_id=?",
		Args:  in.RoleId,
		Order: "id asc",
	})

	if err != nil {
		return -1, "", err
	}

	funcData := func(data interface{}) []map[string]interface{} {
		d := utility.InterfaceToSlice(data)
		insertData := make([]map[string]interface{}, len(d))
		for index := range d {
			insertData[index] = map[string]interface{}{
				"permission_id": d[index],
				"role_id":       in.RoleId,
			}
		}

		return insertData
	}
	// 查询数据为空
	if code == 1 {
		// 拼接输入数据写入信息
		insertData := funcData(in.PermissionIds)
		code, message, err := utility.DBInsert(dao.UcSystemPermissionRelation.Ctx(ctx), utility.DBInsertInput{Data: insertData})
		if code != 0 {
			return code, message, err
		}
	}

	// 提取已分配权限ID
	oldPermissionIds := utility.ArrayColumn(utility.MapsStrStr(out), "permission_id")
	// 分割提交权限数据
	newPermissionIds := utility.ArrayCast[uint16](in.PermissionIds, func(v uint16) string {
		return gconv.String(v)
	})

	// 需要删除的数据
	deletePermissionIds := utility.ArrayDiff[string](oldPermissionIds, newPermissionIds)
	if len(deletePermissionIds) != 0 {
		code, message, err = utility.DBDelByWhere(dao.UcSystemPermissionRelation.Ctx(ctx), utility.DBDelByWhereInput{
			Where: "in (?) and role_id=?",
			Args:  g.Slice{deletePermissionIds, in.RoleId},
		})
		if code != 0 {
			return code, message, err
		}
	}

	// 需要添加的数据
	insertPermissionIds := utility.ArrayDiff[string](newPermissionIds, oldPermissionIds)
	if len(insertPermissionIds) != 0 {
		return utility.DBInsert(dao.UcSystemPermissionRelation.Ctx(ctx), utility.DBInsertInput{Data: funcData(insertPermissionIds)})
	}

	return 1, "数据信息未变动", nil
}

// GetPermissionAll
//
// @Title 所有开启权限列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 14:36:19
// @receiver s
// @param ctx
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) GetPermissionAll(ctx context.Context) (code int32, message string, output []*system_master.GetPermissionAllOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.UcSystemPermission.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "id,fid,name,module,uri",
		Where: "status=0 and type in (1,2)",
		Order: "fid asc, sort desc",
	}, &output)

	if code != 0 {
		return code, message, nil, err
	}

	return code, message, output, nil
}

// GetPermissionByRoleId
//
// @Title 获取权限信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 17:04:36
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) GetPermissionByRoleId(ctx context.Context, roleId uint16) (code int32, message string, output []*system_master.GetPermissionAllOutput, err error) {
	// 查询所有可用权限列表
	code, message, permissionAll, err := s.GetPermissionAll(ctx)
	if code != 0 {
		return code, message, nil, err
	}

	// 查询被关联权限列表
	code, message, permissionRelation, err := utility.DBGetAllMapByWhere(dao.UcSystemPermissionRelation.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "permission_id",
		Where: "role_id=?",
		Args:  roleId,
		Order: "id asc",
	})
	if code != 0 {
		return code, message, nil, err
	}

	// 已有权限关联
	permissionIds := utility.ArrayColumnCast(permissionRelation, "permission_id", func(t interface{}) uint32 {
		return gconv.Uint32(t)
	})
	permissionIdsFunc := utility.ArraySetMapKey(permissionIds)

	// 遍历成树形
	m := make(map[uint32]*system_master.GetPermissionAllOutput)
	tree := make([]*system_master.GetPermissionAllOutput, 0)
	for i, item := range permissionAll {
		// 初始化被选中权限
		if permissionIdsFunc(item.Id) {
			permissionAll[i].Checked = 1
		}

		m[item.Id] = permissionAll[i]
		if item.Fid == 0 {
			tree = append(tree, permissionAll[i])
		}
	}
	for _, item := range permissionAll {
		if item.Fid != 0 {
			parent := m[item.Fid]
			parent.Children = append(parent.Children, m[item.Id])
		}
	}

	return 0, "查询成功", tree, nil
}

// GetMenuAll
//
// @Title 后台导航列表
// @Description 后台导航列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-17 15:10:55
// @receiver s
// @param ctx
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) GetMenuAll(ctx context.Context) (code int32, message string, output []*system_master.GetMenuAllOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.UcSystemPermission.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "id, fid, name, uri",
		Where: "status=0 and type=1",
		Order: "fid asc, sort desc",
	}, &output)

	return code, message, output, err
}

// GetMasterMenu
//
// @Title 获取个管理员Menu信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-17 18:14:32
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) GetMasterMenu(ctx context.Context, accountId uint64) (code int32, message string, output []*system_master.GetMenuAllOutput, err error) {
	// 查询所有可以用菜单
	code, message, menu, err := s.GetMenuAll(ctx)
	if code != 0 {
		return code, message, output, err
	}

	// 查询已分配权限
	code, message, permission, err := s.GetRolePermissionByAccountId(ctx, accountId)
	if code != 0 {
		return code, message, output, err
	}

	// 权限返回信息菜单列表
	permissionFunc := utility.ArraySetMapKey[uint32](permission)
	for index := range menu {
		if menu[index].Fid == 0 || permissionFunc(menu[index].Id) {
			output = append(output, menu[index])
		}
	}

	return 0, "查询成功", output, nil
}

// GetRolePermissionByAccountId
//
// @Title 获取用户权限
// @Description 按用户ID获取用户权限
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-17 16:07:08
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterAuth) GetRolePermissionByAccountId(ctx context.Context, accountId uint64) (code int32, message string, output []uint32, err error) {
	// 查询管理员关联角色
	code, _, out, err := utility.DBGetAllMapByWhere(dao.UcSystemMasterRoleRelation.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "role_id",
		Where: "account_id = ?",
		Args:  accountId,
	})

	if err != nil {
		return code, "", nil, err
	}

	if code != 0 {
		return code, "角色信息错误", nil, nil
	}

	// 角色信息数组
	roleIds := utility.ArrayColumn(out, "role_id")
	// 查询角色关联权限信息
	code, _, out, err = utility.DBGetAllMapByWhere(dao.UcSystemPermissionRelation.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "permission_id",
		Where: "role_id in (?)",
		Args:  roleIds,
	})
	if err != nil {
		return code, "", nil, err
	}

	if code != 0 {
		return code, "权限信息错误", nil, nil
	}

	// 角色关联权限IDS
	permissionIds := utility.ArrayColumnCast(out, "permission_id", func(v interface{}) uint32 {
		return gconv.Uint32(v)
	})

	return code, "查询成功", permissionIds, nil
}
