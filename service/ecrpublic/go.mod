module github.com/Qualstor/aws-sdk-go-v2/service/ecrpublic

go 1.15

require (
	github.com/Qualstor/aws-sdk-go-v2 v1.16.16
	github.com/Qualstor/aws-sdk-go-v2/internal/configsources v1.1.15
	github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 v2.4.9
	github.com/aws/smithy-go v1.11.3
)

replace github.com/Qualstor/aws-sdk-go-v2 => ../../

replace github.com/Qualstor/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
