package envutils

import (
	"fmt"
	"os"
)

func MustEnv(s string) string {
	result := os.Getenv(s)
	if result == "" {
		panic(fmt.Sprintf("No env variable for %s", s))
	}
	return result
}
