package database_test

import (
	"my-joke-book/database"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type HashPasswordTest struct {
	in   string
	want bool
}

var HashPasswordTests = []HashPasswordTest{
	{"password", true},
	{"", true},
	{"p", true},
}

func TestHashPassword(t *testing.T) {
	for _, tt := range HashPasswordTests {
		hash, err := database.HashPassword(tt.in)
		if tt.want {
			if err != nil {
				t.Errorf("export error")
			}
			if hash_err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(tt.in)); hash_err != nil {
				t.Errorf("Hashed text is not match")
			}
		} else {
			if err == nil {
				t.Errorf("Error string is not error")
			}
		}
	}
}
