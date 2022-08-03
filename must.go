package must

import (
	"errors"
	"fmt"
	"os"
)

var ErrEnvironmentVariableIsNotSet = errors.New("environment variable is not set")

func Getenv(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		panic(fmt.Errorf("%s: %v", key, ErrEnvironmentVariableIsNotSet))
	}

	return value
}
