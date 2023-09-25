package tests

import (
	"testing"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

func TestRemoveEmptyStringEmptyString(t *testing.T) {
	empty_string := []string{""}
	result := tls.RemoveEmptyString(empty_string)

	if len(result) != 0 {
		t.Errorf("string array should have 0 lenght, go %v", len(result))
	}
}

func TestRemoveEmptyStringNotEmptyString(t *testing.T) {
	empty_string := []string{"asd", "asd", "wersfds"}
	result := tls.RemoveEmptyString(empty_string)

	if len(result) != 3 {
		t.Errorf("string array should have 3 lenght, go %v", len(result))
	}
}
