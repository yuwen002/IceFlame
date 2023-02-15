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
func Intersect[T IfInt | IfUint | string](a, b []T) []T {
	inter := make([]T, 0)
	mp := make(map[T]bool)

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}
