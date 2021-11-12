package env

import (
	"os"
	"testing"
)

func TestBool(t *testing.T) {
	tests := []struct {
		key      string
		strValue string
		expected bool
	}{
		{"TEST_BOOL_TRUE_1", "1", true},
		{"TEST_BOOL_FALSE_0", "0", false},
		{"TEST_BOOL_TRUE_YES", "yes", true},
		{"TEST_BOOL_FALSE_NO", "no", false},
		{"TEST_BOOL_TRUE_ON", "on", true},
		{"TEST_BOOL_FALSE_OFF", "off", false},
		{"TEST_BOOL_TRUE_TRUE", "true", true},
		{"TEST_BOOL_FALSE_FALSE", "false", false},
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
		boolValue, err := Bool(test.key, !test.expected)
		if err != nil {
			t.Fatalf("Bool(%q) returned error: %v", test.key, err)
		}
		if boolValue != test.expected {
			t.Errorf("Bool(%q) = %v, want %v", test.key, boolValue, test.expected)
		}
	}

	t.Log("  Bool(\"TEST_BOOL_NOT_EXIST\")")
	if err := os.Unsetenv("TEST_BOOL_NOT_EXIST"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bValue, err := Bool("TEST_BOOL_NOT_EXIST", true)
	if err != nil {
		t.Fatalf("Bool(\"TEST_BOOL_NOT_EXIST\", true) returned error: %v", err)
	}
	if bValue != true {
		t.Errorf("Bool(\"TEST_BOOL_NOT_EXIST\", true) = %v, want %v", bValue, true)
	}

	t.Log("  Bool(\"TEST_BOOL_ERROR\", false)")
	if err := os.Setenv("TEST_BOOL_ERROR", "error"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bValue, err = Bool("TEST_BOOL_ERROR", false)
	if err == nil {
		t.Errorf("Bool(\"TEST_BOOL_ERROR\", false) returned no error, want error")
	}
}

func TestInt64(t *testing.T) {
	const envNameInt64 = "TEST_INT64"
	const envNameInt64Invalid = "TEST_INT64_INVALID"

	if err := os.Setenv(envNameInt64, "123456789"); err != nil {
		t.Fatalf("cannot set env %s", envNameInt64)
	}
	if err := os.Unsetenv(envNameInt64Invalid); err != nil {
		t.Fatalf("cannot unset env %s", envNameInt64Invalid)
	}

	t.Log("TestEnvGetInt64")
	t.Logf("  Int64(\"%s\")\n", envNameInt64)
	int64Value, err := Int64(envNameInt64, 0)
	if err != nil {
		t.Fatalf("Int64(\"%s\") failed: %s", envNameInt64, err)
	}
	if int64Value != 123456789 {
		t.Errorf("Int64(\"%s\") returned %d", envNameInt64, int64Value)
	}

	t.Logf("  Int64(\"%s\")\n", envNameInt64Invalid)
	invalidValue, err := Int64(envNameInt64Invalid, 0)
	if err != nil {
		t.Fatalf("Int64(\"%s\") failed: %s", envNameInt64Invalid, err)
	}
	if invalidValue != 0 {
		t.Errorf("Int64(\"%s\") returned %d", envNameInt64Invalid, invalidValue)
	}

	t.Logf("  Int64(\"%s\")\n", envNameInt64Invalid)
	invalidValue, err = Int64(envNameInt64Invalid, 5)
	if err != nil {
		t.Fatalf("Int64(\"%s\") failed: %s", envNameInt64Invalid, err)
	}
	if invalidValue != 5 {
		t.Errorf("Int64(\"%s\") returned %d", envNameInt64Invalid, invalidValue)
	}
}

func TestInt(t *testing.T) {
	const envNameInt = "TEST_INT"
	const envNameIntInvalid = "TEST_INT_INVALID"

	if err := os.Setenv(envNameInt, "123456789"); err != nil {
		t.Fatalf("cannot set env %s", envNameInt)
	}
	if err := os.Unsetenv(envNameIntInvalid); err != nil {
		t.Fatalf("cannot unset env %s", envNameIntInvalid)
	}

	t.Log("TestInt")
	t.Logf("  Int(\"%s\")\n", envNameInt)
	intValue, err := Int(envNameInt, 0)
	if err != nil {
		t.Fatalf("Int(\"%s\") failed: %s", envNameInt, err)
	}
	if intValue != 123456789 {
		t.Errorf("Int(\"%s\") returned %d", envNameInt, intValue)
	}

	t.Logf("  Int(\"%s\")\n", envNameIntInvalid)
	invalidValue, err := Int(envNameIntInvalid, 0)
	if err != nil {
		t.Fatalf("Int(\"%s\") failed: %s", envNameIntInvalid, err)
	}
	if invalidValue != 0 {
		t.Errorf("Int(\"%s\") returned %d", envNameIntInvalid, invalidValue)
	}

	t.Logf("  Int(\"%s\")\n", envNameIntInvalid)
	invalidValue, err = Int(envNameIntInvalid, 5)
	if err != nil {
		t.Fatalf("Int(\"%s\") failed: %s", envNameIntInvalid, err)
	}
	if invalidValue != 5 {
		t.Errorf("Int(\"%s\") returned %d", envNameIntInvalid, invalidValue)
	}
}

func TestStr(t *testing.T) {
	const envNameStr = "TEST_STR"
	const envNameStrInvalid = "TEST_STR_INVALID"

	if err := os.Setenv(envNameStr, "test"); err != nil {
		t.Fatalf("cannot set env %s", envNameStr)
	}
	if err := os.Unsetenv(envNameStrInvalid); err != nil {
		t.Fatalf("cannot unset env %s", envNameStrInvalid)
	}

	t.Log("TestEnvGetStr")
	t.Logf("  Str(\"%s\")\n", envNameStr)
	strValue, err := Str(envNameStr, "")
	if err != nil {
		t.Fatalf("Str(\"%s\") failed: %s", envNameStr, err)
	}
	if strValue != "test" {
		t.Errorf("Str(\"%s\") returned %s", envNameStr, strValue)
	}

	t.Logf("  Str(\"%s\")\n", envNameStrInvalid)
	invalidValue, err := Str(envNameStrInvalid, "default")
	if err != nil {
		t.Fatalf("Str(\"%s\") failed: %s", envNameStrInvalid, err)
	}
	if invalidValue != "default" {
		t.Errorf("Str(\"%s\") returned %s", envNameStrInvalid, invalidValue)
	}
}
