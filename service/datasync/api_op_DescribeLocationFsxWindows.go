// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata about an Amazon FSx for Windows File Server location, such as
// information about its path.
func (c *Client) DescribeLocationFsxWindows(ctx context.Context, params *DescribeLocationFsxWindowsInput, optFns ...func(*Options)) (*DescribeLocationFsxWindowsOutput, error) {
	if params == nil {
		params = &DescribeLocationFsxWindowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationFsxWindows", params, optFns, c.addOperationDescribeLocationFsxWindowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationFsxWindowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocationFsxWindowsInput struct {

	// The Amazon Resource Name (ARN) of the FSx for Windows File Server location to
	// describe.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

type DescribeLocationFsxWindowsOutput struct {

	// The time that the FSx for Windows File Server location was created.
	CreationTime *time.Time

	// The name of the Windows domain that the FSx for Windows File Server belongs to.
	Domain *string

	// The Amazon Resource Name (ARN) of the FSx for Windows File Server location that
	// was described.
	LocationArn *string

	// The URL of the FSx for Windows File Server location that was described.
	LocationUri *string

	// The Amazon Resource Names (ARNs) of the security groups that are configured for
	// the FSx for Windows File Server file system.
	SecurityGroupArns []string

	// The user who has the permissions to access files and folders in the FSx for
	// Windows File Server file system.
	User *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationFsxWindowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationFsxWindows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationFsxWindows{}, middleware.After)
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
	if err = addOpDescribeLocationFsxWindowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationFsxWindows(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationFsxWindows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeLocationFsxWindows",
	}
}
