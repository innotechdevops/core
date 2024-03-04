package gojwt_test

import (
	"github.com/golang-jwt/jwt"
	"github.com/innotechdevops/core/gojwt"
	"testing"
)

type jwtPayload struct {
	Sub string `json:"sub"`
	Iss string `json:"iss"`
	Exp int64  `json:"exp"`
}

func (j jwtPayload) ToMapClaims() jwt.MapClaims {
	claims := gojwt.NewClaims()
	claims["sub"] = j.Sub
	claims["iss"] = j.Iss
	claims["exp"] = j.Exp
	return claims
}

func TestTokenExpired(t *testing.T) {
	// Given
	expiredMinute := 1

	// When
	actual := gojwt.TokenExpired(expiredMinute)

	// Then
	if actual < 0 {
		t.Error("Error", actual)
	}
}

func TestGenerate(t *testing.T) {
	payload := jwtPayload{
		Sub: "696969",
		Iss: "innotech",
		Exp: 9999999999,
	}

	type args struct {
		claims jwt.MapClaims
		secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Generate JWT",
			args: args{claims: payload.ToMapClaims(), secret: "test"},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTksImlzcyI6Imlubm90ZWNoIiwic3ViIjoiNjk2OTY5In0.8C9LAGqnpwqWwH8mrqrJJEICdl1ClHVTXk8PyQf6Z2c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gojwt.Generate(tt.args.claims, tt.args.secret); got != tt.want {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPayload(t *testing.T) {
	secret := "test"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTksImlzcyI6Imlubm90ZWNoIiwic3ViIjoiNjk2OTY5In0.8C9LAGqnpwqWwH8mrqrJJEICdl1ClHVTXk8PyQf6Z2c"
	payload := jwtPayload{
		Sub: "696969",
		Iss: "innotech",
		Exp: 9999999999,
	}

	got, _ := gojwt.GetPayload[jwtPayload](token, secret)

	if got.Sub != payload.Sub || got.Iss != payload.Iss {
		t.Errorf("GetPayload() = %v, want %v", got, payload)
	}
}

func TestGetString(t *testing.T) {
	secret := "test"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTksImlzcyI6Imlubm90ZWNoIiwic3ViIjoiNjk2OTY5In0.8C9LAGqnpwqWwH8mrqrJJEICdl1ClHVTXk8PyQf6Z2c"

	type args struct {
		token  string
		key    string
		secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Get string by key", args: args{token: token, key: "iss", secret: secret}, want: "innotech"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gojwt.GetString(tt.args.token, tt.args.key, tt.args.secret); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	secret := "test"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTksImlzcyI6Imlubm90ZWNoIiwic3ViIjoiNjk2OTY5In0.8C9LAGqnpwqWwH8mrqrJJEICdl1ClHVTXk8PyQf6Z2c"

	type args struct {
		tokenStr string
		secret   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Verify JWT", args: args{tokenStr: token, secret: secret}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := gojwt.Verify(tt.args.tokenStr, tt.args.secret); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
