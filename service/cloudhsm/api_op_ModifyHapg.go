// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/), the AWS
// CloudHSM Classic User Guide
// (https://docs.aws.amazon.com/cloudhsm/classic/userguide/), and the AWS CloudHSM
// Classic API Reference
// (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/). For information
// about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide
// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/), and the AWS CloudHSM
// API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
// Modifies an existing high-availability partition group.
func (c *Client) ModifyHapg(ctx context.Context, params *ModifyHapgInput, optFns ...func(*Options)) (*ModifyHapgOutput, error) {
	if params == nil {
		params = &ModifyHapgInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyHapg", params, optFns, c.addOperationModifyHapgMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyHapgOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyHapgInput struct {

	// The ARN of the high-availability partition group to modify.
	//
	// This member is required.
	HapgArn *string

	// The new label for the high-availability partition group.
	Label *string

	// The list of partition serial numbers to make members of the high-availability
	// partition group.
	PartitionSerialList []string

	noSmithyDocumentSerde
}

type ModifyHapgOutput struct {

	// The ARN of the high-availability partition group.
	HapgArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyHapgMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyHapg{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyHapg{}, middleware.After)
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
	if err = addOpModifyHapgValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyHapg(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyHapg(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "ModifyHapg",
	}
}
