// Code generated by smithy-go-codegen DO NOT EDIT.

package kafkaconnect

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/kafkaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns summary information about the connector.
func (c *Client) DescribeConnector(ctx context.Context, params *DescribeConnectorInput, optFns ...func(*Options)) (*DescribeConnectorOutput, error) {
	if params == nil {
		params = &DescribeConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnector", params, optFns, c.addOperationDescribeConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectorInput struct {

	// The Amazon Resource Name (ARN) of the connector that you want to describe.
	//
	// This member is required.
	ConnectorArn *string

	noSmithyDocumentSerde
}

type DescribeConnectorOutput struct {

	// Information about the capacity of the connector, whether it is auto scaled or
	// provisioned.
	Capacity *types.CapacityDescription

	// The Amazon Resource Name (ARN) of the connector.
	ConnectorArn *string

	// A map of keys to values that represent the configuration for the connector.
	ConnectorConfiguration map[string]string

	// A summary description of the connector.
	ConnectorDescription *string

	// The name of the connector.
	ConnectorName *string

	// The state of the connector.
	ConnectorState types.ConnectorState

	// The time the connector was created.
	CreationTime *time.Time

	// The current version of the connector.
	CurrentVersion *string

	// The Apache Kafka cluster that the connector is connected to.
	KafkaCluster *types.KafkaClusterDescription

	// The type of client authentication used to connect to the Apache Kafka cluster.
	// The value is NONE when no client authentication is used.
	KafkaClusterClientAuthentication *types.KafkaClusterClientAuthenticationDescription

	// Details of encryption in transit to the Apache Kafka cluster.
	KafkaClusterEncryptionInTransit *types.KafkaClusterEncryptionInTransitDescription

	// The version of Kafka Connect. It has to be compatible with both the Apache Kafka
	// cluster's version and the plugins.
	KafkaConnectVersion *string

	// Details about delivering logs to Amazon CloudWatch Logs.
	LogDelivery *types.LogDeliveryDescription

	// Specifies which plugins were used for this connector.
	Plugins []types.PluginDescription

	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access
	// Amazon Web Services resources.
	ServiceExecutionRoleArn *string

	// Details about the state of a connector.
	StateDescription *types.StateDescription

	// Specifies which worker configuration was used for the connector.
	WorkerConfiguration *types.WorkerConfigurationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConnector{}, middleware.After)
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
	if err = addOpDescribeConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafkaconnect",
		OperationName: "DescribeConnector",
	}
}
