package utility

import (
	"fmt"
	"reflect"

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

// ArrayDiff
//
// @Title 数组差集
// @Description 数组差集 泛型版本
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-14 00:25:55
// @param array1
// @param array2
// @return []T
func ArrayDiff[T IfInt | IfUint | string](array1, array2 []T) []T {
	m := make(map[T]bool)

	for _, x := range array2 {
		m[x] = true
	}

	var diff []T

	for _, x := range array1 {
		if !m[x] {
			diff = append(diff, x)
		}
	}

	return diff
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

// ArrayCast
//
// @Title 强制装换成其他类型
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 16:46:19
// @param array
// @param f
// @return []U
func ArrayCast[T IfInt | IfUint | string | interface{}, U IfInt | IfUint | string](array []T, f func(T) U) []U {
	result := make([]U, len(array))
	for i, v := range array {
		result[i] = f(v)
	}
	return result
}

// ArrayColumnCast
//
// @Title 取出map数组其中一列,并强制装换成其他类型
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 16:45:33
// @param data
// @param column
// @param f
// @return []U
func ArrayColumnCast[T IfInt | IfUint | string | interface{}, U IfInt | IfUint | string](data []map[string]T, column string, f func(T) U) []U {
	columnValues := make([]U, len(data))
	for i, item := range data {
		columnValues[i] = f(item[column])
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

// InArray
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 16:20:33
// @param val
// @param array
// @return bool
func InArray(val interface{}, array interface{}) bool {
	value := reflect.ValueOf(array)
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			if value.Index(i).Interface() == val {
				return true
			}
		}
	case reflect.Map:
		if value.MapIndex(reflect.ValueOf(val)).IsValid() {
			return true
		}
	}

	return false
}

// ArraySetMapKey
//
// @Title 从给定切片中创建一个新的集合
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 16:37:30
// @param slice
// @return func(T) bool
func ArraySetMapKey[T IfInt | IfUint | string](slice []T) func(T) bool {
	set := make(map[T]bool)

	// 将每个元素添加到集合中
	for _, val := range slice {
		set[val] = true
	}

	// 返回一个闭包，该闭包接受一个元素并检查该元素是否在集合中
	return func(elem T) bool {
		_, exists := set[elem]
		return exists
	}
}

// MapDiff
//
// @Title Map差集
// @Description Map差集 泛型版本
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-13 18:22:19
// @param map1
// @param map2
// @return map[string]T
func MapDiff[T IfInt | IfUint | string](map1, map2 map[string]T) map[string]T {
	result := make(map[string]T)
	for key, value := range map1 {
		if map2Value, ok := map2[key]; !ok || value != map2Value {
			result[key] = value
		}
	}
	for key, value := range map2 {
		if _, ok := map1[key]; !ok {
			result[key] = value
		}
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
// @Title Map类型转换
// @Description []map[string]interface{} 转换成 []map[string]string
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

// MapStrStr
//
// @Title Map类型转换
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-29 00:07:09
// @param data
// @return map[string]string
func MapStrStr(data map[string]interface{}) map[string]string {
	result := make(map[string]string, len(data))
	for k, v := range data {
		result[k] = fmt.Sprintf("%v", v)
	}
	return result
}

// MapsCast
//
// @Title Map类型转换
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-14 00:45:12
// @param maps
// @param key
// @return []map[string]T
//func MapsCast[T string](maps []map[string]interface{}) []map[string]T {
//	t := reflect.TypeOf(T{}).Kind()
//	result := make([]map[string]T, len(maps))
//	for i, val := range maps {
//		temp := make(map[string]T)
//		for k, v := range val {
//			if val, ok := v.(string); ok {
//				temp[k] = reflect.Interface{}T(val)
//			}
//		}
//		result[i] = temp
//	}
//	return result
//}

// InterfaceToSlice
//
// @Title interface转数组
// @Description  interface{} 转换为 []interface{}
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-14 18:14:20
// @param data
// @return []interface{}
func InterfaceToSlice(data interface{}) []interface{} {
	if slice, ok := data.([]interface{}); ok {
		return slice
	}

	// 如果 data 不是 []interface{} 类型，尝试将其转换为 []interface{}
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Slice {
		length := value.Len()
		slice := make([]interface{}, length)
		for i := 0; i < length; i++ {
			slice[i] = value.Index(i).Interface()
		}
		return slice
	}

	return nil
}
