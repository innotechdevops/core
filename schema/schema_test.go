package schema_test

import (
	"testing"

	"github.com/innotechdevops/core/schema"
)

func TestJoinInt64(t *testing.T) {
	// Given
	list := []int64{1, 2, 3, 4, 5, 6}

	// When
	actual := schema.Join(list, ",")

	// Then
	if actual != "1,2,3,4,5,6" {
		t.Error("Error:", actual)
	}
}
