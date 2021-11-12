package env

import "github.com/pkg/errors"

func parseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "Y", "y", "on", "ON", "On":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "N", "n", "off", "OFF", "Off":
		return false, nil
	}
	return false, errors.New("SyntaxError parsing bool")
}
