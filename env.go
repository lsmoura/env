package env

import (
	"github.com/pkg/errors"
	"os"
	"strconv"
)

func Bool(key string, defaultValue bool) (bool, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	boolValue, err := parseBool(value)
	if err != nil {
		return defaultValue, errors.Wrap(err, "strconv.ParseBool")
	}

	return boolValue, nil
}

func Int(key string, defaultValue int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue, errors.Wrap(err, "strconv.Atoi")
	}

	return intValue, nil
}

func Int64(key string, defaultValue int64) (int64, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	int64Value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue, errors.Wrap(err, "strconv.ParseInt")
	}

	return int64Value, nil
}

func Str(key string, defaultValue string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	return value, nil
}
