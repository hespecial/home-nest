package tool

import (
	"fmt"
	"strconv"
)

// Fen2Yuan 分转元
func Fen2Yuan(fen int64) float64 {
	yuan, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(fen)/100), 64)
	return yuan
}
