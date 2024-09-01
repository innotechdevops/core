package mapping_test

import (
	"github.com/innotechdevops/core/mapping"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"testing"
)

type Input struct {
	FieldFloat32 float32 `json:"field_float32"`
	FieldFloat64 float64 `json:"field_float64"`
	FieldInt     int     `json:"field_int"`
	FieldInt8    int8    `json:"field_int8"`
	FieldInt16   int16   `json:"field_int16"`
	FieldInt32   int32   `json:"field_int32"`
	FieldInt64   int64   `json:"field_int64"`
	FieldUint    uint    `json:"field_uint"`
	FieldUint8   uint8   `json:"field_uint8"`
	FieldUint16  uint16  `json:"field_uint16"`
	FieldUint32  uint32  `json:"field_uint32"`
	FieldUint64  uint64  `json:"field_uint64"`
	FieldString  string  `json:"field_string"`
	FieldBool    bool    `json:"field_bool"`
	FieldAny     any     `json:"field_any"`
	FieldError   error   `json:"field_error"`
}

type Output struct {
	FieldFloat32 float32 `json:"field_float32"`
	FieldFloat64 float64 `json:"field_float64"`
	FieldInt     int     `json:"field_int"`
	FieldInt8    int8    `json:"field_int8"`
	FieldInt16   int16   `json:"field_int16"`
	FieldInt32   int32   `json:"field_int32"`
	FieldInt64   int64   `json:"field_int64"`
	FieldUint    uint    `json:"field_uint"`
	FieldUint8   uint8   `json:"field_uint8"`
	FieldUint16  uint16  `json:"field_uint16"`
	FieldUint32  uint32  `json:"field_uint32"`
	FieldUint64  uint64  `json:"field_uint64"`
	FieldString  string  `json:"field_string"`
	FieldBool    bool    `json:"field_bool"`
	FieldAny     any     `json:"field_any"`
	FieldError   error   `json:"field_error"`
}

func Test_Into(t *testing.T) {
	input := Input{
		FieldFloat32: 3.14,
		FieldFloat64: 3.14159265359,
		FieldInt:     42,
		FieldInt8:    -8,
		FieldInt16:   16,
		FieldInt32:   32,
		FieldInt64:   64,
		FieldUint:    42,
		FieldUint8:   8,
		FieldUint16:  16,
		FieldUint32:  32,
		FieldUint64:  64,
		FieldString:  "example string",
		FieldBool:    true,
		FieldAny:     "any type of value", // Can be any value, like a string, int, etc.
		FieldError:   errors.New("example error"),
	}

	var output Output
	mapping.Into(input, &output)

	if output.FieldFloat32 != input.FieldFloat32 {
		t.Error("Error output FieldFloat32 is", output.FieldFloat32)
	}
	if output.FieldFloat64 != input.FieldFloat64 {
		t.Error("Error output FieldFloat64 is", output.FieldFloat64)
	}
	if output.FieldInt != input.FieldInt {
		t.Error("Error output FieldInt is", output.FieldInt)
	}
	if output.FieldInt8 != input.FieldInt8 {
		t.Error("Error output FieldInt8 is", output.FieldInt8)
	}
	if output.FieldInt16 != input.FieldInt16 {
		t.Error("Error output FieldInt16 is", output.FieldInt16)
	}
	if output.FieldInt32 != input.FieldInt32 {
		t.Error("Error output FieldInt32 is", output.FieldInt32)
	}
	if output.FieldInt64 != input.FieldInt64 {
		t.Error("Error output FieldInt64 is", output.FieldInt64)
	}
	if output.FieldUint != input.FieldUint {
		t.Error("Error output FieldUint is", output.FieldUint)
	}
	if output.FieldUint8 != input.FieldUint8 {
		t.Error("Error output FieldUint8 is", output.FieldUint8)
	}
	if output.FieldUint16 != input.FieldUint16 {
		t.Error("Error output FieldUint16 is", output.FieldUint16)
	}
	if output.FieldUint32 != input.FieldUint32 {
		t.Error("Error output FieldUint32 is", output.FieldUint32)
	}
	if output.FieldUint64 != input.FieldUint64 {
		t.Error("Error output FieldUint64 is", output.FieldUint64)
	}
	if output.FieldString != input.FieldString {
		t.Error("Error output FieldString is", output.FieldString)
	}
	if output.FieldBool != input.FieldBool {
		t.Error("Error output FieldBool is", output.FieldBool)
	}
	if output.FieldAny != input.FieldAny {
		t.Error("Error output FieldAny is", output.FieldAny)
	}
	if !errors.Is(output.FieldError, input.FieldError) {
		t.Error("Error output FieldError is", output.FieldError)
	}
}

func Benchmark_Manual_Large_Data(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := Input{
			FieldFloat32: 3.14,
			FieldFloat64: 3.14159265359,
			FieldInt:     42,
			FieldInt8:    -8,
			FieldInt16:   16,
			FieldInt32:   32,
			FieldInt64:   64,
			FieldUint:    42,
			FieldUint8:   8,
			FieldUint16:  16,
			FieldUint32:  32,
			FieldUint64:  64,
			FieldString:  "d20b0c50263e33b57ea684260dd3f746888141a04dd4df808c475120bd234b17ebd1f668e4419d2329fc2bbf585ce37a4eec732d0f65913777be5acc33da58f2c262e76c4fa3c4e1ca3949193710b8a7ab31467434417d12e161e06fda2542123ac6c7f296a9a0cddebedca1d531459d7c855af340fd1961744d73f9db4f0e943f8799d2288f42abcfce144b237fe98ae6dd8993f64de0909c706892d6ba3e4cfc8890e39a74befca711de32b3f4a81085c710b8c52696e85ae0a646b766f4d1911daca43a58490eca987b55516dca81a27e6f43a9e4c9171032dae60abb597c616d6bd27c487b1375c0f3f8dee4774dfaea8e24dc2ce3e6aa465e30dd2b82bd9063773ee65aa86dbfd4c8503fa3bf991583b876f0dac417aa4fcda79dde30916eca16ceb2b8dcdca34ea977ac39802659702295b5efb5ca5ce91cf6a1930dd5308bab8436ce9ba555b20d22be1dfd2acb30328b6e1fa6f5f1c8e9baca46a2971e972983cdafa69d8e6db70e3fce7f6d4a486d6428addf4bde83b8acd0a711ee6af49ed6366d5847ce56e5718e13e2f2f99f1fd260331c561d34f5130f98e61b1ea18b020d7865fc2d38a3be1bc4e1316afcf27e609c5d8716859a3af655944e969fa6c9cba8d70fb66e4e4302c59d118048cc4807a64ba16a20d8d975d79b2636cb47b348c4685575ea13669d924d75c4a6c61bc5a8376f0c53591e7a25f54cfeb664bdd0e6e4716f7e6232f63f1ec4bb1b65f347ba1f66d5d3fb2e9336a4fd605651627c6802c7b5328b791d16387408c6981c836ea5432ffcf44adf253ef2f8af81a46f2617d463e6c9adda42544324b875a8ada57dbb1fa392ce9778787225809c1ca07021b646048d0fee5885124db54de03ab2263207df9f5455bc2e4f51b33c6659cc0023e77605a8104c57d7f20d457a87b6c028e31f5421b25de0f24813ce122ad5b5157a851d3b0b888ab0c4e2b5f12d581be4716edd93e2f701c8d05be791731dc8f0c00bc01020bd213803e4209fcc10d08128ed1a1ac98d38af818561a2141a7c4f046ca9283a74d8277f85231e5cc3dff9046e08379152f5d01445924f691bed702907ea9abbaf6343d400204cd81260794b5dc5ff72eaee255e40c1f70e018b11fe74599509e2bf4cd306339e03a839d393d7b5c3a28498c819e1dbbbb6c05b13112caff4f317fbd6d21324a0bdfbac8938dd7669b2874951617a02dad6400891ee9bb55f24e1879d4f90b369e1f9c8080c660ebcdac429c2e8302fb7b68b69f6a0083d1d2265cb5e71f43c4bc2e12a98dff892d33811510fa44ce6461626b409d7c4886f8eea1a01f9b779179943f79ab986cba75ccad88c08a7b94d7df77204b41708a925656a4106d721df4469b4e692c80705c35a69a300981d63e525392f",
			FieldBool:    true,
			FieldAny:     "any type of value", // Can be any value, like a string, int, etc.
			FieldError:   errors.New("example error"),
		}

		output := Output{
			FieldFloat32: input.FieldFloat32,
			FieldFloat64: input.FieldFloat64,
			FieldInt:     input.FieldInt,
			FieldInt8:    input.FieldInt8,
			FieldInt16:   input.FieldInt16,
			FieldInt32:   input.FieldInt32,
			FieldInt64:   input.FieldInt64,
			FieldUint:    input.FieldUint,
			FieldUint8:   input.FieldUint8,
			FieldUint16:  input.FieldUint16,
			FieldUint32:  input.FieldUint32,
			FieldUint64:  input.FieldUint64,
			FieldString:  input.FieldString,
			FieldBool:    input.FieldBool,
			FieldAny:     input.FieldAny,
			FieldError:   input.FieldError,
		}

		if output.FieldFloat32 != input.FieldFloat32 {
			b.Error("Error output FieldFloat32 is", output.FieldFloat32)
		}
		if output.FieldFloat64 != input.FieldFloat64 {
			b.Error("Error output FieldFloat64 is", output.FieldFloat64)
		}
		if output.FieldInt != input.FieldInt {
			b.Error("Error output FieldInt is", output.FieldInt)
		}
		if output.FieldInt8 != input.FieldInt8 {
			b.Error("Error output FieldInt8 is", output.FieldInt8)
		}
		if output.FieldInt16 != input.FieldInt16 {
			b.Error("Error output FieldInt16 is", output.FieldInt16)
		}
		if output.FieldInt32 != input.FieldInt32 {
			b.Error("Error output FieldInt32 is", output.FieldInt32)
		}
		if output.FieldInt64 != input.FieldInt64 {
			b.Error("Error output FieldInt64 is", output.FieldInt64)
		}
		if output.FieldUint != input.FieldUint {
			b.Error("Error output FieldUint is", output.FieldUint)
		}
		if output.FieldUint8 != input.FieldUint8 {
			b.Error("Error output FieldUint8 is", output.FieldUint8)
		}
		if output.FieldUint16 != input.FieldUint16 {
			b.Error("Error output FieldUint16 is", output.FieldUint16)
		}
		if output.FieldUint32 != input.FieldUint32 {
			b.Error("Error output FieldUint32 is", output.FieldUint32)
		}
		if output.FieldUint64 != input.FieldUint64 {
			b.Error("Error output FieldUint64 is", output.FieldUint64)
		}
		if output.FieldString != input.FieldString {
			b.Error("Error output FieldString is", output.FieldString)
		}
		if output.FieldBool != input.FieldBool {
			b.Error("Error output FieldBool is", output.FieldBool)
		}
		if output.FieldAny != input.FieldAny {
			b.Error("Error output FieldAny is", output.FieldAny)
		}
		if !errors.Is(output.FieldError, input.FieldError) {
			b.Error("Error output FieldError is", output.FieldError)
		}
	}
}

func Benchmark_Manual_Small_Data(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := Input{
			FieldFloat32: 3.14,
			FieldFloat64: 3.14159265359,
			FieldInt:     42,
			FieldInt8:    -8,
			FieldInt16:   16,
			FieldInt32:   32,
			FieldInt64:   64,
			FieldUint:    42,
			FieldUint8:   8,
			FieldUint16:  16,
			FieldUint32:  32,
			FieldUint64:  64,
			FieldString:  "example",
			FieldBool:    true,
			FieldAny:     "any type of value", // Can be any value, like a string, int, etc.
			FieldError:   errors.New("example error"),
		}

		output := Output{
			FieldFloat32: input.FieldFloat32,
			FieldFloat64: input.FieldFloat64,
			FieldInt:     input.FieldInt,
			FieldInt8:    input.FieldInt8,
			FieldInt16:   input.FieldInt16,
			FieldInt32:   input.FieldInt32,
			FieldInt64:   input.FieldInt64,
			FieldUint:    input.FieldUint,
			FieldUint8:   input.FieldUint8,
			FieldUint16:  input.FieldUint16,
			FieldUint32:  input.FieldUint32,
			FieldUint64:  input.FieldUint64,
			FieldString:  input.FieldString,
			FieldBool:    input.FieldBool,
			FieldAny:     input.FieldAny,
			FieldError:   input.FieldError,
		}

		if output.FieldFloat32 != input.FieldFloat32 {
			b.Error("Error output FieldFloat32 is", output.FieldFloat32)
		}
		if output.FieldFloat64 != input.FieldFloat64 {
			b.Error("Error output FieldFloat64 is", output.FieldFloat64)
		}
		if output.FieldInt != input.FieldInt {
			b.Error("Error output FieldInt is", output.FieldInt)
		}
		if output.FieldInt8 != input.FieldInt8 {
			b.Error("Error output FieldInt8 is", output.FieldInt8)
		}
		if output.FieldInt16 != input.FieldInt16 {
			b.Error("Error output FieldInt16 is", output.FieldInt16)
		}
		if output.FieldInt32 != input.FieldInt32 {
			b.Error("Error output FieldInt32 is", output.FieldInt32)
		}
		if output.FieldInt64 != input.FieldInt64 {
			b.Error("Error output FieldInt64 is", output.FieldInt64)
		}
		if output.FieldUint != input.FieldUint {
			b.Error("Error output FieldUint is", output.FieldUint)
		}
		if output.FieldUint8 != input.FieldUint8 {
			b.Error("Error output FieldUint8 is", output.FieldUint8)
		}
		if output.FieldUint16 != input.FieldUint16 {
			b.Error("Error output FieldUint16 is", output.FieldUint16)
		}
		if output.FieldUint32 != input.FieldUint32 {
			b.Error("Error output FieldUint32 is", output.FieldUint32)
		}
		if output.FieldUint64 != input.FieldUint64 {
			b.Error("Error output FieldUint64 is", output.FieldUint64)
		}
		if output.FieldString != input.FieldString {
			b.Error("Error output FieldString is", output.FieldString)
		}
		if output.FieldBool != input.FieldBool {
			b.Error("Error output FieldBool is", output.FieldBool)
		}
		if output.FieldAny != input.FieldAny {
			b.Error("Error output FieldAny is", output.FieldAny)
		}
		if !errors.Is(output.FieldError, input.FieldError) {
			b.Error("Error output FieldError is", output.FieldError)
		}
	}
}

func Benchmark_TinyInto_Large_Data(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := Input{
			FieldFloat32: 3.14,
			FieldFloat64: 3.14159265359,
			FieldInt:     42,
			FieldInt8:    -8,
			FieldInt16:   16,
			FieldInt32:   32,
			FieldInt64:   64,
			FieldUint:    42,
			FieldUint8:   8,
			FieldUint16:  16,
			FieldUint32:  32,
			FieldUint64:  64,
			FieldString:  "d20b0c50263e33b57ea684260dd3f746888141a04dd4df808c475120bd234b17ebd1f668e4419d2329fc2bbf585ce37a4eec732d0f65913777be5acc33da58f2c262e76c4fa3c4e1ca3949193710b8a7ab31467434417d12e161e06fda2542123ac6c7f296a9a0cddebedca1d531459d7c855af340fd1961744d73f9db4f0e943f8799d2288f42abcfce144b237fe98ae6dd8993f64de0909c706892d6ba3e4cfc8890e39a74befca711de32b3f4a81085c710b8c52696e85ae0a646b766f4d1911daca43a58490eca987b55516dca81a27e6f43a9e4c9171032dae60abb597c616d6bd27c487b1375c0f3f8dee4774dfaea8e24dc2ce3e6aa465e30dd2b82bd9063773ee65aa86dbfd4c8503fa3bf991583b876f0dac417aa4fcda79dde30916eca16ceb2b8dcdca34ea977ac39802659702295b5efb5ca5ce91cf6a1930dd5308bab8436ce9ba555b20d22be1dfd2acb30328b6e1fa6f5f1c8e9baca46a2971e972983cdafa69d8e6db70e3fce7f6d4a486d6428addf4bde83b8acd0a711ee6af49ed6366d5847ce56e5718e13e2f2f99f1fd260331c561d34f5130f98e61b1ea18b020d7865fc2d38a3be1bc4e1316afcf27e609c5d8716859a3af655944e969fa6c9cba8d70fb66e4e4302c59d118048cc4807a64ba16a20d8d975d79b2636cb47b348c4685575ea13669d924d75c4a6c61bc5a8376f0c53591e7a25f54cfeb664bdd0e6e4716f7e6232f63f1ec4bb1b65f347ba1f66d5d3fb2e9336a4fd605651627c6802c7b5328b791d16387408c6981c836ea5432ffcf44adf253ef2f8af81a46f2617d463e6c9adda42544324b875a8ada57dbb1fa392ce9778787225809c1ca07021b646048d0fee5885124db54de03ab2263207df9f5455bc2e4f51b33c6659cc0023e77605a8104c57d7f20d457a87b6c028e31f5421b25de0f24813ce122ad5b5157a851d3b0b888ab0c4e2b5f12d581be4716edd93e2f701c8d05be791731dc8f0c00bc01020bd213803e4209fcc10d08128ed1a1ac98d38af818561a2141a7c4f046ca9283a74d8277f85231e5cc3dff9046e08379152f5d01445924f691bed702907ea9abbaf6343d400204cd81260794b5dc5ff72eaee255e40c1f70e018b11fe74599509e2bf4cd306339e03a839d393d7b5c3a28498c819e1dbbbb6c05b13112caff4f317fbd6d21324a0bdfbac8938dd7669b2874951617a02dad6400891ee9bb55f24e1879d4f90b369e1f9c8080c660ebcdac429c2e8302fb7b68b69f6a0083d1d2265cb5e71f43c4bc2e12a98dff892d33811510fa44ce6461626b409d7c4886f8eea1a01f9b779179943f79ab986cba75ccad88c08a7b94d7df77204b41708a925656a4106d721df4469b4e692c80705c35a69a300981d63e525392f",
			FieldBool:    true,
			FieldAny:     "any type of value", // Can be any value, like a string, int, etc.
			FieldError:   errors.New("example error"),
		}

		var output Output
		_ = mapping.TinyInto(input, &output)

		if output.FieldFloat32 != input.FieldFloat32 {
			b.Error("Error output FieldFloat32 is", output.FieldFloat32)
		}
		if output.FieldFloat64 != input.FieldFloat64 {
			b.Error("Error output FieldFloat64 is", output.FieldFloat64)
		}
		if output.FieldInt != input.FieldInt {
			b.Error("Error output FieldInt is", output.FieldInt)
		}
		if output.FieldInt8 != input.FieldInt8 {
			b.Error("Error output FieldInt8 is", output.FieldInt8)
		}
		if output.FieldInt16 != input.FieldInt16 {
			b.Error("Error output FieldInt16 is", output.FieldInt16)
		}
		if output.FieldInt32 != input.FieldInt32 {
			b.Error("Error output FieldInt32 is", output.FieldInt32)
		}
		if output.FieldInt64 != input.FieldInt64 {
			b.Error("Error output FieldInt64 is", output.FieldInt64)
		}
		if output.FieldUint != input.FieldUint {
			b.Error("Error output FieldUint is", output.FieldUint)
		}
		if output.FieldUint8 != input.FieldUint8 {
			b.Error("Error output FieldUint8 is", output.FieldUint8)
		}
		if output.FieldUint16 != input.FieldUint16 {
			b.Error("Error output FieldUint16 is", output.FieldUint16)
		}
		if output.FieldUint32 != input.FieldUint32 {
			b.Error("Error output FieldUint32 is", output.FieldUint32)
		}
		if output.FieldUint64 != input.FieldUint64 {
			b.Error("Error output FieldUint64 is", output.FieldUint64)
		}
		if output.FieldString != input.FieldString {
			b.Error("Error output FieldString is", output.FieldString)
		}
		if output.FieldBool != input.FieldBool {
			b.Error("Error output FieldBool is", output.FieldBool)
		}
		if output.FieldAny != input.FieldAny {
			b.Error("Error output FieldAny is", output.FieldAny)
		}
	}
}

func Benchmark_TinyInto_Small_Data(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := Input{
			FieldFloat32: 3.14,
			FieldFloat64: 3.14159265359,
			FieldInt:     42,
			FieldInt8:    -8,
			FieldInt16:   16,
			FieldInt32:   32,
			FieldInt64:   64,
			FieldUint:    42,
			FieldUint8:   8,
			FieldUint16:  16,
			FieldUint32:  32,
			FieldUint64:  64,
			FieldString:  "example",
			FieldBool:    true,
			FieldAny:     "any type of value", // Can be any value, like a string, int, etc.
			FieldError:   errors.New("example error"),
		}

		var output Output
		_ = mapping.TinyInto(input, &output)

		if output.FieldFloat32 != input.FieldFloat32 {
			b.Error("Error output FieldFloat32 is", output.FieldFloat32)
		}
		if output.FieldFloat64 != input.FieldFloat64 {
			b.Error("Error output FieldFloat64 is", output.FieldFloat64)
		}
		if output.FieldInt != input.FieldInt {
			b.Error("Error output FieldInt is", output.FieldInt)
		}
		if output.FieldInt8 != input.FieldInt8 {
			b.Error("Error output FieldInt8 is", output.FieldInt8)
		}
		if output.FieldInt16 != input.FieldInt16 {
			b.Error("Error output FieldInt16 is", output.FieldInt16)
		}
		if output.FieldInt32 != input.FieldInt32 {
			b.Error("Error output FieldInt32 is", output.FieldInt32)
		}
		if output.FieldInt64 != input.FieldInt64 {
			b.Error("Error output FieldInt64 is", output.FieldInt64)
		}
		if output.FieldUint != input.FieldUint {
			b.Error("Error output FieldUint is", output.FieldUint)
		}
		if output.FieldUint8 != input.FieldUint8 {
			b.Error("Error output FieldUint8 is", output.FieldUint8)
		}
		if output.FieldUint16 != input.FieldUint16 {
			b.Error("Error output FieldUint16 is", output.FieldUint16)
		}
		if output.FieldUint32 != input.FieldUint32 {
			b.Error("Error output FieldUint32 is", output.FieldUint32)
		}
		if output.FieldUint64 != input.FieldUint64 {
			b.Error("Error output FieldUint64 is", output.FieldUint64)
		}
		if output.FieldString != input.FieldString {
			b.Error("Error output FieldString is", output.FieldString)
		}
		if output.FieldBool != input.FieldBool {
			b.Error("Error output FieldBool is", output.FieldBool)
		}
		if output.FieldAny != input.FieldAny {
			b.Error("Error output FieldAny is", output.FieldAny)
		}
	}
}

func Benchmark_Into_Large_Data(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := Input{
			FieldFloat32: 3.14,
			FieldFloat64: 3.14159265359,
			FieldInt:     42,
			FieldInt8:    -8,
			FieldInt16:   16,
			FieldInt32:   32,
			FieldInt64:   64,
			FieldUint:    42,
			FieldUint8:   8,
			FieldUint16:  16,
			FieldUint32:  32,
			FieldUint64:  64,
			FieldString:  "d20b0c50263e33b57ea684260dd3f746888141a04dd4df808c475120bd234b17ebd1f668e4419d2329fc2bbf585ce37a4eec732d0f65913777be5acc33da58f2c262e76c4fa3c4e1ca3949193710b8a7ab31467434417d12e161e06fda2542123ac6c7f296a9a0cddebedca1d531459d7c855af340fd1961744d73f9db4f0e943f8799d2288f42abcfce144b237fe98ae6dd8993f64de0909c706892d6ba3e4cfc8890e39a74befca711de32b3f4a81085c710b8c52696e85ae0a646b766f4d1911daca43a58490eca987b55516dca81a27e6f43a9e4c9171032dae60abb597c616d6bd27c487b1375c0f3f8dee4774dfaea8e24dc2ce3e6aa465e30dd2b82bd9063773ee65aa86dbfd4c8503fa3bf991583b876f0dac417aa4fcda79dde30916eca16ceb2b8dcdca34ea977ac39802659702295b5efb5ca5ce91cf6a1930dd5308bab8436ce9ba555b20d22be1dfd2acb30328b6e1fa6f5f1c8e9baca46a2971e972983cdafa69d8e6db70e3fce7f6d4a486d6428addf4bde83b8acd0a711ee6af49ed6366d5847ce56e5718e13e2f2f99f1fd260331c561d34f5130f98e61b1ea18b020d7865fc2d38a3be1bc4e1316afcf27e609c5d8716859a3af655944e969fa6c9cba8d70fb66e4e4302c59d118048cc4807a64ba16a20d8d975d79b2636cb47b348c4685575ea13669d924d75c4a6c61bc5a8376f0c53591e7a25f54cfeb664bdd0e6e4716f7e6232f63f1ec4bb1b65f347ba1f66d5d3fb2e9336a4fd605651627c6802c7b5328b791d16387408c6981c836ea5432ffcf44adf253ef2f8af81a46f2617d463e6c9adda42544324b875a8ada57dbb1fa392ce9778787225809c1ca07021b646048d0fee5885124db54de03ab2263207df9f5455bc2e4f51b33c6659cc0023e77605a8104c57d7f20d457a87b6c028e31f5421b25de0f24813ce122ad5b5157a851d3b0b888ab0c4e2b5f12d581be4716edd93e2f701c8d05be791731dc8f0c00bc01020bd213803e4209fcc10d08128ed1a1ac98d38af818561a2141a7c4f046ca9283a74d8277f85231e5cc3dff9046e08379152f5d01445924f691bed702907ea9abbaf6343d400204cd81260794b5dc5ff72eaee255e40c1f70e018b11fe74599509e2bf4cd306339e03a839d393d7b5c3a28498c819e1dbbbb6c05b13112caff4f317fbd6d21324a0bdfbac8938dd7669b2874951617a02dad6400891ee9bb55f24e1879d4f90b369e1f9c8080c660ebcdac429c2e8302fb7b68b69f6a0083d1d2265cb5e71f43c4bc2e12a98dff892d33811510fa44ce6461626b409d7c4886f8eea1a01f9b779179943f79ab986cba75ccad88c08a7b94d7df77204b41708a925656a4106d721df4469b4e692c80705c35a69a300981d63e525392f",
			FieldBool:    true,
			FieldAny:     "any type of value", // Can be any value, like a string, int, etc.
			FieldError:   errors.New("example error"),
		}

		var output Output
		mapping.Into(input, &output)

		if output.FieldFloat32 != input.FieldFloat32 {
			b.Error("Error output FieldFloat32 is", output.FieldFloat32)
		}
		if output.FieldFloat64 != input.FieldFloat64 {
			b.Error("Error output FieldFloat64 is", output.FieldFloat64)
		}
		if output.FieldInt != input.FieldInt {
			b.Error("Error output FieldInt is", output.FieldInt)
		}
		if output.FieldInt8 != input.FieldInt8 {
			b.Error("Error output FieldInt8 is", output.FieldInt8)
		}
		if output.FieldInt16 != input.FieldInt16 {
			b.Error("Error output FieldInt16 is", output.FieldInt16)
		}
		if output.FieldInt32 != input.FieldInt32 {
			b.Error("Error output FieldInt32 is", output.FieldInt32)
		}
		if output.FieldInt64 != input.FieldInt64 {
			b.Error("Error output FieldInt64 is", output.FieldInt64)
		}
		if output.FieldUint != input.FieldUint {
			b.Error("Error output FieldUint is", output.FieldUint)
		}
		if output.FieldUint8 != input.FieldUint8 {
			b.Error("Error output FieldUint8 is", output.FieldUint8)
		}
		if output.FieldUint16 != input.FieldUint16 {
			b.Error("Error output FieldUint16 is", output.FieldUint16)
		}
		if output.FieldUint32 != input.FieldUint32 {
			b.Error("Error output FieldUint32 is", output.FieldUint32)
		}
		if output.FieldUint64 != input.FieldUint64 {
			b.Error("Error output FieldUint64 is", output.FieldUint64)
		}
		if output.FieldString != input.FieldString {
			b.Error("Error output FieldString is", output.FieldString)
		}
		if output.FieldBool != input.FieldBool {
			b.Error("Error output FieldBool is", output.FieldBool)
		}
		if output.FieldAny != input.FieldAny {
			b.Error("Error output FieldAny is", output.FieldAny)
		}
		if !errors.Is(output.FieldError, input.FieldError) {
			b.Error("Error output FieldError is", output.FieldError)
		}
	}
}

func Benchmark_Into_Small_Data(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := Input{
			FieldFloat32: 3.14,
			FieldFloat64: 3.14159265359,
			FieldInt:     42,
			FieldInt8:    -8,
			FieldInt16:   16,
			FieldInt32:   32,
			FieldInt64:   64,
			FieldUint:    42,
			FieldUint8:   8,
			FieldUint16:  16,
			FieldUint32:  32,
			FieldUint64:  64,
			FieldString:  "example",
			FieldBool:    true,
			FieldAny:     "any type of value", // Can be any value, like a string, int, etc.
			FieldError:   errors.New("example error"),
		}

		var output Output
		mapping.Into(input, &output)

		if output.FieldFloat32 != input.FieldFloat32 {
			b.Error("Error output FieldFloat32 is", output.FieldFloat32)
		}
		if output.FieldFloat64 != input.FieldFloat64 {
			b.Error("Error output FieldFloat64 is", output.FieldFloat64)
		}
		if output.FieldInt != input.FieldInt {
			b.Error("Error output FieldInt is", output.FieldInt)
		}
		if output.FieldInt8 != input.FieldInt8 {
			b.Error("Error output FieldInt8 is", output.FieldInt8)
		}
		if output.FieldInt16 != input.FieldInt16 {
			b.Error("Error output FieldInt16 is", output.FieldInt16)
		}
		if output.FieldInt32 != input.FieldInt32 {
			b.Error("Error output FieldInt32 is", output.FieldInt32)
		}
		if output.FieldInt64 != input.FieldInt64 {
			b.Error("Error output FieldInt64 is", output.FieldInt64)
		}
		if output.FieldUint != input.FieldUint {
			b.Error("Error output FieldUint is", output.FieldUint)
		}
		if output.FieldUint8 != input.FieldUint8 {
			b.Error("Error output FieldUint8 is", output.FieldUint8)
		}
		if output.FieldUint16 != input.FieldUint16 {
			b.Error("Error output FieldUint16 is", output.FieldUint16)
		}
		if output.FieldUint32 != input.FieldUint32 {
			b.Error("Error output FieldUint32 is", output.FieldUint32)
		}
		if output.FieldUint64 != input.FieldUint64 {
			b.Error("Error output FieldUint64 is", output.FieldUint64)
		}
		if output.FieldString != input.FieldString {
			b.Error("Error output FieldString is", output.FieldString)
		}
		if output.FieldBool != input.FieldBool {
			b.Error("Error output FieldBool is", output.FieldBool)
		}
		if output.FieldAny != input.FieldAny {
			b.Error("Error output FieldAny is", output.FieldAny)
		}
		if !errors.Is(output.FieldError, input.FieldError) {
			b.Error("Error output FieldError is", output.FieldError)
		}
	}
}

func Benchmark_MapStructure_Small_Data(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := Input{
			FieldFloat32: 3.14,
			FieldFloat64: 3.14159265359,
			FieldInt:     42,
			FieldInt8:    -8,
			FieldInt16:   16,
			FieldInt32:   32,
			FieldInt64:   64,
			FieldUint:    42,
			FieldUint8:   8,
			FieldUint16:  16,
			FieldUint32:  32,
			FieldUint64:  64,
			FieldString:  "example string",
			FieldBool:    true,
			FieldAny:     "any type of value", // Can be any value, like a string, int, etc.
			FieldError:   errors.New("example error"),
		}

		var output Output
		_ = mapstructure.Decode(input, &output)

		if output.FieldFloat32 != input.FieldFloat32 {
			b.Error("Error output FieldFloat32 is", output.FieldFloat32)
		}
		if output.FieldFloat64 != input.FieldFloat64 {
			b.Error("Error output FieldFloat64 is", output.FieldFloat64)
		}
		if output.FieldInt != input.FieldInt {
			b.Error("Error output FieldInt is", output.FieldInt)
		}
		if output.FieldInt8 != input.FieldInt8 {
			b.Error("Error output FieldInt8 is", output.FieldInt8)
		}
		if output.FieldInt16 != input.FieldInt16 {
			b.Error("Error output FieldInt16 is", output.FieldInt16)
		}
		if output.FieldInt32 != input.FieldInt32 {
			b.Error("Error output FieldInt32 is", output.FieldInt32)
		}
		if output.FieldInt64 != input.FieldInt64 {
			b.Error("Error output FieldInt64 is", output.FieldInt64)
		}
		if output.FieldUint != input.FieldUint {
			b.Error("Error output FieldUint is", output.FieldUint)
		}
		if output.FieldUint8 != input.FieldUint8 {
			b.Error("Error output FieldUint8 is", output.FieldUint8)
		}
		if output.FieldUint16 != input.FieldUint16 {
			b.Error("Error output FieldUint16 is", output.FieldUint16)
		}
		if output.FieldUint32 != input.FieldUint32 {
			b.Error("Error output FieldUint32 is", output.FieldUint32)
		}
		if output.FieldUint64 != input.FieldUint64 {
			b.Error("Error output FieldUint64 is", output.FieldUint64)
		}
		if output.FieldString != input.FieldString {
			b.Error("Error output FieldString is", output.FieldString)
		}
		if output.FieldBool != input.FieldBool {
			b.Error("Error output FieldBool is", output.FieldBool)
		}
		if output.FieldAny != input.FieldAny {
			b.Error("Error output FieldAny is", output.FieldAny)
		}
		if !errors.Is(output.FieldError, input.FieldError) {
			b.Error("Error output FieldError is", output.FieldError)
		}
	}
}
