package env

import (
	"github.com/pkg/errors"
	"os"
	"strconv"
)

func BoolValue(key string, defaultValue bool, dst *bool) error {
	value := os.Getenv(key)
	if value == "" {
		*dst = defaultValue
		return nil
	}

	boolValue, err := parseBool(value)
	if err != nil {
		return errors.Wrap(err, "strconv.ParseBool")
	}

	*dst = boolValue
	return nil
}

func IntValue(key string, defaultValue int, dst *int) error {
	value := os.Getenv(key)
	if value == "" {
		*dst = defaultValue
		return nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return errors.Wrap(err, "strconv.Atoi")
	}

	*dst = intValue
	return nil
}

func Int64Value(key string, defaultValue int64, dst *int64) error {
	value := os.Getenv(key)
	if value == "" {
		*dst = defaultValue
		return nil
	}

	int64Value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errors.Wrap(err, "strconv.ParseInt")
	}

	*dst = int64Value
	return nil
}

func StrValue(key string, defaultValue string, dst *string) error {
	value := os.Getenv(key)
	if value == "" {
		*dst = defaultValue
		return nil
	}

	*dst = value
	return nil
}
