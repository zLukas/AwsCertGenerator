package tests

import (
	"strings"
	"testing"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/input"
)

func TestParse_input_default_file(t *testing.T) {
	var test_cfg = input.Config{}
	err := test_cfg.ParseInput()
	if err != nil {
		t.Errorf("expected nil got error %s", err)
	}

}

func TestParse_input_no_input_file(t *testing.T) {
	var test_cfg = input.Config{CfgFilePath: "tls_no_exist.yaml"}
	err := test_cfg.ParseInput()
	if err == nil {
		t.Error("expected error, got nil")
		return
	}
	if strings.Contains(err.Error(), "Error while reading config file") == false {
		t.Errorf("expected :\"Error while reading config file...\" got %s", err)
	}
}
