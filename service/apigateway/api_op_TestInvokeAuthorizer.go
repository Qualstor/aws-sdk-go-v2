// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Simulate the execution of an Authorizer in your RestApi with headers,
// parameters, and an incoming request body.
func (c *Client) TestInvokeAuthorizer(ctx context.Context, params *TestInvokeAuthorizerInput, optFns ...func(*Options)) (*TestInvokeAuthorizerOutput, error) {
	if params == nil {
		params = &TestInvokeAuthorizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestInvokeAuthorizer", params, optFns, c.addOperationTestInvokeAuthorizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestInvokeAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Make a request to simulate the invocation of an Authorizer.
type TestInvokeAuthorizerInput struct {

	// Specifies a test invoke authorizer request's Authorizer ID.
	//
	// This member is required.
	AuthorizerId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// A key-value map of additional context variables.
	AdditionalContext map[string]string

	// The simulated request body of an incoming invocation request.
	Body *string

	// A key-value map of headers to simulate an incoming invocation request. This is
	// where the incoming authorization token, or identity source, should be specified.
	Headers map[string]string

	// The headers as a map from string to list of values to simulate an incoming
	// invocation request. This is where the incoming authorization token, or identity
	// source, may be specified.
	MultiValueHeaders map[string][]string

	// The URI path, including query string, of the simulated invocation request. Use
	// this to specify path parameters and query string parameters.
	PathWithQueryString *string

	// A key-value map of stage variables to simulate an invocation on a deployed
	// Stage.
	StageVariables map[string]string

	noSmithyDocumentSerde
}

// Represents the response of the test invoke request for a custom Authorizer
type TestInvokeAuthorizerOutput struct {

	// The authorization response.
	Authorization map[string][]string

	// The open identity claims, with any supported custom attributes, returned from
	// the Cognito Your User Pool configured for the API.
	Claims map[string]string

	// The HTTP status code that the client would have received. Value is 0 if the
	// authorizer succeeded.
	ClientStatus int32

	// The execution latency of the test authorizer request.
	Latency int64

	// The API Gateway execution log for the test authorizer request.
	Log *string

	// The JSON policy document returned by the Authorizer
	Policy *string

	// The principal identity returned by the Authorizer
	PrincipalId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestInvokeAuthorizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestInvokeAuthorizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestInvokeAuthorizer{}, middleware.After)
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
	if err = addOpTestInvokeAuthorizerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestInvokeAuthorizer(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opTestInvokeAuthorizer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "TestInvokeAuthorizer",
	}
}
