package stringx_test

import (
	"testing"

	"github.com/innotechdevops/core/stringx"
)

func TestParseConnectionString(t *testing.T) {
	conn := "host=ftp.xyz.com:2022;username=abc;password=xyz"

	config := struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := stringx.ParseConnectionString(conn, &config)
	if err != nil || config.Host != "ftp.xyz.com:2022" || config.Username != "abc" || config.Password != "xyz" {
		t.Error("Parse error:", err)
	}
}

func TestTrimToLower(t *testing.T) {
	fixedRate := "Fixed Rate"

	actual := stringx.TrimToLower(fixedRate)

	if actual != "fixedrate" {
		t.Error("Error", actual)
	}
}

func TestIsNotNull(t *testing.T) {
	value := "test"
	if !stringx.IsNotNull(&value) {
		t.Errorf("Expected IsNotNull to return true for non-nil value")
	}

	var nilValue *int
	if stringx.IsNotNull(nilValue) {
		t.Errorf("Expected IsNotNull to return false for nil value")
	}
}

func TestNormalizeText(t *testing.T) {
	text := "Hello, World!  \nThis is a test.\t\n"
	normalized := stringx.NormalizeText(text)
	if normalized != "Hello, World! This is a test." {
		t.Errorf("Expected 'Hello, World! This is a test.', got '%s'", normalized)
	}
}
