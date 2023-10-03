package aws

import( 
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
)
func LogIntoDb(client dBClient) (*session.Session, error) {
	sess , err:= client.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, fmt.Errorf("Cannot log into DB: %s", err.Error())
	}
	return sess, nil
}


func UploadItem(item TableRecord) (error){
  return nil
}