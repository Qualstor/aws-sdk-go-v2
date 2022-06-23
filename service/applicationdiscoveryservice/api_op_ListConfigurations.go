// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of configuration items as specified by the value passed to the
// required parameter configurationType. Optional filtering may be applied to
// refine search results.
func (c *Client) ListConfigurations(ctx context.Context, params *ListConfigurationsInput, optFns ...func(*Options)) (*ListConfigurationsOutput, error) {
	if params == nil {
		params = &ListConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurations", params, optFns, c.addOperationListConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationsInput struct {

	// A valid configuration identified by Application Discovery Service.
	//
	// This member is required.
	ConfigurationType types.ConfigurationItemType

	// You can filter the request using various logical operators and a key-value
	// format. For example: {"key": "serverType", "value": "webServer"} For a complete
	// list of filter options and guidance about using them with this action, see Using
	// the ListConfigurations Action
	// (https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-api-queries.html#ListConfigurations)
	// in the Amazon Web Services Application Discovery Service User Guide.
	Filters []types.Filter

	// The total number of items to return. The maximum value is 100.
	MaxResults int32

	// Token to retrieve the next set of results. For example, if a previous call to
	// ListConfigurations returned 100 items, but you set
	// ListConfigurationsRequest$maxResults to 10, you received a set of 10 results
	// along with a token. Use that token in this query to get the next set of 10.
	NextToken *string

	// Certain filter criteria return output that can be sorted in ascending or
	// descending order. For a list of output characteristics for each filter, see
	// Using the ListConfigurations Action
	// (https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-api-queries.html#ListConfigurations)
	// in the Amazon Web Services Application Discovery Service User Guide.
	OrderBy []types.OrderByElement

	noSmithyDocumentSerde
}

type ListConfigurationsOutput struct {

	// Returns configuration details, including the configuration ID, attribute names,
	// and attribute values.
	Configurations []map[string]string

	// Token to retrieve the next set of results. For example, if your call to
	// ListConfigurations returned 100 items, but you set
	// ListConfigurationsRequest$maxResults to 10, you received a set of 10 results
	// along with this token. Use this token in the next query to retrieve the next set
	// of 10.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConfigurations{}, middleware.After)
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
	if err = addOpListConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "ListConfigurations",
	}
}
