package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash to be not empty")
	}

	if hash == password {
		t.Errorf("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePasswords(hash, password) {
		t.Errorf("expected passwords to match")
	}

	if ComparePasswords(hash, "wrongpassword") {
		t.Errorf("expected passwords to not match")
	}
}
