package cryptox_test

import (
	"fmt"
	"testing"

	"github.com/innotechdevops/core/cryptox"
)

func TestArgon2Hash(t *testing.T) {
	pass, _ := cryptox.Argon2Hash("atv.admin")
	fmt.Println(pass)
}

func TestArgon2Compare(t *testing.T) {
	passHash := "$argon2id$v=19$m=65536,t=1,p=4$DIP3B6ZMkBJxn8JUWHZPoA$2zDJkltjMS3nhmJSpPVSGwM5thm817eY8Rt6eFd/JOE"
	result, _ := cryptox.Argon2Compare("atv.admin", passHash)
	fmt.Println(result)
}
