package utility

import (
	"fmt"

	"github.com/gogf/gf/v2/util/gconv"
)

// ArrayIntersect
//
// @Title 数组交集
// @Description 数组交集 泛型版本
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-14 17:35:46
// @param a
// @param b
// @return []T
func ArrayIntersect[T IfInt | IfUint | string](array1, array2 []T) []T {
	res := make([]T, 0)
	tmp := make(map[T]bool, len(array1))

	for _, s := range array1 {
		if _, ok := tmp[s]; !ok {
			tmp[s] = true
		}
	}
	for _, s := range array2 {
		if _, ok := tmp[s]; ok {
			res = append(res, s)
		}
	}

	return res
}

// ArrayColumn
//
// @Title 取出map数组其中一列
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-08 13:40:16
// @param data
// @param column
// @return []T
func ArrayColumn[T IfInt | IfUint | string | interface{}](data []map[string]T, column string) []T {
	columnValues := make([]T, len(data))
	for i, item := range data {
		columnValues[i] = item[column]
	}
	return columnValues
}

// ArrayCombine
//
// @Title
// @Description 通过合并两个数组来创建一个新数组，其中的一个数组元素为键名，另一个数组元素为键值
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-08 15:03:35
// @param keys
// @param values
// @return map[string]interface{}
func ArrayCombine(keys []string, values []interface{}) map[string]interface{} {
	length := len(keys)
	if length != len(values) {
		panic("ArrayCombine: 数组长度不匹配")
	}

	result := make(map[string]interface{}, length)
	for i := 0; i < length; i++ {
		result[keys[i]] = values[i]
	}

	return result
}

// MapsFromColumns
//
// @Title 将一个map数组元素为键名，另一个数组元素为键值
// @Description 将一个map数组元素为键名，另一个数组元素为键值组成一个新数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-08 18:06:02
// @param maps
// @param keyCol
// @param valueCol
// @return map[string]interface{}
func MapsFromColumns(maps []map[string]interface{}, keyCol, valueCol string) map[string]interface{} {
	result := make(map[string]interface{}, len(maps))
	for _, m := range maps {
		result[gconv.String(m[keyCol])] = m[valueCol]
	}
	return result
}

// MapsStrStr
//
// @Title  []map[string]interface{} 转换成 []map[string]string
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-08 17:44:01
// @param maps
// @return []map[string]string
func MapsStrStr(maps []map[string]interface{}) []map[string]string {
	result := make([]map[string]string, len(maps))
	for i, m := range maps {
		temp := make(map[string]string)
		for k, v := range m {
			// 将 v 转换为 string 类型，并存储到 temp 中
			temp[k] = fmt.Sprintf("%v", v)
		}
		result[i] = temp
	}
	return result
}
