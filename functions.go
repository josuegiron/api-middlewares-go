package apirest

import (
	"net/http"
)

// Check doc...
func Check(err error) bool {
	if err != nil {
		return true
	}
	return false
}

type contextKey int

// Vars doc ...
func Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(0); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

var (
	// Username doc ...
	Username = "test"
	// Password doc ...
	Password = "test"
)

func validate(username, password string) bool {
	if username == Username && password == Password {
		return true
	}
	return false
}
