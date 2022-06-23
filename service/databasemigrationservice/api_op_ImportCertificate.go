// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads the specified certificate.
func (c *Client) ImportCertificate(ctx context.Context, params *ImportCertificateInput, optFns ...func(*Options)) (*ImportCertificateOutput, error) {
	if params == nil {
		params = &ImportCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportCertificate", params, optFns, c.addOperationImportCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportCertificateInput struct {

	// A customer-assigned name for the certificate. Identifiers must begin with a
	// letter and must contain only ASCII letters, digits, and hyphens. They can't end
	// with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	CertificateIdentifier *string

	// The contents of a .pem file, which contains an X.509 certificate.
	CertificatePem *string

	// The location of an imported Oracle Wallet certificate for use with SSL. Provide
	// the name of a .sso file using the fileb:// prefix. You can't provide the
	// certificate inline. Example: filebase64("${path.root}/rds-ca-2019-root.sso")
	CertificateWallet []byte

	// The tags associated with the certificate.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ImportCertificateOutput struct {

	// The certificate to be uploaded.
	Certificate *types.Certificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportCertificate{}, middleware.After)
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
	if err = addOpImportCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "ImportCertificate",
	}
}
