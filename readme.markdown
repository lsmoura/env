# ENV

env is a simple package to retrieve environment variables with default values and convert them to
the desired type

To install: `go get github.com/lsmoura/env`

## API Reference

There are two sets of functions.

The first of functions return the environment value and a possible error. They always take two parameters:
a string with the environment variable name and another value with the default value if the environment
variable is not present.

* `Bool(string, bool) (bool, error)`
* `Int(string, int) (int, error)`
* `Int64(string, int64) (int64, error)`
* `Str(string, string) (string, error)`

The second set of functions takes a pointer to the variable where you want the value to be stored as a
third parameter and only returns an error object.

* `BoolValue(string, bool, *bool) error` 
* `IntValue(string, int, *int) error`
* `Int64Value(string, int64, *int64) error`
* `StrValue(string, string, *string) error`

  
### A word about booleans

Booleans can understand many different values. These values all result int "true": 
`"1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "Y", "y", "on", "ON", "On"` while
all of these result in "false":
`"0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "N", "n", "off", "OFF", "Off"`.

If the value is empty, the default value will be used, and if it's not any of those cases, an
error will be returned.

## Author

This package was written by [Sergio Moura](https://sergio.moura.ca)


## License

[Unlicense](https://opensource.org/licenses/Unlicense)
