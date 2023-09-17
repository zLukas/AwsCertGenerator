package tls


// https://github.com/wardviaene/golang-for-devops-course/blob/main/tls-demo/pkg

type SelfSignedCert struct{
	caAuthority CACert
	cert Cert
	privateKey pem.PrivateKey 
}

func (ssc *SelfSignedCert) GetCertCa(){
	return ssc.caAuthority
}

func (ssc *SelfSignedCert) GetPrivKey(){
	return ssc.privateKey
}


func (ssc *SelfSignedCert) GenerateKey(){

}