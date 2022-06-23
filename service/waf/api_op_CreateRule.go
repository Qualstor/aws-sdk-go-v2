// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/waf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Creates a Rule, which contains the IPSet objects, ByteMatchSet
// objects, and other predicates that identify the requests that you want to block.
// If you add more than one predicate to a Rule, a request must match all of the
// specifications to be allowed or blocked. For example, suppose that you add the
// following to a Rule:
//
// * An IPSet that matches the IP address 192.0.2.44/32
//
// * A
// ByteMatchSet that matches BadBot in the User-Agent header
//
// You then add the Rule
// to a WebACL and specify that you want to blocks requests that satisfy the Rule.
// For a request to be blocked, it must come from the IP address 192.0.2.44 and the
// User-Agent header in the request must contain the value BadBot. To create and
// configure a Rule, perform the following steps:
//
// * Create and update the
// predicates that you want to include in the Rule. For more information, see
// CreateByteMatchSet, CreateIPSet, and CreateSqlInjectionMatchSet.
//
// * Use
// GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateRule request.
//
// * Submit a CreateRule request.
//
// * Use
// GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRule request.
//
// * Submit an UpdateRule request to specify
// the predicates that you want to include in the Rule.
//
// * Create and update a
// WebACL that contains the Rule. For more information, see CreateWebACL.
//
// For more
// information about how to use the AWS WAF API to allow or block HTTP requests,
// see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) {
	if params == nil {
		params = &CreateRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRule", params, optFns, c.addOperationCreateRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description for the metrics for this Rule. The name can
	// contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128
	// and minimum length one. It can't contain whitespace or metric names reserved for
	// AWS WAF, including "All" and "Default_Action." You can't change the name of the
	// metric after you create the Rule.
	//
	// This member is required.
	MetricName *string

	// A friendly name or description of the Rule. You can't change the name of a Rule
	// after you create it.
	//
	// This member is required.
	Name *string

	//
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateRuleOutput struct {

	// The ChangeToken that you used to submit the CreateRule request. You can also use
	// this value to query the status of the request. For more information, see
	// GetChangeTokenStatus.
	ChangeToken *string

	// The Rule returned in the CreateRule response.
	Rule *types.Rule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRule{}, middleware.After)
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
	if err = addOpCreateRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "CreateRule",
	}
}
