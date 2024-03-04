package gojwt

import (
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt"
	"github.com/innotechdevops/timex"
	"github.com/pkg/errors"
	"strings"
	"time"
)

const TypeBearer = "Bearer"

type model[T any] struct {
	Data T
}

func TokenExpired(minute int) int64 {
	return time.Now().In(timex.GetTimeZone(timex.TimeZoneAsiaBangkok)).Add(time.Minute * time.Duration(minute)).Unix()
}

func GetString(token string, key string, secret string) string {
	if token != "" {
		tokenStr := strings.Replace(token, TypeBearer, "", -1)
		hmacSecret := []byte(secret)
		jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})

		if err != nil {
			return ""
		}

		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
			uid := claims[key]
			return uid.(string)
		}
	}

	return ""
}

func GetPayload[T any](token string, secret string) (T, error) {
	if token != "" {
		tokenStr := strings.Replace(token, "Bearer ", "", -1)
		return getPayloadFromToken[T](tokenStr, secret)
	}
	payload := model[T]{}
	return payload.Data, errors.New("token is empty")
}

func getPayloadFromToken[T any](tokenStr string, secret string) (T, error) {
	payload := model[T]{}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return payload.Data, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if p, err2 := json.Marshal(claims); err2 == nil {
			if err2 = json.Unmarshal(p, &payload.Data); err2 == nil {
				return payload.Data, nil
			}
		}
	}
	return payload.Data, errors.New("token invalid")
}

func Verify(tokenStr string, secret string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err == nil && token.Valid {
		return nil
	}
	return err
}

func NewClaims() jwt.MapClaims {
	// Create token
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := jwtToken.Claims.(jwt.MapClaims)
	return claims
}

func Generate(claims jwt.MapClaims, secret string) string {
	// Set claims
	payload := jwt.MapClaims{}
	for k, v := range claims {
		payload[k] = v
	}

	// Generate encoded token and send it as response.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return tokenStr
}
