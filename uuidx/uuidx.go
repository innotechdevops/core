package uuidx

import (
	"github.com/google/uuid"
	"github.com/innotechdevops/core/pointer"
)

func NewID() string {
	return pointer.Deref(NewIDPtr(), "")
}

func NewIDPtr() *string {
	uid, err := uuid.NewV7()
	if err != nil {
		return nil
	}
	return pointer.New(uid.String())
}

func IsValid(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
