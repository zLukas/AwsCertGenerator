# AWS infrastructure

<img width="871" alt="aws infra" src="https://github.com/zLukas/AwsCertGenerator/doc/aws_infra.png">

## directory structure
* modules
  * dynamodb - aws dynamodb table generic module
  * lambda - aws lambda generic module
* main.tf - CertGeneraot infrastructure, using above modules
* test-event.json - sample lambda event to test lambda deployment
