// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/service/iot/types"
)

func ExampleAssetPropertyVariant_outputUsage() {
	var union types.AssetPropertyVariant
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AssetPropertyVariantMemberBooleanValue:
		_ = v.Value // Value is string

	case *types.AssetPropertyVariantMemberDoubleValue:
		_ = v.Value // Value is string

	case *types.AssetPropertyVariantMemberIntegerValue:
		_ = v.Value // Value is string

	case *types.AssetPropertyVariantMemberStringValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string
var _ *string
var _ *string
