// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of objects associated with a profile of a given
// ProfileObjectType.
func (c *Client) ListProfileObjects(ctx context.Context, params *ListProfileObjectsInput, optFns ...func(*Options)) (*ListProfileObjectsOutput, error) {
	if params == nil {
		params = &ListProfileObjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfileObjects", params, optFns, c.addOperationListProfileObjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfileObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProfileObjectsInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The name of the profile object type.
	//
	// This member is required.
	ObjectTypeName *string

	// The unique identifier of a customer profile.
	//
	// This member is required.
	ProfileId *string

	// The maximum number of objects returned per page.
	MaxResults *int32

	// The pagination token from the previous call to ListProfileObjects.
	NextToken *string

	// Applies a filter to the response to include profile objects with the specified
	// index values. This filter is only supported for ObjectTypeName _asset, _case and
	// _order.
	ObjectFilter *types.ObjectFilter

	noSmithyDocumentSerde
}

type ListProfileObjectsOutput struct {

	// The list of ListProfileObject instances.
	Items []types.ListProfileObjectsItem

	// The pagination token from the previous call to ListProfileObjects.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProfileObjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProfileObjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProfileObjects{}, middleware.After)
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
	if err = addOpListProfileObjectsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProfileObjects(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListProfileObjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "ListProfileObjects",
	}
}
