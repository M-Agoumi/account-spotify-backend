package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// JSONResponse sets the header and writes a JSON response
func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}

// JSONError writes an error message in JSON format
func JSONError(w http.ResponseWriter, statusCode int, message string) {
	JSONResponse(w, statusCode, map[string]string{"error": message})
}

// DecodeJSONBody decodes the JSON body of a request into the given struct, returning a generic error message on failure.
func DecodeJSONBody(r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // Ensure that only the defined fields in the struct are accepted

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("invalid request payload")
		case errors.As(err, &unmarshalTypeError):
			return fmt.Errorf("invalid request payload")
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			return fmt.Errorf("invalid request payload")
		case err.Error() == "http: request body too large":
			return fmt.Errorf("invalid request payload")
		default:
			return fmt.Errorf("invalid request payload")
		}
	}
	return nil
}

type ValidationError struct {
	Errors []string
}

func (v *ValidationError) Error() string {
	return strings.Join(v.Errors, ", ")
}

func NewValidationError() *ValidationError {
	return &ValidationError{Errors: []string{}}
}

func (v *ValidationError) Add(err string) {
	v.Errors = append(v.Errors, err)
}

func (v *ValidationError) HasErrors() bool {
	return len(v.Errors) > 0
}
