module github.com/Qualstor/aws-sdk-go-v2/example/service/s3/usingPrivateLink

go 1.15

require (
	github.com/Qualstor/aws-sdk-go-v2 v1.16.15
	github.com/Qualstor/aws-sdk-go-v2/config v1.15.11
	github.com/Qualstor/aws-sdk-go-v2/service/s3 v1.26.11
	github.com/Qualstor/aws-sdk-go-v2/service/s3control v1.21.7
)

replace github.com/Qualstor/aws-sdk-go-v2 => ../../../../

replace github.com/Qualstor/aws-sdk-go-v2/aws/protocol/eventstream => ../../../../aws/protocol/eventstream/

replace github.com/Qualstor/aws-sdk-go-v2/config => ../../../../config/

replace github.com/Qualstor/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/Qualstor/aws-sdk-go-v2/feature/ec2/imds => ../../../../feature/ec2/imds/

replace github.com/Qualstor/aws-sdk-go-v2/internal/configsources => ../../../../internal/configsources/

replace github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 => ../../../../internal/endpoints/v2/

replace github.com/Qualstor/aws-sdk-go-v2/internal/ini => ../../../../internal/ini/

replace github.com/Qualstor/aws-sdk-go-v2/internal/v4a => ../../../../internal/v4a/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/accept-encoding => ../../../../service/internal/accept-encoding/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/checksum => ../../../../service/internal/checksum/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/s3shared => ../../../../service/internal/s3shared/

replace github.com/Qualstor/aws-sdk-go-v2/service/s3 => ../../../../service/s3/

replace github.com/Qualstor/aws-sdk-go-v2/service/s3control => ../../../../service/s3control/

replace github.com/Qualstor/aws-sdk-go-v2/service/sso => ../../../../service/sso/

replace github.com/Qualstor/aws-sdk-go-v2/service/sts => ../../../../service/sts/
