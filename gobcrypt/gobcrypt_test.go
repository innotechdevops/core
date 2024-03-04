package gobcrypt_test

import (
	"github.com/innotechdevops/core/gobcrypt"
	"testing"
)

func TestComparePassword(t *testing.T) {
	type args struct {
		hashedPwd string
		plainPwd  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Compare password", args: args{hashedPwd: "$2a$10$9qA7kNGX0oDcJtE1ugiTaeKRSPsAvlXhRruPDxXHabJjTNWyOZJg2", plainPwd: "69696969"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gobcrypt.ComparePassword(tt.args.hashedPwd, tt.args.plainPwd); got != tt.want {
				t.Errorf("ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Hash password", args: args{password: "69696969"}, want: "", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gobcrypt.HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("HashPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}
