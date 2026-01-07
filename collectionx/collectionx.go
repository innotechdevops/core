package collectionx

import "github.com/samber/lo"

func First[T any](list []T) *T {
	if len(list) == 0 {
		return nil
	}
	return &list[0]
}

func Last[T any](list []T) *T {
	size := len(list)
	if size == 0 {
		return nil
	}
	return &list[size-1]
}

// MapToSlice
// How to use:
// slices := MapToSlice(map[string]*int{"key1": pointer.Int(1), "key2": pointer.Int(2)})
func MapToSlice[T any](sums map[string]*T) []T {
	result := make([]T, 0, len(sums))
	for _, typed := range sums {
		result = append(result, *typed)
	}
	return result
}

// SliceToMap
// How to use:
// sums := []int{1, 2, 3}
// keyFunc := func(i int) string { return fmt.Sprintf("key-%d", i) }
// result := SliceToMap(sums, keyFunc)
func SliceToMap[T any](sums []T, keyFunc func(T) string) map[string]*T {
	result := make(map[string]*T, len(sums))
	for i := range sums {
		key := keyFunc(sums[i])
		result[key] = &sums[i]
	}
	return result
}

func Find[T any](collection []T, predicate func(item T) bool) (T, bool) {
	return lo.Find(collection, predicate)
}

func RemoveIndex(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
}

func IsIn[T comparable](arr []T, val T) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

func IsEmptyList[T any](value *[]T) bool {
	return value == nil || len(*value) == 0
}

func IsNotEmptyList[T any](value *[]T) bool {
	return !IsEmptyList(value)
}
