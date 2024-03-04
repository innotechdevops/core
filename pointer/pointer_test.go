package pointer_test

import (
	"github.com/innotechdevops/core/pointer"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *T
	}
	tests := []testCase[string]{
		{name: "New string pointer", args: args[string]{value: "hello"}, want: pointer.New[string]("hello")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.New(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
