package mapping

import (
	"fmt"
	"reflect"

	"github.com/goccy/go-json"
)

/*
Benchmark:
	goos: darwin
	goarch: arm64
	pkg: github.com/innotechdevops/core/mapping
	Benchmark_Manual_Large_Data-10          	 3120075	       374.7 ns/op
	Benchmark_Manual_Small_Data-10          	 3552823	       331.5 ns/op
	Benchmark_TinyInto_Large_Data-10        	  334605	      3643 ns/op
	Benchmark_TinyInto_Small_Data-10        	  843255	      1316 ns/op
	Benchmark_Into_Large_Data-10            	  487419	      2302 ns/op
	Benchmark_Into_Small_Data-10            	  514274	      2277 ns/op
	Benchmark_MapStructure_Small_Data-10    	  103213	     12178 ns/op
	PASS
*/

/*
Into the function for copy data from input into output struct, Fast when large data, Single-Level Struct Support
Benchmark:

	Benchmark_Manual-10          	 3096486	       364.7 ns/op
	Benchmark_Into-10            	  487270	      2312 ns/op

How to use:

	type Input struct {
		Field1 int `json:"field1"`
		Field2 int `json:"field2"`
		Field3 int `json:"field3"`
		Field4 int `json:"field4"`
	}

	type Output struct {
		FieldFloat32  float32 `json:"field_float32"`
		FieldFloat64  float64 `json:"field_float64"`
		FieldInt      int     `json:"field_int"`
		FieldInt8     int8    `json:"field_int8"`
		FieldInt16    int16   `json:"field_int16"`
		FieldInt32    int32   `json:"field_int32"`
		FieldInt64    int64   `json:"field_int64"`
		FieldUint     uint    `json:"field_uint"`
		FieldUint8    uint8   `json:"field_uint8"`
		FieldUint16   uint16  `json:"field_uint16"`
		FieldUint32   uint32  `json:"field_uint32"`
		FieldUint64   uint64  `json:"field_uint64"`
		FieldString   string  `json:"field_string"`
		FieldBool     bool    `json:"field_bool"`
		FieldAny      any     `json:"field_any"`
		FieldError    error   `json:"field_error"`
	}

	input := Input{
		Field1: 1,
		Field2: 2,
		Field3: 3,
		Field4: 4,
	}

	var output Output
	Into(input, &output)
	fmt.Printf("%+v\n", output) // Output: {Field1:1 Field2:2}
*/
func Into[S any, D any](src S, dst *D) {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst).Elem()

	for i := 0; i < dstVal.NumField(); i++ {
		dstField := dstVal.Type().Field(i)
		srcField := srcVal.FieldByName(dstField.Name)

		if srcField.IsValid() && srcField.Type().AssignableTo(dstField.Type) {
			dstVal.Field(i).Set(srcField)
		}
	}
}

// TinyInto unsupported error type, Fast when small data, Multiple-Level Struct Support
func TinyInto[D any](src any, dst *D) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

// KeyValue Define the KeyValue interface
type KeyValue[K comparable, V any] interface {
	Key() K
	Value() V
}

// ToMap converts a slice of KeyValue into a map.
func ToMap[K comparable, V any, T KeyValue[K, V]](data []T) map[K]V {
	m := make(map[K]V)
	for _, v := range data {
		m[v.Key()] = v.Value()
	}
	return m
}

func ToList[S any, T any](data []S, onSelect func(value S) T) []T {
	result := make([]T, len(data))
	for _, v := range data {
		result = append(result, onSelect(v))
	}
	return result
}

func ToAsOfDate(year *[]int, month *[]string) []string {
	asOfDate := []string{}
	for _, y := range *year {
		for _, m := range *month {
			asOfDate = append(asOfDate, fmt.Sprintf("%d-%s-01", y, m))
		}
	}
	return asOfDate
}
