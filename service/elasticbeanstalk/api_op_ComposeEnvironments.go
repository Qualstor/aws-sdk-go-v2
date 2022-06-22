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

// Create or update a group of environments that each run a separate component of a
// single application. Takes a list of version labels that specify application
// source bundles for each of the environments to create or update. The name of
// each environment and other required information must be included in the source
// bundles in an environment manifest named env.yaml. See Compose Environments
// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/environment-mgmt-compose.html)
// for details.
func (c *Client) ComposeEnvironments(ctx context.Context, params *ComposeEnvironmentsInput, optFns ...func(*Options)) (*ComposeEnvironmentsOutput, error) {
	if params == nil {
		params = &ComposeEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ComposeEnvironments", params, optFns, c.addOperationComposeEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ComposeEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to create or update a group of environments.
type ComposeEnvironmentsInput struct {

	// The name of the application to which the specified source bundles belong.
	ApplicationName *string

	// The name of the group to which the target environments belong. Specify a group
	// name only if the environment name defined in each target environment's manifest
	// ends with a + (plus) character. See Environment Manifest (env.yaml)
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/environment-cfg-manifest.html)
	// for details.
	GroupName *string

	// A list of version labels, specifying one or more application source bundles that
	// belong to the target application. Each source bundle must include an environment
	// manifest that specifies the name of the environment and the name of the solution
	// stack to use, and optionally can specify environment links to create.
	VersionLabels []string

	noSmithyDocumentSerde
}

// Result message containing a list of environment descriptions.
type ComposeEnvironmentsOutput struct {

	// Returns an EnvironmentDescription list.
	Environments []types.EnvironmentDescription

	// In a paginated request, the token that you can pass in a subsequent request to
	// get the next response page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationComposeEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpComposeEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpComposeEnvironments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opComposeEnvironments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opComposeEnvironments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "ComposeEnvironments",
	}
}
