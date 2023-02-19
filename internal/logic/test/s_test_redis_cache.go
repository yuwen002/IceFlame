package test

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/internal/dao"
	"ice_flame/internal/model"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insRedisCacheTest = sRedisCacheTest{}

func NewRedisCacheTest() *sRedisCacheTest {
	return &insRedisCacheTest
}

func init() {
	service.RegisterRedisCacheTest(NewRedisCacheTest())
}

// sRedisCacheTest
// @Description: Redis缓存测试
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-18 11:18:40
type sRedisCacheTest struct {
}

// TestExistsSetData
//
// @Title 测试Redis集合数据是否存在
// @Description 测试Redis集合数据是否存在
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-18 10:27:31
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sRedisCacheTest) TestExistsSetData(ctx context.Context, in model.TestExistsSetDataInput) (code int32, message string, err error) {
	return utility.ExistsSetData(utility.RedisExistsData{
		Key:   "db_test:exists_id",
		Value: in.Id,
	}, func(condition interface{}) (code int32, message string, err error) {
		code, message, _, err = utility.DBGetMapById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{
			Field: "id",
			Where: condition,
		})

		return code, message, err
	})
}

// TestCastHashData
//
// @Title 测试Redis 类型转换
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-19 01:17:27
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sRedisCacheTest) TestCastHashData(ctx context.Context, in model.TestCastHashDataInput) (code int32, message string, output interface{}, err error) {
	code, message, output, err = utility.CastHashData(utility.RedisCastData{
		Config: "test",
		Key:    "db_test:cast_id",
		Field:  gconv.String(in.Id),
	}, func(condition interface{}) (code int32, message string, output interface{}, err error) {
		code, message, output, err = utility.DBGetMapById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{
			Field: "id",
			Where: condition,
		})

		if code != 0 {
			return code, message, output, err
		}
		unionId := output.(map[string]interface{})
		return code, message, unionId["union_id"], err
	})

}
