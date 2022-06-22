// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the user name for the specified IAM user. A login profile is created
// when you create a password for the user to access the Amazon Web Services
// Management Console. If the user does not exist or does not have a password, the
// operation returns a 404 (NoSuchEntity) error. If you create an IAM user with
// access to the console, the CreateDate reflects the date you created the initial
// password for the user. If you create an IAM user with programmatic access, and
// then later add a password for the user to access the Amazon Web Services
// Management Console, the CreateDate reflects the initial password creation date.
// A user with programmatic access does not have a login profile unless you create
// a password for the user to access the Amazon Web Services Management Console.
func (c *Client) GetLoginProfile(ctx context.Context, params *GetLoginProfileInput, optFns ...func(*Options)) (*GetLoginProfileOutput, error) {
	if params == nil {
		params = &GetLoginProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLoginProfile", params, optFns, c.addOperationGetLoginProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLoginProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoginProfileInput struct {

	// The name of the user whose login profile you want to retrieve. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

// Contains the response to a successful GetLoginProfile request.
type GetLoginProfileOutput struct {

	// A structure containing the user name and the profile creation date for the user.
	//
	// This member is required.
	LoginProfile *types.LoginProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLoginProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetLoginProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetLoginProfile{}, middleware.After)
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
	if err = addOpGetLoginProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoginProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLoginProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetLoginProfile",
	}
}
