module github.com/Qualstor/aws-sdk-go-v2/service/dynamodb

go 1.15

require (
	github.com/Qualstor/aws-sdk-go-v2 v1.16.16
	github.com/Qualstor/aws-sdk-go-v2/internal/configsources v1.1.15
	github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 v2.4.9
	github.com/Qualstor/aws-sdk-go-v2/service/internal/accept-encoding v1.9.5
	github.com/Qualstor/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.6
	github.com/aws/smithy-go v1.11.3
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/Qualstor/aws-sdk-go-v2 => ../../

replace github.com/Qualstor/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/accept-encoding => ../../service/internal/accept-encoding/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../service/internal/endpoint-discovery/
