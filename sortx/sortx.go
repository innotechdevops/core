package sortx

import (
	"strconv"
	"strings"
)

func SortZeroPrefix(order string, valueLeft, valueRight string) bool {
	// Convert to integers for proper numerical comparison
	storeNoI, errI := strconv.Atoi(strings.TrimPrefix(valueLeft, "0"))
	storeNoJ, errJ := strconv.Atoi(strings.TrimPrefix(valueRight, "0"))
	if errI != nil || errJ != nil {
		// Handle error if conversion fails, fallback to lexicographical
		if order == "asc" {
			return valueLeft < valueRight
		}
		return valueLeft > valueRight
	}
	if order == "asc" {
		return storeNoI < storeNoJ
	}
	return storeNoI > storeNoJ
}

func SortString(order string, valueLeft, valueRight string) bool {
	if order == "desc" {
		return valueLeft > valueRight
	}
	return valueLeft < valueRight
}

func SortFloat(order string, valueLeft, valueRight float64) bool {
	if order == "desc" {
		return valueLeft > valueRight
	}
	return valueLeft < valueRight
}
