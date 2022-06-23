// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a health check. Amazon Route 53 does not prevent you from deleting a
// health check even if the health check is associated with one or more resource
// record sets. If you delete a health check and you don't update the associated
// resource record sets, the future status of the health check can't be predicted
// and may change. This will affect the routing of DNS queries for your DNS
// failover configuration. For more information, see Replacing and Deleting Health
// Checks
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating-deleting.html#health-checks-deleting.html)
// in the Amazon Route 53 Developer Guide. If you're using Cloud Map and you
// configured Cloud Map to create a Route 53 health check when you register an
// instance, you can't use the Route 53 DeleteHealthCheck command to delete the
// health check. The health check is deleted automatically when you deregister the
// instance; there can be a delay of several hours before the health check is
// deleted from Route 53.
func (c *Client) DeleteHealthCheck(ctx context.Context, params *DeleteHealthCheckInput, optFns ...func(*Options)) (*DeleteHealthCheckOutput, error) {
	if params == nil {
		params = &DeleteHealthCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteHealthCheck", params, optFns, c.addOperationDeleteHealthCheckMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This action deletes a health check.
type DeleteHealthCheckInput struct {

	// The ID of the health check that you want to delete.
	//
	// This member is required.
	HealthCheckId *string

	noSmithyDocumentSerde
}

// An empty element.
type DeleteHealthCheckOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteHealthCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteHealthCheck{}, middleware.After)
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
	if err = addOpDeleteHealthCheckValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHealthCheck(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteHealthCheck(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteHealthCheck",
	}
}
