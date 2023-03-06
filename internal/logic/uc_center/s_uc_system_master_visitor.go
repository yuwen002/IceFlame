package uc_center

import (
	"context"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var insUcSystemMasterVisitor = sUcSystemMasterVisitor{
	VisitCategoryRKey: dao.UcSystemMasterVisitCategory.Table() + ":list_visit_category",
	redisConfig:       "uc_center",
}

func UcSystemMasterVisitorLogs() *sUcSystemMasterVisitor {
	return &insUcSystemMasterVisitor
}

func init() {
	service.RegisterUcSystemMasterVisitor(UcSystemMasterVisitorLogs())
}

// sUcSystemMasterVisitor
// @Description: 用户访问记录结构体
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 15:11:14
type sUcSystemMasterVisitor struct {
	VisitCategoryRKey string
	redisConfig       string
}

// AddVisitCategory
//
// @Title 添加访问类型分类
// @Description 添加访问类型分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 17:09:59
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterVisitor) AddVisitCategory(ctx context.Context, in system_master.AddVisitCategoryInput) (code int32, message string, err error) {
	// 数据写入
	code, message, lastInsertId, err := utility.DBInsertAndGetId(dao.UcSystemMasterVisitCategory.Ctx(ctx), utility.DBInsertInput{Data: in})
	if code != 0 {
		return code, message, err
	}

	// 缓存信息Redis缓存
	code, message, err = utility.RCSetHashId(utility.RedisHashIdData{
		Config: s.redisConfig,
		Key:    s.VisitCategoryRKey,
		Id:     gconv.String(lastInsertId),
		Data:   in.Title,
	})
	if err != nil {
		return code, message, err
	}

	return 0, "数据写入成功", nil
}

// ModifyVisitCategoryById
//
// @Title 修改访问类型分类
// @Description 修改访问类型分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 17:14:18
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterVisitor) ModifyVisitCategoryById(ctx context.Context, in system_master.ModifyVisitCategoryByIdInput) (code int32, message string, err error) {
	// 修改数据
	code, message, err = utility.DBModifyById(dao.UcSystemMasterVisitCategory.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  g.Map{"title": in.Title},
		Where: in.Id,
	})

	if code != 0 {
		return code, message, err
	}

	// 缓存信息Redis缓存
	code, message, err = utility.RCSetHashId(utility.RedisHashIdData{
		Config: s.redisConfig,
		Key:    s.VisitCategoryRKey,
		Id:     gconv.String(in.Id),
		Data:   in.Title,
	})
	if err != nil {
		return code, message, err
	}

	return 0, "修改数据成功", nil
}

// ListVisitCategory
//
// @Title 显示访问类型列表
// @Description 显示访问类型列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 17:31:21
// @receiver s
// @param ctx
func (s *sUcSystemMasterVisitor) ListVisitCategory(ctx context.Context, in system_master.ListVisitCategoryInput) (code int32, message string, output []*system_master.ListVisitCategoryOutput, err error) {
	var out []*system_master.ListVisitCategoryOutput
	code, message, err = utility.DBGetAllStructByWhere(dao.UcSystemMasterVisitCategory.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field:    "id, title",
		Order:    "id asc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &out)

	if code != 0 {
		return code, message, nil, err
	}

	return code, message, out, nil
}

// GetRCacheVisitCategory
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 15:53:56
// @receiver s
// @param ctx
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterVisitor) GetRCacheVisitCategory(ctx context.Context) (code int32, message string, output map[string]interface{}, err error) {
	code, message, output, err = utility.RCGetMapsHashAll(utility.RedisHashIdData{
		Config: s.redisConfig,
		Key:    s.VisitCategoryRKey,
	})
	if err != nil {
		return -1, "", nil, err
	}

	if code == 1 {
		// 查询数据中数据
		code, message, out, err := utility.DBGetAllMapByWhere(dao.UcSystemMasterVisitCategory.Ctx(ctx), utility.DBGetAllByWhereInput{
			Field: "id, title",
			Order: "id asc",
		})
		if code != 0 {
			return code, message, nil, err
		}

		// 缓存信息写入
		for index := range out {
			id := gconv.String(out[index]["id"])
			_, _, err := utility.RCSetHashId(utility.RedisHashIdData{
				Config: s.redisConfig,
				Key:    s.VisitCategoryRKey,
				Id:     id,
				Data:   out[index]["title"],
			})
			if err != nil {
				return 0, "", nil, err
			}
			output[id] = out[index]["title"]
		}

		return 0, "取出数据成功", output, nil
	}

	return 0, "取出数据成功", output, nil
}

// GetRCacheVisitCategoryById
//
// @Title 按ID获取缓存信息
// @Description 按ID获取缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-02 20:39:20
// @receiver s
// @param id
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterVisitor) GetRCacheVisitCategoryById(id string) (code int32, message string, output map[string]interface{}, err error) {
	return utility.RCGetMapHashId(utility.RedisHashIdData{Id: id})
}

// DelRCacheVisitCategory
//
// @Title 删除缓存信息
// @Description  删除缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 11:53:25
// @receiver s
// @param key
// @return code
// @return message
// @return err
func (s *sUcSystemMasterVisitor) DelRCacheVisitCategory() (code int32, message string, err error) {
	return utility.RCDelKey(utility.RedisDelKey{
		Config: s.redisConfig,
		Key:    s.VisitCategoryRKey,
	})
}

// DelRCacheVisitCategoryById
//
// @Title 按ID删除缓存信息
// @Description 按ID删除缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 11:54:15
// @receiver s
// @param id
// @return code
// @return message
// @return err
func (s *sUcSystemMasterVisitor) DelRCacheVisitCategoryById(id string) (code int32, message string, err error) {
	return utility.RCDelHashId(utility.RedisHashIdData{
		Config: s.redisConfig,
		Key:    s.VisitCategoryRKey,
		Id:     id,
	})
}

// AddVisitorLogs
//
// @Title 访问日志信息写入
// @Description 访问日志信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-03 23:18:31
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sUcSystemMasterVisitor) AddVisitorLogs(ctx context.Context, in system_master.AddVisitorLogsInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.UcSystemMasterVisitorLogs.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// ListVisitorLogs
//
// @Title 访问日志信息列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-04 00:01:46
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sUcSystemMasterVisitor) ListVisitorLogs(ctx context.Context, in system_master.ListVisitorLogsInput) (code int32, message string, output []*system_master.ListVisitorLogsOutput, err error) {
	// 判断查询条件
	var condition []string

	// 用户ID
	if in.AccountId > 0 {
		condition = append(condition, "account_id="+gconv.String(in.AccountId))
	}

	// 判断系统平台
	if in.OsCategory > 0 {
		condition = append(condition, "os_category="+gconv.String(in.OsCategory))
	}

	// 判断访问类型
	if in.VisitCategory > 0 {
		condition = append(condition, "visit_category="+gconv.String(in.VisitCategory))
	}

	where := ""
	if len(condition) > 0 {
		where = strings.Join(condition, " and ")
	}

	code, message, err = utility.DBGetAllStructByWhere(dao.UcSystemMasterVisitorLogs.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field:        "*, INET6_NTOA(ip_long) as ip_long",
		Where:        where,
		Order:        "id desc",
		PageType:     1,
		DBPagination: utility.DBPagination{Page: in.Page, Size: in.Size},
	}, &output)

	if code != 0 {
		return code, message, nil, err
	}

	code, message, visitCategory, err := s.GetRCacheVisitCategory(ctx)
	if code != 0 {
		return code, message, nil, err
	}

	for index := range output {
		output[index].OsCategoryName = utility.GetOsCategoryName(output[index].OsCategory)
		output[index].VisitCategoryName = gconv.String(visitCategory[gconv.String(output[index].VisitCategory)])
	}

	return code, message, output, err
}
