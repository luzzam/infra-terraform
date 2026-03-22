package helpers

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// GetEnvironmentVariable retrieves an environment variable and returns its value.
// If the variable is not set, it returns an empty string.
func GetEnvironmentVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if!exists {
		return ""
	}
	return value
}

// GetEnvironmentVariableInt retrieves an environment variable, attempts to parse it as an integer,
// and returns the result. If the variable is not set or the parsing fails, it returns an error.
func GetEnvironmentVariableInt(key string) (int, error) {
	value := GetEnvironmentVariable(key)
	if value == "" {
		return 0, fmt.Errorf("environment variable %s not set", key)
	}
	parsed, err := strconv.Atoi(value)
	if err!= nil {
		return 0, fmt.Errorf("failed to parse environment variable %s as int: %w", key, err)
	}
	return parsed, nil
}

// GetGoVersion retrieves the Go version and returns its value.
func GetGoVersion() string {
	return runtime.Version()
}

// GetOSName retrieves the OS name and returns its value.
func GetOSName() string {
	return runtime.GOOS
}

// GetPlatform retrieves the platform name and returns its value.
func GetPlatform() string {
	return runtime.GOARCH
}

// IsWindows checks if the current platform is Windows.
func IsWindows() bool {
	return strings.ToLower(GetOSName()) == "windows"
}