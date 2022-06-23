// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Information on the shipping label of a Snow device that is being returned to
// Amazon Web Services.
func (c *Client) DescribeReturnShippingLabel(ctx context.Context, params *DescribeReturnShippingLabelInput, optFns ...func(*Options)) (*DescribeReturnShippingLabelOutput, error) {
	if params == nil {
		params = &DescribeReturnShippingLabelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReturnShippingLabel", params, optFns, c.addOperationDescribeReturnShippingLabelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReturnShippingLabelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReturnShippingLabelInput struct {

	// The automatically generated ID for a job, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeReturnShippingLabelOutput struct {

	// The expiration date of the current return shipping label.
	ExpirationDate *time.Time

	// The pre-signed Amazon S3 URI used to download the return shipping label.
	ReturnShippingLabelURI *string

	// The status information of the task on a Snow device that is being returned to
	// Amazon Web Services.
	Status types.ShippingLabelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReturnShippingLabelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReturnShippingLabel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReturnShippingLabel{}, middleware.After)
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
	if err = addOpDescribeReturnShippingLabelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReturnShippingLabel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReturnShippingLabel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "DescribeReturnShippingLabel",
	}
}
