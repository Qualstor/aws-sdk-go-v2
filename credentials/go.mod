module github.com/Qualstor/aws-sdk-go-v2/credentials

go 1.15

require (
	github.com/Qualstor/aws-sdk-go-v2 v1.16.16
	github.com/Qualstor/aws-sdk-go-v2/feature/ec2/imds v1.12.9
	github.com/Qualstor/aws-sdk-go-v2/service/sso v1.11.12
	github.com/Qualstor/aws-sdk-go-v2/service/sts v1.16.10
	github.com/aws/smithy-go v1.11.3
	github.com/google/go-cmp v0.5.8
)

replace github.com/Qualstor/aws-sdk-go-v2 => ../

replace github.com/Qualstor/aws-sdk-go-v2/feature/ec2/imds => ../feature/ec2/imds/

replace github.com/Qualstor/aws-sdk-go-v2/internal/configsources => ../internal/configsources/

replace github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 => ../internal/endpoints/v2/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/presigned-url => ../service/internal/presigned-url/

replace github.com/Qualstor/aws-sdk-go-v2/service/sso => ../service/sso/

replace github.com/Qualstor/aws-sdk-go-v2/service/sts => ../service/sts/
