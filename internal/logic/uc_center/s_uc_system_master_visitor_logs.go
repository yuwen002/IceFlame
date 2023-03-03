package uc_center

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insUcSystemMasterVisitorLogs = sUcSystemMasterVisitorLogs{
	VisitCategoryRKey: dao.UcSystemMasterVisitCategory.Table() + ":list_visit_category",
	redisConfig:       "uc_center",
}

func UcSystemMasterVisitorLogs() *sUcSystemMasterVisitorLogs {
	return &insUcSystemMasterVisitorLogs
}

func init() {
	service.RegisterUcSystemMasterVisitorLogs(UcSystemMasterVisitorLogs())
}

// sUcSystemMasterVisitorLogs
// @Description: 用户访问记录结构体
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 15:11:14
type sUcSystemMasterVisitorLogs struct {
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
func (s *sUcSystemMasterVisitorLogs) AddVisitCategory(ctx context.Context, in system_master.AddVisitCategoryInput) (code int32, message string, err error) {
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
func (s *sUcSystemMasterVisitorLogs) ModifyVisitCategoryById(ctx context.Context, in system_master.ModifyVisitCategoryByIdInput) (code int32, message string, err error) {
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
func (s *sUcSystemMasterVisitorLogs) ListVisitCategory(ctx context.Context, in system_master.ListVisitCategoryInput) (code int32, message string, output []*system_master.ListVisitCategoryOutput, err error) {
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
func (s *sUcSystemMasterVisitorLogs) GetRCacheVisitCategory(ctx context.Context) (code int32, message string, output []map[string]interface{}, err error) {
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
			_, _, err := utility.RCSetHashId(utility.RedisHashIdData{
				Config: s.redisConfig,
				Key:    s.VisitCategoryRKey,
				Id:     gconv.String(out[index]["id"]),
				Data:   out[index]["title"],
			})
			if err != nil {
				return 0, "", nil, err
			}
		}

		return 0, "取出数据成功", out, nil
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
func (s *sUcSystemMasterVisitorLogs) GetRCacheVisitCategoryById(id string) (code int32, message string, output map[string]interface{}, err error) {
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
func (s *sUcSystemMasterVisitorLogs) DelRCacheVisitCategory(key string) (code int32, message string, err error) {
	return utility.RCDelKey(utility.RedisDelKey{
		Config: s.redisConfig,
		Key:    key,
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
func (s *sUcSystemMasterVisitorLogs) DelRCacheVisitCategoryById(key string) (code int32, message string, err error) {
	return utility.RCDelHashId(utility.RedisHashIdData{Key: key})
}

func (s *sUcSystemMasterVisitorLogs) AddVisitorLogs(ctx context.Context) {

}
