// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a 2048-bit RSA key pair and issues an X.509 certificate using the issued
// public key. You can also call CreateKeysAndCertificate over MQTT from a device,
// for more information, see Provisioning MQTT API
// (https://docs.aws.amazon.com/iot/latest/developerguide/provision-wo-cert.html#provision-mqtt-api).
// Note This is the only time IoT issues the private key for this certificate, so
// it is important to keep it in a secure location. Requires permission to access
// the CreateKeysAndCertificate
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CreateKeysAndCertificate(ctx context.Context, params *CreateKeysAndCertificateInput, optFns ...func(*Options)) (*CreateKeysAndCertificateOutput, error) {
	if params == nil {
		params = &CreateKeysAndCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKeysAndCertificate", params, optFns, c.addOperationCreateKeysAndCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeysAndCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreateKeysAndCertificate operation. Requires permission to
// access the CreateKeysAndCertificateRequest
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
type CreateKeysAndCertificateInput struct {

	// Specifies whether the certificate is active.
	SetAsActive bool

	noSmithyDocumentSerde
}

// The output of the CreateKeysAndCertificate operation.
type CreateKeysAndCertificateOutput struct {

	// The ARN of the certificate.
	CertificateArn *string

	// The ID of the certificate. IoT issues a default subject name for the certificate
	// (for example, IoT Certificate).
	CertificateId *string

	// The certificate data, in PEM format.
	CertificatePem *string

	// The generated key pair.
	KeyPair *types.KeyPair

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKeysAndCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKeysAndCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKeysAndCertificate{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKeysAndCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateKeysAndCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateKeysAndCertificate",
	}
}
