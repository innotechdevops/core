package generate_test

import (
	"testing"

	"github.com/innotechdevops/core/generate"
)

func TestPassword(t *testing.T) {
	pass, err := generate.Password(12)
	if err != nil && len(pass) == 0 {
		t.Error(err)
	}
}
