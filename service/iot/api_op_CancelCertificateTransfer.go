// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a pending transfer for the specified certificate. Note Only the transfer
// source account can use this operation to cancel a transfer. (Transfer
// destinations can use RejectCertificateTransfer instead.) After transfer, IoT
// returns the certificate to the source account in the INACTIVE state. After the
// destination account has accepted the transfer, the transfer cannot be cancelled.
// After a certificate transfer is cancelled, the status of the certificate changes
// from PENDING_TRANSFER to INACTIVE. Requires permission to access the
// CancelCertificateTransfer
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CancelCertificateTransfer(ctx context.Context, params *CancelCertificateTransferInput, optFns ...func(*Options)) (*CancelCertificateTransferOutput, error) {
	if params == nil {
		params = &CancelCertificateTransferInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelCertificateTransfer", params, optFns, c.addOperationCancelCertificateTransferMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelCertificateTransferOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CancelCertificateTransfer operation.
type CancelCertificateTransferInput struct {

	// The ID of the certificate. (The last part of the certificate ARN contains the
	// certificate ID.)
	//
	// This member is required.
	CertificateId *string

	noSmithyDocumentSerde
}

type CancelCertificateTransferOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelCertificateTransferMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelCertificateTransfer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelCertificateTransfer{}, middleware.After)
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
	if err = addOpCancelCertificateTransferValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelCertificateTransfer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelCertificateTransfer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CancelCertificateTransfer",
	}
}
