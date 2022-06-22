// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enable or disable collection of reputation metrics for emails that you send
// using a particular configuration set in a specific AWS Region.
func (c *Client) PutConfigurationSetReputationOptions(ctx context.Context, params *PutConfigurationSetReputationOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetReputationOptionsOutput, error) {
	if params == nil {
		params = &PutConfigurationSetReputationOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigurationSetReputationOptions", params, optFns, c.addOperationPutConfigurationSetReputationOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigurationSetReputationOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to enable or disable tracking of reputation metrics for a
// configuration set.
type PutConfigurationSetReputationOptionsInput struct {

	// The name of the configuration set that you want to enable or disable reputation
	// metric tracking for.
	//
	// This member is required.
	ConfigurationSetName *string

	// If true, tracking of reputation metrics is enabled for the configuration set. If
	// false, tracking of reputation metrics is disabled for the configuration set.
	ReputationMetricsEnabled bool

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutConfigurationSetReputationOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutConfigurationSetReputationOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutConfigurationSetReputationOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutConfigurationSetReputationOptions{}, middleware.After)
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
	if err = addOpPutConfigurationSetReputationOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationSetReputationOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutConfigurationSetReputationOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutConfigurationSetReputationOptions",
	}
}
