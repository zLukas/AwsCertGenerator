package aws

import "github.com/aws/aws-sdk-go/aws/session"

func LogIntoDb(client dBClient) *session.Session {
	sess := client.Must(client.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}
