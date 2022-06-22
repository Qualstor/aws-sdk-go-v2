module github.com/Qualstor/aws-sdk-go-v2/service/transcribestreaming/internal/testing

go 1.15

require (
	github.com/Qualstor/aws-sdk-go-v2 v1.16.5
	github.com/Qualstor/aws-sdk-go-v2/aws/protocol/eventstream v1.4.2
	github.com/Qualstor/aws-sdk-go-v2/service/internal/eventstreamtesting v1.0.19
	github.com/Qualstor/aws-sdk-go-v2/service/transcribestreaming v1.6.8
	github.com/aws/smithy-go v1.11.3
	github.com/google/go-cmp v0.5.8
)

replace github.com/Qualstor/aws-sdk-go-v2 => ../../../../

replace github.com/Qualstor/aws-sdk-go-v2/aws/protocol/eventstream => ../../../../aws/protocol/eventstream/

replace github.com/Qualstor/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/Qualstor/aws-sdk-go-v2/feature/ec2/imds => ../../../../feature/ec2/imds/

replace github.com/Qualstor/aws-sdk-go-v2/internal/configsources => ../../../../internal/configsources/

replace github.com/Qualstor/aws-sdk-go-v2/internal/endpoints/v2 => ../../../../internal/endpoints/v2/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/eventstreamtesting => ../../../../service/internal/eventstreamtesting/

replace github.com/Qualstor/aws-sdk-go-v2/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace github.com/Qualstor/aws-sdk-go-v2/service/sso => ../../../../service/sso/

replace github.com/Qualstor/aws-sdk-go-v2/service/sts => ../../../../service/sts/

replace github.com/Qualstor/aws-sdk-go-v2/service/transcribestreaming => ../../../../service/transcribestreaming/
