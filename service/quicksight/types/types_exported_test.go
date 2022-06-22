// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/service/quicksight/types"
)

func ExampleDataSourceParameters_outputUsage() {
	var union types.DataSourceParameters
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DataSourceParametersMemberAmazonElasticsearchParameters:
		_ = v.Value // Value is types.AmazonElasticsearchParameters

	case *types.DataSourceParametersMemberAmazonOpenSearchParameters:
		_ = v.Value // Value is types.AmazonOpenSearchParameters

	case *types.DataSourceParametersMemberAthenaParameters:
		_ = v.Value // Value is types.AthenaParameters

	case *types.DataSourceParametersMemberAuroraParameters:
		_ = v.Value // Value is types.AuroraParameters

	case *types.DataSourceParametersMemberAuroraPostgreSqlParameters:
		_ = v.Value // Value is types.AuroraPostgreSqlParameters

	case *types.DataSourceParametersMemberAwsIotAnalyticsParameters:
		_ = v.Value // Value is types.AwsIotAnalyticsParameters

	case *types.DataSourceParametersMemberExasolParameters:
		_ = v.Value // Value is types.ExasolParameters

	case *types.DataSourceParametersMemberJiraParameters:
		_ = v.Value // Value is types.JiraParameters

	case *types.DataSourceParametersMemberMariaDbParameters:
		_ = v.Value // Value is types.MariaDbParameters

	case *types.DataSourceParametersMemberMySqlParameters:
		_ = v.Value // Value is types.MySqlParameters

	case *types.DataSourceParametersMemberOracleParameters:
		_ = v.Value // Value is types.OracleParameters

	case *types.DataSourceParametersMemberPostgreSqlParameters:
		_ = v.Value // Value is types.PostgreSqlParameters

	case *types.DataSourceParametersMemberPrestoParameters:
		_ = v.Value // Value is types.PrestoParameters

	case *types.DataSourceParametersMemberRdsParameters:
		_ = v.Value // Value is types.RdsParameters

	case *types.DataSourceParametersMemberRedshiftParameters:
		_ = v.Value // Value is types.RedshiftParameters

	case *types.DataSourceParametersMemberS3Parameters:
		_ = v.Value // Value is types.S3Parameters

	case *types.DataSourceParametersMemberServiceNowParameters:
		_ = v.Value // Value is types.ServiceNowParameters

	case *types.DataSourceParametersMemberSnowflakeParameters:
		_ = v.Value // Value is types.SnowflakeParameters

	case *types.DataSourceParametersMemberSparkParameters:
		_ = v.Value // Value is types.SparkParameters

	case *types.DataSourceParametersMemberSqlServerParameters:
		_ = v.Value // Value is types.SqlServerParameters

	case *types.DataSourceParametersMemberTeradataParameters:
		_ = v.Value // Value is types.TeradataParameters

	case *types.DataSourceParametersMemberTwitterParameters:
		_ = v.Value // Value is types.TwitterParameters

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AmazonElasticsearchParameters
var _ *types.MariaDbParameters
var _ *types.RdsParameters
var _ *types.AmazonOpenSearchParameters
var _ *types.RedshiftParameters
var _ *types.OracleParameters
var _ *types.JiraParameters
var _ *types.TeradataParameters
var _ *types.MySqlParameters
var _ *types.ExasolParameters
var _ *types.SnowflakeParameters
var _ *types.SqlServerParameters
var _ *types.PostgreSqlParameters
var _ *types.ServiceNowParameters
var _ *types.PrestoParameters
var _ *types.AuroraParameters
var _ *types.S3Parameters
var _ *types.TwitterParameters
var _ *types.AuroraPostgreSqlParameters
var _ *types.AwsIotAnalyticsParameters
var _ *types.AthenaParameters
var _ *types.SparkParameters

func ExamplePhysicalTable_outputUsage() {
	var union types.PhysicalTable
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PhysicalTableMemberCustomSql:
		_ = v.Value // Value is types.CustomSql

	case *types.PhysicalTableMemberRelationalTable:
		_ = v.Value // Value is types.RelationalTable

	case *types.PhysicalTableMemberS3Source:
		_ = v.Value // Value is types.S3Source

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3Source
var _ *types.CustomSql
var _ *types.RelationalTable

func ExampleTransformOperation_outputUsage() {
	var union types.TransformOperation
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TransformOperationMemberCastColumnTypeOperation:
		_ = v.Value // Value is types.CastColumnTypeOperation

	case *types.TransformOperationMemberCreateColumnsOperation:
		_ = v.Value // Value is types.CreateColumnsOperation

	case *types.TransformOperationMemberFilterOperation:
		_ = v.Value // Value is types.FilterOperation

	case *types.TransformOperationMemberProjectOperation:
		_ = v.Value // Value is types.ProjectOperation

	case *types.TransformOperationMemberRenameColumnOperation:
		_ = v.Value // Value is types.RenameColumnOperation

	case *types.TransformOperationMemberTagColumnOperation:
		_ = v.Value // Value is types.TagColumnOperation

	case *types.TransformOperationMemberUntagColumnOperation:
		_ = v.Value // Value is types.UntagColumnOperation

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.CreateColumnsOperation
var _ *types.FilterOperation
var _ *types.ProjectOperation
var _ *types.CastColumnTypeOperation
var _ *types.TagColumnOperation
var _ *types.RenameColumnOperation
var _ *types.UntagColumnOperation
