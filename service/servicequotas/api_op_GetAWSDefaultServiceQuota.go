// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the default value for the specified quota. The default value does not
// reflect any quota increases.
func (c *Client) GetAWSDefaultServiceQuota(ctx context.Context, params *GetAWSDefaultServiceQuotaInput, optFns ...func(*Options)) (*GetAWSDefaultServiceQuotaOutput, error) {
	if params == nil {
		params = &GetAWSDefaultServiceQuotaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAWSDefaultServiceQuota", params, optFns, c.addOperationGetAWSDefaultServiceQuotaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAWSDefaultServiceQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAWSDefaultServiceQuotaInput struct {

	// The quota identifier.
	//
	// This member is required.
	QuotaCode *string

	// The service identifier.
	//
	// This member is required.
	ServiceCode *string

	noSmithyDocumentSerde
}

type GetAWSDefaultServiceQuotaOutput struct {

	// Information about the quota.
	Quota *types.ServiceQuota

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAWSDefaultServiceQuotaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAWSDefaultServiceQuota{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAWSDefaultServiceQuota{}, middleware.After)
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
	if err = addOpGetAWSDefaultServiceQuotaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAWSDefaultServiceQuota(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAWSDefaultServiceQuota(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "GetAWSDefaultServiceQuota",
	}
}
