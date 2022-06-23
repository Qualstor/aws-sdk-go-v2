// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/Qualstor/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/Qualstor/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Returns configuration information about the specified Multi-Region Access Point.
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around managing Multi-Region Access Points,
// see Managing Multi-Region Access Points
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/ManagingMultiRegionAccessPoints.html)
// in the Amazon S3 User Guide. The following actions are related to
// GetMultiRegionAccessPoint:
//
// * CreateMultiRegionAccessPoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateMultiRegionAccessPoint.html)
//
// *
// DeleteMultiRegionAccessPoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteMultiRegionAccessPoint.html)
//
// *
// DescribeMultiRegionAccessPointOperation
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html)
//
// *
// ListMultiRegionAccessPoints
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListMultiRegionAccessPoints.html)
func (c *Client) GetMultiRegionAccessPoint(ctx context.Context, params *GetMultiRegionAccessPointInput, optFns ...func(*Options)) (*GetMultiRegionAccessPointOutput, error) {
	if params == nil {
		params = &GetMultiRegionAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMultiRegionAccessPoint", params, optFns, c.addOperationGetMultiRegionAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMultiRegionAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMultiRegionAccessPointInput struct {

	// The Amazon Web Services account ID for the owner of the Multi-Region Access
	// Point.
	//
	// This member is required.
	AccountId *string

	// The name of the Multi-Region Access Point whose configuration information you
	// want to receive. The name of the Multi-Region Access Point is different from the
	// alias. For more information about the distinction between the name and the alias
	// of an Multi-Region Access Point, see Managing Multi-Region Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/CreatingMultiRegionAccessPoints.html#multi-region-access-point-naming)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetMultiRegionAccessPointOutput struct {

	// A container element containing the details of the requested Multi-Region Access
	// Point.
	AccessPoint *types.MultiRegionAccessPointReport

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMultiRegionAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetMultiRegionAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetMultiRegionAccessPoint{}, middleware.After)
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetMultiRegionAccessPointMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMultiRegionAccessPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMultiRegionAccessPoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addGetMultiRegionAccessPointUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGetMultiRegionAccessPointMiddleware struct {
}

func (*endpointPrefix_opGetMultiRegionAccessPointMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMultiRegionAccessPointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetMultiRegionAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetMultiRegionAccessPointMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetMultiRegionAccessPointMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetMultiRegionAccessPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetMultiRegionAccessPoint",
	}
}

func copyGetMultiRegionAccessPointInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetMultiRegionAccessPointInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetMultiRegionAccessPointInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillGetMultiRegionAccessPointAccountID(input interface{}, v string) error {
	in := input.(*GetMultiRegionAccessPointInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetMultiRegionAccessPointUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetMultiRegionAccessPointInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
