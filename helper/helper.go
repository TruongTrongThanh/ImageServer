package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// WriteError ...
func WriteError(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	WriteJSON(w, GetOrDefault(message, "Internal Server Error"))
}

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
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		panic(err)
	}
}

// IsMultipartFormData ...
func IsMultipartFormData(r http.Request) bool {
	contentType := r.Header.Get("content-type")
	if strings.Contains(contentType, "multipart/form-data") {
		return true
	}
	return false
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

// GetHTTP ...
func GetHTTP() string {
	isSSL, err := strconv.ParseBool(os.Getenv("ssl"))
	if err != nil {
		log.Println("ssl setting in goenv is wrong")
		return "http"
	}
	if isSSL {
		return "https"
	} else {
		return "http"
	}
}

// GetFileServerHost ...
func GetFileServerHost() string {
	return fmt.Sprintf(
		"%s://%s:%s%s",
		GetHTTP(),
		os.Getenv("hostname"),
		os.Getenv("port"),
		os.Getenv("FileServerPath"))
}

// StringSliceContains ...
func StringSliceContains(slice []string, item string) bool {
	for _, val := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// SimpleHash create subdirectories name for storing image
func SimpleHash(i int) string {
	return fmt.Sprint(i % 50)
}

// SplitFilename ...
func SplitFilename(s string) (string, string) {
	return GetBasename(s), GetExt(s)
}

// GetBasename ...
func GetBasename(s string) string {
	slice := strings.Split(s, ".")
	if len(slice) > 0 {
		return strings.Join(slice[:len(slice)-1], "")
	}
	return slice[0]
}

// GetExt ...
func GetExt(s string) string {
	slice := strings.Split(s, ".")
	if len(slice) > 0 {
		return slice[len(slice)-1]
	} else {
		return ""
	}
}
