package pkg

import (
	"log"
	"math"
	"strconv"
	"strings"
)

// Convert: return default instead of err
func ToInt(s string, d int64) int64 {
	num, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return d
	}
	return num
}
func ToFloat(s string, precision int, f float64) float64 {
	num, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Println(err)
		return f
	}
	if precision > 0 {
		// 四舍五入
		shift := math.Pow(10, float64(precision))
		return math.Round(num*shift) / shift
	}
	return num
}
func ToBool(s string, d bool) bool {
	switch strings.ToLower(s) {
	case "true":
		return true
	case "false":
		return false
	default:
		return d
	}
}

// Unicode
func UnEscapeUnicodeString(s string) string {
	q_str := strconv.Quote(s)
	rp_str := strings.Replace(q_str, `\\u`, `\u`, -1)
	uq_str, err := strconv.Unquote(rp_str)
	if err != nil {
		return err.Error()
	}
	return uq_str
}

func UnEscapeUnicodeBytes(bytes []byte) string {
	return UnEscapeUnicodeString(string(bytes))
}

// compare
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// truncate
func TruncateStr(s string, w string) string {
	// Find the last index of the dot character.
	lastIndex := strings.LastIndex(s, w)
	if lastIndex == -1 {
		// No dot found, return the original string.
		return s
	}
	// Return the substring up to (but not including) the dot.
	return s[:lastIndex]
}

// Ternary 模拟三目运算符的函数
func Ternary[T comparable](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func Trim[T comparable](array []T, e T) []T {
	if len(array) < 1 {
		return array
	}
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] != e {
			return array[:i+1]
		}
	}
	return array
}
func MaskSensitive(source string, unMaskLen int) string {
	// aaaa, 3 > aaa###
	// aaaaa, 6 > aaaa###
	// a, 3 > a
	if len(source) > 2 {
		validLen := len(source)
		if validLen > unMaskLen {
			validLen = unMaskLen
		} else {
			validLen -= 1
		}
		return source[:validLen] + "###"
	}
	return source
}
