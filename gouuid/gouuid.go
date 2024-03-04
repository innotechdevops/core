package gouuid

import "github.com/google/uuid"

func IsValid(uid string) bool {
	_, err := uuid.Parse(uid)
	return err == nil
}

func NewV7() string {
	v, err := uuid.NewV7()
	if err != nil {
		return ""
	}
	return v.String()
}
