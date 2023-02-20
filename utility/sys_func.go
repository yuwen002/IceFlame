package utility

import (
	"bytes"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash
//
// @Title 创建密码的散列
// @Description 密码hash
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-20 15:26:24
// @param password
// @return string
// @return error
func PasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

// PasswordVerify
//
// @Title 验证密码是否和散列值匹配
// @Description 验证密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-20 15:26:50
// @param password
// @param hash
// @return bool
func PasswordVerify(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}

	return true
}

// ConvertUnitMul
//
// @Title 单位转换 乘法
// @Description  单位转换 乘法
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 23:17:55
// @param num
// @param unit
// @return int64
func ConvertUnitMul(num float64, unit int64) int64 {
	n1 := decimal.NewFromFloat(num)
	n2 := decimal.NewFromInt(unit)
	result := n1.Mul(n2).IntPart()

	return result
}

// ConvertUnitDiv
//
// @Title 单位转换 除法
// @Description  单位转换 除法
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 23:21:36
// @param num
// @param unit
// @return float64
func ConvertUnitDiv(num float64, unit int64) float64 {
	n1 := decimal.NewFromFloat(num)
	n2 := decimal.NewFromInt(unit)
	result, _ := n1.Div(n2).Float64()

	return result
}

// ConvertFenToYuan
//
// @Title 分变元
// @Description  分变元
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 23:21:59
// @param money
// @return float64
func ConvertFenToYuan(money int64) float64 {
	n1 := decimal.NewFromInt(money)
	n2 := decimal.NewFromInt(100)
	result, _ := n1.Div(n2).Float64()

	return result
}

// ConvertYuanToFen
//
// @Title 元变分
// @Description  元变分
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 23:22:18
// @param money
// @return int64
func ConvertYuanToFen(money float64) int64 {
	n1 := decimal.NewFromFloat(money)
	n2 := decimal.NewFromInt(100)
	result := n1.Mul(n2).IntPart()

	return result
}

// ConvertLiToYuan
//
// @Title 里变元
// @Description 里变元
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 23:22:48
// @param money
// @return float64
func ConvertLiToYuan(money int64) float64 {
	n1 := decimal.NewFromInt(money)
	n2 := decimal.NewFromInt(100)
	result, _ := n1.Div(n2).Float64()

	return result
}

// ConvertYuanToLi
//
// @Title 元变里
// @Description 元变里
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 23:22:59
// @param money
// @return int64
func ConvertYuanToLi(money float64) int64 {
	n1 := decimal.NewFromFloat(money)
	n2 := decimal.NewFromInt(100)
	result := n1.Mul(n2).IntPart()

	return result
}

// ConcatString
//
// @Title 字符串拼接
// @Description 字符串拼接
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 23:23:13
// @param args
// @return string
func ConcatString(args ...string) string {
	var stringBuffer bytes.Buffer
	if len(args) <= 1 {
		return ""
	}

	for _, v := range args {
		stringBuffer.WriteString(v)
	}

	return stringBuffer.String()
}

// ConvertUnitPercentMul
//
// @Title 百分比转换
// @Description 百分比转换
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 23:23:25
// @param num
// @param unit
// @param percent
// @return float64
func ConvertUnitPercentMul(num float64, unit int64, percent float64) float64 {
	// 被乘数
	n1 := decimal.NewFromFloat(num)
	// 转成百分比
	n2 := decimal.NewFromFloat(percent)
	n2 = n2.Div(decimal.NewFromInt(unit))
	// 计算乘法
	result, _ := n1.Mul(n2).Float64()

	return result
}
