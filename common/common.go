package common

import (
	"encoding/base64"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/innotechdevops/timex"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

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

func ClearCommaInline(text string) string {
	nText := strings.ReplaceAll(text, ",", "")
	nText = strings.ReplaceAll(nText, "\n", "")
	return nText
}

func SplitDash(input string) []string {
	if input == "" {
		return []string{}
	}
	return strings.Split(input, "-")
}

func ParseConnectionString[T any](connectionString string, target *T) error {
	cfg := make(map[string]string)
	pairs := strings.Split(connectionString, ";")

	for _, pair := range pairs {
		keyValue := strings.SplitN(pair, "=", 2)
		if len(keyValue) != 2 {
			return fmt.Errorf("invalid key-value pair: %s", pair)
		}
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		cfg[key] = value
	}

	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, target); err != nil {
		return err
	}
	return nil
}

func JoinInt64(source []int64, sep string) string {
	stringNumbers := make([]string, len(source))
	for i, num := range source {
		stringNumbers[i] = fmt.Sprint(num)
	}
	return strings.Join(stringNumbers, sep)
}

func RemoveDoubleQuote(text string) string {
	text = ClearUnicode(text)
	first := strings.Index(text, "\"")
	last := strings.LastIndex(text, "\"")
	if first == 0 && last == (len(text)-1) {
		text = text[1:last]
	}
	return text
}

func ClearUnicode(text string) string {
	regex := regexp.MustCompile("^\ufeff")
	result := regex.ReplaceAllString(text, "")
	return strings.TrimSpace(result)
}

func TrimLower(text string) string {
	return strings.TrimSpace(strings.ToLower(text))
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

func IsFloat(t reflect.Type) bool {
	if t.Kind() == reflect.Float32 || t == reflect.TypeOf(float64(0)) {
		return true
	}
	return false
}

func IsInt(t reflect.Type) bool {
	if t.Kind() == reflect.Int32 || t.Kind() == reflect.Int64 {
		return true
	}
	return false
}

func Rand(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func IsEmpty(val string) bool {
	if val == "" {
		return true
	} else {
		return false
	}
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

func RemoveExtension(filename string) string {
	spl := strings.Split(filename, ".")
	if len(spl) > 1 {
		return spl[0]
	}
	return filename
}

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func Find(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func FindInt(a []int64, x int64) bool {
	for _, n := range a {
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

func Contains(path string, contain ...string) bool {
	found := 0
	for _, c := range contain {
		if strings.Contains(path, c) {
			found++
		}
	}
	return found > 0
}

func Match(find string, source ...string) bool {
	found := 0
	for _, c := range source {
		if find == c {
			found++
		}
	}
	return found > 0
}

func IsExpired(jwt string, key string) bool {
	token := strings.Split(jwt, ".")
	payload, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(token[1])
	if err != nil {
		return true
	}

	m := map[string]interface{}{}
	if err := json.Unmarshal(payload, &m); err != nil {
		return true
	}

	if timestamp, ok := m[key].(float64); ok {
		current := timex.Now().Unix()
		if current < int64(timestamp) {
			return false
		}
	}
	return true
}

func ToFloat64(value interface{}) float64 {
	if value == nil {
		return 0.0
	}
	return value.(float64)
}

func RemoveIndex(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
}

func HasCode(id string, index int) string {
	return fmt.Sprintf("%sX%d", id, index)
}

func Base64FromByteToByte(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func Base64FromByte(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Base64FromString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func FloatDecimal(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func TrimToLower(text string) string {
	return strings.ToLower(strings.ReplaceAll(text, " ", ""))
}

func GetCodeFromName(str string) string {
	lowerStr := strings.ToLower(str)
	code := strings.ReplaceAll(lowerStr, " ", "")
	return code
}
