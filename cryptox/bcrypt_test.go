package cryptox_test

import (
	"fmt"
	"testing"

	"github.com/innotechdevops/core/cryptox"
)

func TestBcryptHash(t *testing.T) {
	hash, _ := cryptox.BcryptHash("atv.admin")
	fmt.Println(hash)
}

func TestBcryptCompare(t *testing.T) {
	hash := "$2a$14$bC1NJxyMLt5HAaSWKYyrAu37.np8WwW9DmaXsZqQXQhEaRuQIy9aG"
	result := cryptox.BcryptCompare(hash, "atv.admin")
	if !result {
		t.Error("Password not match")
	}
}
