// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Describes the specified images (AMIs, AKIs, and ARIs) available to you or all of
// the images available to you. The images available to you include public images,
// private images that you own, and private images owned by other Amazon Web
// Services accounts for which you have explicit launch permissions. Recently
// deregistered images appear in the returned results for a short interval and then
// return empty results. After all instances that reference a deregistered AMI are
// terminated, specifying the ID of the image will eventually return an error
// indicating that the AMI ID cannot be found.
func (c *Client) DescribeImages(ctx context.Context, params *DescribeImagesInput, optFns ...func(*Options)) (*DescribeImagesOutput, error) {
	if params == nil {
		params = &DescribeImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImages", params, optFns, c.addOperationDescribeImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImagesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Scopes the images by users with explicit launch permissions. Specify an Amazon
	// Web Services account ID, self (the sender of the request), or all (public
	// AMIs).
	//
	// * If you specify an Amazon Web Services account ID that is not your own,
	// only AMIs shared with that specific Amazon Web Services account ID are returned.
	// However, AMIs that are shared with the account’s organization or organizational
	// unit (OU) are not returned.
	//
	// * If you specify self or your own Amazon Web
	// Services account ID, AMIs shared with your account are returned. In addition,
	// AMIs that are shared with the organization or OU of which you are member are
	// also returned.
	//
	// * If you specify all, all public AMIs are returned.
	ExecutableUsers []string

	// The filters.
	//
	// * architecture - The image architecture (i386 | x86_64 |
	// arm64).
	//
	// * block-device-mapping.delete-on-termination - A Boolean value that
	// indicates whether the Amazon EBS volume is deleted on instance termination.
	//
	// *
	// block-device-mapping.device-name - The device name specified in the block device
	// mapping (for example, /dev/sdh or xvdh).
	//
	// * block-device-mapping.snapshot-id -
	// The ID of the snapshot used for the Amazon EBS volume.
	//
	// *
	// block-device-mapping.volume-size - The volume size of the Amazon EBS volume, in
	// GiB.
	//
	// * block-device-mapping.volume-type - The volume type of the Amazon EBS
	// volume (io1 | io2 | gp2 | gp3 | sc1 | st1 | standard).
	//
	// *
	// block-device-mapping.encrypted - A Boolean that indicates whether the Amazon EBS
	// volume is encrypted.
	//
	// * creation-date - The time when the image was created, in
	// the ISO 8601 format in the UTC time zone (YYYY-MM-DDThh:mm:ss.sssZ), for
	// example, 2021-09-29T11:04:43.305Z. You can use a wildcard (*), for example,
	// 2021-09-29T*, which matches an entire day.
	//
	// * description - The description of
	// the image (provided during image creation).
	//
	// * ena-support - A Boolean that
	// indicates whether enhanced networking with ENA is enabled.
	//
	// * hypervisor - The
	// hypervisor type (ovm | xen).
	//
	// * image-id - The ID of the image.
	//
	// * image-type -
	// The image type (machine | kernel | ramdisk).
	//
	// * is-public - A Boolean that
	// indicates whether the image is public.
	//
	// * kernel-id - The kernel ID.
	//
	// *
	// manifest-location - The location of the image manifest.
	//
	// * name - The name of
	// the AMI (provided during image creation).
	//
	// * owner-alias - The owner alias
	// (amazon | aws-marketplace). The valid aliases are defined in an
	// Amazon-maintained list. This is not the Amazon Web Services account alias that
	// can be set using the IAM console. We recommend that you use the Owner request
	// parameter instead of this filter.
	//
	// * owner-id - The Amazon Web Services account
	// ID of the owner. We recommend that you use the Owner request parameter instead
	// of this filter.
	//
	// * platform - The platform. To only list Windows-based AMIs, use
	// windows.
	//
	// * product-code - The product code.
	//
	// * product-code.type - The type of
	// the product code (marketplace).
	//
	// * ramdisk-id - The RAM disk ID.
	//
	// *
	// root-device-name - The device name of the root device volume (for example,
	// /dev/sda1).
	//
	// * root-device-type - The type of the root device volume (ebs |
	// instance-store).
	//
	// * state - The state of the image (available | pending |
	// failed).
	//
	// * state-reason-code - The reason code for the state change.
	//
	// *
	// state-reason-message - The message for the state change.
	//
	// * sriov-net-support -
	// A value of simple indicates that enhanced networking with the Intel 82599 VF
	// interface is enabled.
	//
	// * tag: - The key/value combination of a tag assigned to
	// the resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	// * tag-key - The key of a tag assigned to the resource. Use this filter
	// to find all resources assigned a tag with a specific key, regardless of the tag
	// value.
	//
	// * virtualization-type - The virtualization type (paravirtual | hvm).
	Filters []types.Filter

	// The image IDs. Default: Describes all images available to you.
	ImageIds []string

	// If true, all deprecated AMIs are included in the response. If false, no
	// deprecated AMIs are included in the response. If no value is specified, the
	// default value is false. If you are the AMI owner, all deprecated AMIs appear in
	// the response regardless of the value (true or false) that you set for this
	// parameter.
	IncludeDeprecated *bool

	// Scopes the results to images with the specified owners. You can specify a
	// combination of Amazon Web Services account IDs, self, amazon, and
	// aws-marketplace. If you omit this parameter, the results include all images for
	// which you have launch permissions, regardless of ownership.
	Owners []string

	noSmithyDocumentSerde
}

type DescribeImagesOutput struct {

	// Information about the images.
	Images []types.Image

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeImages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImages(options.Region), middleware.Before); err != nil {
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

// DescribeImagesAPIClient is a client that implements the DescribeImages
// operation.
type DescribeImagesAPIClient interface {
	DescribeImages(context.Context, *DescribeImagesInput, ...func(*Options)) (*DescribeImagesOutput, error)
}

var _ DescribeImagesAPIClient = (*Client)(nil)

// ImageAvailableWaiterOptions are waiter options for ImageAvailableWaiter
type ImageAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ImageAvailableWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ImageAvailableWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeImagesInput, *DescribeImagesOutput, error) (bool, error)
}

// ImageAvailableWaiter defines the waiters for ImageAvailable
type ImageAvailableWaiter struct {
	client DescribeImagesAPIClient

	options ImageAvailableWaiterOptions
}

// NewImageAvailableWaiter constructs a ImageAvailableWaiter.
func NewImageAvailableWaiter(client DescribeImagesAPIClient, optFns ...func(*ImageAvailableWaiterOptions)) *ImageAvailableWaiter {
	options := ImageAvailableWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = imageAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ImageAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ImageAvailable waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ImageAvailableWaiter) Wait(ctx context.Context, params *DescribeImagesInput, maxWaitDur time.Duration, optFns ...func(*ImageAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ImageAvailable waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *ImageAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeImagesInput, maxWaitDur time.Duration, optFns ...func(*ImageAvailableWaiterOptions)) (*DescribeImagesOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeImages(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ImageAvailable waiter")
}

func imageAvailableStateRetryable(ctx context.Context, input *DescribeImagesInput, output *DescribeImagesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Images[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.ImageState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.ImageState value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Images[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "failed"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.ImageState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.ImageState value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// ImageExistsWaiterOptions are waiter options for ImageExistsWaiter
type ImageExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ImageExistsWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ImageExistsWaiter will use default max delay of 120 seconds. Note that
	// MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeImagesInput, *DescribeImagesOutput, error) (bool, error)
}

// ImageExistsWaiter defines the waiters for ImageExists
type ImageExistsWaiter struct {
	client DescribeImagesAPIClient

	options ImageExistsWaiterOptions
}

// NewImageExistsWaiter constructs a ImageExistsWaiter.
func NewImageExistsWaiter(client DescribeImagesAPIClient, optFns ...func(*ImageExistsWaiterOptions)) *ImageExistsWaiter {
	options := ImageExistsWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = imageExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ImageExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ImageExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ImageExistsWaiter) Wait(ctx context.Context, params *DescribeImagesInput, maxWaitDur time.Duration, optFns ...func(*ImageExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ImageExists waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *ImageExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeImagesInput, maxWaitDur time.Duration, optFns ...func(*ImageExistsWaiterOptions)) (*DescribeImagesOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeImages(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ImageExists waiter")
}

func imageExistsStateRetryable(ctx context.Context, input *DescribeImagesInput, output *DescribeImagesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("length(Images[]) > `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "InvalidAMIID.NotFound" == apiErr.ErrorCode() {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeImages",
	}
}
