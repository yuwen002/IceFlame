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

	return 0, message, nil
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
	_, err = utility.InitRedis(s.redisConfig).HSet(ctx, s.VisitCategoryRKey, g.Map{gconv.String(in.Id): in.Title})
	if err != nil {
		return code, message, err
	}

	return 0, message, nil
}

// ListVisitCategory
//
// @Title 显示访问类型列表
// @Description 显示访问类型列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 17:31:21
// @receiver s
// @param ctx
func (s *sUcSystemMasterVisitorLogs) ListVisitCategory(ctx context.Context) (code int32, message string, output []*system_master.ListVisitCategoryOutput, err error) {
	var out []*system_master.ListVisitCategoryOutput
	code, message, err = utility.DBGetAllStructByWhere(dao.UcSystemMasterVisitCategory.Ctx(ctx), utility.DBGetAllByWhereInput{
		Field: "id, title",
		Order: "id asc",
	}, &out)

	if code != 0 {
		return code, message, nil, err
	}

	return code, message, out, nil
}

func (s *sUcSystemMasterVisitorLogs) GetRCacheVisitCategory(ctx context.Context) (code int32, message string, output []map[string]string, err error) {
	redis := utility.InitRedis(s.redisConfig)
	res, err := redis.HGetAll(ctx, s.VisitCategoryRKey)
	if err != nil {
		return -1, "", nil, err
	}

	if res.IsEmpty() {
		code, message, out, err := s.ListVisitCategory(ctx)
		if code != 0 {
			return code, message, nil, err
		}

		for index := range out {
			value := map[string]string{gconv.String(out[index].Id): out[index].Title}
			output = append(output, value)
			redis.HSet(ctx, s.VisitCategoryRKey, gconv.Map(value))
		}
	}
}

func (s *sUcSystemMasterVisitorLogs) AddVisitorLogs(ctx context.Context) {

}
