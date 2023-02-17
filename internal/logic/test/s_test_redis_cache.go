package test

import (
	"context"
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

type sRedisCacheTest struct {
}

func (s *sRedisCacheTest) TestExistsSetData(ctx context.Context, in model.TestExistsSetDataInput) (code int32, message string, err error) {
	var f = func(condition interface{}) (code int32, message string, err error) {
		code, message, _, err = utility.DBGetMapById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{
			Field: "id",
			Where: condition,
		})

		return code, message, err
	}
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
