package numeric

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

func FillValue[T any](record int, value T) []T {
	data := make([]T, record)
	for i := 0; i < record; i++ {
		data[i] = value
	}
	return data
}

func FillNumberString(size int) []string {
	m := []string{}
	for i := 1; i <= size; i++ {
		m = append(m, fmt.Sprint(i))
	}
	return m
}

func FillZeroFloat32(size int) []float32 {
	m := []float32{}
	for i := 1; i <= size; i++ {
		m = append(m, 0)
	}
	return m
}

func FillZeroFloat64(size int) []float64 {
	m := []float64{}
	for i := 1; i <= size; i++ {
		m = append(m, 0)
	}
	return m
}

func Format2Digit(num float64) string {
	return strconv.FormatFloat(num, 'f', 2, 64)
}

func AppendZeroInt(value int) string {
	n := fmt.Sprint(value)
	if value <= 9 {
		n = fmt.Sprintf("0%s", n)
	}
	return n
}

func JoinInt64(source []int64, sep string) string {
	stringNumbers := make([]string, len(source))
	for i, num := range source {
		stringNumbers[i] = fmt.Sprint(num)
	}
	return strings.Join(stringNumbers, sep)
}

func F64ToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func I64ToString(num int64) string {
	return strconv.FormatInt(num, 64)
}

func F64ToStringDyn(num float64) string {
	return fmt.Sprintf("%v", num)
}

func Rand(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func IsDigit(s string) bool {
	s = strings.TrimPrefix(s, "-")

	if len(s) == 0 {
		return false
	}

	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func FindInt(a []int64, x int64) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func FindIntPtr(a *[]int64, x int64) bool {
	if a == nil {
		return false
	}
	for _, n := range *a {
		if x == n {
			return true
		}
	}
	return false
}

func ParseInt(num string) (int, error) {
	return strconv.Atoi(num)
}

func ParseInt64(num string) (int64, error) {
	return strconv.ParseInt(num, 10, 64)
}

func ParseInt64Safety(num string) int64 {
	vi, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0
	}
	return vi
}

func ParseFloat64(num string) (float64, error) {
	return strconv.ParseFloat(num, 64)
}

func ParseFloat64Safety(num string) float64 {
	vf, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0
	}
	return vf
}

func ParseFloat32Safety(num string) float32 {
	vf, err := strconv.ParseFloat(num, 32)
	if err != nil {
		return 0
	}
	return float32(vf)
}

func ToFloat64(value interface{}) float64 {
	if value == nil {
		return 0.0
	}
	return value.(float64)
}

func SumFloat64(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total
}

func MoreThanZeroFloat64(value float64) bool {
	return value > 0
}
