package generate

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func Password(length int) (string, error) {
	if length < 10 || length > 16 {
		return "", errors.New("password length must be between 10 and 16 characters")
	}

	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "@#&!$%^*()-_=+[]{}<>?"

	all := lower + upper + digits + special

	password := make([]rune, 0, length)

	password = append(password, RandomChar(lower))
	password = append(password, RandomChar(upper))
	password = append(password, RandomChar(digits))
	password = append(password, RandomChar(special))

	for len(password) < length {
		password = append(password, RandomChar(all))
	}

	ShuffleRunes(password)

	return string(password), nil
}

func RandomChar(charset string) rune {
	max := big.NewInt(int64(len(charset)))
	n, _ := rand.Int(rand.Reader, max)
	return rune(charset[n.Int64()])
}

func ShuffleRunes(runes []rune) {
	for i := len(runes) - 1; i > 0; i-- {
		j := SecureRandInt(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func SecureRandInt(n int) int {
	max := big.NewInt(int64(n))
	num, _ := rand.Int(rand.Reader, max)
	return int(num.Int64())
}
