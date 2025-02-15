package auth

import "testing"

func TestCreateJwt(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJwt(secret, 1)
	if err != nil {
		t.Errorf("error creating jwt: %v", err)
	}

	if token == "" {
		t.Errorf("expected token to be not empty")
	}
}
