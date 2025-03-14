package tool

import (
	"math/rand"
)

const (
	KcRandKindNum   = 0 // 纯数字
	KcRandKindLower = 1 // 小写字母
	KcRandKindUpper = 2 // 大写字母
	KcRandKindAll   = 3 // 数字、大小写字母
)

// Krand 随机字符串
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
