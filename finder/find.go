package finder

import "strings"

func Contains(path string, contain ...string) bool {
	found := 0
	for _, c := range contain {
		if strings.Contains(path, c) {
			found++
		}
	}
	return found > 0
}

func Contain[T any](list []T, value T, fn func(a, b T) bool) bool {
	for _, it := range list {
		if fn(it, value) {
			return true
		}
	}
	return false
}

func Match(maps map[string]bool, value string) bool {
	return maps[strings.ToLower(value)]
}

func MatchString(maps map[string]string, value string) string {
	return maps[value]
}
