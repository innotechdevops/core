package finder_test

import (
	"github.com/innotechdevops/core/finder"
	"testing"
)

func TestMatch(t *testing.T) {
	m := map[string]bool{
		"asc":  true,
		"desc": true,
	}
	actual := finder.Match(m, "ASC")

	if !actual {
		t.Errorf("Actual %v, Expected: true", actual)
	}
}
