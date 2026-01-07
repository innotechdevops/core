package stringx

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

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

func TrimToLower(text string) string {
	return strings.ToLower(strings.ReplaceAll(text, " ", ""))
}

func IsEmpty(value string) bool {
	return value == ""
}

func IsNotEmpty(value string) bool {
	return !IsEmpty(value)
}

func IsNotEmptyPtr(value *string) bool {
	if value == nil {
		return false
	}
	return !IsEmpty(*value)
}

func IsNotNull[T any](value *T) bool {
	if value == nil {
		return false
	}
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		return !v.IsNil()
	}
	return true
}

func GetStackTrace(e interface{}) string {
	buf := make([]byte, 1<<20)
	buf = buf[:runtime.Stack(buf, false)]
	msg := fmt.Sprintf("panic: %v\n%s\n", e, buf)
	return msg
}

func TrimBy(data string, separator string) string {
	if separator == "" {
		return data
	}
	return strings.ReplaceAll(data, separator, "")
}

var spaceRegex = regexp.MustCompile(`\s+`)

func NormalizeText(text string) string {
	// 0. Early return for empty strings
	if text == "" {
		return ""
	}

	// 1. Unicode normalization (NFC - recommended for most use cases)
	normalized := norm.NFC.String(text)

	// 2. Width normalization (fullwidth → halfwidth)
	normalized, _, _ = transform.String(width.Narrow, normalized)

	// 3. Trim spaces
	normalized = strings.TrimSpace(normalized)

	// 4. Remove/replace problematic unicode characters
	normalized = strings.ReplaceAll(normalized, "\u00A0", " ") // non-breaking space
	normalized = strings.ReplaceAll(normalized, "\u2009", " ") // thin space
	normalized = strings.ReplaceAll(normalized, "\u200A", " ") // hair space
	normalized = strings.ReplaceAll(normalized, "\u2028", " ") // line separator
	normalized = strings.ReplaceAll(normalized, "\u2029", " ") // paragraph separator
	normalized = strings.ReplaceAll(normalized, "\u200B", "")  // zero-width space
	normalized = strings.ReplaceAll(normalized, "\u200C", "")  // zero-width non-joiner
	normalized = strings.ReplaceAll(normalized, "\u200D", "")  // zero-width joiner
	normalized = strings.ReplaceAll(normalized, "\u2060", "")  // word joiner
	normalized = strings.ReplaceAll(normalized, "\uFEFF", "")  // byte order mark (BOM)
	normalized = strings.ReplaceAll(normalized, "\u061C", "")  // Arabic letter mark

	// 5. Replace multiple whitespace with single space
	normalized = spaceRegex.ReplaceAllString(normalized, " ")

	// 6. Final trim
	normalized = strings.TrimSpace(normalized)

	// 7. Return empty string if result is only whitespace
	if normalized == "" {
		return ""
	}

	return normalized
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StringInSliceTrimLower(a string, list []string) bool {
	for _, b := range list {
		if TrimLower(b) == TrimLower(a) {
			return true
		}
	}
	return false
}

func QueryUnescape(encoded string) string {
	decoded, _ := url.QueryUnescape(encoded)
	return decoded
}
