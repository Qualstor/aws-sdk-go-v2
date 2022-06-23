module github.com/Qualstor/aws-sdk-go-v2/service/s3/internal/configtesting

go 1.15

require (
	github.com/Qualstor/aws-sdk-go-v2/config v1.15.16
	github.com/Qualstor/aws-sdk-go-v2/service/internal/s3shared v1.13.9
)

replace github.com/Qualstor/aws-sdk-go-v2 => ../../../../

replace github.com/Qualstor/aws-sdk-go-v2/config => ../../../../config/

replace github.com/Qualstor/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/Qualstor/aws-sdk-go-v2/feature/ec2/imds => ../../../../feature/ec2/imds/

replace github.com/Qualstor/aws-sdk-go-v2/internal/configsources => ../../../../internal/configsources/

replace github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 => ../../../../internal/endpoints/v2/

replace github.com/Qualstor/aws-sdk-go-v2/internal/ini => ../../../../internal/ini/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/s3shared => ../../../../service/internal/s3shared/

replace github.com/Qualstor/aws-sdk-go-v2/service/sso => ../../../../service/sso/

replace github.com/Qualstor/aws-sdk-go-v2/service/sts => ../../../../service/sts/
