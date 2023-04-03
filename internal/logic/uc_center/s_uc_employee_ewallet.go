package uc_center

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
	"ice_flame/utility"

	"github.com/gogf/gf/v2/frame/g"
)

var insUcEmployeeEwallet = sUcEmployeeEwallet{
	prefix:      "EMP_",
	redisConfig: "uc_center",
}

func UcEmployeeEwallet() *sUcEmployeeEwallet {
	return &insUcEmployeeEwallet
}

func init() {
	service.RegisterUcEmployeeEwallet(UcEmployeeEwallet())
}

type sUcEmployeeEwallet struct {
	prefix      string
	redisConfig string
}

// ExistsEwallet
//
// @Title 验证钱包是否存在
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 20:12:30
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return err
func (s *sUcEmployeeEwallet) ExistsEwallet(ctx context.Context, accountId uint64) (code int32, message string, err error) {
	data := utility.RedisExistsData{
		Config: s.redisConfig,
		Key:    dao.UcEmployeeEwallet.Table() + ":exists_account_id",
		Value:  accountId,
	}

	code, message, err = utility.RCExistsSetData(data, func(condition interface{}) (code int32, message string, err error) {
		code, message, _, err = utility.DBGetOneMapByWhere(dao.UcEmployeeEwallet.Ctx(ctx), utility.DBGetOneByWhereInput{
			Field: "id",
			Where: "account_id = ?",
			Args:  condition,
		})

		return code, message, err
	})

	return code, message, err
}

// SetHash
//
// @Title 对钱包进行hash
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 20:51:04
// @receiver s
// @param in
// @return code
// @return message
// @return hash
// @return err
func (s *sUcEmployeeEwallet) SetHash(in system_master.EmployeeEwalletHash) (code int32, message string, hash string, err error) {
	data, err := json.Marshal(in)
	if err != nil {
		return -1, "", "", err
	}
	sum := sha256.Sum256(data)
	hash = hex.EncodeToString(sum[:])
	return 0, "hash成功", hash, nil
}

// CreateEmployeeEwallet
//
// @Title 新建员工钱包
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 20:08:37
// @receiver s
// @param ctx
// @param accountId
// @return code
// @return message
// @return err
func (s *sUcEmployeeEwallet) CreateEmployeeEwallet(ctx context.Context, accountId uint64) (code int32, message string, err error) {
	code, message, err = s.ExistsEwallet(ctx, accountId)
	if code != 1 {
		return code, message, err
	}
	in := system_master.EmployeeEwalletHash{
		AccountId:  accountId,
		Money:      0,
		Balance:    0,
		TotalMoney: 0,
	}
	_, _, hash, err := s.SetHash(in)
	if err != nil {
		return -1, "对钱包hash错误", err
	}

	return utility.DBInsert(dao.UcEmployeeEwallet.Ctx(ctx), utility.DBInsertInput{Data: g.Map{
		"account_id":   accountId,
		"money":        0,
		"balance":      0,
		"total_money":  0,
		"hash_ewallet": hash,
	}})
}
