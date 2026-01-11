package csvx

import (
	"encoding/csv"
	"mime/multipart"

	c "github.com/prongbang/csvx"
)

func Convert[T any](data []T, ignoreDoubleQuote ...bool) string {
	return c.Convert[T](data, ignoreDoubleQuote...)
}

func ManualConvert[T any](data []T, headers []string, onRecord func(data T) []string) string {
	return c.ManualConvert[T](data, headers, onRecord)
}

func TryConvert[T any](data []T, ignoreDoubleQuote ...bool) string {
	return c.TryConvert[T](data, ignoreDoubleQuote...)
}

func FileHeaderReader(fileHeader *multipart.FileHeader) ([][]string, error) {
	return c.FileHeaderReader(fileHeader)
}

func ParserString[T any](rows [][]string) []T {
	return c.ParserString[T](rows)
}

func Parser[T any](rows [][]string) []T {
	return c.Parser[T](rows)
}

func ParserFunc(excludeHeader bool, rows [][]string, onRecord func([]string) error) error {
	return c.ParserFunc(excludeHeader, rows, onRecord)
}

func ParserByReader[T any](ir *csv.Reader, delimiter ...rune) []T {
	return c.ParserByReader[T](ir, delimiter...)
}
