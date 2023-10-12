# AwsCertGenerator

Cert Generator is a hobby project that aim to automate TLS self signed certificates. Cert generator can be used locally as standalone binary, as well as aws deployment.

# AWS infrasructure
## Prerequisities
In order to create AWS infrastructure, use following policy: [aws-permission.json](doc/aws-permissions.json)  
## steps
1. build go binary from `main.go` in  `/src/cert-generator/` dir:
   `$  GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go`
2. zip pacakge to boostrap.zip  
3. place .zip file in `/terraform` dir.  
4. customise terraform variables if needed.
5. run terraform apply command.

