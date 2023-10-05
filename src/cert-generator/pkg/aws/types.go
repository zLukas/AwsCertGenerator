package aws

type CertItem struct {
	PrivateKey []byte `dynamodbav:"PrivateKey"`
	Cert       []byte `dynamodbav:"Cert"`
}

type TableRecord struct {
	CaCert       CertItem `dynamodbav:"Ca"`
	CeCert       CertItem `dynamodbav:"Ce"`
	Name         string   `dynamodbav:"Name"`
	CreationDate string   `dynamodbav:"CreationDate"`
}
