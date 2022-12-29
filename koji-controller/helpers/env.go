package helpers

import (
    "os"
)

// retrieves an env variable by name if available
// returns the provided defaultValue otherwise
func GetEnv(name string, defaultValue string) string {
    if value, ok := os.LookupEnv(name); ok {
        return value
    }

    return defaultValue
}
