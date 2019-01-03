package helper

import (
	"encoding/json"
	"net/http"
	"reflect"
)

// WriteOK ...
func WriteOK(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusOK)
	WriteJSON(w, message)
}

// WriteNotFound ...
func WriteNotFound(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusNotFound)
	WriteJSON(w, message)
}

// WriteBadRequest ...
func WriteBadRequest(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	WriteJSON(w, message)
}

// WriteJSON ...
func WriteJSON(w http.ResponseWriter, message interface{}) {
	json.NewEncoder(w).Encode(message)
}

// GetOrDefault return default value in case of value is zero value
func GetOrDefault(val, defaultVal interface{}) interface{} {
	ty := reflect.TypeOf(val)
	zeroVal := reflect.Zero(ty).Interface()

	if val == zeroVal {
		return defaultVal
	}
	return val
}
