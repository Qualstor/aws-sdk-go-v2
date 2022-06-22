// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disconnects the custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// from its associated CloudHSM cluster. While a custom key store is disconnected,
// you can manage the custom key store and its KMS keys, but you cannot create or
// use KMS keys in the custom key store. You can reconnect the custom key store at
// any time. While a custom key store is disconnected, all attempts to create KMS
// keys in the custom key store or to use existing KMS keys in cryptographic
// operations
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
// will fail. This action can prevent users from storing and accessing sensitive
// data. To find the connection state of a custom key store, use the
// DescribeCustomKeyStores operation. To reconnect a custom key store, use the
// ConnectCustomKeyStore operation. If the operation succeeds, it returns a JSON
// object with no properties. This operation is part of the Custom Key Store
// feature
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in KMS, which combines the convenience and extensive integration of KMS
// with the isolation and control of a single-tenant key store. Cross-account use:
// No. You cannot perform this operation on a custom key store in a different
// Amazon Web Services account. Required permissions: kms:DisconnectCustomKeyStore
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (IAM policy) Related operations:
//
// * ConnectCustomKeyStore
//
// *
// CreateCustomKeyStore
//
// * DeleteCustomKeyStore
//
// * DescribeCustomKeyStores
//
// *
// UpdateCustomKeyStore
func (c *Client) DisconnectCustomKeyStore(ctx context.Context, params *DisconnectCustomKeyStoreInput, optFns ...func(*Options)) (*DisconnectCustomKeyStoreOutput, error) {
	if params == nil {
		params = &DisconnectCustomKeyStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisconnectCustomKeyStore", params, optFns, c.addOperationDisconnectCustomKeyStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisconnectCustomKeyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisconnectCustomKeyStoreInput struct {

	// Enter the ID of the custom key store you want to disconnect. To find the ID of a
	// custom key store, use the DescribeCustomKeyStores operation.
	//
	// This member is required.
	CustomKeyStoreId *string

	noSmithyDocumentSerde
}

type DisconnectCustomKeyStoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisconnectCustomKeyStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisconnectCustomKeyStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisconnectCustomKeyStore{}, middleware.After)
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
	if err = addOpDisconnectCustomKeyStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisconnectCustomKeyStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisconnectCustomKeyStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "DisconnectCustomKeyStore",
	}
}
