// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of the configuration sets associated with your Amazon SES
// account in the current AWS Region. For information about using configuration
// sets, see Monitoring Your Amazon SES Sending Activity
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second. This operation will return up to 1,000 configuration sets each
// time it is run. If your Amazon SES account has more than 1,000 configuration
// sets, this operation will also return a NextToken element. You can then execute
// the ListConfigurationSets operation again, passing the NextToken parameter and
// the value of the NextToken element to retrieve additional results.
func (c *Client) ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsInput, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) {
	if params == nil {
		params = &ListConfigurationSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurationSets", params, optFns, c.addOperationListConfigurationSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list the configuration sets associated with your AWS
// account. Configuration sets enable you to publish email sending events. For
// information about using configuration sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type ListConfigurationSetsInput struct {

	// The number of configuration sets to return.
	MaxItems *int32

	// A token returned from a previous call to ListConfigurationSets to indicate the
	// position of the configuration set in the configuration set list.
	NextToken *string

	noSmithyDocumentSerde
}

// A list of configuration sets associated with your AWS account. Configuration
// sets enable you to publish email sending events. For information about using
// configuration sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type ListConfigurationSetsOutput struct {

	// A list of configuration sets.
	ConfigurationSets []types.ConfigurationSet

	// A token indicating that there are additional configuration sets available to be
	// listed. Pass this token to successive calls of ListConfigurationSets.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConfigurationSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListConfigurationSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListConfigurationSets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationSets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListConfigurationSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListConfigurationSets",
	}
}
