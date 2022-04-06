package utils

import (
	"os"
	"strings"
)

func IsDevEnvironment() bool {
	return strings.Contains(strings.ToUpper(os.Getenv("ENVIRONMENT")), "DEV")
}