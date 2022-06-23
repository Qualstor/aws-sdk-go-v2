// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates an embed URL that you can use to embed an Amazon QuickSight dashboard
// in your website, without having to register any reader users. Before you use
// this action, make sure that you have configured the dashboards and permissions.
// The following rules apply to the generated URL:
//
// * It contains a temporary
// bearer token. It is valid for 5 minutes after it is generated. Once redeemed
// within this period, it cannot be re-used again.
//
// * The URL validity period
// should not be confused with the actual session lifetime that can be customized
// using the SessionLifetimeInMinutes
// (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GenerateEmbedUrlForAnonymousUser.html#QS-GenerateEmbedUrlForAnonymousUser-request-SessionLifetimeInMinutes)
// parameter. The resulting user session is valid for 15 minutes (minimum) to 10
// hours (maximum). The default session duration is 10 hours.
//
// * You are charged
// only when the URL is used or there is interaction with Amazon QuickSight.
//
// For
// more information, see Embedded Analytics
// (https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics.html) in
// the Amazon QuickSight User Guide. For more information about the high-level
// steps for embedding and for an interactive demo of the ways you can customize
// embedding, visit the Amazon QuickSight Developer Portal
// (https://docs.aws.amazon.com/quicksight/latest/user/quicksight-dev-portal.html).
func (c *Client) GenerateEmbedUrlForAnonymousUser(ctx context.Context, params *GenerateEmbedUrlForAnonymousUserInput, optFns ...func(*Options)) (*GenerateEmbedUrlForAnonymousUserOutput, error) {
	if params == nil {
		params = &GenerateEmbedUrlForAnonymousUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateEmbedUrlForAnonymousUser", params, optFns, c.addOperationGenerateEmbedUrlForAnonymousUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateEmbedUrlForAnonymousUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateEmbedUrlForAnonymousUserInput struct {

	// The Amazon Resource Names for the Amazon QuickSight resources that the user is
	// authorized to access during the lifetime of the session. If you choose Dashboard
	// embedding experience, pass the list of dashboard ARNs in the account that you
	// want the user to be able to view. Currently, you can pass up to 25 dashboard
	// ARNs in each API call.
	//
	// This member is required.
	AuthorizedResourceArns []string

	// The ID for the Amazon Web Services account that contains the dashboard that
	// you're embedding.
	//
	// This member is required.
	AwsAccountId *string

	// The configuration of the experience you are embedding.
	//
	// This member is required.
	ExperienceConfiguration *types.AnonymousUserEmbeddingExperienceConfiguration

	// The Amazon QuickSight namespace that the anonymous user virtually belongs to. If
	// you are not using an Amazon QuickSight custom namespace, set this to default.
	//
	// This member is required.
	Namespace *string

	// How many minutes the session is valid. The session lifetime must be in [15-600]
	// minutes range.
	SessionLifetimeInMinutes *int64

	// The session tags used for row-level security. Before you use this parameter,
	// make sure that you have configured the relevant datasets using the
	// DataSet$RowLevelPermissionTagConfiguration parameter so that session tags can be
	// used to provide row-level security. These are not the tags used for the Amazon
	// Web Services resource tagging feature. For more information, see Using Row-Level
	// Security (RLS) with Tags
	// (https://docs.aws.amazon.com/quicksight/latest/user/quicksight-dev-rls-tags.html).
	SessionTags []types.SessionTag

	noSmithyDocumentSerde
}

type GenerateEmbedUrlForAnonymousUserOutput struct {

	// The embed URL for the dashboard.
	//
	// This member is required.
	EmbedUrl *string

	// The Amazon Web Services request ID for this operation.
	//
	// This member is required.
	RequestId *string

	// The HTTP status of the request.
	//
	// This member is required.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateEmbedUrlForAnonymousUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGenerateEmbedUrlForAnonymousUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGenerateEmbedUrlForAnonymousUser{}, middleware.After)
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
	if err = addOpGenerateEmbedUrlForAnonymousUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateEmbedUrlForAnonymousUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateEmbedUrlForAnonymousUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "GenerateEmbedUrlForAnonymousUser",
	}
}
