package env

import (
	"os"
	"testing"
)

func TestBoolValue(t *testing.T) {
	tests := []struct {
		key      string
		strValue string
		expected bool
	}{
		{"TEST_BOOL_VALUE_TRUE_1", "1", true},
		{"TEST_BOOL_VALUE_FALSE_0", "0", false},
		{"TEST_BOOL_VALUE_TRUE_YES", "yes", true},
		{"TEST_BOOL_VALUE_FALSE_NO", "no", false},
		{"TEST_BOOL_VALUE_TRUE_ON", "on", true},
		{"TEST_BOOL_VALUE_FALSE_OFF", "off", false},
		{"TEST_BOOL_VALUE_TRUE_TRUE", "true", true},
		{"TEST_BOOL_VALUE_FALSE_FALSE", "false", false},
	}

	for _, test := range tests {
		if err := os.Setenv(test.key, test.strValue); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	t.Log("TestBoolValue")
	for _, test := range tests {
		var boolValue bool
		t.Logf("  Bool(\"%s\")\n", test.key)
		if err := BoolValue(test.key, !test.expected, &boolValue); err != nil {
			t.Errorf("BoolValue(%q) returned error: %v", test.key, err)
		}
		if boolValue != test.expected {
			t.Errorf("BoolValue(%q) = %v, want %v", test.key, boolValue, test.expected)
		}
	}

	t.Log("  Bool(\"TEST_BOOL_NOT_EXIST\")")
	if err := os.Unsetenv("TEST_BOOL_NOT_EXIST"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var bValue bool
	if err := BoolValue("TEST_BOOL_NOT_EXIST", true, &bValue); err != nil {
		t.Fatalf("Bool(\"TEST_BOOL_NOT_EXIST\", true) returned error: %v", err)
	}
	if bValue != true {
		t.Errorf("Bool(\"TEST_BOOL_NOT_EXIST\", true) = %v, want %v", bValue, true)
	}

	t.Log("  Bool(\"TEST_BOOL_ERROR\", false)")
	if err := os.Setenv("TEST_BOOL_ERROR", "error"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := BoolValue("TEST_BOOL_ERROR", false, &bValue); err == nil {
		t.Errorf("BoolValue(\"TEST_BOOL_ERROR\", false) returned no error, want error")
	}
}

func TestInt64Value(t *testing.T) {
	const envNameInt64 = "TEST_INT64_VALUE"
	const envNameInt64Invalid = "TEST_INT64_VALUE_INVALID"

	if err := os.Setenv(envNameInt64, "123456789"); err != nil {
		t.Fatalf("cannot set env %s", envNameInt64)
	}
	if err := os.Unsetenv(envNameInt64Invalid); err != nil {
		t.Fatalf("cannot unset env %s", envNameInt64Invalid)
	}

	t.Log("TestEnvGetInt64Value")
	t.Logf("  Int64Value(\"%s\")\n", envNameInt64)
	var int64Value int64
	err := Int64Value(envNameInt64, 0, &int64Value)
	if err != nil {
		t.Fatalf("Int64Value(\"%s\") failed: %s", envNameInt64, err)
	}
	if int64Value != 123456789 {
		t.Errorf("Int64Value(\"%s\") returned %d", envNameInt64, int64Value)
	}

	t.Logf("  Int64Value(\"%s\")\n", envNameInt64Invalid)
	var invalidValue int64
	err = Int64Value(envNameInt64Invalid, 0, &invalidValue)
	if err != nil {
		t.Fatalf("Int64Value(\"%s\") failed: %s", envNameInt64Invalid, err)
	}
	if invalidValue != 0 {
		t.Errorf("Int64Value(\"%s\") returned %d", envNameInt64Invalid, invalidValue)
	}

	t.Logf("  Int64Value(\"%s\")\n", envNameInt64Invalid)
	err = Int64Value(envNameInt64Invalid, 5, &invalidValue)
	if err != nil {
		t.Fatalf("Int64Value(\"%s\") failed: %s", envNameInt64Invalid, err)
	}
	if invalidValue != 5 {
		t.Errorf("Int64Value(\"%s\") returned %d", envNameInt64Invalid, invalidValue)
	}
}

func TestIntValue(t *testing.T) {
	const envNameInt = "TEST_INT_VALUE"
	const envNameIntInvalid = "TEST_INT_VALUE_INVALID"

	if err := os.Setenv(envNameInt, "123456789"); err != nil {
		t.Fatalf("cannot set env %s", envNameInt)
	}
	if err := os.Unsetenv(envNameIntInvalid); err != nil {
		t.Fatalf("cannot unset env %s", envNameIntInvalid)
	}

	t.Log("TestIntValue")
	t.Logf("  IntValue(\"%s\")\n", envNameInt)
	var intValue int
	err := IntValue(envNameInt, 0, &intValue)
	if err != nil {
		t.Fatalf("IntValue(\"%s\") failed: %s", envNameInt, err)
	}
	if intValue != 123456789 {
		t.Errorf("IntValue(\"%s\") returned %d", envNameInt, intValue)
	}

	t.Logf("  IntValue(\"%s\")\n", envNameIntInvalid)
	var invalidValue int
	err = IntValue(envNameIntInvalid, 0, &invalidValue)
	if err != nil {
		t.Fatalf("IntValue(\"%s\") failed: %s", envNameIntInvalid, err)
	}
	if invalidValue != 0 {
		t.Errorf("IntValue(\"%s\") returned %d", envNameIntInvalid, invalidValue)
	}

	t.Logf("  IntValue(\"%s\")\n", envNameIntInvalid)
	err = IntValue(envNameIntInvalid, 5, &invalidValue)
	if err != nil {
		t.Fatalf("IntValue(\"%s\") failed: %s", envNameIntInvalid, err)
	}
	if invalidValue != 5 {
		t.Errorf("IntValue(\"%s\") returned %d", envNameIntInvalid, invalidValue)
	}
}

func TestStrValue(t *testing.T) {
	const envNameStr = "TEST_STR_VALUE"
	const envNameStrInvalid = "TEST_STR_VALUE_INVALID"

	if err := os.Setenv(envNameStr, "test"); err != nil {
		t.Fatalf("cannot set env %s", envNameStr)
	}
	if err := os.Unsetenv(envNameStrInvalid); err != nil {
		t.Fatalf("cannot unset env %s", envNameStrInvalid)
	}

	var strValue string

	t.Log("TestStrValue")
	t.Logf("  StrValue(\"%s\")\n", envNameStr)
	if err := StrValue(envNameStr, "", &strValue); err != nil {
		t.Fatalf("StrValue(\"%s\") failed: %s", envNameStr, err)
	}
	if strValue != "test" {
		t.Errorf("StrValue(\"%s\") returned %s", envNameStr, strValue)
	}

	t.Logf("  StrValue(\"%s\")\n", envNameStrInvalid)
	if err := StrValue(envNameStrInvalid, "default", &strValue); err != nil {
		t.Fatalf("StrValue(\"%s\") failed: %s", envNameStrInvalid, err)
	}
	if strValue != "default" {
		t.Errorf("StrValue(\"%s\") returned %s", envNameStrInvalid, strValue)
	}
}
