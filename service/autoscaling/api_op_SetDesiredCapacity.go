// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the size of the specified Auto Scaling group. If a scale-in activity occurs
// as a result of a new DesiredCapacity value that is lower than the current size
// of the group, the Auto Scaling group uses its termination policy to determine
// which instances to terminate. For more information, see Manual scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-manual-scaling.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) SetDesiredCapacity(ctx context.Context, params *SetDesiredCapacityInput, optFns ...func(*Options)) (*SetDesiredCapacityOutput, error) {
	if params == nil {
		params = &SetDesiredCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetDesiredCapacity", params, optFns, c.addOperationSetDesiredCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetDesiredCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetDesiredCapacityInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The desired capacity is the initial capacity of the Auto Scaling group after
	// this operation completes and the capacity it attempts to maintain.
	//
	// This member is required.
	DesiredCapacity *int32

	// Indicates whether Amazon EC2 Auto Scaling waits for the cooldown period to
	// complete before initiating a scaling activity to set your Auto Scaling group to
	// its new capacity. By default, Amazon EC2 Auto Scaling does not honor the
	// cooldown period during manual scaling activities.
	HonorCooldown *bool

	noSmithyDocumentSerde
}

type SetDesiredCapacityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetDesiredCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetDesiredCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetDesiredCapacity{}, middleware.After)
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
	if err = addOpSetDesiredCapacityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetDesiredCapacity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetDesiredCapacity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "SetDesiredCapacity",
	}
}
