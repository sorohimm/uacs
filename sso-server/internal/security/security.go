package security

import (
	"crypto/sha256"
)

type Security struct{}

func (s *Security) HashPassword(password string) string {
	bytePwd := sha256.Sum256([]byte(password))
	return string(bytePwd[:])
}

func (s *Security) GenerateSalt() {

}
