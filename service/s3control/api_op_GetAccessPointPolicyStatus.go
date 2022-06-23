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

// Indicates whether the specified access point currently has a policy that allows
// public access. For more information about public access through access points,
// see Managing Data Access with Amazon S3 access points
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html) in
// the Amazon S3 User Guide.
func (c *Client) GetAccessPointPolicyStatus(ctx context.Context, params *GetAccessPointPolicyStatusInput, optFns ...func(*Options)) (*GetAccessPointPolicyStatusOutput, error) {
	if params == nil {
		params = &GetAccessPointPolicyStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessPointPolicyStatus", params, optFns, c.addOperationGetAccessPointPolicyStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessPointPolicyStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessPointPolicyStatusInput struct {

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point whose policy status you want to retrieve.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetAccessPointPolicyStatusOutput struct {

	// Indicates the current policy status of the specified access point.
	PolicyStatus *types.PolicyStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessPointPolicyStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessPointPolicyStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessPointPolicyStatus{}, middleware.After)
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
	if err = addEndpointPrefix_opGetAccessPointPolicyStatusMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAccessPointPolicyStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessPointPolicyStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addGetAccessPointPolicyStatusUpdateEndpoint(stack, options); err != nil {
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

type endpointPrefix_opGetAccessPointPolicyStatusMiddleware struct {
}

func (*endpointPrefix_opGetAccessPointPolicyStatusMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAccessPointPolicyStatusMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetAccessPointPolicyStatusInput)
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
func addEndpointPrefix_opGetAccessPointPolicyStatusMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetAccessPointPolicyStatusMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessPointPolicyStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetAccessPointPolicyStatus",
	}
}

func copyGetAccessPointPolicyStatusInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetAccessPointPolicyStatusInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetAccessPointPolicyStatusInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillGetAccessPointPolicyStatusAccountID(input interface{}, v string) error {
	in := input.(*GetAccessPointPolicyStatusInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetAccessPointPolicyStatusUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetAccessPointPolicyStatusInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
