// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of distribution IDs for distributions that have a cache behavior
// that’s associated with the specified response headers policy. You can optionally
// specify the maximum number of items to receive in the response. If the total
// number of items in the list exceeds the maximum that you specify, or the default
// maximum, the response is paginated. To get the next page of items, send a
// subsequent request that specifies the NextMarker value from the current response
// as the Marker value in the subsequent request.
func (c *Client) ListDistributionsByResponseHeadersPolicyId(ctx context.Context, params *ListDistributionsByResponseHeadersPolicyIdInput, optFns ...func(*Options)) (*ListDistributionsByResponseHeadersPolicyIdOutput, error) {
	if params == nil {
		params = &ListDistributionsByResponseHeadersPolicyIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDistributionsByResponseHeadersPolicyId", params, optFns, c.addOperationListDistributionsByResponseHeadersPolicyIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDistributionsByResponseHeadersPolicyIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDistributionsByResponseHeadersPolicyIdInput struct {

	// The ID of the response headers policy whose associated distribution IDs you want
	// to list.
	//
	// This member is required.
	ResponseHeadersPolicyId *string

	// Use this field when paginating results to indicate where to begin in your list
	// of distribution IDs. The response includes distribution IDs in the list that
	// occur after the marker. To get the next page of the list, set this field’s value
	// to the value of NextMarker from the current page’s response.
	Marker *string

	// The maximum number of distribution IDs that you want to get in the response.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListDistributionsByResponseHeadersPolicyIdOutput struct {

	// A list of distribution IDs.
	DistributionIdList *types.DistributionIdList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDistributionsByResponseHeadersPolicyIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListDistributionsByResponseHeadersPolicyId{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListDistributionsByResponseHeadersPolicyId{}, middleware.After)
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
	if err = addOpListDistributionsByResponseHeadersPolicyIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDistributionsByResponseHeadersPolicyId(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDistributionsByResponseHeadersPolicyId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListDistributionsByResponseHeadersPolicyId",
	}
}
