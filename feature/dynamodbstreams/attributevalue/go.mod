module github.com/Qualstor/aws-sdk-go-v2/feature/dynamodbstreams/attributevalue

go 1.15

require (
	github.com/Qualstor/aws-sdk-go-v2 v1.16.5
	github.com/Qualstor/aws-sdk-go-v2/service/dynamodb v1.15.6
	github.com/Qualstor/aws-sdk-go-v2/service/dynamodbstreams v1.13.6
	github.com/aws/smithy-go v1.11.3
	github.com/google/go-cmp v0.5.8
)

replace github.com/Qualstor/aws-sdk-go-v2 => ../../../

replace github.com/Qualstor/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/Qualstor/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/Qualstor/aws-sdk-go-v2/service/dynamodbstreams => ../../../service/dynamodbstreams/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/
