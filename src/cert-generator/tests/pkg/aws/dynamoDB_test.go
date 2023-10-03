package tests

import (
	"testing"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/aws"
	"github.com/zLukas/CloudTools/src/cert-generator/tests/pkg/mocks"
)


func TestLogIntoDbSuccesful(t *testing.T){
	client := mocks.DbClientMock{LogInOk: true}

	sess, err := aws.LogIntoDb(&client)

	if sess == nil {
		t.Errorf("session should be nil, got %v", sess)
	}
	if err != nil {
		t.Errorf("error should be nil, got %s", err.Error())
	}
}


func TestLogIntoDbFail(t *testing.T){
	client := mocks.DbClientMock{LogInOk: false}

	sess, err := aws.LogIntoDb(&client)

	if sess != nil {
		t.Errorf("session should be nil, got %v", sess)
	}
	if err == nil {
		t.Errorf("error should be error type got nil")
	}
}