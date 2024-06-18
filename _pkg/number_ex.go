package pkg

import (
	"math"
)

func TruncateNum[T float64 | float32 | int32 | int](num T, precision ...int) float64 {
	// 截断小数部分，再除以100. 123.65 > 1.230000
	if float64(num) == 0.0 {
		return 0.0
	}
	if len(precision) == 0 {
		precision[0] = 2
		precision[1] = 2
	}
	mShift := math.Pow(10, float64(precision[0]))
	shift := math.Pow(10, float64(precision[1]))
	truncated := math.Trunc(float64(num)*mShift) / shift
	return float64(truncated)
}
func RoundNum[T float64 | float32 | int32](num T, precision ...int) float64 {
	// 四舍五入小数部分，再除以100.123.65 > 1.240000
	if len(precision) == 0 {
		precision[0] = 2
		precision[1] = 2
	}
	mShift := math.Pow(10, float64(precision[0]))
	shift := math.Pow(10, float64(precision[1]))
	truncated := math.Round(float64(num)*mShift) / shift
	return float64(truncated)
}
func CeilNum[T float64 | float32 | int32](num T, precision ...int) float64 {
	// 向上
	if len(precision) == 0 {
		precision[0] = 2
		precision[1] = 2
	}
	mShift := math.Pow(10, float64(precision[0]))
	shift := math.Pow(10, float64(precision[1]))
	truncated := math.Ceil(float64(num)*mShift) / shift
	return float64(truncated)
}
func ToYuan[T float64 | float32 | int32 | int](num T, precision ...int) float64 {
	return TruncateNum(num, 0, 2)
}

func ToFen[T float64 | float32 | int32 | int](num T, precision ...int) int32 {
	return int32(TruncateNum(num, 2, 0))
}
