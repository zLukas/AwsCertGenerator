package tests

import (
	"reflect"
	"testing"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

var (
	testPem pemMock = pemMock{
		encodePass: true,
		decodePass: true,
	}
	testRsa rsaMock = rsaMock{
		generateKeyPass: true,
	}

	testX509 x509Mock = x509Mock{
		createcertificatesPass: true,
		parseCertificatePass:   true,
	}
)

func SetUp(t *testing.T) {
	tls.Ix509 = &testX509
	tls.Ipem = &testPem
	tls.Irsa = &testRsa
}

func TearDown() {
	testPem.decodePass = true
	testPem.encodePass = true
	testRsa.generateKeyPass = true
	testRsa.generateKeyPass = true
	testX509.createcertificatesPass = true
	testX509.parseCertificatePass = true
}

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

func TestPemToX509DecodeFail(t *testing.T) {
	SetUp(t)
	defer TearDown()
	testPem.decodePass = false

	cert, err := tls.PemToX509(nil)
	if cert != nil {
		t.Errorf("cert should be nil, got %v", cert)
	}
	if err.Error() != "failed to parse certificate PEM" {
		t.Errorf(" error message should be \"failed to parse certificate PEM\", got %v", err)
	}

}
func TestPemToX509DecodePass(t *testing.T) {
	SetUp(t)
	defer TearDown()
	cert, err := tls.PemToX509(nil)

	if reflect.TypeOf(cert).String() != string("*x509.Certificate") {
		t.Errorf("cert should be type *x509.Certificate, got %T", cert)
	}
	if err != nil {
		t.Errorf(" error should be nil, got %v", err)
	}
}

func TestPemToX509ParseFail(t *testing.T) {
	SetUp(t)
	defer TearDown()
	testX509.parseCertificatePass = false
	cert, err := tls.PemToX509(nil)

	if cert != nil {
		t.Errorf("cert should be nil, got %v", cert)
	}
	if err.Error() != "FAIL" {
		t.Errorf(" error message should be \"FAIL\", got %v", err)
	}
}

func TestPemToX509ParsePass(t *testing.T) {
	SetUp(t)
	defer TearDown()
	cert, err := tls.PemToX509(nil)

	if cert == nil {
		t.Errorf("cert should be \"*x509.Certificate\" type, got %v", cert)
	}
	if err != nil {
		t.Errorf(" error message should be nil, got %v", err)
	}
}

func TestCreateCACertBytesNullTemplate(t *testing.T) {
	SetUp(t)
	defer TearDown()

	key, cert, err := tls.CreateCACertBytes(nil)
	if key != nil {
		t.Errorf(" key should be nil, got %v", err)
	}
	if cert != nil {
		t.Errorf(" cert should be nil, got %v", err)
	}
	if err == nil {
		t.Errorf(" error should be %s got nil", "tempalte must contain CommonName Field")
	} else if err.Error() != "template is nil" {
		t.Errorf(" error should be \"template is nil\" got %s", err)
	}
}

func TestCreateCACertBytesNullNoCommonName(t *testing.T) {
	SetUp(t)
	defer TearDown()
	ca := &tls.CACert{}
	key, cert, err := tls.CreateCACertBytes(ca)
	if key != nil {
		t.Errorf(" key should be <nil>, got <%v>", err)
	}
	if cert != nil {
		t.Errorf(" cert should be <nil>, got <%v>", err)
	}
	if err == nil {
		t.Errorf(" error should be %s got nil", "tempalte must contain CommonName Field")
	} else if err.Error() != "template must contain CommonName Field" {
		t.Errorf(" error should be <template must contain CommonName Field> got <%s>", err)
	}
}
