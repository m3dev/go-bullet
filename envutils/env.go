package envutils

import (
	"fmt"
	"os"
)

func MustGetenv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("failed to get env '%s'", key)
	}
	return val, nil
}
