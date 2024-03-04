package gouuid

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		uid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Is valid uuid", args: args{uid: "018e0a47-56c9-75dc-9cd2-2cc36a042bf0"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.uid); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewV7(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "New uuid v7", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewV7(); got == tt.want {
				t.Errorf("NewV7() = %v, want %v", got, tt.want)
			}
		})
	}
}
