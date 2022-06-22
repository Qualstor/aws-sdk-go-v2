// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a stickiness policy with sticky session lifetimes controlled by the
// lifetime of the browser (user-agent) or a specified expiration period. This
// policy can be associated only with HTTP/HTTPS listeners. When a load balancer
// implements this policy, the load balancer uses a special cookie to track the
// instance for each request. When the load balancer receives a request, it first
// checks to see if this cookie is present in the request. If so, the load balancer
// sends the request to the application server specified in the cookie. If not, the
// load balancer sends the request to a server that is chosen based on the existing
// load-balancing algorithm. A cookie is inserted into the response for binding
// subsequent requests from the same user to that server. The validity of the
// cookie is based on the cookie expiration time, which is specified in the policy
// configuration. For more information, see Duration-Based Session Stickiness
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-duration)
// in the Classic Load Balancers Guide.
func (c *Client) CreateLBCookieStickinessPolicy(ctx context.Context, params *CreateLBCookieStickinessPolicyInput, optFns ...func(*Options)) (*CreateLBCookieStickinessPolicyOutput, error) {
	if params == nil {
		params = &CreateLBCookieStickinessPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLBCookieStickinessPolicy", params, optFns, c.addOperationCreateLBCookieStickinessPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLBCookieStickinessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateLBCookieStickinessPolicy.
type CreateLBCookieStickinessPolicyInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The name of the policy being created. Policy names must consist of alphanumeric
	// characters and dashes (-). This name must be unique within the set of policies
	// for this load balancer.
	//
	// This member is required.
	PolicyName *string

	// The time period, in seconds, after which the cookie should be considered stale.
	// If you do not specify this parameter, the default value is 0, which indicates
	// that the sticky session should last for the duration of the browser session.
	CookieExpirationPeriod *int64

	noSmithyDocumentSerde
}

// Contains the output for CreateLBCookieStickinessPolicy.
type CreateLBCookieStickinessPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLBCookieStickinessPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLBCookieStickinessPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLBCookieStickinessPolicy{}, middleware.After)
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
	if err = addOpCreateLBCookieStickinessPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLBCookieStickinessPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLBCookieStickinessPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateLBCookieStickinessPolicy",
	}
}
