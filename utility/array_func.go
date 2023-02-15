package utility

// Intersect
//
// @Title 数组交集
// @Description 数组交集 泛型版本
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-14 17:35:46
// @param a
// @param b
// @return []T
func Intersect[T IfInt | IfUint | string](array1, array2 []T) []T {
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
