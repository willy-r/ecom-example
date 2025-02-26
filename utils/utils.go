package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/willy-r/ecom-example/types"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("request body is required")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func PermissionDenied(w http.ResponseWriter) {
	WriteError(w, http.StatusUnauthorized, fmt.Errorf("permission denied"))
}

func ScanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
	if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt); err != nil {
		return nil, err
	}
	return u, nil
}

func ScanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	p := new(types.Product)
	if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Image, &p.Price, &p.Quantity, &p.CreatedAt); err != nil {
		return nil, err
	}
	return p, nil
}
