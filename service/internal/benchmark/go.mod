module github.com/Qualstor/aws-sdk-go-v2/service/internal/benchmark

go 1.15

require (
	github.com/Qualstor/aws-sdk-go-v2 v1.16.16
	github.com/Qualstor/aws-sdk-go-v2/service/dynamodb v1.15.7
	github.com/Qualstor/aws-sdk-go-v2/service/lexruntimeservice v1.12.6
	github.com/Qualstor/aws-sdk-go-v2/service/s3 v1.26.14
	github.com/Qualstor/aws-sdk-go-v2/service/schemas v1.14.6
	github.com/aws/aws-sdk-go v1.44.28
	github.com/aws/smithy-go v1.11.3
)

replace github.com/Qualstor/aws-sdk-go-v2 => ../../../

replace github.com/Qualstor/aws-sdk-go-v2/aws/protocol/eventstream => ../../../aws/protocol/eventstream/

replace github.com/Qualstor/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/Qualstor/aws-sdk-go-v2/internal/v4a => ../../../internal/v4a/

replace github.com/Qualstor/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/checksum => ../../../service/internal/checksum/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/s3shared => ../../../service/internal/s3shared/

replace github.com/Qualstor/aws-sdk-go-v2/service/lexruntimeservice => ../../../service/lexruntimeservice/

replace github.com/Qualstor/aws-sdk-go-v2/service/s3 => ../../../service/s3/

replace github.com/Qualstor/aws-sdk-go-v2/service/schemas => ../../../service/schemas/
