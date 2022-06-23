module github.com/Qualstor/aws-sdk-go-v2/example/service/dynamodb/listItems

go 1.15

require (
	github.com/Qualstor/aws-sdk-go-v2 v1.16.14
	github.com/Qualstor/aws-sdk-go-v2/config v1.15.12
	github.com/Qualstor/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.9.4
	github.com/Qualstor/aws-sdk-go-v2/service/dynamodb v1.15.7
)

replace github.com/Qualstor/aws-sdk-go-v2 => ../../../../

replace github.com/Qualstor/aws-sdk-go-v2/config => ../../../../config/

replace github.com/Qualstor/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/Qualstor/aws-sdk-go-v2/feature/dynamodb/attributevalue => ../../../../feature/dynamodb/attributevalue/

replace github.com/Qualstor/aws-sdk-go-v2/feature/ec2/imds => ../../../../feature/ec2/imds/

replace github.com/Qualstor/aws-sdk-go-v2/internal/configsources => ../../../../internal/configsources/

replace github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 => ../../../../internal/endpoints/v2/

replace github.com/Qualstor/aws-sdk-go-v2/internal/ini => ../../../../internal/ini/

replace github.com/Qualstor/aws-sdk-go-v2/service/dynamodb => ../../../../service/dynamodb/

replace github.com/Qualstor/aws-sdk-go-v2/service/dynamodbstreams => ../../../../service/dynamodbstreams/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/accept-encoding => ../../../../service/internal/accept-encoding/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../../../service/internal/endpoint-discovery/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace github.com/Qualstor/aws-sdk-go-v2/service/sso => ../../../../service/sso/

replace github.com/Qualstor/aws-sdk-go-v2/service/sts => ../../../../service/sts/
