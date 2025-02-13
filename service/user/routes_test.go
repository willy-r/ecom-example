package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/willy-r/ecom-example/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &MockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "invalid",
			Password:  "password",
		}
		marshalledPayload, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register a user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "valid@email.com",
			Password:  "password",
		}
		marshalledPayload, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type MockUserStore struct{}

func (m *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *MockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *MockUserStore) CreateUser(user types.User) error {
	return nil
}
