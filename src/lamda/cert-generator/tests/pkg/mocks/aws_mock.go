package mocks

import (
	"fmt"
 	"github.com/aws/aws-sdk-go/aws/session"
)


type DbClientMock struct {
	LogInOk bool
}

func (d *DbClientMock)NewSessionWithOptions(opts session.Options) (*session.Session, error){
	if d.LogInOk {
		return &session.Session{}, nil
	} 
	return nil, fmt.Errorf("Cannot log into DB:")
}