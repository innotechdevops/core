package common_test

import (
	"fmt"
	"github.com/innotechdevops/core/common"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestParseConnectionString(t *testing.T) {
	conn := "host=ftp.xyz.com:2022;username=abc;password=xyz"

	config := struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := common.ParseConnectionString(conn, &config)
	if err != nil || config.Host != "ftp.xyz.com:2022" || config.Username != "abc" || config.Password != "xyz" {
		t.Error("Parse error:", err)
	}
}

func TestTrimToLower(t *testing.T) {
	fixedRate := "Fixed Rate"

	actual := common.TrimToLower(fixedRate)

	if actual != "fixedrate" {
		t.Error("Error", actual)
	}
}

func TestIsFloat32(t *testing.T) {
	f := reflect.TypeOf(float64(3.14))

	actual := common.IsFloat(f)

	if !actual {
		t.Error("Is not float")
	}
}

func TestIsFloat64(t *testing.T) {
	f := reflect.TypeOf(float32(3.14))

	actual := common.IsFloat(f)

	if !actual {
		t.Error("Is not float")
	}
}

func TestF64ToString(t *testing.T) {
	f := 11.5200000000186265332323434343545
	expected := "11.520000000018626"

	actual := common.F64ToString(f)

	if actual != expected {
		t.Error("Convert error", actual)
	}
}

func TestF64ToStringDyn(t *testing.T) {
	f := 11.5200186265332323434343545325657697832
	expected := "11.520018626533233"

	actual := common.F64ToStringDyn(f)

	if actual != expected {
		t.Error("Convert error", actual)
	}
}

func TestParseNumEToNumber(t *testing.T) {
	// define a float with a large number of digits
	f := 123456789.12345678987876543245754255657

	// convert float to string using FormatFloat
	s := strconv.FormatFloat(f, 'f', -1, 64)

	// print the resulting string
	fmt.Println(s)
}

func TestTrimSpace(t *testing.T) {
	text := " Hello "

	actual := strings.TrimSpace(text)

	if actual != "Hello" {
		t.Error("Cannot trim space", actual)
	}
}

func TestIsDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Should true", args{"-1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.IsDigit(tt.args.s); got != tt.want {
				t.Errorf("IsDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
