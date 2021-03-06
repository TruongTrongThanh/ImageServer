package goenv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LoadEnv ...
func LoadEnv() {
	bytes, readErr := ioutil.ReadFile("env.json")
	if readErr != nil {
		panic(readErr)
	}
	env := make(map[string]interface{})
	json.Unmarshal(bytes, &env)
	for key, value := range env {
		os.Setenv(key, fmt.Sprint(value))
	}
}
