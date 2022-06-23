// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Add a method to an existing Resource resource.
func (c *Client) PutMethod(ctx context.Context, params *PutMethodInput, optFns ...func(*Options)) (*PutMethodOutput, error) {
	if params == nil {
		params = &PutMethodInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMethod", params, optFns, c.addOperationPutMethodMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMethodOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to add a method to an existing Resource resource.
type PutMethodInput struct {

	// The method's authorization type. Valid values are NONE for open access, AWS_IAM
	// for using AWS IAM permissions, CUSTOM for using a custom authorizer, or
	// COGNITO_USER_POOLS for using a Cognito user pool.
	//
	// This member is required.
	AuthorizationType *string

	// Specifies the method request's HTTP method type.
	//
	// This member is required.
	HttpMethod *string

	// The Resource identifier for the new Method resource.
	//
	// This member is required.
	ResourceId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// Specifies whether the method required a valid ApiKey.
	ApiKeyRequired bool

	// A list of authorization scopes configured on the method. The scopes are used
	// with a COGNITO_USER_POOLS authorizer to authorize the method invocation. The
	// authorization works by matching the method scopes against the scopes parsed from
	// the access token in the incoming request. The method invocation is authorized if
	// any method scopes matches a claimed scope in the access token. Otherwise, the
	// invocation is not authorized. When the method scope is configured, the client
	// must provide an access token instead of an identity token for authorization
	// purposes.
	AuthorizationScopes []string

	// Specifies the identifier of an Authorizer to use on this Method, if the type is
	// CUSTOM or COGNITO_USER_POOLS. The authorizer identifier is generated by API
	// Gateway when you created the authorizer.
	AuthorizerId *string

	// A human-friendly operation identifier for the method. For example, you can
	// assign the operationName of ListPets for the GET /pets method in the PetStore
	// example.
	OperationName *string

	// Specifies the Model resources used for the request's content type. Request
	// models are represented as a key/value map, with a content type as the key and a
	// Model name as the value.
	RequestModels map[string]string

	// A key-value map defining required or optional method request parameters that can
	// be accepted by API Gateway. A key defines a method request parameter name
	// matching the pattern of method.request.{location}.{name}, where location is
	// querystring, path, or header and name is a valid and unique parameter name. The
	// value associated with the key is a Boolean flag indicating whether the parameter
	// is required (true) or optional (false). The method request parameter names
	// defined here are available in Integration to be mapped to integration request
	// parameters or body-mapping templates.
	RequestParameters map[string]bool

	// The identifier of a RequestValidator for validating the method request.
	RequestValidatorId *string

	noSmithyDocumentSerde
}

// Represents a client-facing interface by which the client calls the API to access
// back-end resources. A Method resource is integrated with an Integration
// resource. Both consist of a request and one or more responses. The method
// request takes the client input that is passed to the back end through the
// integration request. A method response returns the output from the back end to
// the client through an integration response. A method request is embodied in a
// Method resource, whereas an integration request is embodied in an Integration
// resource. On the other hand, a method response is represented by a
// MethodResponse resource, whereas an integration response is represented by an
// IntegrationResponse resource.
type PutMethodOutput struct {

	// A boolean flag specifying whether a valid ApiKey is required to invoke this
	// method.
	ApiKeyRequired *bool

	// A list of authorization scopes configured on the method. The scopes are used
	// with a COGNITO_USER_POOLS authorizer to authorize the method invocation. The
	// authorization works by matching the method scopes against the scopes parsed from
	// the access token in the incoming request. The method invocation is authorized if
	// any method scopes matches a claimed scope in the access token. Otherwise, the
	// invocation is not authorized. When the method scope is configured, the client
	// must provide an access token instead of an identity token for authorization
	// purposes.
	AuthorizationScopes []string

	// The method's authorization type. Valid values are NONE for open access, AWS_IAM
	// for using AWS IAM permissions, CUSTOM for using a custom authorizer, or
	// COGNITO_USER_POOLS for using a Cognito user pool.
	AuthorizationType *string

	// The identifier of an Authorizer to use on this method. The authorizationType
	// must be CUSTOM.
	AuthorizerId *string

	// The method's HTTP verb.
	HttpMethod *string

	// Gets the method's integration responsible for passing the client-submitted
	// request to the back end and performing necessary transformations to make the
	// request compliant with the back end.
	MethodIntegration *types.Integration

	// Gets a method response associated with a given HTTP status code.
	MethodResponses map[string]types.MethodResponse

	// A human-friendly operation identifier for the method. For example, you can
	// assign the operationName of ListPets for the GET /pets method in the PetStore
	// example.
	OperationName *string

	// A key-value map specifying data schemas, represented by Model resources, (as the
	// mapped value) of the request payloads of given content types (as the mapping
	// key).
	RequestModels map[string]string

	// A key-value map defining required or optional method request parameters that can
	// be accepted by API Gateway. A key is a method request parameter name matching
	// the pattern of method.request.{location}.{name}, where location is querystring,
	// path, or header and name is a valid and unique parameter name. The value
	// associated with the key is a Boolean flag indicating whether the parameter is
	// required (true) or optional (false). The method request parameter names defined
	// here are available in Integration to be mapped to integration request parameters
	// or templates.
	RequestParameters map[string]bool

	// The identifier of a RequestValidator for request validation.
	RequestValidatorId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutMethodMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutMethod{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutMethod{}, middleware.After)
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
	if err = addOpPutMethodValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMethod(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutMethod(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "PutMethod",
	}
}
