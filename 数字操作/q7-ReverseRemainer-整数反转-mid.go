package 数字操作

import (
	"math"
)

func reverse(x int) int {
	revNum := 0
	if x == 0 {
		return 0
	}
	flag := 1
	if x < 0{
		flag *= -1
		x *= -1
	}
	for x!=0{
		modnum := x % 10
		x /= 10
		revNum = revNum * 10 + modnum
	}
	if x > math.MaxInt32{
		return 0
	}
	return revNum * flag
}

