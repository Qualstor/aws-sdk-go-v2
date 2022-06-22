// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a unique asymmetric data key pair for use outside of KMS. This operation
// returns a plaintext public key, a plaintext private key, and a copy of the
// private key that is encrypted under the symmetric encryption KMS key you
// specify. You can use the data key pair to perform asymmetric cryptography and
// implement digital signatures outside of KMS. The bytes in the keys are random;
// they not related to the caller or to the KMS key that is used to encrypt the
// private key. You can use the public key that GenerateDataKeyPair returns to
// encrypt data or verify a signature outside of KMS. Then, store the encrypted
// private key with the data. When you are ready to decrypt data or sign a message,
// you can use the Decrypt operation to decrypt the encrypted private key. To
// generate a data key pair, you must specify a symmetric encryption KMS key to
// encrypt the private key in a data key pair. You cannot use an asymmetric KMS key
// or a KMS key in a custom key store. To get the type and origin of your KMS key,
// use the DescribeKey operation. Use the KeyPairSpec parameter to choose an RSA or
// Elliptic Curve (ECC) data key pair. KMS recommends that your use ECC key pairs
// for signing, and use RSA key pairs for either encryption or signing, but not
// both. However, KMS cannot enforce any restrictions on the use of data key pairs
// outside of KMS. If you are using the data key pair to encrypt data, or for any
// operation where you don't immediately need a private key, consider using the
// GenerateDataKeyPairWithoutPlaintext operation.
// GenerateDataKeyPairWithoutPlaintext returns a plaintext public key and an
// encrypted private key, but omits the plaintext private key that you need only to
// decrypt ciphertext or sign a message. Later, when you need to decrypt the data
// or sign a message, use the Decrypt operation to decrypt the encrypted private
// key in the data key pair. GenerateDataKeyPair returns a unique data key pair for
// each request. The bytes in the keys are random; they are not related to the
// caller or the KMS key that is used to encrypt the private key. The public key is
// a DER-encoded X.509 SubjectPublicKeyInfo, as specified in RFC 5280
// (https://tools.ietf.org/html/rfc5280). The private key is a DER-encoded PKCS8
// PrivateKeyInfo, as specified in RFC 5958 (https://tools.ietf.org/html/rfc5958).
// You can use an optional encryption context to add additional security to the
// encryption operation. If you specify an EncryptionContext, you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the
// encrypted data key. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException. For more information, see Encryption Context
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// in the Key Management Service Developer Guide. The KMS key that you use for this
// operation must be in a compatible key state. For details, see Key states of KMS
// keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in
// the Key Management Service Developer Guide. Cross-account use: Yes. To perform
// this operation with a KMS key in a different Amazon Web Services account,
// specify the key ARN or alias ARN in the value of the KeyId parameter. Required
// permissions: kms:GenerateDataKeyPair
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//
// * Decrypt
//
// * Encrypt
//
// * GenerateDataKey
//
// *
// GenerateDataKeyPairWithoutPlaintext
//
// * GenerateDataKeyWithoutPlaintext
func (c *Client) GenerateDataKeyPair(ctx context.Context, params *GenerateDataKeyPairInput, optFns ...func(*Options)) (*GenerateDataKeyPairOutput, error) {
	if params == nil {
		params = &GenerateDataKeyPairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateDataKeyPair", params, optFns, c.addOperationGenerateDataKeyPairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateDataKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateDataKeyPairInput struct {

	// Specifies the symmetric encryption KMS key that encrypts the private key in the
	// data key pair. You cannot specify an asymmetric KMS key or a KMS key in a custom
	// key store. To get the type and origin of your KMS key, use the DescribeKey
	// operation. To specify a KMS key, use its key ID, key ARN, alias name, or alias
	// ARN. When using an alias name, prefix it with "alias/". To specify a KMS key in
	// a different Amazon Web Services account, you must use the key ARN or alias ARN.
	// For example:
	//
	// * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// *
	// Alias name: alias/ExampleAlias
	//
	// * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key
	// ARN for a KMS key, use ListKeys or DescribeKey. To get the alias name and alias
	// ARN, use ListAliases.
	//
	// This member is required.
	KeyId *string

	// Determines the type of data key pair that is generated. The KMS rule that
	// restricts the use of asymmetric RSA KMS keys to encrypt and decrypt or to sign
	// and verify (but not both), and the rule that permits you to use ECC KMS keys
	// only to sign and verify, are not effective on data key pairs, which are used
	// outside of KMS.
	//
	// This member is required.
	KeyPairSpec types.DataKeyPairSpec

	// Specifies the encryption context that will be used when encrypting the private
	// key in the data key pair. An encryption context is a collection of non-secret
	// key-value pairs that represent additional authenticated data. When you use an
	// encryption context to encrypt data, you must specify the same (an exact
	// case-sensitive match) encryption context to decrypt the data. An encryption
	// context is supported only on operations with symmetric encryption KMS keys. On
	// operations with symmetric encryption KMS keys, an encryption context is
	// optional, but it is strongly recommended. For more information, see Encryption
	// context
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the Key Management Service Developer Guide.
	EncryptionContext map[string]string

	// A list of grant tokens. Use a grant token when your permission to call this
	// operation comes from a new grant that has not yet achieved eventual consistency.
	// For more information, see Grant token
	// (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token
	// (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantTokens []string

	noSmithyDocumentSerde
}

type GenerateDataKeyPairOutput struct {

	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the KMS key that encrypted the private key.
	KeyId *string

	// The type of data key pair that was generated.
	KeyPairSpec types.DataKeyPairSpec

	// The encrypted copy of the private key. When you use the HTTP API or the Amazon
	// Web Services CLI, the value is Base64-encoded. Otherwise, it is not
	// Base64-encoded.
	PrivateKeyCiphertextBlob []byte

	// The plaintext copy of the private key. When you use the HTTP API or the Amazon
	// Web Services CLI, the value is Base64-encoded. Otherwise, it is not
	// Base64-encoded.
	PrivateKeyPlaintext []byte

	// The public key (in plaintext). When you use the HTTP API or the Amazon Web
	// Services CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	PublicKey []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateDataKeyPairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateDataKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateDataKeyPair{}, middleware.After)
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
	if err = addOpGenerateDataKeyPairValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateDataKeyPair(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateDataKeyPair(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "GenerateDataKeyPair",
	}
}
