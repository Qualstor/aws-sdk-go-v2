// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an application version for the specified application. You can create an
// application version from a source bundle in Amazon S3, a commit in AWS
// CodeCommit, or the output of an AWS CodeBuild build as follows: Specify a commit
// in an AWS CodeCommit repository with SourceBuildInformation. Specify a build in
// an AWS CodeBuild with SourceBuildInformation and BuildConfiguration. Specify a
// source bundle in S3 with SourceBundle Omit both SourceBuildInformation and
// SourceBundle to use the default sample application. After you create an
// application version with a specified Amazon S3 bucket and key location, you
// can't change that Amazon S3 location. If you change the Amazon S3 location, you
// receive an exception when you attempt to launch an environment from the
// application version.
func (c *Client) CreateApplicationVersion(ctx context.Context, params *CreateApplicationVersionInput, optFns ...func(*Options)) (*CreateApplicationVersionOutput, error) {
	if params == nil {
		params = &CreateApplicationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApplicationVersion", params, optFns, c.addOperationCreateApplicationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateApplicationVersionInput struct {

	// The name of the application. If no application is found with this name, and
	// AutoCreateApplication is false, returns an InvalidParameterValue error.
	//
	// This member is required.
	ApplicationName *string

	// A label identifying this version. Constraint: Must be unique per application. If
	// an application version already exists with this label for the specified
	// application, AWS Elastic Beanstalk returns an InvalidParameterValue error.
	//
	// This member is required.
	VersionLabel *string

	// Set to true to create an application with the specified name if it doesn't
	// already exist.
	AutoCreateApplication *bool

	// Settings for an AWS CodeBuild build.
	BuildConfiguration *types.BuildConfiguration

	// A description of this application version.
	Description *string

	// Pre-processes and validates the environment manifest (env.yaml) and
	// configuration files (*.config files in the .ebextensions folder) in the source
	// bundle. Validating configuration files can identify issues prior to deploying
	// the application version to an environment. You must turn processing on for
	// application versions that you create using AWS CodeBuild or AWS CodeCommit. For
	// application versions built from a source bundle in Amazon S3, processing is
	// optional. The Process option validates Elastic Beanstalk configuration files. It
	// doesn't validate your application's configuration files, like proxy server or
	// Docker configuration.
	Process *bool

	// Specify a commit in an AWS CodeCommit Git repository to use as the source code
	// for the application version.
	SourceBuildInformation *types.SourceBuildInformation

	// The Amazon S3 bucket and key that identify the location of the source bundle for
	// this version. The Amazon S3 bucket must be in the same region as the
	// environment. Specify a source bundle in S3 or a commit in an AWS CodeCommit
	// repository (with SourceBuildInformation), but not both. If neither SourceBundle
	// nor SourceBuildInformation are provided, Elastic Beanstalk uses a sample
	// application.
	SourceBundle *types.S3Location

	// Specifies the tags applied to the application version. Elastic Beanstalk applies
	// these tags only to the application version. Environments that use the
	// application version don't inherit the tags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Result message wrapping a single description of an application version.
type CreateApplicationVersionOutput struct {

	// The ApplicationVersionDescription of the application version.
	ApplicationVersion *types.ApplicationVersionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApplicationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateApplicationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateApplicationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateApplicationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplicationVersion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateApplicationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "CreateApplicationVersion",
	}
}
